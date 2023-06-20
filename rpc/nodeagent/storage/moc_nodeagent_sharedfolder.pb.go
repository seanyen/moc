// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_sharedfolder.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SharedFolderRequest struct {
	SharedFolderSystems  []*SharedFolder  `protobuf:"bytes,1,rep,name=SharedFolderSystems,proto3" json:"SharedFolderSystems,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SharedFolderRequest) Reset()         { *m = SharedFolderRequest{} }
func (m *SharedFolderRequest) String() string { return proto.CompactTextString(m) }
func (*SharedFolderRequest) ProtoMessage()    {}
func (*SharedFolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db807b23661f235a, []int{0}
}

func (m *SharedFolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedFolderRequest.Unmarshal(m, b)
}
func (m *SharedFolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedFolderRequest.Marshal(b, m, deterministic)
}
func (m *SharedFolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedFolderRequest.Merge(m, src)
}
func (m *SharedFolderRequest) XXX_Size() int {
	return xxx_messageInfo_SharedFolderRequest.Size(m)
}
func (m *SharedFolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedFolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SharedFolderRequest proto.InternalMessageInfo

func (m *SharedFolderRequest) GetSharedFolderSystems() []*SharedFolder {
	if m != nil {
		return m.SharedFolderSystems
	}
	return nil
}

func (m *SharedFolderRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type SharedFolderResponse struct {
	SharedFolderSystems  []*SharedFolder     `protobuf:"bytes,1,rep,name=SharedFolderSystems,proto3" json:"SharedFolderSystems,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SharedFolderResponse) Reset()         { *m = SharedFolderResponse{} }
func (m *SharedFolderResponse) String() string { return proto.CompactTextString(m) }
func (*SharedFolderResponse) ProtoMessage()    {}
func (*SharedFolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db807b23661f235a, []int{1}
}

func (m *SharedFolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedFolderResponse.Unmarshal(m, b)
}
func (m *SharedFolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedFolderResponse.Marshal(b, m, deterministic)
}
func (m *SharedFolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedFolderResponse.Merge(m, src)
}
func (m *SharedFolderResponse) XXX_Size() int {
	return xxx_messageInfo_SharedFolderResponse.Size(m)
}
func (m *SharedFolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedFolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SharedFolderResponse proto.InternalMessageInfo

func (m *SharedFolderResponse) GetSharedFolderSystems() []*SharedFolder {
	if m != nil {
		return m.SharedFolderSystems
	}
	return nil
}

func (m *SharedFolderResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SharedFolderResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SharedFolder struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ContainerName        string         `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	FolderName           string         `protobuf:"bytes,4,opt,name=folderName,proto3" json:"folderName,omitempty"`
	ReadOnly             bool           `protobuf:"varint,5,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	Path                 string         `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	VirtualmachineName   string         `protobuf:"bytes,7,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	GuestMountPath       string         `protobuf:"bytes,8,opt,name=guestMountPath,proto3" json:"guestMountPath,omitempty"`
	MountTag             string         `protobuf:"bytes,9,opt,name=mountTag,proto3" json:"mountTag,omitempty"`
	Status               *common.Status `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,11,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SharedFolder) Reset()         { *m = SharedFolder{} }
func (m *SharedFolder) String() string { return proto.CompactTextString(m) }
func (*SharedFolder) ProtoMessage()    {}
func (*SharedFolder) Descriptor() ([]byte, []int) {
	return fileDescriptor_db807b23661f235a, []int{2}
}

func (m *SharedFolder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedFolder.Unmarshal(m, b)
}
func (m *SharedFolder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedFolder.Marshal(b, m, deterministic)
}
func (m *SharedFolder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedFolder.Merge(m, src)
}
func (m *SharedFolder) XXX_Size() int {
	return xxx_messageInfo_SharedFolder.Size(m)
}
func (m *SharedFolder) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedFolder.DiscardUnknown(m)
}

var xxx_messageInfo_SharedFolder proto.InternalMessageInfo

func (m *SharedFolder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SharedFolder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SharedFolder) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *SharedFolder) GetFolderName() string {
	if m != nil {
		return m.FolderName
	}
	return ""
}

func (m *SharedFolder) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *SharedFolder) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SharedFolder) GetVirtualmachineName() string {
	if m != nil {
		return m.VirtualmachineName
	}
	return ""
}

func (m *SharedFolder) GetGuestMountPath() string {
	if m != nil {
		return m.GuestMountPath
	}
	return ""
}

func (m *SharedFolder) GetMountTag() string {
	if m != nil {
		return m.MountTag
	}
	return ""
}

func (m *SharedFolder) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SharedFolder) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *SharedFolder) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedFolderRequest)(nil), "moc.nodeagent.storage.SharedFolderRequest")
	proto.RegisterType((*SharedFolderResponse)(nil), "moc.nodeagent.storage.SharedFolderResponse")
	proto.RegisterType((*SharedFolder)(nil), "moc.nodeagent.storage.SharedFolder")
}

func init() { proto.RegisterFile("moc_nodeagent_sharedfolder.proto", fileDescriptor_db807b23661f235a) }

var fileDescriptor_db807b23661f235a = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x18, 0x85, 0xc9, 0xd6, 0x95, 0xd5, 0xdd, 0x2a, 0xcd, 0x0c, 0x30, 0x41, 0x9b, 0xaa, 0x0e, 0xa1,
	0x0a, 0xa4, 0x04, 0x15, 0x5e, 0x80, 0xa1, 0x21, 0xb8, 0x60, 0x43, 0x59, 0xe1, 0x82, 0x9b, 0xc9,
	0x75, 0xdc, 0x34, 0x5a, 0xec, 0x3f, 0xd8, 0xce, 0x50, 0xdf, 0x00, 0x89, 0xb7, 0xe0, 0x1d, 0x10,
	0xf7, 0x3c, 0x19, 0xca, 0x9f, 0x10, 0xd2, 0xd1, 0x8b, 0xdd, 0x70, 0xd5, 0xfa, 0x9c, 0xcf, 0x27,
	0x27, 0xce, 0x6f, 0x32, 0x54, 0x20, 0x2e, 0x34, 0xc4, 0x92, 0x27, 0x52, 0xbb, 0x0b, 0xbb, 0xe0,
	0x46, 0xc6, 0x73, 0xc8, 0x62, 0x69, 0x82, 0xdc, 0x80, 0x03, 0x7a, 0x57, 0x81, 0x08, 0x1a, 0x22,
	0xb0, 0x0e, 0x0c, 0x4f, 0xa4, 0x7f, 0x98, 0x00, 0x24, 0x99, 0x0c, 0x11, 0x9a, 0x15, 0xf3, 0xf0,
	0x8b, 0xe1, 0x79, 0x2e, 0x8d, 0xad, 0xb6, 0xf9, 0xf7, 0xcb, 0x60, 0x01, 0x4a, 0x81, 0xae, 0x7f,
	0x6a, 0xe3, 0xa0, 0x65, 0x68, 0x70, 0xe9, 0x3c, 0x15, 0xdc, 0xa5, 0x8d, 0xfd, 0xf0, 0x7a, 0xae,
	0x54, 0xb9, 0x5b, 0x56, 0xe6, 0xe8, 0xbb, 0x47, 0xee, 0x9c, 0x63, 0xc5, 0xd7, 0x58, 0x31, 0x92,
	0x9f, 0x0b, 0x69, 0x1d, 0xfd, 0xb0, 0x2a, 0x9f, 0x2f, 0xad, 0x93, 0xca, 0x32, 0x6f, 0xb8, 0x39,
	0xee, 0x4f, 0x8e, 0x82, 0xb5, 0x6f, 0x10, 0xac, 0x04, 0xad, 0xdb, 0x4f, 0x5f, 0x90, 0xdd, 0xb3,
	0x5c, 0x1a, 0xac, 0x37, 0x5d, 0xe6, 0x92, 0x6d, 0x0c, 0xbd, 0xf1, 0x60, 0x32, 0xc0, 0xc0, 0xc6,
	0x89, 0x56, 0xa1, 0xd1, 0x4f, 0x8f, 0xec, 0xaf, 0x96, 0xb4, 0x39, 0x68, 0x2b, 0xff, 0x57, 0xcb,
	0x09, 0xe9, 0x46, 0xd2, 0x16, 0x99, 0xc3, 0x7a, 0xfd, 0x89, 0x1f, 0x54, 0x47, 0x18, 0xfc, 0x39,
	0xc2, 0xe0, 0x18, 0x20, 0xfb, 0xc8, 0xb3, 0x42, 0x46, 0x35, 0x49, 0xf7, 0xc9, 0xd6, 0x89, 0x31,
	0x60, 0xd8, 0xe6, 0xd0, 0x1b, 0xf7, 0xa2, 0x6a, 0x31, 0xfa, 0xb6, 0x49, 0x76, 0xda, 0x4f, 0xa0,
	0x94, 0x74, 0x34, 0x57, 0x92, 0x79, 0x48, 0xe1, 0x7f, 0x3a, 0x20, 0x1b, 0x69, 0x8c, 0x8f, 0xea,
	0x45, 0x1b, 0x69, 0x4c, 0x1f, 0x91, 0x5d, 0x01, 0xda, 0xf1, 0x54, 0x4b, 0x73, 0x5a, 0xc2, 0x55,
	0xe4, 0xaa, 0x48, 0x0f, 0x09, 0xa9, 0xa6, 0x0a, 0x91, 0x0e, 0x22, 0x2d, 0x85, 0xfa, 0x64, 0xdb,
	0x48, 0x1e, 0x9f, 0xe9, 0x6c, 0xc9, 0xb6, 0x86, 0xde, 0x78, 0x3b, 0x6a, 0xd6, 0x94, 0x91, 0x4e,
	0xce, 0xdd, 0x82, 0x75, 0xcb, 0x5d, 0xc7, 0x9d, 0xaf, 0x3f, 0x98, 0x17, 0xa1, 0x42, 0x03, 0x42,
	0xaf, 0x52, 0xe3, 0x0a, 0x9e, 0x29, 0x2e, 0x16, 0xa9, 0x96, 0x98, 0x7e, 0x1b, 0xd3, 0xd7, 0x38,
	0xf4, 0x31, 0x19, 0x24, 0xe5, 0xc0, 0xbc, 0x83, 0x42, 0xbb, 0xf7, 0x65, 0xe6, 0x36, 0xb2, 0xd7,
	0xd4, 0xb2, 0x8d, 0x2a, 0x17, 0x53, 0x9e, 0xb0, 0x1e, 0x12, 0xcd, 0x9a, 0x1e, 0x91, 0xae, 0x75,
	0xdc, 0x15, 0x96, 0x11, 0x3c, 0xee, 0x3e, 0x7e, 0xb8, 0x73, 0x94, 0xa2, 0xda, 0x2a, 0x21, 0xa9,
	0x5d, 0xea, 0x96, 0xac, 0xdf, 0x82, 0x4e, 0x50, 0x8a, 0x6a, 0x8b, 0x1e, 0x90, 0x8e, 0xe3, 0x89,
	0x65, 0x3b, 0x88, 0xf4, 0x10, 0x99, 0xf2, 0xc4, 0x46, 0x28, 0x4f, 0x7e, 0x79, 0x64, 0xaf, 0xfd,
	0x35, 0x5e, 0x96, 0x73, 0x41, 0x05, 0xe9, 0xbe, 0xd5, 0x57, 0x70, 0x29, 0xe9, 0x93, 0x9b, 0x4c,
	0x4c, 0x75, 0x41, 0xfc, 0xa7, 0x37, 0x62, 0xab, 0x39, 0x1d, 0xdd, 0xa2, 0x6f, 0xc8, 0xde, 0xab,
	0x85, 0x14, 0x97, 0xa7, 0xad, 0xfb, 0x49, 0xef, 0xfd, 0x33, 0x57, 0x27, 0xe5, 0xd5, 0xf4, 0x1f,
	0x60, 0x76, 0x1b, 0xfd, 0x9b, 0x74, 0xfc, 0xec, 0x53, 0x90, 0xa4, 0x6e, 0x51, 0xcc, 0x02, 0x01,
	0x2a, 0x54, 0xa9, 0x30, 0x60, 0x61, 0xee, 0x42, 0x05, 0x22, 0x34, 0xb9, 0x08, 0x9b, 0x4a, 0x61,
	0x5d, 0x69, 0xd6, 0xc5, 0xf8, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x70, 0xf8, 0x48,
	0x9a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SharedFolderAgentClient is the client API for SharedFolderAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedFolderAgentClient interface {
	Invoke(ctx context.Context, in *SharedFolderRequest, opts ...grpc.CallOption) (*SharedFolderResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type sharedFolderAgentClient struct {
	cc *grpc.ClientConn
}

func NewSharedFolderAgentClient(cc *grpc.ClientConn) SharedFolderAgentClient {
	return &sharedFolderAgentClient{cc}
}

func (c *sharedFolderAgentClient) Invoke(ctx context.Context, in *SharedFolderRequest, opts ...grpc.CallOption) (*SharedFolderResponse, error) {
	out := new(SharedFolderResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.SharedFolderAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedFolderAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.SharedFolderAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedFolderAgentServer is the server API for SharedFolderAgent service.
type SharedFolderAgentServer interface {
	Invoke(context.Context, *SharedFolderRequest) (*SharedFolderResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedSharedFolderAgentServer can be embedded to have forward compatible implementations.
type UnimplementedSharedFolderAgentServer struct {
}

func (*UnimplementedSharedFolderAgentServer) Invoke(ctx context.Context, req *SharedFolderRequest) (*SharedFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedSharedFolderAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterSharedFolderAgentServer(s *grpc.Server, srv SharedFolderAgentServer) {
	s.RegisterService(&_SharedFolderAgent_serviceDesc, srv)
}

func _SharedFolderAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedFolderAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.SharedFolderAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedFolderAgentServer).Invoke(ctx, req.(*SharedFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedFolderAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedFolderAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.SharedFolderAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedFolderAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedFolderAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.SharedFolderAgent",
	HandlerType: (*SharedFolderAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _SharedFolderAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _SharedFolderAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_sharedfolder.proto",
}
