// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/proto/view/ios/segmentview.proto

package ios

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SegmentView struct {
	Value     int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Titles    []string `protobuf:"bytes,2,rep,name=titles" json:"titles,omitempty"`
	Momentary bool     `protobuf:"varint,3,opt,name=momentary" json:"momentary,omitempty"`
	Enabled   bool     `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *SegmentView) Reset()                    { *m = SegmentView{} }
func (m *SegmentView) String() string            { return proto.CompactTextString(m) }
func (*SegmentView) ProtoMessage()               {}
func (*SegmentView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SegmentView) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *SegmentView) GetTitles() []string {
	if m != nil {
		return m.Titles
	}
	return nil
}

func (m *SegmentView) GetMomentary() bool {
	if m != nil {
		return m.Momentary
	}
	return false
}

func (m *SegmentView) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SegmentViewEvent struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *SegmentViewEvent) Reset()                    { *m = SegmentViewEvent{} }
func (m *SegmentViewEvent) String() string            { return proto.CompactTextString(m) }
func (*SegmentViewEvent) ProtoMessage()               {}
func (*SegmentViewEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SegmentViewEvent) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*SegmentView)(nil), "matcha.view.ios.SegmentView")
	proto.RegisterType((*SegmentViewEvent)(nil), "matcha.view.ios.SegmentViewEvent")
}

func init() {
	proto.RegisterFile("gomatcha.io/matcha/proto/view/ios/segmentview.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x8a, 0xf2, 0x4b, 0xf2,
	0xf5, 0xcb, 0x32, 0x53, 0xcb, 0xf5, 0x33, 0xf3, 0x8b, 0xf5, 0x8b, 0x53, 0xd3, 0x73, 0x53, 0xf3,
	0x4a, 0x40, 0x7c, 0x3d, 0xb0, 0x94, 0x10, 0x3f, 0x54, 0x0b, 0x58, 0x28, 0x33, 0xbf, 0x58, 0xa9,
	0x98, 0x8b, 0x3b, 0x18, 0xa2, 0x2a, 0x2c, 0x33, 0xb5, 0x5c, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31,
	0xa7, 0x34, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x2b,
	0xc9, 0x2c, 0xc9, 0x49, 0x2d, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x64,
	0xb8, 0x38, 0x73, 0xf3, 0x41, 0x7a, 0x13, 0x8b, 0x2a, 0x25, 0x98, 0x15, 0x18, 0x35, 0x38, 0x82,
	0x10, 0x02, 0x42, 0x12, 0x5c, 0xec, 0xa9, 0x79, 0x89, 0x49, 0x39, 0xa9, 0x29, 0x12, 0x2c, 0x60,
	0x39, 0x18, 0x57, 0x49, 0x83, 0x4b, 0x00, 0xc9, 0x52, 0xd7, 0xb2, 0xd4, 0xbc, 0x12, 0xec, 0x36,
	0x3b, 0xb9, 0x72, 0x29, 0x66, 0xe6, 0xeb, 0xc1, 0x7d, 0x0a, 0xa5, 0xc0, 0x7e, 0x81, 0xfb, 0xc1,
	0x89, 0x37, 0x20, 0x09, 0xc9, 0xb8, 0x28, 0xe6, 0xcc, 0xfc, 0xe2, 0x45, 0x4c, 0xdc, 0xbe, 0x60,
	0xb5, 0x99, 0xfe, 0xc1, 0x01, 0x4e, 0x49, 0x6c, 0x60, 0x1d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x88, 0x9f, 0x28, 0x0e, 0x34, 0x01, 0x00, 0x00,
}