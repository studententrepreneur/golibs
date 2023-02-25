// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: certificate.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CertificateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Issuer string                 `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Name   string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url    string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CertificateMessage) Reset() {
	*x = CertificateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateMessage) ProtoMessage() {}

func (x *CertificateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateMessage.ProtoReflect.Descriptor instead.
func (*CertificateMessage) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateMessage) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *CertificateMessage) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *CertificateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CertificateMessage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_certificate_proto protoreflect.FileDescriptor

var file_certificate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x72, 0x65, 0x70, 0x72, 0x65, 0x6e, 0x65, 0x75, 0x72, 0x2f, 0x67, 0x6f, 0x6c, 0x69,
	0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_proto_rawDescOnce sync.Once
	file_certificate_proto_rawDescData = file_certificate_proto_rawDesc
)

func file_certificate_proto_rawDescGZIP() []byte {
	file_certificate_proto_rawDescOnce.Do(func() {
		file_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_proto_rawDescData)
	})
	return file_certificate_proto_rawDescData
}

var file_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_certificate_proto_goTypes = []interface{}{
	(*CertificateMessage)(nil),    // 0: CertificateMessage
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_certificate_proto_depIdxs = []int32{
	1, // 0: CertificateMessage.date:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_certificate_proto_init() }
func file_certificate_proto_init() {
	if File_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateMessage); i {
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
			RawDescriptor: file_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_certificate_proto_goTypes,
		DependencyIndexes: file_certificate_proto_depIdxs,
		MessageInfos:      file_certificate_proto_msgTypes,
	}.Build()
	File_certificate_proto = out.File
	file_certificate_proto_rawDesc = nil
	file_certificate_proto_goTypes = nil
	file_certificate_proto_depIdxs = nil
}