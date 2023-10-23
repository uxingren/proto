// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: it/v1/user.proto

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

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Uuid     string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserInfoRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetUserInfoRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserInfoRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type GetUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Uuid             string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid"`
	UserName         string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Mobile           string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`
	NickName         string `protobuf:"bytes,5,opt,name=nick_name,json=nickName,proto3" json:"nick_name"`
	Sign             string `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign"`
	Avatar           string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar"`
	Password         string `protobuf:"bytes,8,opt,name=password,proto3" json:"password"`
	Author           string `protobuf:"bytes,9,opt,name=author,proto3" json:"author"`
	ColumnUrl        string `protobuf:"bytes,10,opt,name=column_url,json=columnUrl,proto3" json:"column_url"`
	SmscodeMobile    string `protobuf:"bytes,11,opt,name=smscode_mobile,json=smscodeMobile,proto3" json:"smscode_mobile"`
	Smscode          string `protobuf:"bytes,12,opt,name=smscode,proto3" json:"smscode"`
	SmscodeExpTime   int64  `protobuf:"varint,13,opt,name=smscode_exp_time,json=smscodeExpTime,proto3" json:"smscode_exp_time"`
	SmscodeTodayTime uint32 `protobuf:"varint,14,opt,name=smscode_today_time,json=smscodeTodayTime,proto3" json:"smscode_today_time"`
}

func (x *GetUserInfoReply) Reset() {
	*x = GetUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReply) ProtoMessage() {}

func (x *GetUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReply.ProtoReflect.Descriptor instead.
func (*GetUserInfoReply) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoReply) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserInfoReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetUserInfoReply) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserInfoReply) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetUserInfoReply) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GetUserInfoReply) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetUserInfoReply) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetUserInfoReply) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetUserInfoReply) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetUserInfoReply) GetColumnUrl() string {
	if x != nil {
		return x.ColumnUrl
	}
	return ""
}

func (x *GetUserInfoReply) GetSmscodeMobile() string {
	if x != nil {
		return x.SmscodeMobile
	}
	return ""
}

func (x *GetUserInfoReply) GetSmscode() string {
	if x != nil {
		return x.Smscode
	}
	return ""
}

func (x *GetUserInfoReply) GetSmscodeExpTime() int64 {
	if x != nil {
		return x.SmscodeExpTime
	}
	return 0
}

func (x *GetUserInfoReply) GetSmscodeTodayTime() uint32 {
	if x != nil {
		return x.SmscodeTodayTime
	}
	return 0
}

type GetTopUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
}

func (x *GetTopUserRequest) Reset() {
	*x = GetTopUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopUserRequest) ProtoMessage() {}

func (x *GetTopUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopUserRequest.ProtoReflect.Descriptor instead.
func (*GetTopUserRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetTopUserRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTopUserRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetTopUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*GetTopUserReply_Info `protobuf:"bytes,1,rep,name=item,proto3" json:"item"`
}

func (x *GetTopUserReply) Reset() {
	*x = GetTopUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopUserReply) ProtoMessage() {}

func (x *GetTopUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopUserReply.ProtoReflect.Descriptor instead.
func (*GetTopUserReply) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetTopUserReply) GetItem() []*GetTopUserReply_Info {
	if x != nil {
		return x.Item
	}
	return nil
}

type SetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	NickName  string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name"`
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author"`
	Sign      string `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign"`
	Avatar    string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar"`
	Password  string `protobuf:"bytes,6,opt,name=password,proto3" json:"password"`
	LoginTime int64  `protobuf:"varint,7,opt,name=login_time,json=loginTime,proto3" json:"login_time"`
	LoginIp   string `protobuf:"bytes,8,opt,name=login_ip,json=loginIp,proto3" json:"login_ip"`
	PostTime  int64  `protobuf:"varint,9,opt,name=post_time,json=postTime,proto3" json:"post_time"`
	Mobile    string `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile"`
}

func (x *SetUserInfoRequest) Reset() {
	*x = SetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserInfoRequest) ProtoMessage() {}

func (x *SetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*SetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *SetUserInfoRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SetUserInfoRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SetUserInfoRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *SetUserInfoRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *SetUserInfoRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SetUserInfoRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SetUserInfoRequest) GetLoginTime() int64 {
	if x != nil {
		return x.LoginTime
	}
	return 0
}

func (x *SetUserInfoRequest) GetLoginIp() string {
	if x != nil {
		return x.LoginIp
	}
	return ""
}

func (x *SetUserInfoRequest) GetPostTime() int64 {
	if x != nil {
		return x.PostTime
	}
	return 0
}

func (x *SetUserInfoRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SetUserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
}

func (x *SetUserInfoReply) Reset() {
	*x = SetUserInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserInfoReply) ProtoMessage() {}

func (x *SetUserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserInfoReply.ProtoReflect.Descriptor instead.
func (*SetUserInfoReply) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *SetUserInfoReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type BindMobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Mobile  string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile"`
	Smscode string `protobuf:"bytes,3,opt,name=smscode,proto3" json:"smscode"`
}

func (x *BindMobileRequest) Reset() {
	*x = BindMobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindMobileRequest) ProtoMessage() {}

func (x *BindMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindMobileRequest.ProtoReflect.Descriptor instead.
func (*BindMobileRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *BindMobileRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BindMobileRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *BindMobileRequest) GetSmscode() string {
	if x != nil {
		return x.Smscode
	}
	return ""
}

type BindMobileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
}

func (x *BindMobileReply) Reset() {
	*x = BindMobileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindMobileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindMobileReply) ProtoMessage() {}

func (x *BindMobileReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindMobileReply.ProtoReflect.Descriptor instead.
func (*BindMobileReply) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *BindMobileReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	ClientIp string `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *AddUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddUserRequest) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type AddUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
}

func (x *AddUserReply) Reset() {
	*x = AddUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserReply) ProtoMessage() {}

func (x *AddUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserReply.ProtoReflect.Descriptor instead.
func (*AddUserReply) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *AddUserReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetTopUserReply_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	NickName  string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name"`
	Avatar    string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar"`
	ColumnUrl string `protobuf:"bytes,4,opt,name=column_url,json=columnUrl,proto3" json:"column_url"`
	PostTime  int64  `protobuf:"varint,5,opt,name=post_time,json=postTime,proto3" json:"post_time"`
}

func (x *GetTopUserReply_Info) Reset() {
	*x = GetTopUserReply_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_it_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopUserReply_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopUserReply_Info) ProtoMessage() {}

func (x *GetTopUserReply_Info) ProtoReflect() protoreflect.Message {
	mi := &file_it_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopUserReply_Info.ProtoReflect.Descriptor instead.
func (*GetTopUserReply_Info) Descriptor() ([]byte, []int) {
	return file_it_v1_user_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetTopUserReply_Info) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetTopUserReply_Info) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GetTopUserReply_Info) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetTopUserReply_Info) GetColumnUrl() string {
	if x != nil {
		return x.ColumnUrl
	}
	return ""
}

func (x *GetTopUserReply_Info) GetPostTime() int64 {
	if x != nil {
		return x.PostTime
	}
	return 0
}

var File_it_v1_user_proto protoreflect.FileDescriptor

var file_it_v1_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22,
	0xa9, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x55,
	0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6d, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65,
	0x78, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73,
	0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6d, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xd5,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x1a, 0x90, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5e,
	0x0a, 0x11, 0x42, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x29,
	0x0a, 0x0f, 0x42, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x66, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x70, 0x22, 0x26, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xbd, 0x03, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x59, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x51, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x70,
	0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x73, 0x65, 0x74, 0x12, 0x5c,
	0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x4b, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_it_v1_user_proto_rawDescOnce sync.Once
	file_it_v1_user_proto_rawDescData = file_it_v1_user_proto_rawDesc
)

func file_it_v1_user_proto_rawDescGZIP() []byte {
	file_it_v1_user_proto_rawDescOnce.Do(func() {
		file_it_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_it_v1_user_proto_rawDescData)
	})
	return file_it_v1_user_proto_rawDescData
}

var file_it_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_it_v1_user_proto_goTypes = []interface{}{
	(*GetUserInfoRequest)(nil),   // 0: it.v1.GetUserInfoRequest
	(*GetUserInfoReply)(nil),     // 1: it.v1.GetUserInfoReply
	(*GetTopUserRequest)(nil),    // 2: it.v1.GetTopUserRequest
	(*GetTopUserReply)(nil),      // 3: it.v1.GetTopUserReply
	(*SetUserInfoRequest)(nil),   // 4: it.v1.SetUserInfoRequest
	(*SetUserInfoReply)(nil),     // 5: it.v1.SetUserInfoReply
	(*BindMobileRequest)(nil),    // 6: it.v1.BindMobileRequest
	(*BindMobileReply)(nil),      // 7: it.v1.BindMobileReply
	(*AddUserRequest)(nil),       // 8: it.v1.AddUserRequest
	(*AddUserReply)(nil),         // 9: it.v1.AddUserReply
	(*GetTopUserReply_Info)(nil), // 10: it.v1.GetTopUserReply.Info
}
var file_it_v1_user_proto_depIdxs = []int32{
	10, // 0: it.v1.GetTopUserReply.item:type_name -> it.v1.GetTopUserReply.Info
	0,  // 1: it.v1.User.GetUserInfo:input_type -> it.v1.GetUserInfoRequest
	2,  // 2: it.v1.User.GetTopUser:input_type -> it.v1.GetTopUserRequest
	4,  // 3: it.v1.User.SetUserInfo:input_type -> it.v1.SetUserInfoRequest
	6,  // 4: it.v1.User.BindMobile:input_type -> it.v1.BindMobileRequest
	8,  // 5: it.v1.User.AddUser:input_type -> it.v1.AddUserRequest
	1,  // 6: it.v1.User.GetUserInfo:output_type -> it.v1.GetUserInfoReply
	3,  // 7: it.v1.User.GetTopUser:output_type -> it.v1.GetTopUserReply
	5,  // 8: it.v1.User.SetUserInfo:output_type -> it.v1.SetUserInfoReply
	7,  // 9: it.v1.User.BindMobile:output_type -> it.v1.BindMobileReply
	9,  // 10: it.v1.User.AddUser:output_type -> it.v1.AddUserReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_it_v1_user_proto_init() }
func file_it_v1_user_proto_init() {
	if File_it_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_it_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRequest); i {
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
		file_it_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReply); i {
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
		file_it_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopUserRequest); i {
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
		file_it_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopUserReply); i {
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
		file_it_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserInfoRequest); i {
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
		file_it_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserInfoReply); i {
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
		file_it_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindMobileRequest); i {
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
		file_it_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindMobileReply); i {
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
		file_it_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserRequest); i {
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
		file_it_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserReply); i {
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
		file_it_v1_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopUserReply_Info); i {
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
			RawDescriptor: file_it_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_it_v1_user_proto_goTypes,
		DependencyIndexes: file_it_v1_user_proto_depIdxs,
		MessageInfos:      file_it_v1_user_proto_msgTypes,
	}.Build()
	File_it_v1_user_proto = out.File
	file_it_v1_user_proto_rawDesc = nil
	file_it_v1_user_proto_goTypes = nil
	file_it_v1_user_proto_depIdxs = nil
}