// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: math/v1/math.proto

package mathv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_math_v1_math_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *AddRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_math_v1_math_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type MultiplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *MultiplyRequest) Reset() {
	*x = MultiplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1_math_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyRequest) ProtoMessage() {}

func (x *MultiplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1_math_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyRequest.ProtoReflect.Descriptor instead.
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return file_math_v1_math_proto_rawDescGZIP(), []int{2}
}

func (x *MultiplyRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *MultiplyRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type MultiplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *MultiplyResponse) Reset() {
	*x = MultiplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_v1_math_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyResponse) ProtoMessage() {}

func (x *MultiplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_v1_math_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyResponse.ProtoReflect.Descriptor instead.
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return file_math_v1_math_proto_rawDescGZIP(), []int{3}
}

func (x *MultiplyResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_math_v1_math_proto protoreflect.FileDescriptor

var file_math_v1_math_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x22, 0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x22, 0x39, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x24, 0x0a,
	0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x32, 0xad, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a,
	0x22, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x58, 0x0a, 0x08, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x79, 0x42, 0x8a, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x68, 0x6f,
	0x6e, 0x67, 0x70, 0x69, 0x73, 0x75, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x61, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4d, 0x61,
	0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4d, 0x61, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x13, 0x4d, 0x61, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4d, 0x61, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_math_v1_math_proto_rawDescOnce sync.Once
	file_math_v1_math_proto_rawDescData = file_math_v1_math_proto_rawDesc
)

func file_math_v1_math_proto_rawDescGZIP() []byte {
	file_math_v1_math_proto_rawDescOnce.Do(func() {
		file_math_v1_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_math_v1_math_proto_rawDescData)
	})
	return file_math_v1_math_proto_rawDescData
}

var file_math_v1_math_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_math_v1_math_proto_goTypes = []any{
	(*AddRequest)(nil),       // 0: math.v1.AddRequest
	(*AddResponse)(nil),      // 1: math.v1.AddResponse
	(*MultiplyRequest)(nil),  // 2: math.v1.MultiplyRequest
	(*MultiplyResponse)(nil), // 3: math.v1.MultiplyResponse
}
var file_math_v1_math_proto_depIdxs = []int32{
	0, // 0: math.v1.MathService.Add:input_type -> math.v1.AddRequest
	2, // 1: math.v1.MathService.Multiply:input_type -> math.v1.MultiplyRequest
	1, // 2: math.v1.MathService.Add:output_type -> math.v1.AddResponse
	3, // 3: math.v1.MathService.Multiply:output_type -> math.v1.MultiplyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_math_v1_math_proto_init() }
func file_math_v1_math_proto_init() {
	if File_math_v1_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_math_v1_math_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddRequest); i {
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
		file_math_v1_math_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddResponse); i {
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
		file_math_v1_math_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MultiplyRequest); i {
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
		file_math_v1_math_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MultiplyResponse); i {
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
			RawDescriptor: file_math_v1_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_math_v1_math_proto_goTypes,
		DependencyIndexes: file_math_v1_math_proto_depIdxs,
		MessageInfos:      file_math_v1_math_proto_msgTypes,
	}.Build()
	File_math_v1_math_proto = out.File
	file_math_v1_math_proto_rawDesc = nil
	file_math_v1_math_proto_goTypes = nil
	file_math_v1_math_proto_depIdxs = nil
}
