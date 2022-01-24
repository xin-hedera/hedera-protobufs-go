// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: consensus_create_topic.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// See [ConsensusService.createTopic()](#proto.ConsensusService)
type ConsensusCreateTopicTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Short publicly visible memo about the topic. No guarantee of uniqueness.
	Memo string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// Access control for updateTopic/deleteTopic.
	// Anyone can increase the topic's expirationTime via ConsensusService.updateTopic(), regardless of the adminKey.
	// If no adminKey is specified, updateTopic may only be used to extend the topic's expirationTime, and deleteTopic
	// is disallowed.
	AdminKey *Key `protobuf:"bytes,2,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// Access control for submitMessage.
	// If unspecified, no access control is performed on ConsensusService.submitMessage (all submissions are allowed).
	SubmitKey *Key `protobuf:"bytes,3,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	//*
	// The initial lifetime of the topic and the amount of time to attempt to extend the topic's lifetime by
	// automatically at the topic's expirationTime, if the autoRenewAccount is configured (once autoRenew functionality
	// is supported by HAPI).
	// Limited to MIN_AUTORENEW_PERIOD and MAX_AUTORENEW_PERIOD value by server-side configuration.
	// Required.
	AutoRenewPeriod *Duration `protobuf:"bytes,6,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// Optional account to be used at the topic's expirationTime to extend the life of the topic (once autoRenew
	// functionality is supported by HAPI).
	// The topic lifetime will be extended up to a maximum of the autoRenewPeriod or however long the topic
	// can be extended using all funds on the account (whichever is the smaller duration/amount and if any extension
	// is possible with the account's funds).
	// If specified, there must be an adminKey and the autoRenewAccount must sign this transaction.
	AutoRenewAccount *AccountID `protobuf:"bytes,7,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
}

func (x *ConsensusCreateTopicTransactionBody) Reset() {
	*x = ConsensusCreateTopicTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_create_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusCreateTopicTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusCreateTopicTransactionBody) ProtoMessage() {}

func (x *ConsensusCreateTopicTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_create_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusCreateTopicTransactionBody.ProtoReflect.Descriptor instead.
func (*ConsensusCreateTopicTransactionBody) Descriptor() ([]byte, []int) {
	return file_consensus_create_topic_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusCreateTopicTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ConsensusCreateTopicTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ConsensusCreateTopicTransactionBody) GetSubmitKey() *Key {
	if x != nil {
		return x.SubmitKey
	}
	return nil
}

func (x *ConsensusCreateTopicTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *ConsensusCreateTopicTransactionBody) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

var File_consensus_create_topic_proto protoreflect.FileDescriptor

var file_consensus_create_topic_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x23, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x10, 0x61,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_create_topic_proto_rawDescOnce sync.Once
	file_consensus_create_topic_proto_rawDescData = file_consensus_create_topic_proto_rawDesc
)

func file_consensus_create_topic_proto_rawDescGZIP() []byte {
	file_consensus_create_topic_proto_rawDescOnce.Do(func() {
		file_consensus_create_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_create_topic_proto_rawDescData)
	})
	return file_consensus_create_topic_proto_rawDescData
}

var file_consensus_create_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_consensus_create_topic_proto_goTypes = []interface{}{
	(*ConsensusCreateTopicTransactionBody)(nil), // 0: proto.ConsensusCreateTopicTransactionBody
	(*Key)(nil),       // 1: proto.Key
	(*Duration)(nil),  // 2: proto.Duration
	(*AccountID)(nil), // 3: proto.AccountID
}
var file_consensus_create_topic_proto_depIdxs = []int32{
	1, // 0: proto.ConsensusCreateTopicTransactionBody.adminKey:type_name -> proto.Key
	1, // 1: proto.ConsensusCreateTopicTransactionBody.submitKey:type_name -> proto.Key
	2, // 2: proto.ConsensusCreateTopicTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	3, // 3: proto.ConsensusCreateTopicTransactionBody.autoRenewAccount:type_name -> proto.AccountID
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_consensus_create_topic_proto_init() }
func file_consensus_create_topic_proto_init() {
	if File_consensus_create_topic_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consensus_create_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusCreateTopicTransactionBody); i {
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
			RawDescriptor: file_consensus_create_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_create_topic_proto_goTypes,
		DependencyIndexes: file_consensus_create_topic_proto_depIdxs,
		MessageInfos:      file_consensus_create_topic_proto_msgTypes,
	}.Build()
	File_consensus_create_topic_proto = out.File
	file_consensus_create_topic_proto_rawDesc = nil
	file_consensus_create_topic_proto_goTypes = nil
	file_consensus_create_topic_proto_depIdxs = nil
}
