// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api/v3/auth.proto

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

// SignUpClient is the client API for SignUp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignUpClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
}

type signUpClient struct {
	cc grpc.ClientConnInterface
}

func NewSignUpClient(cc grpc.ClientConnInterface) SignUpClient {
	return &signUpClient{cc}
}

func (c *signUpClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/auth.SignUp/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignUpServer is the server API for SignUp service.
// All implementations must embed UnimplementedSignUpServer
// for forward compatibility
type SignUpServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	mustEmbedUnimplementedSignUpServer()
}

// UnimplementedSignUpServer must be embedded to have forward compatible implementations.
type UnimplementedSignUpServer struct {
}

func (UnimplementedSignUpServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedSignUpServer) mustEmbedUnimplementedSignUpServer() {}

// UnsafeSignUpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignUpServer will
// result in compilation errors.
type UnsafeSignUpServer interface {
	mustEmbedUnimplementedSignUpServer()
}

func RegisterSignUpServer(s grpc.ServiceRegistrar, srv SignUpServer) {
	s.RegisterService(&SignUp_ServiceDesc, srv)
}

func _SignUp_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignUpServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.SignUp/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignUpServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignUp_ServiceDesc is the grpc.ServiceDesc for SignUp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignUp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.SignUp",
	HandlerType: (*SignUpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _SignUp_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}

// LogInClient is the client API for LogIn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogInClient interface {
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
}

type logInClient struct {
	cc grpc.ClientConnInterface
}

func NewLogInClient(cc grpc.ClientConnInterface) LogInClient {
	return &logInClient{cc}
}

func (c *logInClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/auth.LogIn/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogInServer is the server API for LogIn service.
// All implementations must embed UnimplementedLogInServer
// for forward compatibility
type LogInServer interface {
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	mustEmbedUnimplementedLogInServer()
}

// UnimplementedLogInServer must be embedded to have forward compatible implementations.
type UnimplementedLogInServer struct {
}

func (UnimplementedLogInServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedLogInServer) mustEmbedUnimplementedLogInServer() {}

// UnsafeLogInServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogInServer will
// result in compilation errors.
type UnsafeLogInServer interface {
	mustEmbedUnimplementedLogInServer()
}

func RegisterLogInServer(s grpc.ServiceRegistrar, srv LogInServer) {
	s.RegisterService(&LogIn_ServiceDesc, srv)
}

func _LogIn_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogInServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.LogIn/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogInServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogIn_ServiceDesc is the grpc.ServiceDesc for LogIn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogIn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.LogIn",
	HandlerType: (*LogInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIn",
			Handler:    _LogIn_LogIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}

// LogOutClient is the client API for LogOut service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogOutClient interface {
	LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error)
}

type logOutClient struct {
	cc grpc.ClientConnInterface
}

func NewLogOutClient(cc grpc.ClientConnInterface) LogOutClient {
	return &logOutClient{cc}
}

func (c *logOutClient) LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error) {
	out := new(LogOutResponse)
	err := c.cc.Invoke(ctx, "/auth.LogOut/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogOutServer is the server API for LogOut service.
// All implementations must embed UnimplementedLogOutServer
// for forward compatibility
type LogOutServer interface {
	LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error)
	mustEmbedUnimplementedLogOutServer()
}

// UnimplementedLogOutServer must be embedded to have forward compatible implementations.
type UnimplementedLogOutServer struct {
}

func (UnimplementedLogOutServer) LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedLogOutServer) mustEmbedUnimplementedLogOutServer() {}

// UnsafeLogOutServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogOutServer will
// result in compilation errors.
type UnsafeLogOutServer interface {
	mustEmbedUnimplementedLogOutServer()
}

func RegisterLogOutServer(s grpc.ServiceRegistrar, srv LogOutServer) {
	s.RegisterService(&LogOut_ServiceDesc, srv)
}

func _LogOut_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogOutServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.LogOut/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogOutServer).LogOut(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogOut_ServiceDesc is the grpc.ServiceDesc for LogOut service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogOut_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.LogOut",
	HandlerType: (*LogOutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogOut",
			Handler:    _LogOut_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}

// RefreshClient is the client API for Refresh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RefreshClient interface {
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type refreshClient struct {
	cc grpc.ClientConnInterface
}

func NewRefreshClient(cc grpc.ClientConnInterface) RefreshClient {
	return &refreshClient{cc}
}

func (c *refreshClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/auth.Refresh/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefreshServer is the server API for Refresh service.
// All implementations must embed UnimplementedRefreshServer
// for forward compatibility
type RefreshServer interface {
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	mustEmbedUnimplementedRefreshServer()
}

// UnimplementedRefreshServer must be embedded to have forward compatible implementations.
type UnimplementedRefreshServer struct {
}

func (UnimplementedRefreshServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedRefreshServer) mustEmbedUnimplementedRefreshServer() {}

// UnsafeRefreshServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RefreshServer will
// result in compilation errors.
type UnsafeRefreshServer interface {
	mustEmbedUnimplementedRefreshServer()
}

func RegisterRefreshServer(s grpc.ServiceRegistrar, srv RefreshServer) {
	s.RegisterService(&Refresh_ServiceDesc, srv)
}

func _Refresh_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Refresh/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Refresh_ServiceDesc is the grpc.ServiceDesc for Refresh service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Refresh_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Refresh",
	HandlerType: (*RefreshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Refresh",
			Handler:    _Refresh_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}

// RevokeClient is the client API for Revoke service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RevokeClient interface {
	Revoke(ctx context.Context, in *RevokeRequest, opts ...grpc.CallOption) (*RevokeResponse, error)
}

type revokeClient struct {
	cc grpc.ClientConnInterface
}

func NewRevokeClient(cc grpc.ClientConnInterface) RevokeClient {
	return &revokeClient{cc}
}

func (c *revokeClient) Revoke(ctx context.Context, in *RevokeRequest, opts ...grpc.CallOption) (*RevokeResponse, error) {
	out := new(RevokeResponse)
	err := c.cc.Invoke(ctx, "/auth.Revoke/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RevokeServer is the server API for Revoke service.
// All implementations must embed UnimplementedRevokeServer
// for forward compatibility
type RevokeServer interface {
	Revoke(context.Context, *RevokeRequest) (*RevokeResponse, error)
	mustEmbedUnimplementedRevokeServer()
}

// UnimplementedRevokeServer must be embedded to have forward compatible implementations.
type UnimplementedRevokeServer struct {
}

func (UnimplementedRevokeServer) Revoke(context.Context, *RevokeRequest) (*RevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (UnimplementedRevokeServer) mustEmbedUnimplementedRevokeServer() {}

// UnsafeRevokeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RevokeServer will
// result in compilation errors.
type UnsafeRevokeServer interface {
	mustEmbedUnimplementedRevokeServer()
}

func RegisterRevokeServer(s grpc.ServiceRegistrar, srv RevokeServer) {
	s.RegisterService(&Revoke_ServiceDesc, srv)
}

func _Revoke_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokeServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Revoke/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokeServer).Revoke(ctx, req.(*RevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Revoke_ServiceDesc is the grpc.ServiceDesc for Revoke service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Revoke_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Revoke",
	HandlerType: (*RevokeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Revoke",
			Handler:    _Revoke_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}

// AuthMiddlewareClient is the client API for AuthMiddleware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthMiddlewareClient interface {
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authMiddlewareClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthMiddlewareClient(cc grpc.ClientConnInterface) AuthMiddlewareClient {
	return &authMiddlewareClient{cc}
}

func (c *authMiddlewareClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthMiddleware/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthMiddlewareServer is the server API for AuthMiddleware service.
// All implementations must embed UnimplementedAuthMiddlewareServer
// for forward compatibility
type AuthMiddlewareServer interface {
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	mustEmbedUnimplementedAuthMiddlewareServer()
}

// UnimplementedAuthMiddlewareServer must be embedded to have forward compatible implementations.
type UnimplementedAuthMiddlewareServer struct {
}

func (UnimplementedAuthMiddlewareServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthMiddlewareServer) mustEmbedUnimplementedAuthMiddlewareServer() {}

// UnsafeAuthMiddlewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthMiddlewareServer will
// result in compilation errors.
type UnsafeAuthMiddlewareServer interface {
	mustEmbedUnimplementedAuthMiddlewareServer()
}

func RegisterAuthMiddlewareServer(s grpc.ServiceRegistrar, srv AuthMiddlewareServer) {
	s.RegisterService(&AuthMiddleware_ServiceDesc, srv)
}

func _AuthMiddleware_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthMiddlewareServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthMiddleware/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthMiddlewareServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthMiddleware_ServiceDesc is the grpc.ServiceDesc for AuthMiddleware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthMiddleware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthMiddleware",
	HandlerType: (*AuthMiddlewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateToken",
			Handler:    _AuthMiddleware_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/auth.proto",
}
