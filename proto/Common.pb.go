// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Common.proto

package skywalking_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpanType int32

const (
	SpanType_Entry SpanType = 0
	SpanType_Exit  SpanType = 1
	SpanType_Local SpanType = 2
)

var SpanType_name = map[int32]string{
	0: "Entry",
	1: "Exit",
	2: "Local",
}

var SpanType_value = map[string]int32{
	"Entry": 0,
	"Exit":  1,
	"Local": 2,
}

func (x SpanType) String() string {
	return proto.EnumName(SpanType_name, int32(x))
}

func (SpanType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{0}
}

func init() {
	proto.RegisterEnum("skywalking_proto.SpanType", SpanType_name, SpanType_value)
}

func init() { proto.RegisterFile("Common.proto", fileDescriptor_ee72d9a89737215c) }

var fileDescriptor_ee72d9a89737215c = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0xce, 0xae, 0x2c, 0x4f, 0xcc,
	0xc9, 0xce, 0xcc, 0x4b, 0x8f, 0x07, 0x8b, 0x68, 0x69, 0x71, 0x71, 0x04, 0x17, 0x24, 0xe6, 0x85,
	0x54, 0x16, 0xa4, 0x0a, 0x71, 0x72, 0xb1, 0xba, 0xe6, 0x95, 0x14, 0x55, 0x0a, 0x30, 0x08, 0x71,
	0x70, 0xb1, 0xb8, 0x56, 0x64, 0x96, 0x08, 0x30, 0x82, 0x04, 0x7d, 0xf2, 0x93, 0x13, 0x73, 0x04,
	0x98, 0x9c, 0x3c, 0xb8, 0xd4, 0xf3, 0x8b, 0xd2, 0xf5, 0x12, 0x0b, 0x12, 0x93, 0x33, 0x52, 0xf5,
	0x10, 0x46, 0xe9, 0x25, 0x16, 0xe4, 0xea, 0xe5, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0x43, 0x2c,
	0x0a, 0x60, 0x5c, 0xc5, 0x24, 0x15, 0x9c, 0x5d, 0x19, 0x0e, 0x55, 0xe0, 0x07, 0x91, 0x0c, 0x00,
	0xc9, 0x25, 0xe7, 0xe7, 0x24, 0xb1, 0x81, 0x55, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x40, 0x4f, 0xf0, 0x9e, 0x00, 0x00, 0x00,
}