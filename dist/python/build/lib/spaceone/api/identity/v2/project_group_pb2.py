# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/project_group.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n,spaceone/api/identity/v2/project_group.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"i\n\x19\x43reateProjectGroupRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x17\n\x0fparent_group_id\x18\x03 \x01(\t\"j\n\x19UpdateProjectGroupRequest\x12\x18\n\x10project_group_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"M\n\x18\x43hangeParentGroupRequest\x12\x18\n\x10project_group_id\x18\x01 \x01(\t\x12\x17\n\x0fparent_group_id\x18\x02 \x01(\t\"/\n\x13ProjectGroupRequest\x12\x18\n\x10project_group_id\x18\x01 \x01(\t\"C\n\x18UsersProjectGroupRequest\x12\x18\n\x10project_group_id\x18\x01 \x01(\t\x12\r\n\x05users\x18\x02 \x03(\t\"\xa4\x02\n\x10ProjectGroupInfo\x12\x18\n\x10project_group_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\r\n\x05users\x18\x04 \x03(\t\x12\x14\n\x0creference_id\x18\x05 \x01(\t\x12\x12\n\nis_managed\x18\x06 \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x17\n\x0fparent_group_id\x18\x17 \x01(\t\x12\x1a\n\x12trusted_account_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x16\n\x0elast_synced_at\x18  \x01(\t\"\x9c\x01\n\x17ProjectGroupSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x18\n\x10project_group_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x17\n\x0fparent_group_id\x18\x16 \x01(\t\"e\n\x11ProjectGroupsInfo\x12;\n\x07results\x18\x01 \x03(\x0b\x32*.spaceone.api.identity.v2.ProjectGroupInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15ProjectGroupStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xd9\n\n\x0cProjectGroup\x12\x97\x01\n\x06\x63reate\x12\x33.spaceone.api.identity.v2.CreateProjectGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/project-group/create:\x01*\x12\x97\x01\n\x06update\x12\x33.spaceone.api.identity.v2.UpdateProjectGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/project-group/update:\x01*\x12\xb0\x01\n\x13\x63hange_parent_group\x12\x32.spaceone.api.identity.v2.ChangeParentGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\"9\x82\xd3\xe4\x93\x02\x33\"./identity/v2/project-group/change-parent-group:\x01*\x12}\n\x06\x64\x65lete\x12-.spaceone.api.identity.v2.ProjectGroupRequest\x1a\x16.google.protobuf.Empty\",\x82\xd3\xe4\x93\x02&\"!/identity/v2/project-group/delete:\x01*\x12\x9c\x01\n\tadd_users\x12\x32.spaceone.api.identity.v2.UsersProjectGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\"/\x82\xd3\xe4\x93\x02)\"$/identity/v2/project-group/add-users:\x01*\x12\xa2\x01\n\x0cremove_users\x12\x32.spaceone.api.identity.v2.UsersProjectGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\"2\x82\xd3\xe4\x93\x02,\"\'/identity/v2/project-group/remove-users:\x01*\x12\x8b\x01\n\x03get\x12-.spaceone.api.identity.v2.ProjectGroupRequest\x1a*.spaceone.api.identity.v2.ProjectGroupInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v2/project-group/get:\x01*\x12\x92\x01\n\x04list\x12\x31.spaceone.api.identity.v2.ProjectGroupSearchQuery\x1a+.spaceone.api.identity.v2.ProjectGroupsInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/identity/v2/project-group/list:\x01*\x12|\n\x04stat\x12/.spaceone.api.identity.v2.ProjectGroupStatQuery\x1a\x17.google.protobuf.Struct\"*\x82\xd3\xe4\x93\x02$\"\x1f/identity/v2/project-group/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.project_group_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_PROJECTGROUP'].methods_by_name['create']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/project-group/create:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['update']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/project-group/update:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['change_parent_group']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['change_parent_group']._serialized_options = b'\202\323\344\223\0023\"./identity/v2/project-group/change-parent-group:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['delete']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002&\"!/identity/v2/project-group/delete:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['add_users']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['add_users']._serialized_options = b'\202\323\344\223\002)\"$/identity/v2/project-group/add-users:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['remove_users']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['remove_users']._serialized_options = b'\202\323\344\223\002,\"\'/identity/v2/project-group/remove-users:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['get']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v2/project-group/get:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['list']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002$\"\037/identity/v2/project-group/list:\001*'
  _globals['_PROJECTGROUP'].methods_by_name['stat']._loaded_options = None
  _globals['_PROJECTGROUP'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002$\"\037/identity/v2/project-group/stat:\001*'
  _globals['_CREATEPROJECTGROUPREQUEST']._serialized_start=197
  _globals['_CREATEPROJECTGROUPREQUEST']._serialized_end=302
  _globals['_UPDATEPROJECTGROUPREQUEST']._serialized_start=304
  _globals['_UPDATEPROJECTGROUPREQUEST']._serialized_end=410
  _globals['_CHANGEPARENTGROUPREQUEST']._serialized_start=412
  _globals['_CHANGEPARENTGROUPREQUEST']._serialized_end=489
  _globals['_PROJECTGROUPREQUEST']._serialized_start=491
  _globals['_PROJECTGROUPREQUEST']._serialized_end=538
  _globals['_USERSPROJECTGROUPREQUEST']._serialized_start=540
  _globals['_USERSPROJECTGROUPREQUEST']._serialized_end=607
  _globals['_PROJECTGROUPINFO']._serialized_start=610
  _globals['_PROJECTGROUPINFO']._serialized_end=902
  _globals['_PROJECTGROUPSEARCHQUERY']._serialized_start=905
  _globals['_PROJECTGROUPSEARCHQUERY']._serialized_end=1061
  _globals['_PROJECTGROUPSINFO']._serialized_start=1063
  _globals['_PROJECTGROUPSINFO']._serialized_end=1164
  _globals['_PROJECTGROUPSTATQUERY']._serialized_start=1166
  _globals['_PROJECTGROUPSTATQUERY']._serialized_end=1243
  _globals['_PROJECTGROUP']._serialized_start=1246
  _globals['_PROJECTGROUP']._serialized_end=2615
# @@protoc_insertion_point(module_scope)
