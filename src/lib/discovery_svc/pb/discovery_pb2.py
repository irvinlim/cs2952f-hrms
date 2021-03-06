# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: discovery.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='discovery.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\tdiscovery',
  serialized_pb=b'\n\x0f\x64iscovery.proto\"\xa6\x01\n\x08Hospital\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x16\n\x0eregisteredTime\x18\x03 \x01(\x03\x12\x13\n\x0bgatewayAddr\x18\x04 \x01(\t\x12\x11\n\tpublicKey\x18\x05 \x01(\x0c\x12\x1d\n\x15\x63onsistentStorageAddr\x18\x06 \x01(\t\x12!\n\npublicKeys\x18\x07 \x03(\x0b\x32\r.DiscoveryKey\"\r\n\x0bInfoRequest\"T\n\x0cInfoResponse\x12\x16\n\x0eregisteredTime\x18\x01 \x01(\x03\x12\n\n\x02id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x12\n\nprivateKey\x18\x04 \x01(\x0c\"\r\n\x0bListRequest\",\n\x0cListResponse\x12\x1c\n\thospitals\x18\x01 \x03(\x0b\x32\t.Hospital\"K\n\x0c\x44iscoveryKey\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x0c\x12\x0e\n\x06public\x18\x03 \x01(\x08\x12\x0e\n\x06scheme\x18\x04 \x01(\t\"-\n\rGetKeyRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06public\x18\x02 \x01(\x08\";\n\x0eGetKeyResponse\x12\r\n\x05\x66ound\x18\x01 \x01(\x08\x12\x1a\n\x03key\x18\x02 \x01(\x0b\x32\r.DiscoveryKey\"+\n\rPutKeyRequest\x12\x1a\n\x03key\x18\x01 \x01(\x0b\x32\r.DiscoveryKey\"+\n\x0ePutKeyResponse\x12\n\n\x02ok\x18\x01 \x01(\x08\x12\r\n\x05\x65rror\x18\x02 \x01(\t2\xc7\x01\n\x11HospitalDiscovery\x12(\n\x07GetInfo\x12\x0c.InfoRequest\x1a\r.InfoResponse\"\x00\x12.\n\rListHospitals\x12\x0c.ListRequest\x1a\r.ListResponse\"\x00\x12+\n\x06GetKey\x12\x0e.GetKeyRequest\x1a\x0f.GetKeyResponse\"\x00\x12+\n\x06PutKey\x12\x0e.PutKeyRequest\x1a\x0f.PutKeyResponse\"\x00\x42\x0bZ\tdiscoveryb\x06proto3'
)




_HOSPITAL = _descriptor.Descriptor(
  name='Hospital',
  full_name='Hospital',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='Hospital.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='Hospital.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='registeredTime', full_name='Hospital.registeredTime', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gatewayAddr', full_name='Hospital.gatewayAddr', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='publicKey', full_name='Hospital.publicKey', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='consistentStorageAddr', full_name='Hospital.consistentStorageAddr', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='publicKeys', full_name='Hospital.publicKeys', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=186,
)


_INFOREQUEST = _descriptor.Descriptor(
  name='InfoRequest',
  full_name='InfoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=188,
  serialized_end=201,
)


_INFORESPONSE = _descriptor.Descriptor(
  name='InfoResponse',
  full_name='InfoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='registeredTime', full_name='InfoResponse.registeredTime', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id', full_name='InfoResponse.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='InfoResponse.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='privateKey', full_name='InfoResponse.privateKey', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=203,
  serialized_end=287,
)


_LISTREQUEST = _descriptor.Descriptor(
  name='ListRequest',
  full_name='ListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=289,
  serialized_end=302,
)


_LISTRESPONSE = _descriptor.Descriptor(
  name='ListResponse',
  full_name='ListResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hospitals', full_name='ListResponse.hospitals', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=304,
  serialized_end=348,
)


_DISCOVERYKEY = _descriptor.Descriptor(
  name='DiscoveryKey',
  full_name='DiscoveryKey',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='DiscoveryKey.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='DiscoveryKey.value', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='public', full_name='DiscoveryKey.public', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scheme', full_name='DiscoveryKey.scheme', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=350,
  serialized_end=425,
)


_GETKEYREQUEST = _descriptor.Descriptor(
  name='GetKeyRequest',
  full_name='GetKeyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='GetKeyRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='public', full_name='GetKeyRequest.public', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=427,
  serialized_end=472,
)


_GETKEYRESPONSE = _descriptor.Descriptor(
  name='GetKeyResponse',
  full_name='GetKeyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='found', full_name='GetKeyResponse.found', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key', full_name='GetKeyResponse.key', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=474,
  serialized_end=533,
)


_PUTKEYREQUEST = _descriptor.Descriptor(
  name='PutKeyRequest',
  full_name='PutKeyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='PutKeyRequest.key', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=535,
  serialized_end=578,
)


_PUTKEYRESPONSE = _descriptor.Descriptor(
  name='PutKeyResponse',
  full_name='PutKeyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ok', full_name='PutKeyResponse.ok', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='PutKeyResponse.error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=580,
  serialized_end=623,
)

_HOSPITAL.fields_by_name['publicKeys'].message_type = _DISCOVERYKEY
_LISTRESPONSE.fields_by_name['hospitals'].message_type = _HOSPITAL
_GETKEYRESPONSE.fields_by_name['key'].message_type = _DISCOVERYKEY
_PUTKEYREQUEST.fields_by_name['key'].message_type = _DISCOVERYKEY
DESCRIPTOR.message_types_by_name['Hospital'] = _HOSPITAL
DESCRIPTOR.message_types_by_name['InfoRequest'] = _INFOREQUEST
DESCRIPTOR.message_types_by_name['InfoResponse'] = _INFORESPONSE
DESCRIPTOR.message_types_by_name['ListRequest'] = _LISTREQUEST
DESCRIPTOR.message_types_by_name['ListResponse'] = _LISTRESPONSE
DESCRIPTOR.message_types_by_name['DiscoveryKey'] = _DISCOVERYKEY
DESCRIPTOR.message_types_by_name['GetKeyRequest'] = _GETKEYREQUEST
DESCRIPTOR.message_types_by_name['GetKeyResponse'] = _GETKEYRESPONSE
DESCRIPTOR.message_types_by_name['PutKeyRequest'] = _PUTKEYREQUEST
DESCRIPTOR.message_types_by_name['PutKeyResponse'] = _PUTKEYRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Hospital = _reflection.GeneratedProtocolMessageType('Hospital', (_message.Message,), {
  'DESCRIPTOR' : _HOSPITAL,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:Hospital)
  })
_sym_db.RegisterMessage(Hospital)

InfoRequest = _reflection.GeneratedProtocolMessageType('InfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _INFOREQUEST,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:InfoRequest)
  })
_sym_db.RegisterMessage(InfoRequest)

InfoResponse = _reflection.GeneratedProtocolMessageType('InfoResponse', (_message.Message,), {
  'DESCRIPTOR' : _INFORESPONSE,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:InfoResponse)
  })
_sym_db.RegisterMessage(InfoResponse)

ListRequest = _reflection.GeneratedProtocolMessageType('ListRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTREQUEST,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:ListRequest)
  })
_sym_db.RegisterMessage(ListRequest)

ListResponse = _reflection.GeneratedProtocolMessageType('ListResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTRESPONSE,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:ListResponse)
  })
_sym_db.RegisterMessage(ListResponse)

DiscoveryKey = _reflection.GeneratedProtocolMessageType('DiscoveryKey', (_message.Message,), {
  'DESCRIPTOR' : _DISCOVERYKEY,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:DiscoveryKey)
  })
_sym_db.RegisterMessage(DiscoveryKey)

GetKeyRequest = _reflection.GeneratedProtocolMessageType('GetKeyRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETKEYREQUEST,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:GetKeyRequest)
  })
_sym_db.RegisterMessage(GetKeyRequest)

GetKeyResponse = _reflection.GeneratedProtocolMessageType('GetKeyResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETKEYRESPONSE,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:GetKeyResponse)
  })
_sym_db.RegisterMessage(GetKeyResponse)

PutKeyRequest = _reflection.GeneratedProtocolMessageType('PutKeyRequest', (_message.Message,), {
  'DESCRIPTOR' : _PUTKEYREQUEST,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:PutKeyRequest)
  })
_sym_db.RegisterMessage(PutKeyRequest)

PutKeyResponse = _reflection.GeneratedProtocolMessageType('PutKeyResponse', (_message.Message,), {
  'DESCRIPTOR' : _PUTKEYRESPONSE,
  '__module__' : 'discovery_pb2'
  # @@protoc_insertion_point(class_scope:PutKeyResponse)
  })
_sym_db.RegisterMessage(PutKeyResponse)


DESCRIPTOR._options = None

_HOSPITALDISCOVERY = _descriptor.ServiceDescriptor(
  name='HospitalDiscovery',
  full_name='HospitalDiscovery',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=626,
  serialized_end=825,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetInfo',
    full_name='HospitalDiscovery.GetInfo',
    index=0,
    containing_service=None,
    input_type=_INFOREQUEST,
    output_type=_INFORESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListHospitals',
    full_name='HospitalDiscovery.ListHospitals',
    index=1,
    containing_service=None,
    input_type=_LISTREQUEST,
    output_type=_LISTRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetKey',
    full_name='HospitalDiscovery.GetKey',
    index=2,
    containing_service=None,
    input_type=_GETKEYREQUEST,
    output_type=_GETKEYRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PutKey',
    full_name='HospitalDiscovery.PutKey',
    index=3,
    containing_service=None,
    input_type=_PUTKEYREQUEST,
    output_type=_PUTKEYRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_HOSPITALDISCOVERY)

DESCRIPTOR.services_by_name['HospitalDiscovery'] = _HOSPITALDISCOVERY

# @@protoc_insertion_point(module_scope)
