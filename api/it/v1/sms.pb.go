// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: it/v1/sms.proto

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

type SmsCodeSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile"`
}

func (x *SmsCodeSendRequest) Reset() {
	*x = SmsCodeSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsCodeSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsCodeSendRequest) ProtoMessage() {}

func (x *SmsCodeSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsCodeSendRequest.ProtoReflect.Descriptor instead.
func (*SmsCodeSendRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SmsCodeSendRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SmsCodeSendRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SmsCodeSendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
}

func (x *SmsCodeSendReply) Reset() {
	*x = SmsCodeSendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsCodeSendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsCodeSendReply) ProtoMessage() {}

func (x *SmsCodeSendReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsCodeSendReply.ProtoReflect.Descriptor instead.
func (*SmsCodeSendReply) Descriptor() ([]byte, []int) {
	return file_it_v1_sms_proto_rawDescGZIP(), []int{1}
}

func (x *SmsCodeSendReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_it_v1_sms_proto protoreflect.FileDescriptor

var file_it_v1_sms_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x2a, 0x0a,
	0x10, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x5e, 0x0a, 0x03, 0x53, 0x4d, 0x53,
	0x12, 0x57, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_it_v1_sms_proto_rawDescOnce sync.Once
	file_it_v1_sms_proto_rawDescData = file_it_v1_sms_proto_rawDesc
)

func file_it_v1_sms_proto_rawDescGZIP() []byte {
	file_it_v1_sms_proto_rawDescOnce.Do(func() {
		file_it_v1_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_it_v1_sms_proto_rawDescData)
	})
	return file_it_v1_sms_proto_rawDescData
}

var file_it_v1_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_it_v1_sms_proto_goTypes = []interface{}{
	(*SmsCodeSendRequest)(nil), // 0: it.v1.SmsCodeSendRequest
	(*SmsCodeSendReply)(nil),   // 1: it.v1.SmsCodeSendReply
}
var file_it_v1_sms_proto_depIdxs = []int32{
	0, // 0: it.v1.SMS.SendSmsCode:input_type -> it.v1.SmsCodeSendRequest
	1, // 1: it.v1.SMS.SendSmsCode:output_type -> it.v1.SmsCodeSendReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_it_v1_sms_proto_init() }
func file_it_v1_sms_proto_init() {
	if File_it_v1_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_it_v1_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsCodeSendRequest); i {
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
		file_it_v1_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsCodeSendReply); i {
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
			RawDescriptor: file_it_v1_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_it_v1_sms_proto_goTypes,
		DependencyIndexes: file_it_v1_sms_proto_depIdxs,
		MessageInfos:      file_it_v1_sms_proto_msgTypes,
	}.Build()
	File_it_v1_sms_proto = out.File
	file_it_v1_sms_proto_rawDesc = nil
	file_it_v1_sms_proto_goTypes = nil
	file_it_v1_sms_proto_depIdxs = nil
}
