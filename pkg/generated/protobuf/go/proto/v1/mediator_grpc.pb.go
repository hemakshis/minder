// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1/mediator.proto

package _go

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

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	CheckHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) CheckHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.HealthService/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	CheckHealth(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) CheckHealth(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.HealthService/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).CheckHealth(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _HealthService_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// GitHubWebhookServiceClient is the client API for GitHubWebhookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitHubWebhookServiceClient interface {
	HandleGitHubWebhook(ctx context.Context, in *GitHubWebhookRequest, opts ...grpc.CallOption) (*GitHubWebhookResponse, error)
}

type gitHubWebhookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGitHubWebhookServiceClient(cc grpc.ClientConnInterface) GitHubWebhookServiceClient {
	return &gitHubWebhookServiceClient{cc}
}

func (c *gitHubWebhookServiceClient) HandleGitHubWebhook(ctx context.Context, in *GitHubWebhookRequest, opts ...grpc.CallOption) (*GitHubWebhookResponse, error) {
	out := new(GitHubWebhookResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.GitHubWebhookService/HandleGitHubWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitHubWebhookServiceServer is the server API for GitHubWebhookService service.
// All implementations must embed UnimplementedGitHubWebhookServiceServer
// for forward compatibility
type GitHubWebhookServiceServer interface {
	HandleGitHubWebhook(context.Context, *GitHubWebhookRequest) (*GitHubWebhookResponse, error)
	mustEmbedUnimplementedGitHubWebhookServiceServer()
}

// UnimplementedGitHubWebhookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGitHubWebhookServiceServer struct {
}

func (UnimplementedGitHubWebhookServiceServer) HandleGitHubWebhook(context.Context, *GitHubWebhookRequest) (*GitHubWebhookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGitHubWebhook not implemented")
}
func (UnimplementedGitHubWebhookServiceServer) mustEmbedUnimplementedGitHubWebhookServiceServer() {}

// UnsafeGitHubWebhookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitHubWebhookServiceServer will
// result in compilation errors.
type UnsafeGitHubWebhookServiceServer interface {
	mustEmbedUnimplementedGitHubWebhookServiceServer()
}

func RegisterGitHubWebhookServiceServer(s grpc.ServiceRegistrar, srv GitHubWebhookServiceServer) {
	s.RegisterService(&GitHubWebhookService_ServiceDesc, srv)
}

func _GitHubWebhookService_HandleGitHubWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitHubWebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitHubWebhookServiceServer).HandleGitHubWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.GitHubWebhookService/HandleGitHubWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitHubWebhookServiceServer).HandleGitHubWebhook(ctx, req.(*GitHubWebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GitHubWebhookService_ServiceDesc is the grpc.ServiceDesc for GitHubWebhookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitHubWebhookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.GitHubWebhookService",
	HandlerType: (*GitHubWebhookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleGitHubWebhook",
			Handler:    _GitHubWebhookService_HandleGitHubWebhook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthServiceClient interface {
	GetAuthorizationURL(ctx context.Context, in *AuthorizationURLRequest, opts ...grpc.CallOption) (*AuthorizationURLResponse, error)
	ExchangeCodeForTokenCLI(ctx context.Context, in *CodeExchangeRequestCLI, opts ...grpc.CallOption) (*CodeExchangeResponseCLI, error)
	ExchangeCodeForTokenWEB(ctx context.Context, in *CodeExchangeRequestWEB, opts ...grpc.CallOption) (*CodeExchangeResponseWEB, error)
}

type oAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthServiceClient(cc grpc.ClientConnInterface) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) GetAuthorizationURL(ctx context.Context, in *AuthorizationURLRequest, opts ...grpc.CallOption) (*AuthorizationURLResponse, error) {
	out := new(AuthorizationURLResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.OAuthService/GetAuthorizationURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) ExchangeCodeForTokenCLI(ctx context.Context, in *CodeExchangeRequestCLI, opts ...grpc.CallOption) (*CodeExchangeResponseCLI, error) {
	out := new(CodeExchangeResponseCLI)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.OAuthService/ExchangeCodeForTokenCLI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) ExchangeCodeForTokenWEB(ctx context.Context, in *CodeExchangeRequestWEB, opts ...grpc.CallOption) (*CodeExchangeResponseWEB, error) {
	out := new(CodeExchangeResponseWEB)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.OAuthService/ExchangeCodeForTokenWEB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
// All implementations must embed UnimplementedOAuthServiceServer
// for forward compatibility
type OAuthServiceServer interface {
	GetAuthorizationURL(context.Context, *AuthorizationURLRequest) (*AuthorizationURLResponse, error)
	ExchangeCodeForTokenCLI(context.Context, *CodeExchangeRequestCLI) (*CodeExchangeResponseCLI, error)
	ExchangeCodeForTokenWEB(context.Context, *CodeExchangeRequestWEB) (*CodeExchangeResponseWEB, error)
	mustEmbedUnimplementedOAuthServiceServer()
}

// UnimplementedOAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (UnimplementedOAuthServiceServer) GetAuthorizationURL(context.Context, *AuthorizationURLRequest) (*AuthorizationURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizationURL not implemented")
}
func (UnimplementedOAuthServiceServer) ExchangeCodeForTokenCLI(context.Context, *CodeExchangeRequestCLI) (*CodeExchangeResponseCLI, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeCodeForTokenCLI not implemented")
}
func (UnimplementedOAuthServiceServer) ExchangeCodeForTokenWEB(context.Context, *CodeExchangeRequestWEB) (*CodeExchangeResponseWEB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeCodeForTokenWEB not implemented")
}
func (UnimplementedOAuthServiceServer) mustEmbedUnimplementedOAuthServiceServer() {}

// UnsafeOAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthServiceServer will
// result in compilation errors.
type UnsafeOAuthServiceServer interface {
	mustEmbedUnimplementedOAuthServiceServer()
}

func RegisterOAuthServiceServer(s grpc.ServiceRegistrar, srv OAuthServiceServer) {
	s.RegisterService(&OAuthService_ServiceDesc, srv)
}

func _OAuthService_GetAuthorizationURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).GetAuthorizationURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.OAuthService/GetAuthorizationURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).GetAuthorizationURL(ctx, req.(*AuthorizationURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_ExchangeCodeForTokenCLI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeExchangeRequestCLI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).ExchangeCodeForTokenCLI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.OAuthService/ExchangeCodeForTokenCLI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).ExchangeCodeForTokenCLI(ctx, req.(*CodeExchangeRequestCLI))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_ExchangeCodeForTokenWEB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeExchangeRequestWEB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).ExchangeCodeForTokenWEB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.OAuthService/ExchangeCodeForTokenWEB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).ExchangeCodeForTokenWEB(ctx, req.(*CodeExchangeRequestWEB))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthService_ServiceDesc is the grpc.ServiceDesc for OAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthorizationURL",
			Handler:    _OAuthService_GetAuthorizationURL_Handler,
		},
		{
			MethodName: "ExchangeCodeForTokenCLI",
			Handler:    _OAuthService_ExchangeCodeForTokenCLI_Handler,
		},
		{
			MethodName: "ExchangeCodeForTokenWEB",
			Handler:    _OAuthService_ExchangeCodeForTokenWEB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// LogInServiceClient is the client API for LogInService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogInServiceClient interface {
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
}

type logInServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogInServiceClient(cc grpc.ClientConnInterface) LogInServiceClient {
	return &logInServiceClient{cc}
}

func (c *logInServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.LogInService/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogInServiceServer is the server API for LogInService service.
// All implementations must embed UnimplementedLogInServiceServer
// for forward compatibility
type LogInServiceServer interface {
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	mustEmbedUnimplementedLogInServiceServer()
}

// UnimplementedLogInServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogInServiceServer struct {
}

func (UnimplementedLogInServiceServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedLogInServiceServer) mustEmbedUnimplementedLogInServiceServer() {}

// UnsafeLogInServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogInServiceServer will
// result in compilation errors.
type UnsafeLogInServiceServer interface {
	mustEmbedUnimplementedLogInServiceServer()
}

func RegisterLogInServiceServer(s grpc.ServiceRegistrar, srv LogInServiceServer) {
	s.RegisterService(&LogInService_ServiceDesc, srv)
}

func _LogInService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogInServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.LogInService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogInServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogInService_ServiceDesc is the grpc.ServiceDesc for LogInService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogInService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.LogInService",
	HandlerType: (*LogInServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIn",
			Handler:    _LogInService_LogIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// LogOutServiceClient is the client API for LogOutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogOutServiceClient interface {
	LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error)
}

type logOutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogOutServiceClient(cc grpc.ClientConnInterface) LogOutServiceClient {
	return &logOutServiceClient{cc}
}

func (c *logOutServiceClient) LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error) {
	out := new(LogOutResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.LogOutService/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogOutServiceServer is the server API for LogOutService service.
// All implementations must embed UnimplementedLogOutServiceServer
// for forward compatibility
type LogOutServiceServer interface {
	LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error)
	mustEmbedUnimplementedLogOutServiceServer()
}

// UnimplementedLogOutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogOutServiceServer struct {
}

func (UnimplementedLogOutServiceServer) LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedLogOutServiceServer) mustEmbedUnimplementedLogOutServiceServer() {}

// UnsafeLogOutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogOutServiceServer will
// result in compilation errors.
type UnsafeLogOutServiceServer interface {
	mustEmbedUnimplementedLogOutServiceServer()
}

func RegisterLogOutServiceServer(s grpc.ServiceRegistrar, srv LogOutServiceServer) {
	s.RegisterService(&LogOutService_ServiceDesc, srv)
}

func _LogOutService_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogOutServiceServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.LogOutService/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogOutServiceServer).LogOut(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogOutService_ServiceDesc is the grpc.ServiceDesc for LogOutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogOutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.LogOutService",
	HandlerType: (*LogOutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogOut",
			Handler:    _LogOutService_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// AuthVerifyServiceClient is the client API for AuthVerifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthVerifyServiceClient interface {
	Verify(ctx context.Context, in *AuthVerifyRequest, opts ...grpc.CallOption) (*AuthVerifyResponse, error)
}

type authVerifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthVerifyServiceClient(cc grpc.ClientConnInterface) AuthVerifyServiceClient {
	return &authVerifyServiceClient{cc}
}

func (c *authVerifyServiceClient) Verify(ctx context.Context, in *AuthVerifyRequest, opts ...grpc.CallOption) (*AuthVerifyResponse, error) {
	out := new(AuthVerifyResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.AuthVerifyService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthVerifyServiceServer is the server API for AuthVerifyService service.
// All implementations must embed UnimplementedAuthVerifyServiceServer
// for forward compatibility
type AuthVerifyServiceServer interface {
	Verify(context.Context, *AuthVerifyRequest) (*AuthVerifyResponse, error)
	mustEmbedUnimplementedAuthVerifyServiceServer()
}

// UnimplementedAuthVerifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthVerifyServiceServer struct {
}

func (UnimplementedAuthVerifyServiceServer) Verify(context.Context, *AuthVerifyRequest) (*AuthVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthVerifyServiceServer) mustEmbedUnimplementedAuthVerifyServiceServer() {}

// UnsafeAuthVerifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthVerifyServiceServer will
// result in compilation errors.
type UnsafeAuthVerifyServiceServer interface {
	mustEmbedUnimplementedAuthVerifyServiceServer()
}

func RegisterAuthVerifyServiceServer(s grpc.ServiceRegistrar, srv AuthVerifyServiceServer) {
	s.RegisterService(&AuthVerifyService_ServiceDesc, srv)
}

func _AuthVerifyService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthVerifyServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.AuthVerifyService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthVerifyServiceServer).Verify(ctx, req.(*AuthVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthVerifyService_ServiceDesc is the grpc.ServiceDesc for AuthVerifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthVerifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.AuthVerifyService",
	HandlerType: (*AuthVerifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _AuthVerifyService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// VulnerabilitiesServiceClient is the client API for VulnerabilitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VulnerabilitiesServiceClient interface {
	GetVulnerabilities(ctx context.Context, in *GetVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetVulnerabilitiesResponse, error)
	GetVulnerabilityById(ctx context.Context, in *GetVulnerabilityByIdRequest, opts ...grpc.CallOption) (*Vulnerability, error)
}

type vulnerabilitiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVulnerabilitiesServiceClient(cc grpc.ClientConnInterface) VulnerabilitiesServiceClient {
	return &vulnerabilitiesServiceClient{cc}
}

func (c *vulnerabilitiesServiceClient) GetVulnerabilities(ctx context.Context, in *GetVulnerabilitiesRequest, opts ...grpc.CallOption) (*GetVulnerabilitiesResponse, error) {
	out := new(GetVulnerabilitiesResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.VulnerabilitiesService/GetVulnerabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilitiesServiceClient) GetVulnerabilityById(ctx context.Context, in *GetVulnerabilityByIdRequest, opts ...grpc.CallOption) (*Vulnerability, error) {
	out := new(Vulnerability)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.VulnerabilitiesService/GetVulnerabilityById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerabilitiesServiceServer is the server API for VulnerabilitiesService service.
// All implementations must embed UnimplementedVulnerabilitiesServiceServer
// for forward compatibility
type VulnerabilitiesServiceServer interface {
	GetVulnerabilities(context.Context, *GetVulnerabilitiesRequest) (*GetVulnerabilitiesResponse, error)
	GetVulnerabilityById(context.Context, *GetVulnerabilityByIdRequest) (*Vulnerability, error)
	mustEmbedUnimplementedVulnerabilitiesServiceServer()
}

// UnimplementedVulnerabilitiesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVulnerabilitiesServiceServer struct {
}

func (UnimplementedVulnerabilitiesServiceServer) GetVulnerabilities(context.Context, *GetVulnerabilitiesRequest) (*GetVulnerabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilities not implemented")
}
func (UnimplementedVulnerabilitiesServiceServer) GetVulnerabilityById(context.Context, *GetVulnerabilityByIdRequest) (*Vulnerability, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilityById not implemented")
}
func (UnimplementedVulnerabilitiesServiceServer) mustEmbedUnimplementedVulnerabilitiesServiceServer() {
}

// UnsafeVulnerabilitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VulnerabilitiesServiceServer will
// result in compilation errors.
type UnsafeVulnerabilitiesServiceServer interface {
	mustEmbedUnimplementedVulnerabilitiesServiceServer()
}

func RegisterVulnerabilitiesServiceServer(s grpc.ServiceRegistrar, srv VulnerabilitiesServiceServer) {
	s.RegisterService(&VulnerabilitiesService_ServiceDesc, srv)
}

func _VulnerabilitiesService_GetVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVulnerabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilitiesServiceServer).GetVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.VulnerabilitiesService/GetVulnerabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilitiesServiceServer).GetVulnerabilities(ctx, req.(*GetVulnerabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VulnerabilitiesService_GetVulnerabilityById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVulnerabilityByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilitiesServiceServer).GetVulnerabilityById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.VulnerabilitiesService/GetVulnerabilityById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilitiesServiceServer).GetVulnerabilityById(ctx, req.(*GetVulnerabilityByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VulnerabilitiesService_ServiceDesc is the grpc.ServiceDesc for VulnerabilitiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VulnerabilitiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.VulnerabilitiesService",
	HandlerType: (*VulnerabilitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVulnerabilities",
			Handler:    _VulnerabilitiesService_GetVulnerabilities_Handler,
		},
		{
			MethodName: "GetVulnerabilityById",
			Handler:    _VulnerabilitiesService_GetVulnerabilityById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// SecretsServiceClient is the client API for SecretsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretsServiceClient interface {
	GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error)
	GetSecretById(ctx context.Context, in *GetSecretByIdRequest, opts ...grpc.CallOption) (*Secret, error)
}

type secretsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsServiceClient(cc grpc.ClientConnInterface) SecretsServiceClient {
	return &secretsServiceClient{cc}
}

func (c *secretsServiceClient) GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error) {
	out := new(GetSecretsResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.SecretsService/GetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsServiceClient) GetSecretById(ctx context.Context, in *GetSecretByIdRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.SecretsService/GetSecretById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsServiceServer is the server API for SecretsService service.
// All implementations must embed UnimplementedSecretsServiceServer
// for forward compatibility
type SecretsServiceServer interface {
	GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error)
	GetSecretById(context.Context, *GetSecretByIdRequest) (*Secret, error)
	mustEmbedUnimplementedSecretsServiceServer()
}

// UnimplementedSecretsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretsServiceServer struct {
}

func (UnimplementedSecretsServiceServer) GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedSecretsServiceServer) GetSecretById(context.Context, *GetSecretByIdRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretById not implemented")
}
func (UnimplementedSecretsServiceServer) mustEmbedUnimplementedSecretsServiceServer() {}

// UnsafeSecretsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretsServiceServer will
// result in compilation errors.
type UnsafeSecretsServiceServer interface {
	mustEmbedUnimplementedSecretsServiceServer()
}

func RegisterSecretsServiceServer(s grpc.ServiceRegistrar, srv SecretsServiceServer) {
	s.RegisterService(&SecretsService_ServiceDesc, srv)
}

func _SecretsService_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.SecretsService/GetSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).GetSecrets(ctx, req.(*GetSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsService_GetSecretById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).GetSecretById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.SecretsService/GetSecretById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).GetSecretById(ctx, req.(*GetSecretByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretsService_ServiceDesc is the grpc.ServiceDesc for SecretsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.SecretsService",
	HandlerType: (*SecretsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecrets",
			Handler:    _SecretsService_GetSecrets_Handler,
		},
		{
			MethodName: "GetSecretById",
			Handler:    _SecretsService_GetSecretById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}

// BranchProtectionServiceClient is the client API for BranchProtectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchProtectionServiceClient interface {
	GetBranchProtection(ctx context.Context, in *GetBranchProtectionRequest, opts ...grpc.CallOption) (*GetBranchProtectionResponse, error)
}

type branchProtectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchProtectionServiceClient(cc grpc.ClientConnInterface) BranchProtectionServiceClient {
	return &branchProtectionServiceClient{cc}
}

func (c *branchProtectionServiceClient) GetBranchProtection(ctx context.Context, in *GetBranchProtectionRequest, opts ...grpc.CallOption) (*GetBranchProtectionResponse, error) {
	out := new(GetBranchProtectionResponse)
	err := c.cc.Invoke(ctx, "/dev.stacklok.mediator.v1.BranchProtectionService/GetBranchProtection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchProtectionServiceServer is the server API for BranchProtectionService service.
// All implementations must embed UnimplementedBranchProtectionServiceServer
// for forward compatibility
type BranchProtectionServiceServer interface {
	GetBranchProtection(context.Context, *GetBranchProtectionRequest) (*GetBranchProtectionResponse, error)
	mustEmbedUnimplementedBranchProtectionServiceServer()
}

// UnimplementedBranchProtectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBranchProtectionServiceServer struct {
}

func (UnimplementedBranchProtectionServiceServer) GetBranchProtection(context.Context, *GetBranchProtectionRequest) (*GetBranchProtectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchProtection not implemented")
}
func (UnimplementedBranchProtectionServiceServer) mustEmbedUnimplementedBranchProtectionServiceServer() {
}

// UnsafeBranchProtectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchProtectionServiceServer will
// result in compilation errors.
type UnsafeBranchProtectionServiceServer interface {
	mustEmbedUnimplementedBranchProtectionServiceServer()
}

func RegisterBranchProtectionServiceServer(s grpc.ServiceRegistrar, srv BranchProtectionServiceServer) {
	s.RegisterService(&BranchProtectionService_ServiceDesc, srv)
}

func _BranchProtectionService_GetBranchProtection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchProtectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchProtectionServiceServer).GetBranchProtection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.stacklok.mediator.v1.BranchProtectionService/GetBranchProtection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchProtectionServiceServer).GetBranchProtection(ctx, req.(*GetBranchProtectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchProtectionService_ServiceDesc is the grpc.ServiceDesc for BranchProtectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchProtectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dev.stacklok.mediator.v1.BranchProtectionService",
	HandlerType: (*BranchProtectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBranchProtection",
			Handler:    _BranchProtectionService_GetBranchProtection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/mediator.proto",
}
