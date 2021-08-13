// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: v1/payment.proto

package v1

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

type PaymentAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentAuthReq) Reset() {
	*x = PaymentAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentAuthReq) ProtoMessage() {}

func (x *PaymentAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentAuthReq.ProtoReflect.Descriptor instead.
func (*PaymentAuthReq) Descriptor() ([]byte, []int) {
	return file_v1_payment_proto_rawDescGZIP(), []int{0}
}

type PaymentAuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentAuthReply) Reset() {
	*x = PaymentAuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentAuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentAuthReply) ProtoMessage() {}

func (x *PaymentAuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentAuthReply.ProtoReflect.Descriptor instead.
func (*PaymentAuthReply) Descriptor() ([]byte, []int) {
	return file_v1_payment_proto_rawDescGZIP(), []int{1}
}

var File_v1_payment_proto protoreflect.FileDescriptor

var file_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x5e, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x53, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_payment_proto_rawDescOnce sync.Once
	file_v1_payment_proto_rawDescData = file_v1_payment_proto_rawDesc
)

func file_v1_payment_proto_rawDescGZIP() []byte {
	file_v1_payment_proto_rawDescOnce.Do(func() {
		file_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_payment_proto_rawDescData)
	})
	return file_v1_payment_proto_rawDescData
}

var file_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_payment_proto_goTypes = []interface{}{
	(*PaymentAuthReq)(nil),   // 0: cart.service.v1.PaymentAuthReq
	(*PaymentAuthReply)(nil), // 1: cart.service.v1.PaymentAuthReply
}
var file_v1_payment_proto_depIdxs = []int32{
	0, // 0: cart.service.v1.Payment.PaymentAuth:input_type -> cart.service.v1.PaymentAuthReq
	1, // 1: cart.service.v1.Payment.PaymentAuth:output_type -> cart.service.v1.PaymentAuthReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_payment_proto_init() }
func file_v1_payment_proto_init() {
	if File_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentAuthReq); i {
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
		file_v1_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentAuthReply); i {
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
			RawDescriptor: file_v1_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_payment_proto_goTypes,
		DependencyIndexes: file_v1_payment_proto_depIdxs,
		MessageInfos:      file_v1_payment_proto_msgTypes,
	}.Build()
	File_v1_payment_proto = out.File
	file_v1_payment_proto_rawDesc = nil
	file_v1_payment_proto_goTypes = nil
	file_v1_payment_proto_depIdxs = nil
}
