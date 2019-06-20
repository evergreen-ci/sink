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
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x4f, 0xdb, 0x4c,
	0x10, 0xfd, 0x9c, 0x90, 0xc4, 0x9e, 0xf0, 0x85, 0xb0, 0xf4, 0x87, 0x15, 0x5a, 0x35, 0x8a, 0x7a,
	0xb0, 0x7a, 0x08, 0x55, 0x7a, 0x28, 0xd0, 0x72, 0x08, 0x08, 0x2a, 0x5a, 0x37, 0x91, 0x4c, 0x2a,
	0xaa, 0x5e, 0xa2, 0x8d, 0x3d, 0xb8, 0x2e, 0xf6, 0xae, 0xb5, 0xbb, 0x41, 0xe5, 0x2f, 0xef, 0xa5,
	0x87, 0x6a, 0xd7, 0x4e, 0x02, 0xa8, 0x02, 0xf5, 0xe6, 0x79, 0x6f, 0xde, 0xcc, 0xec, 0xbc, 0x31,
	0x6c, 0xce, 0xe6, 0x49, 0x1a, 0xa5, 0x3c, 0x8e, 0x51, 0xf4, 0x73, 0xc1, 0x15, 0x27, 0xb5, 0x10,
	0x23, 0x2a, 0x3a, 0x2f, 0x62, 0xce, 0xe3, 0x14, 0x77, 0x0c, 0x38, 0x9b, 0x5f, 0xec, 0xa8, 0x24,
	0x43, 0xa9, 0x68, 0x96, 0x17, 0x79, 0xbd, 0x5f, 0x55, 0x68, 0xf8, 0x3c, 0x3e, 0x65, 0x17, 0x9c,
	0xb8, 0xd0, 0xc8, 0x05, 0xff, 0x81, 0xa1, 0x72, 0xad, 0xae, 0xe5, 0x39, 0xc1, 0x22, 0xd4, 0xcc,
	0x15, 0x0a, 0x99, 0x70, 0xe6, 0x56, 0x0a, 0xa6, 0x0c, 0x0d, 0x43, 0x45, 0x42, 0x99, 0x72, 0xab,
	0x25, 0x53, 0x84, 0x64, 0x1b, 0x1c, 0x45, 0xe5, 0xe5, 0x94, 0xd1, 0x0c, 0xdd, 0x35, 0xc3, 0xd9,
	0x1a, 0x18, 0xd1, 0x0c, 0xc9, 0x53, 0x68, 0x18, 0x32, 0x89, 0xdc, 0x9a, 0xa1, 0xea, 0x3a, 0x3c,
	0x8d, 0xc8, 0x33, 0x70, 0xf0, 0x27, 0x86, 0x73, 0xa5, 0x7b, 0xd5, 0xbb, 0x96, 0x57, 0x0b, 0x56,
	0x80, 0xa9, 0x89, 0x52, 0x15, 0x35, 0x1b, 0x65, 0x4d, 0x94, 0xca, 0xd4, 0x7c, 0x04, 0x35, 0x25,
	0x12, 0x9a, 0xba, 0xb6, 0x91, 0x15, 0x81, 0x96, 0xe4, 0x82, 0x87, 0x85, 0xc4, 0x29, 0x24, 0x1a,
	0x30, 0x12, 0x0f, 0xea, 0x17, 0x5c, 0x64, 0x54, 0xb9, 0xd0, 0xb5, 0xbc, 0xd6, 0xa0, 0xdd, 0x37,
	0x6b, 0xeb, 0xfb, 0x3c, 0x3e, 0x31, 0x78, 0x50, 0xf2, 0xe4, 0x1d, 0x38, 0x54, 0xc4, 0xf3, 0x0c,
	0x99, 0x92, 0x6e, 0xb3, 0x5b, 0xf5, 0x9a, 0x83, 0xe7, 0xab, 0x64, 0xbd, 0xbe, 0xfe, 0x70, 0xc1,
	0x1f, 0x33, 0x25, 0xae, 0x83, 0x55, 0x3e, 0xe9, 0x80, 0x9d, 0xd1, 0x84, 0xa5, 0x09, 0x43, 0x77,
	0xbd, 0x6b, 0x79, 0x76, 0xb0, 0x8c, 0xc9, 0x1e, 0x40, 0x28, 0x90, 0x2a, 0x8c, 0xa6, 0x54, 0xb9,
	0xff, 0x77, 0x2d, 0xaf, 0x39, 0xe8, 0xf4, 0x0b, 0xdb, 0xfa, 0x0b, 0xdb, 0xfa, 0x93, 0x85, 0x6d,
	0x81, 0x53, 0x66, 0x0f, 0x55, 0xe7, 0x3d, 0xb4, 0x6e, 0xf7, 0x24, 0x6d, 0xa8, 0x5e, 0xe2, 0x75,
	0xe9, 0x9e, 0xfe, 0xd4, 0x4b, 0xb9, 0xa2, 0xe9, 0x1c, 0x4b, 0xdf, 0x8a, 0x60, 0xbf, 0xb2, 0x6b,
	0xf5, 0x5e, 0x83, 0xed, 0xf3, 0xd8, 0x4f, 0x18, 0x4a, 0xf2, 0x12, 0x6a, 0x7a, 0x18, 0xe9, 0x5a,
	0xe6, 0x65, 0xad, 0xd5, 0xcb, 0x34, 0x1f, 0x14, 0x64, 0xef, 0xdc, 0x9c, 0x8a, 0x46, 0xc8, 0x2e,
	0x38, 0xcb, 0x4b, 0x32, 0xed, 0x1e, 0x18, 0x7a, 0x99, 0x4c, 0x08, 0xac, 0x45, 0x54, 0xd1, 0x72,
	0x1e, 0xf3, 0xdd, 0xfb, 0x0e, 0xe0, 0xf3, 0xf8, 0x98, 0x45, 0xe6, 0x0c, 0xb7, 0xf5, 0x09, 0x24,
	0x6a, 0x1a, 0xf2, 0x08, 0x4d, 0xed, 0x5a, 0x60, 0x6b, 0xe0, 0x88, 0x47, 0x48, 0x0e, 0x60, 0x3d,
	0xe4, 0x59, 0x9e, 0x62, 0xb9, 0xb0, 0xca, 0x83, 0xbd, 0x9b, 0xcb, 0xfc, 0xa1, 0xea, 0x9d, 0xc0,
	0xd6, 0xe1, 0xea, 0x5f, 0x09, 0x50, 0xe6, 0x9c, 0x49, 0x24, 0x8f, 0xa1, 0x9e, 0xf2, 0x58, 0x5f,
	0x63, 0xb1, 0xba, 0x5a, 0xca, 0xe3, 0xd3, 0x48, 0x1f, 0xb7, 0x9c, 0x87, 0x21, 0x4a, 0x69, 0xfa,
	0xd8, 0xc1, 0x22, 0x7c, 0x35, 0x03, 0x67, 0x79, 0x23, 0xe4, 0x09, 0x10, 0x7f, 0xfc, 0x61, 0x7a,
	0x32, 0x0e, 0x3e, 0x0f, 0x27, 0xd3, 0x2f, 0xa3, 0x4f, 0xa3, 0xf1, 0xf9, 0xa8, 0xfd, 0x1f, 0xd9,
	0x82, 0x8d, 0x1b, 0xf8, 0xe4, 0xf8, 0xeb, 0xa4, 0x6d, 0xdd, 0x01, 0x3f, 0x9e, 0x8d, 0x47, 0xed,
	0xca, 0x1d, 0xf0, 0x50, 0x83, 0xd5, 0xc1, 0x6f, 0x0b, 0x9a, 0x37, 0x86, 0x25, 0x6f, 0xc1, 0x39,
	0x32, 0xde, 0xfb, 0x3c, 0x26, 0xad, 0xdb, 0xc7, 0xd7, 0xe9, 0x94, 0xf1, 0xdf, 0x5e, 0x77, 0x00,
	0xad, 0x61, 0x9e, 0x23, 0x8b, 0x96, 0x7e, 0x6f, 0xdc, 0x36, 0x58, 0xde, 0x2b, 0xdf, 0x07, 0xe7,
	0x4c, 0x09, 0xa4, 0x99, 0xee, 0xfb, 0x2f, 0x4a, 0xcf, 0x22, 0x7b, 0x60, 0x1f, 0xa5, 0x5c, 0x9a,
	0x91, 0x37, 0x57, 0xd2, 0xd2, 0xea, 0xfb, 0xc4, 0x87, 0xf0, 0xcd, 0x4e, 0x98, 0x42, 0xc1, 0x68,
	0x3a, 0xab, 0x1b, 0x5f, 0xdf, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x0b, 0xc7, 0x40, 0xe9,
	0x04, 0x00, 0x00,
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
	StreamLog(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamLogClient, error)
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

func (c *buildloggerClient) StreamLog(ctx context.Context, opts ...grpc.CallOption) (Buildlogger_StreamLogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Buildlogger_serviceDesc.Streams[0], "/cedar.Buildlogger/StreamLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildloggerStreamLogClient{stream}
	return x, nil
}

type Buildlogger_StreamLogClient interface {
	Send(*LogLines) error
	CloseAndRecv() (*BuildloggerResponse, error)
	grpc.ClientStream
}

type buildloggerStreamLogClient struct {
	grpc.ClientStream
}

func (x *buildloggerStreamLogClient) Send(m *LogLines) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildloggerStreamLogClient) CloseAndRecv() (*BuildloggerResponse, error) {
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
	StreamLog(Buildlogger_StreamLogServer) error
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
func (*UnimplementedBuildloggerServer) StreamLog(srv Buildlogger_StreamLogServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLog not implemented")
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

func _Buildlogger_StreamLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildloggerServer).StreamLog(&buildloggerStreamLogServer{stream})
}

type Buildlogger_StreamLogServer interface {
	SendAndClose(*BuildloggerResponse) error
	Recv() (*LogLines, error)
	grpc.ServerStream
}

type buildloggerStreamLogServer struct {
	grpc.ServerStream
}

func (x *buildloggerStreamLogServer) SendAndClose(m *BuildloggerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildloggerStreamLogServer) Recv() (*LogLines, error) {
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
			StreamName:    "StreamLog",
			Handler:       _Buildlogger_StreamLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "buildlogger.proto",
}
