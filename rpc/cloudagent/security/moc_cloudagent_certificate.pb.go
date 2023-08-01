// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_certificate.proto

package security

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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CSRRequest struct {
	CSRs                 []*CertificateSigningRequest `protobuf:"bytes,1,rep,name=CSRs,proto3" json:"CSRs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CSRRequest) Reset()         { *m = CSRRequest{} }
func (m *CSRRequest) String() string { return proto.CompactTextString(m) }
func (*CSRRequest) ProtoMessage()    {}
func (*CSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{1}
}

func (m *CSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSRRequest.Unmarshal(m, b)
}
func (m *CSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSRRequest.Marshal(b, m, deterministic)
}
func (m *CSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRRequest.Merge(m, src)
}
func (m *CSRRequest) XXX_Size() int {
	return xxx_messageInfo_CSRRequest.Size(m)
}
func (m *CSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CSRRequest proto.InternalMessageInfo

func (m *CSRRequest) GetCSRs() []*CertificateSigningRequest {
	if m != nil {
		return m.CSRs
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{2}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CertificateSigningRequest struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Csr                  string              `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	OldCertificate       string              `protobuf:"bytes,3,opt,name=oldCertificate,proto3" json:"oldCertificate,omitempty"`
	Status               *common.Status      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CaName               string              `protobuf:"bytes,5,opt,name=caName,proto3" json:"caName,omitempty"`
	GroupName            string              `protobuf:"bytes,6,opt,name=groupName,proto3" json:"groupName,omitempty"`
	VaultName            string              `protobuf:"bytes,7,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	LocationName         string              `protobuf:"bytes,8,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Identity             string              `protobuf:"bytes,9,opt,name=identity,proto3" json:"identity,omitempty"`
	Validity             int64               `protobuf:"varint,10,opt,name=validity,proto3" json:"validity,omitempty"`
	IsCA                 *wrappers.BoolValue `protobuf:"bytes,11,opt,name=isCA,proto3" json:"isCA,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateSigningRequest) Reset()         { *m = CertificateSigningRequest{} }
func (m *CertificateSigningRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateSigningRequest) ProtoMessage()    {}
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{3}
}

func (m *CertificateSigningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateSigningRequest.Unmarshal(m, b)
}
func (m *CertificateSigningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateSigningRequest.Marshal(b, m, deterministic)
}
func (m *CertificateSigningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateSigningRequest.Merge(m, src)
}
func (m *CertificateSigningRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateSigningRequest.Size(m)
}
func (m *CertificateSigningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateSigningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateSigningRequest proto.InternalMessageInfo

func (m *CertificateSigningRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateSigningRequest) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

func (m *CertificateSigningRequest) GetOldCertificate() string {
	if m != nil {
		return m.OldCertificate
	}
	return ""
}

func (m *CertificateSigningRequest) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateSigningRequest) GetCaName() string {
	if m != nil {
		return m.CaName
	}
	return ""
}

func (m *CertificateSigningRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *CertificateSigningRequest) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *CertificateSigningRequest) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *CertificateSigningRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CertificateSigningRequest) GetValidity() int64 {
	if m != nil {
		return m.Validity
	}
	return 0
}

func (m *CertificateSigningRequest) GetIsCA() *wrappers.BoolValue {
	if m != nil {
		return m.IsCA
	}
	return nil
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          string          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.cloudagent.security.CertificateType" json:"type,omitempty"`
	GroupName            string          `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	VaultName            string          `protobuf:"bytes,9,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	LocationName         string          `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Tags                 *common.Tags    `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_373194e28b9c267a, []int{4}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *Certificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Certificate) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Certificate) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *Certificate) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.cloudagent.security.CertificateRequest")
	proto.RegisterType((*CSRRequest)(nil), "moc.cloudagent.security.CSRRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.cloudagent.security.CertificateResponse")
	proto.RegisterType((*CertificateSigningRequest)(nil), "moc.cloudagent.security.CertificateSigningRequest")
	proto.RegisterType((*Certificate)(nil), "moc.cloudagent.security.Certificate")
}

func init() { proto.RegisterFile("moc_cloudagent_certificate.proto", fileDescriptor_373194e28b9c267a) }

var fileDescriptor_373194e28b9c267a = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0x13, 0xd7, 0xbf, 0x66, 0x52, 0xe5, 0x57, 0x2d, 0xa8, 0x35, 0x11, 0xa0, 0x28, 0x45,
	0x28, 0x40, 0x65, 0x23, 0x73, 0xe5, 0xd2, 0x84, 0x7f, 0x27, 0x90, 0x36, 0x85, 0x03, 0x20, 0x2a,
	0xc7, 0x9e, 0x98, 0x95, 0x6c, 0xaf, 0xd9, 0x5d, 0xb7, 0xca, 0x13, 0xc0, 0xab, 0xf0, 0x00, 0x3c,
	0x1a, 0x17, 0x4e, 0xc8, 0x6b, 0x27, 0x71, 0xaa, 0x86, 0x06, 0x89, 0x9e, 0xe2, 0x9d, 0x6f, 0xe6,
	0x9b, 0x99, 0xef, 0xcb, 0x2e, 0xf4, 0x12, 0x1e, 0x9c, 0x06, 0x31, 0xcf, 0x43, 0x3f, 0xc2, 0x54,
	0x9d, 0x06, 0x28, 0x14, 0x9b, 0xb2, 0xc0, 0x57, 0xe8, 0x64, 0x82, 0x2b, 0x4e, 0x0e, 0x12, 0x1e,
	0x38, 0xcb, 0x0c, 0x47, 0x62, 0x90, 0x0b, 0xa6, 0x66, 0xdd, 0xbb, 0x11, 0xe7, 0x51, 0x8c, 0xae,
	0x4e, 0x9b, 0xe4, 0x53, 0xf7, 0x5c, 0xf8, 0x59, 0x86, 0x42, 0x96, 0x85, 0xdd, 0x03, 0x4d, 0xcd,
	0x93, 0x84, 0xa7, 0xd5, 0x4f, 0x09, 0xf4, 0x3f, 0x01, 0x19, 0x2d, 0xdb, 0x50, 0xfc, 0x92, 0xa3,
	0x54, 0xe4, 0x15, 0xec, 0xd6, 0xa2, 0xd2, 0x36, 0x7a, 0xcd, 0x41, 0xdb, 0xbb, 0xe7, 0xac, 0x69,
	0xef, 0xd4, 0x29, 0x56, 0x2a, 0xfb, 0x27, 0x00, 0xa3, 0x31, 0x9d, 0xf3, 0xbe, 0x00, 0x73, 0x34,
	0xa6, 0x73, 0x3e, 0x6f, 0x13, 0xbe, 0x31, 0x8b, 0x52, 0x96, 0x46, 0x15, 0x03, 0xd5, 0xf5, 0xfd,
	0xef, 0x06, 0xdc, 0x58, 0x19, 0x5b, 0x66, 0x3c, 0x95, 0xf8, 0xef, 0xe6, 0x26, 0x1e, 0x58, 0x14,
	0x65, 0x1e, 0x2b, 0xbb, 0xd1, 0x33, 0x06, 0x6d, 0xaf, 0xeb, 0x94, 0x0a, 0x3b, 0x73, 0x85, 0x9d,
	0x21, 0xe7, 0xf1, 0x3b, 0x3f, 0xce, 0x91, 0x56, 0x99, 0xe4, 0x26, 0x6c, 0x3f, 0x17, 0x82, 0x0b,
	0xbb, 0xd9, 0x33, 0x06, 0x2d, 0x5a, 0x1e, 0xfa, 0xbf, 0x1a, 0x70, 0x6b, 0xed, 0x3e, 0x84, 0x80,
	0x99, 0xfa, 0x09, 0xda, 0x86, 0x2e, 0xd1, 0xdf, 0x64, 0x1f, 0x9a, 0x81, 0x14, 0xba, 0x71, 0x6b,
	0x68, 0x7e, 0xfb, 0x61, 0x1b, 0xb4, 0x08, 0x90, 0x23, 0xe8, 0xf0, 0x38, 0xac, 0x71, 0x95, 0x8d,
	0xaa, 0x94, 0x0b, 0x18, 0x39, 0x04, 0x4b, 0x2a, 0x5f, 0xe5, 0xd2, 0x36, 0xf5, 0x06, 0x6d, 0xad,
	0xc2, 0x58, 0x87, 0x68, 0x05, 0x91, 0x7d, 0xb0, 0x02, 0xff, 0x75, 0x31, 0xc0, 0xb6, 0x1e, 0xa0,
	0x3a, 0x91, 0xdb, 0xd0, 0x8a, 0x04, 0xcf, 0x33, 0x0d, 0x59, 0x1a, 0x5a, 0x06, 0x0a, 0xf4, 0xcc,
	0xcf, 0x63, 0xa5, 0xd1, 0xff, 0x4a, 0x74, 0x11, 0x20, 0x7d, 0xd8, 0x8d, 0x79, 0xe0, 0x2b, 0xc6,
	0x53, 0x9d, 0xb0, 0xa3, 0x13, 0x56, 0x62, 0xa4, 0x0b, 0x3b, 0x2c, 0xc4, 0x54, 0x31, 0x35, 0xb3,
	0x5b, 0x1a, 0x5f, 0x9c, 0x0b, 0xec, 0xcc, 0x8f, 0x59, 0x58, 0x60, 0xd0, 0x33, 0x06, 0x4d, 0xba,
	0x38, 0x13, 0x07, 0x4c, 0x26, 0x47, 0xc7, 0x76, 0xfb, 0x4a, 0x53, 0x74, 0x5e, 0xff, 0x67, 0x03,
	0xda, 0x75, 0x51, 0x2e, 0x93, 0xbb, 0x03, 0x0d, 0x16, 0x96, 0x6a, 0xd3, 0x06, 0x0b, 0x8b, 0xed,
	0x52, 0xae, 0x86, 0x38, 0xe5, 0xa2, 0x54, 0xb8, 0x49, 0x97, 0x81, 0x62, 0xba, 0x94, 0xab, 0xe3,
	0xa9, 0x42, 0xa1, 0x85, 0x6d, 0xd2, 0xc5, 0x99, 0xdc, 0x87, 0x76, 0xed, 0xce, 0x96, 0x92, 0x56,
	0xee, 0xd4, 0x81, 0x9a, 0x35, 0xd6, 0x7a, 0x6b, 0x9e, 0x82, 0xa9, 0x66, 0x59, 0xa9, 0x6f, 0xc7,
	0x1b, 0x6c, 0xf2, 0x1f, 0x3e, 0x99, 0x65, 0x48, 0x75, 0xd5, 0xaa, 0x81, 0x3b, 0x7f, 0x34, 0xb0,
	0x75, 0x95, 0x81, 0x70, 0x89, 0x81, 0x77, 0xc0, 0x54, 0x7e, 0x24, 0x2b, 0x23, 0x5a, 0x7a, 0xba,
	0x13, 0x3f, 0x92, 0x54, 0x87, 0x1f, 0x3e, 0x80, 0xff, 0x2f, 0xcc, 0x45, 0x00, 0xac, 0x51, 0xcc,
	0x30, 0x55, 0x7b, 0x5b, 0xc5, 0xf7, 0x18, 0xc5, 0x19, 0x8a, 0x3d, 0xc3, 0xfb, 0x6a, 0xc2, 0x5e,
	0x2d, 0xf7, 0xb8, 0xd8, 0x8e, 0x24, 0xd0, 0x19, 0x09, 0xf4, 0x15, 0xbe, 0x11, 0x6f, 0xb3, 0xb0,
	0xd0, 0xec, 0xd1, 0x46, 0x97, 0xb8, 0xbc, 0x55, 0xdd, 0xa3, 0xcd, 0x92, 0xcb, 0x57, 0xa3, 0xbf,
	0x45, 0x26, 0xd0, 0x7c, 0x89, 0xea, 0x7a, 0x7b, 0x20, 0x58, 0xcf, 0x30, 0xc6, 0xeb, 0x5e, 0xe5,
	0x03, 0x98, 0xc5, 0x13, 0x43, 0x0e, 0xd7, 0xd7, 0x2d, 0xde, 0xe3, 0xbf, 0x26, 0xff, 0x08, 0xdb,
	0x14, 0x53, 0x3c, 0xbf, 0x16, 0xf6, 0xa1, 0xf7, 0xfe, 0x71, 0xc4, 0xd4, 0xe7, 0x7c, 0xe2, 0x04,
	0x3c, 0x71, 0x13, 0x16, 0x08, 0x2e, 0xf9, 0x54, 0xb9, 0x09, 0x0f, 0x5c, 0x91, 0x05, 0xee, 0x92,
	0xc9, 0x9d, 0x33, 0x4d, 0x2c, 0x7d, 0xf5, 0x9f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x84, 0xa5,
	0xb9, 0x84, 0x3c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.CertificateAgent/Renew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Sign(context.Context, *CSRRequest) (*CertificateResponse, error)
	Renew(context.Context, *CSRRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCertificateAgentServer) Sign(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (*UnimplementedCertificateAgentServer) Renew(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Renew not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Sign(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.CertificateAgent/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Renew(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _CertificateAgent_Sign_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _CertificateAgent_Renew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_certificate.proto",
}
