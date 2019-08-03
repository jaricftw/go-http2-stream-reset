// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/access_reason.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible access reasons.
type AccessReasonEnum_AccessReason int32

const (
	// Not specified.
	AccessReasonEnum_UNSPECIFIED AccessReasonEnum_AccessReason = 0
	// Used for return value only. Represents value unknown in this version.
	AccessReasonEnum_UNKNOWN AccessReasonEnum_AccessReason = 1
	// The resource is owned by the user.
	AccessReasonEnum_OWNED AccessReasonEnum_AccessReason = 2
	// The resource is shared to the user.
	AccessReasonEnum_SHARED AccessReasonEnum_AccessReason = 3
	// The resource is licensed to the user.
	AccessReasonEnum_LICENSED AccessReasonEnum_AccessReason = 4
	// The user subscribed to the resource.
	AccessReasonEnum_SUBSCRIBED AccessReasonEnum_AccessReason = 5
	// The resource is accessible to the user.
	AccessReasonEnum_AFFILIATED AccessReasonEnum_AccessReason = 6
)

var AccessReasonEnum_AccessReason_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OWNED",
	3: "SHARED",
	4: "LICENSED",
	5: "SUBSCRIBED",
	6: "AFFILIATED",
}
var AccessReasonEnum_AccessReason_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"OWNED":       2,
	"SHARED":      3,
	"LICENSED":    4,
	"SUBSCRIBED":  5,
	"AFFILIATED":  6,
}

func (x AccessReasonEnum_AccessReason) String() string {
	return proto.EnumName(AccessReasonEnum_AccessReason_name, int32(x))
}
func (AccessReasonEnum_AccessReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_access_reason_fdd0044bbadbed8a, []int{0, 0}
}

// Indicates the way the resource such as user list is related to a user.
type AccessReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessReasonEnum) Reset()         { *m = AccessReasonEnum{} }
func (m *AccessReasonEnum) String() string { return proto.CompactTextString(m) }
func (*AccessReasonEnum) ProtoMessage()    {}
func (*AccessReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_access_reason_fdd0044bbadbed8a, []int{0}
}
func (m *AccessReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessReasonEnum.Unmarshal(m, b)
}
func (m *AccessReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessReasonEnum.Marshal(b, m, deterministic)
}
func (dst *AccessReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessReasonEnum.Merge(dst, src)
}
func (m *AccessReasonEnum) XXX_Size() int {
	return xxx_messageInfo_AccessReasonEnum.Size(m)
}
func (m *AccessReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AccessReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AccessReasonEnum)(nil), "google.ads.googleads.v1.enums.AccessReasonEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AccessReasonEnum_AccessReason", AccessReasonEnum_AccessReason_name, AccessReasonEnum_AccessReason_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/access_reason.proto", fileDescriptor_access_reason_fdd0044bbadbed8a)
}

var fileDescriptor_access_reason_fdd0044bbadbed8a = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4e, 0xf2, 0x30,
	0x18, 0xfd, 0x37, 0x7e, 0x50, 0x0b, 0xd1, 0xda, 0x4b, 0x23, 0x17, 0xf0, 0x00, 0x5d, 0x16, 0xef,
	0xea, 0x55, 0xc7, 0x0a, 0x2e, 0x92, 0x41, 0x98, 0x40, 0x62, 0x96, 0x98, 0xca, 0x96, 0x85, 0x04,
	0x5a, 0xe4, 0x03, 0xde, 0xc0, 0x17, 0xf1, 0xd2, 0x47, 0xf1, 0x41, 0xbc, 0xf0, 0x29, 0xcc, 0x5a,
	0x21, 0xdc, 0xe8, 0x4d, 0x73, 0xfa, 0x7d, 0xe7, 0x9c, 0x7c, 0xe7, 0x20, 0xbf, 0xd0, 0xba, 0x58,
	0xe4, 0x9e, 0xcc, 0xc0, 0xb3, 0xb0, 0x44, 0x3b, 0xdf, 0xcb, 0xd5, 0x76, 0x09, 0x9e, 0x9c, 0xcd,
	0x72, 0x80, 0xa7, 0x75, 0x2e, 0x41, 0x2b, 0xba, 0x5a, 0xeb, 0x8d, 0x26, 0x4d, 0xcb, 0xa3, 0x32,
	0x03, 0x7a, 0x90, 0xd0, 0x9d, 0x4f, 0x8d, 0xe4, 0xea, 0x7a, 0xef, 0xb8, 0x9a, 0x7b, 0x52, 0x29,
	0xbd, 0x91, 0x9b, 0xb9, 0x56, 0x60, 0xc5, 0xed, 0x57, 0x07, 0x61, 0x6e, 0x4c, 0x47, 0xc6, 0x53,
	0xa8, 0xed, 0xb2, 0xfd, 0x82, 0x1a, 0xc7, 0x33, 0x72, 0x81, 0xea, 0xe3, 0x38, 0x19, 0x8a, 0x4e,
	0xd4, 0x8d, 0x44, 0x88, 0xff, 0x91, 0x3a, 0x3a, 0x19, 0xc7, 0xf7, 0xf1, 0x60, 0x1a, 0x63, 0x87,
	0x9c, 0xa1, 0xea, 0x60, 0x1a, 0x8b, 0x10, 0xbb, 0x04, 0xa1, 0x5a, 0x72, 0xc7, 0x47, 0x22, 0xc4,
	0x15, 0xd2, 0x40, 0xa7, 0xfd, 0xa8, 0x23, 0xe2, 0x44, 0x84, 0xf8, 0x3f, 0x39, 0x47, 0x28, 0x19,
	0x07, 0x49, 0x67, 0x14, 0x05, 0x22, 0xc4, 0xd5, 0xf2, 0xcf, 0xbb, 0xdd, 0xa8, 0x1f, 0xf1, 0x07,
	0x11, 0xe2, 0x5a, 0xf0, 0xe9, 0xa0, 0xd6, 0x4c, 0x2f, 0xe9, 0x9f, 0x59, 0x82, 0xcb, 0xe3, 0xb3,
	0x86, 0x65, 0x80, 0xa1, 0xf3, 0x18, 0xfc, 0x68, 0x0a, 0xbd, 0x90, 0xaa, 0xa0, 0x7a, 0x5d, 0x78,
	0x45, 0xae, 0x4c, 0xbc, 0x7d, 0x85, 0xab, 0x39, 0xfc, 0xd2, 0xe8, 0xad, 0x79, 0xdf, 0xdc, 0x4a,
	0x8f, 0xf3, 0x77, 0xb7, 0xd9, 0xb3, 0x56, 0x3c, 0x03, 0x6a, 0x61, 0x89, 0x26, 0x3e, 0x2d, 0x6b,
	0x81, 0x8f, 0xfd, 0x3e, 0xe5, 0x19, 0xa4, 0x87, 0x7d, 0x3a, 0xf1, 0x53, 0xb3, 0xff, 0x72, 0x5b,
	0x76, 0xc8, 0x18, 0xcf, 0x80, 0xb1, 0x03, 0x83, 0xb1, 0x89, 0xcf, 0x98, 0xe1, 0x3c, 0xd7, 0xcc,
	0x61, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0x92, 0xba, 0x1a, 0xe9, 0x01, 0x00, 0x00,
}
