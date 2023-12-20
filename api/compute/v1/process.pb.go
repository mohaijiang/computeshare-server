// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: api/compute/v1/process.proto

package v1

import (
	v1 "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
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

type CreateInstanceProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance       *CreateInstanceRequest            `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	NetworkMapping []*v1.CreateNetworkMappingRequest `protobuf:"bytes,2,rep,name=networkMapping,proto3" json:"networkMapping,omitempty"`
}

func (x *CreateInstanceProcessRequest) Reset() {
	*x = CreateInstanceProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstanceProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceProcessRequest) ProtoMessage() {}

func (x *CreateInstanceProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceProcessRequest.ProtoReflect.Descriptor instead.
func (*CreateInstanceProcessRequest) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_process_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInstanceProcessRequest) GetInstance() *CreateInstanceRequest {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *CreateInstanceProcessRequest) GetNetworkMapping() []*v1.CreateNetworkMappingRequest {
	if x != nil {
		return x.NetworkMapping
	}
	return nil
}

type CreateInstanceProcessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateInstanceProcessReply) Reset() {
	*x = CreateInstanceProcessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_compute_v1_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstanceProcessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceProcessReply) ProtoMessage() {}

func (x *CreateInstanceProcessReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_compute_v1_process_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceProcessReply.ProtoReflect.Descriptor instead.
func (*CreateInstanceProcessReply) Descriptor() ([]byte, []int) {
	return file_api_compute_v1_process_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInstanceProcessReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateInstanceProcessReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_compute_v1_process_proto protoreflect.FileDescriptor

var file_api_compute_v1_process_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x4a, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e,
	0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x4f, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x68, 0x61, 0x69, 0x6a, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_compute_v1_process_proto_rawDescOnce sync.Once
	file_api_compute_v1_process_proto_rawDescData = file_api_compute_v1_process_proto_rawDesc
)

func file_api_compute_v1_process_proto_rawDescGZIP() []byte {
	file_api_compute_v1_process_proto_rawDescOnce.Do(func() {
		file_api_compute_v1_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_compute_v1_process_proto_rawDescData)
	})
	return file_api_compute_v1_process_proto_rawDescData
}

var file_api_compute_v1_process_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_compute_v1_process_proto_goTypes = []interface{}{
	(*CreateInstanceProcessRequest)(nil),   // 0: api.compute.v1.CreateInstanceProcessRequest
	(*CreateInstanceProcessReply)(nil),     // 1: api.compute.v1.CreateInstanceProcessReply
	(*CreateInstanceRequest)(nil),          // 2: api.compute.v1.CreateInstanceRequest
	(*v1.CreateNetworkMappingRequest)(nil), // 3: api.network_mapping.v1.CreateNetworkMappingRequest
}
var file_api_compute_v1_process_proto_depIdxs = []int32{
	2, // 0: api.compute.v1.CreateInstanceProcessRequest.instance:type_name -> api.compute.v1.CreateInstanceRequest
	3, // 1: api.compute.v1.CreateInstanceProcessRequest.networkMapping:type_name -> api.network_mapping.v1.CreateNetworkMappingRequest
	0, // 2: api.compute.v1.Process.CreateInstanceProcess:input_type -> api.compute.v1.CreateInstanceProcessRequest
	1, // 3: api.compute.v1.Process.CreateInstanceProcess:output_type -> api.compute.v1.CreateInstanceProcessReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_compute_v1_process_proto_init() }
func file_api_compute_v1_process_proto_init() {
	if File_api_compute_v1_process_proto != nil {
		return
	}
	file_api_compute_v1_compute_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_compute_v1_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstanceProcessRequest); i {
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
		file_api_compute_v1_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstanceProcessReply); i {
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
			RawDescriptor: file_api_compute_v1_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_compute_v1_process_proto_goTypes,
		DependencyIndexes: file_api_compute_v1_process_proto_depIdxs,
		MessageInfos:      file_api_compute_v1_process_proto_msgTypes,
	}.Build()
	File_api_compute_v1_process_proto = out.File
	file_api_compute_v1_process_proto_rawDesc = nil
	file_api_compute_v1_process_proto_goTypes = nil
	file_api_compute_v1_process_proto_depIdxs = nil
}
