// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.11.2
// source: rpc/service.proto

package pb

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

// Error
type ErrorResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode uint32 `protobuf:"varint,1,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Detail  string `protobuf:"bytes,3,opt,name=Detail,proto3" json:"Detail,omitempty"`
}

func (x *ErrorResp) Reset() {
	*x = ErrorResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResp) ProtoMessage() {}

func (x *ErrorResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResp.ProtoReflect.Descriptor instead.
func (*ErrorResp) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResp) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *ErrorResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *ErrorResp) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

// Version
type VersionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionReq) Reset() {
	*x = VersionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionReq) ProtoMessage() {}

func (x *VersionReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionReq.ProtoReflect.Descriptor instead.
func (*VersionReq) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{1}
}

type VersionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *VersionResp) Reset() {
	*x = VersionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResp) ProtoMessage() {}

func (x *VersionResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResp.ProtoReflect.Descriptor instead.
func (*VersionResp) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *VersionResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// DocumentCreate 创建文档
type DocumentCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocID  string `protobuf:"bytes,1,opt,name=DocID,proto3" json:"DocID,omitempty"`
	DocURL string `protobuf:"bytes,2,opt,name=DocURL,proto3" json:"DocURL,omitempty"`
}

func (x *DocumentCreateReq) Reset() {
	*x = DocumentCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentCreateReq) ProtoMessage() {}

func (x *DocumentCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentCreateReq.ProtoReflect.Descriptor instead.
func (*DocumentCreateReq) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *DocumentCreateReq) GetDocID() string {
	if x != nil {
		return x.DocID
	}
	return ""
}

func (x *DocumentCreateReq) GetDocURL() string {
	if x != nil {
		return x.DocURL
	}
	return ""
}

type DocumentCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocID string `protobuf:"bytes,1,opt,name=DocID,proto3" json:"DocID,omitempty"`
}

func (x *DocumentCreateResp) Reset() {
	*x = DocumentCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentCreateResp) ProtoMessage() {}

func (x *DocumentCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentCreateResp.ProtoReflect.Descriptor instead.
func (*DocumentCreateResp) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *DocumentCreateResp) GetDocID() string {
	if x != nil {
		return x.DocID
	}
	return ""
}

// Search 文档检索
type SearchItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocID string  `protobuf:"bytes,1,opt,name=DocID,proto3" json:"DocID,omitempty"`
	Score float32 `protobuf:"fixed32,2,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *SearchItem) Reset() {
	*x = SearchItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchItem) ProtoMessage() {}

func (x *SearchItem) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchItem.ProtoReflect.Descriptor instead.
func (*SearchItem) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *SearchItem) GetDocID() string {
	if x != nil {
		return x.DocID
	}
	return ""
}

func (x *SearchItem) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID      string   `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	KeyWords []string `protobuf:"bytes,2,rep,name=KeyWords,proto3" json:"KeyWords,omitempty"`
	DocIDs   []string `protobuf:"bytes,3,rep,name=DocIDs,proto3" json:"DocIDs,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{6}
}

func (x *SearchReq) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *SearchReq) GetKeyWords() []string {
	if x != nil {
		return x.KeyWords
	}
	return nil
}

func (x *SearchReq) GetDocIDs() []string {
	if x != nil {
		return x.DocIDs
	}
	return nil
}

type SearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID  string        `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	List []*SearchItem `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *SearchResp) Reset() {
	*x = SearchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResp) ProtoMessage() {}

func (x *SearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResp.ProtoReflect.Descriptor instead.
func (*SearchResp) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{7}
}

func (x *SearchResp) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *SearchResp) GetList() []*SearchItem {
	if x != nil {
		return x.List
	}
	return nil
}

var File_rpc_service_proto protoreflect.FileDescriptor

var file_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x22, 0x55, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0x0c, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x22, 0x27, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x11, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x63, 0x55, 0x52, 0x4c, 0x22, 0x2a,
	0x0a, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x0a, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x6f, 0x63, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0x51, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x44, 0x6f, 0x63, 0x49, 0x44, 0x73, 0x22, 0x48, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xbf, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x36,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x63, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_service_proto_rawDescOnce sync.Once
	file_rpc_service_proto_rawDescData = file_rpc_service_proto_rawDesc
)

func file_rpc_service_proto_rawDescGZIP() []byte {
	file_rpc_service_proto_rawDescOnce.Do(func() {
		file_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_service_proto_rawDescData)
	})
	return file_rpc_service_proto_rawDescData
}

var file_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_service_proto_goTypes = []interface{}{
	(*ErrorResp)(nil),          // 0: retrieve.ErrorResp
	(*VersionReq)(nil),         // 1: retrieve.VersionReq
	(*VersionResp)(nil),        // 2: retrieve.VersionResp
	(*DocumentCreateReq)(nil),  // 3: retrieve.DocumentCreateReq
	(*DocumentCreateResp)(nil), // 4: retrieve.DocumentCreateResp
	(*SearchItem)(nil),         // 5: retrieve.SearchItem
	(*SearchReq)(nil),          // 6: retrieve.SearchReq
	(*SearchResp)(nil),         // 7: retrieve.SearchResp
}
var file_rpc_service_proto_depIdxs = []int32{
	5, // 0: retrieve.SearchResp.List:type_name -> retrieve.SearchItem
	1, // 1: retrieve.Retrieve.Version:input_type -> retrieve.VersionReq
	3, // 2: retrieve.Retrieve.CreateDoc:input_type -> retrieve.DocumentCreateReq
	6, // 3: retrieve.Retrieve.Search:input_type -> retrieve.SearchReq
	2, // 4: retrieve.Retrieve.Version:output_type -> retrieve.VersionResp
	4, // 5: retrieve.Retrieve.CreateDoc:output_type -> retrieve.DocumentCreateResp
	7, // 6: retrieve.Retrieve.Search:output_type -> retrieve.SearchResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_service_proto_init() }
func file_rpc_service_proto_init() {
	if File_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResp); i {
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
		file_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionReq); i {
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
		file_rpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResp); i {
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
		file_rpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentCreateReq); i {
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
		file_rpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentCreateResp); i {
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
		file_rpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchItem); i {
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
		file_rpc_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_rpc_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResp); i {
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
			RawDescriptor: file_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_service_proto_goTypes,
		DependencyIndexes: file_rpc_service_proto_depIdxs,
		MessageInfos:      file_rpc_service_proto_msgTypes,
	}.Build()
	File_rpc_service_proto = out.File
	file_rpc_service_proto_rawDesc = nil
	file_rpc_service_proto_goTypes = nil
	file_rpc_service_proto_depIdxs = nil
}
