// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gnomobiletypes.proto

package gnomobiletypes

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

const (
	GnomobileService_SetRemote_FullMethodName       = "/gnomobile.v1.GnomobileService/SetRemote"
	GnomobileService_SetChainID_FullMethodName      = "/gnomobile.v1.GnomobileService/SetChainID"
	GnomobileService_SetNameOrBech32_FullMethodName = "/gnomobile.v1.GnomobileService/SetNameOrBech32"
	GnomobileService_SetPassword_FullMethodName     = "/gnomobile.v1.GnomobileService/SetPassword"
	GnomobileService_ListKeyInfo_FullMethodName     = "/gnomobile.v1.GnomobileService/ListKeyInfo"
	GnomobileService_CreateAccount_FullMethodName   = "/gnomobile.v1.GnomobileService/CreateAccount"
	GnomobileService_SelectAccount_FullMethodName   = "/gnomobile.v1.GnomobileService/SelectAccount"
	GnomobileService_Query_FullMethodName           = "/gnomobile.v1.GnomobileService/Query"
	GnomobileService_Call_FullMethodName            = "/gnomobile.v1.GnomobileService/Call"
)

// GnomobileServiceClient is the client API for GnomobileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GnomobileServiceClient interface {
	// Set the connection addresse for the remote node. If you don't call this,
	// the default is "127.0.0.1:26657"
	SetRemote(ctx context.Context, in *SetRemote_Request, opts ...grpc.CallOption) (*SetRemote_Reply, error)
	// Set the chain ID for the remote node. If you don't call this, the default
	// is "dev"
	SetChainID(ctx context.Context, in *SetChainID_Request, opts ...grpc.CallOption) (*SetChainID_Reply, error)
	// Set the nameOrBech32 for the account in the keybase, used for later
	// operations
	SetNameOrBech32(ctx context.Context, in *SetNameOrBech32_Request, opts ...grpc.CallOption) (*SetNameOrBech32_Reply, error)
	// Set the password for the account in the keybase, used for later operations
	SetPassword(ctx context.Context, in *SetPassword_Request, opts ...grpc.CallOption) (*SetPassword_Reply, error)
	// Get the keys informations in the keybase
	ListKeyInfo(ctx context.Context, in *ListKeyInfo_Request, opts ...grpc.CallOption) (*ListKeyInfo_Reply, error)
	// Create a new account the keybase using the name an password specified by
	// SetAccount
	CreateAccount(ctx context.Context, in *CreateAccount_Request, opts ...grpc.CallOption) (*CreateAccount_Reply, error)
	// SelectAccount selects the account to use for later operations
	SelectAccount(ctx context.Context, in *SelectAccount_Request, opts ...grpc.CallOption) (*SelectAccount_Reply, error)
	// Make an ABCI query to the remote node.
	Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (*Query_Reply, error)
	// Call a specific realm function.
	Call(ctx context.Context, in *Call_Request, opts ...grpc.CallOption) (*Call_Reply, error)
}

type gnomobileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGnomobileServiceClient(cc grpc.ClientConnInterface) GnomobileServiceClient {
	return &gnomobileServiceClient{cc}
}

func (c *gnomobileServiceClient) SetRemote(ctx context.Context, in *SetRemote_Request, opts ...grpc.CallOption) (*SetRemote_Reply, error) {
	out := new(SetRemote_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_SetRemote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) SetChainID(ctx context.Context, in *SetChainID_Request, opts ...grpc.CallOption) (*SetChainID_Reply, error) {
	out := new(SetChainID_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_SetChainID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) SetNameOrBech32(ctx context.Context, in *SetNameOrBech32_Request, opts ...grpc.CallOption) (*SetNameOrBech32_Reply, error) {
	out := new(SetNameOrBech32_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_SetNameOrBech32_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) SetPassword(ctx context.Context, in *SetPassword_Request, opts ...grpc.CallOption) (*SetPassword_Reply, error) {
	out := new(SetPassword_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_SetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) ListKeyInfo(ctx context.Context, in *ListKeyInfo_Request, opts ...grpc.CallOption) (*ListKeyInfo_Reply, error) {
	out := new(ListKeyInfo_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_ListKeyInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) CreateAccount(ctx context.Context, in *CreateAccount_Request, opts ...grpc.CallOption) (*CreateAccount_Reply, error) {
	out := new(CreateAccount_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) SelectAccount(ctx context.Context, in *SelectAccount_Request, opts ...grpc.CallOption) (*SelectAccount_Reply, error) {
	out := new(SelectAccount_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_SelectAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (*Query_Reply, error) {
	out := new(Query_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) Call(ctx context.Context, in *Call_Request, opts ...grpc.CallOption) (*Call_Reply, error) {
	out := new(Call_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GnomobileServiceServer is the server API for GnomobileService service.
// All implementations must embed UnimplementedGnomobileServiceServer
// for forward compatibility
type GnomobileServiceServer interface {
	// Set the connection addresse for the remote node. If you don't call this,
	// the default is "127.0.0.1:26657"
	SetRemote(context.Context, *SetRemote_Request) (*SetRemote_Reply, error)
	// Set the chain ID for the remote node. If you don't call this, the default
	// is "dev"
	SetChainID(context.Context, *SetChainID_Request) (*SetChainID_Reply, error)
	// Set the nameOrBech32 for the account in the keybase, used for later
	// operations
	SetNameOrBech32(context.Context, *SetNameOrBech32_Request) (*SetNameOrBech32_Reply, error)
	// Set the password for the account in the keybase, used for later operations
	SetPassword(context.Context, *SetPassword_Request) (*SetPassword_Reply, error)
	// Get the keys informations in the keybase
	ListKeyInfo(context.Context, *ListKeyInfo_Request) (*ListKeyInfo_Reply, error)
	// Create a new account the keybase using the name an password specified by
	// SetAccount
	CreateAccount(context.Context, *CreateAccount_Request) (*CreateAccount_Reply, error)
	// SelectAccount selects the account to use for later operations
	SelectAccount(context.Context, *SelectAccount_Request) (*SelectAccount_Reply, error)
	// Make an ABCI query to the remote node.
	Query(context.Context, *Query_Request) (*Query_Reply, error)
	// Call a specific realm function.
	Call(context.Context, *Call_Request) (*Call_Reply, error)
	mustEmbedUnimplementedGnomobileServiceServer()
}

// UnimplementedGnomobileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGnomobileServiceServer struct {
}

func (UnimplementedGnomobileServiceServer) SetRemote(context.Context, *SetRemote_Request) (*SetRemote_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRemote not implemented")
}
func (UnimplementedGnomobileServiceServer) SetChainID(context.Context, *SetChainID_Request) (*SetChainID_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChainID not implemented")
}
func (UnimplementedGnomobileServiceServer) SetNameOrBech32(context.Context, *SetNameOrBech32_Request) (*SetNameOrBech32_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNameOrBech32 not implemented")
}
func (UnimplementedGnomobileServiceServer) SetPassword(context.Context, *SetPassword_Request) (*SetPassword_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPassword not implemented")
}
func (UnimplementedGnomobileServiceServer) ListKeyInfo(context.Context, *ListKeyInfo_Request) (*ListKeyInfo_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeyInfo not implemented")
}
func (UnimplementedGnomobileServiceServer) CreateAccount(context.Context, *CreateAccount_Request) (*CreateAccount_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedGnomobileServiceServer) SelectAccount(context.Context, *SelectAccount_Request) (*SelectAccount_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectAccount not implemented")
}
func (UnimplementedGnomobileServiceServer) Query(context.Context, *Query_Request) (*Query_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedGnomobileServiceServer) Call(context.Context, *Call_Request) (*Call_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedGnomobileServiceServer) mustEmbedUnimplementedGnomobileServiceServer() {}

// UnsafeGnomobileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GnomobileServiceServer will
// result in compilation errors.
type UnsafeGnomobileServiceServer interface {
	mustEmbedUnimplementedGnomobileServiceServer()
}

func RegisterGnomobileServiceServer(s grpc.ServiceRegistrar, srv GnomobileServiceServer) {
	s.RegisterService(&GnomobileService_ServiceDesc, srv)
}

func _GnomobileService_SetRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRemote_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).SetRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_SetRemote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).SetRemote(ctx, req.(*SetRemote_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_SetChainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetChainID_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).SetChainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_SetChainID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).SetChainID(ctx, req.(*SetChainID_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_SetNameOrBech32_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNameOrBech32_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).SetNameOrBech32(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_SetNameOrBech32_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).SetNameOrBech32(ctx, req.(*SetNameOrBech32_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPassword_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_SetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).SetPassword(ctx, req.(*SetPassword_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_ListKeyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyInfo_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).ListKeyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_ListKeyInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).ListKeyInfo(ctx, req.(*ListKeyInfo_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccount_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).CreateAccount(ctx, req.(*CreateAccount_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_SelectAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectAccount_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).SelectAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_SelectAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).SelectAccount(ctx, req.(*SelectAccount_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).Query(ctx, req.(*Query_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Call_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).Call(ctx, req.(*Call_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GnomobileService_ServiceDesc is the grpc.ServiceDesc for GnomobileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GnomobileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnomobile.v1.GnomobileService",
	HandlerType: (*GnomobileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRemote",
			Handler:    _GnomobileService_SetRemote_Handler,
		},
		{
			MethodName: "SetChainID",
			Handler:    _GnomobileService_SetChainID_Handler,
		},
		{
			MethodName: "SetNameOrBech32",
			Handler:    _GnomobileService_SetNameOrBech32_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _GnomobileService_SetPassword_Handler,
		},
		{
			MethodName: "ListKeyInfo",
			Handler:    _GnomobileService_ListKeyInfo_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _GnomobileService_CreateAccount_Handler,
		},
		{
			MethodName: "SelectAccount",
			Handler:    _GnomobileService_SelectAccount_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _GnomobileService_Query_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _GnomobileService_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gnomobiletypes.proto",
}