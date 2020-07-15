// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/annotation.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
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

// AddAnnotationRequest is a params to add new annotation.
type AddAnnotationRequest struct {
	// An annotation description. Required.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Tags are used to filter annotations.
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	// Used for annotate node.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Used for annotate services.
	ServiceNames         []string `protobuf:"bytes,4,rep,name=service_names,json=serviceNames,proto3" json:"service_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAnnotationRequest) Reset()         { *m = AddAnnotationRequest{} }
func (m *AddAnnotationRequest) String() string { return proto.CompactTextString(m) }
func (*AddAnnotationRequest) ProtoMessage()    {}
func (*AddAnnotationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea15ba3ce14c56e, []int{0}
}

func (m *AddAnnotationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAnnotationRequest.Unmarshal(m, b)
}
func (m *AddAnnotationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAnnotationRequest.Marshal(b, m, deterministic)
}
func (m *AddAnnotationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAnnotationRequest.Merge(m, src)
}
func (m *AddAnnotationRequest) XXX_Size() int {
	return xxx_messageInfo_AddAnnotationRequest.Size(m)
}
func (m *AddAnnotationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAnnotationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAnnotationRequest proto.InternalMessageInfo

func (m *AddAnnotationRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AddAnnotationRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AddAnnotationRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AddAnnotationRequest) GetServiceNames() []string {
	if m != nil {
		return m.ServiceNames
	}
	return nil
}

type AddAnnotationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAnnotationResponse) Reset()         { *m = AddAnnotationResponse{} }
func (m *AddAnnotationResponse) String() string { return proto.CompactTextString(m) }
func (*AddAnnotationResponse) ProtoMessage()    {}
func (*AddAnnotationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dea15ba3ce14c56e, []int{1}
}

func (m *AddAnnotationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAnnotationResponse.Unmarshal(m, b)
}
func (m *AddAnnotationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAnnotationResponse.Marshal(b, m, deterministic)
}
func (m *AddAnnotationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAnnotationResponse.Merge(m, src)
}
func (m *AddAnnotationResponse) XXX_Size() int {
	return xxx_messageInfo_AddAnnotationResponse.Size(m)
}
func (m *AddAnnotationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAnnotationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAnnotationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddAnnotationRequest)(nil), "management.AddAnnotationRequest")
	proto.RegisterType((*AddAnnotationResponse)(nil), "management.AddAnnotationResponse")
}

func init() { proto.RegisterFile("managementpb/annotation.proto", fileDescriptor_dea15ba3ce14c56e) }

var fileDescriptor_dea15ba3ce14c56e = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xd9, 0xb6, 0x14, 0x1b, 0x5a, 0x0f, 0x41, 0x71, 0x59, 0xff, 0xd5, 0xed, 0xa5, 0x8a,
	0xbb, 0x41, 0x05, 0x0f, 0xde, 0xda, 0x07, 0xf0, 0xb0, 0x27, 0xf1, 0x22, 0x69, 0x33, 0xc4, 0x60,
	0x93, 0x59, 0x37, 0x69, 0xeb, 0x4d, 0xf0, 0xe4, 0xc5, 0x93, 0x8f, 0xe6, 0x03, 0x08, 0xe2, 0x83,
	0x48, 0xd3, 0xd2, 0xad, 0x7f, 0xf0, 0x94, 0xe1, 0xfb, 0x7e, 0x93, 0x49, 0xbe, 0x21, 0xbb, 0x9a,
	0x1b, 0x2e, 0x41, 0x83, 0x71, 0xf9, 0x80, 0x71, 0x63, 0xd0, 0x71, 0xa7, 0xd0, 0xa4, 0x79, 0x81,
	0x0e, 0x29, 0x29, 0xed, 0xe8, 0x5c, 0x2a, 0x77, 0x3b, 0x1e, 0xa4, 0x43, 0xd4, 0x4c, 0x4f, 0x95,
	0xbb, 0xc3, 0x29, 0x93, 0x98, 0x78, 0x30, 0x99, 0xf0, 0x91, 0x12, 0xdc, 0x61, 0x61, 0xd9, 0xb2,
	0x9c, 0xdf, 0x11, 0xed, 0x48, 0x44, 0x39, 0x02, 0xc6, 0x73, 0xb5, 0x32, 0xc0, 0x2e, 0xdc, 0x63,
	0x7f, 0x0c, 0x13, 0x09, 0x26, 0xb1, 0x53, 0x2e, 0x25, 0x14, 0x0c, 0x73, 0x4f, 0xfc, 0xa6, 0xe3,
	0xe7, 0x80, 0x6c, 0xf4, 0x84, 0xe8, 0x2d, 0x8d, 0x0c, 0xee, 0xc7, 0x60, 0x1d, 0x8d, 0x48, 0xcd,
	0xc1, 0x83, 0x0b, 0x83, 0x76, 0xd0, 0x6d, 0xf4, 0xeb, 0x1f, 0xef, 0xfb, 0x95, 0xab, 0x20, 0xf3,
	0x1a, 0xa5, 0xa4, 0xe6, 0xb8, 0xb4, 0x61, 0xa5, 0x5d, 0xed, 0x36, 0x32, 0x5f, 0xd3, 0x6d, 0xd2,
	0x30, 0x28, 0xe0, 0xc6, 0x70, 0x0d, 0x61, 0x75, 0xd6, 0x94, 0xad, 0xcd, 0x84, 0x4b, 0xae, 0x81,
	0x76, 0x48, 0xcb, 0x42, 0x31, 0x51, 0xc3, 0xb9, 0x6f, 0xc3, 0x9a, 0xef, 0x6c, 0x2e, 0xc4, 0x19,
	0x63, 0xe3, 0x2d, 0xb2, 0xf9, 0xe3, 0x25, 0x36, 0x47, 0x63, 0xe1, 0xf4, 0x25, 0x20, 0xa4, 0x94,
	0xe9, 0x23, 0x69, 0x7d, 0xe3, 0x68, 0x3b, 0x2d, 0x43, 0x4d, 0xff, 0xfa, 0x4c, 0x74, 0xf0, 0x0f,
	0x31, 0x1f, 0x12, 0x1f, 0x3e, 0xbd, 0x7d, 0xbe, 0x56, 0x3a, 0xf1, 0x1e, 0x9b, 0x9c, 0xb0, 0x92,
	0x66, 0x25, 0x6a, 0x59, 0x4f, 0x88, 0x8b, 0xe0, 0xa8, 0xbf, 0x7e, 0xdd, 0x5c, 0x5d, 0xf2, 0xa0,
	0xee, 0xa3, 0x3c, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x77, 0x80, 0x45, 0x85, 0xfb, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnnotationClient is the client API for Annotation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnnotationClient interface {
	// AddAnnotation adds annotation.
	AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error)
}

type annotationClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnotationClient(cc grpc.ClientConnInterface) AnnotationClient {
	return &annotationClient{cc}
}

func (c *annotationClient) AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error) {
	out := new(AddAnnotationResponse)
	err := c.cc.Invoke(ctx, "/management.Annotation/AddAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnotationServer is the server API for Annotation service.
type AnnotationServer interface {
	// AddAnnotation adds annotation.
	AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error)
}

// UnimplementedAnnotationServer can be embedded to have forward compatible implementations.
type UnimplementedAnnotationServer struct {
}

func (*UnimplementedAnnotationServer) AddAnnotation(ctx context.Context, req *AddAnnotationRequest) (*AddAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnnotation not implemented")
}

func RegisterAnnotationServer(s *grpc.Server, srv AnnotationServer) {
	s.RegisterService(&_Annotation_serviceDesc, srv)
}

func _Annotation_AddAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnotationServer).AddAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Annotation/AddAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnotationServer).AddAnnotation(ctx, req.(*AddAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Annotation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Annotation",
	HandlerType: (*AnnotationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAnnotation",
			Handler:    _Annotation_AddAnnotation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/annotation.proto",
}