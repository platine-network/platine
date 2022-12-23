// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platine/epoch/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryCurrentEpochRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (m *QueryCurrentEpochRequest) Reset()         { *m = QueryCurrentEpochRequest{} }
func (m *QueryCurrentEpochRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochRequest) ProtoMessage()    {}
func (*QueryCurrentEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{0}
}
func (m *QueryCurrentEpochRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochRequest.Merge(m, src)
}
func (m *QueryCurrentEpochRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochRequest proto.InternalMessageInfo

func (m *QueryCurrentEpochRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type QueryCurrentEpochResponse struct {
	CurrentEpoch int64 `protobuf:"varint,1,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
}

func (m *QueryCurrentEpochResponse) Reset()         { *m = QueryCurrentEpochResponse{} }
func (m *QueryCurrentEpochResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochResponse) ProtoMessage()    {}
func (*QueryCurrentEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{1}
}
func (m *QueryCurrentEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochResponse.Merge(m, src)
}
func (m *QueryCurrentEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochResponse proto.InternalMessageInfo

func (m *QueryCurrentEpochResponse) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

type QueryAllEpochRequest struct {
}

func (m *QueryAllEpochRequest) Reset()         { *m = QueryAllEpochRequest{} }
func (m *QueryAllEpochRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllEpochRequest) ProtoMessage()    {}
func (*QueryAllEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{2}
}
func (m *QueryAllEpochRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllEpochRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllEpochRequest.Merge(m, src)
}
func (m *QueryAllEpochRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllEpochRequest proto.InternalMessageInfo

type QueryAllEpochResponse struct {
	Epochs []Epoch `protobuf:"bytes,1,rep,name=Epochs,proto3" json:"Epochs"`
}

func (m *QueryAllEpochResponse) Reset()         { *m = QueryAllEpochResponse{} }
func (m *QueryAllEpochResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllEpochResponse) ProtoMessage()    {}
func (*QueryAllEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{3}
}
func (m *QueryAllEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllEpochResponse.Merge(m, src)
}
func (m *QueryAllEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllEpochResponse proto.InternalMessageInfo

func (m *QueryAllEpochResponse) GetEpochs() []Epoch {
	if m != nil {
		return m.Epochs
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCurrentEpochRequest)(nil), "platinenetwork.platine.epoch.QueryCurrentEpochRequest")
	proto.RegisterType((*QueryCurrentEpochResponse)(nil), "platinenetwork.platine.epoch.QueryCurrentEpochResponse")
	proto.RegisterType((*QueryAllEpochRequest)(nil), "platinenetwork.platine.epoch.QueryAllEpochRequest")
	proto.RegisterType((*QueryAllEpochResponse)(nil), "platinenetwork.platine.epoch.QueryAllEpochResponse")
}

func init() { proto.RegisterFile("platine/epoch/query.proto", fileDescriptor_b446b43f3e004603) }

var fileDescriptor_b446b43f3e004603 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xc8, 0x49, 0x2c,
	0xc9, 0xcc, 0x4b, 0xd5, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0xd0, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0x4a, 0xe5, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65,
	0xeb, 0x41, 0xb9, 0x7a, 0x60, 0x95, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x85, 0xfa, 0x20,
	0x16, 0x44, 0x8f, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e,
	0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x54, 0x16, 0xcd, 0x32, 0x30,
	0x09, 0x91, 0x52, 0xb2, 0xe2, 0x92, 0x08, 0x04, 0xd9, 0xed, 0x5c, 0x5a, 0x54, 0x94, 0x9a, 0x57,
	0xe2, 0x0a, 0x92, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0xca, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcb, 0x4c, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42,
	0x12, 0x51, 0x72, 0xe0, 0x92, 0xc4, 0xa2, 0xb7, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x99,
	0x8b, 0x37, 0x19, 0x22, 0x1e, 0x0f, 0xb6, 0x0f, 0xac, 0x9f, 0x39, 0x88, 0x27, 0x19, 0x49, 0xb1,
	0x92, 0x18, 0x97, 0x08, 0xd8, 0x04, 0xc7, 0x9c, 0x1c, 0x64, 0x9b, 0x95, 0xa2, 0xb8, 0x44, 0xd1,
	0xc4, 0xa1, 0xa6, 0x3a, 0x72, 0xb1, 0x81, 0x05, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d,
	0x94, 0xf5, 0xf0, 0x05, 0x96, 0x1e, 0x58, 0xad, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x50,
	0x8d, 0x46, 0x47, 0x98, 0xb8, 0x58, 0xc1, 0x86, 0x0b, 0x2d, 0x65, 0xe4, 0xe2, 0x41, 0x76, 0xbb,
	0x90, 0x19, 0x7e, 0xd3, 0x70, 0x05, 0x94, 0x94, 0x39, 0xc9, 0xfa, 0x20, 0xde, 0x51, 0x52, 0x69,
	0xba, 0xfc, 0x64, 0x32, 0x93, 0x9c, 0x90, 0x8c, 0x3e, 0x6a, 0x0c, 0xa1, 0x84, 0x9c, 0xd0, 0x24,
	0x46, 0x2e, 0x0e, 0xb0, 0x3e, 0xc7, 0x9c, 0x1c, 0x21, 0x23, 0x22, 0xec, 0x42, 0x0b, 0x4e, 0x29,
	0x63, 0x92, 0xf4, 0x40, 0xdd, 0x26, 0x0b, 0x76, 0x9b, 0xb8, 0x90, 0xa8, 0x3e, 0x96, 0xd4, 0x53,
	0xec, 0xe4, 0x75, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x06, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0x30, 0xad, 0xba, 0x50, 0x8b, 0xe1, 0x46, 0x55,
	0x40, 0x0d, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xa7, 0x45, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0xb1, 0x0d, 0x2e, 0x15, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries current Epoch.
	CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error)
	// Queries a list of Epoch items.
	EpochAll(ctx context.Context, in *QueryAllEpochRequest, opts ...grpc.CallOption) (*QueryAllEpochResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error) {
	out := new(QueryCurrentEpochResponse)
	err := c.cc.Invoke(ctx, "/platinenetwork.platine.epoch.Query/CurrentEpoch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EpochAll(ctx context.Context, in *QueryAllEpochRequest, opts ...grpc.CallOption) (*QueryAllEpochResponse, error) {
	out := new(QueryAllEpochResponse)
	err := c.cc.Invoke(ctx, "/platinenetwork.platine.epoch.Query/EpochAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries current Epoch.
	CurrentEpoch(context.Context, *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error)
	// Queries a list of Epoch items.
	EpochAll(context.Context, *QueryAllEpochRequest) (*QueryAllEpochResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CurrentEpoch(ctx context.Context, req *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentEpoch not implemented")
}
func (*UnimplementedQueryServer) EpochAll(ctx context.Context, req *QueryAllEpochRequest) (*QueryAllEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CurrentEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platinenetwork.platine.epoch.Query/CurrentEpoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentEpoch(ctx, req.(*QueryCurrentEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EpochAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platinenetwork.platine.epoch.Query/EpochAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochAll(ctx, req.(*QueryAllEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platinenetwork.platine.epoch.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentEpoch",
			Handler:    _Query_CurrentEpoch_Handler,
		},
		{
			MethodName: "EpochAll",
			Handler:    _Query_EpochAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platine/epoch/query.proto",
}

func (m *QueryCurrentEpochRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrentEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllEpochRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllEpochRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllEpochRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for iNdEx := len(m.Epochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCurrentEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrentEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		n += 1 + sovQuery(uint64(m.CurrentEpoch))
	}
	return n
}

func (m *QueryAllEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for _, e := range m.Epochs {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCurrentEpochRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentEpochRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCurrentEpochResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllEpochRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllEpochRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllEpochRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllEpochResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epochs = append(m.Epochs, Epoch{})
			if err := m.Epochs[len(m.Epochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
