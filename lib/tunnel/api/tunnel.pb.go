// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tunnel.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Frame struct {
	// Types that are valid to be assigned to Message:
	//	*Frame_DialRequest
	//	*Frame_Data
	Message              isFrame_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}
func (m *Frame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return m.Size()
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

type isFrame_Message interface {
	isFrame_Message()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Frame_DialRequest struct {
	DialRequest *DialRequest `protobuf:"bytes,1,opt,name=DialRequest,proto3,oneof" json:"DialRequest,omitempty"`
}
type Frame_Data struct {
	Data *Data `protobuf:"bytes,2,opt,name=Data,proto3,oneof" json:"Data,omitempty"`
}

func (*Frame_DialRequest) isFrame_Message() {}
func (*Frame_Data) isFrame_Message()        {}

func (m *Frame) GetMessage() isFrame_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Frame) GetDialRequest() *DialRequest {
	if x, ok := m.GetMessage().(*Frame_DialRequest); ok {
		return x.DialRequest
	}
	return nil
}

func (m *Frame) GetData() *Data {
	if x, ok := m.GetMessage().(*Frame_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Frame) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Frame_DialRequest)(nil),
		(*Frame_Data)(nil),
	}
}

type DialRequest struct {
	ServerID             string   `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	ConnType             string   `protobuf:"bytes,2,opt,name=ConnType,proto3" json:"ConnType,omitempty"`
	To                   *Addr    `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DialRequest) Reset()         { *m = DialRequest{} }
func (m *DialRequest) String() string { return proto.CompactTextString(m) }
func (*DialRequest) ProtoMessage()    {}
func (*DialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}
func (m *DialRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DialRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DialRequest.Merge(m, src)
}
func (m *DialRequest) XXX_Size() int {
	return m.Size()
}
func (m *DialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DialRequest proto.InternalMessageInfo

func (m *DialRequest) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *DialRequest) GetConnType() string {
	if m != nil {
		return m.ConnType
	}
	return ""
}

func (m *DialRequest) GetTo() *Addr {
	if m != nil {
		return m.To
	}
	return nil
}

type Addr struct {
	Network              string   `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{2}
}
func (m *Addr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return m.Size()
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Addr) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Data struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{3}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return m.Size()
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*Frame)(nil), "api.Frame")
	proto.RegisterType((*DialRequest)(nil), "api.DialRequest")
	proto.RegisterType((*Addr)(nil), "api.Addr")
	proto.RegisterType((*Data)(nil), "api.Data")
}

func init() { proto.RegisterFile("tunnel.proto", fileDescriptor_6f51ddaa7891a711) }

var fileDescriptor_6f51ddaa7891a711 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0xa6, 0x5f, 0x66, 0x1a, 0x50, 0x16, 0x0f, 0xb1, 0x48, 0x94, 0xe0, 0xa1, 0xa7, 0x20,
	0x55, 0x10, 0xbc, 0x19, 0x83, 0xd4, 0x83, 0x1e, 0xd6, 0xfc, 0x00, 0x57, 0x3b, 0x48, 0xb0, 0x66,
	0xe3, 0xee, 0x56, 0xe9, 0x3f, 0xf4, 0xe8, 0x4f, 0x90, 0xfc, 0x12, 0xd9, 0x49, 0x5b, 0xd2, 0xdb,
	0xbc, 0x79, 0x6f, 0xdf, 0x9b, 0x7d, 0x10, 0xd8, 0x65, 0x59, 0xe2, 0x22, 0xa9, 0xb4, 0xb2, 0x8a,
	0x77, 0x65, 0x55, 0xc4, 0x05, 0xf4, 0xef, 0xb4, 0xfc, 0x40, 0x7e, 0x09, 0xa3, 0xac, 0x90, 0x0b,
	0x81, 0x9f, 0x4b, 0x34, 0x36, 0x64, 0xa7, 0x6c, 0x32, 0x9a, 0x1e, 0x24, 0xb2, 0x2a, 0x92, 0xd6,
	0x7e, 0xd6, 0x11, 0x6d, 0x19, 0x3f, 0x81, 0x5e, 0x26, 0xad, 0x0c, 0x3d, 0x92, 0xfb, 0x8d, 0x5c,
	0x5a, 0x39, 0xeb, 0x08, 0x22, 0x52, 0x1f, 0x86, 0x0f, 0x68, 0x8c, 0x7c, 0xc3, 0xf8, 0x79, 0x27,
	0x81, 0x8f, 0x61, 0xef, 0x09, 0xf5, 0x17, 0xea, 0xfb, 0x8c, 0xd2, 0x7c, 0xb1, 0xc5, 0x8e, 0xbb,
	0x55, 0x65, 0x99, 0xaf, 0x2a, 0x24, 0x6b, 0x5f, 0x6c, 0x31, 0x3f, 0x02, 0x2f, 0x57, 0x61, 0xb7,
	0x15, 0x78, 0x33, 0x9f, 0x6b, 0xe1, 0xe5, 0x2a, 0xbe, 0x86, 0x9e, 0x9b, 0x79, 0x08, 0xc3, 0x47,
	0xb4, 0xdf, 0x4a, 0xbf, 0xaf, 0x9d, 0x37, 0xd0, 0x31, 0x4e, 0x81, 0xc6, 0xac, 0x7d, 0x37, 0x30,
	0x3e, 0x6e, 0x7e, 0xc2, 0x0f, 0xa1, 0x9f, 0xae, 0x2c, 0x1a, 0x7a, 0x19, 0x88, 0x06, 0x4c, 0xaf,
	0x60, 0x3f, 0xa7, 0xee, 0x50, 0xbb, 0x23, 0x8b, 0x57, 0xe4, 0x67, 0x30, 0x68, 0x56, 0x1c, 0xe8,
	0x0a, 0xaa, 0x71, 0xdc, 0x9a, 0x27, 0xec, 0x9c, 0xa5, 0xc1, 0x4f, 0x1d, 0xb1, 0xdf, 0x3a, 0x62,
	0x7f, 0x75, 0xc4, 0x5e, 0x06, 0xd4, 0xfc, 0xc5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xd5,
	0x08, 0x80, 0x89, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TunnelerServiceClient is the client API for TunnelerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TunnelerServiceClient interface {
	Tunnel(ctx context.Context, opts ...grpc.CallOption) (TunnelerService_TunnelClient, error)
}

type tunnelerServiceClient struct {
	cc *grpc.ClientConn
}

func NewTunnelerServiceClient(cc *grpc.ClientConn) TunnelerServiceClient {
	return &tunnelerServiceClient{cc}
}

func (c *tunnelerServiceClient) Tunnel(ctx context.Context, opts ...grpc.CallOption) (TunnelerService_TunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TunnelerService_serviceDesc.Streams[0], "/api.TunnelerService/Tunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelerServiceTunnelClient{stream}
	return x, nil
}

type TunnelerService_TunnelClient interface {
	Send(*Frame) error
	Recv() (*Frame, error)
	grpc.ClientStream
}

type tunnelerServiceTunnelClient struct {
	grpc.ClientStream
}

func (x *tunnelerServiceTunnelClient) Send(m *Frame) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelerServiceTunnelClient) Recv() (*Frame, error) {
	m := new(Frame)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelerServiceServer is the server API for TunnelerService service.
type TunnelerServiceServer interface {
	Tunnel(TunnelerService_TunnelServer) error
}

// UnimplementedTunnelerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelerServiceServer struct {
}

func (*UnimplementedTunnelerServiceServer) Tunnel(srv TunnelerService_TunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method Tunnel not implemented")
}

func RegisterTunnelerServiceServer(s *grpc.Server, srv TunnelerServiceServer) {
	s.RegisterService(&_TunnelerService_serviceDesc, srv)
}

func _TunnelerService_Tunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelerServiceServer).Tunnel(&tunnelerServiceTunnelServer{stream})
}

type TunnelerService_TunnelServer interface {
	Send(*Frame) error
	Recv() (*Frame, error)
	grpc.ServerStream
}

type tunnelerServiceTunnelServer struct {
	grpc.ServerStream
}

func (x *tunnelerServiceTunnelServer) Send(m *Frame) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelerServiceTunnelServer) Recv() (*Frame, error) {
	m := new(Frame)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TunnelerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TunnelerService",
	HandlerType: (*TunnelerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tunnel",
			Handler:       _TunnelerService_Tunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tunnel.proto",
}

func (m *Frame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Frame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Frame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Message != nil {
		{
			size := m.Message.Size()
			i -= size
			if _, err := m.Message.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Frame_DialRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Frame_DialRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DialRequest != nil {
		{
			size, err := m.DialRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTunnel(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Frame_Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Frame_Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTunnel(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *DialRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DialRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DialRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.To != nil {
		{
			size, err := m.To.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTunnel(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnType) > 0 {
		i -= len(m.ConnType)
		copy(dAtA[i:], m.ConnType)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.ConnType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Addr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Addr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Addr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintTunnel(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTunnel(dAtA []byte, offset int, v uint64) int {
	offset -= sovTunnel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Frame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		n += m.Message.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Frame_DialRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DialRequest != nil {
		l = m.DialRequest.Size()
		n += 1 + l + sovTunnel(uint64(l))
	}
	return n
}
func (m *Frame_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTunnel(uint64(l))
	}
	return n
}
func (m *DialRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	l = len(m.ConnType)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.To != nil {
		l = m.To.Size()
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Addr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + sovTunnel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTunnel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTunnel(x uint64) (n int) {
	return sovTunnel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Frame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Frame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Frame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DialRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DialRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &Frame_DialRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Data{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &Frame_Data{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DialRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DialRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DialRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.To == nil {
				m.To = &Addr{}
			}
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Addr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Addr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Addr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTunnel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTunnel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTunnel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTunnel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTunnel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTunnel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTunnel
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTunnel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTunnel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTunnel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTunnel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTunnel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTunnel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTunnel = fmt.Errorf("proto: unexpected end of group")
)
