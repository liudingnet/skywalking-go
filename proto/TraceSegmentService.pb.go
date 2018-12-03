// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TraceSegmentService.proto

package skywalking_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RefType int32

const (
	RefType_CrossProcess RefType = 0
	RefType_CrossThread  RefType = 1
)

var RefType_name = map[int32]string{
	0: "CrossProcess",
	1: "CrossThread",
}

var RefType_value = map[string]int32{
	"CrossProcess": 0,
	"CrossThread":  1,
}

func (x RefType) String() string {
	return proto.EnumName(RefType_name, int32(x))
}

func (RefType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{0}
}

type SpanLayer int32

const (
	SpanLayer_Unknown      SpanLayer = 0
	SpanLayer_Database     SpanLayer = 1
	SpanLayer_RPCFramework SpanLayer = 2
	SpanLayer_Http         SpanLayer = 3
	SpanLayer_MQ           SpanLayer = 4
	SpanLayer_Cache        SpanLayer = 5
)

var SpanLayer_name = map[int32]string{
	0: "Unknown",
	1: "Database",
	2: "RPCFramework",
	3: "Http",
	4: "MQ",
	5: "Cache",
}

var SpanLayer_value = map[string]int32{
	"Unknown":      0,
	"Database":     1,
	"RPCFramework": 2,
	"Http":         3,
	"MQ":           4,
	"Cache":        5,
}

func (x SpanLayer) String() string {
	return proto.EnumName(SpanLayer_name, int32(x))
}

func (SpanLayer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{1}
}

type UpstreamSegment struct {
	GlobalTraceIds       []*UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment              []byte      `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpstreamSegment) Reset()         { *m = UpstreamSegment{} }
func (m *UpstreamSegment) String() string { return proto.CompactTextString(m) }
func (*UpstreamSegment) ProtoMessage()    {}
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{0}
}

func (m *UpstreamSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSegment.Unmarshal(m, b)
}
func (m *UpstreamSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSegment.Marshal(b, m, deterministic)
}
func (m *UpstreamSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSegment.Merge(m, src)
}
func (m *UpstreamSegment) XXX_Size() int {
	return xxx_messageInfo_UpstreamSegment.Size(m)
}
func (m *UpstreamSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSegment.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSegment proto.InternalMessageInfo

func (m *UpstreamSegment) GetGlobalTraceIds() []*UniqueId {
	if m != nil {
		return m.GlobalTraceIds
	}
	return nil
}

func (m *UpstreamSegment) GetSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type UniqueId struct {
	IdParts              []int64  `protobuf:"varint,1,rep,packed,name=idParts,proto3" json:"idParts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueId) Reset()         { *m = UniqueId{} }
func (m *UniqueId) String() string { return proto.CompactTextString(m) }
func (*UniqueId) ProtoMessage()    {}
func (*UniqueId) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{1}
}

func (m *UniqueId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueId.Unmarshal(m, b)
}
func (m *UniqueId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueId.Marshal(b, m, deterministic)
}
func (m *UniqueId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueId.Merge(m, src)
}
func (m *UniqueId) XXX_Size() int {
	return xxx_messageInfo_UniqueId.Size(m)
}
func (m *UniqueId) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueId.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueId proto.InternalMessageInfo

func (m *UniqueId) GetIdParts() []int64 {
	if m != nil {
		return m.IdParts
	}
	return nil
}

type TraceSegmentObject struct {
	TraceSegmentId        *UniqueId     `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans                 []*SpanObject `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ApplicationId         int32         `protobuf:"varint,3,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApplicationInstanceId int32         `protobuf:"varint,4,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	IsSizeLimited         bool          `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *TraceSegmentObject) Reset()         { *m = TraceSegmentObject{} }
func (m *TraceSegmentObject) String() string { return proto.CompactTextString(m) }
func (*TraceSegmentObject) ProtoMessage()    {}
func (*TraceSegmentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{2}
}

func (m *TraceSegmentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceSegmentObject.Unmarshal(m, b)
}
func (m *TraceSegmentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceSegmentObject.Marshal(b, m, deterministic)
}
func (m *TraceSegmentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceSegmentObject.Merge(m, src)
}
func (m *TraceSegmentObject) XXX_Size() int {
	return xxx_messageInfo_TraceSegmentObject.Size(m)
}
func (m *TraceSegmentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceSegmentObject.DiscardUnknown(m)
}

var xxx_messageInfo_TraceSegmentObject proto.InternalMessageInfo

func (m *TraceSegmentObject) GetTraceSegmentId() *UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *TraceSegmentObject) GetSpans() []*SpanObject {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *TraceSegmentObject) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *TraceSegmentObject) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentObject) GetIsSizeLimited() bool {
	if m != nil {
		return m.IsSizeLimited
	}
	return false
}

type TraceSegmentReference struct {
	RefType                     RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=skywalking_proto.RefType" json:"refType,omitempty"`
	ParentTraceSegmentId        *UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId                int32     `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentApplicationInstanceId int32     `protobuf:"varint,4,opt,name=parentApplicationInstanceId,proto3" json:"parentApplicationInstanceId,omitempty"`
	NetworkAddress              string    `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId            int32     `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryApplicationInstanceId  int32     `protobuf:"varint,7,opt,name=entryApplicationInstanceId,proto3" json:"entryApplicationInstanceId,omitempty"`
	EntryServiceName            string    `protobuf:"bytes,8,opt,name=entryServiceName,proto3" json:"entryServiceName,omitempty"`
	EntryServiceId              int32     `protobuf:"varint,9,opt,name=entryServiceId,proto3" json:"entryServiceId,omitempty"`
	ParentServiceName           string    `protobuf:"bytes,10,opt,name=parentServiceName,proto3" json:"parentServiceName,omitempty"`
	ParentServiceId             int32     `protobuf:"varint,11,opt,name=parentServiceId,proto3" json:"parentServiceId,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}  `json:"-"`
	XXX_unrecognized            []byte    `json:"-"`
	XXX_sizecache               int32     `json:"-"`
}

func (m *TraceSegmentReference) Reset()         { *m = TraceSegmentReference{} }
func (m *TraceSegmentReference) String() string { return proto.CompactTextString(m) }
func (*TraceSegmentReference) ProtoMessage()    {}
func (*TraceSegmentReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{3}
}

func (m *TraceSegmentReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceSegmentReference.Unmarshal(m, b)
}
func (m *TraceSegmentReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceSegmentReference.Marshal(b, m, deterministic)
}
func (m *TraceSegmentReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceSegmentReference.Merge(m, src)
}
func (m *TraceSegmentReference) XXX_Size() int {
	return xxx_messageInfo_TraceSegmentReference.Size(m)
}
func (m *TraceSegmentReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceSegmentReference.DiscardUnknown(m)
}

var xxx_messageInfo_TraceSegmentReference proto.InternalMessageInfo

func (m *TraceSegmentReference) GetRefType() RefType {
	if m != nil {
		return m.RefType
	}
	return RefType_CrossProcess
}

func (m *TraceSegmentReference) GetParentTraceSegmentId() *UniqueId {
	if m != nil {
		return m.ParentTraceSegmentId
	}
	return nil
}

func (m *TraceSegmentReference) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *TraceSegmentReference) GetParentApplicationInstanceId() int32 {
	if m != nil {
		return m.ParentApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentReference) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *TraceSegmentReference) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func (m *TraceSegmentReference) GetEntryApplicationInstanceId() int32 {
	if m != nil {
		return m.EntryApplicationInstanceId
	}
	return 0
}

func (m *TraceSegmentReference) GetEntryServiceName() string {
	if m != nil {
		return m.EntryServiceName
	}
	return ""
}

func (m *TraceSegmentReference) GetEntryServiceId() int32 {
	if m != nil {
		return m.EntryServiceId
	}
	return 0
}

func (m *TraceSegmentReference) GetParentServiceName() string {
	if m != nil {
		return m.ParentServiceName
	}
	return ""
}

func (m *TraceSegmentReference) GetParentServiceId() int32 {
	if m != nil {
		return m.ParentServiceId
	}
	return 0
}

type SpanObject struct {
	SpanId               int32                    `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId         int32                    `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime            int64                    `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64                    `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs                 []*TraceSegmentReference `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId      int32                    `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName        string                   `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId               int32                    `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer                 string                   `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType             SpanType                 `protobuf:"varint,10,opt,name=spanType,proto3,enum=skywalking_proto.SpanType" json:"spanType,omitempty"`
	SpanLayer            SpanLayer                `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=skywalking_proto.SpanLayer" json:"spanLayer,omitempty"`
	ComponentId          int32                    `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component            string                   `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError              bool                     `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags                 []*KeyWithStringValue    `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*LogMessage            `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SpanObject) Reset()         { *m = SpanObject{} }
func (m *SpanObject) String() string { return proto.CompactTextString(m) }
func (*SpanObject) ProtoMessage()    {}
func (*SpanObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{4}
}

func (m *SpanObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanObject.Unmarshal(m, b)
}
func (m *SpanObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanObject.Marshal(b, m, deterministic)
}
func (m *SpanObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanObject.Merge(m, src)
}
func (m *SpanObject) XXX_Size() int {
	return xxx_messageInfo_SpanObject.Size(m)
}
func (m *SpanObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanObject.DiscardUnknown(m)
}

var xxx_messageInfo_SpanObject proto.InternalMessageInfo

func (m *SpanObject) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanObject) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SpanObject) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *SpanObject) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *SpanObject) GetRefs() []*TraceSegmentReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *SpanObject) GetOperationNameId() int32 {
	if m != nil {
		return m.OperationNameId
	}
	return 0
}

func (m *SpanObject) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *SpanObject) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SpanObject) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *SpanObject) GetSpanType() SpanType {
	if m != nil {
		return m.SpanType
	}
	return SpanType_Entry
}

func (m *SpanObject) GetSpanLayer() SpanLayer {
	if m != nil {
		return m.SpanLayer
	}
	return SpanLayer_Unknown
}

func (m *SpanObject) GetComponentId() int32 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

func (m *SpanObject) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SpanObject) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SpanObject) GetTags() []*KeyWithStringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SpanObject) GetLogs() []*LogMessage {
	if m != nil {
		return m.Logs
	}
	return nil
}

type LogMessage struct {
	Time                 int64                 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []*KeyWithStringValue `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc055cab3bd8ed31, []int{5}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogMessage) GetData() []*KeyWithStringValue {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("skywalking_proto.RefType", RefType_name, RefType_value)
	proto.RegisterEnum("skywalking_proto.SpanLayer", SpanLayer_name, SpanLayer_value)
	proto.RegisterType((*UpstreamSegment)(nil), "skywalking_proto.UpstreamSegment")
	proto.RegisterType((*UniqueId)(nil), "skywalking_proto.UniqueId")
	proto.RegisterType((*TraceSegmentObject)(nil), "skywalking_proto.TraceSegmentObject")
	proto.RegisterType((*TraceSegmentReference)(nil), "skywalking_proto.TraceSegmentReference")
	proto.RegisterType((*SpanObject)(nil), "skywalking_proto.SpanObject")
	proto.RegisterType((*LogMessage)(nil), "skywalking_proto.LogMessage")
}

func init() { proto.RegisterFile("TraceSegmentService.proto", fileDescriptor_bc055cab3bd8ed31) }

var fileDescriptor_bc055cab3bd8ed31 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x6f, 0x8f, 0xdb, 0xc4,
	0x13, 0x3e, 0x27, 0xce, 0x25, 0x99, 0xa4, 0x39, 0xff, 0xf6, 0x47, 0x91, 0x9b, 0xf6, 0x45, 0x88,
	0x4e, 0x34, 0x3a, 0x55, 0x11, 0x4a, 0x11, 0x02, 0x21, 0x21, 0xda, 0x2b, 0xa8, 0x11, 0xd7, 0x23,
	0x38, 0x39, 0x2a, 0xf1, 0x06, 0xed, 0xd9, 0x73, 0x39, 0x93, 0x78, 0xd7, 0xec, 0x6e, 0x39, 0x85,
	0x0f, 0xc1, 0x07, 0xe1, 0x63, 0xf0, 0x91, 0xf8, 0x04, 0x68, 0xc7, 0xce, 0x1f, 0x27, 0xe6, 0xca,
	0x3b, 0xcf, 0x33, 0xcf, 0xcc, 0x8e, 0x9f, 0xd9, 0x99, 0x85, 0x47, 0x33, 0xc5, 0x43, 0x9c, 0xe2,
	0x3c, 0x41, 0x61, 0xa6, 0xa8, 0x7e, 0x8b, 0x43, 0x1c, 0xa6, 0x4a, 0x1a, 0xc9, 0x3c, 0xbd, 0x58,
	0xdd, 0xf1, 0xe5, 0x22, 0x16, 0xf3, 0x9f, 0x09, 0xe9, 0xb6, 0xcf, 0x65, 0x92, 0x48, 0x91, 0xf9,
	0xbb, 0xde, 0x2b, 0x79, 0x27, 0xb4, 0x51, 0xc8, 0x93, 0x1c, 0xf1, 0xbf, 0xc3, 0xd5, 0xdb, 0xd8,
	0xdc, 0x4e, 0x8d, 0x8a, 0xc5, 0xfc, 0x47, 0xbe, 0x7c, 0x97, 0xe7, 0xea, 0x4b, 0x38, 0xb9, 0x4a,
	0x33, 0x6e, 0x7e, 0x16, 0x7b, 0x09, 0x9d, 0xf9, 0x52, 0x5e, 0xf3, 0x25, 0x55, 0x30, 0x8e, 0xb4,
	0xef, 0xf4, 0xaa, 0x83, 0xd6, 0xa8, 0x3b, 0xdc, 0x3f, 0x77, 0x78, 0x25, 0xe2, 0x5f, 0xdf, 0xe1,
	0x38, 0x0a, 0xf6, 0x22, 0x98, 0x0f, 0x75, 0x9d, 0xa5, 0xf3, 0x2b, 0x3d, 0x67, 0xd0, 0x0e, 0xd6,
	0x66, 0xff, 0x14, 0x1a, 0xeb, 0x28, 0xcb, 0x8a, 0xa3, 0x09, 0x57, 0x26, 0x3b, 0xa2, 0x1a, 0xac,
	0xcd, 0xfe, 0x1f, 0x15, 0x60, 0xbb, 0x02, 0x7c, 0x7f, 0xfd, 0x0b, 0x86, 0x54, 0x9a, 0xd9, 0x41,
	0xc7, 0x91, 0xef, 0xf4, 0x9c, 0xf7, 0x95, 0x56, 0x8c, 0x60, 0x23, 0xa8, 0xe9, 0x94, 0x0b, 0xed,
	0x57, 0xe8, 0xaf, 0x9e, 0x1c, 0x86, 0x4e, 0x53, 0x2e, 0xb2, 0x03, 0x83, 0x8c, 0xca, 0x4e, 0xe1,
	0x01, 0x4f, 0xd3, 0x65, 0x1c, 0x72, 0x13, 0x4b, 0x31, 0x8e, 0xfc, 0x6a, 0xcf, 0x19, 0xd4, 0x82,
	0x22, 0xc8, 0x3e, 0x85, 0x87, 0xbb, 0x80, 0xd0, 0x86, 0x0b, 0x2b, 0x87, 0xef, 0x12, 0xbb, 0xdc,
	0x69, 0x73, 0xc7, 0x7a, 0x1a, 0xff, 0x8e, 0x17, 0x71, 0x12, 0x1b, 0x8c, 0xfc, 0x5a, 0xcf, 0x19,
	0x34, 0x82, 0x22, 0xd8, 0xff, 0xcb, 0x85, 0x87, 0xbb, 0x82, 0x04, 0x78, 0x83, 0x0a, 0x45, 0x88,
	0xec, 0x39, 0xd4, 0x15, 0xde, 0xcc, 0x56, 0x29, 0x92, 0x18, 0x9d, 0xd1, 0xa3, 0xc3, 0x3f, 0x0a,
	0x32, 0x42, 0xb0, 0x66, 0xb2, 0x4b, 0xf8, 0x20, 0xe5, 0x0a, 0x85, 0x99, 0x15, 0xe5, 0xac, 0xbc,
	0x57, 0xce, 0xd2, 0x38, 0xd6, 0x87, 0x76, 0x86, 0x5b, 0xed, 0x36, 0xfa, 0x14, 0x30, 0xf6, 0x35,
	0x3c, 0xce, 0xec, 0x17, 0xf7, 0x88, 0x74, 0x1f, 0x85, 0x7d, 0x0c, 0x1d, 0x81, 0xe6, 0x4e, 0xaa,
	0xc5, 0x8b, 0x28, 0x52, 0xa8, 0x35, 0x69, 0xd5, 0x0c, 0xf6, 0x50, 0x76, 0x06, 0x5e, 0x11, 0x19,
	0x47, 0xfe, 0x31, 0xa5, 0x3f, 0xc0, 0xd9, 0x57, 0xd0, 0x45, 0x61, 0xd4, 0xaa, 0xbc, 0xa8, 0x3a,
	0x45, 0xdd, 0xc3, 0xb0, 0x67, 0x91, 0x37, 0x1f, 0xd1, 0x4b, 0x9e, 0xa0, 0xdf, 0xa0, 0xaa, 0x0e,
	0x70, 0x5b, 0xff, 0x2e, 0x36, 0x8e, 0xfc, 0x26, 0xe5, 0xdf, 0x43, 0xd9, 0x33, 0xf8, 0x5f, 0xae,
	0xdc, 0x4e, 0x52, 0xa0, 0xa4, 0x87, 0x0e, 0x36, 0x80, 0x93, 0x02, 0x38, 0x8e, 0xfc, 0x16, 0xa5,
	0xdd, 0x87, 0xfb, 0x7f, 0xbb, 0x00, 0xdb, 0xcb, 0xcd, 0x3e, 0x84, 0x63, 0x9d, 0xb5, 0xcb, 0x21,
	0x7e, 0x6e, 0x1d, 0x34, 0xb3, 0x52, 0xd2, 0xcc, 0x27, 0xd0, 0xd4, 0x86, 0x2b, 0x33, 0x8b, 0x13,
	0xa4, 0x6e, 0x57, 0x83, 0x2d, 0x60, 0x07, 0x1b, 0x45, 0x44, 0x3e, 0x97, 0x7c, 0x6b, 0x93, 0x7d,
	0x09, 0xae, 0xc2, 0x1b, 0xdb, 0x38, 0x3b, 0x7c, 0x4f, 0x0f, 0x2f, 0x5a, 0xe9, 0x25, 0x0f, 0x28,
	0xc8, 0xfe, 0xa9, 0x4c, 0x51, 0x51, 0x0b, 0xec, 0xaf, 0x6f, 0xda, 0xba, 0x0f, 0xdb, 0xa1, 0x2a,
	0x40, 0xd4, 0xc8, 0x66, 0x50, 0x04, 0xad, 0x00, 0x29, 0xa2, 0x1a, 0x47, 0xd4, 0xb1, 0x5a, 0x90,
	0x5b, 0x8c, 0x81, 0x6b, 0xbf, 0xa8, 0x3b, 0xcd, 0x80, 0xbe, 0xd9, 0x67, 0xd0, 0xb0, 0xf2, 0xd0,
	0x9c, 0x01, 0xcd, 0x59, 0xb7, 0x7c, 0x73, 0xd0, 0xa0, 0x6d, 0xb8, 0xec, 0x0b, 0x68, 0xda, 0xef,
	0x0b, 0xbe, 0x42, 0x45, 0x7d, 0xe9, 0x8c, 0x1e, 0x97, 0x07, 0x12, 0x25, 0xd8, 0xb2, 0x59, 0x0f,
	0x5a, 0xa1, 0x4c, 0x52, 0x29, 0xb2, 0xd9, 0x6c, 0x53, 0x8d, 0xbb, 0x90, 0xed, 0xc2, 0xc6, 0xf4,
	0x1f, 0x50, 0xb5, 0x5b, 0x80, 0xd6, 0xab, 0xfe, 0x46, 0x29, 0xa9, 0xfc, 0x0e, 0xed, 0x94, 0xb5,
	0xc9, 0x3e, 0x07, 0xd7, 0xf0, 0xb9, 0xf6, 0x4f, 0xa8, 0x0b, 0xa7, 0x87, 0xf5, 0x1c, 0xbe, 0x17,
	0x01, 0x45, 0xb0, 0x4f, 0xc0, 0x5d, 0xca, 0xb9, 0xf6, 0xbd, 0x7f, 0x5b, 0x9e, 0x17, 0x72, 0xfe,
	0x06, 0xb5, 0xe6, 0x73, 0x0c, 0x88, 0xd9, 0xff, 0x09, 0x60, 0x8b, 0x59, 0x69, 0x8d, 0xbd, 0x16,
	0x0e, 0x5d, 0x0b, 0xfa, 0xb6, 0xd5, 0x44, 0xdc, 0xf0, 0x7c, 0x21, 0xff, 0xc7, 0x6a, 0x6c, 0xc4,
	0xd9, 0x33, 0xa8, 0xe7, 0xab, 0x8d, 0x79, 0xd0, 0x3e, 0x57, 0x52, 0xeb, 0x89, 0x92, 0x21, 0x6a,
	0xed, 0x1d, 0xb1, 0x13, 0x68, 0x11, 0x32, 0xbb, 0x55, 0xc8, 0x23, 0xcf, 0x39, 0xbb, 0x82, 0xe6,
	0x46, 0x67, 0xd6, 0x82, 0xfa, 0x95, 0x58, 0x08, 0x79, 0x27, 0xbc, 0x23, 0xd6, 0x86, 0xc6, 0x2b,
	0x6e, 0xf8, 0x35, 0xd7, 0xe8, 0x39, 0x36, 0x55, 0x30, 0x39, 0xff, 0x56, 0xf1, 0x04, 0xed, 0xae,
	0xf0, 0x2a, 0xac, 0x01, 0xee, 0x6b, 0x63, 0x52, 0xaf, 0xca, 0x8e, 0xa1, 0xf2, 0xe6, 0x07, 0xcf,
	0x65, 0x4d, 0xa8, 0x9d, 0xf3, 0xf0, 0x16, 0xbd, 0xda, 0x08, 0xe1, 0xff, 0x25, 0x6f, 0x35, 0xbb,
	0x84, 0x7a, 0x28, 0x97, 0x4b, 0x3b, 0x68, 0x1f, 0x95, 0xec, 0xd3, 0xe2, 0xa3, 0xdb, 0x2d, 0x51,
	0x72, 0xfb, 0x8a, 0xf7, 0x8f, 0x06, 0xce, 0xcb, 0xd7, 0xf0, 0x54, 0xaa, 0xf9, 0x90, 0xa7, 0xf6,
	0xd8, 0x1d, 0xf6, 0x90, 0xa7, 0xc9, 0x30, 0x5f, 0x6b, 0xd9, 0xa3, 0x3e, 0x71, 0xfe, 0xac, 0x74,
	0xa7, 0x8b, 0xd5, 0xdb, 0x9c, 0x70, 0x99, 0x39, 0x27, 0xd6, 0x17, 0xca, 0xe5, 0xf5, 0x31, 0xb1,
	0x9e, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd5, 0x60, 0x54, 0x63, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceSegmentServiceClient is the client API for TraceSegmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSegmentServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error)
}

type traceSegmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceSegmentServiceClient(cc *grpc.ClientConn) TraceSegmentServiceClient {
	return &traceSegmentServiceClient{cc}
}

func (c *traceSegmentServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceSegmentService_serviceDesc.Streams[0], "/skywalking_proto.TraceSegmentService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentService_CollectClient interface {
	Send(*UpstreamSegment) error
	CloseAndRecv() (*Downstream, error)
	grpc.ClientStream
}

type traceSegmentServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentServiceCollectClient) Send(m *UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectClient) CloseAndRecv() (*Downstream, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Downstream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentServiceServer is the server API for TraceSegmentService service.
type TraceSegmentServiceServer interface {
	Collect(TraceSegmentService_CollectServer) error
}

func RegisterTraceSegmentServiceServer(s *grpc.Server, srv TraceSegmentServiceServer) {
	s.RegisterService(&_TraceSegmentService_serviceDesc, srv)
}

func _TraceSegmentService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentServiceServer).Collect(&traceSegmentServiceCollectServer{stream})
}

type TraceSegmentService_CollectServer interface {
	SendAndClose(*Downstream) error
	Recv() (*UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentServiceCollectServer) SendAndClose(m *Downstream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectServer) Recv() (*UpstreamSegment, error) {
	m := new(UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceSegmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking_proto.TraceSegmentService",
	HandlerType: (*TraceSegmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "TraceSegmentService.proto",
}
