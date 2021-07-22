// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IpcConnectionContext.proto

/*
Package hadoop_common is a generated protocol buffer package.

It is generated from these files:
	IpcConnectionContext.proto
	ZKFCProtocol.proto
	RefreshAuthorizationPolicyProtocol.proto
	GenericRefreshProtocol.proto
	RpcHeader.proto
	TraceAdmin.proto
	RefreshCallQueueProtocol.proto
	Security.proto
	ProtobufRpcEngine.proto
	RefreshUserMappingsProtocol.proto
	GetUserMappingsProtocol.proto
	HAServiceProtocol.proto
	ProtocolInfo.proto

It has these top-level messages:
	UserInformationProto
	IpcConnectionContextProto
	CedeActiveRequestProto
	CedeActiveResponseProto
	GracefulFailoverRequestProto
	GracefulFailoverResponseProto
	RefreshServiceAclRequestProto
	RefreshServiceAclResponseProto
	GenericRefreshRequestProto
	GenericRefreshResponseProto
	GenericRefreshResponseCollectionProto
	RPCTraceInfoProto
	RPCCallerContextProto
	RpcRequestHeaderProto
	RpcResponseHeaderProto
	RpcSaslProto
	ListSpanReceiversRequestProto
	SpanReceiverListInfo
	ListSpanReceiversResponseProto
	ConfigPair
	AddSpanReceiverRequestProto
	AddSpanReceiverResponseProto
	RemoveSpanReceiverRequestProto
	RemoveSpanReceiverResponseProto
	RefreshCallQueueRequestProto
	RefreshCallQueueResponseProto
	TokenProto
	GetDelegationTokenRequestProto
	GetDelegationTokenResponseProto
	RenewDelegationTokenRequestProto
	RenewDelegationTokenResponseProto
	CancelDelegationTokenRequestProto
	CancelDelegationTokenResponseProto
	RequestHeaderProto
	RefreshUserToGroupsMappingsRequestProto
	RefreshUserToGroupsMappingsResponseProto
	RefreshSuperUserGroupsConfigurationRequestProto
	RefreshSuperUserGroupsConfigurationResponseProto
	GetGroupsForUserRequestProto
	GetGroupsForUserResponseProto
	HAStateChangeRequestInfoProto
	MonitorHealthRequestProto
	MonitorHealthResponseProto
	TransitionToActiveRequestProto
	TransitionToActiveResponseProto
	TransitionToStandbyRequestProto
	TransitionToStandbyResponseProto
	GetServiceStatusRequestProto
	GetServiceStatusResponseProto
	GetProtocolVersionsRequestProto
	ProtocolVersionProto
	GetProtocolVersionsResponseProto
	GetProtocolSignatureRequestProto
	GetProtocolSignatureResponseProto
	ProtocolSignatureProto
*/
package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Spec for UserInformationProto is specified in ProtoUtil#makeIpcConnectionContext
type UserInformationProto struct {
	EffectiveUser    *string `protobuf:"bytes,1,opt,name=effectiveUser" json:"effectiveUser,omitempty"`
	RealUser         *string `protobuf:"bytes,2,opt,name=realUser" json:"realUser,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserInformationProto) Reset()                    { *m = UserInformationProto{} }
func (m *UserInformationProto) String() string            { return proto.CompactTextString(m) }
func (*UserInformationProto) ProtoMessage()               {}
func (*UserInformationProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInformationProto) GetEffectiveUser() string {
	if m != nil && m.EffectiveUser != nil {
		return *m.EffectiveUser
	}
	return ""
}

func (m *UserInformationProto) GetRealUser() string {
	if m != nil && m.RealUser != nil {
		return *m.RealUser
	}
	return ""
}

// *
// The connection context is sent as part of the connection establishment.
// It establishes the context for ALL Rpc calls within the connection.
type IpcConnectionContextProto struct {
	// UserInfo beyond what is determined as part of security handshake
	// at connection time (kerberos, tokens etc).
	UserInfo *UserInformationProto `protobuf:"bytes,2,opt,name=userInfo" json:"userInfo,omitempty"`
	// Protocol name for next rpc layer.
	// The client created a proxy with this protocol name
	Protocol         *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IpcConnectionContextProto) Reset()                    { *m = IpcConnectionContextProto{} }
func (m *IpcConnectionContextProto) String() string            { return proto.CompactTextString(m) }
func (*IpcConnectionContextProto) ProtoMessage()               {}
func (*IpcConnectionContextProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IpcConnectionContextProto) GetUserInfo() *UserInformationProto {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *IpcConnectionContextProto) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf2, 0x2c, 0x48, 0x76,
	0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x73, 0xce, 0xcf, 0x2b, 0x49, 0xad, 0x28,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0x48, 0x4c, 0xc9, 0xcf, 0x2f, 0xd0, 0x4b,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x53, 0x8a, 0xe0, 0x12, 0x09, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x4b,
	0xcb, 0x2f, 0xca, 0x4d, 0x04, 0x29, 0x0f, 0x00, 0x2b, 0x53, 0xe1, 0xe2, 0x4d, 0x4d, 0x4b, 0x03,
	0x19, 0x50, 0x96, 0x0a, 0x52, 0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x2a, 0x28, 0x24,
	0xc5, 0xc5, 0x51, 0x94, 0x9a, 0x98, 0x03, 0x56, 0xc0, 0x04, 0x56, 0x00, 0xe7, 0x2b, 0x55, 0x70,
	0x49, 0x62, 0x73, 0x06, 0xc4, 0x78, 0x7b, 0x2e, 0x8e, 0x52, 0xa8, 0xb5, 0x60, 0x8d, 0xdc, 0x46,
	0xca, 0x7a, 0x28, 0x0e, 0xd3, 0xc3, 0xe6, 0xaa, 0x20, 0xb8, 0x26, 0x90, 0xcd, 0x60, 0xff, 0x24,
	0xe7, 0xe7, 0x48, 0x30, 0x43, 0x6c, 0x86, 0xf1, 0x9d, 0xec, 0xb9, 0xe4, 0xf2, 0x8b, 0xd2, 0xf5,
	0x12, 0x0b, 0x12, 0x93, 0x33, 0x52, 0x61, 0xc6, 0x66, 0x16, 0x24, 0x43, 0x42, 0x20, 0xa9, 0x34,
	0xcd, 0x49, 0x0a, 0xa7, 0xcb, 0x8a, 0x17, 0x30, 0x32, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x05, 0xde, 0x49, 0x40, 0x01, 0x00, 0x00,
}
