// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/dates.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A date range.
type DateRange struct {
	// The start date, in yyyy-mm-dd format.
	StartDate *wrappers.StringValue `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format.
	EndDate              *wrappers.StringValue `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_dates_68fc4d867dec9fcb, []int{0}
}
func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (dst *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(dst, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DateRange) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.ads.googleads.v0.common.DateRange")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/dates.proto", fileDescriptor_dates_68fc4d867dec9fcb)
}

var fileDescriptor_dates_68fc4d867dec9fcb = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xc9, 0x09, 0x6a, 0xe3, 0x76, 0x93, 0x88, 0x14, 0xe9, 0x24, 0x0e, 0x5f, 0x0e, 0x1d,
	0x1c, 0x3a, 0x5d, 0x2d, 0x74, 0x2d, 0x15, 0x6e, 0x90, 0x03, 0x49, 0x9b, 0xcf, 0x50, 0xb8, 0xcb,
	0x77, 0x24, 0x69, 0x9d, 0xfd, 0x2b, 0x8e, 0xfe, 0x0a, 0x67, 0x7f, 0x95, 0x24, 0xb9, 0xeb, 0xa6,
	0x74, 0xba, 0x17, 0xee, 0x79, 0x9f, 0x97, 0x7c, 0xfc, 0x4e, 0x13, 0xe9, 0x06, 0x85, 0x54, 0x4e,
	0xa4, 0x18, 0xd2, 0xbe, 0x10, 0x1b, 0x6a, 0x5b, 0x32, 0x42, 0x49, 0x8f, 0x0e, 0x3a, 0x4b, 0x9e,
	0xf2, 0x71, 0x02, 0x40, 0x2a, 0x07, 0x07, 0x16, 0xf6, 0x05, 0x24, 0xf6, 0xaa, 0xff, 0x2f, 0x22,
	0xbd, 0xde, 0xbd, 0x89, 0x77, 0x2b, 0xbb, 0x0e, 0x6d, 0xdf, 0x9f, 0x7c, 0x30, 0x3e, 0x9a, 0x4b,
	0x8f, 0x2b, 0x69, 0x34, 0xe6, 0x53, 0xce, 0x9d, 0x97, 0xd6, 0xbf, 0x86, 0x89, 0x4b, 0x76, 0xc3,
	0x6e, 0x2f, 0xee, 0xaf, 0x7b, 0x2f, 0x0c, 0x0a, 0x78, 0xf6, 0x76, 0x6b, 0x74, 0x25, 0x9b, 0x1d,
	0xae, 0x46, 0x91, 0x0f, 0x86, 0xfc, 0x91, 0x9f, 0xa3, 0x51, 0xa9, 0x9a, 0x1d, 0x51, 0x3d, 0x43,
	0xa3, 0x42, 0x71, 0xf6, 0xcd, 0xf8, 0x64, 0x43, 0x2d, 0xfc, 0xff, 0x94, 0x19, 0x0f, 0xb0, 0x5b,
	0x06, 0xd5, 0x92, 0xbd, 0xcc, 0x7b, 0x5a, 0x53, 0x23, 0x8d, 0x06, 0xb2, 0x5a, 0x68, 0x34, 0x71,
	0x68, 0x38, 0x5a, 0xb7, 0x75, 0x7f, 0xdd, 0x70, 0x9a, 0x3e, 0x9f, 0xd9, 0xc9, 0xa2, 0x2c, 0xbf,
	0xb2, 0xf1, 0x22, 0xc9, 0x4a, 0xe5, 0x20, 0xc5, 0x90, 0xaa, 0x02, 0x9e, 0x22, 0xf6, 0x33, 0x00,
	0x75, 0xa9, 0x5c, 0x7d, 0x00, 0xea, 0xaa, 0xa8, 0x13, 0xb0, 0x3e, 0x8d, 0xc3, 0x0f, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x29, 0x33, 0x41, 0x4f, 0xbb, 0x01, 0x00, 0x00,
}
