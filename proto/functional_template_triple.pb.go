// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.20.3
// source: functional_template.proto

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

// FunctionalTemplateClient is the client API for FunctionalTemplate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FunctionalTemplateClient interface {
	Add(ctx context.Context, in *FunctionalTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Update(ctx context.Context, in *FunctionalTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Query(ctx context.Context, in *QueryFunctionalTemplateRequest, opts ...grpc_go.CallOption) (*QueryFunctionalTemplateResponse, common.ErrorWithAttachment)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllFunctionalTemplateResponse, common.ErrorWithAttachment)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetFunctionalTemplateDetailResponse, common.ErrorWithAttachment)
	Copy(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
}

type functionalTemplateClient struct {
	cc *triple.TripleConn
}

type FunctionalTemplateClientImpl struct {
	Add       func(ctx context.Context, in *FunctionalTemplateInfo) (*CommonResponse, error)
	Update    func(ctx context.Context, in *FunctionalTemplateInfo) (*CommonResponse, error)
	Delete    func(ctx context.Context, in *DelRequest) (*CommonResponse, error)
	Query     func(ctx context.Context, in *QueryFunctionalTemplateRequest) (*QueryFunctionalTemplateResponse, error)
	GetAll    func(ctx context.Context, in *GetAllRequest) (*GetAllFunctionalTemplateResponse, error)
	GetDetail func(ctx context.Context, in *GetDetailRequest) (*GetFunctionalTemplateDetailResponse, error)
	Copy      func(ctx context.Context, in *GetDetailRequest) (*CommonResponse, error)
	Enable    func(ctx context.Context, in *EnableRequest) (*CommonResponse, error)
}

func (c *FunctionalTemplateClientImpl) GetDubboStub(cc *triple.TripleConn) FunctionalTemplateClient {
	return NewFunctionalTemplateClient(cc)
}

func (c *FunctionalTemplateClientImpl) XXX_InterfaceName() string {
	return "curd.FunctionalTemplate"
}

func NewFunctionalTemplateClient(cc *triple.TripleConn) FunctionalTemplateClient {
	return &functionalTemplateClient{cc}
}

func (c *functionalTemplateClient) Add(ctx context.Context, in *FunctionalTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Add", in, out)
}

func (c *functionalTemplateClient) Update(ctx context.Context, in *FunctionalTemplateInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Update", in, out)
}

func (c *functionalTemplateClient) Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Delete", in, out)
}

func (c *functionalTemplateClient) Query(ctx context.Context, in *QueryFunctionalTemplateRequest, opts ...grpc_go.CallOption) (*QueryFunctionalTemplateResponse, common.ErrorWithAttachment) {
	out := new(QueryFunctionalTemplateResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

func (c *functionalTemplateClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllFunctionalTemplateResponse, common.ErrorWithAttachment) {
	out := new(GetAllFunctionalTemplateResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetAll", in, out)
}

func (c *functionalTemplateClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetFunctionalTemplateDetailResponse, common.ErrorWithAttachment) {
	out := new(GetFunctionalTemplateDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetDetail", in, out)
}

func (c *functionalTemplateClient) Copy(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Copy", in, out)
}

func (c *functionalTemplateClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Enable", in, out)
}

// FunctionalTemplateServer is the server API for FunctionalTemplate service.
// All implementations must embed UnimplementedFunctionalTemplateServer
// for forward compatibility
type FunctionalTemplateServer interface {
	Add(context.Context, *FunctionalTemplateInfo) (*CommonResponse, error)
	Update(context.Context, *FunctionalTemplateInfo) (*CommonResponse, error)
	Delete(context.Context, *DelRequest) (*CommonResponse, error)
	Query(context.Context, *QueryFunctionalTemplateRequest) (*QueryFunctionalTemplateResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllFunctionalTemplateResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetFunctionalTemplateDetailResponse, error)
	Copy(context.Context, *GetDetailRequest) (*CommonResponse, error)
	Enable(context.Context, *EnableRequest) (*CommonResponse, error)
	mustEmbedUnimplementedFunctionalTemplateServer()
}

// UnimplementedFunctionalTemplateServer must be embedded to have forward compatible implementations.
type UnimplementedFunctionalTemplateServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedFunctionalTemplateServer) Add(context.Context, *FunctionalTemplateInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedFunctionalTemplateServer) Update(context.Context, *FunctionalTemplateInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFunctionalTemplateServer) Delete(context.Context, *DelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFunctionalTemplateServer) Query(context.Context, *QueryFunctionalTemplateRequest) (*QueryFunctionalTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedFunctionalTemplateServer) GetAll(context.Context, *GetAllRequest) (*GetAllFunctionalTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedFunctionalTemplateServer) GetDetail(context.Context, *GetDetailRequest) (*GetFunctionalTemplateDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedFunctionalTemplateServer) Copy(context.Context, *GetDetailRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedFunctionalTemplateServer) Enable(context.Context, *EnableRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (s *UnimplementedFunctionalTemplateServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedFunctionalTemplateServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedFunctionalTemplateServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &FunctionalTemplate_ServiceDesc
}
func (s *UnimplementedFunctionalTemplateServer) XXX_InterfaceName() string {
	return "curd.FunctionalTemplate"
}

func (UnimplementedFunctionalTemplateServer) mustEmbedUnimplementedFunctionalTemplateServer() {}

// UnsafeFunctionalTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FunctionalTemplateServer will
// result in compilation errors.
type UnsafeFunctionalTemplateServer interface {
	mustEmbedUnimplementedFunctionalTemplateServer()
}

func RegisterFunctionalTemplateServer(s grpc_go.ServiceRegistrar, srv FunctionalTemplateServer) {
	s.RegisterService(&FunctionalTemplate_ServiceDesc, srv)
}

func _FunctionalTemplate_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionalTemplateInfo)
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

func _FunctionalTemplate_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionalTemplateInfo)
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

func _FunctionalTemplate_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
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

func _FunctionalTemplate_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFunctionalTemplateRequest)
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

func _FunctionalTemplate_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
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

func _FunctionalTemplate_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
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

func _FunctionalTemplate_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
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

func _FunctionalTemplate_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
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
	invo := invocation.NewRPCInvocation("Enable", args, invAttachment)
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

// FunctionalTemplate_ServiceDesc is the grpc_go.ServiceDesc for FunctionalTemplate service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var FunctionalTemplate_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "curd.FunctionalTemplate",
	HandlerType: (*FunctionalTemplateServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _FunctionalTemplate_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FunctionalTemplate_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FunctionalTemplate_Delete_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _FunctionalTemplate_Query_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _FunctionalTemplate_GetAll_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _FunctionalTemplate_GetDetail_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _FunctionalTemplate_Copy_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _FunctionalTemplate_Enable_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "functional_template.proto",
}
