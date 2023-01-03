// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: trainer.proto

package trainer

import (
	reflect "reflect"
	sync "sync"

	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IsHourAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *IsHourAvailableRequest) Reset() {
	*x = IsHourAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableRequest) ProtoMessage() {}

func (x *IsHourAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableRequest.ProtoReflect.Descriptor instead.
func (*IsHourAvailableRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *IsHourAvailableRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type IsHourAvailableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvailable bool `protobuf:"varint,1,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
}

func (x *IsHourAvailableResponse) Reset() {
	*x = IsHourAvailableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableResponse) ProtoMessage() {}

func (x *IsHourAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableResponse.ProtoReflect.Descriptor instead.
func (*IsHourAvailableResponse) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *IsHourAvailableResponse) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type UpdateHourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *UpdateHourRequest) Reset() {
	*x = UpdateHourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHourRequest) ProtoMessage() {}

func (x *UpdateHourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHourRequest.ProtoReflect.Descriptor instead.
func (*UpdateHourRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHourRequest) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_trainer_proto protoreflect.FileDescriptor

var file_trainer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x3c, 0x0a, 0x17, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x43,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x32, 0xc5, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x72, 0x65, 0x65, 0x44,
	0x6f, 0x74, 0x73, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x77, 0x69, 0x6c, 0x64, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x6f, 0x75, 0x74, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x64, 0x64, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trainer_proto_rawDescOnce sync.Once
	file_trainer_proto_rawDescData = file_trainer_proto_rawDesc
)

func file_trainer_proto_rawDescGZIP() []byte {
	file_trainer_proto_rawDescOnce.Do(func() {
		file_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_trainer_proto_rawDescData)
	})
	return file_trainer_proto_rawDescData
}

var file_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trainer_proto_goTypes = []interface{}{
	(*IsHourAvailableRequest)(nil),  // 0: trainer.IsHourAvailableRequest
	(*IsHourAvailableResponse)(nil), // 1: trainer.IsHourAvailableResponse
	(*UpdateHourRequest)(nil),       // 2: trainer.UpdateHourRequest
	(*timestamp.Timestamp)(nil),     // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_trainer_proto_depIdxs = []int32{
	3, // 0: trainer.IsHourAvailableRequest.time:type_name -> google.protobuf.Timestamp
	3, // 1: trainer.UpdateHourRequest.time:type_name -> google.protobuf.Timestamp
	0, // 2: trainer.TrainerService.IsHourAvailable:input_type -> trainer.IsHourAvailableRequest
	2, // 3: trainer.TrainerService.ScheduleTraining:input_type -> trainer.UpdateHourRequest
	2, // 4: trainer.TrainerService.CancelTraining:input_type -> trainer.UpdateHourRequest
	2, // 5: trainer.TrainerService.MakeHourAvailable:input_type -> trainer.UpdateHourRequest
	1, // 6: trainer.TrainerService.IsHourAvailable:output_type -> trainer.IsHourAvailableResponse
	4, // 7: trainer.TrainerService.ScheduleTraining:output_type -> google.protobuf.Empty
	4, // 8: trainer.TrainerService.CancelTraining:output_type -> google.protobuf.Empty
	4, // 9: trainer.TrainerService.MakeHourAvailable:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trainer_proto_init() }
func file_trainer_proto_init() {
	if File_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trainer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableRequest); i {
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
		file_trainer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableResponse); i {
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
		file_trainer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHourRequest); i {
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
			RawDescriptor: file_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trainer_proto_goTypes,
		DependencyIndexes: file_trainer_proto_depIdxs,
		MessageInfos:      file_trainer_proto_msgTypes,
	}.Build()
	File_trainer_proto = out.File
	file_trainer_proto_rawDesc = nil
	file_trainer_proto_goTypes = nil
	file_trainer_proto_depIdxs = nil
}
