// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: pkg/proto/docs.proto

package dpb

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

type Operation_Type int32

const (
	Operation_INSERT Operation_Type = 0
	Operation_DELETE Operation_Type = 1
)

// Enum value maps for Operation_Type.
var (
	Operation_Type_name = map[int32]string{
		0: "INSERT",
		1: "DELETE",
	}
	Operation_Type_value = map[string]int32{
		"INSERT": 0,
		"DELETE": 1,
	}
)

func (x Operation_Type) Enum() *Operation_Type {
	p := new(Operation_Type)
	*p = x
	return p
}

func (x Operation_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_docs_proto_enumTypes[0].Descriptor()
}

func (Operation_Type) Type() protoreflect.EnumType {
	return &file_pkg_proto_docs_proto_enumTypes[0]
}

func (x Operation_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation_Type.Descriptor instead.
func (Operation_Type) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{0, 0}
}

type Message_Type int32

const (
	Message_DOC_SYNC  Message_Type = 0
	Message_DOC_REQ   Message_Type = 1
	Message_SITE_ID   Message_Type = 2
	Message_JOIN      Message_Type = 3
	Message_USERS     Message_Type = 4
	Message_OPERATION Message_Type = 5
)

// Enum value maps for Message_Type.
var (
	Message_Type_name = map[int32]string{
		0: "DOC_SYNC",
		1: "DOC_REQ",
		2: "SITE_ID",
		3: "JOIN",
		4: "USERS",
		5: "OPERATION",
	}
	Message_Type_value = map[string]int32{
		"DOC_SYNC":  0,
		"DOC_REQ":   1,
		"SITE_ID":   2,
		"JOIN":      3,
		"USERS":     4,
		"OPERATION": 5,
	}
)

func (x Message_Type) Enum() *Message_Type {
	p := new(Message_Type)
	*p = x
	return p
}

func (x Message_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_docs_proto_enumTypes[1].Descriptor()
}

func (Message_Type) Type() protoreflect.EnumType {
	return &file_pkg_proto_docs_proto_enumTypes[1]
}

func (x Message_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Type.Descriptor instead.
func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{1, 0}
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType Operation_Type `protobuf:"varint,1,opt,name=operationType,proto3,enum=docs.Operation_Type" json:"operationType,omitempty"`
	Position      int32          `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	Value         string         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_docs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_docs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetOperationType() Operation_Type {
	if x != nil {
		return x.OperationType
	}
	return Operation_INSERT
}

func (x *Operation) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Operation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType Message_Type `protobuf:"varint,1,opt,name=messageType,proto3,enum=docs.Message_Type" json:"messageType,omitempty"`
	Id          string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"` // ID represents the client's UUID.
	Username    string       `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Text represents the body of the message. This is being used for joining messages, the siteID, and the list of active users.
	Text string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// Operation represents the CRDT operation.
	Operation *Operation `protobuf:"bytes,5,opt,name=operation,proto3" json:"operation,omitempty"`
	// Document represents the client's document.
	//
	//	This is not used frequently, and should be only used when necessary, due to the large size of documents.
	Document *Document `protobuf:"bytes,6,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_docs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_docs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMessageType() Message_Type {
	if x != nil {
		return x.MessageType
	}
	return Message_DOC_SYNC
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetOperation() *Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *Message) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document []*Document `protobuf:"bytes,1,rep,name=document,proto3" json:"document,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_docs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_docs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{2}
}

func (x *Document) GetDocument() []*Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Visible    bool   `protobuf:"varint,2,opt,name=visible,proto3" json:"visible,omitempty"`
	Value      string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	IdPrevious string `protobuf:"bytes,4,opt,name=idPrevious,proto3" json:"idPrevious,omitempty"`
	IdNext     string `protobuf:"bytes,5,opt,name=idNext,proto3" json:"idNext,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_docs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_docs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{3}
}

func (x *Character) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Character) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Character) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Character) GetIdPrevious() string {
	if x != nil {
		return x.IdPrevious
	}
	return ""
}

func (x *Character) GetIdNext() string {
	if x != nil {
		return x.IdNext
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_docs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_docs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_docs_proto_rawDescGZIP(), []int{4}
}

func (x *MessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pkg_proto_docs_proto protoreflect.FileDescriptor

var file_pkg_proto_docs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x22, 0x99, 0x01, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x22, 0xae, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x63, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x64, 0x6f, 0x63, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x08, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f,
	0x63, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x4f, 0x43, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x4f, 0x43, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49, 0x54, 0x45,
	0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x22, 0x36, 0x0a, 0x08, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x83, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0x44, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0d, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x15, 0x2e, 0x64, 0x6f, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x3b, 0x64, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_docs_proto_rawDescOnce sync.Once
	file_pkg_proto_docs_proto_rawDescData = file_pkg_proto_docs_proto_rawDesc
)

func file_pkg_proto_docs_proto_rawDescGZIP() []byte {
	file_pkg_proto_docs_proto_rawDescOnce.Do(func() {
		file_pkg_proto_docs_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_docs_proto_rawDescData)
	})
	return file_pkg_proto_docs_proto_rawDescData
}

var file_pkg_proto_docs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_proto_docs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_docs_proto_goTypes = []any{
	(Operation_Type)(0),     // 0: docs.Operation.Type
	(Message_Type)(0),       // 1: docs.Message.Type
	(*Operation)(nil),       // 2: docs.Operation
	(*Message)(nil),         // 3: docs.Message
	(*Document)(nil),        // 4: docs.Document
	(*Character)(nil),       // 5: docs.Character
	(*MessageResponse)(nil), // 6: docs.MessageResponse
}
var file_pkg_proto_docs_proto_depIdxs = []int32{
	0, // 0: docs.Operation.operationType:type_name -> docs.Operation.Type
	1, // 1: docs.Message.messageType:type_name -> docs.Message.Type
	2, // 2: docs.Message.operation:type_name -> docs.Operation
	4, // 3: docs.Message.document:type_name -> docs.Document
	4, // 4: docs.Document.document:type_name -> docs.Document
	3, // 5: docs.DocsService.SendMessage:input_type -> docs.Message
	6, // 6: docs.DocsService.SendMessage:output_type -> docs.MessageResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_docs_proto_init() }
func file_pkg_proto_docs_proto_init() {
	if File_pkg_proto_docs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_docs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Operation); i {
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
		file_pkg_proto_docs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_pkg_proto_docs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Document); i {
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
		file_pkg_proto_docs_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Character); i {
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
		file_pkg_proto_docs_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MessageResponse); i {
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
			RawDescriptor: file_pkg_proto_docs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_docs_proto_goTypes,
		DependencyIndexes: file_pkg_proto_docs_proto_depIdxs,
		EnumInfos:         file_pkg_proto_docs_proto_enumTypes,
		MessageInfos:      file_pkg_proto_docs_proto_msgTypes,
	}.Build()
	File_pkg_proto_docs_proto = out.File
	file_pkg_proto_docs_proto_rawDesc = nil
	file_pkg_proto_docs_proto_goTypes = nil
	file_pkg_proto_docs_proto_depIdxs = nil
}