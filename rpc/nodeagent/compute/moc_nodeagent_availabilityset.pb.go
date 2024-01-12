// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_availabilityset.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type AvailabilitySetRequest struct {
	// avset field
	AvailabilitySets []*AvailabilitySet `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	// common field
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AvailabilitySetRequest) Reset()         { *m = AvailabilitySetRequest{} }
func (m *AvailabilitySetRequest) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetRequest) ProtoMessage()    {}
func (*AvailabilitySetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{0}
}

func (m *AvailabilitySetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetRequest.Unmarshal(m, b)
}
func (m *AvailabilitySetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetRequest.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetRequest.Merge(m, src)
}
func (m *AvailabilitySetRequest) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetRequest.Size(m)
}
func (m *AvailabilitySetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetRequest proto.InternalMessageInfo

func (m *AvailabilitySetRequest) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type AvailabilitySetResponse struct {
	// avset field
	AvailabilitySets []*AvailabilitySet `protobuf:"bytes,1,rep,name=AvailabilitySets,proto3" json:"AvailabilitySets,omitempty"`
	// common fields
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AvailabilitySetResponse) Reset()         { *m = AvailabilitySetResponse{} }
func (m *AvailabilitySetResponse) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySetResponse) ProtoMessage()    {}
func (*AvailabilitySetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{1}
}

func (m *AvailabilitySetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySetResponse.Unmarshal(m, b)
}
func (m *AvailabilitySetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySetResponse.Marshal(b, m, deterministic)
}
func (m *AvailabilitySetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySetResponse.Merge(m, src)
}
func (m *AvailabilitySetResponse) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySetResponse.Size(m)
}
func (m *AvailabilitySetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySetResponse proto.InternalMessageInfo

func (m *AvailabilitySetResponse) GetAvailabilitySets() []*AvailabilitySet {
	if m != nil {
		return m.AvailabilitySets
	}
	return nil
}

func (m *AvailabilitySetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AvailabilitySetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type AvailabilitySet struct {
	// identifier fields
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	LocationName string `protobuf:"bytes,3,opt,name=locationName,proto3" json:"locationName,omitempty"`
	GroupName    string `protobuf:"bytes,4,opt,name=groupName,proto3" json:"groupName,omitempty"`
	// common fields
	Status *common.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Tags   *common.Tags   `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	// avset fields
	PlatformFaultDomainCount uint32            `protobuf:"varint,7,opt,name=platformFaultDomainCount,proto3" json:"platformFaultDomainCount,omitempty"`
	VirtualMachines          []*VirtualMachine `protobuf:"bytes,8,rep,name=virtualMachines,proto3" json:"virtualMachines,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *AvailabilitySet) Reset()         { *m = AvailabilitySet{} }
func (m *AvailabilitySet) String() string { return proto.CompactTextString(m) }
func (*AvailabilitySet) ProtoMessage()    {}
func (*AvailabilitySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d615df2311cce965, []int{2}
}

func (m *AvailabilitySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilitySet.Unmarshal(m, b)
}
func (m *AvailabilitySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilitySet.Marshal(b, m, deterministic)
}
func (m *AvailabilitySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilitySet.Merge(m, src)
}
func (m *AvailabilitySet) XXX_Size() int {
	return xxx_messageInfo_AvailabilitySet.Size(m)
}
func (m *AvailabilitySet) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilitySet.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilitySet proto.InternalMessageInfo

func (m *AvailabilitySet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AvailabilitySet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AvailabilitySet) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *AvailabilitySet) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *AvailabilitySet) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AvailabilitySet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *AvailabilitySet) GetPlatformFaultDomainCount() uint32 {
	if m != nil {
		return m.PlatformFaultDomainCount
	}
	return 0
}

func (m *AvailabilitySet) GetVirtualMachines() []*VirtualMachine {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func init() {
	proto.RegisterType((*AvailabilitySetRequest)(nil), "moc.nodeagent.compute.AvailabilitySetRequest")
	proto.RegisterType((*AvailabilitySetResponse)(nil), "moc.nodeagent.compute.AvailabilitySetResponse")
	proto.RegisterType((*AvailabilitySet)(nil), "moc.nodeagent.compute.AvailabilitySet")
}

func init() {
	proto.RegisterFile("moc_nodeagent_availabilityset.proto", fileDescriptor_d615df2311cce965)
}

var fileDescriptor_d615df2311cce965 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x9a, 0x1a, 0xb2, 0xa1, 0x29, 0x5a, 0x15, 0x6a, 0x45, 0x80, 0x22, 0x57, 0xa0,
	0x5c, 0x58, 0x23, 0xc3, 0x89, 0x5b, 0xcb, 0x1f, 0x89, 0x03, 0x54, 0xda, 0x56, 0x3d, 0x70, 0xa9,
	0x36, 0xce, 0xc6, 0x5d, 0xb1, 0xbb, 0xb3, 0xec, 0x9f, 0xa0, 0x1e, 0x79, 0x15, 0x5e, 0x82, 0x57,
	0xe0, 0xb1, 0x50, 0xd6, 0x26, 0x10, 0x97, 0x4a, 0xbd, 0x70, 0x4a, 0x3c, 0xf3, 0x9b, 0xcf, 0xe3,
	0xf9, 0x3e, 0x74, 0xa0, 0xa0, 0x3a, 0xd7, 0x30, 0xe7, 0xac, 0xe6, 0xda, 0x9f, 0xb3, 0x25, 0x13,
	0x92, 0xcd, 0x84, 0x14, 0xfe, 0xd2, 0x71, 0x4f, 0x8c, 0x05, 0x0f, 0xf8, 0xbe, 0x82, 0x8a, 0xac,
	0x21, 0x52, 0x81, 0x32, 0xc1, 0xf3, 0xf1, 0xe3, 0x1a, 0xa0, 0x96, 0xbc, 0x88, 0xd0, 0x2c, 0x2c,
	0x8a, 0xaf, 0x96, 0x19, 0xc3, 0xad, 0x6b, 0xc6, 0xc6, 0xf9, 0xa6, 0xf6, 0x52, 0x58, 0x1f, 0x98,
	0x54, 0xac, 0xba, 0x10, 0x9a, 0xb7, 0xcc, 0xfe, 0x8a, 0xa9, 0x40, 0x29, 0xd0, 0xed, 0x4f, 0xd3,
	0xc8, 0xbf, 0x27, 0xe8, 0xc1, 0xe1, 0x5f, 0xdb, 0x9c, 0x70, 0x4f, 0xf9, 0x97, 0xc0, 0x9d, 0xc7,
	0x14, 0xdd, 0xeb, 0x74, 0x5c, 0x96, 0x4c, 0xb6, 0xa6, 0xc3, 0xf2, 0x29, 0xf9, 0xe7, 0xa6, 0xa4,
	0x2b, 0x74, 0x65, 0x1e, 0xbf, 0x44, 0x3b, 0xc7, 0x86, 0x5b, 0xe6, 0x05, 0xe8, 0xd3, 0x4b, 0xc3,
	0xb3, 0xde, 0x24, 0x99, 0x8e, 0xca, 0x51, 0x14, 0x5c, 0x77, 0xe8, 0x26, 0x94, 0xff, 0x48, 0xd0,
	0xfe, 0x95, 0x25, 0x9d, 0x01, 0xed, 0xf8, 0x7f, 0xd9, 0xb2, 0x44, 0x29, 0xe5, 0x2e, 0x48, 0x1f,
	0xd7, 0x1b, 0x96, 0x63, 0xd2, 0x58, 0x40, 0x7e, 0x5b, 0x40, 0x8e, 0x00, 0xe4, 0x19, 0x93, 0x81,
	0xd3, 0x96, 0xc4, 0x7b, 0x68, 0xfb, 0xad, 0xb5, 0x60, 0xb3, 0xad, 0x49, 0x32, 0x1d, 0xd0, 0xe6,
	0x21, 0xff, 0xd9, 0x43, 0xbb, 0x1d, 0x79, 0x8c, 0x51, 0x5f, 0x33, 0xc5, 0xb3, 0x24, 0x82, 0xf1,
	0x3f, 0x1e, 0xa1, 0x9e, 0x98, 0xc7, 0xb7, 0x0d, 0x68, 0x4f, 0xcc, 0x71, 0x8e, 0xee, 0x4a, 0xa8,
	0xe2, 0x05, 0x3e, 0xae, 0xd8, 0x46, 0x74, 0xa3, 0x86, 0x1f, 0xa2, 0x41, 0x6d, 0x21, 0x98, 0x08,
	0xf4, 0x23, 0xf0, 0xa7, 0x80, 0x0f, 0x50, 0xea, 0x3c, 0xf3, 0xc1, 0x65, 0xdb, 0xf1, 0x1b, 0x86,
	0xf1, 0x1a, 0x27, 0xb1, 0x44, 0xdb, 0x16, 0x7e, 0x84, 0xfa, 0x9e, 0xd5, 0x2e, 0x4b, 0x23, 0x32,
	0x88, 0xc8, 0x29, 0xab, 0x1d, 0x8d, 0x65, 0xfc, 0x0a, 0x65, 0x46, 0x32, 0xbf, 0x00, 0xab, 0xde,
	0xb1, 0x20, 0xfd, 0x1b, 0x50, 0x4c, 0xe8, 0xd7, 0x10, 0xb4, 0xcf, 0x6e, 0x4f, 0x92, 0xe9, 0x0e,
	0xbd, 0xb6, 0x8f, 0x8f, 0xd1, 0x6e, 0x9b, 0xc4, 0x0f, 0x4d, 0x12, 0x5d, 0x76, 0x27, 0xda, 0xf2,
	0xe4, 0x1a, 0x5b, 0xce, 0x36, 0x68, 0xda, 0x9d, 0x2e, 0xbf, 0x25, 0x68, 0xaf, 0x73, 0xca, 0xc3,
	0x95, 0x00, 0x16, 0x28, 0x7d, 0xaf, 0x97, 0xf0, 0x99, 0xe3, 0x67, 0x37, 0x74, 0xbc, 0x09, 0xf8,
	0x98, 0xdc, 0x14, 0x6f, 0xa2, 0x96, 0xdf, 0x3a, 0x7a, 0xfe, 0x89, 0xd4, 0xc2, 0x5f, 0x84, 0xd9,
	0x8a, 0x2d, 0x94, 0xa8, 0x2c, 0x38, 0x58, 0xf8, 0x42, 0x41, 0x55, 0x58, 0x53, 0x15, 0x6b, 0xad,
	0xa2, 0xd5, 0x9a, 0xa5, 0x31, 0x32, 0x2f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xc9, 0x80,
	0x2a, 0x01, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AvailabilitySetAgentClient is the client API for AvailabilitySetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AvailabilitySetAgentClient interface {
	Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error)
}

type availabilitySetAgentClient struct {
	cc *grpc.ClientConn
}

func NewAvailabilitySetAgentClient(cc *grpc.ClientConn) AvailabilitySetAgentClient {
	return &availabilitySetAgentClient{cc}
}

func (c *availabilitySetAgentClient) Invoke(ctx context.Context, in *AvailabilitySetRequest, opts ...grpc.CallOption) (*AvailabilitySetResponse, error) {
	out := new(AvailabilitySetResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvailabilitySetAgentServer is the server API for AvailabilitySetAgent service.
type AvailabilitySetAgentServer interface {
	Invoke(context.Context, *AvailabilitySetRequest) (*AvailabilitySetResponse, error)
}

// UnimplementedAvailabilitySetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAvailabilitySetAgentServer struct {
}

func (*UnimplementedAvailabilitySetAgentServer) Invoke(ctx context.Context, req *AvailabilitySetRequest) (*AvailabilitySetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterAvailabilitySetAgentServer(s *grpc.Server, srv AvailabilitySetAgentServer) {
	s.RegisterService(&_AvailabilitySetAgent_serviceDesc, srv)
}

func _AvailabilitySetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailabilitySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.compute.AvailabilitySetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailabilitySetAgentServer).Invoke(ctx, req.(*AvailabilitySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AvailabilitySetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.compute.AvailabilitySetAgent",
	HandlerType: (*AvailabilitySetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _AvailabilitySetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_availabilityset.proto",
}
