// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/oauth_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/oauth_services.proto", fileDescriptor_10930ff381158870)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/oauth_services.proto", fileDescriptor_10930ff381158870)
}

var fileDescriptor_10930ff381158870 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbb, 0x6e, 0xdb, 0x30,
	0x14, 0x86, 0xa1, 0xa2, 0xf5, 0xa0, 0x02, 0x1d, 0x38, 0x74, 0x50, 0x2f, 0x40, 0x3d, 0x14, 0xad,
	0x5b, 0x93, 0x45, 0xd5, 0xce, 0x85, 0x7b, 0x2f, 0x7a, 0x03, 0xdc, 0x4e, 0x5d, 0x0c, 0x4a, 0x3e,
	0xa6, 0x08, 0xc9, 0xa4, 0x4a, 0x1e, 0xd9, 0x70, 0x0c, 0x2f, 0x79, 0x85, 0x4c, 0x59, 0xb2, 0x66,
	0xca, 0x94, 0x17, 0xc8, 0x3b, 0xe4, 0x11, 0x92, 0x07, 0x09, 0x44, 0xc9, 0x46, 0x14, 0xc4, 0x88,
	0x81, 0x64, 0x10, 0x74, 0xa8, 0xf3, 0xff, 0x07, 0x1f, 0xff, 0x23, 0xff, 0x69, 0xa6, 0x0d, 0x9f,
	0x72, 0xd5, 0xb5, 0xc8, 0xe3, 0x94, 0xf1, 0x5c, 0x32, 0xcd, 0x0b, 0x4c, 0x06, 0x16, 0xcc, 0x44,
	0xc6, 0x60, 0x69, 0x6e, 0x34, 0x6a, 0x72, 0x0f, 0x51, 0xd1, 0x5a, 0x4b, 0x27, 0x61, 0xd0, 0x15,
	0x12, 0x93, 0x22, 0xa2, 0xb1, 0x1e, 0x33, 0xa1, 0x85, 0x66, 0x4e, 0x16, 0x15, 0x23, 0x77, 0x72,
	0x07, 0x57, 0x55, 0xf6, 0xe0, 0xa1, 0xd0, 0x5a, 0x64, 0xe0, 0xe6, 0x73, 0xa5, 0x34, 0x72, 0x94,
	0x5a, 0xd5, 0xc3, 0x83, 0x07, 0x75, 0x77, 0x35, 0x03, 0xc6, 0x39, 0xce, 0xea, 0xe6, 0xa3, 0x35,
	0x84, 0x55, 0xfb, 0xf5, 0xe1, 0x1d, 0x3f, 0xf8, 0xdd, 0x2b, 0x30, 0x29, 0x1f, 0x6d, 0xe4, 0x96,
	0x9b, 0xdc, 0x07, 0x21, 0x2d, 0x9a, 0x19, 0xd9, 0xf5, 0xfc, 0xdb, 0x3f, 0xa4, 0x45, 0xf2, 0x86,
	0x36, 0x6f, 0x40, 0xcb, 0xaf, 0xce, 0xf8, 0x21, 0x93, 0xa0, 0xb0, 0x61, 0xb7, 0x7d, 0xf8, 0x5f,
	0x80, 0xc5, 0xe0, 0xf9, 0x45, 0xd7, 0x5a, 0x47, 0xfb, 0xd5, 0xf6, 0xf1, 0xe9, 0xce, 0xad, 0x0e,
	0x79, 0xc6, 0x0a, 0x0b, 0xc6, 0xb2, 0x79, 0xf9, 0x1a, 0xc8, 0xa1, 0xa5, 0x75, 0xb1, 0x60, 0xbc,
	0xe1, 0x20, 0x07, 0x9e, 0xef, 0x97, 0x14, 0x7f, 0x75, 0x0a, 0xca, 0x92, 0x97, 0x6b, 0x09, 0x7b,
	0x71, 0x0c, 0xd6, 0x56, 0xb2, 0x25, 0xd9, 0x93, 0x4b, 0xc9, 0xce, 0x2b, 0xdb, 0xbf, 0x1c, 0xd1,
	0x57, 0xf2, 0x79, 0x53, 0x22, 0x36, 0x8f, 0xdd, 0xcd, 0x9c, 0x66, 0x55, 0x2e, 0x18, 0x56, 0x80,
	0x7b, 0x9e, 0xdf, 0xfa, 0x08, 0x19, 0x20, 0x90, 0x70, 0xd3, 0x5c, 0xbe, 0x0d, 0x41, 0xa1, 0x1c,
	0x49, 0x30, 0x36, 0xb8, 0x4f, 0xab, 0x3d, 0xd3, 0xe5, 0x9e, 0xe9, 0xa7, 0x72, 0xcf, 0xed, 0x2f,
	0x8e, 0xb3, 0xd7, 0x79, 0x77, 0x4d, 0x4e, 0xb2, 0xef, 0xf9, 0x77, 0x2b, 0x40, 0x97, 0x00, 0x79,
	0x71, 0x55, 0x46, 0x9b, 0xd0, 0xfd, 0x71, 0x74, 0x3f, 0x3b, 0xdf, 0x6f, 0x26, 0x45, 0x36, 0x97,
	0xc3, 0xc5, 0xfb, 0xb7, 0x47, 0x27, 0x8f, 0xbd, 0x7f, 0x4c, 0x68, 0x8a, 0x09, 0x60, 0x22, 0x95,
	0xb0, 0x54, 0x01, 0x4e, 0xb5, 0x49, 0x59, 0xf3, 0x7f, 0x9f, 0x84, 0x2c, 0x4f, 0x05, 0x43, 0x54,
	0x79, 0x14, 0xb5, 0x1c, 0x5b, 0x78, 0x16, 0x00, 0x00, 0xff, 0xff, 0x81, 0xd5, 0xb4, 0xaf, 0xb6,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OAuthAuthorizationRegistryClient is the client API for OAuthAuthorizationRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OAuthAuthorizationRegistryClient interface {
	// List OAuth clients that are authorized by the user.
	List(ctx context.Context, in *ListOAuthClientAuthorizationsRequest, opts ...grpc.CallOption) (*OAuthClientAuthorizations, error)
	// List OAuth access tokens issued to the OAuth client on behalf of the user.
	ListTokens(ctx context.Context, in *ListOAuthAccessTokensRequest, opts ...grpc.CallOption) (*OAuthAccessTokens, error)
	// Delete (de-authorize) an OAuth client for the user.
	Delete(ctx context.Context, in *OAuthClientAuthorizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
	// Delete (invalidate) an OAuth access token.
	DeleteToken(ctx context.Context, in *OAuthAccessTokenIdentifiers, opts ...grpc.CallOption) (*types.Empty, error)
}

type oAuthAuthorizationRegistryClient struct {
	cc *grpc.ClientConn
}

func NewOAuthAuthorizationRegistryClient(cc *grpc.ClientConn) OAuthAuthorizationRegistryClient {
	return &oAuthAuthorizationRegistryClient{cc}
}

func (c *oAuthAuthorizationRegistryClient) List(ctx context.Context, in *ListOAuthClientAuthorizationsRequest, opts ...grpc.CallOption) (*OAuthClientAuthorizations, error) {
	out := new(OAuthClientAuthorizations)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OAuthAuthorizationRegistry/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthAuthorizationRegistryClient) ListTokens(ctx context.Context, in *ListOAuthAccessTokensRequest, opts ...grpc.CallOption) (*OAuthAccessTokens, error) {
	out := new(OAuthAccessTokens)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OAuthAuthorizationRegistry/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthAuthorizationRegistryClient) Delete(ctx context.Context, in *OAuthClientAuthorizationIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OAuthAuthorizationRegistry/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthAuthorizationRegistryClient) DeleteToken(ctx context.Context, in *OAuthAccessTokenIdentifiers, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.OAuthAuthorizationRegistry/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthAuthorizationRegistryServer is the server API for OAuthAuthorizationRegistry service.
type OAuthAuthorizationRegistryServer interface {
	// List OAuth clients that are authorized by the user.
	List(context.Context, *ListOAuthClientAuthorizationsRequest) (*OAuthClientAuthorizations, error)
	// List OAuth access tokens issued to the OAuth client on behalf of the user.
	ListTokens(context.Context, *ListOAuthAccessTokensRequest) (*OAuthAccessTokens, error)
	// Delete (de-authorize) an OAuth client for the user.
	Delete(context.Context, *OAuthClientAuthorizationIdentifiers) (*types.Empty, error)
	// Delete (invalidate) an OAuth access token.
	DeleteToken(context.Context, *OAuthAccessTokenIdentifiers) (*types.Empty, error)
}

// UnimplementedOAuthAuthorizationRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedOAuthAuthorizationRegistryServer struct {
}

func (*UnimplementedOAuthAuthorizationRegistryServer) List(ctx context.Context, req *ListOAuthClientAuthorizationsRequest) (*OAuthClientAuthorizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedOAuthAuthorizationRegistryServer) ListTokens(ctx context.Context, req *ListOAuthAccessTokensRequest) (*OAuthAccessTokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}
func (*UnimplementedOAuthAuthorizationRegistryServer) Delete(ctx context.Context, req *OAuthClientAuthorizationIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedOAuthAuthorizationRegistryServer) DeleteToken(ctx context.Context, req *OAuthAccessTokenIdentifiers) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}

func RegisterOAuthAuthorizationRegistryServer(s *grpc.Server, srv OAuthAuthorizationRegistryServer) {
	s.RegisterService(&_OAuthAuthorizationRegistry_serviceDesc, srv)
}

func _OAuthAuthorizationRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOAuthClientAuthorizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthAuthorizationRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OAuthAuthorizationRegistry/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthAuthorizationRegistryServer).List(ctx, req.(*ListOAuthClientAuthorizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthAuthorizationRegistry_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOAuthAccessTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthAuthorizationRegistryServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OAuthAuthorizationRegistry/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthAuthorizationRegistryServer).ListTokens(ctx, req.(*ListOAuthAccessTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthAuthorizationRegistry_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthClientAuthorizationIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthAuthorizationRegistryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OAuthAuthorizationRegistry/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthAuthorizationRegistryServer).Delete(ctx, req.(*OAuthClientAuthorizationIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthAuthorizationRegistry_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthAccessTokenIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthAuthorizationRegistryServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.OAuthAuthorizationRegistry/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthAuthorizationRegistryServer).DeleteToken(ctx, req.(*OAuthAccessTokenIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _OAuthAuthorizationRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.OAuthAuthorizationRegistry",
	HandlerType: (*OAuthAuthorizationRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _OAuthAuthorizationRegistry_List_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _OAuthAuthorizationRegistry_ListTokens_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OAuthAuthorizationRegistry_Delete_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _OAuthAuthorizationRegistry_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/oauth_services.proto",
}
