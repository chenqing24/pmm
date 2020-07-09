// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/metrics_names.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MetricsNamesRequest is emty.
type MetricsNamesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsNamesRequest) Reset()         { *m = MetricsNamesRequest{} }
func (m *MetricsNamesRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsNamesRequest) ProtoMessage()    {}
func (*MetricsNamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed854621c154122c, []int{0}
}

func (m *MetricsNamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsNamesRequest.Unmarshal(m, b)
}
func (m *MetricsNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsNamesRequest.Marshal(b, m, deterministic)
}
func (m *MetricsNamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsNamesRequest.Merge(m, src)
}
func (m *MetricsNamesRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsNamesRequest.Size(m)
}
func (m *MetricsNamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsNamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsNamesRequest proto.InternalMessageInfo

// MetricsNamesReply is map of stored metrics:
// key is root of metric name in db (Ex:. [m_]query_time[_sum]);
// value - Human readable name of metrics.
type MetricsNamesReply struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricsNamesReply) Reset()         { *m = MetricsNamesReply{} }
func (m *MetricsNamesReply) String() string { return proto.CompactTextString(m) }
func (*MetricsNamesReply) ProtoMessage()    {}
func (*MetricsNamesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed854621c154122c, []int{1}
}

func (m *MetricsNamesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsNamesReply.Unmarshal(m, b)
}
func (m *MetricsNamesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsNamesReply.Marshal(b, m, deterministic)
}
func (m *MetricsNamesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsNamesReply.Merge(m, src)
}
func (m *MetricsNamesReply) XXX_Size() int {
	return xxx_messageInfo_MetricsNamesReply.Size(m)
}
func (m *MetricsNamesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsNamesReply.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsNamesReply proto.InternalMessageInfo

func (m *MetricsNamesReply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsNamesRequest)(nil), "qan.v1beta1.MetricsNamesRequest")
	proto.RegisterType((*MetricsNamesReply)(nil), "qan.v1beta1.MetricsNamesReply")
	proto.RegisterMapType((map[string]string)(nil), "qan.v1beta1.MetricsNamesReply.DataEntry")
}

func init() {
	proto.RegisterFile("qanpb/metrics_names.proto", fileDescriptor_ed854621c154122c)
}

var fileDescriptor_ed854621c154122c = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4c, 0xcc, 0x2b,
	0x48, 0xd2, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0x8e, 0xcf, 0x4b, 0xcc, 0x4d, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x4c, 0xcc, 0xd3, 0x2b, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x2a, 0x95, 0xd2, 0x01, 0x53, 0xc9,
	0xba, 0xe9, 0xa9, 0x79, 0xba, 0xc5, 0xe5, 0x89, 0xe9, 0xe9, 0xa9, 0x45, 0xfa, 0xf9, 0x05, 0x60,
	0x15, 0x98, 0xaa, 0x95, 0x44, 0xb9, 0x84, 0x7d, 0x21, 0xf6, 0xf9, 0x81, 0xac, 0x0b, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xea, 0x62, 0xe4, 0x12, 0x44, 0x15, 0x2f, 0xc8, 0xa9, 0x14, 0xb2,
	0xe1, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xd0, 0x43,
	0x72, 0x94, 0x1e, 0x86, 0x6a, 0x3d, 0x97, 0xc4, 0x92, 0x44, 0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20,
	0xb0, 0x2e, 0x29, 0x73, 0x2e, 0x4e, 0xb8, 0x90, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a,
	0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x8d, 0xda, 0x19, 0xb9, 0x78, 0x90,
	0x8d, 0x17, 0x2a, 0xe7, 0xe2, 0x77, 0x4f, 0x2d, 0x41, 0x11, 0x52, 0xc0, 0xe3, 0x18, 0xb0, 0x97,
	0xa4, 0xe4, 0xf0, 0x3b, 0x57, 0x49, 0xa9, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x32, 0x4a, 0xe2, 0xfa,
	0x65, 0x06, 0xfa, 0x85, 0x89, 0x79, 0xfa, 0x68, 0x56, 0x58, 0x31, 0x6a, 0x39, 0xb1, 0x47, 0xb1,
	0x82, 0xe3, 0x28, 0x89, 0x0d, 0x1c, 0x7a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xce,
	0x83, 0x16, 0xb3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricsNamesClient is the client API for MetricsNames service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsNamesClient interface {
	// GetMetricsNames gets map of metrics names.
	GetMetricsNames(ctx context.Context, in *MetricsNamesRequest, opts ...grpc.CallOption) (*MetricsNamesReply, error)
}

type metricsNamesClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsNamesClient(cc grpc.ClientConnInterface) MetricsNamesClient {
	return &metricsNamesClient{cc}
}

func (c *metricsNamesClient) GetMetricsNames(ctx context.Context, in *MetricsNamesRequest, opts ...grpc.CallOption) (*MetricsNamesReply, error) {
	out := new(MetricsNamesReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.MetricsNames/GetMetricsNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsNamesServer is the server API for MetricsNames service.
type MetricsNamesServer interface {
	// GetMetricsNames gets map of metrics names.
	GetMetricsNames(context.Context, *MetricsNamesRequest) (*MetricsNamesReply, error)
}

// UnimplementedMetricsNamesServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsNamesServer struct {
}

func (*UnimplementedMetricsNamesServer) GetMetricsNames(ctx context.Context, req *MetricsNamesRequest) (*MetricsNamesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsNames not implemented")
}

func RegisterMetricsNamesServer(s *grpc.Server, srv MetricsNamesServer) {
	s.RegisterService(&_MetricsNames_serviceDesc, srv)
}

func _MetricsNames_GetMetricsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsNamesServer).GetMetricsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.MetricsNames/GetMetricsNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsNamesServer).GetMetricsNames(ctx, req.(*MetricsNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsNames_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.MetricsNames",
	HandlerType: (*MetricsNamesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricsNames",
			Handler:    _MetricsNames_GetMetricsNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/metrics_names.proto",
}
