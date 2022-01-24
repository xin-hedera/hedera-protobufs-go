// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: response.proto

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
// A single response, which is returned from the node to the client, after the client sent the node
// a query. This includes all responses.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*Response_GetByKey
	//	*Response_GetBySolidityID
	//	*Response_ContractCallLocal
	//	*Response_ContractGetBytecodeResponse
	//	*Response_ContractGetInfo
	//	*Response_ContractGetRecordsResponse
	//	*Response_CryptogetAccountBalance
	//	*Response_CryptoGetAccountRecords
	//	*Response_CryptoGetInfo
	//	*Response_CryptoGetLiveHash
	//	*Response_CryptoGetProxyStakers
	//	*Response_FileGetContents
	//	*Response_FileGetInfo
	//	*Response_TransactionGetReceipt
	//	*Response_TransactionGetRecord
	//	*Response_TransactionGetFastRecord
	//	*Response_ConsensusGetTopicInfo
	//	*Response_NetworkGetVersionInfo
	//	*Response_TokenGetInfo
	//	*Response_ScheduleGetInfo
	//	*Response_TokenGetAccountNftInfos
	//	*Response_TokenGetNftInfo
	//	*Response_TokenGetNftInfos
	//	*Response_NetworkGetExecutionTime
	Response isResponse_Response `protobuf_oneof:"response"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *Response) GetGetByKey() *GetByKeyResponse {
	if x, ok := x.GetResponse().(*Response_GetByKey); ok {
		return x.GetByKey
	}
	return nil
}

func (x *Response) GetGetBySolidityID() *GetBySolidityIDResponse {
	if x, ok := x.GetResponse().(*Response_GetBySolidityID); ok {
		return x.GetBySolidityID
	}
	return nil
}

func (x *Response) GetContractCallLocal() *ContractCallLocalResponse {
	if x, ok := x.GetResponse().(*Response_ContractCallLocal); ok {
		return x.ContractCallLocal
	}
	return nil
}

func (x *Response) GetContractGetBytecodeResponse() *ContractGetBytecodeResponse {
	if x, ok := x.GetResponse().(*Response_ContractGetBytecodeResponse); ok {
		return x.ContractGetBytecodeResponse
	}
	return nil
}

func (x *Response) GetContractGetInfo() *ContractGetInfoResponse {
	if x, ok := x.GetResponse().(*Response_ContractGetInfo); ok {
		return x.ContractGetInfo
	}
	return nil
}

func (x *Response) GetContractGetRecordsResponse() *ContractGetRecordsResponse {
	if x, ok := x.GetResponse().(*Response_ContractGetRecordsResponse); ok {
		return x.ContractGetRecordsResponse
	}
	return nil
}

func (x *Response) GetCryptogetAccountBalance() *CryptoGetAccountBalanceResponse {
	if x, ok := x.GetResponse().(*Response_CryptogetAccountBalance); ok {
		return x.CryptogetAccountBalance
	}
	return nil
}

func (x *Response) GetCryptoGetAccountRecords() *CryptoGetAccountRecordsResponse {
	if x, ok := x.GetResponse().(*Response_CryptoGetAccountRecords); ok {
		return x.CryptoGetAccountRecords
	}
	return nil
}

func (x *Response) GetCryptoGetInfo() *CryptoGetInfoResponse {
	if x, ok := x.GetResponse().(*Response_CryptoGetInfo); ok {
		return x.CryptoGetInfo
	}
	return nil
}

func (x *Response) GetCryptoGetLiveHash() *CryptoGetLiveHashResponse {
	if x, ok := x.GetResponse().(*Response_CryptoGetLiveHash); ok {
		return x.CryptoGetLiveHash
	}
	return nil
}

func (x *Response) GetCryptoGetProxyStakers() *CryptoGetStakersResponse {
	if x, ok := x.GetResponse().(*Response_CryptoGetProxyStakers); ok {
		return x.CryptoGetProxyStakers
	}
	return nil
}

func (x *Response) GetFileGetContents() *FileGetContentsResponse {
	if x, ok := x.GetResponse().(*Response_FileGetContents); ok {
		return x.FileGetContents
	}
	return nil
}

func (x *Response) GetFileGetInfo() *FileGetInfoResponse {
	if x, ok := x.GetResponse().(*Response_FileGetInfo); ok {
		return x.FileGetInfo
	}
	return nil
}

func (x *Response) GetTransactionGetReceipt() *TransactionGetReceiptResponse {
	if x, ok := x.GetResponse().(*Response_TransactionGetReceipt); ok {
		return x.TransactionGetReceipt
	}
	return nil
}

func (x *Response) GetTransactionGetRecord() *TransactionGetRecordResponse {
	if x, ok := x.GetResponse().(*Response_TransactionGetRecord); ok {
		return x.TransactionGetRecord
	}
	return nil
}

func (x *Response) GetTransactionGetFastRecord() *TransactionGetFastRecordResponse {
	if x, ok := x.GetResponse().(*Response_TransactionGetFastRecord); ok {
		return x.TransactionGetFastRecord
	}
	return nil
}

func (x *Response) GetConsensusGetTopicInfo() *ConsensusGetTopicInfoResponse {
	if x, ok := x.GetResponse().(*Response_ConsensusGetTopicInfo); ok {
		return x.ConsensusGetTopicInfo
	}
	return nil
}

func (x *Response) GetNetworkGetVersionInfo() *NetworkGetVersionInfoResponse {
	if x, ok := x.GetResponse().(*Response_NetworkGetVersionInfo); ok {
		return x.NetworkGetVersionInfo
	}
	return nil
}

func (x *Response) GetTokenGetInfo() *TokenGetInfoResponse {
	if x, ok := x.GetResponse().(*Response_TokenGetInfo); ok {
		return x.TokenGetInfo
	}
	return nil
}

func (x *Response) GetScheduleGetInfo() *ScheduleGetInfoResponse {
	if x, ok := x.GetResponse().(*Response_ScheduleGetInfo); ok {
		return x.ScheduleGetInfo
	}
	return nil
}

func (x *Response) GetTokenGetAccountNftInfos() *TokenGetAccountNftInfosResponse {
	if x, ok := x.GetResponse().(*Response_TokenGetAccountNftInfos); ok {
		return x.TokenGetAccountNftInfos
	}
	return nil
}

func (x *Response) GetTokenGetNftInfo() *TokenGetNftInfoResponse {
	if x, ok := x.GetResponse().(*Response_TokenGetNftInfo); ok {
		return x.TokenGetNftInfo
	}
	return nil
}

func (x *Response) GetTokenGetNftInfos() *TokenGetNftInfosResponse {
	if x, ok := x.GetResponse().(*Response_TokenGetNftInfos); ok {
		return x.TokenGetNftInfos
	}
	return nil
}

func (x *Response) GetNetworkGetExecutionTime() *NetworkGetExecutionTimeResponse {
	if x, ok := x.GetResponse().(*Response_NetworkGetExecutionTime); ok {
		return x.NetworkGetExecutionTime
	}
	return nil
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_GetByKey struct {
	//*
	// Get all entities associated with a given key
	GetByKey *GetByKeyResponse `protobuf:"bytes,1,opt,name=getByKey,proto3,oneof"`
}

type Response_GetBySolidityID struct {
	//*
	// Get the IDs in the format used in transactions, given the format used in Solidity
	GetBySolidityID *GetBySolidityIDResponse `protobuf:"bytes,2,opt,name=getBySolidityID,proto3,oneof"`
}

type Response_ContractCallLocal struct {
	//*
	// Response to call a function of a smart contract instance
	ContractCallLocal *ContractCallLocalResponse `protobuf:"bytes,3,opt,name=contractCallLocal,proto3,oneof"`
}

type Response_ContractGetBytecodeResponse struct {
	//*
	// Get the bytecode for a smart contract instance
	ContractGetBytecodeResponse *ContractGetBytecodeResponse `protobuf:"bytes,5,opt,name=contractGetBytecodeResponse,proto3,oneof"`
}

type Response_ContractGetInfo struct {
	//*
	// Get information about a smart contract instance
	ContractGetInfo *ContractGetInfoResponse `protobuf:"bytes,4,opt,name=contractGetInfo,proto3,oneof"`
}

type Response_ContractGetRecordsResponse struct {
	//*
	// Get all existing records for a smart contract instance
	ContractGetRecordsResponse *ContractGetRecordsResponse `protobuf:"bytes,6,opt,name=contractGetRecordsResponse,proto3,oneof"`
}

type Response_CryptogetAccountBalance struct {
	//*
	// Get the current balance in a cryptocurrency account
	CryptogetAccountBalance *CryptoGetAccountBalanceResponse `protobuf:"bytes,7,opt,name=cryptogetAccountBalance,proto3,oneof"`
}

type Response_CryptoGetAccountRecords struct {
	//*
	// Get all the records that currently exist for transactions involving an account
	CryptoGetAccountRecords *CryptoGetAccountRecordsResponse `protobuf:"bytes,8,opt,name=cryptoGetAccountRecords,proto3,oneof"`
}

type Response_CryptoGetInfo struct {
	//*
	// Get all information about an account
	CryptoGetInfo *CryptoGetInfoResponse `protobuf:"bytes,9,opt,name=cryptoGetInfo,proto3,oneof"`
}

type Response_CryptoGetLiveHash struct {
	//*
	// Contains a livehash associated to an account
	CryptoGetLiveHash *CryptoGetLiveHashResponse `protobuf:"bytes,10,opt,name=cryptoGetLiveHash,proto3,oneof"`
}

type Response_CryptoGetProxyStakers struct {
	//*
	// Get all the accounts that proxy stake to a given account, and how much they proxy stake
	CryptoGetProxyStakers *CryptoGetStakersResponse `protobuf:"bytes,11,opt,name=cryptoGetProxyStakers,proto3,oneof"`
}

type Response_FileGetContents struct {
	//*
	// Get the contents of a file (the bytes stored in it)
	FileGetContents *FileGetContentsResponse `protobuf:"bytes,12,opt,name=fileGetContents,proto3,oneof"`
}

type Response_FileGetInfo struct {
	//*
	// Get information about a file, such as its expiration date
	FileGetInfo *FileGetInfoResponse `protobuf:"bytes,13,opt,name=fileGetInfo,proto3,oneof"`
}

type Response_TransactionGetReceipt struct {
	//*
	// Get a receipt for a transaction
	TransactionGetReceipt *TransactionGetReceiptResponse `protobuf:"bytes,14,opt,name=transactionGetReceipt,proto3,oneof"`
}

type Response_TransactionGetRecord struct {
	//*
	// Get a record for a transaction
	TransactionGetRecord *TransactionGetRecordResponse `protobuf:"bytes,15,opt,name=transactionGetRecord,proto3,oneof"`
}

type Response_TransactionGetFastRecord struct {
	//*
	// Get a record for a transaction (lasts 180 seconds)
	TransactionGetFastRecord *TransactionGetFastRecordResponse `protobuf:"bytes,16,opt,name=transactionGetFastRecord,proto3,oneof"`
}

type Response_ConsensusGetTopicInfo struct {
	//*
	// Parameters of and state of a consensus topic..
	ConsensusGetTopicInfo *ConsensusGetTopicInfoResponse `protobuf:"bytes,150,opt,name=consensusGetTopicInfo,proto3,oneof"`
}

type Response_NetworkGetVersionInfo struct {
	//*
	// Semantic versions of Hedera Services and HAPI proto
	NetworkGetVersionInfo *NetworkGetVersionInfoResponse `protobuf:"bytes,151,opt,name=networkGetVersionInfo,proto3,oneof"`
}

type Response_TokenGetInfo struct {
	//*
	// Get all information about a token
	TokenGetInfo *TokenGetInfoResponse `protobuf:"bytes,152,opt,name=tokenGetInfo,proto3,oneof"`
}

type Response_ScheduleGetInfo struct {
	//*
	// Get all information about a schedule entity
	ScheduleGetInfo *ScheduleGetInfoResponse `protobuf:"bytes,153,opt,name=scheduleGetInfo,proto3,oneof"`
}

type Response_TokenGetAccountNftInfos struct {
	//*
	// A list of the NFTs associated with the account
	TokenGetAccountNftInfos *TokenGetAccountNftInfosResponse `protobuf:"bytes,154,opt,name=tokenGetAccountNftInfos,proto3,oneof"`
}

type Response_TokenGetNftInfo struct {
	//*
	// All information about an NFT
	TokenGetNftInfo *TokenGetNftInfoResponse `protobuf:"bytes,155,opt,name=tokenGetNftInfo,proto3,oneof"`
}

type Response_TokenGetNftInfos struct {
	//*
	// A list of the NFTs for the token
	TokenGetNftInfos *TokenGetNftInfosResponse `protobuf:"bytes,156,opt,name=tokenGetNftInfos,proto3,oneof"`
}

type Response_NetworkGetExecutionTime struct {
	//*
	// Execution times of "sufficiently recent" transactions
	NetworkGetExecutionTime *NetworkGetExecutionTimeResponse `protobuf:"bytes,157,opt,name=networkGetExecutionTime,proto3,oneof"`
}

func (*Response_GetByKey) isResponse_Response() {}

func (*Response_GetBySolidityID) isResponse_Response() {}

func (*Response_ContractCallLocal) isResponse_Response() {}

func (*Response_ContractGetBytecodeResponse) isResponse_Response() {}

func (*Response_ContractGetInfo) isResponse_Response() {}

func (*Response_ContractGetRecordsResponse) isResponse_Response() {}

func (*Response_CryptogetAccountBalance) isResponse_Response() {}

func (*Response_CryptoGetAccountRecords) isResponse_Response() {}

func (*Response_CryptoGetInfo) isResponse_Response() {}

func (*Response_CryptoGetLiveHash) isResponse_Response() {}

func (*Response_CryptoGetProxyStakers) isResponse_Response() {}

func (*Response_FileGetContents) isResponse_Response() {}

func (*Response_FileGetInfo) isResponse_Response() {}

func (*Response_TransactionGetReceipt) isResponse_Response() {}

func (*Response_TransactionGetRecord) isResponse_Response() {}

func (*Response_TransactionGetFastRecord) isResponse_Response() {}

func (*Response_ConsensusGetTopicInfo) isResponse_Response() {}

func (*Response_NetworkGetVersionInfo) isResponse_Response() {}

func (*Response_TokenGetInfo) isResponse_Response() {}

func (*Response_ScheduleGetInfo) isResponse_Response() {}

func (*Response_TokenGetAccountNftInfos) isResponse_Response() {}

func (*Response_TokenGetNftInfo) isResponse_Response() {}

func (*Response_TokenGetNftInfos) isResponse_Response() {}

func (*Response_NetworkGetExecutionTime) isResponse_Response() {}

var File_response_proto protoreflect.FileDescriptor

var file_response_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x79, 0x5f, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x67,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x61,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x6e, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x10, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x67, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f,
	0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x50, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x66, 0x0a, 0x1b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x63,
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x17,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x0d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x50, 0x0a, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x57, 0x0a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x0f,
	0x66, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x69, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5c, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x59, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x14, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x65, 0x0a, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x18,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x5d, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x97, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x98, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x99, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x63, 0x0a, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x9a, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x00, 0x52, 0x17, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x4b, 0x0a, 0x0f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x9b, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47,
	0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4e, 0x0a, 0x10, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x9c, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65,
	0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x63, 0x0a, 0x17, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x9d, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61,
	0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_response_proto_rawDescOnce sync.Once
	file_response_proto_rawDescData = file_response_proto_rawDesc
)

func file_response_proto_rawDescGZIP() []byte {
	file_response_proto_rawDescOnce.Do(func() {
		file_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_proto_rawDescData)
	})
	return file_response_proto_rawDescData
}

var file_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_proto_goTypes = []interface{}{
	(*Response)(nil),                         // 0: proto.Response
	(*GetByKeyResponse)(nil),                 // 1: proto.GetByKeyResponse
	(*GetBySolidityIDResponse)(nil),          // 2: proto.GetBySolidityIDResponse
	(*ContractCallLocalResponse)(nil),        // 3: proto.ContractCallLocalResponse
	(*ContractGetBytecodeResponse)(nil),      // 4: proto.ContractGetBytecodeResponse
	(*ContractGetInfoResponse)(nil),          // 5: proto.ContractGetInfoResponse
	(*ContractGetRecordsResponse)(nil),       // 6: proto.ContractGetRecordsResponse
	(*CryptoGetAccountBalanceResponse)(nil),  // 7: proto.CryptoGetAccountBalanceResponse
	(*CryptoGetAccountRecordsResponse)(nil),  // 8: proto.CryptoGetAccountRecordsResponse
	(*CryptoGetInfoResponse)(nil),            // 9: proto.CryptoGetInfoResponse
	(*CryptoGetLiveHashResponse)(nil),        // 10: proto.CryptoGetLiveHashResponse
	(*CryptoGetStakersResponse)(nil),         // 11: proto.CryptoGetStakersResponse
	(*FileGetContentsResponse)(nil),          // 12: proto.FileGetContentsResponse
	(*FileGetInfoResponse)(nil),              // 13: proto.FileGetInfoResponse
	(*TransactionGetReceiptResponse)(nil),    // 14: proto.TransactionGetReceiptResponse
	(*TransactionGetRecordResponse)(nil),     // 15: proto.TransactionGetRecordResponse
	(*TransactionGetFastRecordResponse)(nil), // 16: proto.TransactionGetFastRecordResponse
	(*ConsensusGetTopicInfoResponse)(nil),    // 17: proto.ConsensusGetTopicInfoResponse
	(*NetworkGetVersionInfoResponse)(nil),    // 18: proto.NetworkGetVersionInfoResponse
	(*TokenGetInfoResponse)(nil),             // 19: proto.TokenGetInfoResponse
	(*ScheduleGetInfoResponse)(nil),          // 20: proto.ScheduleGetInfoResponse
	(*TokenGetAccountNftInfosResponse)(nil),  // 21: proto.TokenGetAccountNftInfosResponse
	(*TokenGetNftInfoResponse)(nil),          // 22: proto.TokenGetNftInfoResponse
	(*TokenGetNftInfosResponse)(nil),         // 23: proto.TokenGetNftInfosResponse
	(*NetworkGetExecutionTimeResponse)(nil),  // 24: proto.NetworkGetExecutionTimeResponse
}
var file_response_proto_depIdxs = []int32{
	1,  // 0: proto.Response.getByKey:type_name -> proto.GetByKeyResponse
	2,  // 1: proto.Response.getBySolidityID:type_name -> proto.GetBySolidityIDResponse
	3,  // 2: proto.Response.contractCallLocal:type_name -> proto.ContractCallLocalResponse
	4,  // 3: proto.Response.contractGetBytecodeResponse:type_name -> proto.ContractGetBytecodeResponse
	5,  // 4: proto.Response.contractGetInfo:type_name -> proto.ContractGetInfoResponse
	6,  // 5: proto.Response.contractGetRecordsResponse:type_name -> proto.ContractGetRecordsResponse
	7,  // 6: proto.Response.cryptogetAccountBalance:type_name -> proto.CryptoGetAccountBalanceResponse
	8,  // 7: proto.Response.cryptoGetAccountRecords:type_name -> proto.CryptoGetAccountRecordsResponse
	9,  // 8: proto.Response.cryptoGetInfo:type_name -> proto.CryptoGetInfoResponse
	10, // 9: proto.Response.cryptoGetLiveHash:type_name -> proto.CryptoGetLiveHashResponse
	11, // 10: proto.Response.cryptoGetProxyStakers:type_name -> proto.CryptoGetStakersResponse
	12, // 11: proto.Response.fileGetContents:type_name -> proto.FileGetContentsResponse
	13, // 12: proto.Response.fileGetInfo:type_name -> proto.FileGetInfoResponse
	14, // 13: proto.Response.transactionGetReceipt:type_name -> proto.TransactionGetReceiptResponse
	15, // 14: proto.Response.transactionGetRecord:type_name -> proto.TransactionGetRecordResponse
	16, // 15: proto.Response.transactionGetFastRecord:type_name -> proto.TransactionGetFastRecordResponse
	17, // 16: proto.Response.consensusGetTopicInfo:type_name -> proto.ConsensusGetTopicInfoResponse
	18, // 17: proto.Response.networkGetVersionInfo:type_name -> proto.NetworkGetVersionInfoResponse
	19, // 18: proto.Response.tokenGetInfo:type_name -> proto.TokenGetInfoResponse
	20, // 19: proto.Response.scheduleGetInfo:type_name -> proto.ScheduleGetInfoResponse
	21, // 20: proto.Response.tokenGetAccountNftInfos:type_name -> proto.TokenGetAccountNftInfosResponse
	22, // 21: proto.Response.tokenGetNftInfo:type_name -> proto.TokenGetNftInfoResponse
	23, // 22: proto.Response.tokenGetNftInfos:type_name -> proto.TokenGetNftInfosResponse
	24, // 23: proto.Response.networkGetExecutionTime:type_name -> proto.NetworkGetExecutionTimeResponse
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_response_proto_init() }
func file_response_proto_init() {
	if File_response_proto != nil {
		return
	}
	file_get_by_key_proto_init()
	file_get_by_solidity_id_proto_init()
	file_contract_call_local_proto_init()
	file_contract_get_bytecode_proto_init()
	file_contract_get_info_proto_init()
	file_contract_get_records_proto_init()
	file_crypto_get_account_balance_proto_init()
	file_crypto_get_account_records_proto_init()
	file_crypto_get_info_proto_init()
	file_crypto_get_live_hash_proto_init()
	file_crypto_get_stakers_proto_init()
	file_file_get_contents_proto_init()
	file_file_get_info_proto_init()
	file_transaction_get_receipt_proto_init()
	file_transaction_get_record_proto_init()
	file_transaction_get_fast_record_proto_init()
	file_consensus_get_topic_info_proto_init()
	file_network_get_version_info_proto_init()
	file_network_get_execution_time_proto_init()
	file_token_get_account_nft_infos_proto_init()
	file_token_get_info_proto_init()
	file_token_get_nft_info_proto_init()
	file_token_get_nft_infos_proto_init()
	file_schedule_get_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_response_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Response_GetByKey)(nil),
		(*Response_GetBySolidityID)(nil),
		(*Response_ContractCallLocal)(nil),
		(*Response_ContractGetBytecodeResponse)(nil),
		(*Response_ContractGetInfo)(nil),
		(*Response_ContractGetRecordsResponse)(nil),
		(*Response_CryptogetAccountBalance)(nil),
		(*Response_CryptoGetAccountRecords)(nil),
		(*Response_CryptoGetInfo)(nil),
		(*Response_CryptoGetLiveHash)(nil),
		(*Response_CryptoGetProxyStakers)(nil),
		(*Response_FileGetContents)(nil),
		(*Response_FileGetInfo)(nil),
		(*Response_TransactionGetReceipt)(nil),
		(*Response_TransactionGetRecord)(nil),
		(*Response_TransactionGetFastRecord)(nil),
		(*Response_ConsensusGetTopicInfo)(nil),
		(*Response_NetworkGetVersionInfo)(nil),
		(*Response_TokenGetInfo)(nil),
		(*Response_ScheduleGetInfo)(nil),
		(*Response_TokenGetAccountNftInfos)(nil),
		(*Response_TokenGetNftInfo)(nil),
		(*Response_TokenGetNftInfos)(nil),
		(*Response_NetworkGetExecutionTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_proto_goTypes,
		DependencyIndexes: file_response_proto_depIdxs,
		MessageInfos:      file_response_proto_msgTypes,
	}.Build()
	File_response_proto = out.File
	file_response_proto_rawDesc = nil
	file_response_proto_goTypes = nil
	file_response_proto_depIdxs = nil
}
