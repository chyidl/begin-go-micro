# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: storage_message.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import memory_message_pb2 as memory__message__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='storage_message.proto',
  package='pcbook',
  syntax='proto3',
  serialized_options=b'Z\t./;pcbook',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15storage_message.proto\x12\x06pcbook\x1a\x14memory_message.proto\"z\n\x07Storage\x12&\n\x06\x64river\x18\x01 \x01(\x0e\x32\x16.pcbook.Storage.Driver\x12\x1e\n\x06memory\x18\x02 \x01(\x0b\x32\x0e.pcbook.Memory\"\'\n\x06\x44river\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03HDD\x10\x01\x12\x07\n\x03SSD\x10\x02\x42\x0bZ\t./;pcbookb\x06proto3'
  ,
  dependencies=[memory__message__pb2.DESCRIPTOR,])



_STORAGE_DRIVER = _descriptor.EnumDescriptor(
  name='Driver',
  full_name='pcbook.Storage.Driver',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HDD', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SSD', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=138,
  serialized_end=177,
)
_sym_db.RegisterEnumDescriptor(_STORAGE_DRIVER)


_STORAGE = _descriptor.Descriptor(
  name='Storage',
  full_name='pcbook.Storage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='driver', full_name='pcbook.Storage.driver', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='memory', full_name='pcbook.Storage.memory', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _STORAGE_DRIVER,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=177,
)

_STORAGE.fields_by_name['driver'].enum_type = _STORAGE_DRIVER
_STORAGE.fields_by_name['memory'].message_type = memory__message__pb2._MEMORY
_STORAGE_DRIVER.containing_type = _STORAGE
DESCRIPTOR.message_types_by_name['Storage'] = _STORAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Storage = _reflection.GeneratedProtocolMessageType('Storage', (_message.Message,), {
  'DESCRIPTOR' : _STORAGE,
  '__module__' : 'storage_message_pb2'
  # @@protoc_insertion_point(class_scope:pcbook.Storage)
  })
_sym_db.RegisterMessage(Storage)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
