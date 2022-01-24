// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: token_update.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
// At consensus, updates an already created token to the given values.
//
// If no value is given for a field, that field is left unchanged. For an immutable tokens (that is,
// a token without an admin key), only the expiry may be updated. Setting any other field in that
// case will cause the transaction status to resolve to TOKEN_IS_IMMUTABLE.
//
// --- Signing Requirements ---
// 1. Whether or not a token has an admin key, its expiry can be extended with only the transaction
//    payer's signature.
// 2. Updating any other field of a mutable token requires the admin key's signature.
// 3. If a new admin key is set, this new key must sign <b>unless</b> it is exactly an empty
//    <tt>KeyList</tt>. This special sentinel key removes the existing admin key and causes the
//    token to become immutable. (Other <tt>Key</tt> structures without a constituent
//    <tt>Ed25519</tt> key will be rejected with <tt>INVALID_ADMIN_KEY</tt>.)
// 4. If a new treasury is set, the new treasury account's key must sign the transaction.
//
// --- Nft Requirements ---
// 1. If a non fungible token has a positive treasury balance, the operation will abort with
//    CURRENT_TREASURY_STILL_OWNS_NFTS.
type TokenUpdateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The Token to be updated
	Token *TokenID `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//*
	// The new publicly visible Token symbol, limited to a UTF-8 encoding of length
	// <tt>tokens.maxTokenNameUtf8Bytes</tt>.
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	//*
	// The new publicly visible name of the Token, limited to a UTF-8 encoding of length
	// <tt>tokens.maxSymbolUtf8Bytes</tt>.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	//*
	// The new Treasury account of the Token. If the provided treasury account is not existing or
	// deleted, the response will be INVALID_TREASURY_ACCOUNT_FOR_TOKEN. If successful, the Token
	// balance held in the previous Treasury Account is transferred to the new one.
	Treasury *AccountID `protobuf:"bytes,4,opt,name=treasury,proto3" json:"treasury,omitempty"`
	//*
	// The new admin key of the Token. If Token is immutable, transaction will resolve to
	// TOKEN_IS_IMMUTABlE.
	AdminKey *Key `protobuf:"bytes,5,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// The new KYC key of the Token. If Token does not have currently a KYC key, transaction will
	// resolve to TOKEN_HAS_NO_KYC_KEY.
	KycKey *Key `protobuf:"bytes,6,opt,name=kycKey,proto3" json:"kycKey,omitempty"`
	//*
	// The new Freeze key of the Token. If the Token does not have currently a Freeze key,
	// transaction will resolve to TOKEN_HAS_NO_FREEZE_KEY.
	FreezeKey *Key `protobuf:"bytes,7,opt,name=freezeKey,proto3" json:"freezeKey,omitempty"`
	//*
	// The new Wipe key of the Token. If the Token does not have currently a Wipe key, transaction
	// will resolve to TOKEN_HAS_NO_WIPE_KEY.
	WipeKey *Key `protobuf:"bytes,8,opt,name=wipeKey,proto3" json:"wipeKey,omitempty"`
	//*
	// The new Supply key of the Token. If the Token does not have currently a Supply key,
	// transaction will resolve to TOKEN_HAS_NO_SUPPLY_KEY.
	SupplyKey *Key `protobuf:"bytes,9,opt,name=supplyKey,proto3" json:"supplyKey,omitempty"`
	//*
	// The new account which will be automatically charged to renew the token's expiration, at
	// autoRenewPeriod interval.
	AutoRenewAccount *AccountID `protobuf:"bytes,10,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	//*
	// The new interval at which the auto-renew account will be charged to extend the token's
	// expiry.
	AutoRenewPeriod *Duration `protobuf:"bytes,11,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// The new expiry time of the token. Expiry can be updated even if admin key is not set. If the
	// provided expiry is earlier than the current token expiry, transaction wil resolve to
	// INVALID_EXPIRATION_TIME
	Expiry *Timestamp `protobuf:"bytes,12,opt,name=expiry,proto3" json:"expiry,omitempty"`
	//*
	// If set, the new memo to be associated with the token (UTF-8 encoding max 100 bytes)
	Memo *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// If set, the new key to use to update the token's custom fee schedule; if the token does not
	// currently have this key, transaction will resolve to TOKEN_HAS_NO_FEE_SCHEDULE_KEY
	FeeScheduleKey *Key `protobuf:"bytes,14,opt,name=fee_schedule_key,json=feeScheduleKey,proto3" json:"fee_schedule_key,omitempty"`
	//*
	// The Key which can pause and unpause the Token. If the Token does not currently have a pause key,
	// transaction will resolve to TOKEN_HAS_NO_PAUSE_KEY
	PauseKey *Key `protobuf:"bytes,15,opt,name=pause_key,json=pauseKey,proto3" json:"pause_key,omitempty"`
}

func (x *TokenUpdateTransactionBody) Reset() {
	*x = TokenUpdateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenUpdateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenUpdateTransactionBody) ProtoMessage() {}

func (x *TokenUpdateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_token_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenUpdateTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenUpdateTransactionBody) Descriptor() ([]byte, []int) {
	return file_token_update_proto_rawDescGZIP(), []int{0}
}

func (x *TokenUpdateTransactionBody) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenUpdateTransactionBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenUpdateTransactionBody) GetTreasury() *AccountID {
	if x != nil {
		return x.Treasury
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetKycKey() *Key {
	if x != nil {
		return x.KycKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetFreezeKey() *Key {
	if x != nil {
		return x.FreezeKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetWipeKey() *Key {
	if x != nil {
		return x.WipeKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetSupplyKey() *Key {
	if x != nil {
		return x.SupplyKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetExpiry() *Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetMemo() *wrapperspb.StringValue {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetFeeScheduleKey() *Key {
	if x != nil {
		return x.FeeScheduleKey
	}
	return nil
}

func (x *TokenUpdateTransactionBody) GetPauseKey() *Key {
	if x != nil {
		return x.PauseKey
	}
	return nil
}

var File_token_update_proto protoreflect.FileDescriptor

var file_token_update_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x05, 0x0a, 0x1a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x08, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x26,
	0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x72,
	0x65, 0x65, 0x7a, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65,
	0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x28, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x34, 0x0a, 0x10, 0x66, 0x65, 0x65,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x0e, 0x66, 0x65, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x27, 0x0a, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08,
	0x70, 0x61, 0x75, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_update_proto_rawDescOnce sync.Once
	file_token_update_proto_rawDescData = file_token_update_proto_rawDesc
)

func file_token_update_proto_rawDescGZIP() []byte {
	file_token_update_proto_rawDescOnce.Do(func() {
		file_token_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_update_proto_rawDescData)
	})
	return file_token_update_proto_rawDescData
}

var file_token_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_update_proto_goTypes = []interface{}{
	(*TokenUpdateTransactionBody)(nil), // 0: proto.TokenUpdateTransactionBody
	(*TokenID)(nil),                    // 1: proto.TokenID
	(*AccountID)(nil),                  // 2: proto.AccountID
	(*Key)(nil),                        // 3: proto.Key
	(*Duration)(nil),                   // 4: proto.Duration
	(*Timestamp)(nil),                  // 5: proto.Timestamp
	(*wrapperspb.StringValue)(nil),     // 6: google.protobuf.StringValue
}
var file_token_update_proto_depIdxs = []int32{
	1,  // 0: proto.TokenUpdateTransactionBody.token:type_name -> proto.TokenID
	2,  // 1: proto.TokenUpdateTransactionBody.treasury:type_name -> proto.AccountID
	3,  // 2: proto.TokenUpdateTransactionBody.adminKey:type_name -> proto.Key
	3,  // 3: proto.TokenUpdateTransactionBody.kycKey:type_name -> proto.Key
	3,  // 4: proto.TokenUpdateTransactionBody.freezeKey:type_name -> proto.Key
	3,  // 5: proto.TokenUpdateTransactionBody.wipeKey:type_name -> proto.Key
	3,  // 6: proto.TokenUpdateTransactionBody.supplyKey:type_name -> proto.Key
	2,  // 7: proto.TokenUpdateTransactionBody.autoRenewAccount:type_name -> proto.AccountID
	4,  // 8: proto.TokenUpdateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	5,  // 9: proto.TokenUpdateTransactionBody.expiry:type_name -> proto.Timestamp
	6,  // 10: proto.TokenUpdateTransactionBody.memo:type_name -> google.protobuf.StringValue
	3,  // 11: proto.TokenUpdateTransactionBody.fee_schedule_key:type_name -> proto.Key
	3,  // 12: proto.TokenUpdateTransactionBody.pause_key:type_name -> proto.Key
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_token_update_proto_init() }
func file_token_update_proto_init() {
	if File_token_update_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenUpdateTransactionBody); i {
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
			RawDescriptor: file_token_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_update_proto_goTypes,
		DependencyIndexes: file_token_update_proto_depIdxs,
		MessageInfos:      file_token_update_proto_msgTypes,
	}.Build()
	File_token_update_proto = out.File
	file_token_update_proto_rawDesc = nil
	file_token_update_proto_goTypes = nil
	file_token_update_proto_depIdxs = nil
}
