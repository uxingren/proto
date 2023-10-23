// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: it/v1/category.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCateAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCateAllRequest) Reset() {
	*x = GetCateAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCateAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCateAllRequest) ProtoMessage() {}

func (x *GetCateAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCateAllRequest.ProtoReflect.Descriptor instead.
func (*GetCateAllRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_category_proto_rawDescGZIP(), []int{0}
}

type GetCateAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetCateAllReply_Nested `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
}

func (x *GetCateAllReply) Reset() {
	*x = GetCateAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCateAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCateAllReply) ProtoMessage() {}

func (x *GetCateAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCateAllReply.ProtoReflect.Descriptor instead.
func (*GetCateAllReply) Descriptor() ([]byte, []int) {
	return file_it_v1_category_proto_rawDescGZIP(), []int{1}
}

func (x *GetCateAllReply) GetItems() []*GetCateAllReply_Nested {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetCateAllReply_Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid"`
	Pid  uint32 `protobuf:"varint,4,opt,name=pid,proto3" json:"pid"`
}

func (x *GetCateAllReply_Child) Reset() {
	*x = GetCateAllReply_Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCateAllReply_Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCateAllReply_Child) ProtoMessage() {}

func (x *GetCateAllReply_Child) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCateAllReply_Child.ProtoReflect.Descriptor instead.
func (*GetCateAllReply_Child) Descriptor() ([]byte, []int) {
	return file_it_v1_category_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetCateAllReply_Child) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCateAllReply_Child) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCateAllReply_Child) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetCateAllReply_Child) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type GetCateAllReply_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name  string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Uuid  string                   `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid"`
	Child []*GetCateAllReply_Child `protobuf:"bytes,4,rep,name=child,proto3" json:"child"`
}

func (x *GetCateAllReply_Nested) Reset() {
	*x = GetCateAllReply_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCateAllReply_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCateAllReply_Nested) ProtoMessage() {}

func (x *GetCateAllReply_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCateAllReply_Nested.ProtoReflect.Descriptor instead.
func (*GetCateAllReply_Nested) Descriptor() ([]byte, []int) {
	return file_it_v1_category_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetCateAllReply_Nested) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCateAllReply_Nested) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCateAllReply_Nested) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetCateAllReply_Nested) GetChild() []*GetCateAllReply_Child {
	if x != nil {
		return x.Child
	}
	return nil
}

var File_it_v1_category_proto protoreflect.FileDescriptor

var file_it_v1_category_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x8f, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x51, 0x0a, 0x05, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x1a, 0x74, 0x0a, 0x06,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x32,
	0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x05, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x32, 0x65, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x59,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x6c, 0x6c,
	0x12, 0x18, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x6c, 0x6c, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_it_v1_category_proto_rawDescOnce sync.Once
	file_it_v1_category_proto_rawDescData = file_it_v1_category_proto_rawDesc
)

func file_it_v1_category_proto_rawDescGZIP() []byte {
	file_it_v1_category_proto_rawDescOnce.Do(func() {
		file_it_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_it_v1_category_proto_rawDescData)
	})
	return file_it_v1_category_proto_rawDescData
}

var file_it_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_it_v1_category_proto_goTypes = []interface{}{
	(*GetCateAllRequest)(nil),      // 0: it.v1.GetCateAllRequest
	(*GetCateAllReply)(nil),        // 1: it.v1.GetCateAllReply
	(*GetCateAllReply_Child)(nil),  // 2: it.v1.GetCateAllReply.Child
	(*GetCateAllReply_Nested)(nil), // 3: it.v1.GetCateAllReply.Nested
}
var file_it_v1_category_proto_depIdxs = []int32{
	3, // 0: it.v1.GetCateAllReply.items:type_name -> it.v1.GetCateAllReply.Nested
	2, // 1: it.v1.GetCateAllReply.Nested.child:type_name -> it.v1.GetCateAllReply.Child
	0, // 2: it.v1.Category.GetCategoryAll:input_type -> it.v1.GetCateAllRequest
	1, // 3: it.v1.Category.GetCategoryAll:output_type -> it.v1.GetCateAllReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_it_v1_category_proto_init() }
func file_it_v1_category_proto_init() {
	if File_it_v1_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_it_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCateAllRequest); i {
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
		file_it_v1_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCateAllReply); i {
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
		file_it_v1_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCateAllReply_Child); i {
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
		file_it_v1_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCateAllReply_Nested); i {
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
			RawDescriptor: file_it_v1_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_it_v1_category_proto_goTypes,
		DependencyIndexes: file_it_v1_category_proto_depIdxs,
		MessageInfos:      file_it_v1_category_proto_msgTypes,
	}.Build()
	File_it_v1_category_proto = out.File
	file_it_v1_category_proto_rawDesc = nil
	file_it_v1_category_proto_goTypes = nil
	file_it_v1_category_proto_depIdxs = nil
}