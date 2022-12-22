// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platine/epoch/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryGetEpochRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetEpochRequest) Reset()         { *m = QueryGetEpochRequest{} }
func (m *QueryGetEpochRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetEpochRequest) ProtoMessage()    {}
func (*QueryGetEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{0}
}
func (m *QueryGetEpochRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEpochRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEpochRequest.Merge(m, src)
}
func (m *QueryGetEpochRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEpochRequest proto.InternalMessageInfo

func (m *QueryGetEpochRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetEpochResponse struct {
	Epoch Epoch `protobuf:"bytes,1,opt,name=Epoch,proto3" json:"Epoch"`
}

func (m *QueryGetEpochResponse) Reset()         { *m = QueryGetEpochResponse{} }
func (m *QueryGetEpochResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetEpochResponse) ProtoMessage()    {}
func (*QueryGetEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b446b43f3e004603, []int{1}
}
func (m *QueryGetEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEpochResponse.Merge(m, src)
}
func (m *QueryGetEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEpochResponse proto.InternalMessageInfo

func (m *QueryGetEpochResponse) GetEpoch() Epoch {
	if m != nil {
		return m.Epoch
	}
	return Epoch{}
}

type QueryAllEpochRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
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

func (m *QueryAllEpochRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllEpochResponse struct {
	Epoch      []Epoch             `protobuf:"bytes,1,rep,name=Epoch,proto3" json:"Epoch"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
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

func (m *QueryAllEpochResponse) GetEpoch() []Epoch {
	if m != nil {
		return m.Epoch
	}
	return nil
}

func (m *QueryAllEpochResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetEpochRequest)(nil), "platinenetwork.platine.epoch.QueryGetEpochRequest")
	proto.RegisterType((*QueryGetEpochResponse)(nil), "platinenetwork.platine.epoch.QueryGetEpochResponse")
	proto.RegisterType((*QueryAllEpochRequest)(nil), "platinenetwork.platine.epoch.QueryAllEpochRequest")
	proto.RegisterType((*QueryAllEpochResponse)(nil), "platinenetwork.platine.epoch.QueryAllEpochResponse")
}

func init() { proto.RegisterFile("platine/epoch/query.proto", fileDescriptor_b446b43f3e004603) }

var fileDescriptor_b446b43f3e004603 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x2d, 0x2f, 0x19, 0x43, 0x83, 0x1d, 0x44, 0x06, 0x5b, 0x08, 0xde, 0xf0, 0x46, 0xf6,
	0x5f, 0x5a, 0x92, 0x0f, 0x30, 0x12, 0xd8, 0x02, 0x3b, 0x6d, 0x3e, 0x8d, 0x1d, 0x06, 0x72, 0x22,
	0x1c, 0x31, 0xc7, 0x72, 0x22, 0x65, 0x5b, 0x18, 0xbd, 0xf4, 0x13, 0x14, 0xfa, 0x09, 0x42, 0xbf,
	0x4c, 0x8e, 0x29, 0xbd, 0xf4, 0x54, 0x4a, 0xd2, 0x0f, 0x52, 0x2c, 0x09, 0x92, 0xb8, 0x69, 0x9a,
	0xf6, 0x62, 0x6c, 0xbd, 0xcf, 0xf3, 0xbe, 0xbf, 0xc7, 0xaf, 0x0d, 0x9f, 0xa6, 0x31, 0x55, 0x3c,
	0x61, 0x84, 0xa5, 0xa2, 0xd3, 0x23, 0x83, 0x11, 0x1b, 0x8e, 0x71, 0x3a, 0x14, 0x4a, 0xa0, 0x8a,
	0x2d, 0x25, 0x4c, 0xfd, 0x15, 0xc3, 0xdf, 0xd8, 0x3e, 0x62, 0xad, 0x2c, 0x97, 0x22, 0x11, 0x09,
	0x2d, 0x24, 0xd9, 0x9d, 0xf1, 0x94, 0x2b, 0x91, 0x10, 0x51, 0xcc, 0x08, 0x4d, 0x39, 0xa1, 0x49,
	0x22, 0x14, 0x55, 0x5c, 0x24, 0xd2, 0x56, 0xdf, 0x76, 0x84, 0xec, 0x0b, 0x49, 0x42, 0x2a, 0x99,
	0x19, 0x45, 0xfe, 0xd4, 0x42, 0xa6, 0x68, 0x8d, 0xa4, 0x34, 0xe2, 0x89, 0x16, 0x5b, 0x6d, 0x0e,
	0x4c, 0x5f, 0x4d, 0xc9, 0xaf, 0xc2, 0xd2, 0xf7, 0xcc, 0xdc, 0x66, 0xea, 0x73, 0x76, 0x1c, 0xb0,
	0xc1, 0x88, 0x49, 0x85, 0x1e, 0x41, 0x97, 0x77, 0x9f, 0x80, 0xe7, 0xe0, 0x75, 0x21, 0x70, 0x79,
	0xd7, 0xff, 0x01, 0x1f, 0xe7, 0x74, 0x32, 0x15, 0x89, 0x64, 0xe8, 0x13, 0x2c, 0xea, 0x03, 0xad,
	0x7d, 0x58, 0x7f, 0x81, 0xb7, 0x25, 0xc5, 0x5a, 0xda, 0x2a, 0x4c, 0xcf, 0x9e, 0x39, 0x81, 0xf1,
	0xf9, 0xbf, 0x2c, 0x41, 0x33, 0x8e, 0xd7, 0x08, 0xbe, 0x40, 0xb8, 0x0c, 0x62, 0xbb, 0x57, 0xb1,
	0x49, 0x8d, 0xb3, 0xd4, 0xd8, 0xbc, 0x60, 0x9b, 0x1a, 0x7f, 0xa3, 0x11, 0xb3, 0xde, 0x60, 0xc5,
	0xe9, 0x4f, 0x80, 0x45, 0x5f, 0x0e, 0xb8, 0x8a, 0x7e, 0xef, 0x2e, 0xe8, 0xa8, 0xbd, 0x86, 0xe8,
	0x6a, 0xc4, 0x57, 0x37, 0x22, 0x9a, 0xe9, 0xab, 0x8c, 0xf5, 0x63, 0x17, 0x16, 0x35, 0x23, 0x3a,
	0x02, 0x16, 0x0a, 0xd5, 0xb7, 0xe3, 0x6c, 0xda, 0x5a, 0xb9, 0x71, 0x2b, 0x8f, 0x01, 0xf1, 0x6b,
	0xfb, 0x27, 0x17, 0x87, 0xee, 0x3b, 0xf4, 0x86, 0x58, 0xf5, 0x07, 0xeb, 0x26, 0x1b, 0x3e, 0x1b,
	0xf2, 0x9f, 0x77, 0xf7, 0xd0, 0x04, 0xc0, 0x07, 0xba, 0x49, 0x33, 0x8e, 0x77, 0x02, 0xcd, 0x2d,
	0x77, 0x27, 0xd0, 0xfc, 0xbe, 0xfc, 0xf7, 0x1a, 0xb4, 0x8a, 0x5e, 0xee, 0x02, 0xda, 0xfa, 0x3a,
	0x9d, 0x7b, 0x60, 0x36, 0xf7, 0xc0, 0xf9, 0xdc, 0x03, 0x07, 0x0b, 0xcf, 0x99, 0x2d, 0x3c, 0xe7,
	0x74, 0xe1, 0x39, 0x3f, 0x3f, 0x46, 0x5c, 0xf5, 0x46, 0x21, 0xee, 0x88, 0xfe, 0xb5, 0x9d, 0xfe,
	0xd9, 0x5e, 0x6a, 0x9c, 0x32, 0x19, 0xde, 0xd7, 0x3f, 0x4b, 0xe3, 0x32, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0x9a, 0x57, 0x80, 0xe2, 0x03, 0x00, 0x00,
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
	// Queries a Epoch by id.
	Epoch(ctx context.Context, in *QueryGetEpochRequest, opts ...grpc.CallOption) (*QueryGetEpochResponse, error)
	// Queries a list of Epoch items.
	EpochAll(ctx context.Context, in *QueryAllEpochRequest, opts ...grpc.CallOption) (*QueryAllEpochResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Epoch(ctx context.Context, in *QueryGetEpochRequest, opts ...grpc.CallOption) (*QueryGetEpochResponse, error) {
	out := new(QueryGetEpochResponse)
	err := c.cc.Invoke(ctx, "/platinenetwork.platine.epoch.Query/Epoch", in, out, opts...)
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
	// Queries a Epoch by id.
	Epoch(context.Context, *QueryGetEpochRequest) (*QueryGetEpochResponse, error)
	// Queries a list of Epoch items.
	EpochAll(context.Context, *QueryAllEpochRequest) (*QueryAllEpochResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Epoch(ctx context.Context, req *QueryGetEpochRequest) (*QueryGetEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Epoch not implemented")
}
func (*UnimplementedQueryServer) EpochAll(ctx context.Context, req *QueryAllEpochRequest) (*QueryAllEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Epoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Epoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platinenetwork.platine.epoch.Query/Epoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Epoch(ctx, req.(*QueryGetEpochRequest))
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
			MethodName: "Epoch",
			Handler:    _Query_Epoch_Handler,
		},
		{
			MethodName: "EpochAll",
			Handler:    _Query_EpochAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platine/epoch/query.proto",
}

func (m *QueryGetEpochRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEpochRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEpochRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Epoch.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Epoch) > 0 {
		for iNdEx := len(m.Epoch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epoch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Epoch.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epoch) > 0 {
		for _, e := range m.Epoch {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetEpochRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetEpochRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEpochRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
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
func (m *QueryGetEpochResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
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
			if err := m.Epoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
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
			m.Epoch = append(m.Epoch, Epoch{})
			if err := m.Epoch[len(m.Epoch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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