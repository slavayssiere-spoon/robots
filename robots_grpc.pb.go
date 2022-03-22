// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: robots.proto

package robots

import (
	context "context"
	groups "github.com/slavayssiere-spoon/groups"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RobotsClient is the client API for Robots service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotsClient interface {
	GetAll(ctx context.Context, in *Robots, opts ...grpc.CallOption) (*Robots, error)
	GetGraph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*groups.Groups, error)
	Get(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error)
	GetDirectusToken(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*DirectusToken, error)
	GetMyDirectusToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DirectusToken, error)
	GetByGroup(ctx context.Context, in *groups.Group, opts ...grpc.CallOption) (*Robots, error)
	GetSAFile(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*SaFile, error)
	Add(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error)
	SendToRobot(ctx context.Context, in *RobotMessage, opts ...grpc.CallOption) (*RobotMessage, error)
	Update(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error)
	UpdateMacAddress(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error)
}

type robotsClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotsClient(cc grpc.ClientConnInterface) RobotsClient {
	return &robotsClient{cc}
}

func (c *robotsClient) GetAll(ctx context.Context, in *Robots, opts ...grpc.CallOption) (*Robots, error) {
	out := new(Robots)
	err := c.cc.Invoke(ctx, "/robots.robots/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) GetGraph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*groups.Groups, error) {
	out := new(groups.Groups)
	err := c.cc.Invoke(ctx, "/robots.robots/GetGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) Get(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error) {
	out := new(Robot)
	err := c.cc.Invoke(ctx, "/robots.robots/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) GetDirectusToken(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*DirectusToken, error) {
	out := new(DirectusToken)
	err := c.cc.Invoke(ctx, "/robots.robots/GetDirectusToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) GetMyDirectusToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DirectusToken, error) {
	out := new(DirectusToken)
	err := c.cc.Invoke(ctx, "/robots.robots/GetMyDirectusToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) GetByGroup(ctx context.Context, in *groups.Group, opts ...grpc.CallOption) (*Robots, error) {
	out := new(Robots)
	err := c.cc.Invoke(ctx, "/robots.robots/GetByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) GetSAFile(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*SaFile, error) {
	out := new(SaFile)
	err := c.cc.Invoke(ctx, "/robots.robots/GetSAFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) Add(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error) {
	out := new(Robot)
	err := c.cc.Invoke(ctx, "/robots.robots/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) SendToRobot(ctx context.Context, in *RobotMessage, opts ...grpc.CallOption) (*RobotMessage, error) {
	out := new(RobotMessage)
	err := c.cc.Invoke(ctx, "/robots.robots/SendToRobot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) Update(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error) {
	out := new(Robot)
	err := c.cc.Invoke(ctx, "/robots.robots/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotsClient) UpdateMacAddress(ctx context.Context, in *Robot, opts ...grpc.CallOption) (*Robot, error) {
	out := new(Robot)
	err := c.cc.Invoke(ctx, "/robots.robots/UpdateMacAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotsServer is the server API for Robots service.
// All implementations must embed UnimplementedRobotsServer
// for forward compatibility
type RobotsServer interface {
	GetAll(context.Context, *Robots) (*Robots, error)
	GetGraph(context.Context, *emptypb.Empty) (*groups.Groups, error)
	Get(context.Context, *Robot) (*Robot, error)
	GetDirectusToken(context.Context, *Robot) (*DirectusToken, error)
	GetMyDirectusToken(context.Context, *emptypb.Empty) (*DirectusToken, error)
	GetByGroup(context.Context, *groups.Group) (*Robots, error)
	GetSAFile(context.Context, *Robot) (*SaFile, error)
	Add(context.Context, *Robot) (*Robot, error)
	SendToRobot(context.Context, *RobotMessage) (*RobotMessage, error)
	Update(context.Context, *Robot) (*Robot, error)
	UpdateMacAddress(context.Context, *Robot) (*Robot, error)
	mustEmbedUnimplementedRobotsServer()
}

// UnimplementedRobotsServer must be embedded to have forward compatible implementations.
type UnimplementedRobotsServer struct {
}

func (UnimplementedRobotsServer) GetAll(context.Context, *Robots) (*Robots, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRobotsServer) GetGraph(context.Context, *emptypb.Empty) (*groups.Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}
func (UnimplementedRobotsServer) Get(context.Context, *Robot) (*Robot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRobotsServer) GetDirectusToken(context.Context, *Robot) (*DirectusToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirectusToken not implemented")
}
func (UnimplementedRobotsServer) GetMyDirectusToken(context.Context, *emptypb.Empty) (*DirectusToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyDirectusToken not implemented")
}
func (UnimplementedRobotsServer) GetByGroup(context.Context, *groups.Group) (*Robots, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGroup not implemented")
}
func (UnimplementedRobotsServer) GetSAFile(context.Context, *Robot) (*SaFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSAFile not implemented")
}
func (UnimplementedRobotsServer) Add(context.Context, *Robot) (*Robot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedRobotsServer) SendToRobot(context.Context, *RobotMessage) (*RobotMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToRobot not implemented")
}
func (UnimplementedRobotsServer) Update(context.Context, *Robot) (*Robot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRobotsServer) UpdateMacAddress(context.Context, *Robot) (*Robot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMacAddress not implemented")
}
func (UnimplementedRobotsServer) mustEmbedUnimplementedRobotsServer() {}

// UnsafeRobotsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotsServer will
// result in compilation errors.
type UnsafeRobotsServer interface {
	mustEmbedUnimplementedRobotsServer()
}

func RegisterRobotsServer(s grpc.ServiceRegistrar, srv RobotsServer) {
	s.RegisterService(&Robots_ServiceDesc, srv)
}

func _Robots_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robots)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetAll(ctx, req.(*Robots))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetGraph(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).Get(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_GetDirectusToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetDirectusToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetDirectusToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetDirectusToken(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_GetMyDirectusToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetMyDirectusToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetMyDirectusToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetMyDirectusToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_GetByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(groups.Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetByGroup(ctx, req.(*groups.Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_GetSAFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).GetSAFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/GetSAFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).GetSAFile(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).Add(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_SendToRobot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).SendToRobot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/SendToRobot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).SendToRobot(ctx, req.(*RobotMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).Update(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Robots_UpdateMacAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Robot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotsServer).UpdateMacAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robots.robots/UpdateMacAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotsServer).UpdateMacAddress(ctx, req.(*Robot))
	}
	return interceptor(ctx, in, info, handler)
}

// Robots_ServiceDesc is the grpc.ServiceDesc for Robots service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Robots_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "robots.robots",
	HandlerType: (*RobotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Robots_GetAll_Handler,
		},
		{
			MethodName: "GetGraph",
			Handler:    _Robots_GetGraph_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Robots_Get_Handler,
		},
		{
			MethodName: "GetDirectusToken",
			Handler:    _Robots_GetDirectusToken_Handler,
		},
		{
			MethodName: "GetMyDirectusToken",
			Handler:    _Robots_GetMyDirectusToken_Handler,
		},
		{
			MethodName: "GetByGroup",
			Handler:    _Robots_GetByGroup_Handler,
		},
		{
			MethodName: "GetSAFile",
			Handler:    _Robots_GetSAFile_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Robots_Add_Handler,
		},
		{
			MethodName: "SendToRobot",
			Handler:    _Robots_SendToRobot_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Robots_Update_Handler,
		},
		{
			MethodName: "UpdateMacAddress",
			Handler:    _Robots_UpdateMacAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "robots.proto",
}
