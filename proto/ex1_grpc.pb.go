// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/ex1.proto

package proto

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
	Exercise1_Student_FullMethodName = "/proto.exercise1/Student"
	Exercise1_Teacher_FullMethodName = "/proto.exercise1/Teacher"
	Exercise1_Course_FullMethodName  = "/proto.exercise1/Course"
)

// Exercise1Client is the client API for Exercise1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Exercise1Client interface {
	// Student service endpoints:
	Student(ctx context.Context, in *StudentID, opts ...grpc.CallOption) (*StudentInformation, error)
	// Teacher service endpoints:
	Teacher(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*TeacherInformation, error)
	// Course service endpoints:
	Course(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*CourseInformation, error)
}

type exercise1Client struct {
	cc grpc.ClientConnInterface
}

func NewExercise1Client(cc grpc.ClientConnInterface) Exercise1Client {
	return &exercise1Client{cc}
}

func (c *exercise1Client) Student(ctx context.Context, in *StudentID, opts ...grpc.CallOption) (*StudentInformation, error) {
	out := new(StudentInformation)
	err := c.cc.Invoke(ctx, Exercise1_Student_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exercise1Client) Teacher(ctx context.Context, in *TeacherID, opts ...grpc.CallOption) (*TeacherInformation, error) {
	out := new(TeacherInformation)
	err := c.cc.Invoke(ctx, Exercise1_Teacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exercise1Client) Course(ctx context.Context, in *CourseID, opts ...grpc.CallOption) (*CourseInformation, error) {
	out := new(CourseInformation)
	err := c.cc.Invoke(ctx, Exercise1_Course_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Exercise1Server is the server API for Exercise1 service.
// All implementations must embed UnimplementedExercise1Server
// for forward compatibility
type Exercise1Server interface {
	// Student service endpoints:
	Student(context.Context, *StudentID) (*StudentInformation, error)
	// Teacher service endpoints:
	Teacher(context.Context, *TeacherID) (*TeacherInformation, error)
	// Course service endpoints:
	Course(context.Context, *CourseID) (*CourseInformation, error)
	mustEmbedUnimplementedExercise1Server()
}

// UnimplementedExercise1Server must be embedded to have forward compatible implementations.
type UnimplementedExercise1Server struct {
}

func (UnimplementedExercise1Server) Student(context.Context, *StudentID) (*StudentInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Student not implemented")
}
func (UnimplementedExercise1Server) Teacher(context.Context, *TeacherID) (*TeacherInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teacher not implemented")
}
func (UnimplementedExercise1Server) Course(context.Context, *CourseID) (*CourseInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Course not implemented")
}
func (UnimplementedExercise1Server) mustEmbedUnimplementedExercise1Server() {}

// UnsafeExercise1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Exercise1Server will
// result in compilation errors.
type UnsafeExercise1Server interface {
	mustEmbedUnimplementedExercise1Server()
}

func RegisterExercise1Server(s grpc.ServiceRegistrar, srv Exercise1Server) {
	s.RegisterService(&Exercise1_ServiceDesc, srv)
}

func _Exercise1_Student_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Exercise1Server).Student(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exercise1_Student_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Exercise1Server).Student(ctx, req.(*StudentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exercise1_Teacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Exercise1Server).Teacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exercise1_Teacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Exercise1Server).Teacher(ctx, req.(*TeacherID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exercise1_Course_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Exercise1Server).Course(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Exercise1_Course_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Exercise1Server).Course(ctx, req.(*CourseID))
	}
	return interceptor(ctx, in, info, handler)
}

// Exercise1_ServiceDesc is the grpc.ServiceDesc for Exercise1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exercise1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.exercise1",
	HandlerType: (*Exercise1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Student",
			Handler:    _Exercise1_Student_Handler,
		},
		{
			MethodName: "Teacher",
			Handler:    _Exercise1_Teacher_Handler,
		},
		{
			MethodName: "Course",
			Handler:    _Exercise1_Course_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ex1.proto",
}