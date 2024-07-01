// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: fog.proto

package fogpb

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

type SensorType int32

const (
	SensorType_UNKNOWN   SensorType = 0
	SensorType_VELOCITY  SensorType = 1
	SensorType_GYROSCOPE SensorType = 2
)

// Enum value maps for SensorType.
var (
	SensorType_name = map[int32]string{
		0: "UNKNOWN",
		1: "VELOCITY",
		2: "GYROSCOPE",
	}
	SensorType_value = map[string]int32{
		"UNKNOWN":   0,
		"VELOCITY":  1,
		"GYROSCOPE": 2,
	}
)

func (x SensorType) Enum() *SensorType {
	p := new(SensorType)
	*p = x
	return p
}

func (x SensorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SensorType) Descriptor() protoreflect.EnumDescriptor {
	return file_fog_proto_enumTypes[0].Descriptor()
}

func (SensorType) Type() protoreflect.EnumType {
	return &file_fog_proto_enumTypes[0]
}

func (x SensorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SensorType.Descriptor instead.
func (SensorType) EnumDescriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{0}
}

type SensorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorId  string     `protobuf:"bytes,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	Type      SensorType `protobuf:"varint,2,opt,name=type,proto3,enum=fogpb.SensorType" json:"type,omitempty"`
	Value     string     `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp string     `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SensorData) Reset() {
	*x = SensorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorData) ProtoMessage() {}

func (x *SensorData) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorData.ProtoReflect.Descriptor instead.
func (*SensorData) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{0}
}

func (x *SensorData) GetSensorId() string {
	if x != nil {
		return x.SensorId
	}
	return ""
}

func (x *SensorData) GetType() SensorType {
	if x != nil {
		return x.Type
	}
	return SensorType_UNKNOWN
}

func (x *SensorData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SensorData) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type StreamDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StreamDataRequest) Reset() {
	*x = StreamDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDataRequest) ProtoMessage() {}

func (x *StreamDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDataRequest.ProtoReflect.Descriptor instead.
func (*StreamDataRequest) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{1}
}

func (x *StreamDataRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StreamDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SensorData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamDataResponse) Reset() {
	*x = StreamDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDataResponse) ProtoMessage() {}

func (x *StreamDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDataResponse.ProtoReflect.Descriptor instead.
func (*StreamDataResponse) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{2}
}

func (x *StreamDataResponse) GetData() *SensorData {
	if x != nil {
		return x.Data
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{3}
}

func (x *Position) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type UpdatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *UpdatePositionRequest) Reset() {
	*x = UpdatePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionRequest) ProtoMessage() {}

func (x *UpdatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePositionRequest) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePositionRequest) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type UpdatePositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePositionResponse) Reset() {
	*x = UpdatePositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionResponse) ProtoMessage() {}

func (x *UpdatePositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePositionResponse) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{5}
}

type ProcessDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SensorData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProcessDataRequest) Reset() {
	*x = ProcessDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDataRequest) ProtoMessage() {}

func (x *ProcessDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDataRequest.ProtoReflect.Descriptor instead.
func (*ProcessDataRequest) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{6}
}

func (x *ProcessDataRequest) GetData() *SensorData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProcessDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProcessDataResponse) Reset() {
	*x = ProcessDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDataResponse) ProtoMessage() {}

func (x *ProcessDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDataResponse.ProtoReflect.Descriptor instead.
func (*ProcessDataResponse) Descriptor() ([]byte, []int) {
	return file_fog_proto_rawDescGZIP(), []int{7}
}

var File_fog_proto protoreflect.FileDescriptor

var file_fog_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6f, 0x67,
	0x70, 0x62, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x66,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x7a, 0x22, 0x44, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3b, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x15,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x36, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x4c, 0x4f, 0x43, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x47, 0x59, 0x52, 0x4f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x02, 0x32, 0x54, 0x0a,
	0x0d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x66,
	0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x32, 0x5c, 0x0a, 0x0b, 0x45, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x54, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x2e, 0x66, 0x6f, 0x67, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x6f,
	0x67, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x66, 0x6f,
	0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fog_proto_rawDescOnce sync.Once
	file_fog_proto_rawDescData = file_fog_proto_rawDesc
)

func file_fog_proto_rawDescGZIP() []byte {
	file_fog_proto_rawDescOnce.Do(func() {
		file_fog_proto_rawDescData = protoimpl.X.CompressGZIP(file_fog_proto_rawDescData)
	})
	return file_fog_proto_rawDescData
}

var file_fog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fog_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_fog_proto_goTypes = []interface{}{
	(SensorType)(0),                // 0: fogpb.SensorType
	(*SensorData)(nil),             // 1: fogpb.SensorData
	(*StreamDataRequest)(nil),      // 2: fogpb.StreamDataRequest
	(*StreamDataResponse)(nil),     // 3: fogpb.StreamDataResponse
	(*Position)(nil),               // 4: fogpb.Position
	(*UpdatePositionRequest)(nil),  // 5: fogpb.UpdatePositionRequest
	(*UpdatePositionResponse)(nil), // 6: fogpb.UpdatePositionResponse
	(*ProcessDataRequest)(nil),     // 7: fogpb.ProcessDataRequest
	(*ProcessDataResponse)(nil),    // 8: fogpb.ProcessDataResponse
}
var file_fog_proto_depIdxs = []int32{
	0, // 0: fogpb.SensorData.type:type_name -> fogpb.SensorType
	1, // 1: fogpb.StreamDataResponse.data:type_name -> fogpb.SensorData
	4, // 2: fogpb.UpdatePositionRequest.position:type_name -> fogpb.Position
	1, // 3: fogpb.ProcessDataRequest.data:type_name -> fogpb.SensorData
	2, // 4: fogpb.SensorService.StreamData:input_type -> fogpb.StreamDataRequest
	5, // 5: fogpb.EdgeService.UpdatePosition:input_type -> fogpb.UpdatePositionRequest
	7, // 6: fogpb.CloudService.ProcessData:input_type -> fogpb.ProcessDataRequest
	3, // 7: fogpb.SensorService.StreamData:output_type -> fogpb.StreamDataResponse
	6, // 8: fogpb.EdgeService.UpdatePosition:output_type -> fogpb.UpdatePositionResponse
	8, // 9: fogpb.CloudService.ProcessData:output_type -> fogpb.ProcessDataResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_fog_proto_init() }
func file_fog_proto_init() {
	if File_fog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorData); i {
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
		file_fog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDataRequest); i {
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
		file_fog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDataResponse); i {
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
		file_fog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_fog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePositionRequest); i {
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
		file_fog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePositionResponse); i {
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
		file_fog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDataRequest); i {
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
		file_fog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDataResponse); i {
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
			RawDescriptor: file_fog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_fog_proto_goTypes,
		DependencyIndexes: file_fog_proto_depIdxs,
		EnumInfos:         file_fog_proto_enumTypes,
		MessageInfos:      file_fog_proto_msgTypes,
	}.Build()
	File_fog_proto = out.File
	file_fog_proto_rawDesc = nil
	file_fog_proto_goTypes = nil
	file_fog_proto_depIdxs = nil
}
