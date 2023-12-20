# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/cloud_service_report.proto
# Protobuf Python Version: 4.25.0
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n4spaceone/api/inventory/v1/cloud_service_report.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xd9\x02\n\x0eReportSchedule\x12\x46\n\x05state\x18\x01 \x01(\x0e\x32\x37.spaceone.api.inventory.v1.ReportSchedule.ScheduleState\x12\r\n\x05hours\x18\x02 \x03(\x05\x12I\n\x0c\x64\x61ys_of_week\x18\x03 \x03(\x0e\x32\x33.spaceone.api.inventory.v1.ReportSchedule.DayOfWeek\"C\n\rScheduleState\x12\x17\n\x13SCHEDULE_STATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"`\n\tDayOfWeek\x12\x14\n\x10\x44\x41Y_OF_WEEK_NONE\x10\x00\x12\x07\n\x03MON\x10\x01\x12\x07\n\x03TUE\x10\x02\x12\x07\n\x03WED\x10\x03\x12\x07\n\x03THU\x10\x04\x12\x07\n\x03\x46RI\x10\x05\x12\x07\n\x03SAT\x10\x06\x12\x07\n\x03SUN\x10\x07\"\xe6\x04\n\x1f\x43reateCloudServiceReportRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x33\n\x07options\x18\x02 \x03(\x0b\x32\".spaceone.api.core.v1.ExportOption\x12Z\n\x0b\x66ile_format\x18\x03 \x01(\x0e\x32\x45.spaceone.api.inventory.v1.CreateCloudServiceReportRequest.FileFormat\x12\x10\n\x08timezone\x18\x04 \x01(\t\x12\x10\n\x08language\x18\x05 \x01(\t\x12;\n\x08schedule\x18\x06 \x01(\x0b\x32).spaceone.api.inventory.v1.ReportSchedule\x12\'\n\x06target\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12`\n\x0eresource_group\x18\x14 \x01(\x0e\x32H.spaceone.api.inventory.v1.CreateCloudServiceReportRequest.ResourceGroup\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"6\n\nFileFormat\x12\x14\n\x10NONE_FILE_FORMAT\x10\x00\x12\t\n\x05\x45XCEL\x10\x01\x12\x07\n\x03\x43SV\x10\x02\"C\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"\xbc\x03\n\x1fUpdateCloudServiceReportRequest\x12\x11\n\treport_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x33\n\x07options\x18\x03 \x03(\x0b\x32\".spaceone.api.core.v1.ExportOption\x12Z\n\x0b\x66ile_format\x18\x04 \x01(\x0e\x32\x45.spaceone.api.inventory.v1.UpdateCloudServiceReportRequest.FileFormat\x12\x10\n\x08timezone\x18\x05 \x01(\t\x12\x10\n\x08language\x18\x06 \x01(\t\x12;\n\x08schedule\x18\x07 \x01(\x0b\x32).spaceone.api.inventory.v1.ReportSchedule\x12\'\n\x06target\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\"6\n\nFileFormat\x12\x14\n\x10NONE_FILE_FORMAT\x10\x00\x12\t\n\x05\x45XCEL\x10\x01\x12\x07\n\x03\x43SV\x10\x02\".\n\x19\x43loudServiceReportRequest\x12\x11\n\treport_id\x18\x01 \x01(\t\"|\n\x17\x43loudServiceReportQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x11\n\treport_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"\x9b\x05\n\x16\x43loudServiceReportInfo\x12\x11\n\treport_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x33\n\x07options\x18\x03 \x03(\x0b\x32\".spaceone.api.core.v1.ExportOption\x12Q\n\x0b\x66ile_format\x18\x04 \x01(\x0e\x32<.spaceone.api.inventory.v1.CloudServiceReportInfo.FileFormat\x12\x10\n\x08timezone\x18\x05 \x01(\t\x12\x10\n\x08language\x18\x06 \x01(\t\x12;\n\x08schedule\x18\x07 \x01(\x0b\x32).spaceone.api.inventory.v1.ReportSchedule\x12\'\n\x06target\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12W\n\x0eresource_group\x18\x14 \x01(\x0e\x32?.spaceone.api.inventory.v1.CloudServiceReportInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x14\n\x0clast_sent_at\x18  \x01(\t\"6\n\nFileFormat\x12\x14\n\x10NONE_FILE_FORMAT\x10\x00\x12\t\n\x05\x45XCEL\x10\x01\x12\x07\n\x03\x43SV\x10\x02\"C\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\"r\n\x17\x43loudServiceReportsInfo\x12\x42\n\x07results\x18\x01 \x03(\x0b\x32\x31.spaceone.api.inventory.v1.CloudServiceReportInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"S\n\x1b\x43loudServiceReportStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery2\xe5\x08\n\x12\x43loudServiceReport\x12\xad\x01\n\x06\x63reate\x12:.spaceone.api.inventory.v1.CreateCloudServiceReportRequest\x1a\x31.spaceone.api.inventory.v1.CloudServiceReportInfo\"4\x82\xd3\xe4\x93\x02.\")/inventory/v1/cloud-service-report/create:\x01*\x12\xad\x01\n\x06update\x12:.spaceone.api.inventory.v1.UpdateCloudServiceReportRequest\x1a\x31.spaceone.api.inventory.v1.CloudServiceReportInfo\"4\x82\xd3\xe4\x93\x02.\")/inventory/v1/cloud-service-report/update:\x01*\x12\x8c\x01\n\x06\x64\x65lete\x12\x34.spaceone.api.inventory.v1.CloudServiceReportRequest\x1a\x16.google.protobuf.Empty\"4\x82\xd3\xe4\x93\x02.\")/inventory/v1/cloud-service-report/delete:\x01*\x12\x88\x01\n\x04send\x12\x34.spaceone.api.inventory.v1.CloudServiceReportRequest\x1a\x16.google.protobuf.Empty\"2\x82\xd3\xe4\x93\x02,\"\'/inventory/v1/cloud-service-report/send:\x01*\x12\xa1\x01\n\x03get\x12\x34.spaceone.api.inventory.v1.CloudServiceReportRequest\x1a\x31.spaceone.api.inventory.v1.CloudServiceReportInfo\"1\x82\xd3\xe4\x93\x02+\"&/inventory/v1/cloud-service-report/get:\x01*\x12\xa2\x01\n\x04list\x12\x32.spaceone.api.inventory.v1.CloudServiceReportQuery\x1a\x32.spaceone.api.inventory.v1.CloudServiceReportsInfo\"2\x82\xd3\xe4\x93\x02,\"\'/inventory/v1/cloud-service-report/list:\x01*\x12\x8b\x01\n\x04stat\x12\x36.spaceone.api.inventory.v1.CloudServiceReportStatQuery\x1a\x17.google.protobuf.Struct\"2\x82\xd3\xe4\x93\x02,\"\'/inventory/v1/cloud-service-report/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.inventory.v1.cloud_service_report_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['create']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002.\")/inventory/v1/cloud-service-report/create:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['update']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002.\")/inventory/v1/cloud-service-report/update:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['delete']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002.\")/inventory/v1/cloud-service-report/delete:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['send']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['send']._serialized_options = b'\202\323\344\223\002,\"\'/inventory/v1/cloud-service-report/send:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['get']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002+\"&/inventory/v1/cloud-service-report/get:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['list']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002,\"\'/inventory/v1/cloud-service-report/list:\001*'
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['stat']._options = None
  _globals['_CLOUDSERVICEREPORT'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002,\"\'/inventory/v1/cloud-service-report/stat:\001*'
  _globals['_REPORTSCHEDULE']._serialized_start=207
  _globals['_REPORTSCHEDULE']._serialized_end=552
  _globals['_REPORTSCHEDULE_SCHEDULESTATE']._serialized_start=387
  _globals['_REPORTSCHEDULE_SCHEDULESTATE']._serialized_end=454
  _globals['_REPORTSCHEDULE_DAYOFWEEK']._serialized_start=456
  _globals['_REPORTSCHEDULE_DAYOFWEEK']._serialized_end=552
  _globals['_CREATECLOUDSERVICEREPORTREQUEST']._serialized_start=555
  _globals['_CREATECLOUDSERVICEREPORTREQUEST']._serialized_end=1169
  _globals['_CREATECLOUDSERVICEREPORTREQUEST_FILEFORMAT']._serialized_start=1046
  _globals['_CREATECLOUDSERVICEREPORTREQUEST_FILEFORMAT']._serialized_end=1100
  _globals['_CREATECLOUDSERVICEREPORTREQUEST_RESOURCEGROUP']._serialized_start=1102
  _globals['_CREATECLOUDSERVICEREPORTREQUEST_RESOURCEGROUP']._serialized_end=1169
  _globals['_UPDATECLOUDSERVICEREPORTREQUEST']._serialized_start=1172
  _globals['_UPDATECLOUDSERVICEREPORTREQUEST']._serialized_end=1616
  _globals['_UPDATECLOUDSERVICEREPORTREQUEST_FILEFORMAT']._serialized_start=1046
  _globals['_UPDATECLOUDSERVICEREPORTREQUEST_FILEFORMAT']._serialized_end=1100
  _globals['_CLOUDSERVICEREPORTREQUEST']._serialized_start=1618
  _globals['_CLOUDSERVICEREPORTREQUEST']._serialized_end=1664
  _globals['_CLOUDSERVICEREPORTQUERY']._serialized_start=1666
  _globals['_CLOUDSERVICEREPORTQUERY']._serialized_end=1790
  _globals['_CLOUDSERVICEREPORTINFO']._serialized_start=1793
  _globals['_CLOUDSERVICEREPORTINFO']._serialized_end=2460
  _globals['_CLOUDSERVICEREPORTINFO_FILEFORMAT']._serialized_start=1046
  _globals['_CLOUDSERVICEREPORTINFO_FILEFORMAT']._serialized_end=1100
  _globals['_CLOUDSERVICEREPORTINFO_RESOURCEGROUP']._serialized_start=1102
  _globals['_CLOUDSERVICEREPORTINFO_RESOURCEGROUP']._serialized_end=1169
  _globals['_CLOUDSERVICEREPORTSINFO']._serialized_start=2462
  _globals['_CLOUDSERVICEREPORTSINFO']._serialized_end=2576
  _globals['_CLOUDSERVICEREPORTSTATQUERY']._serialized_start=2578
  _globals['_CLOUDSERVICEREPORTSTATQUERY']._serialized_end=2661
  _globals['_CLOUDSERVICEREPORT']._serialized_start=2664
  _globals['_CLOUDSERVICEREPORT']._serialized_end=3789
# @@protoc_insertion_point(module_scope)
