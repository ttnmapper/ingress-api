// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: lorawan-stack/api/configuration_services.proto

package ttnpb

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

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationClient interface {
	ListFrequencyPlans(ctx context.Context, in *ListFrequencyPlansRequest, opts ...grpc.CallOption) (*ListFrequencyPlansResponse, error)
	// Returns a list of supported LoRaWAN PHY Versions for the given Band ID.
	GetPhyVersions(ctx context.Context, in *GetPhyVersionsRequest, opts ...grpc.CallOption) (*GetPhyVersionsResponse, error)
	ListBands(ctx context.Context, in *ListBandsRequest, opts ...grpc.CallOption) (*ListBandsResponse, error)
}

type configurationClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationClient(cc grpc.ClientConnInterface) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) ListFrequencyPlans(ctx context.Context, in *ListFrequencyPlansRequest, opts ...grpc.CallOption) (*ListFrequencyPlansResponse, error) {
	out := new(ListFrequencyPlansResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Configuration/ListFrequencyPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetPhyVersions(ctx context.Context, in *GetPhyVersionsRequest, opts ...grpc.CallOption) (*GetPhyVersionsResponse, error) {
	out := new(GetPhyVersionsResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Configuration/GetPhyVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) ListBands(ctx context.Context, in *ListBandsRequest, opts ...grpc.CallOption) (*ListBandsResponse, error) {
	out := new(ListBandsResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Configuration/ListBands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
// All implementations must embed UnimplementedConfigurationServer
// for forward compatibility
type ConfigurationServer interface {
	ListFrequencyPlans(context.Context, *ListFrequencyPlansRequest) (*ListFrequencyPlansResponse, error)
	// Returns a list of supported LoRaWAN PHY Versions for the given Band ID.
	GetPhyVersions(context.Context, *GetPhyVersionsRequest) (*GetPhyVersionsResponse, error)
	ListBands(context.Context, *ListBandsRequest) (*ListBandsResponse, error)
	mustEmbedUnimplementedConfigurationServer()
}

// UnimplementedConfigurationServer must be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (UnimplementedConfigurationServer) ListFrequencyPlans(context.Context, *ListFrequencyPlansRequest) (*ListFrequencyPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFrequencyPlans not implemented")
}
func (UnimplementedConfigurationServer) GetPhyVersions(context.Context, *GetPhyVersionsRequest) (*GetPhyVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhyVersions not implemented")
}
func (UnimplementedConfigurationServer) ListBands(context.Context, *ListBandsRequest) (*ListBandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBands not implemented")
}
func (UnimplementedConfigurationServer) mustEmbedUnimplementedConfigurationServer() {}

// UnsafeConfigurationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationServer will
// result in compilation errors.
type UnsafeConfigurationServer interface {
	mustEmbedUnimplementedConfigurationServer()
}

func RegisterConfigurationServer(s grpc.ServiceRegistrar, srv ConfigurationServer) {
	s.RegisterService(&Configuration_ServiceDesc, srv)
}

func _Configuration_ListFrequencyPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFrequencyPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).ListFrequencyPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Configuration/ListFrequencyPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).ListFrequencyPlans(ctx, req.(*ListFrequencyPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetPhyVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhyVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetPhyVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Configuration/GetPhyVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetPhyVersions(ctx, req.(*GetPhyVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_ListBands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).ListBands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Configuration/ListBands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).ListBands(ctx, req.(*ListBandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Configuration_ServiceDesc is the grpc.ServiceDesc for Configuration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Configuration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFrequencyPlans",
			Handler:    _Configuration_ListFrequencyPlans_Handler,
		},
		{
			MethodName: "GetPhyVersions",
			Handler:    _Configuration_GetPhyVersions_Handler,
		},
		{
			MethodName: "ListBands",
			Handler:    _Configuration_ListBands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/configuration_services.proto",
}
