# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello_grpc.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10hello_grpc.proto\"\x19\n\x08MyNumber\x12\r\n\x05value\x18\x01 \x01(\x05\x32\x31\n\tMyService\x12$\n\nMyFunction\x12\t.MyNumber\x1a\t.MyNumber\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'hello_grpc_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_MYNUMBER']._serialized_start=20
  _globals['_MYNUMBER']._serialized_end=45
  _globals['_MYSERVICE']._serialized_start=47
  _globals['_MYSERVICE']._serialized_end=96
# @@protoc_insertion_point(module_scope)
