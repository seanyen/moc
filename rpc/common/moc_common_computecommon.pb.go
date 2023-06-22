// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_computecommon.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type UserType int32

const (
	UserType_ROOT UserType = 0
	UserType_USER UserType = 1
)

var UserType_name = map[int32]string{
	0: "ROOT",
	1: "USER",
}

var UserType_value = map[string]int32{
	"ROOT": 0,
	"USER": 1,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

type ProcessorType int32

const (
	ProcessorType_None    ProcessorType = 0
	ProcessorType_Intel   ProcessorType = 1
	ProcessorType_Intel64 ProcessorType = 2
	ProcessorType_AMD     ProcessorType = 3
	ProcessorType_AMD64   ProcessorType = 4
	ProcessorType_ARM     ProcessorType = 5
	ProcessorType_ARM64   ProcessorType = 6
)

var ProcessorType_name = map[int32]string{
	0: "None",
	1: "Intel",
	2: "Intel64",
	3: "AMD",
	4: "AMD64",
	5: "ARM",
	6: "ARM64",
}

var ProcessorType_value = map[string]int32{
	"None":    0,
	"Intel":   1,
	"Intel64": 2,
	"AMD":     3,
	"AMD64":   4,
	"ARM":     5,
	"ARM64":   6,
}

func (x ProcessorType) String() string {
	return proto.EnumName(ProcessorType_name, int32(x))
}

func (ProcessorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{1}
}

type OperatingSystemBootstrapEngine int32

const (
	OperatingSystemBootstrapEngine_CLOUD_INIT           OperatingSystemBootstrapEngine = 0
	OperatingSystemBootstrapEngine_WINDOWS_ANSWER_FILES OperatingSystemBootstrapEngine = 1
)

var OperatingSystemBootstrapEngine_name = map[int32]string{
	0: "CLOUD_INIT",
	1: "WINDOWS_ANSWER_FILES",
}

var OperatingSystemBootstrapEngine_value = map[string]int32{
	"CLOUD_INIT":           0,
	"WINDOWS_ANSWER_FILES": 1,
}

func (x OperatingSystemBootstrapEngine) String() string {
	return proto.EnumName(OperatingSystemBootstrapEngine_name, int32(x))
}

func (OperatingSystemBootstrapEngine) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{2}
}

type OperatingSystemType int32

const (
	OperatingSystemType_WINDOWS OperatingSystemType = 0
	OperatingSystemType_LINUX   OperatingSystemType = 1
)

var OperatingSystemType_name = map[int32]string{
	0: "WINDOWS",
	1: "LINUX",
}

var OperatingSystemType_value = map[string]int32{
	"WINDOWS": 0,
	"LINUX":   1,
}

func (x OperatingSystemType) String() string {
	return proto.EnumName(OperatingSystemType_name, int32(x))
}

func (OperatingSystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{3}
}

type VirtualMachineSizeType int32

const (
	VirtualMachineSizeType_Default             VirtualMachineSizeType = 0
	VirtualMachineSizeType_Standard_A2_v2      VirtualMachineSizeType = 2
	VirtualMachineSizeType_Standard_A4_v2      VirtualMachineSizeType = 3
	VirtualMachineSizeType_Standard_D2s_v3     VirtualMachineSizeType = 4
	VirtualMachineSizeType_Standard_D4s_v3     VirtualMachineSizeType = 5
	VirtualMachineSizeType_Standard_D8s_v3     VirtualMachineSizeType = 6
	VirtualMachineSizeType_Standard_D16s_v3    VirtualMachineSizeType = 7
	VirtualMachineSizeType_Standard_D32s_v3    VirtualMachineSizeType = 8
	VirtualMachineSizeType_Standard_DS2_v2     VirtualMachineSizeType = 9
	VirtualMachineSizeType_Standard_DS3_v2     VirtualMachineSizeType = 10
	VirtualMachineSizeType_Standard_DS4_v2     VirtualMachineSizeType = 11
	VirtualMachineSizeType_Standard_DS5_v2     VirtualMachineSizeType = 12
	VirtualMachineSizeType_Standard_DS13_v2    VirtualMachineSizeType = 13
	VirtualMachineSizeType_Standard_K8S_v1     VirtualMachineSizeType = 14
	VirtualMachineSizeType_Standard_K8S2_v1    VirtualMachineSizeType = 15
	VirtualMachineSizeType_Standard_K8S3_v1    VirtualMachineSizeType = 16
	VirtualMachineSizeType_Standard_K8S4_v1    VirtualMachineSizeType = 17
	VirtualMachineSizeType_Standard_NK6        VirtualMachineSizeType = 18
	VirtualMachineSizeType_Standard_NK12       VirtualMachineSizeType = 19
	VirtualMachineSizeType_Standard_NV6        VirtualMachineSizeType = 20
	VirtualMachineSizeType_Standard_NV12       VirtualMachineSizeType = 21
	VirtualMachineSizeType_Standard_K8S5_v1    VirtualMachineSizeType = 22
	VirtualMachineSizeType_Standard_DS2_v2_HPN VirtualMachineSizeType = 23
	VirtualMachineSizeType_Standard_DS3_v2_HPN VirtualMachineSizeType = 24
	VirtualMachineSizeType_Standard_DS4_v2_HPN VirtualMachineSizeType = 25
	VirtualMachineSizeType_Standard_F2s_HPN    VirtualMachineSizeType = 26
	VirtualMachineSizeType_Standard_F4s_HPN    VirtualMachineSizeType = 27
	VirtualMachineSizeType_Standard_F8s_HPN    VirtualMachineSizeType = 28
	VirtualMachineSizeType_Standard_F16s_HPN   VirtualMachineSizeType = 29
	VirtualMachineSizeType_Standard_NC4_A2     VirtualMachineSizeType = 30
	VirtualMachineSizeType_Standard_NC8_A2     VirtualMachineSizeType = 31
	VirtualMachineSizeType_Standard_NC16_A2    VirtualMachineSizeType = 32
	VirtualMachineSizeType_Standard_NC32_A2    VirtualMachineSizeType = 33
	VirtualMachineSizeType_Custom_A2           VirtualMachineSizeType = 95
	VirtualMachineSizeType_Custom_NK           VirtualMachineSizeType = 96
	VirtualMachineSizeType_Custom_Gpupv        VirtualMachineSizeType = 97
	VirtualMachineSizeType_Custom              VirtualMachineSizeType = 98
	VirtualMachineSizeType_Unsupported         VirtualMachineSizeType = 99
)

var VirtualMachineSizeType_name = map[int32]string{
	0:  "Default",
	2:  "Standard_A2_v2",
	3:  "Standard_A4_v2",
	4:  "Standard_D2s_v3",
	5:  "Standard_D4s_v3",
	6:  "Standard_D8s_v3",
	7:  "Standard_D16s_v3",
	8:  "Standard_D32s_v3",
	9:  "Standard_DS2_v2",
	10: "Standard_DS3_v2",
	11: "Standard_DS4_v2",
	12: "Standard_DS5_v2",
	13: "Standard_DS13_v2",
	14: "Standard_K8S_v1",
	15: "Standard_K8S2_v1",
	16: "Standard_K8S3_v1",
	17: "Standard_K8S4_v1",
	18: "Standard_NK6",
	19: "Standard_NK12",
	20: "Standard_NV6",
	21: "Standard_NV12",
	22: "Standard_K8S5_v1",
	23: "Standard_DS2_v2_HPN",
	24: "Standard_DS3_v2_HPN",
	25: "Standard_DS4_v2_HPN",
	26: "Standard_F2s_HPN",
	27: "Standard_F4s_HPN",
	28: "Standard_F8s_HPN",
	29: "Standard_F16s_HPN",
	30: "Standard_NC4_A2",
	31: "Standard_NC8_A2",
	32: "Standard_NC16_A2",
	33: "Standard_NC32_A2",
	95: "Custom_A2",
	96: "Custom_NK",
	97: "Custom_Gpupv",
	98: "Custom",
	99: "Unsupported",
}

var VirtualMachineSizeType_value = map[string]int32{
	"Default":             0,
	"Standard_A2_v2":      2,
	"Standard_A4_v2":      3,
	"Standard_D2s_v3":     4,
	"Standard_D4s_v3":     5,
	"Standard_D8s_v3":     6,
	"Standard_D16s_v3":    7,
	"Standard_D32s_v3":    8,
	"Standard_DS2_v2":     9,
	"Standard_DS3_v2":     10,
	"Standard_DS4_v2":     11,
	"Standard_DS5_v2":     12,
	"Standard_DS13_v2":    13,
	"Standard_K8S_v1":     14,
	"Standard_K8S2_v1":    15,
	"Standard_K8S3_v1":    16,
	"Standard_K8S4_v1":    17,
	"Standard_NK6":        18,
	"Standard_NK12":       19,
	"Standard_NV6":        20,
	"Standard_NV12":       21,
	"Standard_K8S5_v1":    22,
	"Standard_DS2_v2_HPN": 23,
	"Standard_DS3_v2_HPN": 24,
	"Standard_DS4_v2_HPN": 25,
	"Standard_F2s_HPN":    26,
	"Standard_F4s_HPN":    27,
	"Standard_F8s_HPN":    28,
	"Standard_F16s_HPN":   29,
	"Standard_NC4_A2":     30,
	"Standard_NC8_A2":     31,
	"Standard_NC16_A2":    32,
	"Standard_NC32_A2":    33,
	"Custom_A2":           95,
	"Custom_NK":           96,
	"Custom_Gpupv":        97,
	"Custom":              98,
	"Unsupported":         99,
}

func (x VirtualMachineSizeType) String() string {
	return proto.EnumName(VirtualMachineSizeType_name, int32(x))
}

func (VirtualMachineSizeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{4}
}

type WinRMProtocolType int32

const (
	WinRMProtocolType_HTTP  WinRMProtocolType = 0
	WinRMProtocolType_HTTPS WinRMProtocolType = 1
)

var WinRMProtocolType_name = map[int32]string{
	0: "HTTP",
	1: "HTTPS",
}

var WinRMProtocolType_value = map[string]int32{
	"HTTP":  0,
	"HTTPS": 1,
}

func (x WinRMProtocolType) String() string {
	return proto.EnumName(WinRMProtocolType_name, int32(x))
}

func (WinRMProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{5}
}

type PowerState int32

const (
	PowerState_Unknown  PowerState = 0
	PowerState_Running  PowerState = 1
	PowerState_Off      PowerState = 2
	PowerState_Paused   PowerState = 3
	PowerState_Critical PowerState = 4
)

var PowerState_name = map[int32]string{
	0: "Unknown",
	1: "Running",
	2: "Off",
	3: "Paused",
	4: "Critical",
}

var PowerState_value = map[string]int32{
	"Unknown":  0,
	"Running":  1,
	"Off":      2,
	"Paused":   3,
	"Critical": 4,
}

func (x PowerState) String() string {
	return proto.EnumName(PowerState_name, int32(x))
}

func (PowerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{6}
}

type VirtualMachineOperation int32

const (
	VirtualMachineOperation_START VirtualMachineOperation = 0
	VirtualMachineOperation_STOP  VirtualMachineOperation = 1
	VirtualMachineOperation_RESET VirtualMachineOperation = 2
)

var VirtualMachineOperation_name = map[int32]string{
	0: "START",
	1: "STOP",
	2: "RESET",
}

var VirtualMachineOperation_value = map[string]int32{
	"START": 0,
	"STOP":  1,
	"RESET": 2,
}

func (x VirtualMachineOperation) String() string {
	return proto.EnumName(VirtualMachineOperation_name, int32(x))
}

func (VirtualMachineOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{7}
}

type VirtualMachineRunCommandExecutionState int32

const (
	VirtualMachineRunCommandExecutionState_ExecutionState_UNKNOWN   VirtualMachineRunCommandExecutionState = 0
	VirtualMachineRunCommandExecutionState_ExecutionState_FAILED    VirtualMachineRunCommandExecutionState = 1
	VirtualMachineRunCommandExecutionState_ExecutionState_SUCCEEDED VirtualMachineRunCommandExecutionState = 2
)

var VirtualMachineRunCommandExecutionState_name = map[int32]string{
	0: "ExecutionState_UNKNOWN",
	1: "ExecutionState_FAILED",
	2: "ExecutionState_SUCCEEDED",
}

var VirtualMachineRunCommandExecutionState_value = map[string]int32{
	"ExecutionState_UNKNOWN":   0,
	"ExecutionState_FAILED":    1,
	"ExecutionState_SUCCEEDED": 2,
}

func (x VirtualMachineRunCommandExecutionState) String() string {
	return proto.EnumName(VirtualMachineRunCommandExecutionState_name, int32(x))
}

func (VirtualMachineRunCommandExecutionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{8}
}

type Architecture int32

const (
	Architecture_x86                Architecture = 0
	Architecture_MIPS               Architecture = 1
	Architecture_Alpha              Architecture = 2
	Architecture_PowerPC            Architecture = 3
	Architecture_ARM_Architecture   Architecture = 5
	Architecture_ia64               Architecture = 6
	Architecture_x64                Architecture = 9
	Architecture_ARM64_Architecture Architecture = 12
)

var Architecture_name = map[int32]string{
	0:  "x86",
	1:  "MIPS",
	2:  "Alpha",
	3:  "PowerPC",
	5:  "ARM_Architecture",
	6:  "ia64",
	9:  "x64",
	12: "ARM64_Architecture",
}

var Architecture_value = map[string]int32{
	"x86":                0,
	"MIPS":               1,
	"Alpha":              2,
	"PowerPC":            3,
	"ARM_Architecture":   5,
	"ia64":               6,
	"x64":                9,
	"ARM64_Architecture": 12,
}

func (x Architecture) String() string {
	return proto.EnumName(Architecture_name, int32(x))
}

func (Architecture) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{9}
}

type SecurityType int32

const (
	SecurityType_NOTCONFIGURED  SecurityType = 0
	SecurityType_TRUSTEDLAUNCH  SecurityType = 1
	SecurityType_CONFIDENTIALVM SecurityType = 2
)

var SecurityType_name = map[int32]string{
	0: "NOTCONFIGURED",
	1: "TRUSTEDLAUNCH",
	2: "CONFIDENTIALVM",
}

var SecurityType_value = map[string]int32{
	"NOTCONFIGURED":  0,
	"TRUSTEDLAUNCH":  1,
	"CONFIDENTIALVM": 2,
}

func (x SecurityType) String() string {
	return proto.EnumName(SecurityType_name, int32(x))
}

func (SecurityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{10}
}

type WinRMListener struct {
	Protocol             WinRMProtocolType `protobuf:"varint,1,opt,name=Protocol,proto3,enum=moc.WinRMProtocolType" json:"Protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WinRMListener) Reset()         { *m = WinRMListener{} }
func (m *WinRMListener) String() string { return proto.CompactTextString(m) }
func (*WinRMListener) ProtoMessage()    {}
func (*WinRMListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

func (m *WinRMListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinRMListener.Unmarshal(m, b)
}
func (m *WinRMListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinRMListener.Marshal(b, m, deterministic)
}
func (m *WinRMListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinRMListener.Merge(m, src)
}
func (m *WinRMListener) XXX_Size() int {
	return xxx_messageInfo_WinRMListener.Size(m)
}
func (m *WinRMListener) XXX_DiscardUnknown() {
	xxx_messageInfo_WinRMListener.DiscardUnknown(m)
}

var xxx_messageInfo_WinRMListener proto.InternalMessageInfo

func (m *WinRMListener) GetProtocol() WinRMProtocolType {
	if m != nil {
		return m.Protocol
	}
	return WinRMProtocolType_HTTP
}

type WinRMConfiguration struct {
	Listeners            []*WinRMListener `protobuf:"bytes,1,rep,name=Listeners,proto3" json:"Listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WinRMConfiguration) Reset()         { *m = WinRMConfiguration{} }
func (m *WinRMConfiguration) String() string { return proto.CompactTextString(m) }
func (*WinRMConfiguration) ProtoMessage()    {}
func (*WinRMConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{1}
}

func (m *WinRMConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinRMConfiguration.Unmarshal(m, b)
}
func (m *WinRMConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinRMConfiguration.Marshal(b, m, deterministic)
}
func (m *WinRMConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinRMConfiguration.Merge(m, src)
}
func (m *WinRMConfiguration) XXX_Size() int {
	return xxx_messageInfo_WinRMConfiguration.Size(m)
}
func (m *WinRMConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WinRMConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WinRMConfiguration proto.InternalMessageInfo

func (m *WinRMConfiguration) GetListeners() []*WinRMListener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

type VirtualMachineCustomSize struct {
	CpuCount             int32    `protobuf:"varint,1,opt,name=cpuCount,proto3" json:"cpuCount,omitempty"`
	MemoryMB             int32    `protobuf:"varint,2,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	GpuCount             int32    `protobuf:"varint,3,opt,name=gpuCount,proto3" json:"gpuCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineCustomSize) Reset()         { *m = VirtualMachineCustomSize{} }
func (m *VirtualMachineCustomSize) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineCustomSize) ProtoMessage()    {}
func (*VirtualMachineCustomSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{2}
}

func (m *VirtualMachineCustomSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineCustomSize.Unmarshal(m, b)
}
func (m *VirtualMachineCustomSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineCustomSize.Marshal(b, m, deterministic)
}
func (m *VirtualMachineCustomSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineCustomSize.Merge(m, src)
}
func (m *VirtualMachineCustomSize) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineCustomSize.Size(m)
}
func (m *VirtualMachineCustomSize) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineCustomSize.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineCustomSize proto.InternalMessageInfo

func (m *VirtualMachineCustomSize) GetCpuCount() int32 {
	if m != nil {
		return m.CpuCount
	}
	return 0
}

func (m *VirtualMachineCustomSize) GetMemoryMB() int32 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

func (m *VirtualMachineCustomSize) GetGpuCount() int32 {
	if m != nil {
		return m.GpuCount
	}
	return 0
}

type DynamicMemoryConfiguration struct {
	MaximumMemoryMB      uint64   `protobuf:"varint,1,opt,name=maximumMemoryMB,proto3" json:"maximumMemoryMB,omitempty"`
	MinimumMemoryMB      uint64   `protobuf:"varint,2,opt,name=minimumMemoryMB,proto3" json:"minimumMemoryMB,omitempty"`
	TargetMemoryBuffer   uint32   `protobuf:"varint,3,opt,name=targetMemoryBuffer,proto3" json:"targetMemoryBuffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynamicMemoryConfiguration) Reset()         { *m = DynamicMemoryConfiguration{} }
func (m *DynamicMemoryConfiguration) String() string { return proto.CompactTextString(m) }
func (*DynamicMemoryConfiguration) ProtoMessage()    {}
func (*DynamicMemoryConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{3}
}

func (m *DynamicMemoryConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicMemoryConfiguration.Unmarshal(m, b)
}
func (m *DynamicMemoryConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicMemoryConfiguration.Marshal(b, m, deterministic)
}
func (m *DynamicMemoryConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicMemoryConfiguration.Merge(m, src)
}
func (m *DynamicMemoryConfiguration) XXX_Size() int {
	return xxx_messageInfo_DynamicMemoryConfiguration.Size(m)
}
func (m *DynamicMemoryConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicMemoryConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicMemoryConfiguration proto.InternalMessageInfo

func (m *DynamicMemoryConfiguration) GetMaximumMemoryMB() uint64 {
	if m != nil {
		return m.MaximumMemoryMB
	}
	return 0
}

func (m *DynamicMemoryConfiguration) GetMinimumMemoryMB() uint64 {
	if m != nil {
		return m.MinimumMemoryMB
	}
	return 0
}

func (m *DynamicMemoryConfiguration) GetTargetMemoryBuffer() uint32 {
	if m != nil {
		return m.TargetMemoryBuffer
	}
	return 0
}

type VirtualMachineRunCommandInputParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineRunCommandInputParameter) Reset() {
	*m = VirtualMachineRunCommandInputParameter{}
}
func (m *VirtualMachineRunCommandInputParameter) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandInputParameter) ProtoMessage()    {}
func (*VirtualMachineRunCommandInputParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{4}
}

func (m *VirtualMachineRunCommandInputParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandInputParameter.Merge(m, src)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Size(m)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandInputParameter.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandInputParameter proto.InternalMessageInfo

func (m *VirtualMachineRunCommandInputParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineRunCommandInputParameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VirtualMachineRunCommandScriptSource struct {
	Script               string   `protobuf:"bytes,1,opt,name=Script,proto3" json:"Script,omitempty"`
	ScriptURI            string   `protobuf:"bytes,2,opt,name=ScriptURI,proto3" json:"ScriptURI,omitempty"`
	CommandID            string   `protobuf:"bytes,3,opt,name=CommandID,proto3" json:"CommandID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineRunCommandScriptSource) Reset()         { *m = VirtualMachineRunCommandScriptSource{} }
func (m *VirtualMachineRunCommandScriptSource) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandScriptSource) ProtoMessage()    {}
func (*VirtualMachineRunCommandScriptSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{5}
}

func (m *VirtualMachineRunCommandScriptSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandScriptSource.Merge(m, src)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Size(m)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandScriptSource.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandScriptSource proto.InternalMessageInfo

func (m *VirtualMachineRunCommandScriptSource) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *VirtualMachineRunCommandScriptSource) GetScriptURI() string {
	if m != nil {
		return m.ScriptURI
	}
	return ""
}

func (m *VirtualMachineRunCommandScriptSource) GetCommandID() string {
	if m != nil {
		return m.CommandID
	}
	return ""
}

type VirtualMachineRunCommandInstanceView struct {
	ExecutionState       VirtualMachineRunCommandExecutionState `protobuf:"varint,1,opt,name=ExecutionState,proto3,enum=moc.VirtualMachineRunCommandExecutionState" json:"ExecutionState,omitempty"`
	ExitCode             int32                                  `protobuf:"varint,2,opt,name=ExitCode,proto3" json:"ExitCode,omitempty"`
	Output               string                                 `protobuf:"bytes,3,opt,name=Output,proto3" json:"Output,omitempty"`
	Error                string                                 `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *VirtualMachineRunCommandInstanceView) Reset()         { *m = VirtualMachineRunCommandInstanceView{} }
func (m *VirtualMachineRunCommandInstanceView) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandInstanceView) ProtoMessage()    {}
func (*VirtualMachineRunCommandInstanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{6}
}

func (m *VirtualMachineRunCommandInstanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandInstanceView.Merge(m, src)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Size(m)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandInstanceView.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandInstanceView proto.InternalMessageInfo

func (m *VirtualMachineRunCommandInstanceView) GetExecutionState() VirtualMachineRunCommandExecutionState {
	if m != nil {
		return m.ExecutionState
	}
	return VirtualMachineRunCommandExecutionState_ExecutionState_UNKNOWN
}

func (m *VirtualMachineRunCommandInstanceView) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *VirtualMachineRunCommandInstanceView) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *VirtualMachineRunCommandInstanceView) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// describes the the capabilities of this node in terms of what type of VM can be created
// In the future when CVM comes, another bool will be added to check hardware capability.
type VirtualMachineCapabilities struct {
	IsolatedVmCapable    bool     `protobuf:"varint,1,opt,name=IsolatedVmCapable,proto3" json:"IsolatedVmCapable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineCapabilities) Reset()         { *m = VirtualMachineCapabilities{} }
func (m *VirtualMachineCapabilities) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineCapabilities) ProtoMessage()    {}
func (*VirtualMachineCapabilities) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{7}
}

func (m *VirtualMachineCapabilities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineCapabilities.Unmarshal(m, b)
}
func (m *VirtualMachineCapabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineCapabilities.Marshal(b, m, deterministic)
}
func (m *VirtualMachineCapabilities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineCapabilities.Merge(m, src)
}
func (m *VirtualMachineCapabilities) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineCapabilities.Size(m)
}
func (m *VirtualMachineCapabilities) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineCapabilities.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineCapabilities proto.InternalMessageInfo

func (m *VirtualMachineCapabilities) GetIsolatedVmCapable() bool {
	if m != nil {
		return m.IsolatedVmCapable
	}
	return false
}

var E_Sensitivecompute = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50002,
	Name:          "moc.sensitivecompute",
	Tag:           "varint,50002,opt,name=sensitivecompute",
	Filename:      "moc_common_computecommon.proto",
}

func init() {
	proto.RegisterEnum("moc.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("moc.ProcessorType", ProcessorType_name, ProcessorType_value)
	proto.RegisterEnum("moc.OperatingSystemBootstrapEngine", OperatingSystemBootstrapEngine_name, OperatingSystemBootstrapEngine_value)
	proto.RegisterEnum("moc.OperatingSystemType", OperatingSystemType_name, OperatingSystemType_value)
	proto.RegisterEnum("moc.VirtualMachineSizeType", VirtualMachineSizeType_name, VirtualMachineSizeType_value)
	proto.RegisterEnum("moc.WinRMProtocolType", WinRMProtocolType_name, WinRMProtocolType_value)
	proto.RegisterEnum("moc.PowerState", PowerState_name, PowerState_value)
	proto.RegisterEnum("moc.VirtualMachineOperation", VirtualMachineOperation_name, VirtualMachineOperation_value)
	proto.RegisterEnum("moc.VirtualMachineRunCommandExecutionState", VirtualMachineRunCommandExecutionState_name, VirtualMachineRunCommandExecutionState_value)
	proto.RegisterEnum("moc.Architecture", Architecture_name, Architecture_value)
	proto.RegisterEnum("moc.SecurityType", SecurityType_name, SecurityType_value)
	proto.RegisterType((*WinRMListener)(nil), "moc.WinRMListener")
	proto.RegisterType((*WinRMConfiguration)(nil), "moc.WinRMConfiguration")
	proto.RegisterType((*VirtualMachineCustomSize)(nil), "moc.VirtualMachineCustomSize")
	proto.RegisterType((*DynamicMemoryConfiguration)(nil), "moc.DynamicMemoryConfiguration")
	proto.RegisterType((*VirtualMachineRunCommandInputParameter)(nil), "moc.VirtualMachineRunCommandInputParameter")
	proto.RegisterType((*VirtualMachineRunCommandScriptSource)(nil), "moc.VirtualMachineRunCommandScriptSource")
	proto.RegisterType((*VirtualMachineRunCommandInstanceView)(nil), "moc.VirtualMachineRunCommandInstanceView")
	proto.RegisterType((*VirtualMachineCapabilities)(nil), "moc.VirtualMachineCapabilities")
	proto.RegisterExtension(E_Sensitivecompute)
}

func init() { proto.RegisterFile("moc_common_computecommon.proto", fileDescriptor_7d1a061f6c82445b) }

var fileDescriptor_7d1a061f6c82445b = []byte{
	// 1310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xdd, 0x52, 0xe3, 0x36,
	0x1b, 0xc7, 0xc9, 0x07, 0x1f, 0x79, 0xf8, 0x12, 0x82, 0x65, 0xb3, 0xbc, 0x2c, 0x2f, 0x4d, 0x3f,
	0x86, 0x49, 0xdb, 0x50, 0x42, 0x36, 0x43, 0x7b, 0x16, 0x1c, 0xb3, 0x78, 0x21, 0x4e, 0xc6, 0x4e,
	0xc2, 0xb6, 0x27, 0xa9, 0x71, 0x94, 0xa0, 0x69, 0x2c, 0x79, 0x6c, 0x99, 0x85, 0x9d, 0xe9, 0x4c,
	0x8f, 0x7a, 0xdc, 0x4b, 0xd8, 0x9b, 0xe8, 0x25, 0xf4, 0x06, 0x7a, 0x45, 0x1d, 0xc9, 0x09, 0xac,
	0x4d, 0xb7, 0x47, 0xd1, 0xf3, 0xfb, 0x3f, 0xfa, 0xeb, 0xc9, 0x63, 0x49, 0x36, 0xec, 0x79, 0xdc,
	0x1d, 0xb8, 0xdc, 0xf3, 0x38, 0x93, 0x3f, 0x7e, 0x24, 0x48, 0x1c, 0x55, 0xfc, 0x80, 0x0b, 0x8e,
	0x73, 0x1e, 0x77, 0x77, 0xf6, 0xc7, 0x9c, 0x8f, 0x27, 0xe4, 0x50, 0xa1, 0xeb, 0x68, 0x74, 0x38,
	0x24, 0xa1, 0x1b, 0x50, 0x5f, 0xf0, 0x20, 0x4e, 0x2b, 0x69, 0xb0, 0x7a, 0x45, 0x99, 0xd5, 0xba,
	0xa4, 0xa1, 0x20, 0x8c, 0x04, 0xb8, 0x0a, 0x4b, 0x1d, 0xa9, 0xb8, 0x7c, 0x52, 0xcc, 0xec, 0x67,
	0x0e, 0xd6, 0xaa, 0xdb, 0x15, 0x8f, 0xbb, 0x15, 0x95, 0x35, 0x53, 0xba, 0xf7, 0x3e, 0xb1, 0x1e,
	0xf2, 0x4a, 0x67, 0x80, 0x95, 0xac, 0x71, 0x36, 0xa2, 0xe3, 0x28, 0x70, 0x04, 0xe5, 0x0c, 0x7f,
	0x07, 0x85, 0x99, 0x6b, 0x58, 0xcc, 0xec, 0xe7, 0x0e, 0x96, 0xab, 0xf8, 0xd1, 0x6a, 0x26, 0x59,
	0x8f, 0x49, 0x25, 0x06, 0xc5, 0x3e, 0x0d, 0x44, 0xe4, 0x4c, 0x5a, 0x8e, 0x7b, 0x43, 0x19, 0xd1,
	0xa2, 0x50, 0x70, 0xcf, 0xa6, 0xef, 0x09, 0xde, 0x81, 0x25, 0xd7, 0x8f, 0x34, 0x1e, 0x31, 0xa1,
	0xea, 0x9a, 0xb7, 0x1e, 0x62, 0xa9, 0x79, 0xc4, 0xe3, 0xc1, 0x7d, 0xeb, 0xb4, 0x98, 0x8d, 0xb5,
	0x59, 0x2c, 0xb5, 0xf1, 0x6c, 0x5e, 0x2e, 0xd6, 0x66, 0x71, 0xe9, 0x43, 0x06, 0x76, 0x9a, 0xf7,
	0xcc, 0xf1, 0xa8, 0xdb, 0x52, 0xf9, 0xc9, 0x3f, 0x70, 0x00, 0xeb, 0x9e, 0x73, 0x47, 0xbd, 0xc8,
	0x6b, 0xcd, 0xdc, 0xe5, 0xca, 0x79, 0x2b, 0x8d, 0x55, 0x26, 0x65, 0x89, 0xcc, 0xec, 0x34, 0x33,
	0x89, 0x71, 0x05, 0xb0, 0x70, 0x82, 0x31, 0x11, 0x31, 0x39, 0x8d, 0x46, 0x23, 0x12, 0xa8, 0xc2,
	0x56, 0xad, 0x7f, 0x51, 0x4a, 0x6f, 0xe1, 0xab, 0x64, 0x4b, 0xac, 0x88, 0x69, 0xdc, 0xf3, 0x1c,
	0x36, 0x34, 0x98, 0x1f, 0x89, 0x8e, 0x13, 0x38, 0x1e, 0x11, 0x24, 0xc0, 0x18, 0xf2, 0xa6, 0xe3,
	0x11, 0x55, 0x62, 0xc1, 0x52, 0x63, 0xbc, 0x03, 0xf3, 0x7d, 0x67, 0x12, 0x11, 0x55, 0x4d, 0xe1,
	0x34, 0xff, 0xc7, 0x9f, 0xc5, 0x8c, 0x15, 0xa3, 0xd2, 0x6f, 0x19, 0xf8, 0xe2, 0x53, 0xd6, 0xb6,
	0xda, 0x24, 0x36, 0x8f, 0x02, 0x97, 0xe0, 0x5d, 0x58, 0x88, 0xe3, 0xd8, 0x7a, 0xea, 0x32, 0x65,
	0x78, 0x17, 0x0a, 0xf1, 0xa8, 0x67, 0x19, 0xf1, 0x32, 0xd6, 0x23, 0x90, 0xea, 0xac, 0xda, 0xa6,
	0xfa, 0x97, 0x05, 0xeb, 0x11, 0x94, 0xfe, 0xfa, 0x8f, 0x12, 0x0c, 0x16, 0x0a, 0x87, 0xb9, 0xa4,
	0x4f, 0xc9, 0x3b, 0x6c, 0xc3, 0x9a, 0x7e, 0x47, 0xdc, 0x48, 0x3e, 0x16, 0x5b, 0x38, 0x82, 0x4c,
	0xb7, 0xe6, 0xd7, 0x6a, 0x3f, 0x7d, 0xca, 0x22, 0x39, 0xc5, 0x4a, 0x59, 0xc8, 0x9d, 0xa1, 0xdf,
	0x51, 0xa1, 0xf1, 0x21, 0x99, 0xed, 0x9a, 0x59, 0x8c, 0xb7, 0x61, 0xa1, 0x1d, 0x09, 0x3f, 0x12,
	0xd3, 0xa2, 0xa7, 0x11, 0xde, 0x82, 0x79, 0x3d, 0x08, 0x78, 0x50, 0xcc, 0x2b, 0x1c, 0x07, 0xa5,
	0x37, 0xb0, 0x93, 0xda, 0xb7, 0x8e, 0xef, 0x5c, 0xd3, 0x09, 0x15, 0x94, 0x84, 0xf8, 0x1b, 0xd8,
	0x30, 0x42, 0x3e, 0x71, 0x04, 0x19, 0xf6, 0x3d, 0xa5, 0x4c, 0xe2, 0xfa, 0x97, 0xac, 0xa7, 0x42,
	0x79, 0x0f, 0x96, 0x7a, 0x21, 0x09, 0xe4, 0x09, 0xc3, 0x4b, 0x90, 0xb7, 0xda, 0xed, 0x2e, 0x9a,
	0x93, 0xa3, 0x9e, 0xad, 0x5b, 0x28, 0x53, 0xfe, 0x11, 0x56, 0x3b, 0x01, 0x77, 0x49, 0x18, 0xf2,
	0x87, 0x24, 0x93, 0x33, 0x82, 0xe6, 0x70, 0x01, 0xe6, 0x0d, 0x26, 0xc8, 0x04, 0x65, 0xf0, 0x32,
	0x2c, 0xaa, 0x61, 0xbd, 0x86, 0xb2, 0x78, 0x11, 0x72, 0x8d, 0x56, 0x13, 0xe5, 0x64, 0x42, 0xa3,
	0xd5, 0xac, 0xd7, 0x50, 0x5e, 0x31, 0xab, 0x85, 0xe6, 0x15, 0xb3, 0x5a, 0xf5, 0x1a, 0x5a, 0x28,
	0xbf, 0x81, 0xbd, 0xb6, 0x4f, 0xe4, 0xe6, 0x67, 0x63, 0xfb, 0x3e, 0x14, 0xc4, 0x3b, 0xe5, 0x5c,
	0x84, 0x22, 0x70, 0x7c, 0x9d, 0x8d, 0x29, 0x23, 0x78, 0x0d, 0x40, 0xbb, 0x6c, 0xf7, 0x9a, 0x03,
	0xc3, 0x34, 0x64, 0x59, 0x45, 0xd8, 0xba, 0x32, 0xcc, 0x66, 0xfb, 0xca, 0x1e, 0x34, 0x4c, 0xfb,
	0x4a, 0xb7, 0x06, 0x67, 0xc6, 0xa5, 0x6e, 0xa3, 0x4c, 0xf9, 0x5b, 0xd8, 0x4c, 0x79, 0xa9, 0x62,
	0x97, 0x61, 0x71, 0x3a, 0x21, 0xae, 0xf7, 0xd2, 0x30, 0x7b, 0x6f, 0x51, 0xa6, 0xfc, 0x61, 0x01,
	0xb6, 0x93, 0x2d, 0x94, 0x87, 0x7e, 0x36, 0xa5, 0x49, 0x46, 0x4e, 0x34, 0x11, 0x68, 0x0e, 0x63,
	0x58, 0xb3, 0x85, 0xc3, 0x86, 0x4e, 0x30, 0x1c, 0x34, 0xaa, 0x83, 0xdb, 0x2a, 0xca, 0x26, 0x59,
	0x4d, 0xb2, 0x1c, 0xde, 0x84, 0xf5, 0x07, 0xd6, 0xac, 0x86, 0x83, 0xdb, 0x63, 0x94, 0x4f, 0xc2,
	0x9a, 0x82, 0xf3, 0x49, 0x78, 0xa2, 0xe0, 0x02, 0xde, 0x02, 0xf4, 0x08, 0x8f, 0xea, 0x8a, 0x2e,
	0x26, 0xe9, 0x71, 0xec, 0xba, 0x94, 0x34, 0xb0, 0x55, 0x4d, 0x85, 0x14, 0x3c, 0x96, 0x10, 0x52,
	0x50, 0x55, 0xba, 0x9c, 0x82, 0xaf, 0x24, 0x5c, 0x49, 0xae, 0x64, 0x1f, 0xa9, 0xf9, 0xab, 0x89,
	0xd4, 0x8b, 0x13, 0x7b, 0x70, 0x7b, 0x84, 0xd6, 0x12, 0xa9, 0x17, 0x27, 0x72, 0xfd, 0x23, 0xb4,
	0x9e, 0xa6, 0xc7, 0x92, 0xa2, 0x34, 0xad, 0x49, 0xba, 0x81, 0x11, 0xac, 0x3c, 0x50, 0xf3, 0xa2,
	0x8e, 0x30, 0xde, 0x80, 0xd5, 0x8f, 0xc8, 0x51, 0x15, 0x6d, 0x26, 0x93, 0xfa, 0x75, 0xb4, 0x95,
	0x4c, 0xea, 0x1f, 0x55, 0xd1, 0xb3, 0xb4, 0xff, 0x2b, 0xe9, 0xbf, 0x8d, 0x9f, 0xc3, 0x66, 0xaa,
	0x41, 0x83, 0xf3, 0x8e, 0x89, 0x9e, 0xa7, 0x84, 0xe3, 0x99, 0x50, 0x4c, 0x09, 0xb5, 0x99, 0xf0,
	0x22, 0xb1, 0xc0, 0x59, 0x35, 0x54, 0x74, 0x27, 0x49, 0x6b, 0x31, 0xfd, 0x5f, 0x92, 0x9e, 0xc4,
	0x74, 0x17, 0x3f, 0x83, 0x8d, 0x47, 0x2a, 0x9f, 0xac, 0xc4, 0x2f, 0x13, 0xad, 0x35, 0xb5, 0xda,
	0xa0, 0x51, 0x45, 0x7b, 0x29, 0x78, 0x22, 0xe1, 0xff, 0x13, 0xb6, 0xa6, 0x76, 0x54, 0x97, 0x74,
	0x3f, 0x45, 0x8f, 0xab, 0x92, 0x7e, 0x86, 0x57, 0xa1, 0x10, 0xbf, 0xc1, 0x64, 0x38, 0xf8, 0x28,
	0x34, 0x2f, 0xd0, 0xcf, 0xb2, 0xa5, 0xd3, 0xf0, 0xb5, 0x1f, 0xf9, 0xb7, 0xc8, 0xc1, 0x00, 0x0b,
	0x31, 0x41, 0xd7, 0x78, 0x1d, 0x96, 0x7b, 0x2c, 0x8c, 0x7c, 0x9f, 0x07, 0x82, 0x0c, 0x91, 0x5b,
	0x3e, 0x80, 0x8d, 0x27, 0xef, 0x60, 0x79, 0xf8, 0xcf, 0xbb, 0xdd, 0x4e, 0x7c, 0x98, 0xe4, 0x48,
	0x9e, 0x3d, 0x03, 0xa0, 0xc3, 0xdf, 0x91, 0x20, 0xbe, 0xe6, 0x96, 0x61, 0xb1, 0xc7, 0x7e, 0x61,
	0xfc, 0x1d, 0x43, 0x73, 0x32, 0xb0, 0x22, 0xc6, 0x28, 0x1b, 0xa3, 0x8c, 0xbc, 0x03, 0xda, 0xa3,
	0x11, 0xca, 0xca, 0x75, 0x3b, 0x4e, 0x14, 0x92, 0x21, 0xca, 0xe1, 0x15, 0x58, 0xd2, 0x02, 0x2a,
	0xa8, 0xeb, 0x4c, 0x50, 0xbe, 0xfc, 0x3d, 0x3c, 0x4f, 0x1e, 0xcb, 0xe9, 0xa1, 0xe6, 0x4c, 0x2e,
	0x68, 0x77, 0x1b, 0xd6, 0xf4, 0x76, 0xb2, 0xbb, 0xed, 0x0e, 0xca, 0x48, 0x68, 0xe9, 0xb6, 0xde,
	0x45, 0xd9, 0xf2, 0xaf, 0x9f, 0x7e, 0x73, 0x3d, 0xb9, 0x88, 0xb7, 0x93, 0x64, 0xd0, 0x33, 0x2f,
	0xcc, 0xf6, 0x95, 0x89, 0xe6, 0xf0, 0x0b, 0x78, 0x96, 0xd2, 0xce, 0x1a, 0xc6, 0xa5, 0xde, 0x44,
	0x19, 0xbc, 0x0b, 0xc5, 0x94, 0x64, 0xf7, 0x34, 0x4d, 0xd7, 0x9b, 0x7a, 0x13, 0x65, 0xcb, 0xef,
	0x61, 0xa5, 0x11, 0xb8, 0x37, 0x54, 0x10, 0x57, 0x44, 0x01, 0x91, 0x7f, 0xf6, 0xee, 0xa4, 0x1e,
	0x17, 0xdb, 0x32, 0x64, 0x9f, 0xd4, 0xd5, 0x37, 0xf1, 0x6f, 0x1c, 0x94, 0x95, 0x7d, 0x51, 0x2d,
	0xeb, 0x68, 0x28, 0x27, 0x1f, 0x66, 0xc3, 0x6a, 0x0d, 0x3e, 0x9e, 0x8e, 0xe6, 0xe5, 0x3c, 0xea,
	0xc8, 0x7b, 0x52, 0x59, 0xd5, 0x6b, 0xa8, 0x80, 0xb7, 0x01, 0xab, 0xbb, 0x33, 0x99, 0xba, 0x52,
	0x3e, 0x87, 0x15, 0x9b, 0xb8, 0x51, 0x40, 0xc5, 0xbd, 0x7a, 0x4a, 0x1b, 0xb0, 0x6a, 0xb6, 0xbb,
	0x5a, 0xdb, 0x3c, 0x33, 0x5e, 0xf7, 0x2c, 0xbd, 0x89, 0xe6, 0x24, 0xea, 0x5a, 0x3d, 0xbb, 0xab,
	0x37, 0x2f, 0x1b, 0x3d, 0x53, 0x3b, 0x47, 0x19, 0x79, 0x8f, 0xa9, 0x94, 0xa6, 0x6e, 0x76, 0x8d,
	0xc6, 0x65, 0xbf, 0x85, 0xb2, 0x3f, 0x5c, 0x00, 0x0a, 0x09, 0x0b, 0xa9, 0xa0, 0xb7, 0x64, 0xfa,
	0x95, 0x87, 0x5f, 0x56, 0xe2, 0xaf, 0xba, 0xca, 0xec, 0xab, 0xae, 0x72, 0x46, 0xc9, 0x64, 0xd8,
	0xf6, 0x65, 0x23, 0xc2, 0xe2, 0xdf, 0xbf, 0xe7, 0xd4, 0xbb, 0xe5, 0xc9, 0xc4, 0xd3, 0x2f, 0x7f,
	0xfa, 0x7c, 0x4c, 0xc5, 0x4d, 0x74, 0x5d, 0x71, 0xb9, 0x77, 0xe8, 0x51, 0x37, 0xe0, 0x21, 0x1f,
	0x89, 0x43, 0x8f, 0xbb, 0x87, 0x81, 0xef, 0x1e, 0xc6, 0xdf, 0x8f, 0xd7, 0x0b, 0xca, 0xf7, 0xf8,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x12, 0x94, 0x01, 0x62, 0x0a, 0x00, 0x00,
}
