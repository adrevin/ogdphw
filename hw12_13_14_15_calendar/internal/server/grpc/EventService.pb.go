// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: EventService.proto

package internalgrpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type NewEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Time     *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Duration int64                `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	UserId   []byte               `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewEventRequest) Reset() {
	*x = NewEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEventRequest) ProtoMessage() {}

func (x *NewEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEventRequest.ProtoReflect.Descriptor instead.
func (*NewEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{0}
}

func (x *NewEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewEventRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *NewEventRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *NewEventRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

type ChangeEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []byte               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Time     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Duration int64                `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	UserId   []byte               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ChangeEventRequest) Reset() {
	*x = ChangeEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeEventRequest) ProtoMessage() {}

func (x *ChangeEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeEventRequest.ProtoReflect.Descriptor instead.
func (*ChangeEventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeEventRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChangeEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChangeEventRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *ChangeEventRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *ChangeEventRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []byte               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Time     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Duration int64                `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	UserId   []byte               `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{2}
}

func (x *EventResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EventResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EventResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *EventResponse) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventResponse `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{3}
}

func (x *EventsResponse) GetEvents() []*EventResponse {
	if x != nil {
		return x.Events
	}
	return nil
}

type EventIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventIdResponse) Reset() {
	*x = EventIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventIdResponse) ProtoMessage() {}

func (x *EventIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventIdResponse.ProtoReflect.Descriptor instead.
func (*EventIdResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{4}
}

func (x *EventIdResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type EventIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventIdRequest) Reset() {
	*x = EventIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventIdRequest) ProtoMessage() {}

func (x *EventIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventIdRequest.ProtoReflect.Descriptor instead.
func (*EventIdRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{5}
}

func (x *EventIdRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type TimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TimeRequest) Reset() {
	*x = TimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRequest) ProtoMessage() {}

func (x *TimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRequest.ProtoReflect.Descriptor instead.
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{6}
}

func (x *TimeRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_EventService_proto protoreflect.FileDescriptor

var file_EventService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0f, 0x4e, 0x65,
	0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0b,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xfa, 0x02, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x44, 0x61,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x57, 0x65, 0x65, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x73,
	0x12, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x0a, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EventService_proto_rawDescOnce sync.Once
	file_EventService_proto_rawDescData = file_EventService_proto_rawDesc
)

func file_EventService_proto_rawDescGZIP() []byte {
	file_EventService_proto_rawDescOnce.Do(func() {
		file_EventService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EventService_proto_rawDescData)
	})
	return file_EventService_proto_rawDescData
}

var file_EventService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_EventService_proto_goTypes = []interface{}{
	(*NewEventRequest)(nil),     // 0: event.NewEventRequest
	(*ChangeEventRequest)(nil),  // 1: event.ChangeEventRequest
	(*EventResponse)(nil),       // 2: event.EventResponse
	(*EventsResponse)(nil),      // 3: event.EventsResponse
	(*EventIdResponse)(nil),     // 4: event.EventIdResponse
	(*EventIdRequest)(nil),      // 5: event.EventIdRequest
	(*TimeRequest)(nil),         // 6: event.TimeRequest
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_EventService_proto_depIdxs = []int32{
	7,  // 0: event.NewEventRequest.time:type_name -> google.protobuf.Timestamp
	7,  // 1: event.ChangeEventRequest.time:type_name -> google.protobuf.Timestamp
	7,  // 2: event.EventResponse.time:type_name -> google.protobuf.Timestamp
	2,  // 3: event.EventsResponse.Events:type_name -> event.EventResponse
	7,  // 4: event.TimeRequest.time:type_name -> google.protobuf.Timestamp
	0,  // 5: event.Evens.CreateEvent:input_type -> event.NewEventRequest
	1,  // 6: event.Evens.UpdateEvent:input_type -> event.ChangeEventRequest
	5,  // 7: event.Evens.DeleteEvent:input_type -> event.EventIdRequest
	6,  // 8: event.Evens.DayEvens:input_type -> event.TimeRequest
	6,  // 9: event.Evens.WeekEvens:input_type -> event.TimeRequest
	6,  // 10: event.Evens.MonthEvens:input_type -> event.TimeRequest
	4,  // 11: event.Evens.CreateEvent:output_type -> event.EventIdResponse
	8,  // 12: event.Evens.UpdateEvent:output_type -> google.protobuf.Empty
	8,  // 13: event.Evens.DeleteEvent:output_type -> google.protobuf.Empty
	3,  // 14: event.Evens.DayEvens:output_type -> event.EventsResponse
	3,  // 15: event.Evens.WeekEvens:output_type -> event.EventsResponse
	3,  // 16: event.Evens.MonthEvens:output_type -> event.EventsResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_EventService_proto_init() }
func file_EventService_proto_init() {
	if File_EventService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EventService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEventRequest); i {
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
		file_EventService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeEventRequest); i {
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
		file_EventService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_EventService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
		file_EventService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventIdResponse); i {
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
		file_EventService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventIdRequest); i {
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
		file_EventService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRequest); i {
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
			RawDescriptor: file_EventService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_EventService_proto_goTypes,
		DependencyIndexes: file_EventService_proto_depIdxs,
		MessageInfos:      file_EventService_proto_msgTypes,
	}.Build()
	File_EventService_proto = out.File
	file_EventService_proto_rawDesc = nil
	file_EventService_proto_goTypes = nil
	file_EventService_proto_depIdxs = nil
}
