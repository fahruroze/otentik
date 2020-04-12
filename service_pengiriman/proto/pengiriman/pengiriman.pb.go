// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pengiriman/pengiriman.proto

package pengiriman

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Pengiriman struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string       `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	DriverId             string       `protobuf:"bytes,5,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Pengiriman) Reset()         { *m = Pengiriman{} }
func (m *Pengiriman) String() string { return proto.CompactTextString(m) }
func (*Pengiriman) ProtoMessage()    {}
func (*Pengiriman) Descriptor() ([]byte, []int) {
	return fileDescriptor_5489ba00305b307b, []int{0}
}

func (m *Pengiriman) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pengiriman.Unmarshal(m, b)
}
func (m *Pengiriman) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pengiriman.Marshal(b, m, deterministic)
}
func (m *Pengiriman) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pengiriman.Merge(m, src)
}
func (m *Pengiriman) XXX_Size() int {
	return xxx_messageInfo_Pengiriman.Size(m)
}
func (m *Pengiriman) XXX_DiscardUnknown() {
	xxx_messageInfo_Pengiriman.DiscardUnknown(m)
}

var xxx_messageInfo_Pengiriman proto.InternalMessageInfo

func (m *Pengiriman) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pengiriman) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Pengiriman) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Pengiriman) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Pengiriman) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_5489ba00305b307b, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool        `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Pengiriman           *Pengiriman `protobuf:"bytes,2,opt,name=pengiriman,proto3" json:"pengiriman,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5489ba00305b307b, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetPengiriman() *Pengiriman {
	if m != nil {
		return m.Pengiriman
	}
	return nil
}

func init() {
	proto.RegisterType((*Pengiriman)(nil), "pengiriman.Pengiriman")
	proto.RegisterType((*Container)(nil), "pengiriman.Container")
	proto.RegisterType((*Response)(nil), "pengiriman.Response")
}

func init() {
	proto.RegisterFile("proto/pengiriman/pengiriman.proto", fileDescriptor_5489ba00305b307b)
}

var fileDescriptor_5489ba00305b307b = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x6d, 0x9a, 0x26, 0x13, 0x10, 0x1d, 0xb4, 0x2e, 0x7a, 0x30, 0xe6, 0x94, 0x53, 0x85,
	0x88, 0x9e, 0x85, 0x9e, 0x72, 0x93, 0x15, 0x6f, 0x82, 0xc4, 0xcc, 0x50, 0xf7, 0xd0, 0x6c, 0xd8,
	0x4d, 0xeb, 0xcf, 0xf8, 0xb1, 0xd2, 0x8d, 0x69, 0x17, 0xc4, 0xdb, 0xbc, 0x37, 0xf3, 0xe6, 0xbd,
	0x9d, 0x85, 0xdb, 0xce, 0xe8, 0x5e, 0xdf, 0x75, 0xdc, 0xae, 0x95, 0x51, 0x9b, 0xba, 0xf5, 0xca,
	0xa5, 0xeb, 0x21, 0x1c, 0x99, 0xfc, 0x3b, 0x00, 0x78, 0x3e, 0x40, 0x3c, 0x81, 0x89, 0x22, 0x11,
	0x64, 0x41, 0x91, 0xc8, 0x89, 0x22, 0x44, 0x08, 0x89, 0x6d, 0x23, 0x26, 0x8e, 0x71, 0x35, 0x2e,
	0x20, 0xfa, 0x62, 0xb5, 0xfe, 0xec, 0xc5, 0x34, 0x0b, 0x8a, 0x99, 0xfc, 0x45, 0xf8, 0x00, 0xd0,
	0xe8, 0xb6, 0xaf, 0x55, 0xcb, 0xc6, 0x8a, 0x30, 0x9b, 0x16, 0x69, 0x79, 0xb1, 0xf4, 0xdc, 0x57,
	0x63, 0x57, 0x7a, 0x83, 0x78, 0x0d, 0x09, 0x19, 0xb5, 0x63, 0xf3, 0xae, 0x48, 0xcc, 0x9c, 0x4f,
	0x3c, 0x10, 0x15, 0xe5, 0x1a, 0x92, 0x83, 0xea, 0x4f, 0xb8, 0x1b, 0x48, 0x9b, 0xad, 0xed, 0xf5,
	0x66, 0xd0, 0x0e, 0x19, 0x61, 0xa4, 0x2a, 0x42, 0x01, 0xf3, 0x9a, 0xc8, 0xb0, 0xb5, 0x2e, 0x6a,
	0x22, 0x47, 0x88, 0x97, 0x30, 0xdf, 0xda, 0x41, 0x16, 0xba, 0x4e, 0xb4, 0x87, 0x15, 0xe5, 0x6f,
	0x10, 0x4b, 0xb6, 0x9d, 0x6e, 0x2d, 0xef, 0xe5, 0x8d, 0xe1, 0xba, 0xe7, 0xc1, 0x34, 0x96, 0x23,
	0xc4, 0x47, 0xf0, 0x6e, 0xe8, 0x8c, 0xd3, 0x72, 0xe1, 0x3f, 0xf5, 0x78, 0x52, 0xe9, 0x4d, 0x96,
	0xaf, 0x70, 0xf6, 0xc2, 0x66, 0xa7, 0x1a, 0xf6, 0x6e, 0xfe, 0x04, 0xa7, 0x2b, 0xb7, 0xd7, 0xe3,
	0xfe, 0x59, 0x76, 0x75, 0xee, 0xf3, 0x63, 0xd0, 0x8f, 0xc8, 0xfd, 0xeb, 0xfd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x0c, 0xa5, 0xd9, 0xfc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServicePengirimanClient is the client API for ServicePengiriman service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicePengirimanClient interface {
	CreatePengiriman(ctx context.Context, in *Pengiriman, opts ...grpc.CallOption) (*Response, error)
}

type servicePengirimanClient struct {
	cc grpc.ClientConnInterface
}

func NewServicePengirimanClient(cc grpc.ClientConnInterface) ServicePengirimanClient {
	return &servicePengirimanClient{cc}
}

func (c *servicePengirimanClient) CreatePengiriman(ctx context.Context, in *Pengiriman, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pengiriman.ServicePengiriman/CreatePengiriman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicePengirimanServer is the server API for ServicePengiriman service.
type ServicePengirimanServer interface {
	CreatePengiriman(context.Context, *Pengiriman) (*Response, error)
}

// UnimplementedServicePengirimanServer can be embedded to have forward compatible implementations.
type UnimplementedServicePengirimanServer struct {
}

func (*UnimplementedServicePengirimanServer) CreatePengiriman(ctx context.Context, req *Pengiriman) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePengiriman not implemented")
}

func RegisterServicePengirimanServer(s *grpc.Server, srv ServicePengirimanServer) {
	s.RegisterService(&_ServicePengiriman_serviceDesc, srv)
}

func _ServicePengiriman_CreatePengiriman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pengiriman)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicePengirimanServer).CreatePengiriman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pengiriman.ServicePengiriman/CreatePengiriman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicePengirimanServer).CreatePengiriman(ctx, req.(*Pengiriman))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServicePengiriman_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pengiriman.ServicePengiriman",
	HandlerType: (*ServicePengirimanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePengiriman",
			Handler:    _ServicePengiriman_CreatePengiriman_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pengiriman/pengiriman.proto",
}