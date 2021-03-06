"""
Load test for consistent storage operations.
"""

import random
import string
import sys
import time

import grequests
import requests
from gevent.pool import Pool

from lib.discovery_svc import DiscoveryService
from loadtesting_py.card_digester import PatientCardDigester

PATIENT_CARD_DIGESTER = PatientCardDigester()


def random_key():
    # Returns a random 32 byte key.
    # This is the same length as a SHA256 hash.
    return ''.join(random.choice(string.ascii_lowercase) for _ in range(32))


def request(method, addr, data):
    try:
        res = requests.request(method, f'http://{addr}/', json=data)
        res.raise_for_status()
        return res
    except requests.RequestException as e:
        print('[!] Exception in requests (gevent):', str(e))
        print(f"RESULT BODY {res.text}")
        return None
    except Exception as e:
        print('[!] Unknown exception in request() call:', str(e))
        print(f"RESULT BODY {res.text}")
        return None


def dispatch(method, get_data, addrs, executor):
    """
    Dispatch requests asynchronously and return the time taken, as well as the results.

    Arguments:
        method {str} -- Method name.
        get_data {Generator[Dict, None, None]} -- Generator that yields JSON data. Determines total number of requests.
                                                  The same data will be used for all ports per request.
        addrs {List[str]} -- List of addresses that the requests should dispatch to.
    """

    start_time = time.time()

    if executor == 'gevent':
        # Use arbitrary pool size of 1000
        pool = Pool(1000)

        # Spawn greenlets
        greenlets = [pool.spawn(request, method, addr, data) for data in get_data for addr in addrs]

        # Wait until complete
        pool.join()
        end_time = time.time()

        # Get responses
        responses = [g.value for g in greenlets]

    elif executor == 'grequests':
        # Create requests
        reqs = (grequests.request(method, f'http://{addr}/', json=data) for data in get_data for addr in addrs)

        # Map requests to responses
        responses = grequests.map(reqs)
        end_time = time.time()

    else:
        raise Exception(f'Unknown executor: {executor}')

    # Compute total duration as a sum of elapsed time for valid responses
    durations = [res.elapsed.total_seconds() for res in responses if res]
    res_data = [res.json() for res in responses if res]

    actual_duration = end_time - start_time
    return actual_duration, durations, res_data


def run(test, num_requests, hospitals, executor):
    addrs = [hospital['consistent_storage_addr'] for hospital in hospitals]
    ids = [hospital['id'] for hospital in hospitals]

    # Get patient cards from digester; these must be prepopulated beforehand.
    if len(PATIENT_CARD_DIGESTER.cards) < num_requests:
        return Exception("ERROR: Num requests > num patient cards")

    uuid_pubkey_pairs = PATIENT_CARD_DIGESTER.get_uuid_pubkey_pairs()[:num_requests]

    if test == "get":
        print('[!] NOTE: Test may fail if keys do not already exist.')

        # Get UUIDs in random order
        uuids = [uuid_pubkey_pair[0] for uuid_pubkey_pair in uuid_pubkey_pairs]
        random.shuffle(uuids)

        # Get keys, assuming they already exist
        data = ({'key': uuid} for uuid in uuids)
        actual_duration, durations, responses = dispatch('GET', data, addrs, executor)

        # Count successes
        num_success = len([1 for res in responses if res['exists']])
        return actual_duration, durations, num_success

    elif test == "put":
        print('[!] NOTE: Test may fail if keys already exist.')

        # Put the keys
        data = ({'key': uuid, 'value': pubkey} for uuid, pubkey in uuid_pubkey_pairs)
        actual_duration, durations, responses = dispatch('PUT', data, addrs, executor)

        # Count successes
        num_success = len([1 for res in responses if res['ok']])
        return actual_duration, durations, num_success

    elif test == "remove":
        print('[!] NOTE: Test may fail if keys do not already exist.')

        # Get UUIDs in random order
        uuids = [uuid_pubkey_pair[0] for uuid_pubkey_pair in uuid_pubkey_pairs]
        random.shuffle(uuids)

        # Remove all the keys
        data = ({'key': key} for key in uuids)
        actual_duration, durations, responses = dispatch('REMOVE', data, addrs, executor)

        # Count successes
        num_success = len([1 for res in responses if res['removed']])
        return actual_duration, durations, num_success

    elif test == "transfer":
        print('[!] NOTE: Test may fail if keys do not already exist.')

        # Get UUIDs
        uuids = [uuid_pubkey_pair[0] for uuid_pubkey_pair in uuid_pubkey_pairs]

        # Transfer them all to some random hospital
        owners = (random.choice(ids) for _ in range(num_requests))
        data = ({'key': key, 'dest': owner} for key, owner in zip(uuids, owners))
        actual_duration, durations, responses = dispatch('TRANSFER', data, addrs, executor)

        # Count successes
        num_success = len([1 for res in responses if res['transferred']])
        return actual_duration, durations, num_success

    else:
        raise Exception('Unknown test:', test)


def main():
    if len(sys.argv) < 3:
        print('Usage python storage_load_test.py <test> <num_requests> <executor>')
        sys.exit(1)

    test = sys.argv[1].strip().lower()
    num_requests = int(sys.argv[2])

    executor = 'gevent'
    if len(sys.argv) == 4 and sys.argv[3] == 'grequests':
        executor = 'grequests'

    print(f'Using {executor} executor.')

    # Find all hospitals via discovery service
    hospitals = DiscoveryService().list_hospitals()
    num_hospitals = len(hospitals)

    hospital_ids = [hospital['id'] for hospital in hospitals]
    print(f'Discovered {num_hospitals} hospitals: {hospital_ids}')

    # Dispatch requests
    actual_duration, durations, num_success = run(test, num_requests, hospitals, executor)

    # Number of valid responses
    num_ok = len(durations)
    elapsed_duration = sum(durations)

    print(f'Total duration: {actual_duration:.4f}s for {num_requests} requests each at {num_hospitals} hospitals')
    avg_time = actual_duration / num_requests
    print(f'Average time per request: {avg_time:.4f}s')
    print(f'QPS: {1/avg_time:.4f}')
    print()

    print(f'Number of successes: {num_success}')
    print(f'Number of valid responses: {num_ok}')

    if num_ok > 0:
        inflight_time = elapsed_duration / num_ok
        print(f'Average in-flight request time: {inflight_time:.6f}s')


if __name__ == "__main__":
    main()
