"""
Load test for consistent storage operations.
"""

import random
import string
import sys
import time

from gevent.pool import Pool
import requests

from lib.discovery_svc import DiscoveryService

with open('loadtesting_py/pub.key') as f:
    PUBLIC_KEY = f.read()


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
        return None
    except Exception as e:
        print('[!] Unknown exception in request() call:', str(e))
        return None


def dispatch(method, get_data, addrs):
    """
    Dispatch requests asynchronously and return the time taken, as well as the results.

    Arguments:
        method {str} -- Method name.
        get_data {Generator[Dict, None, None]} -- Generator that yields JSON data. Determines total number of requests.
                                                  The same data will be used for all ports per request.
        addrs {List[str]} -- List of addresses that the requests should dispatch to.
    """

    # Use arbitrary pool size of 100
    pool = Pool(1000)

    start_time = time.time()

    # Spawn greenlets
    greenlets = [pool.spawn(request, method, addr, data) for data in get_data for addr in addrs]

    # Wait until complete
    pool.join()

    end_time = time.time()
    actual_duration = end_time - start_time

    # Get responses
    responses = [g.value for g in greenlets]

    # Compute total duration as a sum of elapsed time for valid responses
    durations = [res.elapsed.total_seconds() for res in responses if res]
    res_data = [res.json() for res in responses if res]

    return actual_duration, durations, res_data


def run(test, num_requests, hospitals):
    addrs = [hospital['consistent_storage_addr'] for hospital in hospitals]
    ids = [hospital['id'] for hospital in hospitals]

    if test == "get":
        # Fetch random, non-existent keys
        data = ({'key': random_key()} for _ in range(num_requests))
        actual_duration, durations, responses = dispatch('GET', data, addrs)

        # Count successes
        num_success = len([1 for res in responses if res['exists']])
        return actual_duration, durations, num_success

    elif test == "get_existing":
        # Generate random keys
        keys = [random_key() for _ in range(num_requests)]

        # Put random keys
        data = ({'key': key, 'value': PUBLIC_KEY} for key in keys)
        dispatch('PUT', data, addrs)

        # Now measure getting successfully put keys
        data = ({'key': key} for key in keys)
        actual_duration, durations, responses = dispatch('GET', data, addrs)

        # Count successes
        num_success = len([1 for res in responses if res['exists']])
        return actual_duration, durations, num_success

    elif test == "put":
        # Put random keys
        data = ({'key': random_key(), 'value': PUBLIC_KEY} for _ in range(num_requests))
        actual_duration, durations, responses = dispatch('PUT', data, addrs)

        # Count successes
        num_success = len([1 for res in responses if res['ok']])
        return actual_duration, durations, num_success

    elif test == "remove":
        # Generate random keys
        keys = [random_key() for _ in range(num_requests)]

        # Put keys initially
        data = ({'key': key, 'value': PUBLIC_KEY} for key in keys)
        dispatch('PUT', data, addrs)

        # Now measure the removal
        data = ({'key': key} for key in keys)
        actual_duration, durations, responses = dispatch('REMOVE', data, addrs)

        # Count successes
        num_success = len([1 for res in responses if res['removed']])
        return actual_duration, durations, num_success

    elif test == "transfer":
        # Generate random keys
        keys = [random_key() for _ in range(num_requests)]

        # Put keys initially
        data = ({'key': key, 'value': PUBLIC_KEY} for key in keys)
        dispatch('PUT', data, addrs)

        # Now transfer them all to some random hospital
        owners = (random.choice(ids) for _ in range(num_requests))
        data = ({'key': key, 'owner': owner} for key, owner in zip(keys, owners))
        actual_duration, durations, responses = dispatch('TRANSFER', data, addrs)

        # Count successes
        num_success = len([1 for res in responses if res['transferred']])
        return actual_duration, durations, num_success

    else:
        raise Exception('Unknown test:', test)


def main():
    if len(sys.argv) != 3:
        print('Usage python storage_load_test.py <test> <num_requests>')
        sys.exit(1)

    test = sys.argv[1].strip().lower()
    num_requests = int(sys.argv[2])

    # Find all hospitals via discovery service
    hospitals = DiscoveryService().list_hospitals()
    num_hospitals = len(hospitals)

    hospital_ids = [hospital['id'] for hospital in hospitals]
    print(f'Discovered {num_hospitals} hospitals: {hospital_ids}')

    # Dispatch requests
    actual_duration, durations, num_success = run(test, num_requests, hospitals)

    # Number of valid responses
    num_ok = len(durations)
    elapsed_duration = sum(durations)

    print(f'Total duration: {actual_duration:.4f}s for {num_requests} requests each at {num_hospitals} hospitals')
    print(f'Number of successes: {num_success}')

    avg_time = elapsed_duration / num_ok
    print(f'Number of valid responses: {num_ok}')
    print(f'Total elapsed duration: {elapsed_duration}')
    print(f'Average time per request: {avg_time:.6f}s')
    print(f'QPS: {1/avg_time:.2f}')


if __name__ == "__main__":
    main()
