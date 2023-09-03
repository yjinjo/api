# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/cost.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v1 import query_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/cost_analysis/v1/cost.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\"\xf4\x02\n\x11\x43reateCostRequest\x12\x0c\n\x04\x63ost\x18\x01 \x01(\x02\x12\x16\n\x0eusage_quantity\x18\x02 \x01(\x02\x12\x12\n\nusage_unit\x18\x03 \x01(\x02\x12\x10\n\x08provider\x18\x04 \x01(\t\x12\x13\n\x0bregion_code\x18\x05 \x01(\t\x12\x0f\n\x07product\x18\x06 \x01(\t\x12\x12\n\nusage_type\x18\x07 \x01(\t\x12\x10\n\x08resource\x18\x08 \x01(\t\x12%\n\x04tags\x18\x15 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0f\x61\x64\x64itional_info\x18\x16 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x1a\n\x12service_account_id\x18\x1f \x01(\t\x12\x12\n\nproject_id\x18  \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18! \x01(\t\x12\x11\n\tdomain_id\x18\" \x01(\t\x12\x13\n\x0b\x62illed_date\x18) \x01(\t\"1\n\x0b\x43ostRequest\x12\x0f\n\x07\x63ost_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\"B\n\x0eGetCostRequest\x12\x0f\n\x07\x63ost_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x02 \x01(\t\x12\x0c\n\x04only\x18\x03 \x03(\t\"\xef\x02\n\tCostQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v1.Query\x12\x0f\n\x07\x63ost_id\x18\x02 \x01(\t\x12\x10\n\x08provider\x18\x03 \x01(\t\x12\x13\n\x0bregion_code\x18\x04 \x01(\t\x12\x12\n\nregion_key\x18\x05 \x01(\t\x12\x0f\n\x07product\x18\x06 \x01(\t\x12\x12\n\nusage_type\x18\x07 \x01(\t\x12\x10\n\x08resource\x18\x08 \x01(\t\x12\x1a\n\x12service_account_id\x18\x0b \x01(\t\x12\x12\n\nproject_id\x18\x0c \x01(\t\x12\x18\n\x10project_group_id\x18\r \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18\x0e \x01(\t\x12\x11\n\tdomain_id\x18\x0f \x01(\t\x12\x13\n\x0b\x62illed_year\x18\x15 \x01(\t\x12\x14\n\x0c\x62illed_month\x18\x16 \x01(\t\x12\x13\n\x0b\x62illed_date\x18\x17 \x01(\t\"\xbb\x03\n\x08\x43ostInfo\x12\x0f\n\x07\x63ost_id\x18\x01 \x01(\t\x12\x0c\n\x04\x63ost\x18\x02 \x01(\x02\x12\x16\n\x0eusage_quantity\x18\x03 \x01(\x02\x12\x12\n\nusage_unit\x18\x04 \x01(\t\x12\x10\n\x08provider\x18\x05 \x01(\t\x12\x13\n\x0bregion_code\x18\x06 \x01(\t\x12\x12\n\nregion_key\x18\x07 \x01(\t\x12\x0f\n\x07product\x18\x08 \x01(\t\x12\x12\n\nusage_type\x18\t \x01(\t\x12\x10\n\x08resource\x18\n \x01(\t\x12%\n\x04tags\x18\x15 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x30\n\x0f\x61\x64\x64itional_info\x18\x16 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x1a\n\x12service_account_id\x18\x1f \x01(\t\x12\x12\n\nproject_id\x18  \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18! \x01(\t\x12\x11\n\tdomain_id\x18\" \x01(\t\x12\x13\n\x0b\x62illed_year\x18) \x01(\t\x12\x14\n\x0c\x62illed_month\x18* \x01(\t\x12\x13\n\x0b\x62illed_date\x18+ \x01(\t\"Z\n\tCostsInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.cost_analysis.v1.CostInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"z\n\x10\x43ostAnalyzeQuery\x12;\n\x05query\x18\x01 \x01(\x0b\x32,.spaceone.api.core.v1.TimeSeriesAnalyzeQuery\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t\"p\n\rCostStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v1.StatisticsQuery\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x03 \x01(\t2\x91\x06\n\x04\x43ost\x12\x8d\x01\n\x06\x63reate\x12\x30.spaceone.api.cost_analysis.v1.CreateCostRequest\x1a\'.spaceone.api.cost_analysis.v1.CostInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/cost-analysis/v1/cost/create:\x01*\x12v\n\x06\x64\x65lete\x12*.spaceone.api.cost_analysis.v1.CostRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/cost-analysis/v1/cost/delete:\x01*\x12\x84\x01\n\x03get\x12-.spaceone.api.cost_analysis.v1.GetCostRequest\x1a\'.spaceone.api.cost_analysis.v1.CostInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/cost-analysis/v1/cost/get:\x01*\x12\x82\x01\n\x04list\x12(.spaceone.api.cost_analysis.v1.CostQuery\x1a(.spaceone.api.cost_analysis.v1.CostsInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/cost-analysis/v1/cost/list:\x01*\x12~\n\x07\x61nalyze\x12/.spaceone.api.cost_analysis.v1.CostAnalyzeQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#\"\x1e/cost-analysis/v1/cost/analyze:\x01*\x12u\n\x04stat\x12,.spaceone.api.cost_analysis.v1.CostStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/cost-analysis/v1/cost/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')



_CREATECOSTREQUEST = DESCRIPTOR.message_types_by_name['CreateCostRequest']
_COSTREQUEST = DESCRIPTOR.message_types_by_name['CostRequest']
_GETCOSTREQUEST = DESCRIPTOR.message_types_by_name['GetCostRequest']
_COSTQUERY = DESCRIPTOR.message_types_by_name['CostQuery']
_COSTINFO = DESCRIPTOR.message_types_by_name['CostInfo']
_COSTSINFO = DESCRIPTOR.message_types_by_name['CostsInfo']
_COSTANALYZEQUERY = DESCRIPTOR.message_types_by_name['CostAnalyzeQuery']
_COSTSTATQUERY = DESCRIPTOR.message_types_by_name['CostStatQuery']
CreateCostRequest = _reflection.GeneratedProtocolMessageType('CreateCostRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECOSTREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CreateCostRequest)
  })
_sym_db.RegisterMessage(CreateCostRequest)

CostRequest = _reflection.GeneratedProtocolMessageType('CostRequest', (_message.Message,), {
  'DESCRIPTOR' : _COSTREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostRequest)
  })
_sym_db.RegisterMessage(CostRequest)

GetCostRequest = _reflection.GeneratedProtocolMessageType('GetCostRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCOSTREQUEST,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.GetCostRequest)
  })
_sym_db.RegisterMessage(GetCostRequest)

CostQuery = _reflection.GeneratedProtocolMessageType('CostQuery', (_message.Message,), {
  'DESCRIPTOR' : _COSTQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostQuery)
  })
_sym_db.RegisterMessage(CostQuery)

CostInfo = _reflection.GeneratedProtocolMessageType('CostInfo', (_message.Message,), {
  'DESCRIPTOR' : _COSTINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostInfo)
  })
_sym_db.RegisterMessage(CostInfo)

CostsInfo = _reflection.GeneratedProtocolMessageType('CostsInfo', (_message.Message,), {
  'DESCRIPTOR' : _COSTSINFO,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostsInfo)
  })
_sym_db.RegisterMessage(CostsInfo)

CostAnalyzeQuery = _reflection.GeneratedProtocolMessageType('CostAnalyzeQuery', (_message.Message,), {
  'DESCRIPTOR' : _COSTANALYZEQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostAnalyzeQuery)
  })
_sym_db.RegisterMessage(CostAnalyzeQuery)

CostStatQuery = _reflection.GeneratedProtocolMessageType('CostStatQuery', (_message.Message,), {
  'DESCRIPTOR' : _COSTSTATQUERY,
  '__module__' : 'spaceone.api.cost_analysis.v1.cost_pb2'
  # @@protoc_insertion_point(class_scope:spaceone.api.cost_analysis.v1.CostStatQuery)
  })
_sym_db.RegisterMessage(CostStatQuery)

_COST = DESCRIPTOR.services_by_name['Cost']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _COST.methods_by_name['create']._options = None
  _COST.methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/cost-analysis/v1/cost/create:\001*'
  _COST.methods_by_name['delete']._options = None
  _COST.methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/cost-analysis/v1/cost/delete:\001*'
  _COST.methods_by_name['get']._options = None
  _COST.methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/cost-analysis/v1/cost/get:\001*'
  _COST.methods_by_name['list']._options = None
  _COST.methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/cost-analysis/v1/cost/list:\001*'
  _COST.methods_by_name['analyze']._options = None
  _COST.methods_by_name['analyze']._serialized_options = b'\202\323\344\223\002#\"\036/cost-analysis/v1/cost/analyze:\001*'
  _COST.methods_by_name['stat']._options = None
  _COST.methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/cost-analysis/v1/cost/stat:\001*'
  _CREATECOSTREQUEST._serialized_start=199
  _CREATECOSTREQUEST._serialized_end=571
  _COSTREQUEST._serialized_start=573
  _COSTREQUEST._serialized_end=622
  _GETCOSTREQUEST._serialized_start=624
  _GETCOSTREQUEST._serialized_end=690
  _COSTQUERY._serialized_start=693
  _COSTQUERY._serialized_end=1060
  _COSTINFO._serialized_start=1063
  _COSTINFO._serialized_end=1506
  _COSTSINFO._serialized_start=1508
  _COSTSINFO._serialized_end=1598
  _COSTANALYZEQUERY._serialized_start=1600
  _COSTANALYZEQUERY._serialized_end=1722
  _COSTSTATQUERY._serialized_start=1724
  _COSTSTATQUERY._serialized_end=1836
  _COST._serialized_start=1839
  _COST._serialized_end=2624
# @@protoc_insertion_point(module_scope)
