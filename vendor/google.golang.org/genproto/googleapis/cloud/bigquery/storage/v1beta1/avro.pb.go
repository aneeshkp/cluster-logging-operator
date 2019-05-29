// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1beta1/avro.proto

package storage // import "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Avro schema.
type AvroSchema struct {
	// Json serialized schema, as described at
	// https://avro.apache.org/docs/1.8.1/spec.html
	Schema               string   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvroSchema) Reset()         { *m = AvroSchema{} }
func (m *AvroSchema) String() string { return proto.CompactTextString(m) }
func (*AvroSchema) ProtoMessage()    {}
func (*AvroSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_avro_4770a0731176dc2d, []int{0}
}
func (m *AvroSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvroSchema.Unmarshal(m, b)
}
func (m *AvroSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvroSchema.Marshal(b, m, deterministic)
}
func (dst *AvroSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvroSchema.Merge(dst, src)
}
func (m *AvroSchema) XXX_Size() int {
	return xxx_messageInfo_AvroSchema.Size(m)
}
func (m *AvroSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_AvroSchema.DiscardUnknown(m)
}

var xxx_messageInfo_AvroSchema proto.InternalMessageInfo

func (m *AvroSchema) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

// Avro rows.
type AvroRows struct {
	// Binary serialized rows in a block.
	SerializedBinaryRows []byte `protobuf:"bytes,1,opt,name=serialized_binary_rows,json=serializedBinaryRows,proto3" json:"serialized_binary_rows,omitempty"`
	// The count of rows in the returning block.
	RowCount             int64    `protobuf:"varint,2,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvroRows) Reset()         { *m = AvroRows{} }
func (m *AvroRows) String() string { return proto.CompactTextString(m) }
func (*AvroRows) ProtoMessage()    {}
func (*AvroRows) Descriptor() ([]byte, []int) {
	return fileDescriptor_avro_4770a0731176dc2d, []int{1}
}
func (m *AvroRows) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvroRows.Unmarshal(m, b)
}
func (m *AvroRows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvroRows.Marshal(b, m, deterministic)
}
func (dst *AvroRows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvroRows.Merge(dst, src)
}
func (m *AvroRows) XXX_Size() int {
	return xxx_messageInfo_AvroRows.Size(m)
}
func (m *AvroRows) XXX_DiscardUnknown() {
	xxx_messageInfo_AvroRows.DiscardUnknown(m)
}

var xxx_messageInfo_AvroRows proto.InternalMessageInfo

func (m *AvroRows) GetSerializedBinaryRows() []byte {
	if m != nil {
		return m.SerializedBinaryRows
	}
	return nil
}

func (m *AvroRows) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func init() {
	proto.RegisterType((*AvroSchema)(nil), "google.cloud.bigquery.storage.v1beta1.AvroSchema")
	proto.RegisterType((*AvroRows)(nil), "google.cloud.bigquery.storage.v1beta1.AvroRows")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1beta1/avro.proto", fileDescriptor_avro_4770a0731176dc2d)
}

var fileDescriptor_avro_4770a0731176dc2d = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0xaa, 0x30, 0xd6, 0xe0, 0xa9, 0xc8, 0x18, 0x78, 0x19, 0x43, 0x61, 0x5e, 0x12,
	0x87, 0xde, 0x3c, 0x59, 0xaf, 0x1e, 0xa4, 0xde, 0x04, 0x29, 0x69, 0x17, 0x9e, 0x81, 0xae, 0xdf,
	0xf9, 0xd2, 0x1f, 0xcc, 0xb3, 0x7f, 0xb8, 0x24, 0x8d, 0x78, 0xd4, 0x5b, 0xbe, 0x7c, 0xf3, 0x81,
	0xf7, 0x9e, 0xb8, 0x21, 0x80, 0x1a, 0xa3, 0xea, 0x06, 0xfd, 0x4e, 0x55, 0x96, 0x3e, 0x7a, 0xc3,
	0x47, 0xe5, 0x3a, 0xb0, 0x26, 0xa3, 0x86, 0x6d, 0x65, 0x3a, 0xbd, 0x55, 0x7a, 0x60, 0xc8, 0x03,
	0xa3, 0x43, 0x76, 0x35, 0x09, 0x19, 0x84, 0xfc, 0x11, 0x32, 0x0a, 0x19, 0xc5, 0xfa, 0x52, 0x88,
	0x87, 0x81, 0xf1, 0x52, 0xbf, 0x9b, 0xbd, 0xce, 0x16, 0x62, 0xe6, 0xc2, 0x6b, 0x99, 0xac, 0x92,
	0x4d, 0x5a, 0xc4, 0xb4, 0x7e, 0x13, 0x73, 0xff, 0xab, 0xc0, 0xe8, 0xb2, 0x3b, 0xb1, 0x70, 0x86,
	0xad, 0x6e, 0xec, 0xa7, 0xd9, 0x95, 0x95, 0x6d, 0x35, 0x1f, 0x4b, 0xc6, 0xe8, 0x82, 0x39, 0x2b,
	0xce, 0x7f, 0xdb, 0x3c, 0x94, 0x41, 0x5d, 0x88, 0x94, 0x31, 0x96, 0x35, 0xfa, 0xb6, 0x5b, 0x9e,
	0xac, 0x92, 0xcd, 0x69, 0x31, 0x67, 0x8c, 0x8f, 0x3e, 0xe7, 0x5f, 0x89, 0xb8, 0xae, 0xb1, 0x97,
	0xff, 0x1a, 0x39, 0x4f, 0xfd, 0x28, 0xcf, 0x7e, 0xc9, 0xd7, 0xa7, 0x28, 0x08, 0x8d, 0x6e, 0x49,
	0x82, 0x49, 0x91, 0x69, 0xc3, 0x01, 0xd4, 0x54, 0xe9, 0x83, 0x75, 0x7f, 0x5c, 0xed, 0x3e, 0xe6,
	0x6a, 0x16, 0xe0, 0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xaf, 0x83, 0xb1, 0x6d, 0x01,
	0x00, 0x00,
}