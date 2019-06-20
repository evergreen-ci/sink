// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buildlogger.proto

package internal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LogFormat int32

const (
	LogFormat_LOG_FORMAT_UNKNOWN LogFormat = 0
	LogFormat_LOG_FORMAT_TEXT    LogFormat = 1
	LogFormat_LOG_FORMAT_JSON    LogFormat = 2
	LogFormat_LOG_FORMAT_BSON    LogFormat = 3
)

var LogFormat_name = map[int32]string{
	0: "LOG_FORMAT_UNKNOWN",
	1: "LOG_FORMAT_TEXT",
	2: "LOG_FORMAT_JSON",
	3: "LOG_FORMAT_BSON",
}

var LogFormat_value = map[string]int32{
	"LOG_FORMAT_UNKNOWN": 0,
	"LOG_FORMAT_TEXT":    1,
	"LOG_FORMAT_JSON":    2,
	"LOG_FORMAT_BSON":    3,
}

func (x LogFormat) String() string {
	return proto.EnumName(LogFormat_name, int32(x))
}

func (LogFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{0}
}

type LogInfo struct {
	Project              string               `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Version              string               `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variant              string               `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TaskName             string               `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId               string               `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Execution            int32                `protobuf:"varint,6,opt,name=execution,proto3" json:"execution,omitempty"`
	TestName             string               `protobuf:"bytes,7,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	Trial                int32                `protobuf:"varint,8,opt,name=trial,proto3" json:"trial,omitempty"`
	ProcName             string               `protobuf:"bytes,9,opt,name=proc_name,json=procName,proto3" json:"proc_name,omitempty"`
	Format               LogFormat            `protobuf:"varint,10,opt,name=format,proto3,enum=cedar.LogFormat" json:"format,omitempty"`
	Arguments            map[string]string    `protobuf:"bytes,11,rep,name=arguments,proto3" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mainline             bool                 `protobuf:"varint,12,opt,name=mainline,proto3" json:"mainline,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{0}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *LogInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LogInfo) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *LogInfo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *LogInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *LogInfo) GetExecution() int32 {
	if m != nil {
		return m.Execution
	}
	return 0
}

func (m *LogInfo) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *LogInfo) GetTrial() int32 {
	if m != nil {
		return m.Trial
	}
	return 0
}

func (m *LogInfo) GetProcName() string {
	if m != nil {
		return m.ProcName
	}
	return ""
}

func (m *LogInfo) GetFormat() LogFormat {
	if m != nil {
		return m.Format
	}
	return LogFormat_LOG_FORMAT_UNKNOWN
}

func (m *LogInfo) GetArguments() map[string]string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *LogInfo) GetMainline() bool {
	if m != nil {
		return m.Mainline
	}
	return false
}

func (m *LogInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type LogLines struct {
	Lines                []*LogLine `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LogLines) Reset()         { *m = LogLines{} }
func (m *LogLines) String() string { return proto.CompactTextString(m) }
func (*LogLines) ProtoMessage()    {}
func (*LogLines) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{1}
}

func (m *LogLines) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLines.Unmarshal(m, b)
}
func (m *LogLines) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLines.Marshal(b, m, deterministic)
}
func (m *LogLines) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLines.Merge(m, src)
}
func (m *LogLines) XXX_Size() int {
	return xxx_messageInfo_LogLines.Size(m)
}
func (m *LogLines) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLines.DiscardUnknown(m)
}

var xxx_messageInfo_LogLines proto.InternalMessageInfo

func (m *LogLines) GetLines() []*LogLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

type LogLine struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data                 string               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogLine) Reset()         { *m = LogLine{} }
func (m *LogLine) String() string { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()    {}
func (*LogLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{2}
}

func (m *LogLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLine.Unmarshal(m, b)
}
func (m *LogLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLine.Marshal(b, m, deterministic)
}
func (m *LogLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLine.Merge(m, src)
}
func (m *LogLine) XXX_Size() int {
	return xxx_messageInfo_LogLine.Size(m)
}
func (m *LogLine) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLine.DiscardUnknown(m)
}

var xxx_messageInfo_LogLine proto.InternalMessageInfo

func (m *LogLine) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogLine) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type LogEndInfo struct {
	ExitCode             int32                `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	CompletedAt          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LogEndInfo) Reset()         { *m = LogEndInfo{} }
func (m *LogEndInfo) String() string { return proto.CompactTextString(m) }
func (*LogEndInfo) ProtoMessage()    {}
func (*LogEndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{3}
}

func (m *LogEndInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEndInfo.Unmarshal(m, b)
}
func (m *LogEndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEndInfo.Marshal(b, m, deterministic)
}
func (m *LogEndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEndInfo.Merge(m, src)
}
func (m *LogEndInfo) XXX_Size() int {
	return xxx_messageInfo_LogEndInfo.Size(m)
}
func (m *LogEndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogEndInfo proto.InternalMessageInfo

func (m *LogEndInfo) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *LogEndInfo) GetCompletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CompletedAt
	}
	return nil
}

type BuildloggerResponse struct {
	LogId                string   `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildloggerResponse) Reset()         { *m = BuildloggerResponse{} }
func (m *BuildloggerResponse) String() string { return proto.CompactTextString(m) }
func (*BuildloggerResponse) ProtoMessage()    {}
func (*BuildloggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f5c52c3a3ee6d6, []int{4}
}

func (m *BuildloggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildloggerResponse.Unmarshal(m, b)
}
func (m *BuildloggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildloggerResponse.Marshal(b, m, deterministic)
}
func (m *BuildloggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildloggerResponse.Merge(m, src)
}
func (m *BuildloggerResponse) XXX_Size() int {
	return xxx_messageInfo_BuildloggerResponse.Size(m)
}
func (m *BuildloggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildloggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuildloggerResponse proto.InternalMessageInfo

func (m *BuildloggerResponse) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *BuildloggerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("cedar.LogFormat", LogFormat_name, LogFormat_value)
	proto.RegisterType((*LogInfo)(nil), "cedar.LogInfo")
	proto.RegisterMapType((map[string]string)(nil), "cedar.LogInfo.ArgumentsEntry")
	proto.RegisterType((*LogLines)(nil), "cedar.LogLines")
	proto.RegisterType((*LogLine)(nil), "cedar.LogLine")
	proto.RegisterType((*LogEndInfo)(nil), "cedar.LogEndInfo")
	proto.RegisterType((*BuildloggerResponse)(nil), "cedar.BuildloggerResponse")
}

func init() { proto.RegisterFile("buildlogger.proto", fileDescriptor_c4f5c52c3a3ee6d6) }

var fileDescriptor_c4f5c52c3a3ee6d6 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x7d, 0x4e, 0x48, 0x62, 0xdf, 0xf0, 0x42, 0x18, 0x78, 0xef, 0x59, 0xe1, 0x55, 0x8d, 0xa2,
	0x2e, 0xac, 0x2e, 0x42, 0x95, 0x2e, 0x0a, 0x6d, 0x59, 0x04, 0x14, 0x2a, 0x5a, 0x37, 0x91, 0x4c,
	0x2a, 0xaa, 0x6e, 0xa2, 0x89, 0x7d, 0x71, 0x5d, 0xec, 0x19, 0x6b, 0x66, 0x82, 0xca, 0x2f, 0xe8,
	0x4f, 0xee, 0xb6, 0x9a, 0xb1, 0x93, 0x00, 0xaa, 0x60, 0xd1, 0x9d, 0xcf, 0x39, 0xf7, 0x6b, 0xee,
	0xb9, 0x86, 0xed, 0xf9, 0x22, 0x49, 0xa3, 0x94, 0xc7, 0x31, 0x8a, 0x7e, 0x2e, 0xb8, 0xe2, 0xa4,
	0x16, 0x62, 0x44, 0x45, 0xe7, 0x69, 0xcc, 0x79, 0x9c, 0xe2, 0xbe, 0x21, 0xe7, 0x8b, 0xcb, 0x7d,
	0x95, 0x64, 0x28, 0x15, 0xcd, 0xf2, 0x22, 0xae, 0xf7, 0xb3, 0x0a, 0x0d, 0x9f, 0xc7, 0x67, 0xec,
	0x92, 0x13, 0x17, 0x1a, 0xb9, 0xe0, 0xdf, 0x30, 0x54, 0xae, 0xd5, 0xb5, 0x3c, 0x27, 0x58, 0x42,
	0xad, 0x5c, 0xa3, 0x90, 0x09, 0x67, 0x6e, 0xa5, 0x50, 0x4a, 0x68, 0x14, 0x2a, 0x12, 0xca, 0x94,
	0x5b, 0x2d, 0x95, 0x02, 0x92, 0x3d, 0x70, 0x14, 0x95, 0x57, 0x33, 0x46, 0x33, 0x74, 0x37, 0x8c,
	0x66, 0x6b, 0x62, 0x4c, 0x33, 0x24, 0xff, 0x41, 0xc3, 0x88, 0x49, 0xe4, 0xd6, 0x8c, 0x54, 0xd7,
	0xf0, 0x2c, 0x22, 0xff, 0x83, 0x83, 0xdf, 0x31, 0x5c, 0x28, 0xdd, 0xab, 0xde, 0xb5, 0xbc, 0x5a,
	0xb0, 0x26, 0x4c, 0x4d, 0x94, 0xaa, 0xa8, 0xd9, 0x28, 0x6b, 0xa2, 0x54, 0xa6, 0xe6, 0x2e, 0xd4,
	0x94, 0x48, 0x68, 0xea, 0xda, 0x26, 0xad, 0x00, 0x3a, 0x25, 0x17, 0x3c, 0x2c, 0x52, 0x9c, 0x22,
	0x45, 0x13, 0x26, 0xc5, 0x83, 0xfa, 0x25, 0x17, 0x19, 0x55, 0x2e, 0x74, 0x2d, 0xaf, 0x35, 0x68,
	0xf7, 0xcd, 0xda, 0xfa, 0x3e, 0x8f, 0x4f, 0x0d, 0x1f, 0x94, 0x3a, 0x79, 0x03, 0x0e, 0x15, 0xf1,
	0x22, 0x43, 0xa6, 0xa4, 0xdb, 0xec, 0x56, 0xbd, 0xe6, 0xe0, 0xc9, 0x3a, 0x58, 0xaf, 0xaf, 0x3f,
	0x5c, 0xea, 0x23, 0xa6, 0xc4, 0x4d, 0xb0, 0x8e, 0x27, 0x1d, 0xb0, 0x33, 0x9a, 0xb0, 0x34, 0x61,
	0xe8, 0x6e, 0x76, 0x2d, 0xcf, 0x0e, 0x56, 0x98, 0x1c, 0x02, 0x84, 0x02, 0xa9, 0xc2, 0x68, 0x46,
	0x95, 0xfb, 0x77, 0xd7, 0xf2, 0x9a, 0x83, 0x4e, 0xbf, 0xb0, 0xad, 0xbf, 0xb4, 0xad, 0x3f, 0x5d,
	0xda, 0x16, 0x38, 0x65, 0xf4, 0x50, 0x75, 0xde, 0x42, 0xeb, 0x6e, 0x4f, 0xd2, 0x86, 0xea, 0x15,
	0xde, 0x94, 0xee, 0xe9, 0x4f, 0xbd, 0x94, 0x6b, 0x9a, 0x2e, 0xb0, 0xf4, 0xad, 0x00, 0xaf, 0x2b,
	0x07, 0x56, 0xef, 0x05, 0xd8, 0x3e, 0x8f, 0xfd, 0x84, 0xa1, 0x24, 0xcf, 0xa0, 0xa6, 0x87, 0x91,
	0xae, 0x65, 0x5e, 0xd6, 0x5a, 0xbf, 0x4c, 0xeb, 0x41, 0x21, 0xf6, 0x2e, 0xcc, 0xa9, 0x68, 0x86,
	0x1c, 0x80, 0xb3, 0xba, 0x24, 0xd3, 0xee, 0x91, 0xa1, 0x57, 0xc1, 0x84, 0xc0, 0x46, 0x44, 0x15,
	0x2d, 0xe7, 0x31, 0xdf, 0xbd, 0xaf, 0x00, 0x3e, 0x8f, 0x47, 0x2c, 0x32, 0x67, 0xb8, 0xa7, 0x4f,
	0x20, 0x51, 0xb3, 0x90, 0x47, 0x68, 0x6a, 0xd7, 0x02, 0x5b, 0x13, 0x27, 0x3c, 0x42, 0x72, 0x04,
	0x9b, 0x21, 0xcf, 0xf2, 0x14, 0xcb, 0x85, 0x55, 0x1e, 0xed, 0xdd, 0x5c, 0xc5, 0x0f, 0x55, 0xef,
	0x14, 0x76, 0x8e, 0xd7, 0xff, 0x4a, 0x80, 0x32, 0xe7, 0x4c, 0x22, 0xf9, 0x07, 0xea, 0x29, 0x8f,
	0xf5, 0x35, 0x16, 0xab, 0xab, 0xa5, 0x3c, 0x3e, 0x8b, 0xf4, 0x71, 0xcb, 0x45, 0x18, 0xa2, 0x94,
	0xa6, 0x8f, 0x1d, 0x2c, 0xe1, 0xf3, 0x39, 0x38, 0xab, 0x1b, 0x21, 0xff, 0x02, 0xf1, 0x27, 0xef,
	0x66, 0xa7, 0x93, 0xe0, 0xe3, 0x70, 0x3a, 0xfb, 0x34, 0xfe, 0x30, 0x9e, 0x5c, 0x8c, 0xdb, 0x7f,
	0x91, 0x1d, 0xd8, 0xba, 0xc5, 0x4f, 0x47, 0x9f, 0xa7, 0x6d, 0xeb, 0x1e, 0xf9, 0xfe, 0x7c, 0x32,
	0x6e, 0x57, 0xee, 0x91, 0xc7, 0x9a, 0xac, 0x0e, 0x7e, 0x54, 0xa0, 0x79, 0x6b, 0x58, 0xf2, 0x0a,
	0x9c, 0x13, 0xe3, 0xbd, 0xcf, 0x63, 0xd2, 0xba, 0x7b, 0x7c, 0x9d, 0x4e, 0x89, 0x7f, 0xf7, 0xba,
	0x23, 0x68, 0x0d, 0xf3, 0x1c, 0x59, 0xb4, 0xf2, 0x7b, 0xeb, 0xae, 0xc1, 0xf2, 0xc1, 0xf4, 0x11,
	0xec, 0x9e, 0x2b, 0x81, 0x34, 0xfb, 0x83, 0x22, 0x9e, 0x45, 0x0e, 0xc1, 0x3e, 0x49, 0xb9, 0x34,
	0xd3, 0x6f, 0xaf, 0x53, 0x4b, 0xd7, 0x1f, 0x4a, 0x3e, 0x86, 0x2f, 0x76, 0xc2, 0x14, 0x0a, 0x46,
	0xd3, 0x79, 0xdd, 0x58, 0xfc, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xc1, 0xbe, 0xcd,
	0xf4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildloggerClient is the client API for Buildlogger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildloggerClient interface {
	CreateLog(ctx context.Context, in *LogInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error)
	AppendLogLines(ctx context.Context, in *LogLines, opts ...grpc.CallOption) (*BuildloggerResponse, error)
	StreamAppendLogLines(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamAppendLogLinesClient, error)
	CloseLog(ctx context.Context, in *LogEndInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error)
}

type buildloggerClient struct {
	cc *grpc.ClientConn
}

func NewBuildloggerClient(cc *grpc.ClientConn) BuildloggerClient {
	return &buildloggerClient{cc}
}

func (c *buildloggerClient) CreateLog(ctx context.Context, in *LogInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/CreateLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildloggerClient) AppendLogLines(ctx context.Context, in *LogLines, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/AppendLogLines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildloggerClient) StreamAppendLogLines(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamAppendLogLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Buildlogger_serviceDesc.Streams[0], "/cedar.Buildlogger/StreamAppendLogLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildloggerStreamAppendLogLinesClient{stream}
	return x, nil
}

type Buildlogger_StreamAppendLogLinesClient interface {
	Send(*LogLines) error
	CloseAndRecv() (*BuildloggerResponse, error)
	grpc.ClientStream
}

type buildloggerStreamAppendLogLinesClient struct {
	grpc.ClientStream
}

func (x *buildloggerStreamAppendLogLinesClient) Send(m *LogLines) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildloggerStreamAppendLogLinesClient) CloseAndRecv() (*BuildloggerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildloggerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildloggerClient) CloseLog(ctx context.Context, in *LogEndInfo, opts ...grpc.CallOption) (*BuildloggerResponse, error) {
	out := new(BuildloggerResponse)
	err := c.cc.Invoke(ctx, "/cedar.Buildlogger/CloseLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildloggerServer is the server API for Buildlogger service.
type BuildloggerServer interface {
	CreateLog(context.Context, *LogInfo) (*BuildloggerResponse, error)
	AppendLogLines(context.Context, *LogLines) (*BuildloggerResponse, error)
	StreamAppendLogLines(Buildlogger_StreamAppendLogLinesServer) error
	CloseLog(context.Context, *LogEndInfo) (*BuildloggerResponse, error)
}

// UnimplementedBuildloggerServer can be embedded to have forward compatible implementations.
type UnimplementedBuildloggerServer struct {
}

func (*UnimplementedBuildloggerServer) CreateLog(ctx context.Context, req *LogInfo) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLog not implemented")
}
func (*UnimplementedBuildloggerServer) AppendLogLines(ctx context.Context, req *LogLines) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendLogLines not implemented")
}
func (*UnimplementedBuildloggerServer) StreamAppendLogLines(srv Buildlogger_StreamAppendLogLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAppendLogLines not implemented")
}
func (*UnimplementedBuildloggerServer) CloseLog(ctx context.Context, req *LogEndInfo) (*BuildloggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLog not implemented")
}

func RegisterBuildloggerServer(s *grpc.Server, srv BuildloggerServer) {
	s.RegisterService(&_Buildlogger_serviceDesc, srv)
}

func _Buildlogger_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/CreateLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).CreateLog(ctx, req.(*LogInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildlogger_AppendLogLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogLines)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).AppendLogLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/AppendLogLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).AppendLogLines(ctx, req.(*LogLines))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildlogger_StreamAppendLogLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildloggerServer).StreamAppendLogLines(&buildloggerStreamAppendLogLinesServer{stream})
}

type Buildlogger_StreamAppendLogLinesServer interface {
	SendAndClose(*BuildloggerResponse) error
	Recv() (*LogLines, error)
	grpc.ServerStream
}

type buildloggerStreamAppendLogLinesServer struct {
	grpc.ServerStream
}

func (x *buildloggerStreamAppendLogLinesServer) SendAndClose(m *BuildloggerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildloggerStreamAppendLogLinesServer) Recv() (*LogLines, error) {
	m := new(LogLines)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Buildlogger_CloseLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogEndInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildloggerServer).CloseLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.Buildlogger/CloseLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildloggerServer).CloseLog(ctx, req.(*LogEndInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildlogger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.Buildlogger",
	HandlerType: (*BuildloggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLog",
			Handler:    _Buildlogger_CreateLog_Handler,
		},
		{
			MethodName: "AppendLogLines",
			Handler:    _Buildlogger_AppendLogLines_Handler,
		},
		{
			MethodName: "CloseLog",
			Handler:    _Buildlogger_CloseLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAppendLogLines",
			Handler:       _Buildlogger_StreamAppendLogLines_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "buildlogger.proto",
}
