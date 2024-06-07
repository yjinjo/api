# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/private_widget.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.spaceone/api/dashboard/v1/private_widget.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xbb\x01\n\x1a\x43reatePrivateWidgetRequest\x12\x14\n\x0c\x64\x61shboard_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x13\n\x0bwidget_type\x18\x04 \x01(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xb8\x01\n\x1aUpdatePrivateWidgetRequest\x12\x11\n\twidget_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x13\n\x0bwidget_type\x18\x04 \x01(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\")\n\x14PrivateWidgetRequest\x12\x11\n\twidget_id\x18\x01 \x01(\t\"\x9e\x01\n\x18LoadPrivateWidgetRequest\x12\x11\n\twidget_id\x18\x01 \x01(\t\x12\x15\n\rdata_table_id\x18\x02 \x01(\t\x12\x31\n\x05query\x18\x03 \x01(\x0b\x32\".spaceone.api.core.v2.AnalyzeQuery\x12%\n\x04vars\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"w\n\x12PrivateWidgetQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x14\n\x0c\x64\x61shboard_id\x18\x02 \x01(\t\x12\x11\n\twidget_id\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\"\x91\x02\n\x11PrivateWidgetInfo\x12\x11\n\twidget_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x13\n\x0bwidget_type\x18\x04 \x01(\t\x12(\n\x07options\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07user_id\x18\x16 \x01(\t\x12\x14\n\x0c\x64\x61shboard_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"h\n\x12PrivateWidgetsInfo\x12=\n\x07results\x18\x01 \x03(\x0b\x32,.spaceone.api.dashboard.v1.PrivateWidgetInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x81\x07\n\rPrivateWidget\x12\x9d\x01\n\x06\x63reate\x12\x35.spaceone.api.dashboard.v1.CreatePrivateWidgetRequest\x1a,.spaceone.api.dashboard.v1.PrivateWidgetInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/private-widget/create:\x01*\x12\x9d\x01\n\x06update\x12\x35.spaceone.api.dashboard.v1.UpdatePrivateWidgetRequest\x1a,.spaceone.api.dashboard.v1.PrivateWidgetInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/private-widget/update:\x01*\x12\x81\x01\n\x06\x64\x65lete\x12/.spaceone.api.dashboard.v1.PrivateWidgetRequest\x1a\x16.google.protobuf.Empty\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/private-widget/delete:\x01*\x12\x82\x01\n\x04load\x12\x33.spaceone.api.dashboard.v1.LoadPrivateWidgetRequest\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&\"!/dashboard/v1/private-widget/load:\x01*\x12\x91\x01\n\x03get\x12/.spaceone.api.dashboard.v1.PrivateWidgetRequest\x1a,.spaceone.api.dashboard.v1.PrivateWidgetInfo\"+\x82\xd3\xe4\x93\x02%\" /dashboard/v1/private-widget/get:\x01*\x12\x92\x01\n\x04list\x12-.spaceone.api.dashboard.v1.PrivateWidgetQuery\x1a-.spaceone.api.dashboard.v1.PrivateWidgetsInfo\",\x82\xd3\xe4\x93\x02&\"!/dashboard/v1/private-widget/list:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.private_widget_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_PRIVATEWIDGET'].methods_by_name['create']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/private-widget/create:\001*'
  _globals['_PRIVATEWIDGET'].methods_by_name['update']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/private-widget/update:\001*'
  _globals['_PRIVATEWIDGET'].methods_by_name['delete']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/private-widget/delete:\001*'
  _globals['_PRIVATEWIDGET'].methods_by_name['load']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['load']._serialized_options = b'\202\323\344\223\002&\"!/dashboard/v1/private-widget/load:\001*'
  _globals['_PRIVATEWIDGET'].methods_by_name['get']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002%\" /dashboard/v1/private-widget/get:\001*'
  _globals['_PRIVATEWIDGET'].methods_by_name['list']._loaded_options = None
  _globals['_PRIVATEWIDGET'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002&\"!/dashboard/v1/private-widget/list:\001*'
  _globals['_CREATEPRIVATEWIDGETREQUEST']._serialized_start=201
  _globals['_CREATEPRIVATEWIDGETREQUEST']._serialized_end=388
  _globals['_UPDATEPRIVATEWIDGETREQUEST']._serialized_start=391
  _globals['_UPDATEPRIVATEWIDGETREQUEST']._serialized_end=575
  _globals['_PRIVATEWIDGETREQUEST']._serialized_start=577
  _globals['_PRIVATEWIDGETREQUEST']._serialized_end=618
  _globals['_LOADPRIVATEWIDGETREQUEST']._serialized_start=621
  _globals['_LOADPRIVATEWIDGETREQUEST']._serialized_end=779
  _globals['_PRIVATEWIDGETQUERY']._serialized_start=781
  _globals['_PRIVATEWIDGETQUERY']._serialized_end=900
  _globals['_PRIVATEWIDGETINFO']._serialized_start=903
  _globals['_PRIVATEWIDGETINFO']._serialized_end=1176
  _globals['_PRIVATEWIDGETSINFO']._serialized_start=1178
  _globals['_PRIVATEWIDGETSINFO']._serialized_end=1282
  _globals['_PRIVATEWIDGET']._serialized_start=1285
  _globals['_PRIVATEWIDGET']._serialized_end=2182
# @@protoc_insertion_point(module_scope)
