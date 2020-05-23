// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/access_loggers/file/v4alpha/file.proto

package envoy_extensions_access_loggers_file_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FileAccessLog struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are valid to be assigned to AccessLogFormat:
	//	*FileAccessLog_HiddenEnvoyDeprecatedFormat
	//	*FileAccessLog_HiddenEnvoyDeprecatedJsonFormat
	//	*FileAccessLog_HiddenEnvoyDeprecatedTypedJsonFormat
	//	*FileAccessLog_LogFormat
	AccessLogFormat      isFileAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FileAccessLog) Reset()         { *m = FileAccessLog{} }
func (m *FileAccessLog) String() string { return proto.CompactTextString(m) }
func (*FileAccessLog) ProtoMessage()    {}
func (*FileAccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd2016a39dd49326, []int{0}
}

func (m *FileAccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileAccessLog.Unmarshal(m, b)
}
func (m *FileAccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileAccessLog.Marshal(b, m, deterministic)
}
func (m *FileAccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAccessLog.Merge(m, src)
}
func (m *FileAccessLog) XXX_Size() int {
	return xxx_messageInfo_FileAccessLog.Size(m)
}
func (m *FileAccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_FileAccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_FileAccessLog proto.InternalMessageInfo

func (m *FileAccessLog) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type isFileAccessLog_AccessLogFormat interface {
	isFileAccessLog_AccessLogFormat()
}

type FileAccessLog_HiddenEnvoyDeprecatedFormat struct {
	HiddenEnvoyDeprecatedFormat string `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_format,json=hiddenEnvoyDeprecatedFormat,proto3,oneof"`
}

type FileAccessLog_HiddenEnvoyDeprecatedJsonFormat struct {
	HiddenEnvoyDeprecatedJsonFormat *_struct.Struct `protobuf:"bytes,3,opt,name=hidden_envoy_deprecated_json_format,json=hiddenEnvoyDeprecatedJsonFormat,proto3,oneof"`
}

type FileAccessLog_HiddenEnvoyDeprecatedTypedJsonFormat struct {
	HiddenEnvoyDeprecatedTypedJsonFormat *_struct.Struct `protobuf:"bytes,4,opt,name=hidden_envoy_deprecated_typed_json_format,json=hiddenEnvoyDeprecatedTypedJsonFormat,proto3,oneof"`
}

type FileAccessLog_LogFormat struct {
	LogFormat *v4alpha.SubstitutionFormatString `protobuf:"bytes,5,opt,name=log_format,json=logFormat,proto3,oneof"`
}

func (*FileAccessLog_HiddenEnvoyDeprecatedFormat) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_HiddenEnvoyDeprecatedJsonFormat) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_HiddenEnvoyDeprecatedTypedJsonFormat) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_LogFormat) isFileAccessLog_AccessLogFormat() {}

func (m *FileAccessLog) GetAccessLogFormat() isFileAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

// Deprecated: Do not use.
func (m *FileAccessLog) GetHiddenEnvoyDeprecatedFormat() string {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_HiddenEnvoyDeprecatedFormat); ok {
		return x.HiddenEnvoyDeprecatedFormat
	}
	return ""
}

// Deprecated: Do not use.
func (m *FileAccessLog) GetHiddenEnvoyDeprecatedJsonFormat() *_struct.Struct {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_HiddenEnvoyDeprecatedJsonFormat); ok {
		return x.HiddenEnvoyDeprecatedJsonFormat
	}
	return nil
}

// Deprecated: Do not use.
func (m *FileAccessLog) GetHiddenEnvoyDeprecatedTypedJsonFormat() *_struct.Struct {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_HiddenEnvoyDeprecatedTypedJsonFormat); ok {
		return x.HiddenEnvoyDeprecatedTypedJsonFormat
	}
	return nil
}

func (m *FileAccessLog) GetLogFormat() *v4alpha.SubstitutionFormatString {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_LogFormat); ok {
		return x.LogFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileAccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileAccessLog_HiddenEnvoyDeprecatedFormat)(nil),
		(*FileAccessLog_HiddenEnvoyDeprecatedJsonFormat)(nil),
		(*FileAccessLog_HiddenEnvoyDeprecatedTypedJsonFormat)(nil),
		(*FileAccessLog_LogFormat)(nil),
	}
}

func init() {
	proto.RegisterType((*FileAccessLog)(nil), "envoy.extensions.access_loggers.file.v4alpha.FileAccessLog")
}

func init() {
	proto.RegisterFile("envoy/extensions/access_loggers/file/v4alpha/file.proto", fileDescriptor_fd2016a39dd49326)
}

var fileDescriptor_fd2016a39dd49326 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x49, 0xba, 0x2d, 0xac, 0x11, 0x52, 0x59, 0x0e, 0x5d, 0xb5, 0x50, 0x96, 0x3f, 0x87,
	0x22, 0x21, 0x5b, 0x62, 0x41, 0x95, 0x56, 0x5c, 0x88, 0xa0, 0x02, 0xc4, 0xa1, 0x74, 0x39, 0x13,
	0x79, 0x13, 0xaf, 0xd7, 0x60, 0x3c, 0x91, 0x3d, 0x89, 0xba, 0x37, 0x8e, 0x88, 0x0b, 0x77, 0x3e,
	0x0a, 0x77, 0x24, 0xae, 0x7c, 0x9d, 0x9e, 0x50, 0xec, 0xa4, 0x69, 0xa1, 0x95, 0xca, 0x2d, 0xce,
	0xcc, 0xbc, 0x9f, 0xe7, 0xe5, 0x85, 0xec, 0x0a, 0x53, 0xc1, 0x92, 0x89, 0x43, 0x14, 0xc6, 0x29,
	0x30, 0x8e, 0xf1, 0x2c, 0x13, 0xce, 0xa5, 0x1a, 0xa4, 0x14, 0xd6, 0xb1, 0xb9, 0xd2, 0x82, 0x55,
	0x8f, 0xb9, 0x2e, 0x16, 0xdc, 0x1f, 0x68, 0x61, 0x01, 0x61, 0xf0, 0xd0, 0x0f, 0xd2, 0x6e, 0x90,
	0x9e, 0x1e, 0xa4, 0xbe, 0xb7, 0x19, 0xdc, 0x9c, 0x04, 0x4c, 0x06, 0x66, 0xae, 0x24, 0xcb, 0xc0,
	0x76, 0x9a, 0xae, 0x9c, 0x39, 0x54, 0x58, 0xa2, 0x02, 0x93, 0xce, 0xc1, 0x7e, 0xe2, 0x98, 0x3a,
	0xb4, 0xca, 0xc8, 0x40, 0xda, 0xbc, 0x29, 0x01, 0xa4, 0x16, 0xcc, 0x9f, 0x66, 0xe5, 0x9c, 0x39,
	0xb4, 0x65, 0x86, 0x4d, 0xf5, 0x56, 0x99, 0x17, 0x9c, 0x71, 0x63, 0x00, 0x39, 0xfa, 0x05, 0x1c,
	0x72, 0x2c, 0x5d, 0x53, 0xbe, 0xf3, 0x4f, 0xb9, 0x12, 0xb6, 0xbe, 0x6f, 0xa7, 0xbf, 0x51, 0x71,
	0xad, 0x72, 0x8e, 0x82, 0xb5, 0x0f, 0xa1, 0x70, 0xf7, 0x5b, 0x8f, 0x5c, 0xdb, 0x53, 0x5a, 0x3c,
	0xf3, 0x8b, 0xbd, 0x01, 0x39, 0xd8, 0x22, 0xbd, 0x82, 0xe3, 0x62, 0x18, 0x8d, 0xa2, 0x9d, 0x7e,
	0x72, 0xf9, 0x28, 0xe9, 0xd9, 0x78, 0x14, 0x1d, 0xf8, 0x97, 0x83, 0x57, 0x64, 0x7b, 0xa1, 0xf2,
	0x5c, 0x98, 0xd4, 0x2f, 0x9b, 0xe6, 0xa2, 0xb0, 0x22, 0xe3, 0x28, 0xf2, 0x66, 0xad, 0x61, 0xec,
	0xc7, 0xe2, 0x61, 0xf4, 0xf2, 0xd2, 0xc1, 0x56, 0xe8, 0x7d, 0x51, 0xb7, 0x3e, 0x3f, 0xee, 0xdc,
	0xf3, 0x8d, 0x83, 0x8f, 0xe4, 0xde, 0x79, 0x52, 0x1f, 0xdc, 0xb1, 0x4d, 0xc3, 0x95, 0x51, 0xb4,
	0x73, 0xf5, 0xd1, 0x06, 0x0d, 0x06, 0xd1, 0xd6, 0x20, 0x3a, 0xf5, 0x06, 0x35, 0xa0, 0xdb, 0x67,
	0x82, 0x5e, 0x3b, 0x30, 0x0d, 0xac, 0x22, 0x0f, 0xce, 0x83, 0xe1, 0xb2, 0xf8, 0x0b, 0xd9, 0xbb,
	0x08, 0xf2, 0xfe, 0x99, 0xc8, 0x77, 0xb5, 0xd8, 0x09, 0xee, 0x7b, 0x42, 0x34, 0xc8, 0x56, 0x78,
	0xd5, 0x0b, 0x8f, 0x69, 0x88, 0x55, 0x08, 0x0a, 0xad, 0x83, 0xd2, 0x66, 0x88, 0x4e, 0x4f, 0x04,
	0x25, 0x48, 0x4c, 0x7d, 0x4c, 0x92, 0x2b, 0x47, 0xc9, 0xea, 0xd7, 0x28, 0x5e, 0xaf, 0xd1, 0x7d,
	0x0d, 0x32, 0x14, 0x27, 0x4f, 0xbf, 0xff, 0xfc, 0xb2, 0xbd, 0x4b, 0x9e, 0x5c, 0x2c, 0xa8, 0x63,
	0x7a, 0xea, 0x53, 0x27, 0x37, 0xc8, 0xf5, 0xae, 0xaf, 0xb9, 0x64, 0xf2, 0xf6, 0xc7, 0xe7, 0x5f,
	0xbf, 0xd7, 0xe2, 0xf5, 0x15, 0x32, 0x51, 0x10, 0xae, 0x5a, 0x58, 0x38, 0x5c, 0xd2, 0xff, 0xf9,
	0x19, 0x92, 0x7e, 0x4d, 0xda, 0xaf, 0x9d, 0xdb, 0x8f, 0x66, 0x6b, 0xde, 0xc2, 0xf1, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x43, 0xe0, 0x60, 0x60, 0x89, 0x03, 0x00, 0x00,
}