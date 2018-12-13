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
	proto.RegisterEnum("SpanType", SpanType_name, SpanType_value)
}

func init() { proto.RegisterFile("Common.proto", fileDescriptor_ee72d9a89737215c) }

var fileDescriptor_ee72d9a89737215c = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xd2, 0xe2, 0xe2, 0x08, 0x2e, 0x48, 0xcc,
	0x0b, 0xa9, 0x2c, 0x48, 0x15, 0xe2, 0xe4, 0x62, 0x75, 0xcd, 0x2b, 0x29, 0xaa, 0x14, 0x60, 0x10,
	0xe2, 0xe0, 0x62, 0x71, 0xad, 0xc8, 0x2c, 0x11, 0x60, 0x04, 0x09, 0xfa, 0xe4, 0x27, 0x27, 0xe6,
	0x08, 0x30, 0x39, 0x79, 0x70, 0xa9, 0xe7, 0x17, 0xa5, 0xeb, 0x25, 0x16, 0x24, 0x26, 0x67, 0xa4,
	0xea, 0x15, 0x67, 0x57, 0x96, 0x27, 0xe6, 0x64, 0x67, 0xe6, 0x81, 0x44, 0x72, 0xf5, 0xf2, 0x52,
	0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0x21, 0xc6, 0x06, 0x30, 0xae, 0x62, 0x92, 0x0a, 0xce, 0xae, 0x0c,
	0x87, 0x2a, 0xf0, 0x83, 0x48, 0x06, 0x80, 0xe4, 0x92, 0xf3, 0x73, 0x92, 0xd8, 0xc0, 0xaa, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x7b, 0x0b, 0x6c, 0x8c, 0x00, 0x00, 0x00,
}
