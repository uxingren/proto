// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.12
// source: it/v1/oss.proto

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

const OperationOSSOssUpload = "/it.v1.OSS/OssUpload"

type OSSHTTPServer interface {
	OssUpload(context.Context, *OssUploadRequest) (*OssUploadReply, error)
}

func RegisterOSSHTTPServer(s *http.Server, srv OSSHTTPServer) {
	r := s.Route("/")
	r.POST("/oss/upload", _OSS_OssUpload0_HTTP_Handler(srv))
}

func _OSS_OssUpload0_HTTP_Handler(srv OSSHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OssUploadRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOSSOssUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OssUpload(ctx, req.(*OssUploadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OssUploadReply)
		return ctx.Result(200, reply)
	}
}

type OSSHTTPClient interface {
	OssUpload(ctx context.Context, req *OssUploadRequest, opts ...http.CallOption) (rsp *OssUploadReply, err error)
}

type OSSHTTPClientImpl struct {
	cc *http.Client
}

func NewOSSHTTPClient(client *http.Client) OSSHTTPClient {
	return &OSSHTTPClientImpl{client}
}

func (c *OSSHTTPClientImpl) OssUpload(ctx context.Context, in *OssUploadRequest, opts ...http.CallOption) (*OssUploadReply, error) {
	var out OssUploadReply
	pattern := "/oss/upload"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOSSOssUpload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
