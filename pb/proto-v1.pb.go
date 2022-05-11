// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto-v1.proto

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

type Stuff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Stuff) Reset() {
	*x = Stuff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stuff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stuff) ProtoMessage() {}

func (x *Stuff) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stuff.ProtoReflect.Descriptor instead.
func (*Stuff) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Stuff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stuff *Stuff `protobuf:"bytes,2,opt,name=stuff,proto3" json:"stuff,omitempty"`
}

func (x *GatewayRequest) Reset() {
	*x = GatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRequest) ProtoMessage() {}

func (x *GatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRequest.ProtoReflect.Descriptor instead.
func (*GatewayRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GatewayRequest) GetStuff() *Stuff {
	if x != nil {
		return x.Stuff
	}
	return nil
}

type GatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GatewayResponse) Reset() {
	*x = GatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayResponse) ProtoMessage() {}

func (x *GatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayResponse.ProtoReflect.Descriptor instead.
func (*GatewayResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stuff *Stuff `protobuf:"bytes,2,opt,name=stuff,proto3" json:"stuff,omitempty"`
}

func (x *ApplicationRequest) Reset() {
	*x = ApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationRequest) ProtoMessage() {}

func (x *ApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationRequest.ProtoReflect.Descriptor instead.
func (*ApplicationRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApplicationRequest) GetStuff() *Stuff {
	if x != nil {
		return x.Stuff
	}
	return nil
}

var File_proto_v1_proto protoreflect.FileDescriptor

var file_proto_v1_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1b, 0x0a, 0x05, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a,
	0x0e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x05, 0x73, 0x74, 0x75, 0x66, 0x66, 0x22, 0x23, 0x0a,
	0x0f, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x42, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x66,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x66, 0x66, 0x32, 0x3c, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x31, 0x0a, 0x0a, 0x44, 0x6f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x0f, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x42, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x13, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6d, 0x6f, 0x66, 0x66, 0x61, 0x74, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_proto_rawDescOnce sync.Once
	file_proto_v1_proto_rawDescData = file_proto_v1_proto_rawDesc
)

func file_proto_v1_proto_rawDescGZIP() []byte {
	file_proto_v1_proto_rawDescOnce.Do(func() {
		file_proto_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_proto_rawDescData)
	})
	return file_proto_v1_proto_rawDescData
}

var file_proto_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_proto_goTypes = []interface{}{
	(*Stuff)(nil),              // 0: Stuff
	(*GatewayRequest)(nil),     // 1: GatewayRequest
	(*GatewayResponse)(nil),    // 2: GatewayResponse
	(*ApplicationRequest)(nil), // 3: ApplicationRequest
}
var file_proto_v1_proto_depIdxs = []int32{
	0, // 0: GatewayRequest.stuff:type_name -> Stuff
	0, // 1: ApplicationRequest.stuff:type_name -> Stuff
	1, // 2: Gateway.DoTheThing:input_type -> GatewayRequest
	3, // 3: Application.TheThing:input_type -> ApplicationRequest
	2, // 4: Gateway.DoTheThing:output_type -> GatewayResponse
	2, // 5: Application.TheThing:output_type -> GatewayResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_v1_proto_init() }
func file_proto_v1_proto_init() {
	if File_proto_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stuff); i {
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
		file_proto_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayRequest); i {
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
		file_proto_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayResponse); i {
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
		file_proto_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationRequest); i {
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
			RawDescriptor: file_proto_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_v1_proto_goTypes,
		DependencyIndexes: file_proto_v1_proto_depIdxs,
		MessageInfos:      file_proto_v1_proto_msgTypes,
	}.Build()
	File_proto_v1_proto = out.File
	file_proto_v1_proto_rawDesc = nil
	file_proto_v1_proto_goTypes = nil
	file_proto_v1_proto_depIdxs = nil
}