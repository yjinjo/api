# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/secret/v1/user_secret.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/secret/v1/user_secret.proto\x12\x1bspaceone.api.user_secret.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x88\x01\n\x17\x43reateUserSecretRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tschema_id\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"f\n\x17UpdateUserSecretRequest\x12\x16\n\x0euser_secret_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"+\n\x11UserSecretRequest\x12\x16\n\x0euser_secret_id\x18\x01 \x01(\t\"E\n\x18GetUserSecretDataRequest\x12\x16\n\x0euser_secret_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"o\n\x1bUpdateUserSecretDataRequest\x12\x16\n\x0euser_secret_id\x18\x01 \x01(\t\x12\x11\n\tschema_id\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x88\x01\n\x0fUserSecretQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x16\n\x0euser_secret_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tschema_id\x18\x04 \x01(\t\x12\x10\n\x08provider\x18\x05 \x01(\t\"\x80\x01\n\x12UserSecretDataInfo\x12\x11\n\tencrypted\x18\x01 \x01(\x08\x12\x30\n\x0f\x65ncrypt_options\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x8c\x02\n\x0eUserSecretInfo\x12\x16\n\x0euser_secret_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x11\n\tschema_id\x18\x03 \x01(\t\x12\x10\n\x08provider\x18\x04 \x01(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07user_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"d\n\x0fUserSecretsInfo\x12<\n\x07results\x18\x01 \x03(\x0b\x32+.spaceone.api.user_secret.v1.UserSecretInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"^\n\x13UserSecretStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t2\xd3\x08\n\nUserSecret\x12\x95\x01\n\x06\x63reate\x12\x34.spaceone.api.user_secret.v1.CreateUserSecretRequest\x1a+.spaceone.api.user_secret.v1.UserSecretInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/secret/v1/user-secret/create:\x01*\x12\x95\x01\n\x06update\x12\x34.spaceone.api.user_secret.v1.UpdateUserSecretRequest\x1a+.spaceone.api.user_secret.v1.UserSecretInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/secret/v1/user-secret/update:\x01*\x12z\n\x06\x64\x65lete\x12..spaceone.api.user_secret.v1.UserSecretRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/secret/v1/user-secret/delete:\x01*\x12\x8e\x01\n\x0bupdate_data\x12\x38.spaceone.api.user_secret.v1.UpdateUserSecretDataRequest\x1a\x16.google.protobuf.Empty\"-\x82\xd3\xe4\x93\x02\'\"\"/secret/v1/user-secret/update-data:\x01*\x12t\n\x08get_data\x12\x35.spaceone.api.user_secret.v1.GetUserSecretDataRequest\x1a/.spaceone.api.user_secret.v1.UserSecretDataInfo\"\x00\x12\x89\x01\n\x03get\x12..spaceone.api.user_secret.v1.UserSecretRequest\x1a+.spaceone.api.user_secret.v1.UserSecretInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/secret/v1/user-secret/get:\x01*\x12\x8a\x01\n\x04list\x12,.spaceone.api.user_secret.v1.UserSecretQuery\x1a,.spaceone.api.user_secret.v1.UserSecretsInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/secret/v1/user-secret/list:\x01*\x12y\n\x04stat\x12\x30.spaceone.api.user_secret.v1.UserSecretStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/secret/v1/user-secret/stat:\x01*BBZ@github.com/cloudforet-io/api/dist/go/spaceone/api/user_secret/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.secret.v1.user_secret_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z@github.com/cloudforet-io/api/dist/go/spaceone/api/user_secret/v1'
  _globals['_USERSECRET'].methods_by_name['create']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/secret/v1/user-secret/create:\001*'
  _globals['_USERSECRET'].methods_by_name['update']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\"\"\035/secret/v1/user-secret/update:\001*'
  _globals['_USERSECRET'].methods_by_name['delete']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/secret/v1/user-secret/delete:\001*'
  _globals['_USERSECRET'].methods_by_name['update_data']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['update_data']._serialized_options = b'\202\323\344\223\002\'\"\"/secret/v1/user-secret/update-data:\001*'
  _globals['_USERSECRET'].methods_by_name['get']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/secret/v1/user-secret/get:\001*'
  _globals['_USERSECRET'].methods_by_name['list']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/secret/v1/user-secret/list:\001*'
  _globals['_USERSECRET'].methods_by_name['stat']._loaded_options = None
  _globals['_USERSECRET'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/secret/v1/user-secret/stat:\001*'
  _globals['_CREATEUSERSECRETREQUEST']._serialized_start=197
  _globals['_CREATEUSERSECRETREQUEST']._serialized_end=333
  _globals['_UPDATEUSERSECRETREQUEST']._serialized_start=335
  _globals['_UPDATEUSERSECRETREQUEST']._serialized_end=437
  _globals['_USERSECRETREQUEST']._serialized_start=439
  _globals['_USERSECRETREQUEST']._serialized_end=482
  _globals['_GETUSERSECRETDATAREQUEST']._serialized_start=484
  _globals['_GETUSERSECRETDATAREQUEST']._serialized_end=553
  _globals['_UPDATEUSERSECRETDATAREQUEST']._serialized_start=555
  _globals['_UPDATEUSERSECRETDATAREQUEST']._serialized_end=666
  _globals['_USERSECRETQUERY']._serialized_start=669
  _globals['_USERSECRETQUERY']._serialized_end=805
  _globals['_USERSECRETDATAINFO']._serialized_start=808
  _globals['_USERSECRETDATAINFO']._serialized_end=936
  _globals['_USERSECRETINFO']._serialized_start=939
  _globals['_USERSECRETINFO']._serialized_end=1207
  _globals['_USERSECRETINFO_RESOURCEGROUP']._serialized_start=1127
  _globals['_USERSECRETINFO_RESOURCEGROUP']._serialized_end=1207
  _globals['_USERSECRETSINFO']._serialized_start=1209
  _globals['_USERSECRETSINFO']._serialized_end=1309
  _globals['_USERSECRETSTATQUERY']._serialized_start=1311
  _globals['_USERSECRETSTATQUERY']._serialized_end=1405
  _globals['_USERSECRET']._serialized_start=1408
  _globals['_USERSECRET']._serialized_end=2515
# @@protoc_insertion_point(module_scope)
