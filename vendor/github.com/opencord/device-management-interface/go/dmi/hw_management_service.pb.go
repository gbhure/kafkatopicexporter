// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dmi/hw_management_service.proto

package dmi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PhysicalInventoryRequest struct {
	DeviceUuid           *Uuid    `protobuf:"bytes,1,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhysicalInventoryRequest) Reset()         { *m = PhysicalInventoryRequest{} }
func (m *PhysicalInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*PhysicalInventoryRequest) ProtoMessage()    {}
func (*PhysicalInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{0}
}

func (m *PhysicalInventoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalInventoryRequest.Unmarshal(m, b)
}
func (m *PhysicalInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalInventoryRequest.Marshal(b, m, deterministic)
}
func (m *PhysicalInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalInventoryRequest.Merge(m, src)
}
func (m *PhysicalInventoryRequest) XXX_Size() int {
	return xxx_messageInfo_PhysicalInventoryRequest.Size(m)
}
func (m *PhysicalInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalInventoryRequest proto.InternalMessageInfo

func (m *PhysicalInventoryRequest) GetDeviceUuid() *Uuid {
	if m != nil {
		return m.DeviceUuid
	}
	return nil
}

type PhysicalInventoryResponse struct {
	Status               Status    `protobuf:"varint,1,opt,name=status,proto3,enum=dmi.Status" json:"status,omitempty"`
	Reason               Reason    `protobuf:"varint,2,opt,name=reason,proto3,enum=dmi.Reason" json:"reason,omitempty"`
	Inventory            *Hardware `protobuf:"bytes,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PhysicalInventoryResponse) Reset()         { *m = PhysicalInventoryResponse{} }
func (m *PhysicalInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*PhysicalInventoryResponse) ProtoMessage()    {}
func (*PhysicalInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{1}
}

func (m *PhysicalInventoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalInventoryResponse.Unmarshal(m, b)
}
func (m *PhysicalInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalInventoryResponse.Marshal(b, m, deterministic)
}
func (m *PhysicalInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalInventoryResponse.Merge(m, src)
}
func (m *PhysicalInventoryResponse) XXX_Size() int {
	return xxx_messageInfo_PhysicalInventoryResponse.Size(m)
}
func (m *PhysicalInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalInventoryResponse proto.InternalMessageInfo

func (m *PhysicalInventoryResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNDEFINED_STATUS
}

func (m *PhysicalInventoryResponse) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_UNDEFINED_REASON
}

func (m *PhysicalInventoryResponse) GetInventory() *Hardware {
	if m != nil {
		return m.Inventory
	}
	return nil
}

type HWComponentInfoGetRequest struct {
	DeviceUuid           *Uuid    `protobuf:"bytes,1,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
	ComponentUuid        *Uuid    `protobuf:"bytes,2,opt,name=component_uuid,json=componentUuid,proto3" json:"component_uuid,omitempty"`
	ComponentName        string   `protobuf:"bytes,3,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HWComponentInfoGetRequest) Reset()         { *m = HWComponentInfoGetRequest{} }
func (m *HWComponentInfoGetRequest) String() string { return proto.CompactTextString(m) }
func (*HWComponentInfoGetRequest) ProtoMessage()    {}
func (*HWComponentInfoGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{2}
}

func (m *HWComponentInfoGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWComponentInfoGetRequest.Unmarshal(m, b)
}
func (m *HWComponentInfoGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWComponentInfoGetRequest.Marshal(b, m, deterministic)
}
func (m *HWComponentInfoGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWComponentInfoGetRequest.Merge(m, src)
}
func (m *HWComponentInfoGetRequest) XXX_Size() int {
	return xxx_messageInfo_HWComponentInfoGetRequest.Size(m)
}
func (m *HWComponentInfoGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HWComponentInfoGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HWComponentInfoGetRequest proto.InternalMessageInfo

func (m *HWComponentInfoGetRequest) GetDeviceUuid() *Uuid {
	if m != nil {
		return m.DeviceUuid
	}
	return nil
}

func (m *HWComponentInfoGetRequest) GetComponentUuid() *Uuid {
	if m != nil {
		return m.ComponentUuid
	}
	return nil
}

func (m *HWComponentInfoGetRequest) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

type HWComponentInfoSetRequest struct {
	DeviceUuid           *Uuid                `protobuf:"bytes,1,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
	ComponentUuid        *Uuid                `protobuf:"bytes,2,opt,name=component_uuid,json=componentUuid,proto3" json:"component_uuid,omitempty"`
	ComponentName        string               `protobuf:"bytes,3,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	Changes              *ModifiableComponent `protobuf:"bytes,4,opt,name=changes,proto3" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HWComponentInfoSetRequest) Reset()         { *m = HWComponentInfoSetRequest{} }
func (m *HWComponentInfoSetRequest) String() string { return proto.CompactTextString(m) }
func (*HWComponentInfoSetRequest) ProtoMessage()    {}
func (*HWComponentInfoSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{3}
}

func (m *HWComponentInfoSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWComponentInfoSetRequest.Unmarshal(m, b)
}
func (m *HWComponentInfoSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWComponentInfoSetRequest.Marshal(b, m, deterministic)
}
func (m *HWComponentInfoSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWComponentInfoSetRequest.Merge(m, src)
}
func (m *HWComponentInfoSetRequest) XXX_Size() int {
	return xxx_messageInfo_HWComponentInfoSetRequest.Size(m)
}
func (m *HWComponentInfoSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HWComponentInfoSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HWComponentInfoSetRequest proto.InternalMessageInfo

func (m *HWComponentInfoSetRequest) GetDeviceUuid() *Uuid {
	if m != nil {
		return m.DeviceUuid
	}
	return nil
}

func (m *HWComponentInfoSetRequest) GetComponentUuid() *Uuid {
	if m != nil {
		return m.ComponentUuid
	}
	return nil
}

func (m *HWComponentInfoSetRequest) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

func (m *HWComponentInfoSetRequest) GetChanges() *ModifiableComponent {
	if m != nil {
		return m.Changes
	}
	return nil
}

type HWComponentInfoSetResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=dmi.Status" json:"status,omitempty"`
	Reason               Reason   `protobuf:"varint,2,opt,name=reason,proto3,enum=dmi.Reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HWComponentInfoSetResponse) Reset()         { *m = HWComponentInfoSetResponse{} }
func (m *HWComponentInfoSetResponse) String() string { return proto.CompactTextString(m) }
func (*HWComponentInfoSetResponse) ProtoMessage()    {}
func (*HWComponentInfoSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{4}
}

func (m *HWComponentInfoSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWComponentInfoSetResponse.Unmarshal(m, b)
}
func (m *HWComponentInfoSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWComponentInfoSetResponse.Marshal(b, m, deterministic)
}
func (m *HWComponentInfoSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWComponentInfoSetResponse.Merge(m, src)
}
func (m *HWComponentInfoSetResponse) XXX_Size() int {
	return xxx_messageInfo_HWComponentInfoSetResponse.Size(m)
}
func (m *HWComponentInfoSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HWComponentInfoSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HWComponentInfoSetResponse proto.InternalMessageInfo

func (m *HWComponentInfoSetResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNDEFINED_STATUS
}

func (m *HWComponentInfoSetResponse) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_UNDEFINED_REASON
}

type StartManagingDeviceResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=dmi.Status" json:"status,omitempty"`
	Reason               Reason   `protobuf:"varint,2,opt,name=reason,proto3,enum=dmi.Reason" json:"reason,omitempty"`
	DeviceUuid           *Uuid    `protobuf:"bytes,3,opt,name=device_uuid,json=deviceUuid,proto3" json:"device_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartManagingDeviceResponse) Reset()         { *m = StartManagingDeviceResponse{} }
func (m *StartManagingDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*StartManagingDeviceResponse) ProtoMessage()    {}
func (*StartManagingDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae902e73066286d, []int{5}
}

func (m *StartManagingDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartManagingDeviceResponse.Unmarshal(m, b)
}
func (m *StartManagingDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartManagingDeviceResponse.Marshal(b, m, deterministic)
}
func (m *StartManagingDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartManagingDeviceResponse.Merge(m, src)
}
func (m *StartManagingDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_StartManagingDeviceResponse.Size(m)
}
func (m *StartManagingDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartManagingDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartManagingDeviceResponse proto.InternalMessageInfo

func (m *StartManagingDeviceResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNDEFINED_STATUS
}

func (m *StartManagingDeviceResponse) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_UNDEFINED_REASON
}

func (m *StartManagingDeviceResponse) GetDeviceUuid() *Uuid {
	if m != nil {
		return m.DeviceUuid
	}
	return nil
}

func init() {
	proto.RegisterType((*PhysicalInventoryRequest)(nil), "dmi.PhysicalInventoryRequest")
	proto.RegisterType((*PhysicalInventoryResponse)(nil), "dmi.PhysicalInventoryResponse")
	proto.RegisterType((*HWComponentInfoGetRequest)(nil), "dmi.HWComponentInfoGetRequest")
	proto.RegisterType((*HWComponentInfoSetRequest)(nil), "dmi.HWComponentInfoSetRequest")
	proto.RegisterType((*HWComponentInfoSetResponse)(nil), "dmi.HWComponentInfoSetResponse")
	proto.RegisterType((*StartManagingDeviceResponse)(nil), "dmi.StartManagingDeviceResponse")
}

func init() { proto.RegisterFile("dmi/hw_management_service.proto", fileDescriptor_eae902e73066286d) }

var fileDescriptor_eae902e73066286d = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x65, 0x13, 0xa9, 0x64, 0x62, 0x03, 0x8e, 0x3e, 0x6c, 0x56, 0xb4, 0x65, 0x45, 0x10, 0xa5,
	0xd9, 0x90, 0x3e, 0x89, 0x6f, 0x2a, 0x4d, 0xfa, 0x90, 0x22, 0xbb, 0x94, 0x80, 0x2f, 0x61, 0xb2,
	0x73, 0xb3, 0x19, 0xe8, 0xcc, 0xc4, 0x99, 0xd9, 0x94, 0xfe, 0x86, 0xfe, 0x80, 0x1f, 0xe5, 0xab,
	0xff, 0x22, 0x3b, 0xd3, 0xdd, 0xc4, 0x24, 0x2b, 0x22, 0x0a, 0x7d, 0x4b, 0xee, 0x39, 0xf7, 0x70,
	0xee, 0xde, 0x7b, 0x06, 0x1d, 0x51, 0xce, 0xa2, 0xc5, 0xf5, 0x94, 0x13, 0x41, 0x32, 0xe0, 0x20,
	0xcc, 0x54, 0x83, 0x5a, 0xb1, 0x14, 0x7a, 0x4b, 0x25, 0x8d, 0xc4, 0x4d, 0xca, 0x59, 0xf0, 0xb0,
	0x60, 0xa5, 0x92, 0x73, 0x29, 0xb4, 0xab, 0x07, 0x0f, 0x5c, 0xa3, 0xfb, 0x17, 0x9e, 0x21, 0xff,
	0xe3, 0xe2, 0x46, 0xb3, 0x94, 0x5c, 0x9d, 0x8b, 0x15, 0x08, 0x23, 0xd5, 0x4d, 0x0c, 0x9f, 0x73,
	0xd0, 0x06, 0xbf, 0x42, 0x6d, 0x0a, 0x85, 0xe2, 0x34, 0xcf, 0x19, 0xf5, 0xbd, 0x63, 0xef, 0x65,
	0x7b, 0xd0, 0xea, 0x51, 0xce, 0x7a, 0x97, 0x39, 0xa3, 0x31, 0x72, 0x68, 0xf1, 0x3b, 0xfc, 0xe2,
	0xa1, 0xee, 0x1e, 0x21, 0xbd, 0x94, 0x42, 0x03, 0x7e, 0x8e, 0x0e, 0xb4, 0x21, 0x26, 0xd7, 0x56,
	0xa4, 0x33, 0x68, 0x5b, 0x91, 0xc4, 0x96, 0xe2, 0x5b, 0xa8, 0x20, 0x29, 0x20, 0x5a, 0x0a, 0xbf,
	0xb1, 0x41, 0x8a, 0x6d, 0x29, 0xbe, 0x85, 0xf0, 0x6b, 0xd4, 0x62, 0xa5, 0xbc, 0xdf, 0xb4, 0x8e,
	0x0e, 0x2d, 0x6f, 0x44, 0x14, 0xbd, 0x26, 0x0a, 0xe2, 0x35, 0x1e, 0x7e, 0xf3, 0x50, 0x77, 0x34,
	0x79, 0x2f, 0xf9, 0x52, 0x0a, 0x10, 0xe6, 0x5c, 0xcc, 0xe5, 0x10, 0xcc, 0x5f, 0x8c, 0x87, 0xfb,
	0xa8, 0x93, 0x96, 0x32, 0x8e, 0xde, 0xd8, 0xa6, 0x1f, 0x56, 0x04, 0xdb, 0xf1, 0x62, 0xb3, 0x43,
	0x10, 0x0e, 0xd6, 0x6d, 0x6b, 0x83, 0x76, 0x41, 0x38, 0x84, 0xdf, 0x77, 0x2d, 0x26, 0x77, 0xcb,
	0x22, 0x1e, 0xa0, 0xfb, 0xe9, 0x82, 0x88, 0x0c, 0xb4, 0x7f, 0xcf, 0x2a, 0xfa, 0x56, 0x71, 0x2c,
	0x29, 0x9b, 0x33, 0x32, 0xbb, 0x82, 0xca, 0x7d, 0x5c, 0x12, 0xc3, 0x39, 0x0a, 0xf6, 0x4d, 0xf5,
	0xaf, 0xcf, 0x21, 0xfc, 0xea, 0xa1, 0x27, 0x89, 0x21, 0xca, 0x8c, 0x8b, 0x18, 0x30, 0x91, 0x7d,
	0xb0, 0x1f, 0xe4, 0x3f, 0x1c, 0xde, 0xd6, 0x2a, 0x9a, 0xbf, 0x59, 0xc5, 0xe0, 0x47, 0x03, 0x75,
	0x2f, 0x88, 0x61, 0x2b, 0x18, 0x4d, 0xc6, 0x55, 0x3e, 0x13, 0x17, 0x4f, 0x3c, 0x42, 0x8f, 0xf6,
	0x58, 0xc6, 0xbf, 0x9e, 0x71, 0x70, 0x5c, 0x3a, 0xad, 0x9b, 0xad, 0xef, 0xe1, 0x09, 0x7a, 0x3c,
	0x04, 0xb3, 0x13, 0x3b, 0xfc, 0xd4, 0xf6, 0xd6, 0xe5, 0x3a, 0x78, 0x56, 0x07, 0x57, 0xc2, 0x67,
	0x08, 0x0f, 0xc1, 0x6c, 0x6d, 0x10, 0xbb, 0xbe, 0xda, 0x40, 0x05, 0x1d, 0x8b, 0x57, 0x68, 0xdf,
	0xc3, 0x97, 0x08, 0x27, 0x7f, 0xa8, 0xb3, 0xbe, 0xfa, 0xe0, 0xa8, 0x16, 0x77, 0x06, 0xdf, 0xbd,
	0xfd, 0xf4, 0x26, 0x63, 0x66, 0x91, 0xcf, 0x7a, 0xa9, 0xe4, 0x91, 0x5c, 0x82, 0x48, 0xa5, 0xa2,
	0x91, 0xdb, 0xc0, 0xc9, 0xfa, 0x45, 0x3c, 0x61, 0xc2, 0x80, 0x9a, 0x93, 0x14, 0xa2, 0xd5, 0x69,
	0x94, 0xc9, 0x88, 0x72, 0x36, 0x3b, 0xb0, 0x0f, 0xdf, 0xe9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x23, 0x94, 0xd9, 0x06, 0x41, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NativeHWManagementServiceClient is the client API for NativeHWManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NativeHWManagementServiceClient interface {
	// Initializes context for a device and sets up required states
	StartManagingDevice(ctx context.Context, in *Hardware, opts ...grpc.CallOption) (NativeHWManagementService_StartManagingDeviceClient, error)
	// Get the HW inventory details of the Device
	GetPhysicalInventory(ctx context.Context, in *PhysicalInventoryRequest, opts ...grpc.CallOption) (NativeHWManagementService_GetPhysicalInventoryClient, error)
	// Get the details of a particular HW component
	GetHWComponentInfo(ctx context.Context, in *HWComponentInfoGetRequest, opts ...grpc.CallOption) (NativeHWManagementService_GetHWComponentInfoClient, error)
	// Sets the permissible attributes of a HW component
	SetHWComponentInfo(ctx context.Context, in *HWComponentInfoSetRequest, opts ...grpc.CallOption) (*HWComponentInfoSetResponse, error)
}

type nativeHWManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewNativeHWManagementServiceClient(cc *grpc.ClientConn) NativeHWManagementServiceClient {
	return &nativeHWManagementServiceClient{cc}
}

func (c *nativeHWManagementServiceClient) StartManagingDevice(ctx context.Context, in *Hardware, opts ...grpc.CallOption) (NativeHWManagementService_StartManagingDeviceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NativeHWManagementService_serviceDesc.Streams[0], "/dmi.NativeHWManagementService/StartManagingDevice", opts...)
	if err != nil {
		return nil, err
	}
	x := &nativeHWManagementServiceStartManagingDeviceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NativeHWManagementService_StartManagingDeviceClient interface {
	Recv() (*StartManagingDeviceResponse, error)
	grpc.ClientStream
}

type nativeHWManagementServiceStartManagingDeviceClient struct {
	grpc.ClientStream
}

func (x *nativeHWManagementServiceStartManagingDeviceClient) Recv() (*StartManagingDeviceResponse, error) {
	m := new(StartManagingDeviceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nativeHWManagementServiceClient) GetPhysicalInventory(ctx context.Context, in *PhysicalInventoryRequest, opts ...grpc.CallOption) (NativeHWManagementService_GetPhysicalInventoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NativeHWManagementService_serviceDesc.Streams[1], "/dmi.NativeHWManagementService/GetPhysicalInventory", opts...)
	if err != nil {
		return nil, err
	}
	x := &nativeHWManagementServiceGetPhysicalInventoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NativeHWManagementService_GetPhysicalInventoryClient interface {
	Recv() (*PhysicalInventoryResponse, error)
	grpc.ClientStream
}

type nativeHWManagementServiceGetPhysicalInventoryClient struct {
	grpc.ClientStream
}

func (x *nativeHWManagementServiceGetPhysicalInventoryClient) Recv() (*PhysicalInventoryResponse, error) {
	m := new(PhysicalInventoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nativeHWManagementServiceClient) GetHWComponentInfo(ctx context.Context, in *HWComponentInfoGetRequest, opts ...grpc.CallOption) (NativeHWManagementService_GetHWComponentInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NativeHWManagementService_serviceDesc.Streams[2], "/dmi.NativeHWManagementService/GetHWComponentInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &nativeHWManagementServiceGetHWComponentInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NativeHWManagementService_GetHWComponentInfoClient interface {
	Recv() (*Component, error)
	grpc.ClientStream
}

type nativeHWManagementServiceGetHWComponentInfoClient struct {
	grpc.ClientStream
}

func (x *nativeHWManagementServiceGetHWComponentInfoClient) Recv() (*Component, error) {
	m := new(Component)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nativeHWManagementServiceClient) SetHWComponentInfo(ctx context.Context, in *HWComponentInfoSetRequest, opts ...grpc.CallOption) (*HWComponentInfoSetResponse, error) {
	out := new(HWComponentInfoSetResponse)
	err := c.cc.Invoke(ctx, "/dmi.NativeHWManagementService/SetHWComponentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NativeHWManagementServiceServer is the server API for NativeHWManagementService service.
type NativeHWManagementServiceServer interface {
	// Initializes context for a device and sets up required states
	StartManagingDevice(*Hardware, NativeHWManagementService_StartManagingDeviceServer) error
	// Get the HW inventory details of the Device
	GetPhysicalInventory(*PhysicalInventoryRequest, NativeHWManagementService_GetPhysicalInventoryServer) error
	// Get the details of a particular HW component
	GetHWComponentInfo(*HWComponentInfoGetRequest, NativeHWManagementService_GetHWComponentInfoServer) error
	// Sets the permissible attributes of a HW component
	SetHWComponentInfo(context.Context, *HWComponentInfoSetRequest) (*HWComponentInfoSetResponse, error)
}

func RegisterNativeHWManagementServiceServer(s *grpc.Server, srv NativeHWManagementServiceServer) {
	s.RegisterService(&_NativeHWManagementService_serviceDesc, srv)
}

func _NativeHWManagementService_StartManagingDevice_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Hardware)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NativeHWManagementServiceServer).StartManagingDevice(m, &nativeHWManagementServiceStartManagingDeviceServer{stream})
}

type NativeHWManagementService_StartManagingDeviceServer interface {
	Send(*StartManagingDeviceResponse) error
	grpc.ServerStream
}

type nativeHWManagementServiceStartManagingDeviceServer struct {
	grpc.ServerStream
}

func (x *nativeHWManagementServiceStartManagingDeviceServer) Send(m *StartManagingDeviceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NativeHWManagementService_GetPhysicalInventory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PhysicalInventoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NativeHWManagementServiceServer).GetPhysicalInventory(m, &nativeHWManagementServiceGetPhysicalInventoryServer{stream})
}

type NativeHWManagementService_GetPhysicalInventoryServer interface {
	Send(*PhysicalInventoryResponse) error
	grpc.ServerStream
}

type nativeHWManagementServiceGetPhysicalInventoryServer struct {
	grpc.ServerStream
}

func (x *nativeHWManagementServiceGetPhysicalInventoryServer) Send(m *PhysicalInventoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NativeHWManagementService_GetHWComponentInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HWComponentInfoGetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NativeHWManagementServiceServer).GetHWComponentInfo(m, &nativeHWManagementServiceGetHWComponentInfoServer{stream})
}

type NativeHWManagementService_GetHWComponentInfoServer interface {
	Send(*Component) error
	grpc.ServerStream
}

type nativeHWManagementServiceGetHWComponentInfoServer struct {
	grpc.ServerStream
}

func (x *nativeHWManagementServiceGetHWComponentInfoServer) Send(m *Component) error {
	return x.ServerStream.SendMsg(m)
}

func _NativeHWManagementService_SetHWComponentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HWComponentInfoSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NativeHWManagementServiceServer).SetHWComponentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dmi.NativeHWManagementService/SetHWComponentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NativeHWManagementServiceServer).SetHWComponentInfo(ctx, req.(*HWComponentInfoSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NativeHWManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dmi.NativeHWManagementService",
	HandlerType: (*NativeHWManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHWComponentInfo",
			Handler:    _NativeHWManagementService_SetHWComponentInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartManagingDevice",
			Handler:       _NativeHWManagementService_StartManagingDevice_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPhysicalInventory",
			Handler:       _NativeHWManagementService_GetPhysicalInventory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetHWComponentInfo",
			Handler:       _NativeHWManagementService_GetHWComponentInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dmi/hw_management_service.proto",
}