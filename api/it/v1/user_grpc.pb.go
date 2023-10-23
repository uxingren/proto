// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: it/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_GetUserInfo_FullMethodName = "/it.v1.User/GetUserInfo"
	User_GetTopUser_FullMethodName  = "/it.v1.User/GetTopUser"
	User_SetUserInfo_FullMethodName = "/it.v1.User/SetUserInfo"
	User_BindMobile_FullMethodName  = "/it.v1.User/BindMobile"
	User_AddUser_FullMethodName     = "/it.v1.User/AddUser"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error)
	// 获取用户信息
	GetTopUser(ctx context.Context, in *GetTopUserRequest, opts ...grpc.CallOption) (*GetTopUserReply, error)
	// 设置用户信息
	SetUserInfo(ctx context.Context, in *SetUserInfoRequest, opts ...grpc.CallOption) (*SetUserInfoReply, error)
	// 绑定手机号
	BindMobile(ctx context.Context, in *BindMobileRequest, opts ...grpc.CallOption) (*BindMobileReply, error)
	// 注册用户
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := c.cc.Invoke(ctx, User_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetTopUser(ctx context.Context, in *GetTopUserRequest, opts ...grpc.CallOption) (*GetTopUserReply, error) {
	out := new(GetTopUserReply)
	err := c.cc.Invoke(ctx, User_GetTopUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetUserInfo(ctx context.Context, in *SetUserInfoRequest, opts ...grpc.CallOption) (*SetUserInfoReply, error) {
	out := new(SetUserInfoReply)
	err := c.cc.Invoke(ctx, User_SetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindMobile(ctx context.Context, in *BindMobileRequest, opts ...grpc.CallOption) (*BindMobileReply, error) {
	out := new(BindMobileReply)
	err := c.cc.Invoke(ctx, User_BindMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserReply, error) {
	out := new(AddUserReply)
	err := c.cc.Invoke(ctx, User_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	// 获取用户信息
	GetTopUser(context.Context, *GetTopUserRequest) (*GetTopUserReply, error)
	// 设置用户信息
	SetUserInfo(context.Context, *SetUserInfoRequest) (*SetUserInfoReply, error)
	// 绑定手机号
	BindMobile(context.Context, *BindMobileRequest) (*BindMobileReply, error)
	// 注册用户
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) GetTopUser(context.Context, *GetTopUserRequest) (*GetTopUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUser not implemented")
}
func (UnimplementedUserServer) SetUserInfo(context.Context, *SetUserInfoRequest) (*SetUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserInfo not implemented")
}
func (UnimplementedUserServer) BindMobile(context.Context, *BindMobileRequest) (*BindMobileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindMobile not implemented")
}
func (UnimplementedUserServer) AddUser(context.Context, *AddUserRequest) (*AddUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetTopUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetTopUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetTopUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetTopUser(ctx, req.(*GetTopUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetUserInfo(ctx, req.(*SetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_BindMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindMobile(ctx, req.(*BindMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "it.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "GetTopUser",
			Handler:    _User_GetTopUser_Handler,
		},
		{
			MethodName: "SetUserInfo",
			Handler:    _User_SetUserInfo_Handler,
		},
		{
			MethodName: "BindMobile",
			Handler:    _User_BindMobile_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _User_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "it/v1/user.proto",
}
