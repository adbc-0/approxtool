// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	approx "apisrv/protos/approx"
	auth "apisrv/protos/auth"
	model "apisrv/protos/model"
	user "apisrv/protos/user"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	// Auth Service
	Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error)
	RefreshToken(ctx context.Context, in *auth.RefreshRequest, opts ...grpc.CallOption) (*auth.RefreshResponse, error)
	Logout(ctx context.Context, in *auth.LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// User Service
	ChangeUserPrivilege(ctx context.Context, in *user.ChangePrivilegeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUser(ctx context.Context, in *user.NewUserRequest, opts ...grpc.CallOption) (*user.UserResponse, error)
	DeleteUser(ctx context.Context, in *user.DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllUsers(ctx context.Context, in *user.GetAllUsersRequest, opts ...grpc.CallOption) (*user.GetUsersResponse, error)
	SearchForUsers(ctx context.Context, in *user.SearchRequest, opts ...grpc.CallOption) (*user.SearchResponse, error)
	ChangePassword(ctx context.Context, in *user.ChangePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Model Service
	AddModel(ctx context.Context, in *model.NewModelRequest, opts ...grpc.CallOption) (*model.NewModelResponse, error)
	DeleteModel(ctx context.Context, in *model.DeleteModelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserModels(ctx context.Context, in *model.GetModelsRequest, opts ...grpc.CallOption) (*model.GetModelsResponse, error)
	// Approx Service
	FitCurves(ctx context.Context, in *approx.CurveFitRequest, opts ...grpc.CallOption) (*approx.CurveFitResult, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
	out := new(auth.LoginResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RefreshToken(ctx context.Context, in *auth.RefreshRequest, opts ...grpc.CallOption) (*auth.RefreshResponse, error) {
	out := new(auth.RefreshResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) Logout(ctx context.Context, in *auth.LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.ApiService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ChangeUserPrivilege(ctx context.Context, in *user.ChangePrivilegeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.ApiService/ChangeUserPrivilege", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CreateUser(ctx context.Context, in *user.NewUserRequest, opts ...grpc.CallOption) (*user.UserResponse, error) {
	out := new(user.UserResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteUser(ctx context.Context, in *user.DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.ApiService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetAllUsers(ctx context.Context, in *user.GetAllUsersRequest, opts ...grpc.CallOption) (*user.GetUsersResponse, error) {
	out := new(user.GetUsersResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SearchForUsers(ctx context.Context, in *user.SearchRequest, opts ...grpc.CallOption) (*user.SearchResponse, error) {
	out := new(user.SearchResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/SearchForUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ChangePassword(ctx context.Context, in *user.ChangePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.ApiService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddModel(ctx context.Context, in *model.NewModelRequest, opts ...grpc.CallOption) (*model.NewModelResponse, error) {
	out := new(model.NewModelResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/AddModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) DeleteModel(ctx context.Context, in *model.DeleteModelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/protos.ApiService/DeleteModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUserModels(ctx context.Context, in *model.GetModelsRequest, opts ...grpc.CallOption) (*model.GetModelsResponse, error) {
	out := new(model.GetModelsResponse)
	err := c.cc.Invoke(ctx, "/protos.ApiService/GetUserModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) FitCurves(ctx context.Context, in *approx.CurveFitRequest, opts ...grpc.CallOption) (*approx.CurveFitResult, error) {
	out := new(approx.CurveFitResult)
	err := c.cc.Invoke(ctx, "/protos.ApiService/FitCurves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	// Auth Service
	Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error)
	RefreshToken(context.Context, *auth.RefreshRequest) (*auth.RefreshResponse, error)
	Logout(context.Context, *auth.LogoutRequest) (*emptypb.Empty, error)
	// User Service
	ChangeUserPrivilege(context.Context, *user.ChangePrivilegeRequest) (*emptypb.Empty, error)
	CreateUser(context.Context, *user.NewUserRequest) (*user.UserResponse, error)
	DeleteUser(context.Context, *user.DeleteUserRequest) (*emptypb.Empty, error)
	GetAllUsers(context.Context, *user.GetAllUsersRequest) (*user.GetUsersResponse, error)
	SearchForUsers(context.Context, *user.SearchRequest) (*user.SearchResponse, error)
	ChangePassword(context.Context, *user.ChangePasswordRequest) (*emptypb.Empty, error)
	// Model Service
	AddModel(context.Context, *model.NewModelRequest) (*model.NewModelResponse, error)
	DeleteModel(context.Context, *model.DeleteModelRequest) (*emptypb.Empty, error)
	GetUserModels(context.Context, *model.GetModelsRequest) (*model.GetModelsResponse, error)
	// Approx Service
	FitCurves(context.Context, *approx.CurveFitRequest) (*approx.CurveFitResult, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedApiServiceServer) RefreshToken(context.Context, *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedApiServiceServer) Logout(context.Context, *auth.LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedApiServiceServer) ChangeUserPrivilege(context.Context, *user.ChangePrivilegeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPrivilege not implemented")
}
func (UnimplementedApiServiceServer) CreateUser(context.Context, *user.NewUserRequest) (*user.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedApiServiceServer) DeleteUser(context.Context, *user.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedApiServiceServer) GetAllUsers(context.Context, *user.GetAllUsersRequest) (*user.GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedApiServiceServer) SearchForUsers(context.Context, *user.SearchRequest) (*user.SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForUsers not implemented")
}
func (UnimplementedApiServiceServer) ChangePassword(context.Context, *user.ChangePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedApiServiceServer) AddModel(context.Context, *model.NewModelRequest) (*model.NewModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddModel not implemented")
}
func (UnimplementedApiServiceServer) DeleteModel(context.Context, *model.DeleteModelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModel not implemented")
}
func (UnimplementedApiServiceServer) GetUserModels(context.Context, *model.GetModelsRequest) (*model.GetModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserModels not implemented")
}
func (UnimplementedApiServiceServer) FitCurves(context.Context, *approx.CurveFitRequest) (*approx.CurveFitResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FitCurves not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Login(ctx, req.(*auth.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RefreshToken(ctx, req.(*auth.RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Logout(ctx, req.(*auth.LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ChangeUserPrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.ChangePrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ChangeUserPrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/ChangeUserPrivilege",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ChangeUserPrivilege(ctx, req.(*user.ChangePrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateUser(ctx, req.(*user.NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteUser(ctx, req.(*user.DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetAllUsers(ctx, req.(*user.GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_SearchForUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).SearchForUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/SearchForUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).SearchForUsers(ctx, req.(*user.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).ChangePassword(ctx, req.(*user.ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AddModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.NewModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AddModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/AddModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AddModel(ctx, req.(*model.NewModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_DeleteModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.DeleteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).DeleteModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/DeleteModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).DeleteModel(ctx, req.(*model.DeleteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUserModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUserModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/GetUserModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUserModels(ctx, req.(*model.GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_FitCurves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(approx.CurveFitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).FitCurves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ApiService/FitCurves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).FitCurves(ctx, req.(*approx.CurveFitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ApiService_Login_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _ApiService_RefreshToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ApiService_Logout_Handler,
		},
		{
			MethodName: "ChangeUserPrivilege",
			Handler:    _ApiService_ChangeUserPrivilege_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ApiService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ApiService_DeleteUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _ApiService_GetAllUsers_Handler,
		},
		{
			MethodName: "SearchForUsers",
			Handler:    _ApiService_SearchForUsers_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _ApiService_ChangePassword_Handler,
		},
		{
			MethodName: "AddModel",
			Handler:    _ApiService_AddModel_Handler,
		},
		{
			MethodName: "DeleteModel",
			Handler:    _ApiService_DeleteModel_Handler,
		},
		{
			MethodName: "GetUserModels",
			Handler:    _ApiService_GetUserModels_Handler,
		},
		{
			MethodName: "FitCurves",
			Handler:    _ApiService_FitCurves_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiservice.proto",
}
