// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: notes.proto

package schema

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

type AddNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *AddNoteRequest) Reset() {
	*x = AddNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteRequest) ProtoMessage() {}

func (x *AddNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteRequest.ProtoReflect.Descriptor instead.
func (*AddNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{0}
}

func (x *AddNoteRequest) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type GetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{1}
}

func (x *GetNoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetNotesRequest) Reset() {
	*x = GetNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesRequest) ProtoMessage() {}

func (x *GetNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesRequest.ProtoReflect.Descriptor instead.
func (*GetNotesRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNoteRequest) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type DeleteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteNotesRequest) Reset() {
	*x = DeleteNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotesRequest) ProtoMessage() {}

func (x *DeleteNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotesRequest.ProtoReflect.Descriptor instead.
func (*DeleteNotesRequest) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNotesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type AddNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddNoteResponse) Reset() {
	*x = AddNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNoteResponse) ProtoMessage() {}

func (x *AddNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNoteResponse.ProtoReflect.Descriptor instead.
func (*AddNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{6}
}

type GetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *GetNoteResponse) Reset() {
	*x = GetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteResponse) ProtoMessage() {}

func (x *GetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteResponse.ProtoReflect.Descriptor instead.
func (*GetNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{7}
}

func (x *GetNoteResponse) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type UpdateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNoteResponse) Reset() {
	*x = UpdateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteResponse) ProtoMessage() {}

func (x *UpdateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{8}
}

type GetNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GetNotesResponse) Reset() {
	*x = GetNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesResponse) ProtoMessage() {}

func (x *GetNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesResponse.ProtoReflect.Descriptor instead.
func (*GetNotesResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{9}
}

func (x *GetNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type DeleteNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNoteResponse) Reset() {
	*x = DeleteNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteResponse) ProtoMessage() {}

func (x *DeleteNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{10}
}

type DeleteNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNotesResponse) Reset() {
	*x = DeleteNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotesResponse) ProtoMessage() {}

func (x *DeleteNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotesResponse.ProtoReflect.Descriptor instead.
func (*DeleteNotesResponse) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{11}
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data      string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_notes_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_notes_proto_rawDescGZIP(), []int{12}
}

func (x *Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Note) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Note) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_notes_proto protoreflect.FileDescriptor

var file_notes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0x3c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e,
	0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x0a, 0x04, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x94, 0x04, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x50, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x20,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x23, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_notes_proto_rawDescOnce sync.Once
	file_notes_proto_rawDescData = file_notes_proto_rawDesc
)

func file_notes_proto_rawDescGZIP() []byte {
	file_notes_proto_rawDescOnce.Do(func() {
		file_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_proto_rawDescData)
	})
	return file_notes_proto_rawDescData
}

var file_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_notes_proto_goTypes = []interface{}{
	(*AddNoteRequest)(nil),      // 0: sharenote.schema.AddNoteRequest
	(*GetNoteRequest)(nil),      // 1: sharenote.schema.GetNoteRequest
	(*GetNotesRequest)(nil),     // 2: sharenote.schema.GetNotesRequest
	(*UpdateNoteRequest)(nil),   // 3: sharenote.schema.UpdateNoteRequest
	(*DeleteNoteRequest)(nil),   // 4: sharenote.schema.DeleteNoteRequest
	(*DeleteNotesRequest)(nil),  // 5: sharenote.schema.DeleteNotesRequest
	(*AddNoteResponse)(nil),     // 6: sharenote.schema.AddNoteResponse
	(*GetNoteResponse)(nil),     // 7: sharenote.schema.GetNoteResponse
	(*UpdateNoteResponse)(nil),  // 8: sharenote.schema.UpdateNoteResponse
	(*GetNotesResponse)(nil),    // 9: sharenote.schema.GetNotesResponse
	(*DeleteNoteResponse)(nil),  // 10: sharenote.schema.DeleteNoteResponse
	(*DeleteNotesResponse)(nil), // 11: sharenote.schema.DeleteNotesResponse
	(*Note)(nil),                // 12: sharenote.schema.Note
}
var file_notes_proto_depIdxs = []int32{
	12, // 0: sharenote.schema.AddNoteRequest.note:type_name -> sharenote.schema.Note
	12, // 1: sharenote.schema.UpdateNoteRequest.note:type_name -> sharenote.schema.Note
	12, // 2: sharenote.schema.GetNoteResponse.note:type_name -> sharenote.schema.Note
	12, // 3: sharenote.schema.GetNotesResponse.notes:type_name -> sharenote.schema.Note
	0,  // 4: sharenote.schema.Notes.AddNote:input_type -> sharenote.schema.AddNoteRequest
	1,  // 5: sharenote.schema.Notes.GetNote:input_type -> sharenote.schema.GetNoteRequest
	2,  // 6: sharenote.schema.Notes.GetNotes:input_type -> sharenote.schema.GetNotesRequest
	4,  // 7: sharenote.schema.Notes.DeleteNote:input_type -> sharenote.schema.DeleteNoteRequest
	5,  // 8: sharenote.schema.Notes.DeleteNotes:input_type -> sharenote.schema.DeleteNotesRequest
	3,  // 9: sharenote.schema.Notes.UpdateNote:input_type -> sharenote.schema.UpdateNoteRequest
	6,  // 10: sharenote.schema.Notes.AddNote:output_type -> sharenote.schema.AddNoteResponse
	7,  // 11: sharenote.schema.Notes.GetNote:output_type -> sharenote.schema.GetNoteResponse
	9,  // 12: sharenote.schema.Notes.GetNotes:output_type -> sharenote.schema.GetNotesResponse
	10, // 13: sharenote.schema.Notes.DeleteNote:output_type -> sharenote.schema.DeleteNoteResponse
	11, // 14: sharenote.schema.Notes.DeleteNotes:output_type -> sharenote.schema.DeleteNotesResponse
	8,  // 15: sharenote.schema.Notes.UpdateNote:output_type -> sharenote.schema.UpdateNoteResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_notes_proto_init() }
func file_notes_proto_init() {
	if File_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteRequest); i {
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
		file_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteRequest); i {
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
		file_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesRequest); i {
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
		file_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteRequest); i {
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
		file_notes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteRequest); i {
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
		file_notes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotesRequest); i {
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
		file_notes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNoteResponse); i {
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
		file_notes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteResponse); i {
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
		file_notes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteResponse); i {
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
		file_notes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesResponse); i {
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
		file_notes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteResponse); i {
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
		file_notes_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotesResponse); i {
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
		file_notes_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
			RawDescriptor: file_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_proto_goTypes,
		DependencyIndexes: file_notes_proto_depIdxs,
		MessageInfos:      file_notes_proto_msgTypes,
	}.Build()
	File_notes_proto = out.File
	file_notes_proto_rawDesc = nil
	file_notes_proto_goTypes = nil
	file_notes_proto_depIdxs = nil
}
