// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v3.12.4
// source: api/network_mapping/v1/domain_binding.proto

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

const OperationDomainBindingCreateDomainBinding = "/api.network_mapping.v1.DomainBinding/CreateDomainBinding"
const OperationDomainBindingDeleteDomainBinding = "/api.network_mapping.v1.DomainBinding/DeleteDomainBinding"
const OperationDomainBindingGetDomainBinding = "/api.network_mapping.v1.DomainBinding/GetDomainBinding"
const OperationDomainBindingListDomainBinding = "/api.network_mapping.v1.DomainBinding/ListDomainBinding"
const OperationDomainBindingNsLookup = "/api.network_mapping.v1.DomainBinding/NsLookup"
const OperationDomainBindingUpdateDomainBinding = "/api.network_mapping.v1.DomainBinding/UpdateDomainBinding"

type DomainBindingHTTPServer interface {
	CreateDomainBinding(context.Context, *CreateDomainBindingRequest) (*CreateDomainBindingReply, error)
	DeleteDomainBinding(context.Context, *DeleteDomainBindingRequest) (*DeleteDomainBindingReply, error)
	GetDomainBinding(context.Context, *GetDomainBindingRequest) (*GetDomainBindingReply, error)
	ListDomainBinding(context.Context, *ListDomainBindingRequest) (*ListDomainBindingReply, error)
	NsLookup(context.Context, *NsLookupRequest) (*NsLookupReply, error)
	UpdateDomainBinding(context.Context, *UpdateDomainBindingRequest) (*UpdateDomainBindingReply, error)
}

func RegisterDomainBindingHTTPServer(s *http.Server, srv DomainBindingHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/domain-binding", _DomainBinding_CreateDomainBinding0_HTTP_Handler(srv))
	r.GET("/v1/domain-binding/nslookup", _DomainBinding_NsLookup0_HTTP_Handler(srv))
	r.PUT("/v1/domain-binding/{id}", _DomainBinding_UpdateDomainBinding0_HTTP_Handler(srv))
	r.DELETE("/v1/domain-binding/{id}", _DomainBinding_DeleteDomainBinding0_HTTP_Handler(srv))
	r.GET("/v1/domain-binding/{id}", _DomainBinding_GetDomainBinding0_HTTP_Handler(srv))
	r.GET("/v1/domain-binding", _DomainBinding_ListDomainBinding0_HTTP_Handler(srv))
}

func _DomainBinding_CreateDomainBinding0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDomainBindingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingCreateDomainBinding)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDomainBinding(ctx, req.(*CreateDomainBindingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDomainBindingReply)
		return ctx.Result(200, reply)
	}
}

func _DomainBinding_NsLookup0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NsLookupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingNsLookup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.NsLookup(ctx, req.(*NsLookupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*NsLookupReply)
		return ctx.Result(200, reply)
	}
}

func _DomainBinding_UpdateDomainBinding0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDomainBindingRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingUpdateDomainBinding)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDomainBinding(ctx, req.(*UpdateDomainBindingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDomainBindingReply)
		return ctx.Result(200, reply)
	}
}

func _DomainBinding_DeleteDomainBinding0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDomainBindingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingDeleteDomainBinding)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDomainBinding(ctx, req.(*DeleteDomainBindingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDomainBindingReply)
		return ctx.Result(200, reply)
	}
}

func _DomainBinding_GetDomainBinding0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDomainBindingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingGetDomainBinding)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDomainBinding(ctx, req.(*GetDomainBindingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDomainBindingReply)
		return ctx.Result(200, reply)
	}
}

func _DomainBinding_ListDomainBinding0_HTTP_Handler(srv DomainBindingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDomainBindingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDomainBindingListDomainBinding)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDomainBinding(ctx, req.(*ListDomainBindingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDomainBindingReply)
		return ctx.Result(200, reply)
	}
}

type DomainBindingHTTPClient interface {
	CreateDomainBinding(ctx context.Context, req *CreateDomainBindingRequest, opts ...http.CallOption) (rsp *CreateDomainBindingReply, err error)
	DeleteDomainBinding(ctx context.Context, req *DeleteDomainBindingRequest, opts ...http.CallOption) (rsp *DeleteDomainBindingReply, err error)
	GetDomainBinding(ctx context.Context, req *GetDomainBindingRequest, opts ...http.CallOption) (rsp *GetDomainBindingReply, err error)
	ListDomainBinding(ctx context.Context, req *ListDomainBindingRequest, opts ...http.CallOption) (rsp *ListDomainBindingReply, err error)
	NsLookup(ctx context.Context, req *NsLookupRequest, opts ...http.CallOption) (rsp *NsLookupReply, err error)
	UpdateDomainBinding(ctx context.Context, req *UpdateDomainBindingRequest, opts ...http.CallOption) (rsp *UpdateDomainBindingReply, err error)
}

type DomainBindingHTTPClientImpl struct {
	cc *http.Client
}

func NewDomainBindingHTTPClient(client *http.Client) DomainBindingHTTPClient {
	return &DomainBindingHTTPClientImpl{client}
}

func (c *DomainBindingHTTPClientImpl) CreateDomainBinding(ctx context.Context, in *CreateDomainBindingRequest, opts ...http.CallOption) (*CreateDomainBindingReply, error) {
	var out CreateDomainBindingReply
	pattern := "/v1/domain-binding"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDomainBindingCreateDomainBinding))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainBindingHTTPClientImpl) DeleteDomainBinding(ctx context.Context, in *DeleteDomainBindingRequest, opts ...http.CallOption) (*DeleteDomainBindingReply, error) {
	var out DeleteDomainBindingReply
	pattern := "/v1/domain-binding/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainBindingDeleteDomainBinding))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainBindingHTTPClientImpl) GetDomainBinding(ctx context.Context, in *GetDomainBindingRequest, opts ...http.CallOption) (*GetDomainBindingReply, error) {
	var out GetDomainBindingReply
	pattern := "/v1/domain-binding/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainBindingGetDomainBinding))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainBindingHTTPClientImpl) ListDomainBinding(ctx context.Context, in *ListDomainBindingRequest, opts ...http.CallOption) (*ListDomainBindingReply, error) {
	var out ListDomainBindingReply
	pattern := "/v1/domain-binding"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainBindingListDomainBinding))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainBindingHTTPClientImpl) NsLookup(ctx context.Context, in *NsLookupRequest, opts ...http.CallOption) (*NsLookupReply, error) {
	var out NsLookupReply
	pattern := "/v1/domain-binding/nslookup"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDomainBindingNsLookup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DomainBindingHTTPClientImpl) UpdateDomainBinding(ctx context.Context, in *UpdateDomainBindingRequest, opts ...http.CallOption) (*UpdateDomainBindingReply, error) {
	var out UpdateDomainBindingReply
	pattern := "/v1/domain-binding/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDomainBindingUpdateDomainBinding))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
