# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: serverstreaming.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15serverstreaming.proto\x12\x0fserverstreaming\"\x1a\n\x07Message\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x17\n\x06Number\x12\r\n\x05value\x18\x01 \x01(\x05\x32]\n\x0fServerStreaming\x12J\n\x11GetServerResponse\x12\x17.serverstreaming.Number\x1a\x18.serverstreaming.Message\"\x00\x30\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'serverstreaming_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_MESSAGE']._serialized_start=42
  _globals['_MESSAGE']._serialized_end=68
  _globals['_NUMBER']._serialized_start=70
  _globals['_NUMBER']._serialized_end=93
  _globals['_SERVERSTREAMING']._serialized_start=95
  _globals['_SERVERSTREAMING']._serialized_end=188
# @@protoc_insertion_point(module_scope)
