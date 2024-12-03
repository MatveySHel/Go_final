// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protofile.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: protofile.proto

package pb

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

type AskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotel string `protobuf:"bytes,1,opt,name=hotel,proto3" json:"hotel,omitempty"`
}

func (x *AskRequest) Reset() {
	*x = AskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskRequest) ProtoMessage() {}

func (x *AskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskRequest.ProtoReflect.Descriptor instead.
func (*AskRequest) Descriptor() ([]byte, []int) {
	return file_protofile_proto_rawDescGZIP(), []int{0}
}

func (x *AskRequest) GetHotel() string {
	if x != nil {
		return x.Hotel
	}
	return ""
}

type AskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price int32 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AskResponse) Reset() {
	*x = AskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskResponse) ProtoMessage() {}

func (x *AskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskResponse.ProtoReflect.Descriptor instead.
func (*AskResponse) Descriptor() ([]byte, []int) {
	return file_protofile_proto_rawDescGZIP(), []int{1}
}

func (x *AskResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_protofile_proto protoreflect.FileDescriptor

var file_protofile_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x70, 0x62, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0x23,
	0x0a, 0x0b, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x32, 0x48, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x70,
	0x62, 0x2e, 0x41, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x65,
	0x6d, 0x53, 0x68, 0x61, 0x6d, 0x72, 0x6f, 0x2f, 0x47, 0x6f, 0x5f, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofile_proto_rawDescOnce sync.Once
	file_protofile_proto_rawDescData = file_protofile_proto_rawDesc
)

func file_protofile_proto_rawDescGZIP() []byte {
	file_protofile_proto_rawDescOnce.Do(func() {
		file_protofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofile_proto_rawDescData)
	})
	return file_protofile_proto_rawDescData
}

var file_protofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protofile_proto_goTypes = []any{
	(*AskRequest)(nil),  // 0: pricepb.AskRequest
	(*AskResponse)(nil), // 1: pricepb.AskResponse
}
var file_protofile_proto_depIdxs = []int32{
	0, // 0: pricepb.MessengerServer.askPrice:input_type -> pricepb.AskRequest
	1, // 1: pricepb.MessengerServer.askPrice:output_type -> pricepb.AskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofile_proto_init() }
func file_protofile_proto_init() {
	if File_protofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofile_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AskRequest); i {
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
		file_protofile_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AskResponse); i {
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
			RawDescriptor: file_protofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofile_proto_goTypes,
		DependencyIndexes: file_protofile_proto_depIdxs,
		MessageInfos:      file_protofile_proto_msgTypes,
	}.Build()
	File_protofile_proto = out.File
	file_protofile_proto_rawDesc = nil
	file_protofile_proto_goTypes = nil
	file_protofile_proto_depIdxs = nil
}