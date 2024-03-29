// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: proto/ex1.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Related to students endpoint
type StudentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID int32 `protobuf:"varint,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
}

func (x *StudentID) Reset() {
	*x = StudentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentID) ProtoMessage() {}

func (x *StudentID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentID.ProtoReflect.Descriptor instead.
func (*StudentID) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{0}
}

func (x *StudentID) GetStudentID() int32 {
	if x != nil {
		return x.StudentID
	}
	return 0
}

type StudentInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID   *StudentID `protobuf:"bytes,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
	Enrollments string     `protobuf:"bytes,2,opt,name=enrollments,proto3" json:"enrollments,omitempty"`
	Workload    string     `protobuf:"bytes,3,opt,name=workload,proto3" json:"workload,omitempty"`
}

func (x *StudentInformation) Reset() {
	*x = StudentInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentInformation) ProtoMessage() {}

func (x *StudentInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentInformation.ProtoReflect.Descriptor instead.
func (*StudentInformation) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{1}
}

func (x *StudentInformation) GetStudentID() *StudentID {
	if x != nil {
		return x.StudentID
	}
	return nil
}

func (x *StudentInformation) GetEnrollments() string {
	if x != nil {
		return x.Enrollments
	}
	return ""
}

func (x *StudentInformation) GetWorkload() string {
	if x != nil {
		return x.Workload
	}
	return ""
}

type TeacherID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherID int32 `protobuf:"varint,1,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
}

func (x *TeacherID) Reset() {
	*x = TeacherID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherID) ProtoMessage() {}

func (x *TeacherID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherID.ProtoReflect.Descriptor instead.
func (*TeacherID) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{2}
}

func (x *TeacherID) GetTeacherID() int32 {
	if x != nil {
		return x.TeacherID
	}
	return 0
}

type TeacherInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeacherID              *TeacherID `protobuf:"bytes,1,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	CourseID               *CourseID  `protobuf:"bytes,2,opt,name=courseID,proto3" json:"courseID,omitempty"`
	StudentPopularityScore int32      `protobuf:"varint,3,opt,name=studentPopularityScore,proto3" json:"studentPopularityScore,omitempty"`
}

func (x *TeacherInformation) Reset() {
	*x = TeacherInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherInformation) ProtoMessage() {}

func (x *TeacherInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherInformation.ProtoReflect.Descriptor instead.
func (*TeacherInformation) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{3}
}

func (x *TeacherInformation) GetTeacherID() *TeacherID {
	if x != nil {
		return x.TeacherID
	}
	return nil
}

func (x *TeacherInformation) GetCourseID() *CourseID {
	if x != nil {
		return x.CourseID
	}
	return nil
}

func (x *TeacherInformation) GetStudentPopularityScore() int32 {
	if x != nil {
		return x.StudentPopularityScore
	}
	return 0
}

// Related to courses endpoint
type CourseID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseID int32 `protobuf:"varint,1,opt,name=courseID,proto3" json:"courseID,omitempty"`
}

func (x *CourseID) Reset() {
	*x = CourseID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseID) ProtoMessage() {}

func (x *CourseID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseID.ProtoReflect.Descriptor instead.
func (*CourseID) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{4}
}

func (x *CourseID) GetCourseID() int32 {
	if x != nil {
		return x.CourseID
	}
	return 0
}

type CourseInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseID   *CourseID  `protobuf:"bytes,1,opt,name=courseID,proto3" json:"courseID,omitempty"`
	TeacherID  *TeacherID `protobuf:"bytes,2,opt,name=teacherID,proto3" json:"teacherID,omitempty"`
	CourseName string     `protobuf:"bytes,3,opt,name=courseName,proto3" json:"courseName,omitempty"`
}

func (x *CourseInformation) Reset() {
	*x = CourseInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ex1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseInformation) ProtoMessage() {}

func (x *CourseInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ex1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseInformation.ProtoReflect.Descriptor instead.
func (*CourseInformation) Descriptor() ([]byte, []int) {
	return file_proto_ex1_proto_rawDescGZIP(), []int{5}
}

func (x *CourseInformation) GetCourseID() *CourseID {
	if x != nil {
		return x.CourseID
	}
	return nil
}

func (x *CourseInformation) GetTeacherID() *TeacherID {
	if x != nil {
		return x.TeacherID
	}
	return nil
}

func (x *CourseInformation) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

var File_proto_ex1_proto protoreflect.FileDescriptor

var file_proto_ex1_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x29, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x44, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x26, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xb6, 0x01, 0x0a, 0x09, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x31, 0x12, 0x38, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x44, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4f, 0x4c, 0x61, 0x75, 0x72, 0x73, 0x65, 0x6e, 0x2f, 0x47, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ex1_proto_rawDescOnce sync.Once
	file_proto_ex1_proto_rawDescData = file_proto_ex1_proto_rawDesc
)

func file_proto_ex1_proto_rawDescGZIP() []byte {
	file_proto_ex1_proto_rawDescOnce.Do(func() {
		file_proto_ex1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ex1_proto_rawDescData)
	})
	return file_proto_ex1_proto_rawDescData
}

var file_proto_ex1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_ex1_proto_goTypes = []interface{}{
	(*StudentID)(nil),          // 0: proto.studentID
	(*StudentInformation)(nil), // 1: proto.studentInformation
	(*TeacherID)(nil),          // 2: proto.teacherID
	(*TeacherInformation)(nil), // 3: proto.teacherInformation
	(*CourseID)(nil),           // 4: proto.courseID
	(*CourseInformation)(nil),  // 5: proto.courseInformation
}
var file_proto_ex1_proto_depIdxs = []int32{
	0, // 0: proto.studentInformation.studentID:type_name -> proto.studentID
	2, // 1: proto.teacherInformation.teacherID:type_name -> proto.teacherID
	4, // 2: proto.teacherInformation.courseID:type_name -> proto.courseID
	4, // 3: proto.courseInformation.courseID:type_name -> proto.courseID
	2, // 4: proto.courseInformation.teacherID:type_name -> proto.teacherID
	0, // 5: proto.exercise1.Student:input_type -> proto.studentID
	2, // 6: proto.exercise1.Teacher:input_type -> proto.teacherID
	4, // 7: proto.exercise1.Course:input_type -> proto.courseID
	1, // 8: proto.exercise1.Student:output_type -> proto.studentInformation
	3, // 9: proto.exercise1.Teacher:output_type -> proto.teacherInformation
	5, // 10: proto.exercise1.Course:output_type -> proto.courseInformation
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_ex1_proto_init() }
func file_proto_ex1_proto_init() {
	if File_proto_ex1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ex1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ex1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ex1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ex1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ex1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ex1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ex1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ex1_proto_goTypes,
		DependencyIndexes: file_proto_ex1_proto_depIdxs,
		MessageInfos:      file_proto_ex1_proto_msgTypes,
	}.Build()
	File_proto_ex1_proto = out.File
	file_proto_ex1_proto_rawDesc = nil
	file_proto_ex1_proto_goTypes = nil
	file_proto_ex1_proto_depIdxs = nil
}
