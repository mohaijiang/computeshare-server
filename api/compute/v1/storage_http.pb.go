// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: compute/v1/storage.proto

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

const OperationStorageCreateDir = "/api.compute.v1.Storage/CreateDir"
const OperationStorageDelete = "/api.compute.v1.Storage/Delete"
const OperationStorageDownload = "/api.compute.v1.Storage/Download"
const OperationStorageList = "/api.compute.v1.Storage/List"
const OperationStorageUploadFile = "/api.compute.v1.Storage/UploadFile"

type StorageHTTPServer interface {
	CreateDir(context.Context, *CreateDirRequest) (*CreateDirReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
	Download(context.Context, *DownloadRequest) (*DownloadReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	UploadFile(context.Context, *UploadFileRequest) (*File, error)
}

func RegisterStorageHTTPServer(s *http.Server, srv StorageHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/storage", _Storage_List0_HTTP_Handler(srv))
	r.POST("/v1/storage", _Storage_UploadFile0_HTTP_Handler(srv))
	r.POST("/v1/storage/dir", _Storage_CreateDir0_HTTP_Handler(srv))
	r.GET("/v1/storage/{id}", _Storage_Download0_HTTP_Handler(srv))
	r.DELETE("/v1/storage", _Storage_Delete0_HTTP_Handler(srv))
}

func _Storage_List0_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListReply)
		return ctx.Result(200, reply)
	}
}

func _Storage_UploadFile0_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFile(ctx, req.(*UploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*File)
		return ctx.Result(200, reply)
	}
}

func _Storage_CreateDir0_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageCreateDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDir(ctx, req.(*CreateDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDirReply)
		return ctx.Result(200, reply)
	}
}

func _Storage_Download0_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DownloadRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageDownload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Download(ctx, req.(*DownloadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DownloadReply)
		return ctx.Result(200, reply)
	}
}

func _Storage_Delete0_HTTP_Handler(srv StorageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteReply)
		return ctx.Result(200, reply)
	}
}

type StorageHTTPClient interface {
	CreateDir(ctx context.Context, req *CreateDirRequest, opts ...http.CallOption) (rsp *CreateDirReply, err error)
	Delete(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	Download(ctx context.Context, req *DownloadRequest, opts ...http.CallOption) (rsp *DownloadReply, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListReply, err error)
	UploadFile(ctx context.Context, req *UploadFileRequest, opts ...http.CallOption) (rsp *File, err error)
}

type StorageHTTPClientImpl struct {
	cc *http.Client
}

func NewStorageHTTPClient(client *http.Client) StorageHTTPClient {
	return &StorageHTTPClientImpl{client}
}

func (c *StorageHTTPClientImpl) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...http.CallOption) (*CreateDirReply, error) {
	var out CreateDirReply
	pattern := "/v1/storage/dir"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageCreateDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageHTTPClientImpl) Delete(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*DeleteReply, error) {
	var out DeleteReply
	pattern := "/v1/storage"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageHTTPClientImpl) Download(ctx context.Context, in *DownloadRequest, opts ...http.CallOption) (*DownloadReply, error) {
	var out DownloadReply
	pattern := "/v1/storage/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageDownload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageHTTPClientImpl) List(ctx context.Context, in *ListRequest, opts ...http.CallOption) (*ListReply, error) {
	var out ListReply
	pattern := "/v1/storage"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageHTTPClientImpl) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...http.CallOption) (*File, error) {
	var out File
	pattern := "/v1/storage"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
