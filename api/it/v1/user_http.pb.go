// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.12
// source: it/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserAddUser = "/it.v1.User/AddUser"
const OperationUserBindMobile = "/it.v1.User/BindMobile"
const OperationUserGetTopUser = "/it.v1.User/GetTopUser"
const OperationUserGetUserInfo = "/it.v1.User/GetUserInfo"
const OperationUserSetUserInfo = "/it.v1.User/SetUserInfo"

type UserHTTPServer interface {
	// AddUser 注册用户
	AddUser(context.Context, *AddUserRequest) (*AddUserReply, error)
	// BindMobile 绑定手机号
	BindMobile(context.Context, *BindMobileRequest) (*BindMobileReply, error)
	// GetTopUser 获取用户信息
	GetTopUser(context.Context, *GetTopUserRequest) (*GetTopUserReply, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	// SetUserInfo 设置用户信息
	SetUserInfo(context.Context, *SetUserInfoRequest) (*SetUserInfoReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.GET("/user/info/get", _User_GetUserInfo0_HTTP_Handler(srv))
	r.GET("/user/top", _User_GetTopUser0_HTTP_Handler(srv))
	r.POST("/user/info/set", _User_SetUserInfo0_HTTP_Handler(srv))
	r.POST("/user/mobile/bind", _User_BindMobile0_HTTP_Handler(srv))
	r.POST("/user/add", _User_AddUser0_HTTP_Handler(srv))
}

func _User_GetUserInfo0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetTopUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTopUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetTopUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTopUser(ctx, req.(*GetTopUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTopUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_SetUserInfo0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserInfo(ctx, req.(*SetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserInfoReply)
		return ctx.Result(200, reply)
	}
}

func _User_BindMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindMobileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserBindMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindMobile(ctx, req.(*BindMobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BindMobileReply)
		return ctx.Result(200, reply)
	}
}

func _User_AddUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAddUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddUser(ctx, req.(*AddUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddUserReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	AddUser(ctx context.Context, req *AddUserRequest, opts ...http.CallOption) (rsp *AddUserReply, err error)
	BindMobile(ctx context.Context, req *BindMobileRequest, opts ...http.CallOption) (rsp *BindMobileReply, err error)
	GetTopUser(ctx context.Context, req *GetTopUserRequest, opts ...http.CallOption) (rsp *GetTopUserReply, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoReply, err error)
	SetUserInfo(ctx context.Context, req *SetUserInfoRequest, opts ...http.CallOption) (rsp *SetUserInfoReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) AddUser(ctx context.Context, in *AddUserRequest, opts ...http.CallOption) (*AddUserReply, error) {
	var out AddUserReply
	pattern := "/user/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAddUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) BindMobile(ctx context.Context, in *BindMobileRequest, opts ...http.CallOption) (*BindMobileReply, error) {
	var out BindMobileReply
	pattern := "/user/mobile/bind"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserBindMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetTopUser(ctx context.Context, in *GetTopUserRequest, opts ...http.CallOption) (*GetTopUserReply, error) {
	var out GetTopUserReply
	pattern := "/user/top"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetTopUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoReply, error) {
	var out GetUserInfoReply
	pattern := "/user/info/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SetUserInfo(ctx context.Context, in *SetUserInfoRequest, opts ...http.CallOption) (*SetUserInfoReply, error) {
	var out SetUserInfoReply
	pattern := "/user/info/set"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
