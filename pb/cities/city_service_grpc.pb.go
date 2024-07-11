// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: city_service.proto

package cities

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CitiesService_GetCitylocal_FullMethodName = "/cities.CitiesService/GetCitylocal"
	CitiesService_GetCity_FullMethodName      = "/cities.CitiesService/GetCity"
	CitiesService_Create_FullMethodName       = "/cities.CitiesService/Create"
	CitiesService_Update_FullMethodName       = "/cities.CitiesService/Update"
)

// CitiesServiceClient is the client API for CitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CitiesServiceClient interface {
	GetCitylocal(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error)
	GetCity(ctx context.Context, in *Id, opts ...grpc.CallOption) (*City, error)
	Create(ctx context.Context, in *CityInput, opts ...grpc.CallOption) (*City, error)
	Update(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error)
}

type citiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCitiesServiceClient(cc grpc.ClientConnInterface) CitiesServiceClient {
	return &citiesServiceClient{cc}
}

func (c *citiesServiceClient) GetCitylocal(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(City)
	err := c.cc.Invoke(ctx, CitiesService_GetCitylocal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) GetCity(ctx context.Context, in *Id, opts ...grpc.CallOption) (*City, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(City)
	err := c.cc.Invoke(ctx, CitiesService_GetCity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) Create(ctx context.Context, in *CityInput, opts ...grpc.CallOption) (*City, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(City)
	err := c.cc.Invoke(ctx, CitiesService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citiesServiceClient) Update(ctx context.Context, in *City, opts ...grpc.CallOption) (*City, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(City)
	err := c.cc.Invoke(ctx, CitiesService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CitiesServiceServer is the server API for CitiesService service.
// All implementations should embed UnimplementedCitiesServiceServer
// for forward compatibility
type CitiesServiceServer interface {
	GetCitylocal(context.Context, *City) (*City, error)
	GetCity(context.Context, *Id) (*City, error)
	Create(context.Context, *CityInput) (*City, error)
	Update(context.Context, *City) (*City, error)
}

// UnimplementedCitiesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCitiesServiceServer struct {
}

func (UnimplementedCitiesServiceServer) GetCitylocal(context.Context, *City) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCitylocal not implemented")
}
func (UnimplementedCitiesServiceServer) GetCity(context.Context, *Id) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (UnimplementedCitiesServiceServer) Create(context.Context, *CityInput) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCitiesServiceServer) Update(context.Context, *City) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeCitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CitiesServiceServer will
// result in compilation errors.
type UnsafeCitiesServiceServer interface {
	mustEmbedUnimplementedCitiesServiceServer()
}

func RegisterCitiesServiceServer(s grpc.ServiceRegistrar, srv CitiesServiceServer) {
	s.RegisterService(&CitiesService_ServiceDesc, srv)
}

func _CitiesService_GetCitylocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).GetCitylocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitiesService_GetCitylocal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).GetCitylocal(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitiesService_GetCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).GetCity(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitiesService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).Create(ctx, req.(*CityInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitiesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitiesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitiesService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitiesServiceServer).Update(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

// CitiesService_ServiceDesc is the grpc.ServiceDesc for CitiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CitiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cities.CitiesService",
	HandlerType: (*CitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCitylocal",
			Handler:    _CitiesService_GetCitylocal_Handler,
		},
		{
			MethodName: "GetCity",
			Handler:    _CitiesService_GetCity_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CitiesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CitiesService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city_service.proto",
}
