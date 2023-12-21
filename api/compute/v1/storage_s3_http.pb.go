// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.2
// source: api/compute/v1/storage_s3.proto

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

const OperationStorageS3CreateBucket = "/api.compute.v1.StorageS3/CreateBucket"
const OperationStorageS3CreateS3Key = "/api.compute.v1.StorageS3/CreateS3Key"
const OperationStorageS3DeleteBucket = "/api.compute.v1.StorageS3/DeleteBucket"
const OperationStorageS3EmptyBucket = "/api.compute.v1.StorageS3/EmptyBucket"
const OperationStorageS3GetUserS3User = "/api.compute.v1.StorageS3/GetUserS3User"
const OperationStorageS3ListBucket = "/api.compute.v1.StorageS3/ListBucket"
const OperationStorageS3S3StorageDelete = "/api.compute.v1.StorageS3/S3StorageDelete"
const OperationStorageS3S3StorageDownload = "/api.compute.v1.StorageS3/S3StorageDownload"
const OperationStorageS3S3StorageInBucketList = "/api.compute.v1.StorageS3/S3StorageInBucketList"
const OperationStorageS3S3StorageMkdir = "/api.compute.v1.StorageS3/S3StorageMkdir"
const OperationStorageS3S3StorageUploadFile = "/api.compute.v1.StorageS3/S3StorageUploadFile"

type StorageS3HTTPServer interface {
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketReply, error)
	CreateS3Key(context.Context, *CreateS3KeyRequest) (*CreateS3KeyReply, error)
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketReply, error)
	EmptyBucket(context.Context, *EmptyBucketRequest) (*EmptyBucketReply, error)
	GetUserS3User(context.Context, *GetS3UserRequest) (*GetS3UserReply, error)
	ListBucket(context.Context, *ListBucketRequest) (*ListBucketReply, error)
	S3StorageDelete(context.Context, *S3StorageDeleteRequest) (*S3StorageDeleteReply, error)
	S3StorageDownload(context.Context, *S3StorageDownloadRequest) (*S3StorageDownloadReply, error)
	S3StorageInBucketList(context.Context, *S3StorageInBucketListRequest) (*S3StorageInBucketListReply, error)
	S3StorageMkdir(context.Context, *S3StorageMkdirRequest) (*S3StorageMkdirReply, error)
	S3StorageUploadFile(context.Context, *S3StorageUploadFileRequest) (*S3StorageUploadFileReply, error)
}

func RegisterStorageS3HTTPServer(s *http.Server, srv StorageS3HTTPServer) {
	r := s.Route("/")
	r.POST("/v1/s3user/create/key", _StorageS3_CreateS3Key0_HTTP_Handler(srv))
	r.GET("/v1/s3user", _StorageS3_GetUserS3User0_HTTP_Handler(srv))
	r.POST("/v1/s3bucket", _StorageS3_CreateBucket0_HTTP_Handler(srv))
	r.DELETE("/v1/s3bucket/{bucketName}", _StorageS3_DeleteBucket0_HTTP_Handler(srv))
	r.DELETE("/v1/s3bucket/{bucketName}/empty", _StorageS3_EmptyBucket0_HTTP_Handler(srv))
	r.GET("/v1/s3bucket", _StorageS3_ListBucket0_HTTP_Handler(srv))
	r.GET("/v1/s3bucket/{bucketName}/objects", _StorageS3_S3StorageInBucketList0_HTTP_Handler(srv))
	r.POST("/v1/storage/{bucketName}/objects/upload/ersatz", _StorageS3_S3StorageUploadFile0_HTTP_Handler(srv))
	r.POST("/v1/storage/{bucketName}/mkdir", _StorageS3_S3StorageMkdir0_HTTP_Handler(srv))
	r.GET("/v1/storage/{bucketName}/objects/download/ersatz", _StorageS3_S3StorageDownload0_HTTP_Handler(srv))
	r.DELETE("/v1/storage/{bucketName}/objects/delete", _StorageS3_S3StorageDelete0_HTTP_Handler(srv))
}

func _StorageS3_CreateS3Key0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateS3KeyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3CreateS3Key)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateS3Key(ctx, req.(*CreateS3KeyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateS3KeyReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_GetUserS3User0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetS3UserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3GetUserS3User)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserS3User(ctx, req.(*GetS3UserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetS3UserReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_CreateBucket0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBucketRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3CreateBucket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBucket(ctx, req.(*CreateBucketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBucketReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_DeleteBucket0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBucketRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3DeleteBucket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBucket(ctx, req.(*DeleteBucketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBucketReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_EmptyBucket0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EmptyBucketRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3EmptyBucket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EmptyBucket(ctx, req.(*EmptyBucketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyBucketReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_ListBucket0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBucketRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3ListBucket)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBucket(ctx, req.(*ListBucketRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBucketReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_S3StorageInBucketList0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageInBucketListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageInBucketList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageInBucketList(ctx, req.(*S3StorageInBucketListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageInBucketListReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_S3StorageUploadFile0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageUploadFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageUploadFile(ctx, req.(*S3StorageUploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageUploadFileReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_S3StorageMkdir0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageMkdirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageMkdir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageMkdir(ctx, req.(*S3StorageMkdirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageMkdirReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_S3StorageDownload0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageDownloadRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageDownload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageDownload(ctx, req.(*S3StorageDownloadRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageDownloadReply)
		return ctx.Result(200, reply)
	}
}

func _StorageS3_S3StorageDelete0_HTTP_Handler(srv StorageS3HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in S3StorageDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStorageS3S3StorageDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.S3StorageDelete(ctx, req.(*S3StorageDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*S3StorageDeleteReply)
		return ctx.Result(200, reply)
	}
}

type StorageS3HTTPClient interface {
	CreateBucket(ctx context.Context, req *CreateBucketRequest, opts ...http.CallOption) (rsp *CreateBucketReply, err error)
	CreateS3Key(ctx context.Context, req *CreateS3KeyRequest, opts ...http.CallOption) (rsp *CreateS3KeyReply, err error)
	DeleteBucket(ctx context.Context, req *DeleteBucketRequest, opts ...http.CallOption) (rsp *DeleteBucketReply, err error)
	EmptyBucket(ctx context.Context, req *EmptyBucketRequest, opts ...http.CallOption) (rsp *EmptyBucketReply, err error)
	GetUserS3User(ctx context.Context, req *GetS3UserRequest, opts ...http.CallOption) (rsp *GetS3UserReply, err error)
	ListBucket(ctx context.Context, req *ListBucketRequest, opts ...http.CallOption) (rsp *ListBucketReply, err error)
	S3StorageDelete(ctx context.Context, req *S3StorageDeleteRequest, opts ...http.CallOption) (rsp *S3StorageDeleteReply, err error)
	S3StorageDownload(ctx context.Context, req *S3StorageDownloadRequest, opts ...http.CallOption) (rsp *S3StorageDownloadReply, err error)
	S3StorageInBucketList(ctx context.Context, req *S3StorageInBucketListRequest, opts ...http.CallOption) (rsp *S3StorageInBucketListReply, err error)
	S3StorageMkdir(ctx context.Context, req *S3StorageMkdirRequest, opts ...http.CallOption) (rsp *S3StorageMkdirReply, err error)
	S3StorageUploadFile(ctx context.Context, req *S3StorageUploadFileRequest, opts ...http.CallOption) (rsp *S3StorageUploadFileReply, err error)
}

type StorageS3HTTPClientImpl struct {
	cc *http.Client
}

func NewStorageS3HTTPClient(client *http.Client) StorageS3HTTPClient {
	return &StorageS3HTTPClientImpl{client}
}

func (c *StorageS3HTTPClientImpl) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...http.CallOption) (*CreateBucketReply, error) {
	var out CreateBucketReply
	pattern := "/v1/s3bucket"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageS3CreateBucket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) CreateS3Key(ctx context.Context, in *CreateS3KeyRequest, opts ...http.CallOption) (*CreateS3KeyReply, error) {
	var out CreateS3KeyReply
	pattern := "/v1/s3user/create/key"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageS3CreateS3Key))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...http.CallOption) (*DeleteBucketReply, error) {
	var out DeleteBucketReply
	pattern := "/v1/s3bucket/{bucketName}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3DeleteBucket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) EmptyBucket(ctx context.Context, in *EmptyBucketRequest, opts ...http.CallOption) (*EmptyBucketReply, error) {
	var out EmptyBucketReply
	pattern := "/v1/s3bucket/{bucketName}/empty"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3EmptyBucket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) GetUserS3User(ctx context.Context, in *GetS3UserRequest, opts ...http.CallOption) (*GetS3UserReply, error) {
	var out GetS3UserReply
	pattern := "/v1/s3user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3GetUserS3User))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) ListBucket(ctx context.Context, in *ListBucketRequest, opts ...http.CallOption) (*ListBucketReply, error) {
	var out ListBucketReply
	pattern := "/v1/s3bucket"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3ListBucket))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) S3StorageDelete(ctx context.Context, in *S3StorageDeleteRequest, opts ...http.CallOption) (*S3StorageDeleteReply, error) {
	var out S3StorageDeleteReply
	pattern := "/v1/storage/{bucketName}/objects/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3S3StorageDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) S3StorageDownload(ctx context.Context, in *S3StorageDownloadRequest, opts ...http.CallOption) (*S3StorageDownloadReply, error) {
	var out S3StorageDownloadReply
	pattern := "/v1/storage/{bucketName}/objects/download/ersatz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3S3StorageDownload))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) S3StorageInBucketList(ctx context.Context, in *S3StorageInBucketListRequest, opts ...http.CallOption) (*S3StorageInBucketListReply, error) {
	var out S3StorageInBucketListReply
	pattern := "/v1/s3bucket/{bucketName}/objects"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStorageS3S3StorageInBucketList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) S3StorageMkdir(ctx context.Context, in *S3StorageMkdirRequest, opts ...http.CallOption) (*S3StorageMkdirReply, error) {
	var out S3StorageMkdirReply
	pattern := "/v1/storage/{bucketName}/mkdir"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageS3S3StorageMkdir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StorageS3HTTPClientImpl) S3StorageUploadFile(ctx context.Context, in *S3StorageUploadFileRequest, opts ...http.CallOption) (*S3StorageUploadFileReply, error) {
	var out S3StorageUploadFileReply
	pattern := "/v1/storage/{bucketName}/objects/upload/ersatz"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStorageS3S3StorageUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
