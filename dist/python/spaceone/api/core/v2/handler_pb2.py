# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/core/v2/handler.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"spaceone/api/core/v2/handler.proto\x12\x14spaceone.api.core.v2\x1a\x1cgoogle/protobuf/struct.proto\"\xe6\x01\n\x14\x41uthorizationRequest\x12\r\n\x05scope\x18\x01 \x01(\t\x12H\n\nowner_type\x18\x02 \x01(\x0e\x32\x34.spaceone.api.core.v2.AuthorizationRequest.OwnerType\x12\x10\n\x08\x61udience\x18\x03 \x01(\t\x12\x10\n\x08token_id\x18\x04 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x06 \x01(\t\x12\x11\n\tdomain_id\x18\x07 \x01(\t\"(\n\tOwnerType\x12\x08\n\x04NONE\x10\x00\x12\x08\n\x04USER\x10\x01\x12\x07\n\x03\x41PP\x10\x02\"\xf2\x01\n\x15\x41uthorizationResponse\x12G\n\trole_type\x18\x01 \x01(\x0e\x32\x34.spaceone.api.core.v2.AuthorizationResponse.RoleType\x12\x15\n\ruser_projects\x18\x02 \x03(\t\"y\n\x08RoleType\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06SYSTEM\x10\x01\x12\x10\n\x0cSYSTEM_ADMIN\x10\x02\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x03\x12\x13\n\x0fWORKSPACE_OWNER\x10\x04\x12\x14\n\x10WORKSPACE_MEMBER\x10\x05\x12\x08\n\x04USER\x10\x06\"*\n\x15\x41uthenticationRequest\x12\x11\n\tdomain_id\x18\x01 \x01(\t\"?\n\x16\x41uthenticationResponse\x12\x11\n\tdomain_id\x18\x01 \x01(\t\x12\x12\n\npublic_key\x18\x02 \x01(\t\"y\n\x0c\x45ventRequest\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x10\n\x08resource\x18\x02 \x01(\t\x12\x0c\n\x04verb\x18\x03 \x01(\t\x12\x0e\n\x06status\x18\x04 \x01(\t\x12(\n\x07message\x18\x05 \x01(\x0b\x32\x17.google.protobuf.StructB;Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.core.v2.handler_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2'
  _globals['_AUTHORIZATIONREQUEST']._serialized_start=91
  _globals['_AUTHORIZATIONREQUEST']._serialized_end=321
  _globals['_AUTHORIZATIONREQUEST_OWNERTYPE']._serialized_start=281
  _globals['_AUTHORIZATIONREQUEST_OWNERTYPE']._serialized_end=321
  _globals['_AUTHORIZATIONRESPONSE']._serialized_start=324
  _globals['_AUTHORIZATIONRESPONSE']._serialized_end=566
  _globals['_AUTHORIZATIONRESPONSE_ROLETYPE']._serialized_start=445
  _globals['_AUTHORIZATIONRESPONSE_ROLETYPE']._serialized_end=566
  _globals['_AUTHENTICATIONREQUEST']._serialized_start=568
  _globals['_AUTHENTICATIONREQUEST']._serialized_end=610
  _globals['_AUTHENTICATIONRESPONSE']._serialized_start=612
  _globals['_AUTHENTICATIONRESPONSE']._serialized_end=675
  _globals['_EVENTREQUEST']._serialized_start=677
  _globals['_EVENTREQUEST']._serialized_end=798
# @@protoc_insertion_point(module_scope)
