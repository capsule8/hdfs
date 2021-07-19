// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RpcHeader.proto

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// RpcKind determine the rpcEngine and the serialization of the rpc request
type RpcKindProto int32

const (
	RpcKindProto_RPC_BUILTIN         RpcKindProto = 0
	RpcKindProto_RPC_WRITABLE        RpcKindProto = 1
	RpcKindProto_RPC_PROTOCOL_BUFFER RpcKindProto = 2
)

var RpcKindProto_name = map[int32]string{
	0: "RPC_BUILTIN",
	1: "RPC_WRITABLE",
	2: "RPC_PROTOCOL_BUFFER",
}
var RpcKindProto_value = map[string]int32{
	"RPC_BUILTIN":         0,
	"RPC_WRITABLE":        1,
	"RPC_PROTOCOL_BUFFER": 2,
}

func (x RpcKindProto) Enum() *RpcKindProto {
	p := new(RpcKindProto)
	*p = x
	return p
}
func (x RpcKindProto) String() string {
	return proto.EnumName(RpcKindProto_name, int32(x))
}
func (x *RpcKindProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcKindProto_value, data, "RpcKindProto")
	if err != nil {
		return err
	}
	*x = RpcKindProto(value)
	return nil
}
func (RpcKindProto) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type RpcRequestHeaderProto_OperationProto int32

const (
	RpcRequestHeaderProto_RPC_FINAL_PACKET        RpcRequestHeaderProto_OperationProto = 0
	RpcRequestHeaderProto_RPC_CONTINUATION_PACKET RpcRequestHeaderProto_OperationProto = 1
	RpcRequestHeaderProto_RPC_CLOSE_CONNECTION    RpcRequestHeaderProto_OperationProto = 2
)

var RpcRequestHeaderProto_OperationProto_name = map[int32]string{
	0: "RPC_FINAL_PACKET",
	1: "RPC_CONTINUATION_PACKET",
	2: "RPC_CLOSE_CONNECTION",
}
var RpcRequestHeaderProto_OperationProto_value = map[string]int32{
	"RPC_FINAL_PACKET":        0,
	"RPC_CONTINUATION_PACKET": 1,
	"RPC_CLOSE_CONNECTION":    2,
}

func (x RpcRequestHeaderProto_OperationProto) Enum() *RpcRequestHeaderProto_OperationProto {
	p := new(RpcRequestHeaderProto_OperationProto)
	*p = x
	return p
}
func (x RpcRequestHeaderProto_OperationProto) String() string {
	return proto.EnumName(RpcRequestHeaderProto_OperationProto_name, int32(x))
}
func (x *RpcRequestHeaderProto_OperationProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcRequestHeaderProto_OperationProto_value, data, "RpcRequestHeaderProto_OperationProto")
	if err != nil {
		return err
	}
	*x = RpcRequestHeaderProto_OperationProto(value)
	return nil
}
func (RpcRequestHeaderProto_OperationProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{2, 0}
}

type RpcResponseHeaderProto_RpcStatusProto int32

const (
	RpcResponseHeaderProto_SUCCESS RpcResponseHeaderProto_RpcStatusProto = 0
	RpcResponseHeaderProto_ERROR   RpcResponseHeaderProto_RpcStatusProto = 1
	RpcResponseHeaderProto_FATAL   RpcResponseHeaderProto_RpcStatusProto = 2
)

var RpcResponseHeaderProto_RpcStatusProto_name = map[int32]string{
	0: "SUCCESS",
	1: "ERROR",
	2: "FATAL",
}
var RpcResponseHeaderProto_RpcStatusProto_value = map[string]int32{
	"SUCCESS": 0,
	"ERROR":   1,
	"FATAL":   2,
}

func (x RpcResponseHeaderProto_RpcStatusProto) Enum() *RpcResponseHeaderProto_RpcStatusProto {
	p := new(RpcResponseHeaderProto_RpcStatusProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcStatusProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcStatusProto_name, int32(x))
}
func (x *RpcResponseHeaderProto_RpcStatusProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcStatusProto_value, data, "RpcResponseHeaderProto_RpcStatusProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcStatusProto(value)
	return nil
}
func (RpcResponseHeaderProto_RpcStatusProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{3, 0}
}

type RpcResponseHeaderProto_RpcErrorCodeProto int32

const (
	// Non-fatal Rpc error - connection left open for future rpc calls
	RpcResponseHeaderProto_ERROR_APPLICATION          RpcResponseHeaderProto_RpcErrorCodeProto = 1
	RpcResponseHeaderProto_ERROR_NO_SUCH_METHOD       RpcResponseHeaderProto_RpcErrorCodeProto = 2
	RpcResponseHeaderProto_ERROR_NO_SUCH_PROTOCOL     RpcResponseHeaderProto_RpcErrorCodeProto = 3
	RpcResponseHeaderProto_ERROR_RPC_SERVER           RpcResponseHeaderProto_RpcErrorCodeProto = 4
	RpcResponseHeaderProto_ERROR_SERIALIZING_RESPONSE RpcResponseHeaderProto_RpcErrorCodeProto = 5
	RpcResponseHeaderProto_ERROR_RPC_VERSION_MISMATCH RpcResponseHeaderProto_RpcErrorCodeProto = 6
	// Fatal Server side Rpc error - connection closed
	RpcResponseHeaderProto_FATAL_UNKNOWN                   RpcResponseHeaderProto_RpcErrorCodeProto = 10
	RpcResponseHeaderProto_FATAL_UNSUPPORTED_SERIALIZATION RpcResponseHeaderProto_RpcErrorCodeProto = 11
	RpcResponseHeaderProto_FATAL_INVALID_RPC_HEADER        RpcResponseHeaderProto_RpcErrorCodeProto = 12
	RpcResponseHeaderProto_FATAL_DESERIALIZING_REQUEST     RpcResponseHeaderProto_RpcErrorCodeProto = 13
	RpcResponseHeaderProto_FATAL_VERSION_MISMATCH          RpcResponseHeaderProto_RpcErrorCodeProto = 14
	RpcResponseHeaderProto_FATAL_UNAUTHORIZED              RpcResponseHeaderProto_RpcErrorCodeProto = 15
)

var RpcResponseHeaderProto_RpcErrorCodeProto_name = map[int32]string{
	1:  "ERROR_APPLICATION",
	2:  "ERROR_NO_SUCH_METHOD",
	3:  "ERROR_NO_SUCH_PROTOCOL",
	4:  "ERROR_RPC_SERVER",
	5:  "ERROR_SERIALIZING_RESPONSE",
	6:  "ERROR_RPC_VERSION_MISMATCH",
	10: "FATAL_UNKNOWN",
	11: "FATAL_UNSUPPORTED_SERIALIZATION",
	12: "FATAL_INVALID_RPC_HEADER",
	13: "FATAL_DESERIALIZING_REQUEST",
	14: "FATAL_VERSION_MISMATCH",
	15: "FATAL_UNAUTHORIZED",
}
var RpcResponseHeaderProto_RpcErrorCodeProto_value = map[string]int32{
	"ERROR_APPLICATION":               1,
	"ERROR_NO_SUCH_METHOD":            2,
	"ERROR_NO_SUCH_PROTOCOL":          3,
	"ERROR_RPC_SERVER":                4,
	"ERROR_SERIALIZING_RESPONSE":      5,
	"ERROR_RPC_VERSION_MISMATCH":      6,
	"FATAL_UNKNOWN":                   10,
	"FATAL_UNSUPPORTED_SERIALIZATION": 11,
	"FATAL_INVALID_RPC_HEADER":        12,
	"FATAL_DESERIALIZING_REQUEST":     13,
	"FATAL_VERSION_MISMATCH":          14,
	"FATAL_UNAUTHORIZED":              15,
}

func (x RpcResponseHeaderProto_RpcErrorCodeProto) Enum() *RpcResponseHeaderProto_RpcErrorCodeProto {
	p := new(RpcResponseHeaderProto_RpcErrorCodeProto)
	*p = x
	return p
}
func (x RpcResponseHeaderProto_RpcErrorCodeProto) String() string {
	return proto.EnumName(RpcResponseHeaderProto_RpcErrorCodeProto_name, int32(x))
}
func (x *RpcResponseHeaderProto_RpcErrorCodeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcResponseHeaderProto_RpcErrorCodeProto_value, data, "RpcResponseHeaderProto_RpcErrorCodeProto")
	if err != nil {
		return err
	}
	*x = RpcResponseHeaderProto_RpcErrorCodeProto(value)
	return nil
}
func (RpcResponseHeaderProto_RpcErrorCodeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{3, 1}
}

type RpcSaslProto_SaslState int32

const (
	RpcSaslProto_SUCCESS   RpcSaslProto_SaslState = 0
	RpcSaslProto_NEGOTIATE RpcSaslProto_SaslState = 1
	RpcSaslProto_INITIATE  RpcSaslProto_SaslState = 2
	RpcSaslProto_CHALLENGE RpcSaslProto_SaslState = 3
	RpcSaslProto_RESPONSE  RpcSaslProto_SaslState = 4
	RpcSaslProto_WRAP      RpcSaslProto_SaslState = 5
)

var RpcSaslProto_SaslState_name = map[int32]string{
	0: "SUCCESS",
	1: "NEGOTIATE",
	2: "INITIATE",
	3: "CHALLENGE",
	4: "RESPONSE",
	5: "WRAP",
}
var RpcSaslProto_SaslState_value = map[string]int32{
	"SUCCESS":   0,
	"NEGOTIATE": 1,
	"INITIATE":  2,
	"CHALLENGE": 3,
	"RESPONSE":  4,
	"WRAP":      5,
}

func (x RpcSaslProto_SaslState) Enum() *RpcSaslProto_SaslState {
	p := new(RpcSaslProto_SaslState)
	*p = x
	return p
}
func (x RpcSaslProto_SaslState) String() string {
	return proto.EnumName(RpcSaslProto_SaslState_name, int32(x))
}
func (x *RpcSaslProto_SaslState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcSaslProto_SaslState_value, data, "RpcSaslProto_SaslState")
	if err != nil {
		return err
	}
	*x = RpcSaslProto_SaslState(value)
	return nil
}
func (RpcSaslProto_SaslState) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{4, 0} }

// *
// Used to pass through the information necessary to continue
// a trace after an RPC is made. All we need is the traceid
// (so we know the overarching trace this message is a part of), and
// the id of the current span when this message was sent, so we know
// what span caused the new span we will create when this message is received.
type RPCTraceInfoProto struct {
	TraceId          *int64 `protobuf:"varint,1,opt,name=traceId" json:"traceId,omitempty"`
	ParentId         *int64 `protobuf:"varint,2,opt,name=parentId" json:"parentId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RPCTraceInfoProto) Reset()                    { *m = RPCTraceInfoProto{} }
func (m *RPCTraceInfoProto) String() string            { return proto.CompactTextString(m) }
func (*RPCTraceInfoProto) ProtoMessage()               {}
func (*RPCTraceInfoProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *RPCTraceInfoProto) GetTraceId() int64 {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return 0
}

func (m *RPCTraceInfoProto) GetParentId() int64 {
	if m != nil && m.ParentId != nil {
		return *m.ParentId
	}
	return 0
}

// *
// Used to pass through the call context entry after an RPC is made.
type RPCCallerContextProto struct {
	Context          *string `protobuf:"bytes,1,req,name=context" json:"context,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RPCCallerContextProto) Reset()                    { *m = RPCCallerContextProto{} }
func (m *RPCCallerContextProto) String() string            { return proto.CompactTextString(m) }
func (*RPCCallerContextProto) ProtoMessage()               {}
func (*RPCCallerContextProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *RPCCallerContextProto) GetContext() string {
	if m != nil && m.Context != nil {
		return *m.Context
	}
	return ""
}

func (m *RPCCallerContextProto) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type RpcRequestHeaderProto struct {
	RpcKind  *RpcKindProto                         `protobuf:"varint,1,opt,name=rpcKind,enum=hadoop.common.RpcKindProto" json:"rpcKind,omitempty"`
	RpcOp    *RpcRequestHeaderProto_OperationProto `protobuf:"varint,2,opt,name=rpcOp,enum=hadoop.common.RpcRequestHeaderProto_OperationProto" json:"rpcOp,omitempty"`
	CallId   *int32                                `protobuf:"zigzag32,3,req,name=callId" json:"callId,omitempty"`
	ClientId []byte                                `protobuf:"bytes,4,req,name=clientId" json:"clientId,omitempty"`
	// clientId + callId uniquely identifies a request
	// retry count, 1 means this is the first retry
	RetryCount       *int32                 `protobuf:"zigzag32,5,opt,name=retryCount,def=-1" json:"retryCount,omitempty"`
	TraceInfo        *RPCTraceInfoProto     `protobuf:"bytes,6,opt,name=traceInfo" json:"traceInfo,omitempty"`
	CallerContext    *RPCCallerContextProto `protobuf:"bytes,7,opt,name=callerContext" json:"callerContext,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *RpcRequestHeaderProto) Reset()                    { *m = RpcRequestHeaderProto{} }
func (m *RpcRequestHeaderProto) String() string            { return proto.CompactTextString(m) }
func (*RpcRequestHeaderProto) ProtoMessage()               {}
func (*RpcRequestHeaderProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

const Default_RpcRequestHeaderProto_RetryCount int32 = -1

func (m *RpcRequestHeaderProto) GetRpcKind() RpcKindProto {
	if m != nil && m.RpcKind != nil {
		return *m.RpcKind
	}
	return RpcKindProto_RPC_BUILTIN
}

func (m *RpcRequestHeaderProto) GetRpcOp() RpcRequestHeaderProto_OperationProto {
	if m != nil && m.RpcOp != nil {
		return *m.RpcOp
	}
	return RpcRequestHeaderProto_RPC_FINAL_PACKET
}

func (m *RpcRequestHeaderProto) GetCallId() int32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcRequestHeaderProto) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *RpcRequestHeaderProto) GetRetryCount() int32 {
	if m != nil && m.RetryCount != nil {
		return *m.RetryCount
	}
	return Default_RpcRequestHeaderProto_RetryCount
}

func (m *RpcRequestHeaderProto) GetTraceInfo() *RPCTraceInfoProto {
	if m != nil {
		return m.TraceInfo
	}
	return nil
}

func (m *RpcRequestHeaderProto) GetCallerContext() *RPCCallerContextProto {
	if m != nil {
		return m.CallerContext
	}
	return nil
}

// *
// Rpc Response Header
// +------------------------------------------------------------------+
// | Rpc total response length in bytes (4 bytes int)                 |
// |  (sum of next two parts)                                         |
// +------------------------------------------------------------------+
// | RpcResponseHeaderProto - serialized delimited ie has len         |
// +------------------------------------------------------------------+
// | if request is successful:                                        |
// |   - RpcResponse -  The actual rpc response  bytes follow         |
// |     the response header                                          |
// |     This response is serialized based on RpcKindProto            |
// | if request fails :                                               |
// |   The rpc response header contains the necessary info            |
// +------------------------------------------------------------------+
//
// Note that rpc response header is also used when connection setup fails.
// Ie the response looks like a rpc response with a fake callId.
type RpcResponseHeaderProto struct {
	CallId              *uint32                                   `protobuf:"varint,1,req,name=callId" json:"callId,omitempty"`
	Status              *RpcResponseHeaderProto_RpcStatusProto    `protobuf:"varint,2,req,name=status,enum=hadoop.common.RpcResponseHeaderProto_RpcStatusProto" json:"status,omitempty"`
	ServerIpcVersionNum *uint32                                   `protobuf:"varint,3,opt,name=serverIpcVersionNum" json:"serverIpcVersionNum,omitempty"`
	ExceptionClassName  *string                                   `protobuf:"bytes,4,opt,name=exceptionClassName" json:"exceptionClassName,omitempty"`
	ErrorMsg            *string                                   `protobuf:"bytes,5,opt,name=errorMsg" json:"errorMsg,omitempty"`
	ErrorDetail         *RpcResponseHeaderProto_RpcErrorCodeProto `protobuf:"varint,6,opt,name=errorDetail,enum=hadoop.common.RpcResponseHeaderProto_RpcErrorCodeProto" json:"errorDetail,omitempty"`
	ClientId            []byte                                    `protobuf:"bytes,7,opt,name=clientId" json:"clientId,omitempty"`
	RetryCount          *int32                                    `protobuf:"zigzag32,8,opt,name=retryCount,def=-1" json:"retryCount,omitempty"`
	XXX_unrecognized    []byte                                    `json:"-"`
}

func (m *RpcResponseHeaderProto) Reset()                    { *m = RpcResponseHeaderProto{} }
func (m *RpcResponseHeaderProto) String() string            { return proto.CompactTextString(m) }
func (*RpcResponseHeaderProto) ProtoMessage()               {}
func (*RpcResponseHeaderProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

const Default_RpcResponseHeaderProto_RetryCount int32 = -1

func (m *RpcResponseHeaderProto) GetCallId() uint32 {
	if m != nil && m.CallId != nil {
		return *m.CallId
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetStatus() RpcResponseHeaderProto_RpcStatusProto {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return RpcResponseHeaderProto_SUCCESS
}

func (m *RpcResponseHeaderProto) GetServerIpcVersionNum() uint32 {
	if m != nil && m.ServerIpcVersionNum != nil {
		return *m.ServerIpcVersionNum
	}
	return 0
}

func (m *RpcResponseHeaderProto) GetExceptionClassName() string {
	if m != nil && m.ExceptionClassName != nil {
		return *m.ExceptionClassName
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorMsg() string {
	if m != nil && m.ErrorMsg != nil {
		return *m.ErrorMsg
	}
	return ""
}

func (m *RpcResponseHeaderProto) GetErrorDetail() RpcResponseHeaderProto_RpcErrorCodeProto {
	if m != nil && m.ErrorDetail != nil {
		return *m.ErrorDetail
	}
	return RpcResponseHeaderProto_ERROR_APPLICATION
}

func (m *RpcResponseHeaderProto) GetClientId() []byte {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *RpcResponseHeaderProto) GetRetryCount() int32 {
	if m != nil && m.RetryCount != nil {
		return *m.RetryCount
	}
	return Default_RpcResponseHeaderProto_RetryCount
}

type RpcSaslProto struct {
	Version          *uint32                  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	State            *RpcSaslProto_SaslState  `protobuf:"varint,2,req,name=state,enum=hadoop.common.RpcSaslProto_SaslState" json:"state,omitempty"`
	Token            []byte                   `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	Auths            []*RpcSaslProto_SaslAuth `protobuf:"bytes,4,rep,name=auths" json:"auths,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *RpcSaslProto) Reset()                    { *m = RpcSaslProto{} }
func (m *RpcSaslProto) String() string            { return proto.CompactTextString(m) }
func (*RpcSaslProto) ProtoMessage()               {}
func (*RpcSaslProto) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *RpcSaslProto) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *RpcSaslProto) GetState() RpcSaslProto_SaslState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RpcSaslProto_SUCCESS
}

func (m *RpcSaslProto) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RpcSaslProto) GetAuths() []*RpcSaslProto_SaslAuth {
	if m != nil {
		return m.Auths
	}
	return nil
}

type RpcSaslProto_SaslAuth struct {
	Method           *string `protobuf:"bytes,1,req,name=method" json:"method,omitempty"`
	Mechanism        *string `protobuf:"bytes,2,req,name=mechanism" json:"mechanism,omitempty"`
	Protocol         *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
	ServerId         *string `protobuf:"bytes,4,opt,name=serverId" json:"serverId,omitempty"`
	Challenge        []byte  `protobuf:"bytes,5,opt,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpcSaslProto_SaslAuth) Reset()                    { *m = RpcSaslProto_SaslAuth{} }
func (m *RpcSaslProto_SaslAuth) String() string            { return proto.CompactTextString(m) }
func (*RpcSaslProto_SaslAuth) ProtoMessage()               {}
func (*RpcSaslProto_SaslAuth) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4, 0} }

func (m *RpcSaslProto_SaslAuth) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetMechanism() string {
	if m != nil && m.Mechanism != nil {
		return *m.Mechanism
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetServerId() string {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return ""
}

func (m *RpcSaslProto_SaslAuth) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

var fileDescriptor4 = []byte{
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x3b, 0x49, 0xdb, 0x9c, 0x26, 0xa9, 0x33, 0xdb, 0x76, 0xad, 0x76, 0xb5, 0x1b, 0x05,
	0x90, 0x22, 0x24, 0x22, 0xe8, 0x2e, 0x42, 0x5a, 0x24, 0x24, 0xd7, 0x99, 0x36, 0xb3, 0x75, 0x6d,
	0x33, 0x76, 0x5a, 0xb1, 0x02, 0x45, 0xc6, 0x99, 0x6d, 0x22, 0x12, 0xdb, 0xd8, 0xce, 0x6a, 0x79,
	0x11, 0xae, 0x79, 0x02, 0xc4, 0x03, 0x71, 0xc3, 0x9b, 0xa0, 0x99, 0x89, 0x9d, 0xa6, 0x29, 0x12,
	0x77, 0x73, 0xfe, 0xbe, 0x39, 0x3f, 0xdf, 0x99, 0x81, 0x03, 0x9a, 0x84, 0x43, 0x16, 0x4c, 0x58,
	0xda, 0x4f, 0xd2, 0x38, 0x8f, 0x51, 0x73, 0x1a, 0x4c, 0xe2, 0x38, 0xe9, 0x87, 0xf1, 0x62, 0x11,
	0x47, 0x5d, 0x02, 0x6d, 0xea, 0x9a, 0x7e, 0x1a, 0x84, 0x8c, 0x44, 0xef, 0x63, 0x57, 0xf8, 0xe8,
	0xb0, 0x9b, 0x0b, 0xcd, 0x44, 0x57, 0x3a, 0x4a, 0xaf, 0x42, 0x0b, 0x11, 0x9d, 0xc0, 0x5e, 0x12,
	0xa4, 0x2c, 0xca, 0xc9, 0x44, 0x57, 0x85, 0xa9, 0x94, 0xbb, 0x0e, 0x1c, 0x51, 0xd7, 0x34, 0x83,
	0xf9, 0x9c, 0xa5, 0x66, 0x1c, 0xe5, 0xec, 0x63, 0x5e, 0xc2, 0x85, 0x52, 0xd6, 0x95, 0x8e, 0xda,
	0xab, 0xd3, 0x42, 0x44, 0xcf, 0xa1, 0x9e, 0xcd, 0xee, 0xa2, 0x20, 0x5f, 0xa6, 0x4c, 0xe0, 0x35,
	0xe8, 0x5a, 0xd1, 0xfd, 0xbb, 0x02, 0x47, 0x34, 0x09, 0x29, 0xfb, 0x75, 0xc9, 0xb2, 0x5c, 0x56,
	0x21, 0x11, 0xbf, 0x86, 0xdd, 0x34, 0x09, 0xaf, 0x66, 0x91, 0x4c, 0xb0, 0x75, 0x76, 0xda, 0xdf,
	0x28, 0xab, 0x4f, 0xa5, 0x55, 0x78, 0xd3, 0xc2, 0x17, 0x11, 0xa8, 0xa5, 0x49, 0xe8, 0x24, 0xe2,
	0xaa, 0xd6, 0xd9, 0xab, 0xed, 0xa0, 0xed, 0xbb, 0xfa, 0x4e, 0xc2, 0xd2, 0x20, 0x9f, 0xc5, 0x91,
	0x04, 0x93, 0x08, 0xe8, 0x18, 0x76, 0xc2, 0x60, 0x3e, 0x27, 0x13, 0xbd, 0xd2, 0x51, 0x7b, 0x6d,
	0xba, 0x92, 0x78, 0x83, 0xc2, 0xf9, 0x4c, 0x36, 0xa8, 0xda, 0x51, 0x7b, 0x0d, 0x5a, 0xca, 0xa8,
	0x0b, 0x90, 0xb2, 0x3c, 0xfd, 0xcd, 0x8c, 0x97, 0x51, 0xae, 0xd7, 0x3a, 0x4a, 0xaf, 0xfd, 0x46,
	0xfd, 0xe2, 0x2b, 0x7a, 0x4f, 0x8b, 0xbe, 0x83, 0x7a, 0x5e, 0x0c, 0x43, 0xdf, 0xe9, 0x28, 0xbd,
	0xfd, 0xb3, 0xce, 0xc3, 0x34, 0x1f, 0xce, 0x8b, 0xae, 0x43, 0xd0, 0x5b, 0x68, 0x86, 0xf7, 0x27,
	0xa0, 0xef, 0x0a, 0x8c, 0x4f, 0xb7, 0x31, 0xb6, 0x07, 0x45, 0x37, 0x43, 0xbb, 0x3f, 0x41, 0x6b,
	0xb3, 0x78, 0x74, 0x08, 0x1a, 0x75, 0xcd, 0xf1, 0x05, 0xb1, 0x0d, 0x6b, 0xec, 0x1a, 0xe6, 0x15,
	0xf6, 0xb5, 0x27, 0xe8, 0x14, 0x9e, 0x71, 0xad, 0xe9, 0xd8, 0x3e, 0xb1, 0x47, 0x86, 0x4f, 0x1c,
	0xbb, 0x30, 0x2a, 0x48, 0x87, 0x43, 0x61, 0xb4, 0x1c, 0x0f, 0x73, 0x17, 0x1b, 0x9b, 0xdc, 0x41,
	0x53, 0xbb, 0x7f, 0xed, 0xc0, 0xb1, 0x68, 0x79, 0x96, 0xc4, 0x51, 0xc6, 0xee, 0xcf, 0x77, 0xdd,
	0x5d, 0x4e, 0x98, 0x66, 0xd9, 0x5d, 0x0b, 0x76, 0xb2, 0x3c, 0xc8, 0x97, 0x99, 0xae, 0x76, 0xd4,
	0x5e, 0xeb, 0xec, 0xf5, 0x63, 0x13, 0xdc, 0x82, 0xe3, 0x6a, 0x4f, 0x84, 0xc9, 0x32, 0x57, 0x18,
	0xe8, 0x4b, 0x78, 0x9a, 0xb1, 0xf4, 0x03, 0x4b, 0x49, 0x12, 0xde, 0xb0, 0x34, 0x9b, 0xc5, 0x91,
	0xbd, 0x5c, 0xe8, 0x95, 0x8e, 0xd2, 0x6b, 0xd2, 0xc7, 0x4c, 0xa8, 0x0f, 0x88, 0x7d, 0x0c, 0x59,
	0xc2, 0x3b, 0x62, 0xce, 0x83, 0x2c, 0xb3, 0x83, 0x05, 0xd3, 0xab, 0x1d, 0xa5, 0x57, 0xa7, 0x8f,
	0x58, 0x38, 0x1b, 0x58, 0x9a, 0xc6, 0xe9, 0x75, 0x76, 0x27, 0xe6, 0x5d, 0xa7, 0xa5, 0x8c, 0x7e,
	0x80, 0x7d, 0x71, 0x1e, 0xb0, 0x3c, 0x98, 0xcd, 0xc5, 0xac, 0x5b, 0x67, 0xdf, 0xfc, 0xef, 0x82,
	0x30, 0x8f, 0x35, 0xe3, 0x09, 0x93, 0x35, 0xdd, 0xc7, 0xda, 0x20, 0xe1, 0xae, 0xd8, 0xaa, 0xff,
	0x22, 0xe1, 0xde, 0x63, 0x24, 0xec, 0xbe, 0x82, 0xd6, 0x66, 0xcb, 0xd0, 0x3e, 0xec, 0x7a, 0x23,
	0xd3, 0xc4, 0x9e, 0xa7, 0x3d, 0x41, 0x75, 0xa8, 0x61, 0x4a, 0x1d, 0xaa, 0x29, 0xfc, 0x78, 0x61,
	0xf8, 0x86, 0xa5, 0xa9, 0xdd, 0x7f, 0x54, 0x68, 0x6f, 0xe5, 0x85, 0x8e, 0xa0, 0x2d, 0x7c, 0xc7,
	0x86, 0xeb, 0x5a, 0xc4, 0x14, 0xe4, 0x90, 0xac, 0x90, 0x6a, 0xdb, 0x19, 0x7b, 0x23, 0x73, 0x38,
	0xbe, 0xc6, 0xfe, 0xd0, 0x19, 0x68, 0x2a, 0x3a, 0x81, 0xe3, 0x4d, 0x8b, 0x4b, 0x1d, 0xdf, 0x31,
	0x1d, 0x4b, 0xab, 0x70, 0xfa, 0x49, 0x1b, 0x67, 0x94, 0x87, 0xe9, 0x0d, 0xa6, 0x5a, 0x15, 0xbd,
	0x80, 0x13, 0xa9, 0xf5, 0x30, 0x25, 0x86, 0x45, 0xde, 0x11, 0xfb, 0x72, 0x4c, 0xb1, 0xe7, 0x3a,
	0xb6, 0x87, 0xb5, 0xda, 0xda, 0xce, 0xa3, 0x6e, 0x30, 0xf5, 0x38, 0x3f, 0xaf, 0x89, 0x77, 0x6d,
	0xf8, 0xe6, 0x50, 0xdb, 0x41, 0x6d, 0x68, 0x8a, 0x1a, 0xc6, 0x23, 0xfb, 0xca, 0x76, 0x6e, 0x6d,
	0x0d, 0xd0, 0x27, 0xf0, 0xb2, 0x50, 0x79, 0x23, 0xd7, 0x75, 0xa8, 0x8f, 0x07, 0x25, 0xbc, 0xac,
	0x61, 0x1f, 0x3d, 0x07, 0x5d, 0x3a, 0x11, 0xfb, 0xc6, 0xb0, 0xc8, 0x40, 0xe0, 0x0f, 0xb1, 0x31,
	0xc0, 0x54, 0x6b, 0xa0, 0x97, 0x70, 0x2a, 0xad, 0x03, 0xbc, 0x99, 0xd7, 0xf7, 0x23, 0xec, 0xf9,
	0x5a, 0x93, 0x17, 0x2a, 0x1d, 0xb6, 0x52, 0x6a, 0xa1, 0x63, 0x40, 0xc5, 0xfd, 0xc6, 0xc8, 0x1f,
	0x3a, 0x94, 0xbc, 0xc3, 0x03, 0xed, 0xa0, 0xfb, 0x67, 0x05, 0x1a, 0x7c, 0x32, 0x41, 0x36, 0x2f,
	0x9f, 0xd6, 0x0f, 0x92, 0x9e, 0xe2, 0x21, 0x6c, 0xd2, 0x42, 0x44, 0xdf, 0x42, 0x8d, 0xd3, 0x9c,
	0xad, 0x36, 0xe5, 0xb3, 0x6d, 0x62, 0x95, 0x28, 0x7d, 0x7e, 0xe2, 0xd3, 0x66, 0x54, 0xc6, 0xa0,
	0x43, 0xa8, 0xe5, 0xf1, 0x2f, 0x2c, 0x12, 0xbb, 0xd0, 0xa0, 0x52, 0x40, 0x6f, 0xa0, 0x16, 0x2c,
	0xf3, 0x69, 0xa6, 0x57, 0x3b, 0x95, 0xc7, 0xde, 0x94, 0x87, 0x90, 0xc6, 0x32, 0x9f, 0x52, 0x19,
	0x72, 0xf2, 0xbb, 0x02, 0x7b, 0x85, 0x8e, 0xaf, 0xf7, 0x82, 0xe5, 0xd3, 0x78, 0xb2, 0xfa, 0x0f,
	0x56, 0x12, 0xff, 0x0e, 0x16, 0x2c, 0x9c, 0x06, 0xd1, 0x2c, 0x5b, 0x88, 0xbc, 0xeb, 0x74, 0xad,
	0x10, 0x7f, 0x0f, 0xc7, 0x0e, 0xe3, 0xb9, 0xc8, 0xab, 0x4e, 0x4b, 0x99, 0xdb, 0x56, 0xfb, 0x3a,
	0x59, 0xad, 0x63, 0x29, 0x73, 0xd4, 0x70, 0xca, 0x1f, 0xb6, 0xe8, 0x8e, 0x89, 0x2d, 0x6c, 0xd0,
	0xb5, 0xa2, 0xfb, 0x23, 0xd4, 0xcb, 0xf2, 0x37, 0x69, 0xde, 0x84, 0xba, 0x8d, 0x2f, 0x1d, 0x9f,
	0x18, 0x3e, 0xd6, 0x14, 0xd4, 0x80, 0x3d, 0x62, 0x13, 0x29, 0xa9, 0xdc, 0x68, 0x0e, 0x0d, 0xcb,
	0xc2, 0xf6, 0x25, 0xd6, 0x2a, 0xdc, 0x58, 0x32, 0xae, 0x8a, 0xf6, 0xa0, 0x7a, 0x4b, 0x0d, 0x57,
	0xab, 0x7d, 0xfe, 0x56, 0xcc, 0xab, 0xfc, 0x8a, 0xd0, 0x01, 0xec, 0x73, 0x96, 0x9c, 0x8f, 0x88,
	0xe5, 0x13, 0x5b, 0x7b, 0x82, 0x34, 0x68, 0x70, 0xc5, 0x2d, 0x25, 0xbe, 0x71, 0x6e, 0xf1, 0x7b,
	0x9e, 0xc1, 0x53, 0xae, 0x29, 0x68, 0x3f, 0x3e, 0x1f, 0x5d, 0x5c, 0x60, 0xaa, 0xa9, 0xe7, 0xaf,
	0xe1, 0x45, 0x9c, 0xde, 0xf5, 0x83, 0x24, 0x08, 0xa7, 0xac, 0xe8, 0xfd, 0x2c, 0x09, 0xe5, 0xc7,
	0xfe, 0xf3, 0xf2, 0xfd, 0xf9, 0xfa, 0xb3, 0x17, 0xb7, 0x65, 0x7f, 0x28, 0xca, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x60, 0xab, 0xef, 0x01, 0x08, 0x00, 0x00,
}