import json

from nameko.rpc import RpcProxy, rpc
from nameko.exceptions import RemoteError

from lib import crypto, hasher
from lib.card import Card
from lib.local_storage import LocalStorage
from lib.hospital import get_hospital_name
from lib.medical_record import MedicalRecord


class PatientService:
    name = 'patient_service'
    consistent_storage = RpcProxy('consistent_storage')
    local_storage = LocalStorage()
    hospital_name = get_hospital_name()

    @rpc
    def healthy(self):
        """
        Returns True if service is healthy.
        """
        return True

    @rpc
    def register(self, patient_name, patient_id):
        """
        Registers the patient to this hospital and stores the patient's
        public key and patient ID in the consistent storage.
        Returns an error if the patient is already registered in another hospital.
        """
        print("hello")
        uid = patient_name + patient_id
        hash_uid = hasher.hash(uid)

        # Generate public keys for new patient.
        # Perhaps cache generated keys to prevent DDoS?
        pub_key, priv_key = crypto.generate_keys()

        # Check if hashed UID resides in consistent storage, otherwise store public key
        try:
            self.consistent_storage.put(hash_uid, crypto.b64encode(pub_key))
        except Exception as e:
            # TODO: Raise relevant exception
            raise e

        # Create patient card
        card = Card(patient_name, patient_id, uid, priv_key, self.hospital_name)

        # Store medical record in local storage
        record = MedicalRecord(self.hospital_name, card)
        self.local_storage.insert_item(uid, pub_key, record)

        return card
    
    @rpc
    def read(self, patient_uid):
        """
        Returns encrypted medical records for uid.
        Returns an error if the patient has not registered with a hospital.
        """
        med_records = []

        # Obtain the hashed UID.
        hash_uid = hasher.hash(patient_uid)

        # Get the public key from consistent storage if it exists.
        try:
            pub_key = self.consistent_storage.get(hash_uid)
            # Obtain the encrypted medical records.
            med_records = LocalStorage.get_items(uid, pub_key)
        except Exception as e:
            # TODO: Raise relevant exception
            raise e
        
        return med_records