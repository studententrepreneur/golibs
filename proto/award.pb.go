// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: award.proto

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

type AwardMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Awarder string `protobuf:"bytes,1,opt,name=awarder,proto3" json:"awarder,omitempty"`
	Date    string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Summary string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AwardMessage) Reset() {
	*x = AwardMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_award_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwardMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwardMessage) ProtoMessage() {}

func (x *AwardMessage) ProtoReflect() protoreflect.Message {
	mi := &file_award_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwardMessage.ProtoReflect.Descriptor instead.
func (*AwardMessage) Descriptor() ([]byte, []int) {
	return file_award_proto_rawDescGZIP(), []int{0}
}

func (x *AwardMessage) GetAwarder() string {
	if x != nil {
		return x.Awarder
	}
	return ""
}

func (x *AwardMessage) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AwardMessage) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *AwardMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_award_proto protoreflect.FileDescriptor

var file_award_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x0c, 0x41, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x70, 0x72, 0x65, 0x6e, 0x65, 0x75, 0x72, 0x2f, 0x67, 0x6f,
	0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_award_proto_rawDescOnce sync.Once
	file_award_proto_rawDescData = file_award_proto_rawDesc
)

func file_award_proto_rawDescGZIP() []byte {
	file_award_proto_rawDescOnce.Do(func() {
		file_award_proto_rawDescData = protoimpl.X.CompressGZIP(file_award_proto_rawDescData)
	})
	return file_award_proto_rawDescData
}

var file_award_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_award_proto_goTypes = []interface{}{
	(*AwardMessage)(nil), // 0: AwardMessage
}
var file_award_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_award_proto_init() }
func file_award_proto_init() {
	if File_award_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_award_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwardMessage); i {
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
			RawDescriptor: file_award_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_award_proto_goTypes,
		DependencyIndexes: file_award_proto_depIdxs,
		MessageInfos:      file_award_proto_msgTypes,
	}.Build()
	File_award_proto = out.File
	file_award_proto_rawDesc = nil
	file_award_proto_goTypes = nil
	file_award_proto_depIdxs = nil
}