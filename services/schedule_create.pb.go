// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: schedule_create.proto

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
// Create a new <i>schedule entity</i> (or simply, <i>schedule</i>) in the network's action queue.
// Upon <tt>SUCCESS</tt>, the receipt contains the `ScheduleID` of the created schedule. A schedule
// entity includes a <tt>scheduledTransactionBody</tt> to be executed when the schedule has
// collected enough signing Ed25519 keys to satisfy the scheduled transaction's signing
// requirements. Upon `SUCCESS`, the receipt also includes the <tt>scheduledTransactionID</tt> to
// use to query for the record of the scheduled transaction's execution (if it occurs).
//
// The expiration time of a schedule is always 30 minutes; it remains in state and can be queried
// using <tt>GetScheduleInfo</tt> until expiration, no matter if the scheduled transaction has
// executed or marked deleted.
//
// If the <tt>adminKey</tt> field is omitted, the resulting schedule is immutable. If the
// <tt>adminKey</tt> is set, the <tt>ScheduleDelete</tt> transaction can be used to mark it as
// deleted. The creator may also specify an optional <tt>memo</tt> whose UTF-8 encoding is at most
// 100 bytes and does not include the zero byte is also supported.
//
// When a scheduled transaction whose schedule has collected enough signing keys is executed, the
// network only charges its payer the service fee, and not the node and network fees. If the
// optional <tt>payerAccountID</tt> is set, the network charges this account. Otherwise it charges
// the payer of the originating <tt>ScheduleCreate</tt>.
//
// Two <tt>ScheduleCreate</tt> transactions are <i>identical</i> if they are equal in all their
// fields other than <tt>payerAccountID</tt>.  (Here "equal" should be understood in the sense of
// gRPC object equality in the network software runtime. In particular, a gRPC object with <a
// href="https://developers.google.com/protocol-buffers/docs/proto3#unknowns">unknown fields</a> is
// not equal to a gRPC object without unknown fields, even if they agree on all known fields.)
//
// A <tt>ScheduleCreate</tt> transaction that attempts to re-create an identical schedule already in
// state will receive a receipt with status <tt>IDENTICAL_SCHEDULE_ALREADY_CREATED</tt>; the receipt
// will include the <tt>ScheduleID</tt> of the extant schedule, which may be used in a subsequent
// <tt>ScheduleSign</tt> transaction. (The receipt will also include the <tt>TransactionID</tt> to
// use in querying for the receipt or record of the scheduled transaction.)
//
// Other notable response codes include, <tt>INVALID_ACCOUNT_ID</tt>,
// <tt>UNSCHEDULABLE_TRANSACTION</tt>, <tt>UNRESOLVABLE_REQUIRED_SIGNERS</tt>,
// <tt>INVALID_SIGNATURE</tt>. For more information please see the section of this documentation on
// the <tt>ResponseCode</tt> enum.
type ScheduleCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The scheduled transaction
	ScheduledTransactionBody *SchedulableTransactionBody `protobuf:"bytes,1,opt,name=scheduledTransactionBody,proto3" json:"scheduledTransactionBody,omitempty"`
	//*
	// An optional memo with a UTF-8 encoding of no more than 100 bytes which does not contain the
	// zero byte
	Memo string `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// An optional Hedera key which can be used to sign a ScheduleDelete and remove the schedule
	AdminKey *Key `protobuf:"bytes,3,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// An optional id of the account to be charged the service fee for the scheduled transaction at
	// the consensus time that it executes (if ever); defaults to the ScheduleCreate payer if not
	// given
	PayerAccountID *AccountID `protobuf:"bytes,4,opt,name=payerAccountID,proto3" json:"payerAccountID,omitempty"`
}

func (x *ScheduleCreateTransactionBody) Reset() {
	*x = ScheduleCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleCreateTransactionBody) ProtoMessage() {}

func (x *ScheduleCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*ScheduleCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_schedule_create_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleCreateTransactionBody) GetScheduledTransactionBody() *SchedulableTransactionBody {
	if x != nil {
		return x.ScheduledTransactionBody
	}
	return nil
}

func (x *ScheduleCreateTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ScheduleCreateTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ScheduleCreateTransactionBody) GetPayerAccountID() *AccountID {
	if x != nil {
		return x.PayerAccountID
	}
	return nil
}

var File_schedule_create_proto protoreflect.FileDescriptor

var file_schedule_create_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x1d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x5d, 0x0a, 0x18, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x18, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x26, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61,
	0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_create_proto_rawDescOnce sync.Once
	file_schedule_create_proto_rawDescData = file_schedule_create_proto_rawDesc
)

func file_schedule_create_proto_rawDescGZIP() []byte {
	file_schedule_create_proto_rawDescOnce.Do(func() {
		file_schedule_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_create_proto_rawDescData)
	})
	return file_schedule_create_proto_rawDescData
}

var file_schedule_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schedule_create_proto_goTypes = []interface{}{
	(*ScheduleCreateTransactionBody)(nil), // 0: proto.ScheduleCreateTransactionBody
	(*SchedulableTransactionBody)(nil),    // 1: proto.SchedulableTransactionBody
	(*Key)(nil),                           // 2: proto.Key
	(*AccountID)(nil),                     // 3: proto.AccountID
}
var file_schedule_create_proto_depIdxs = []int32{
	1, // 0: proto.ScheduleCreateTransactionBody.scheduledTransactionBody:type_name -> proto.SchedulableTransactionBody
	2, // 1: proto.ScheduleCreateTransactionBody.adminKey:type_name -> proto.Key
	3, // 2: proto.ScheduleCreateTransactionBody.payerAccountID:type_name -> proto.AccountID
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schedule_create_proto_init() }
func file_schedule_create_proto_init() {
	if File_schedule_create_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_schedulable_transaction_body_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schedule_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleCreateTransactionBody); i {
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
			RawDescriptor: file_schedule_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schedule_create_proto_goTypes,
		DependencyIndexes: file_schedule_create_proto_depIdxs,
		MessageInfos:      file_schedule_create_proto_msgTypes,
	}.Build()
	File_schedule_create_proto = out.File
	file_schedule_create_proto_rawDesc = nil
	file_schedule_create_proto_goTypes = nil
	file_schedule_create_proto_depIdxs = nil
}
