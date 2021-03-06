# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: laptop_message.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import processer_message_pb2 as processer__message__pb2
import memory_message_pb2 as memory__message__pb2
import storage_message_pb2 as storage__message__pb2
import screen_message_pb2 as screen__message__pb2
import keyboard_message_pb2 as keyboard__message__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='laptop_message.proto',
  package='pcbook',
  syntax='proto3',
  serialized_options=b'Z\t./;pcbook',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14laptop_message.proto\x12\x06pcbook\x1a\x17processer_message.proto\x1a\x14memory_message.proto\x1a\x15storage_message.proto\x1a\x14screen_message.proto\x1a\x16keyboard_message.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf7\x02\n\x06Laptop\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05\x62rand\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x18\n\x03\x63pu\x18\x04 \x01(\x0b\x32\x0b.pcbook.CPU\x12\x1b\n\x03ram\x18\x05 \x01(\x0b\x32\x0e.pcbook.Memory\x12\x19\n\x04gpus\x18\x06 \x03(\x0b\x32\x0b.pcbook.GPU\x12!\n\x08storages\x18\x07 \x03(\x0b\x32\x0f.pcbook.Storage\x12\x1e\n\x06screen\x18\x08 \x01(\x0b\x32\x0e.pcbook.Screen\x12\"\n\x08keyboard\x18\t \x01(\x0b\x32\x10.pcbook.Keyboard\x12\x13\n\tweight_kg\x18\n \x01(\x01H\x00\x12\x13\n\tweight_lb\x18\x0b \x01(\x01H\x00\x12\x11\n\tprice_usd\x18\x0c \x01(\x01\x12\x14\n\x0crelease_year\x18\r \x01(\r\x12.\n\nupdated_at\x18\x0e \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\n\x06weightB\x0bZ\t./;pcbookb\x06proto3'
  ,
  dependencies=[processer__message__pb2.DESCRIPTOR,memory__message__pb2.DESCRIPTOR,storage__message__pb2.DESCRIPTOR,screen__message__pb2.DESCRIPTOR,keyboard__message__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_LAPTOP = _descriptor.Descriptor(
  name='Laptop',
  full_name='pcbook.Laptop',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pcbook.Laptop.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='brand', full_name='pcbook.Laptop.brand', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='pcbook.Laptop.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cpu', full_name='pcbook.Laptop.cpu', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ram', full_name='pcbook.Laptop.ram', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gpus', full_name='pcbook.Laptop.gpus', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='storages', full_name='pcbook.Laptop.storages', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='screen', full_name='pcbook.Laptop.screen', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='keyboard', full_name='pcbook.Laptop.keyboard', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='weight_kg', full_name='pcbook.Laptop.weight_kg', index=9,
      number=10, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='weight_lb', full_name='pcbook.Laptop.weight_lb', index=10,
      number=11, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='price_usd', full_name='pcbook.Laptop.price_usd', index=11,
      number=12, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='release_year', full_name='pcbook.Laptop.release_year', index=12,
      number=13, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='pcbook.Laptop.updated_at', index=13,
      number=14, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='weight', full_name='pcbook.Laptop.weight',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=182,
  serialized_end=557,
)

_LAPTOP.fields_by_name['cpu'].message_type = processer__message__pb2._CPU
_LAPTOP.fields_by_name['ram'].message_type = memory__message__pb2._MEMORY
_LAPTOP.fields_by_name['gpus'].message_type = processer__message__pb2._GPU
_LAPTOP.fields_by_name['storages'].message_type = storage__message__pb2._STORAGE
_LAPTOP.fields_by_name['screen'].message_type = screen__message__pb2._SCREEN
_LAPTOP.fields_by_name['keyboard'].message_type = keyboard__message__pb2._KEYBOARD
_LAPTOP.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_LAPTOP.oneofs_by_name['weight'].fields.append(
  _LAPTOP.fields_by_name['weight_kg'])
_LAPTOP.fields_by_name['weight_kg'].containing_oneof = _LAPTOP.oneofs_by_name['weight']
_LAPTOP.oneofs_by_name['weight'].fields.append(
  _LAPTOP.fields_by_name['weight_lb'])
_LAPTOP.fields_by_name['weight_lb'].containing_oneof = _LAPTOP.oneofs_by_name['weight']
DESCRIPTOR.message_types_by_name['Laptop'] = _LAPTOP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Laptop = _reflection.GeneratedProtocolMessageType('Laptop', (_message.Message,), {
  'DESCRIPTOR' : _LAPTOP,
  '__module__' : 'laptop_message_pb2'
  # @@protoc_insertion_point(class_scope:pcbook.Laptop)
  })
_sym_db.RegisterMessage(Laptop)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
