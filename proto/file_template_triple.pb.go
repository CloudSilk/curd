﻿// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: file_template.proto

package curd

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// FileTemplateClient is the client API for FileTemplate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTemplateClient interface {
	Add(ctx context.Context, in *FileTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Update(ctx context.Context, in *FileTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Query(ctx context.Context, in *QueryFileTemplateRequest, opts ...grpc_go.CallOption) (*QueryFileTemplateResponse, common.ErrorWithAttachment)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllFileTemplateResponse, common.ErrorWithAttachment)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetFileTemplateDetailResponse, common.ErrorWithAttachment)
	Copy(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
}

type fileTemplateClient struct {
	cc *triple.TripleConn
}

type FileTemplateClientImpl struct {
	Add       func(ctx context.Context, in *FileTemplateInfo) (*CommonResponse, error)
	Update    func(ctx context.Context, in *FileTemplateInfo) (*CommonResponse, error)
	Delete    func(ctx context.Context, in *DelRequest) (*CommonResponse, error)
	Query     func(ctx context.Context, in *QueryFileTemplateRequest) (*QueryFileTemplateResponse, error)
	GetAll    func(ctx context.Context, in *GetAllRequest) (*GetAllFileTemplateResponse, error)
	GetDetail func(ctx context.Context, in *GetDetailRequest) (*GetFileTemplateDetailResponse, error)
	Copy      func(ctx context.Context, in *GetDetailRequest) (*CommonResponse, error)
}

func (c *FileTemplateClientImpl) GetDubboStub(cc *triple.TripleConn) FileTemplateClient {
	return NewFileTemplateClient(cc)
}

func (c *FileTemplateClientImpl) XXX_InterfaceName() string {
	return "curd.FileTemplate"
}

func NewFileTemplateClient(cc *triple.TripleConn) FileTemplateClient {
	return &fileTemplateClient{cc}
}

func (c *fileTemplateClient) Add(ctx context.Context, in *FileTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Add", in, out)
}

func (c *fileTemplateClient) Update(ctx context.Context, in *FileTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Update", in, out)
}

func (c *fileTemplateClient) Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Delete", in, out)
}

func (c *fileTemplateClient) Query(ctx context.Context, in *QueryFileTemplateRequest, opts ...grpc_go.CallOption) (*QueryFileTemplateResponse, common.ErrorWithAttachment) {
	out := new(QueryFileTemplateResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

func (c *fileTemplateClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllFileTemplateResponse, common.ErrorWithAttachment) {
	out := new(GetAllFileTemplateResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetAll", in, out)
}

func (c *fileTemplateClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetFileTemplateDetailResponse, common.ErrorWithAttachment) {
	out := new(GetFileTemplateDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetDetail", in, out)
}

func (c *fileTemplateClient) Copy(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Copy", in, out)
}

// FileTemplateServer is the server API for FileTemplate service.
// All implementations must embed UnimplementedFileTemplateServer
// for forward compatibility
type FileTemplateServer interface {
	Add(context.Context, *FileTemplateInfo) (*CommonResponse, error)
	Update(context.Context, *FileTemplateInfo) (*CommonResponse, error)
	Delete(context.Context, *DelRequest) (*CommonResponse, error)
	Query(context.Context, *QueryFileTemplateRequest) (*QueryFileTemplateResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllFileTemplateResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetFileTemplateDetailResponse, error)
	Copy(context.Context, *GetDetailRequest) (*CommonResponse, error)
	mustEmbedUnimplementedFileTemplateServer()
}

// UnimplementedFileTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedFileTemplateServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedFileTemplateServer) Add(context.Context, *FileTemplateInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedFileTemplateServer) Update(context.Context, *FileTemplateInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFileTemplateServer) Delete(context.Context, *DelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileTemplateServer) Query(context.Context, *QueryFileTemplateRequest) (*QueryFileTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedFileTemplateServer) GetAll(context.Context, *GetAllRequest) (*GetAllFileTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedFileTemplateServer) GetDetail(context.Context, *GetDetailRequest) (*GetFileTemplateDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedFileTemplateServer) Copy(context.Context, *GetDetailRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (s *UnimplementedFileTemplateServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedFileTemplateServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedFileTemplateServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &FileTemplate_ServiceDesc
}
func (s *UnimplementedFileTemplateServer) XXX_InterfaceName() string {
	return "curd.FileTemplate"
}

func (UnimplementedFileTemplateServer) mustEmbedUnimplementedFileTemplateServer() {}

// UnsafeFileTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTemplateServer will
// result in compilation errors.
type UnsafeFileTemplateServer interface {
	mustEmbedUnimplementedFileTemplateServer()
}

func RegisterFileTemplateServer(s grpc_go.ServiceRegistrar, srv FileTemplateServer) {
	s.RegisterService(&FileTemplate_ServiceDesc, srv)
}

func _FileTemplate_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileTemplateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Add", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileTemplateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Update", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Delete", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFileTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Query", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetAll", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetDetail", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTemplate_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Copy", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// FileTemplate_ServiceDesc is the grpc_go.ServiceDesc for FileTemplate service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTemplate_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "curd.FileTemplate",
	HandlerType: (*FileTemplateServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _FileTemplate_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FileTemplate_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FileTemplate_Delete_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _FileTemplate_Query_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _FileTemplate_GetAll_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _FileTemplate_GetDetail_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _FileTemplate_Copy_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "file_template.proto",
}
