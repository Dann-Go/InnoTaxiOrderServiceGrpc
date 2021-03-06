// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: protos/order.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	TaxiType string `protobuf:"bytes,2,opt,name=TaxiType,proto3" json:"TaxiType,omitempty"`
	FromDest string `protobuf:"bytes,3,opt,name=FromDest,proto3" json:"FromDest,omitempty"`
	ToDest   string `protobuf:"bytes,4,opt,name=ToDest,proto3" json:"ToDest,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRequest) GetTaxiType() string {
	if x != nil {
		return x.TaxiType
	}
	return ""
}

func (x *UserRequest) GetFromDest() string {
	if x != nil {
		return x.FromDest
	}
	return ""
}

func (x *UserRequest) GetToDest() string {
	if x != nil {
		return x.ToDest
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
	UserId   int32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{2}
}

func (x *Ids) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *Ids) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId   int32  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	DriverId int32  `protobuf:"varint,3,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
	FromDest string `protobuf:"bytes,4,opt,name=FromDest,proto3" json:"FromDest,omitempty"`
	ToDest   string `protobuf:"bytes,5,opt,name=ToDest,proto3" json:"ToDest,omitempty"`
	TaxiType string `protobuf:"bytes,6,opt,name=TaxiType,proto3" json:"TaxiType,omitempty"`
	Date     string `protobuf:"bytes,7,opt,name=Date,proto3" json:"Date,omitempty"`
	Status   string `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderResponse) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *OrderResponse) GetFromDest() string {
	if x != nil {
		return x.FromDest
	}
	return ""
}

func (x *OrderResponse) GetToDest() string {
	if x != nil {
		return x.ToDest
	}
	return ""
}

func (x *OrderResponse) GetTaxiType() string {
	if x != nil {
		return x.TaxiType
	}
	return ""
}

func (x *OrderResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *OrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DriverId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
}

func (x *DriverId) Reset() {
	*x = DriverId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverId) ProtoMessage() {}

func (x *DriverId) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverId.ProtoReflect.Descriptor instead.
func (*DriverId) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{4}
}

func (x *DriverId) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type FinishOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *FinishOrderResponse) Reset() {
	*x = FinishOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishOrderResponse) ProtoMessage() {}

func (x *FinishOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishOrderResponse.ProtoReflect.Descriptor instead.
func (*FinishOrderResponse) Descriptor() ([]byte, []int) {
	return file_protos_order_proto_rawDescGZIP(), []int{5}
}

func (x *FinishOrderResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_protos_order_proto protoreflect.FileDescriptor

var file_protos_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d,
	0x44, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d,
	0x44, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x6f, 0x44, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x54, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x54, 0x61, 0x78, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x13,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x90, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x04, 0x2e, 0x49,
	0x64, 0x73, 0x1a, 0x0e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x09, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x6e, 0x6e, 0x2d, 0x47, 0x6f, 0x2f, 0x49, 0x6e, 0x6e, 0x6f, 0x54, 0x61, 0x78,
	0x69, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order_proto_rawDescOnce sync.Once
	file_protos_order_proto_rawDescData = file_protos_order_proto_rawDesc
)

func file_protos_order_proto_rawDescGZIP() []byte {
	file_protos_order_proto_rawDescOnce.Do(func() {
		file_protos_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order_proto_rawDescData)
	})
	return file_protos_order_proto_rawDescData
}

var file_protos_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_order_proto_goTypes = []interface{}{
	(*UserRequest)(nil),         // 0: UserRequest
	(*UserResponse)(nil),        // 1: UserResponse
	(*Ids)(nil),                 // 2: Ids
	(*OrderResponse)(nil),       // 3: OrderResponse
	(*DriverId)(nil),            // 4: DriverId
	(*FinishOrderResponse)(nil), // 5: FinishOrderResponse
}
var file_protos_order_proto_depIdxs = []int32{
	0, // 0: OrderService.CreateOrder:input_type -> UserRequest
	2, // 1: OrderService.GetOrderById:input_type -> Ids
	4, // 2: OrderService.FinishOrder:input_type -> DriverId
	1, // 3: OrderService.CreateOrder:output_type -> UserResponse
	3, // 4: OrderService.GetOrderById:output_type -> OrderResponse
	5, // 5: OrderService.FinishOrder:output_type -> FinishOrderResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_order_proto_init() }
func file_protos_order_proto_init() {
	if File_protos_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_order_proto_goTypes,
		DependencyIndexes: file_protos_order_proto_depIdxs,
		MessageInfos:      file_protos_order_proto_msgTypes,
	}.Build()
	File_protos_order_proto = out.File
	file_protos_order_proto_rawDesc = nil
	file_protos_order_proto_goTypes = nil
	file_protos_order_proto_depIdxs = nil
}
