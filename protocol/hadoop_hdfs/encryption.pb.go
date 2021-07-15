// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encryption.proto

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateEncryptionZoneRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	KeyName          *string `protobuf:"bytes,2,opt,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateEncryptionZoneRequestProto) Reset()         { *m = CreateEncryptionZoneRequestProto{} }
func (m *CreateEncryptionZoneRequestProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneRequestProto) ProtoMessage()    {}
func (*CreateEncryptionZoneRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0}
}

func (m *CreateEncryptionZoneRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *CreateEncryptionZoneRequestProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type CreateEncryptionZoneResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CreateEncryptionZoneResponseProto) Reset()         { *m = CreateEncryptionZoneResponseProto{} }
func (m *CreateEncryptionZoneResponseProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneResponseProto) ProtoMessage()    {}
func (*CreateEncryptionZoneResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{1}
}

type ListEncryptionZonesRequestProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListEncryptionZonesRequestProto) Reset()         { *m = ListEncryptionZonesRequestProto{} }
func (m *ListEncryptionZonesRequestProto) String() string { return proto.CompactTextString(m) }
func (*ListEncryptionZonesRequestProto) ProtoMessage()    {}
func (*ListEncryptionZonesRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{2}
}

func (m *ListEncryptionZonesRequestProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type EncryptionZoneProto struct {
	Id                    *int64                      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Path                  *string                     `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	Suite                 *CipherSuiteProto           `protobuf:"varint,3,req,name=suite,enum=hadoop.hdfs.CipherSuiteProto" json:"suite,omitempty"`
	CryptoProtocolVersion *CryptoProtocolVersionProto `protobuf:"varint,4,req,name=cryptoProtocolVersion,enum=hadoop.hdfs.CryptoProtocolVersionProto" json:"cryptoProtocolVersion,omitempty"`
	KeyName               *string                     `protobuf:"bytes,5,req,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized      []byte                      `json:"-"`
}

func (m *EncryptionZoneProto) Reset()                    { *m = EncryptionZoneProto{} }
func (m *EncryptionZoneProto) String() string            { return proto.CompactTextString(m) }
func (*EncryptionZoneProto) ProtoMessage()               {}
func (*EncryptionZoneProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *EncryptionZoneProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *EncryptionZoneProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *EncryptionZoneProto) GetSuite() CipherSuiteProto {
	if m != nil && m.Suite != nil {
		return *m.Suite
	}
	return CipherSuiteProto_UNKNOWN
}

func (m *EncryptionZoneProto) GetCryptoProtocolVersion() CryptoProtocolVersionProto {
	if m != nil && m.CryptoProtocolVersion != nil {
		return *m.CryptoProtocolVersion
	}
	return CryptoProtocolVersionProto_UNKNOWN_PROTOCOL_VERSION
}

func (m *EncryptionZoneProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type ListEncryptionZonesResponseProto struct {
	Zones            []*EncryptionZoneProto `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	HasMore          *bool                  `protobuf:"varint,2,req,name=hasMore" json:"hasMore,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListEncryptionZonesResponseProto) Reset()         { *m = ListEncryptionZonesResponseProto{} }
func (m *ListEncryptionZonesResponseProto) String() string { return proto.CompactTextString(m) }
func (*ListEncryptionZonesResponseProto) ProtoMessage()    {}
func (*ListEncryptionZonesResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{4}
}

func (m *ListEncryptionZonesResponseProto) GetZones() []*EncryptionZoneProto {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *ListEncryptionZonesResponseProto) GetHasMore() bool {
	if m != nil && m.HasMore != nil {
		return *m.HasMore
	}
	return false
}

type GetEZForPathRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetEZForPathRequestProto) Reset()                    { *m = GetEZForPathRequestProto{} }
func (m *GetEZForPathRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathRequestProto) ProtoMessage()               {}
func (*GetEZForPathRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GetEZForPathRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetEZForPathResponseProto struct {
	Zone             *EncryptionZoneProto `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *GetEZForPathResponseProto) Reset()                    { *m = GetEZForPathResponseProto{} }
func (m *GetEZForPathResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathResponseProto) ProtoMessage()               {}
func (*GetEZForPathResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *GetEZForPathResponseProto) GetZone() *EncryptionZoneProto {
	if m != nil {
		return m.Zone
	}
	return nil
}

var fileDescriptor2 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x99, 0xa4, 0xe5, 0xbd, 0xde, 0x42, 0x29, 0x79, 0x14, 0xe6, 0x3d, 0x78, 0x38, 0x46,
	0xc4, 0x2c, 0x24, 0x60, 0x15, 0xf7, 0xb6, 0x54, 0x37, 0x5a, 0x6a, 0x04, 0x17, 0x05, 0x17, 0x43,
	0x72, 0x35, 0x41, 0xcd, 0x8c, 0x33, 0xd3, 0x45, 0xfd, 0x35, 0xfe, 0x3b, 0xff, 0x86, 0x64, 0x62,
	0x4b, 0x46, 0x82, 0xb8, 0x09, 0x77, 0x66, 0xce, 0x3d, 0xe7, 0xf0, 0x05, 0x86, 0x58, 0xa6, 0x6a,
	0x2d, 0x4d, 0x21, 0xca, 0x58, 0x2a, 0x61, 0x44, 0xd0, 0xcf, 0x79, 0x26, 0x84, 0x8c, 0xf3, 0xec,
	0x5e, 0xff, 0x83, 0xea, 0x5b, 0x3f, 0x84, 0x73, 0x60, 0x53, 0x85, 0xdc, 0xe0, 0x6c, 0xbb, 0xb2,
	0x14, 0x25, 0x26, 0xf8, 0xb2, 0x42, 0x6d, 0x16, 0x76, 0x79, 0x08, 0xbe, 0x56, 0x29, 0x25, 0xcc,
	0x8b, 0x7a, 0x49, 0x35, 0x06, 0x14, 0x7e, 0x3d, 0xe2, 0x7a, 0xce, 0x9f, 0x91, 0x7a, 0x8c, 0x44,
	0xbd, 0x64, 0x73, 0x0c, 0xf7, 0x60, 0xb7, 0xdd, 0x4f, 0x4b, 0x51, 0x6a, 0xb4, 0x86, 0xe1, 0x11,
	0xec, 0x5c, 0x16, 0xda, 0xb8, 0x12, 0xed, 0x64, 0x0e, 0xc0, 0x2b, 0x32, 0x1b, 0xe9, 0x27, 0x5e,
	0x91, 0x85, 0xef, 0x04, 0xfe, 0xb8, 0xfa, 0x56, 0x5d, 0x10, 0x40, 0x47, 0x72, 0x93, 0x53, 0xcf,
	0x96, 0xb5, 0x73, 0x70, 0x0c, 0x5d, 0xbd, 0x2a, 0x0c, 0x52, 0x9f, 0x79, 0xd1, 0x60, 0xfc, 0x3f,
	0x6e, 0xc0, 0x88, 0xa7, 0x85, 0xcc, 0x51, 0xdd, 0x54, 0xef, 0xd6, 0x31, 0xa9, 0xb5, 0xc1, 0x1d,
	0x8c, 0x6c, 0x9a, 0xb0, 0xb7, 0xa9, 0x78, 0xba, 0x45, 0xa5, 0x0b, 0x51, 0xd2, 0x8e, 0x35, 0x39,
	0x70, 0x4d, 0xda, 0x94, 0xb5, 0x5d, 0xbb, 0x4b, 0x93, 0x60, 0xd7, 0x56, 0xdd, 0x12, 0x34, 0xc0,
	0x5a, 0xe1, 0x34, 0x00, 0x06, 0xa7, 0xd0, 0x7d, 0xad, 0x6e, 0x29, 0x61, 0x7e, 0xd4, 0x1f, 0x33,
	0xa7, 0x4c, 0x0b, 0xa6, 0xa4, 0x96, 0x57, 0xa9, 0x39, 0xd7, 0x57, 0x42, 0xa1, 0x05, 0xf4, 0x3b,
	0xd9, 0x1c, 0xc3, 0x43, 0xa0, 0x17, 0x68, 0x66, 0xcb, 0x73, 0xa1, 0x16, 0xdc, 0xe4, 0xdf, 0xff,
	0xff, 0xf0, 0x1a, 0xfe, 0xba, 0xea, 0x66, 0xb9, 0x13, 0xe8, 0x54, 0x69, 0x94, 0x30, 0xf2, 0xa3,
	0x6e, 0x56, 0x3d, 0x39, 0x83, 0x7d, 0xa1, 0x1e, 0x62, 0x2e, 0x79, 0x9a, 0xa3, 0xb3, 0x23, 0x3f,
	0xd1, 0xd5, 0xc3, 0x64, 0xf4, 0x85, 0x8c, 0x35, 0xd1, 0x6f, 0x84, 0x7c, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xb2, 0xc6, 0x30, 0xf7, 0x02, 0x00, 0x00,
}
