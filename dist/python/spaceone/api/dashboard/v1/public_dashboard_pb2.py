# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/public_dashboard.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0spaceone/api/dashboard/v1/public_dashboard.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x82\x06\n\x1c\x43reatePublicDashboardRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0btemplate_id\x18\x02 \x01(\t\x12[\n\rtemplate_type\x18\x03 \x01(\x0e\x32\x44.spaceone.api.dashboard.v1.CreatePublicDashboardRequest.TemplateType\x12+\n\x07layouts\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\t \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12]\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x45.spaceone.api.dashboard.v1.CreatePublicDashboardRequest.ResourceGroup\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"N\n\x0cTemplateType\x12\x16\n\x12TEMPLATE_TYPE_NONE\x10\x00\x12\x0b\n\x07MANAGED\x10\x01\x12\n\n\x06\x43USTOM\x10\x02\x12\r\n\tEXTENSION\x10\x03\"\xc4\x04\n\x1cUpdatePublicDashboardRequest\x12\x1b\n\x13public_dashboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0btemplate_id\x18\x03 \x01(\t\x12[\n\rtemplate_type\x18\x04 \x01(\x0e\x32\x44.spaceone.api.dashboard.v1.UpdatePublicDashboardRequest.TemplateType\x12+\n\x07layouts\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\n \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\"N\n\x0cTemplateType\x12\x16\n\x12TEMPLATE_TYPE_NONE\x10\x00\x12\x0b\n\x07MANAGED\x10\x01\x12\n\n\x06\x43USTOM\x10\x02\x12\r\n\tEXTENSION\x10\x03\"5\n\x16PublicDashboardRequest\x12\x1b\n\x13public_dashboard_id\x18\x01 \x01(\t\"M\n\x1dPublicDashboardVersionRequest\x12\x1b\n\x13public_dashboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x04 \x01(\x05\"}\n!PublicDashboardVersionSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1b\n\x13public_dashboard_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\x05\"\x97\x01\n\x14PublicDashboardQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1b\n\x13public_dashboard_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\xd0\x06\n\x13PublicDashboardInfo\x12\x1b\n\x13public_dashboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0btemplate_id\x18\x03 \x01(\t\x12R\n\rtemplate_type\x18\x04 \x01(\x0e\x32;.spaceone.api.dashboard.v1.PublicDashboardInfo.TemplateType\x12\x0f\n\x07version\x18\x05 \x01(\x05\x12+\n\x07layouts\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0c\x64isplay_info\x18\n \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\x06labels\x18\x0b \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12T\n\x0eresource_group\x18\x14 \x01(\x0e\x32<.spaceone.api.dashboard.v1.PublicDashboardInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"N\n\x0cTemplateType\x12\x16\n\x12TEMPLATE_TYPE_NONE\x10\x00\x12\x0b\n\x07MANAGED\x10\x01\x12\n\n\x06\x43USTOM\x10\x02\x12\r\n\tEXTENSION\x10\x03\"l\n\x14PublicDashboardsInfo\x12?\n\x07results\x18\x01 \x03(\x0b\x32..spaceone.api.dashboard.v1.PublicDashboardInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"P\n\x18PublicDashboardStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\xb8\x02\n\x1aPublicDashboardVersionInfo\x12\x1b\n\x13public_dashboard_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\x05\x12\x0e\n\x06latest\x18\x03 \x01(\x08\x12+\n\x07layouts\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12*\n\tvariables\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12)\n\x08settings\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x31\n\x10variables_schema\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"z\n\x1bPublicDashboardVersionsInfo\x12\x46\n\x07results\x18\x01 \x03(\x0b\x32\x35.spaceone.api.dashboard.v1.PublicDashboardVersionInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\xf0\x0c\n\x0fPublicDashboard\x12\xa3\x01\n\x06\x63reate\x12\x37.spaceone.api.dashboard.v1.CreatePublicDashboardRequest\x1a..spaceone.api.dashboard.v1.PublicDashboardInfo\"0\x82\xd3\xe4\x93\x02*\"%/dashboard/v1/public-dashboard/create:\x01*\x12\xa3\x01\n\x06update\x12\x37.spaceone.api.dashboard.v1.UpdatePublicDashboardRequest\x1a..spaceone.api.dashboard.v1.PublicDashboardInfo\"0\x82\xd3\xe4\x93\x02*\"%/dashboard/v1/public-dashboard/update:\x01*\x12\x85\x01\n\x06\x64\x65lete\x12\x31.spaceone.api.dashboard.v1.PublicDashboardRequest\x1a\x16.google.protobuf.Empty\"0\x82\xd3\xe4\x93\x02*\"%/dashboard/v1/public-dashboard/delete:\x01*\x12\x97\x01\n\x03get\x12\x31.spaceone.api.dashboard.v1.PublicDashboardRequest\x1a..spaceone.api.dashboard.v1.PublicDashboardInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/dashboard/v1/public-dashboard/get:\x01*\x12\x9c\x01\n\x0e\x64\x65lete_version\x12\x38.spaceone.api.dashboard.v1.PublicDashboardVersionRequest\x1a\x16.google.protobuf.Empty\"8\x82\xd3\xe4\x93\x02\x32\"-/dashboard/v1/public-dashboard/delete-version:\x01*\x12\xb4\x01\n\x0erevert_version\x12\x38.spaceone.api.dashboard.v1.PublicDashboardVersionRequest\x1a..spaceone.api.dashboard.v1.PublicDashboardInfo\"8\x82\xd3\xe4\x93\x02\x32\"-/dashboard/v1/public-dashboard/revert-version:\x01*\x12\xb5\x01\n\x0bget_version\x12\x38.spaceone.api.dashboard.v1.PublicDashboardVersionRequest\x1a\x35.spaceone.api.dashboard.v1.PublicDashboardVersionInfo\"5\x82\xd3\xe4\x93\x02/\"*/dashboard/v1/public-dashboard/get-version:\x01*\x12\xbe\x01\n\rlist_versions\x12<.spaceone.api.dashboard.v1.PublicDashboardVersionSearchQuery\x1a\x36.spaceone.api.dashboard.v1.PublicDashboardVersionsInfo\"7\x82\xd3\xe4\x93\x02\x31\",/dashboard/v1/public-dashboard/list-versions:\x01*\x12\x98\x01\n\x04list\x12/.spaceone.api.dashboard.v1.PublicDashboardQuery\x1a/.spaceone.api.dashboard.v1.PublicDashboardsInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/public-dashboard/list:\x01*\x12\x84\x01\n\x04stat\x12\x33.spaceone.api.dashboard.v1.PublicDashboardStatQuery\x1a\x17.google.protobuf.Struct\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/public-dashboard/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.public_dashboard_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_PUBLICDASHBOARD'].methods_by_name['create']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002*\"%/dashboard/v1/public-dashboard/create:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['update']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002*\"%/dashboard/v1/public-dashboard/update:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['delete']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002*\"%/dashboard/v1/public-dashboard/delete:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['get']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\'\"\"/dashboard/v1/public-dashboard/get:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['delete_version']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['delete_version']._serialized_options = b'\202\323\344\223\0022\"-/dashboard/v1/public-dashboard/delete-version:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['revert_version']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['revert_version']._serialized_options = b'\202\323\344\223\0022\"-/dashboard/v1/public-dashboard/revert-version:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['get_version']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['get_version']._serialized_options = b'\202\323\344\223\002/\"*/dashboard/v1/public-dashboard/get-version:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['list_versions']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['list_versions']._serialized_options = b'\202\323\344\223\0021\",/dashboard/v1/public-dashboard/list-versions:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['list']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/public-dashboard/list:\001*'
  _globals['_PUBLICDASHBOARD'].methods_by_name['stat']._loaded_options = None
  _globals['_PUBLICDASHBOARD'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/public-dashboard/stat:\001*'
  _globals['_CREATEPUBLICDASHBOARDREQUEST']._serialized_start=203
  _globals['_CREATEPUBLICDASHBOARDREQUEST']._serialized_end=973
  _globals['_CREATEPUBLICDASHBOARDREQUEST_RESOURCEGROUP']._serialized_start=813
  _globals['_CREATEPUBLICDASHBOARDREQUEST_RESOURCEGROUP']._serialized_end=893
  _globals['_CREATEPUBLICDASHBOARDREQUEST_TEMPLATETYPE']._serialized_start=895
  _globals['_CREATEPUBLICDASHBOARDREQUEST_TEMPLATETYPE']._serialized_end=973
  _globals['_UPDATEPUBLICDASHBOARDREQUEST']._serialized_start=976
  _globals['_UPDATEPUBLICDASHBOARDREQUEST']._serialized_end=1556
  _globals['_UPDATEPUBLICDASHBOARDREQUEST_TEMPLATETYPE']._serialized_start=895
  _globals['_UPDATEPUBLICDASHBOARDREQUEST_TEMPLATETYPE']._serialized_end=973
  _globals['_PUBLICDASHBOARDREQUEST']._serialized_start=1558
  _globals['_PUBLICDASHBOARDREQUEST']._serialized_end=1611
  _globals['_PUBLICDASHBOARDVERSIONREQUEST']._serialized_start=1613
  _globals['_PUBLICDASHBOARDVERSIONREQUEST']._serialized_end=1690
  _globals['_PUBLICDASHBOARDVERSIONSEARCHQUERY']._serialized_start=1692
  _globals['_PUBLICDASHBOARDVERSIONSEARCHQUERY']._serialized_end=1817
  _globals['_PUBLICDASHBOARDQUERY']._serialized_start=1820
  _globals['_PUBLICDASHBOARDQUERY']._serialized_end=1971
  _globals['_PUBLICDASHBOARDINFO']._serialized_start=1974
  _globals['_PUBLICDASHBOARDINFO']._serialized_end=2822
  _globals['_PUBLICDASHBOARDINFO_RESOURCEGROUP']._serialized_start=813
  _globals['_PUBLICDASHBOARDINFO_RESOURCEGROUP']._serialized_end=893
  _globals['_PUBLICDASHBOARDINFO_TEMPLATETYPE']._serialized_start=895
  _globals['_PUBLICDASHBOARDINFO_TEMPLATETYPE']._serialized_end=973
  _globals['_PUBLICDASHBOARDSINFO']._serialized_start=2824
  _globals['_PUBLICDASHBOARDSINFO']._serialized_end=2932
  _globals['_PUBLICDASHBOARDSTATQUERY']._serialized_start=2934
  _globals['_PUBLICDASHBOARDSTATQUERY']._serialized_end=3014
  _globals['_PUBLICDASHBOARDVERSIONINFO']._serialized_start=3017
  _globals['_PUBLICDASHBOARDVERSIONINFO']._serialized_end=3329
  _globals['_PUBLICDASHBOARDVERSIONSINFO']._serialized_start=3331
  _globals['_PUBLICDASHBOARDVERSIONSINFO']._serialized_end=3453
  _globals['_PUBLICDASHBOARD']._serialized_start=3456
  _globals['_PUBLICDASHBOARD']._serialized_end=5104
# @@protoc_insertion_point(module_scope)
