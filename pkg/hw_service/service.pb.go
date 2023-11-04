// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: api/service.proto

package hw_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [Request messages] ---
type PostRequestWithEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Post `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *PostRequestWithEntity) Reset() {
	*x = PostRequestWithEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequestWithEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequestWithEntity) ProtoMessage() {}

func (x *PostRequestWithEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequestWithEntity.ProtoReflect.Descriptor instead.
func (*PostRequestWithEntity) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *PostRequestWithEntity) GetEntity() *Post {
	if x != nil {
		return x.Entity
	}
	return nil
}

type PostRequestWithId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PostRequestWithId) Reset() {
	*x = PostRequestWithId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequestWithId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequestWithId) ProtoMessage() {}

func (x *PostRequestWithId) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequestWithId.ProtoReflect.Descriptor instead.
func (*PostRequestWithId) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *PostRequestWithId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommentRequestWithEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Entity *Comment `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CommentRequestWithEntity) Reset() {
	*x = CommentRequestWithEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequestWithEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequestWithEntity) ProtoMessage() {}

func (x *CommentRequestWithEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequestWithEntity.ProtoReflect.Descriptor instead.
func (*CommentRequestWithEntity) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *CommentRequestWithEntity) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CommentRequestWithEntity) GetEntity() *Comment {
	if x != nil {
		return x.Entity
	}
	return nil
}

type CommentRequestWithId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommentRequestWithId) Reset() {
	*x = CommentRequestWithId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequestWithId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequestWithId) ProtoMessage() {}

func (x *CommentRequestWithId) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequestWithId.ProtoReflect.Descriptor instead.
func (*CommentRequestWithId) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{3}
}

func (x *CommentRequestWithId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// [Response messages] ---
type PostResponseWithEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Post `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *PostResponseWithEntity) Reset() {
	*x = PostResponseWithEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponseWithEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponseWithEntity) ProtoMessage() {}

func (x *PostResponseWithEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponseWithEntity.ProtoReflect.Descriptor instead.
func (*PostResponseWithEntity) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{4}
}

func (x *PostResponseWithEntity) GetEntity() *Post {
	if x != nil {
		return x.Entity
	}
	return nil
}

type CommentResponseWithEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Comment `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CommentResponseWithEntity) Reset() {
	*x = CommentResponseWithEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponseWithEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponseWithEntity) ProtoMessage() {}

func (x *CommentResponseWithEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponseWithEntity.ProtoReflect.Descriptor instead.
func (*CommentResponseWithEntity) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{5}
}

func (x *CommentResponseWithEntity) GetEntity() *Comment {
	if x != nil {
		return x.Entity
	}
	return nil
}

// [Entities] ---
type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Heading    string     `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	Text       string     `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	LikesCount uint64     `protobuf:"varint,4,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
	Comments   []*Comment `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{6}
}

func (x *Post) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetHeading() string {
	if x != nil {
		return x.Heading
	}
	return ""
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetLikesCount() uint64 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text       string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	LikesCount uint64 `protobuf:"varint,4,opt,name=likes_count,json=likesCount,proto3" json:"likes_count,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{7}
}

func (x *Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetLikesCount() uint64 {
	if x != nil {
		return x.LikesCount
	}
	return 0
}

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66,
	0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48,
	0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x4e, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9e, 0x05, 0x0a, 0x0f, 0x48, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a, 0x28, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x28, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a,
	0x01, 0x2a, 0x22, 0x05, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x3a, 0x01, 0x2a, 0x1a, 0x05, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x66, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f,
	0x68, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_service_proto_rawDescOnce sync.Once
	file_api_service_proto_rawDescData = file_api_service_proto_rawDesc
)

func file_api_service_proto_rawDescGZIP() []byte {
	file_api_service_proto_rawDescOnce.Do(func() {
		file_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_proto_rawDescData)
	})
	return file_api_service_proto_rawDescData
}

var file_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_service_proto_goTypes = []interface{}{
	(*PostRequestWithEntity)(nil),     // 0: homework_service.PostRequestWithEntity
	(*PostRequestWithId)(nil),         // 1: homework_service.PostRequestWithId
	(*CommentRequestWithEntity)(nil),  // 2: homework_service.CommentRequestWithEntity
	(*CommentRequestWithId)(nil),      // 3: homework_service.CommentRequestWithId
	(*PostResponseWithEntity)(nil),    // 4: homework_service.PostResponseWithEntity
	(*CommentResponseWithEntity)(nil), // 5: homework_service.CommentResponseWithEntity
	(*Post)(nil),                      // 6: homework_service.Post
	(*Comment)(nil),                   // 7: homework_service.Comment
	(*emptypb.Empty)(nil),             // 8: google.protobuf.Empty
}
var file_api_service_proto_depIdxs = []int32{
	6,  // 0: homework_service.PostRequestWithEntity.entity:type_name -> homework_service.Post
	7,  // 1: homework_service.CommentRequestWithEntity.entity:type_name -> homework_service.Comment
	6,  // 2: homework_service.PostResponseWithEntity.entity:type_name -> homework_service.Post
	7,  // 3: homework_service.CommentResponseWithEntity.entity:type_name -> homework_service.Comment
	7,  // 4: homework_service.Post.comments:type_name -> homework_service.Comment
	1,  // 5: homework_service.HomeworkService.GetPost:input_type -> homework_service.PostRequestWithId
	0,  // 6: homework_service.HomeworkService.AddPost:input_type -> homework_service.PostRequestWithEntity
	0,  // 7: homework_service.HomeworkService.UpdatePost:input_type -> homework_service.PostRequestWithEntity
	1,  // 8: homework_service.HomeworkService.RemovePost:input_type -> homework_service.PostRequestWithId
	2,  // 9: homework_service.HomeworkService.AddComment:input_type -> homework_service.CommentRequestWithEntity
	3,  // 10: homework_service.HomeworkService.RemoveComment:input_type -> homework_service.CommentRequestWithId
	4,  // 11: homework_service.HomeworkService.GetPost:output_type -> homework_service.PostResponseWithEntity
	4,  // 12: homework_service.HomeworkService.AddPost:output_type -> homework_service.PostResponseWithEntity
	8,  // 13: homework_service.HomeworkService.UpdatePost:output_type -> google.protobuf.Empty
	8,  // 14: homework_service.HomeworkService.RemovePost:output_type -> google.protobuf.Empty
	5,  // 15: homework_service.HomeworkService.AddComment:output_type -> homework_service.CommentResponseWithEntity
	8,  // 16: homework_service.HomeworkService.RemoveComment:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequestWithEntity); i {
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
		file_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequestWithId); i {
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
		file_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequestWithEntity); i {
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
		file_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequestWithId); i {
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
		file_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponseWithEntity); i {
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
		file_api_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponseWithEntity); i {
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
		file_api_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_api_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
		MessageInfos:      file_api_service_proto_msgTypes,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
