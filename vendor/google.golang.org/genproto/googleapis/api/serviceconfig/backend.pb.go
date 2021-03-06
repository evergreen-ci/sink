// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/backend.proto

package serviceconfig

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Path Translation specifies how to combine the backend address with the
// request path in order to produce the appropriate forwarding URL for the
// request.
//
// Path Translation is applicable only to HTTP-based backends. Backends which
// do not accept requests over HTTP/HTTPS should leave `path_translation`
// unspecified.
type BackendRule_PathTranslation int32

const (
	BackendRule_PATH_TRANSLATION_UNSPECIFIED BackendRule_PathTranslation = 0
	// Use the backend address as-is, with no modification to the path. If the
	// URL pattern contains variables, the variable names and values will be
	// appended to the query string. If a query string parameter and a URL
	// pattern variable have the same name, this may result in duplicate keys in
	// the query string.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//     Method path:        /api/company/{cid}/user/{uid}
	//     Backend address:    https://example.cloudfunctions.net/getUser
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//     Request path: /api/company/widgetworks/user/johndoe
	//     Translated:
	//     https://example.cloudfunctions.net/getUser?cid=widgetworks&uid=johndoe
	//
	//     Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//     Translated:
	//     https://example.cloudfunctions.net/getUser?timezone=EST&cid=widgetworks&uid=johndoe
	BackendRule_CONSTANT_ADDRESS BackendRule_PathTranslation = 1
	// The request path will be appended to the backend address.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//     Method path:        /api/company/{cid}/user/{uid}
	//     Backend address:    https://example.appspot.com
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//     Request path: /api/company/widgetworks/user/johndoe
	//     Translated:
	//     https://example.appspot.com/api/company/widgetworks/user/johndoe
	//
	//     Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//     Translated:
	//     https://example.appspot.com/api/company/widgetworks/user/johndoe?timezone=EST
	BackendRule_APPEND_PATH_TO_ADDRESS BackendRule_PathTranslation = 2
)

var BackendRule_PathTranslation_name = map[int32]string{
	0: "PATH_TRANSLATION_UNSPECIFIED",
	1: "CONSTANT_ADDRESS",
	2: "APPEND_PATH_TO_ADDRESS",
}

var BackendRule_PathTranslation_value = map[string]int32{
	"PATH_TRANSLATION_UNSPECIFIED": 0,
	"CONSTANT_ADDRESS":             1,
	"APPEND_PATH_TO_ADDRESS":       2,
}

func (x BackendRule_PathTranslation) String() string {
	return proto.EnumName(BackendRule_PathTranslation_name, int32(x))
}

func (BackendRule_PathTranslation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87d0f28daa3f64f0, []int{1, 0}
}

// `Backend` defines the backend configuration for a service.
type Backend struct {
	// A list of API backend rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules                []*BackendRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Backend) Reset()         { *m = Backend{} }
func (m *Backend) String() string { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()    {}
func (*Backend) Descriptor() ([]byte, []int) {
	return fileDescriptor_87d0f28daa3f64f0, []int{0}
}

func (m *Backend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backend.Unmarshal(m, b)
}
func (m *Backend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backend.Marshal(b, m, deterministic)
}
func (m *Backend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backend.Merge(m, src)
}
func (m *Backend) XXX_Size() int {
	return xxx_messageInfo_Backend.Size(m)
}
func (m *Backend) XXX_DiscardUnknown() {
	xxx_messageInfo_Backend.DiscardUnknown(m)
}

var xxx_messageInfo_Backend proto.InternalMessageInfo

func (m *Backend) GetRules() []*BackendRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The address of the API backend.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request. The default
	// varies based on the request protocol and deployment environment.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Minimum deadline in seconds needed for this method. Calls having deadline
	// value lower than this will be rejected.
	MinDeadline float64 `protobuf:"fixed64,4,opt,name=min_deadline,json=minDeadline,proto3" json:"min_deadline,omitempty"`
	// The number of seconds to wait for the completion of a long running
	// operation. The default is no deadline.
	OperationDeadline float64                     `protobuf:"fixed64,5,opt,name=operation_deadline,json=operationDeadline,proto3" json:"operation_deadline,omitempty"`
	PathTranslation   BackendRule_PathTranslation `protobuf:"varint,6,opt,name=path_translation,json=pathTranslation,proto3,enum=google.api.BackendRule_PathTranslation" json:"path_translation,omitempty"`
	// Authentication settings used by the backend.
	//
	// These are typically used to provide service management functionality to
	// a backend served on a publicly-routable URL. The `authentication`
	// details should match the authentication behavior used by the backend.
	//
	// For example, specifying `jwt_audience` implies that the backend expects
	// authentication via a JWT.
	//
	// When authentication is unspecified, a JWT ID token will be generated with
	// the value from [BackendRule.address][google.api.BackendRule.address] as jwt_audience, overrode to the
	// HTTP "Authorization" request header and sent to the backend.
	//
	// Refer to https://developers.google.com/identity/protocols/OpenIDConnect for
	// JWT ID token.
	//
	// Types that are valid to be assigned to Authentication:
	//	*BackendRule_JwtAudience
	//	*BackendRule_DisableAuth
	Authentication       isBackendRule_Authentication `protobuf_oneof:"authentication"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BackendRule) Reset()         { *m = BackendRule{} }
func (m *BackendRule) String() string { return proto.CompactTextString(m) }
func (*BackendRule) ProtoMessage()    {}
func (*BackendRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_87d0f28daa3f64f0, []int{1}
}

func (m *BackendRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendRule.Unmarshal(m, b)
}
func (m *BackendRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendRule.Marshal(b, m, deterministic)
}
func (m *BackendRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendRule.Merge(m, src)
}
func (m *BackendRule) XXX_Size() int {
	return xxx_messageInfo_BackendRule.Size(m)
}
func (m *BackendRule) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendRule.DiscardUnknown(m)
}

var xxx_messageInfo_BackendRule proto.InternalMessageInfo

func (m *BackendRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *BackendRule) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BackendRule) GetDeadline() float64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func (m *BackendRule) GetMinDeadline() float64 {
	if m != nil {
		return m.MinDeadline
	}
	return 0
}

func (m *BackendRule) GetOperationDeadline() float64 {
	if m != nil {
		return m.OperationDeadline
	}
	return 0
}

func (m *BackendRule) GetPathTranslation() BackendRule_PathTranslation {
	if m != nil {
		return m.PathTranslation
	}
	return BackendRule_PATH_TRANSLATION_UNSPECIFIED
}

type isBackendRule_Authentication interface {
	isBackendRule_Authentication()
}

type BackendRule_JwtAudience struct {
	JwtAudience string `protobuf:"bytes,7,opt,name=jwt_audience,json=jwtAudience,proto3,oneof"`
}

type BackendRule_DisableAuth struct {
	DisableAuth bool `protobuf:"varint,8,opt,name=disable_auth,json=disableAuth,proto3,oneof"`
}

func (*BackendRule_JwtAudience) isBackendRule_Authentication() {}

func (*BackendRule_DisableAuth) isBackendRule_Authentication() {}

func (m *BackendRule) GetAuthentication() isBackendRule_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *BackendRule) GetJwtAudience() string {
	if x, ok := m.GetAuthentication().(*BackendRule_JwtAudience); ok {
		return x.JwtAudience
	}
	return ""
}

func (m *BackendRule) GetDisableAuth() bool {
	if x, ok := m.GetAuthentication().(*BackendRule_DisableAuth); ok {
		return x.DisableAuth
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BackendRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BackendRule_JwtAudience)(nil),
		(*BackendRule_DisableAuth)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.api.BackendRule_PathTranslation", BackendRule_PathTranslation_name, BackendRule_PathTranslation_value)
	proto.RegisterType((*Backend)(nil), "google.api.Backend")
	proto.RegisterType((*BackendRule)(nil), "google.api.BackendRule")
}

func init() { proto.RegisterFile("google/api/backend.proto", fileDescriptor_87d0f28daa3f64f0) }

var fileDescriptor_87d0f28daa3f64f0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xeb, 0xa4, 0x6d, 0xc2, 0x24, 0x4a, 0xcd, 0x08, 0xc1, 0xa8, 0x62, 0x61, 0xc2, 0x02,
	0x6f, 0xea, 0x48, 0x65, 0x83, 0xc4, 0x6a, 0x52, 0x1b, 0x1a, 0x09, 0x39, 0xd6, 0xd8, 0x6c, 0xd8,
	0x58, 0x13, 0xfb, 0xe0, 0x4c, 0x71, 0x66, 0x2c, 0x7b, 0x4c, 0xdf, 0x87, 0xa7, 0xe1, 0xb1, 0x90,
	0x2f, 0x75, 0x53, 0xa4, 0x2e, 0xff, 0xff, 0xfb, 0x8e, 0xe6, 0xa2, 0x83, 0x48, 0xa6, 0x54, 0x96,
	0xc3, 0x8a, 0x17, 0x62, 0xb5, 0xe3, 0xc9, 0x2f, 0x90, 0xa9, 0x53, 0x94, 0x4a, 0x2b, 0x8c, 0x3a,
	0xe2, 0xf0, 0x42, 0x2c, 0x3f, 0xa1, 0xc9, 0xba, 0x83, 0xf8, 0x0a, 0x9d, 0x95, 0x75, 0x0e, 0x15,
	0x31, 0xac, 0xb1, 0x3d, 0xbb, 0x7e, 0xe3, 0x3c, 0x6a, 0x4e, 0xef, 0xb0, 0x3a, 0x07, 0xd6, 0x59,
	0xcb, 0xbf, 0x63, 0x34, 0x3b, 0xaa, 0xf1, 0x25, 0x9a, 0x56, 0x90, 0x43, 0xa2, 0x55, 0x49, 0x0c,
	0xcb, 0xb0, 0x5f, 0xb0, 0x21, 0x63, 0x82, 0x26, 0x3c, 0x4d, 0x4b, 0xa8, 0x2a, 0x32, 0x6a, 0xd1,
	0x43, 0x6c, 0xa6, 0x52, 0xe0, 0x69, 0x2e, 0x24, 0x90, 0xb1, 0x65, 0xd8, 0x06, 0x1b, 0x32, 0x7e,
	0x87, 0xe6, 0x07, 0x21, 0xe3, 0x81, 0x9f, 0xb6, 0x7c, 0x76, 0x10, 0xd2, 0x7d, 0x50, 0xae, 0x10,
	0x56, 0x05, 0x94, 0x5c, 0x0b, 0x75, 0x24, 0x9e, 0xb5, 0xe2, 0xcb, 0x81, 0x0c, 0x3a, 0x43, 0x66,
	0xc1, 0xf5, 0x3e, 0xd6, 0x25, 0x97, 0x55, 0xde, 0x32, 0x72, 0x6e, 0x19, 0xf6, 0xe2, 0xfa, 0xc3,
	0x33, 0xaf, 0x75, 0x02, 0xae, 0xf7, 0xd1, 0xa3, 0xce, 0x2e, 0x8a, 0xa7, 0x05, 0x7e, 0x8f, 0xe6,
	0x77, 0xf7, 0x3a, 0xe6, 0x75, 0x2a, 0x40, 0x26, 0x40, 0x26, 0xcd, 0x03, 0x6f, 0x4f, 0xd8, 0xec,
	0xee, 0x5e, 0xd3, 0xbe, 0x6c, 0xa4, 0x54, 0x54, 0x7c, 0x97, 0x43, 0xcc, 0x6b, 0xbd, 0x27, 0x53,
	0xcb, 0xb0, 0xa7, 0x8d, 0xd4, 0xb7, 0xb4, 0xd6, 0xfb, 0x25, 0xa0, 0x8b, 0xff, 0x4e, 0xc3, 0x16,
	0x7a, 0x1b, 0xd0, 0xe8, 0x36, 0x8e, 0x18, 0xf5, 0xc3, 0x6f, 0x34, 0xda, 0x6c, 0xfd, 0xf8, 0xbb,
	0x1f, 0x06, 0xde, 0xcd, 0xe6, 0xcb, 0xc6, 0x73, 0xcd, 0x13, 0xfc, 0x0a, 0x99, 0x37, 0x5b, 0x3f,
	0x8c, 0xa8, 0x1f, 0xc5, 0xd4, 0x75, 0x99, 0x17, 0x86, 0xa6, 0x81, 0x2f, 0xd1, 0x6b, 0x1a, 0x04,
	0x9e, 0xef, 0xc6, 0xdd, 0xf8, 0x76, 0x60, 0xa3, 0xb5, 0x89, 0x16, 0xcd, 0x1d, 0x40, 0x6a, 0x91,
	0xb4, 0xa7, 0xac, 0x25, 0x5a, 0x24, 0xea, 0x70, 0xf4, 0x03, 0xeb, 0x79, 0xff, 0x05, 0x41, 0xb3,
	0x30, 0x81, 0xf1, 0xc3, 0xeb, 0x59, 0xa6, 0x72, 0x2e, 0x33, 0x47, 0x95, 0xd9, 0x2a, 0x03, 0xd9,
	0xae, 0xd3, 0xaa, 0x43, 0xbc, 0x10, 0x55, 0xbb, 0x6b, 0x15, 0x94, 0xbf, 0x45, 0x02, 0x89, 0x92,
	0x3f, 0x45, 0xf6, 0xf9, 0x49, 0xfa, 0x33, 0x3a, 0xfd, 0x4a, 0x83, 0xcd, 0xee, 0xbc, 0x1d, 0xfc,
	0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xac, 0xef, 0x0a, 0xed, 0xa3, 0x02, 0x00, 0x00,
}
