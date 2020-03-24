import os

from redis import Redis

from lib import crypto


class LocalStorage():
    def __init__(self):
        redis_host = os.environ.get('REDIS_HOST', 'localhost')
        redis_port = int(os.environ.get('REDIS_PORT', '6379'))
        self.redis = Redis(host=redis_host, port=redis_port)

    def insert_item(self, uid, public_key, value):
        """
        Inserts a new item into a list in local storage with encrypted data.
        """

        key = crypto.encrypt(uid, public_key)
        value = crypto.encrypt(str(value), public_key)
        self.redis.lpush(key, value)

    def get_items(self, uid, public_key, offset=0, length=None):
        """
        Retrieves encrypted items from a list in local storage,
        optionally with offset and length parameters.
        """

        if length is None:
            length = -1

        key = crypto.encrypt(uid, public_key)
        return self.redis.lrange(key, offset, length)