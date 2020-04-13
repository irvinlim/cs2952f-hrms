import os

from nameko.rpc import rpc

from consistent_storage.base import BaseStorageBackend
from consistent_storage.sagas import SagasBackend

BACKEND = os.getenv('CONSISTENT_STORAGE_BACKEND', 'sagas')


class ConsistentStorageProxy:
    """
    Proxy for consistent storage based on the chosen backend.
    Exposes a simple key-value store interface via RPC.
    The storage semantics is that only one client should be able
    to put a given key; all others should fail subsequently until
    the key is removed from the consistent storage.
    """

    name = 'consistent_storage'

    def __init__(self):
        self.backend = BaseStorageBackend()

        if BACKEND == 'sagas':
            grpc_addr = os.getenv('SAGAS_GRPC_ADDR')
            if not grpc_addr:
                raise Exception('SAGAS_GRPC_ADDR not set')

            self.backend = SagasBackend(grpc_addr)

        elif BACKEND == 'bigchain':
            raise NotImplementedError()

        else:
            raise Exception(f'Invalid consistent storage backend: "{BACKEND}"')

    @rpc
    def healthy(self):
        return True

    @rpc
    def get(self, key: str) -> dict:
        """
        Fetches a key from consistent storage.

        Arguments:
            key {str} -- Key to fetch.

        Returns:
            dict --  Returns a dictionary as follows:

            {
                'exists': [bool],
                'value': [str or NoneType],
                'is_owner': [bool],
                'owner': [str],
            }

            If the key does not exist, the `is_owner` and `owner`
            metadata keys may not be present in the returned dictionary.
        """
        return self.backend.get(key)

    @rpc
    def put(self, key: str, value: str) -> dict:
        """
        Stores a key in the consistent storage in an atomic, linearizable fashion.
        Fails if the key is already present and stored by another owner.

        Arguments:
            key {str} -- Key to store.
            value {str} -- Value to store.

        Returns:
            dict --  Returns a dictionary as follows:

            {
                'ok': [bool],
                'owner': [str],
            }
        """
        return self.backend.put(key, value)

    @rpc
    def remove(self, key: str) -> dict:
        """
        Removes a key in the consistent storage in an atomic, linearizable fashion.
        Fails if the key is not in storage, or is owned by another owner.

        Arguments:
            key {str} -- Key to remove.

        Returns:
            dict --  Returns a dictionary as follows:

            {
                'removed': [bool],
                'error': [str],
            }
        """
        return self.backend.remove(key)
