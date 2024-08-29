# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/role.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/identity/v2/role.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xa9\x01\n\x11\x43reateRoleRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x35\n\trole_type\x18\x02 \x01(\x0e\x32\".spaceone.api.identity.v2.RoleType\x12\x13\n\x0bpermissions\x18\x03 \x03(\t\x12\x13\n\x0bpage_access\x18\x04 \x03(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x83\x01\n\x11UpdateRoleRequest\x12\x0f\n\x07role_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bpermissions\x18\x03 \x03(\t\x12\x13\n\x0bpage_access\x18\x04 \x03(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x1e\n\x0bRoleRequest\x12\x0f\n\x07role_id\x18\x01 \x01(\t\"\xe7\x02\n\x08RoleInfo\x12\x0f\n\x07role_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x35\n\trole_type\x18\x03 \x01(\x0e\x32\".spaceone.api.identity.v2.RoleType\x12\x13\n\x0bpermissions\x18\x04 \x03(\t\x12\x13\n\x0bpage_access\x18\x05 \x03(\t\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nis_managed\x18\x07 \x01(\x08\x12\x37\n\x05state\x18\x08 \x01(\x0e\x32(.spaceone.api.identity.v2.RoleInfo.State\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xd1\x01\n\rBasicRoleInfo\x12\x0f\n\x07role_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12<\n\x05state\x18\x03 \x01(\x0e\x32-.spaceone.api.identity.v2.BasicRoleInfo.State\x12\x35\n\trole_type\x18\x04 \x01(\x0e\x32\".spaceone.api.identity.v2.RoleType\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\x81\x02\n\x0fRoleSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07role_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x35\n\trole_type\x18\x04 \x01(\x0e\x32\".spaceone.api.identity.v2.RoleType\x12>\n\x05state\x18\x05 \x01(\x0e\x32/.spaceone.api.identity.v2.RoleSearchQuery.State\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"U\n\tRolesInfo\x12\x33\n\x07results\x18\x01 \x03(\x0b\x32\".spaceone.api.identity.v2.RoleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"_\n\x0e\x42\x61sicRolesInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.identity.v2.BasicRoleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"E\n\rRoleStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery*Q\n\x08RoleType\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x01\x12\x13\n\x0fWORKSPACE_OWNER\x10\x02\x12\x14\n\x10WORKSPACE_MEMBER\x10\x03\x32\xdd\x08\n\x04Role\x12~\n\x06\x63reate\x12+.spaceone.api.identity.v2.CreateRoleRequest\x1a\".spaceone.api.identity.v2.RoleInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/role/create:\x01*\x12~\n\x06update\x12+.spaceone.api.identity.v2.UpdateRoleRequest\x1a\".spaceone.api.identity.v2.RoleInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/role/update:\x01*\x12x\n\x06\x65nable\x12%.spaceone.api.identity.v2.RoleRequest\x1a\".spaceone.api.identity.v2.RoleInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/role/enable:\x01*\x12z\n\x07\x64isable\x12%.spaceone.api.identity.v2.RoleRequest\x1a\".spaceone.api.identity.v2.RoleInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v2/role/disable:\x01*\x12l\n\x06\x64\x65lete\x12%.spaceone.api.identity.v2.RoleRequest\x1a\x16.google.protobuf.Empty\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/role/delete:\x01*\x12r\n\x03get\x12%.spaceone.api.identity.v2.RoleRequest\x1a\".spaceone.api.identity.v2.RoleInfo\" \x82\xd3\xe4\x93\x02\x1a\"\x15/identity/v2/role/get:\x01*\x12y\n\x04list\x12).spaceone.api.identity.v2.RoleSearchQuery\x1a#.spaceone.api.identity.v2.RolesInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v2/role/list:\x01*\x12\x94\x01\n\x0flist_basic_role\x12).spaceone.api.identity.v2.RoleSearchQuery\x1a(.spaceone.api.identity.v2.BasicRolesInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/role/list-basic-role:\x01*\x12k\n\x04stat\x12\'.spaceone.api.identity.v2.RoleStatQuery\x1a\x17.google.protobuf.Struct\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v2/role/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.role_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_ROLE'].methods_by_name['create']._loaded_options = None
  _globals['_ROLE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/role/create:\001*'
  _globals['_ROLE'].methods_by_name['update']._loaded_options = None
  _globals['_ROLE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/role/update:\001*'
  _globals['_ROLE'].methods_by_name['enable']._loaded_options = None
  _globals['_ROLE'].methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/role/enable:\001*'
  _globals['_ROLE'].methods_by_name['disable']._loaded_options = None
  _globals['_ROLE'].methods_by_name['disable']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v2/role/disable:\001*'
  _globals['_ROLE'].methods_by_name['delete']._loaded_options = None
  _globals['_ROLE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/role/delete:\001*'
  _globals['_ROLE'].methods_by_name['get']._loaded_options = None
  _globals['_ROLE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\032\"\025/identity/v2/role/get:\001*'
  _globals['_ROLE'].methods_by_name['list']._loaded_options = None
  _globals['_ROLE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v2/role/list:\001*'
  _globals['_ROLE'].methods_by_name['list_basic_role']._loaded_options = None
  _globals['_ROLE'].methods_by_name['list_basic_role']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/role/list-basic-role:\001*'
  _globals['_ROLE'].methods_by_name['stat']._loaded_options = None
  _globals['_ROLE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v2/role/stat:\001*'
  _globals['_ROLETYPE']._serialized_start=1615
  _globals['_ROLETYPE']._serialized_end=1696
  _globals['_CREATEROLEREQUEST']._serialized_start=189
  _globals['_CREATEROLEREQUEST']._serialized_end=358
  _globals['_UPDATEROLEREQUEST']._serialized_start=361
  _globals['_UPDATEROLEREQUEST']._serialized_end=492
  _globals['_ROLEREQUEST']._serialized_start=494
  _globals['_ROLEREQUEST']._serialized_end=524
  _globals['_ROLEINFO']._serialized_start=527
  _globals['_ROLEINFO']._serialized_end=886
  _globals['_ROLEINFO_STATE']._serialized_start=842
  _globals['_ROLEINFO_STATE']._serialized_end=886
  _globals['_BASICROLEINFO']._serialized_start=889
  _globals['_BASICROLEINFO']._serialized_end=1098
  _globals['_BASICROLEINFO_STATE']._serialized_start=842
  _globals['_BASICROLEINFO_STATE']._serialized_end=886
  _globals['_ROLESEARCHQUERY']._serialized_start=1101
  _globals['_ROLESEARCHQUERY']._serialized_end=1358
  _globals['_ROLESEARCHQUERY_STATE']._serialized_start=842
  _globals['_ROLESEARCHQUERY_STATE']._serialized_end=886
  _globals['_ROLESINFO']._serialized_start=1360
  _globals['_ROLESINFO']._serialized_end=1445
  _globals['_BASICROLESINFO']._serialized_start=1447
  _globals['_BASICROLESINFO']._serialized_end=1542
  _globals['_ROLESTATQUERY']._serialized_start=1544
  _globals['_ROLESTATQUERY']._serialized_end=1613
  _globals['_ROLE']._serialized_start=1699
  _globals['_ROLE']._serialized_end=2816
# @@protoc_insertion_point(module_scope)
