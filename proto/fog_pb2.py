# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: fog.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tfog.proto\x12\x05\x66ogpb\"H\n\nSensorData\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\r\n\x05value\x18\x03 \x01(\t\x12\x11\n\ttimestamp\x18\x04 \x01(\t\"\x1b\n\rSensorRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"1\n\x0eSensorResponse\x12\x1f\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x11.fogpb.SensorData\".\n\x0b\x45\x64geRequest\x12\x1f\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x11.fogpb.SensorData\"\x0e\n\x0c\x45\x64geResponse\"/\n\x0c\x43loudRequest\x12\x1f\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x11.fogpb.SensorData\"\x1b\n\tCloudData\x12\x0e\n\x06result\x18\x01 \x01(\x05\"/\n\rCloudResponse\x12\x1e\n\x04\x64\x61ta\x18\x01 \x01(\x0b\x32\x10.fogpb.CloudData2L\n\rSensorService\x12;\n\nStreamData\x12\x14.fogpb.SensorRequest\x1a\x15.fogpb.SensorResponse0\x01\x32\r\n\x0b\x45\x64geService2H\n\x0c\x43loudService\x12\x38\n\x0bProcessData\x12\x13.fogpb.CloudRequest\x1a\x14.fogpb.CloudResponseB\nZ\x08./;fogpbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'fog_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\010./;fogpb'
  _globals['_SENSORDATA']._serialized_start=20
  _globals['_SENSORDATA']._serialized_end=92
  _globals['_SENSORREQUEST']._serialized_start=94
  _globals['_SENSORREQUEST']._serialized_end=121
  _globals['_SENSORRESPONSE']._serialized_start=123
  _globals['_SENSORRESPONSE']._serialized_end=172
  _globals['_EDGEREQUEST']._serialized_start=174
  _globals['_EDGEREQUEST']._serialized_end=220
  _globals['_EDGERESPONSE']._serialized_start=222
  _globals['_EDGERESPONSE']._serialized_end=236
  _globals['_CLOUDREQUEST']._serialized_start=238
  _globals['_CLOUDREQUEST']._serialized_end=285
  _globals['_CLOUDDATA']._serialized_start=287
  _globals['_CLOUDDATA']._serialized_end=314
  _globals['_CLOUDRESPONSE']._serialized_start=316
  _globals['_CLOUDRESPONSE']._serialized_end=363
  _globals['_SENSORSERVICE']._serialized_start=365
  _globals['_SENSORSERVICE']._serialized_end=441
  _globals['_EDGESERVICE']._serialized_start=443
  _globals['_EDGESERVICE']._serialized_end=456
  _globals['_CLOUDSERVICE']._serialized_start=458
  _globals['_CLOUDSERVICE']._serialized_end=530
# @@protoc_insertion_point(module_scope)
