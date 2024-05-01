# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v1/user.proto
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
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/identity/v1/user.proto\x12\x18spaceone.api.identity.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"t\n\x03MFA\x12\x31\n\x05state\x18\x01 \x01(\x0e\x32\".spaceone.api.identity.v1.MFAState\x12\x10\n\x08mfa_type\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xb8\x02\n\x11\x43reateUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x35\n\tuser_type\x18\x05 \x01(\x0e\x32\".spaceone.api.identity.v1.UserType\x12\x36\n\x07\x62\x61\x63kend\x18\x06 \x01(\x0e\x32%.spaceone.api.identity.v1.UserBackend\x12\x10\n\x08language\x18\x07 \x01(\t\x12\x10\n\x08timezone\x18\x08 \x01(\t\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\n \x01(\t\x12\x16\n\x0ereset_password\x18\x0b \x01(\x08\"\xc9\x01\n\x11UpdateUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x10\n\x08language\x18\x07 \x01(\t\x12\x10\n\x08timezone\x18\x08 \x01(\t\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\n \x01(\t\x12\x16\n\x0ereset_password\x18\x0b \x01(\x08\"V\n\x12VerifyEmailRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\r\n\x05\x66orce\x18\x03 \x01(\x08\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"N\n\x13\x43onfirmEmailRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x13\n\x0bverify_code\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"~\n\x19SetRequiredActionsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12=\n\x07\x61\x63tions\x18\x02 \x03(\x0e\x32,.spaceone.api.identity.v1.UserRequiredAction\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"r\n\x10\x45nableMFARequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08mfa_type\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"F\n\x11\x44isableMFARequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\r\n\x05\x66orce\x18\x02 \x01(\x08\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"L\n\x11\x43onfirmMFARequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x13\n\x0bverify_code\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"1\n\x0bUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"B\n\x0eGetUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xf6\x01\n\tUserQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05state\x18\x04 \x01(\t\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12\x35\n\tuser_type\x18\x06 \x01(\x0e\x32\".spaceone.api.identity.v1.UserType\x12\x36\n\x07\x62\x61\x63kend\x18\x07 \x01(\x0e\x32%.spaceone.api.identity.v1.UserBackend\x12\x11\n\tdomain_id\x18\x0b \x01(\t\"\xb3\x04\n\x08UserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\x05state\x18\x03 \x01(\x0e\x32(.spaceone.api.identity.v1.UserInfo.State\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x35\n\tuser_type\x18\x05 \x01(\x0e\x32\".spaceone.api.identity.v1.UserType\x12\x36\n\x07\x62\x61\x63kend\x18\x06 \x01(\x0e\x32%.spaceone.api.identity.v1.UserBackend\x12\x10\n\x08language\x18\x07 \x01(\t\x12\x10\n\x08timezone\x18\x08 \x01(\t\x12\x46\n\x10required_actions\x18\t \x03(\x0e\x32,.spaceone.api.identity.v1.UserRequiredAction\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x18\n\x10last_accessed_at\x18\x0b \x01(\t\x12\x12\n\ncreated_at\x18\x0c \x01(\t\x12\x11\n\tdomain_id\x18\r \x01(\t\x12\x16\n\x0e\x65mail_verified\x18\x0e \x01(\x08\x12*\n\x03mfa\x18\x0f \x01(\x0b\x32\x1d.spaceone.api.identity.v1.MFA\"9\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x0b\n\x07PENDING\x10\x03\"U\n\tUsersInfo\x12\x33\n\x07results\x18\x01 \x03(\x0b\x32\".spaceone.api.identity.v1.UserInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"X\n\rUserStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"F\n\x0e\x46indUserSearch\x12\x11\n\x07user_id\x18\x01 \x01(\tH\x00\x12\x11\n\x07keyword\x18\x02 \x01(\tH\x00\x42\x0e\n\x0csearch_alias\"\\\n\rFindUserQuery\x12\x38\n\x06search\x18\x01 \x01(\x0b\x32(.spaceone.api.identity.v1.FindUserSearch\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"c\n\x0c\x46indUserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"]\n\rFindUsersInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.identity.v1.FindUserInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05*8\n\x0bUserBackend\x12\x10\n\x0cNONE_BACKEND\x10\x00\x12\t\n\x05LOCAL\x10\x01\x12\x0c\n\x08\x45XTERNAL\x10\x02*6\n\x08UserType\x12\x12\n\x0eNONE_USER_TYPE\x10\x00\x12\x08\n\x04USER\x10\x01\x12\x0c\n\x08\x41PI_USER\x10\x02*)\n\x12UserRequiredAction\x12\x13\n\x0fUPDATE_PASSWORD\x10\x00*9\n\x08MFAState\x12\x12\n\x0eNONE_MFA_STATE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x32\x89\x11\n\x04User\x12~\n\x06\x63reate\x12+.spaceone.api.identity.v1.CreateUserRequest\x1a\".spaceone.api.identity.v1.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v1/user/create:\x01*\x12~\n\x06update\x12+.spaceone.api.identity.v1.UpdateUserRequest\x1a\".spaceone.api.identity.v1.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v1/user/update:\x01*\x12\x7f\n\x0cverify_email\x12,.spaceone.api.identity.v1.VerifyEmailRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v1/user/verify-email:\x01*\x12\x8e\x01\n\rconfirm_email\x12-.spaceone.api.identity.v1.ConfirmEmailRequest\x1a\".spaceone.api.identity.v1.UserInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/identity/v1/user/confirm-email:\x01*\x12|\n\x0ereset_password\x12%.spaceone.api.identity.v1.UserRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /identity/v1/user/reset-password:\x01*\x12\xa2\x01\n\x14set_required_actions\x12\x33.spaceone.api.identity.v1.SetRequiredActionsRequest\x1a\".spaceone.api.identity.v1.UserInfo\"1\x82\xd3\xe4\x93\x02+\"&/identity/v1/user/set-required-actions:\x01*\x12\x85\x01\n\nenable_mfa\x12*.spaceone.api.identity.v1.EnableMFARequest\x1a\".spaceone.api.identity.v1.UserInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/identity/v1/user/enable-mfa:\x01*\x12\x88\x01\n\x0b\x64isable_mfa\x12+.spaceone.api.identity.v1.DisableMFARequest\x1a\".spaceone.api.identity.v1.UserInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v1/user/disable-mfa:\x01*\x12\x88\x01\n\x0b\x63onfirm_mfa\x12+.spaceone.api.identity.v1.ConfirmMFARequest\x1a\".spaceone.api.identity.v1.UserInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v1/user/confirm-mfa:\x01*\x12x\n\x06\x65nable\x12%.spaceone.api.identity.v1.UserRequest\x1a\".spaceone.api.identity.v1.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v1/user/enable:\x01*\x12z\n\x07\x64isable\x12%.spaceone.api.identity.v1.UserRequest\x1a\".spaceone.api.identity.v1.UserInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v1/user/disable:\x01*\x12l\n\x06\x64\x65lete\x12%.spaceone.api.identity.v1.UserRequest\x1a\x16.google.protobuf.Empty\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v1/user/delete:\x01*\x12u\n\x03get\x12(.spaceone.api.identity.v1.GetUserRequest\x1a\".spaceone.api.identity.v1.UserInfo\" \x82\xd3\xe4\x93\x02\x1a\"\x15/identity/v1/user/get:\x01*\x12s\n\x04list\x12#.spaceone.api.identity.v1.UserQuery\x1a#.spaceone.api.identity.v1.UsersInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v1/user/list:\x01*\x12k\n\x04stat\x12\'.spaceone.api.identity.v1.UserStatQuery\x1a\x17.google.protobuf.Struct\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v1/user/stat:\x01*\x12{\n\x04\x66ind\x12\'.spaceone.api.identity.v1.FindUserQuery\x1a\'.spaceone.api.identity.v1.FindUsersInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v1/user/find:\x01*\x12t\n\x04sync\x12%.spaceone.api.identity.v1.UserRequest\x1a\".spaceone.api.identity.v1.UserInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v1/user/sync:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v1.user_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1'
  _globals['_USER'].methods_by_name['create']._loaded_options = None
  _globals['_USER'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v1/user/create:\001*'
  _globals['_USER'].methods_by_name['update']._loaded_options = None
  _globals['_USER'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v1/user/update:\001*'
  _globals['_USER'].methods_by_name['verify_email']._loaded_options = None
  _globals['_USER'].methods_by_name['verify_email']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v1/user/verify-email:\001*'
  _globals['_USER'].methods_by_name['confirm_email']._loaded_options = None
  _globals['_USER'].methods_by_name['confirm_email']._serialized_options = b'\202\323\344\223\002$\"\037/identity/v1/user/confirm-email:\001*'
  _globals['_USER'].methods_by_name['reset_password']._loaded_options = None
  _globals['_USER'].methods_by_name['reset_password']._serialized_options = b'\202\323\344\223\002%\" /identity/v1/user/reset-password:\001*'
  _globals['_USER'].methods_by_name['set_required_actions']._loaded_options = None
  _globals['_USER'].methods_by_name['set_required_actions']._serialized_options = b'\202\323\344\223\002+\"&/identity/v1/user/set-required-actions:\001*'
  _globals['_USER'].methods_by_name['enable_mfa']._loaded_options = None
  _globals['_USER'].methods_by_name['enable_mfa']._serialized_options = b'\202\323\344\223\002!\"\034/identity/v1/user/enable-mfa:\001*'
  _globals['_USER'].methods_by_name['disable_mfa']._loaded_options = None
  _globals['_USER'].methods_by_name['disable_mfa']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v1/user/disable-mfa:\001*'
  _globals['_USER'].methods_by_name['confirm_mfa']._loaded_options = None
  _globals['_USER'].methods_by_name['confirm_mfa']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v1/user/confirm-mfa:\001*'
  _globals['_USER'].methods_by_name['enable']._loaded_options = None
  _globals['_USER'].methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v1/user/enable:\001*'
  _globals['_USER'].methods_by_name['disable']._loaded_options = None
  _globals['_USER'].methods_by_name['disable']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v1/user/disable:\001*'
  _globals['_USER'].methods_by_name['delete']._loaded_options = None
  _globals['_USER'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v1/user/delete:\001*'
  _globals['_USER'].methods_by_name['get']._loaded_options = None
  _globals['_USER'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\032\"\025/identity/v1/user/get:\001*'
  _globals['_USER'].methods_by_name['list']._loaded_options = None
  _globals['_USER'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v1/user/list:\001*'
  _globals['_USER'].methods_by_name['stat']._loaded_options = None
  _globals['_USER'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v1/user/stat:\001*'
  _globals['_USER'].methods_by_name['find']._loaded_options = None
  _globals['_USER'].methods_by_name['find']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v1/user/find:\001*'
  _globals['_USER'].methods_by_name['sync']._loaded_options = None
  _globals['_USER'].methods_by_name['sync']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v1/user/sync:\001*'
  _globals['_USERBACKEND']._serialized_start=2860
  _globals['_USERBACKEND']._serialized_end=2916
  _globals['_USERTYPE']._serialized_start=2918
  _globals['_USERTYPE']._serialized_end=2972
  _globals['_USERREQUIREDACTION']._serialized_start=2974
  _globals['_USERREQUIREDACTION']._serialized_end=3015
  _globals['_MFASTATE']._serialized_start=3017
  _globals['_MFASTATE']._serialized_end=3074
  _globals['_MFA']._serialized_start=188
  _globals['_MFA']._serialized_end=304
  _globals['_CREATEUSERREQUEST']._serialized_start=307
  _globals['_CREATEUSERREQUEST']._serialized_end=619
  _globals['_UPDATEUSERREQUEST']._serialized_start=622
  _globals['_UPDATEUSERREQUEST']._serialized_end=823
  _globals['_VERIFYEMAILREQUEST']._serialized_start=825
  _globals['_VERIFYEMAILREQUEST']._serialized_end=911
  _globals['_CONFIRMEMAILREQUEST']._serialized_start=913
  _globals['_CONFIRMEMAILREQUEST']._serialized_end=991
  _globals['_SETREQUIREDACTIONSREQUEST']._serialized_start=993
  _globals['_SETREQUIREDACTIONSREQUEST']._serialized_end=1119
  _globals['_ENABLEMFAREQUEST']._serialized_start=1121
  _globals['_ENABLEMFAREQUEST']._serialized_end=1235
  _globals['_DISABLEMFAREQUEST']._serialized_start=1237
  _globals['_DISABLEMFAREQUEST']._serialized_end=1307
  _globals['_CONFIRMMFAREQUEST']._serialized_start=1309
  _globals['_CONFIRMMFAREQUEST']._serialized_end=1385
  _globals['_USERREQUEST']._serialized_start=1387
  _globals['_USERREQUEST']._serialized_end=1436
  _globals['_GETUSERREQUEST']._serialized_start=1438
  _globals['_GETUSERREQUEST']._serialized_end=1504
  _globals['_USERQUERY']._serialized_start=1507
  _globals['_USERQUERY']._serialized_end=1753
  _globals['_USERINFO']._serialized_start=1756
  _globals['_USERINFO']._serialized_end=2319
  _globals['_USERINFO_STATE']._serialized_start=2262
  _globals['_USERINFO_STATE']._serialized_end=2319
  _globals['_USERSINFO']._serialized_start=2321
  _globals['_USERSINFO']._serialized_end=2406
  _globals['_USERSTATQUERY']._serialized_start=2408
  _globals['_USERSTATQUERY']._serialized_end=2496
  _globals['_FINDUSERSEARCH']._serialized_start=2498
  _globals['_FINDUSERSEARCH']._serialized_end=2568
  _globals['_FINDUSERQUERY']._serialized_start=2570
  _globals['_FINDUSERQUERY']._serialized_end=2662
  _globals['_FINDUSERINFO']._serialized_start=2664
  _globals['_FINDUSERINFO']._serialized_end=2763
  _globals['_FINDUSERSINFO']._serialized_start=2765
  _globals['_FINDUSERSINFO']._serialized_end=2858
  _globals['_USER']._serialized_start=3077
  _globals['_USER']._serialized_end=5262
# @@protoc_insertion_point(module_scope)
