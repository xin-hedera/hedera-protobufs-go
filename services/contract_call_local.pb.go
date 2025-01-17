// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: contract_call_local.proto

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
// The log information for an event returned by a smart contract function call. One function call
// may return several such events.
type ContractLoginfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// address of a contract that emitted the event
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	//*
	// bloom filter for a particular log
	Bloom []byte `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`
	//*
	// topics of a particular event
	Topic [][]byte `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	//*
	// event data
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ContractLoginfo) Reset() {
	*x = ContractLoginfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_call_local_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractLoginfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractLoginfo) ProtoMessage() {}

func (x *ContractLoginfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_call_local_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractLoginfo.ProtoReflect.Descriptor instead.
func (*ContractLoginfo) Descriptor() ([]byte, []int) {
	return file_contract_call_local_proto_rawDescGZIP(), []int{0}
}

func (x *ContractLoginfo) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractLoginfo) GetBloom() []byte {
	if x != nil {
		return x.Bloom
	}
	return nil
}

func (x *ContractLoginfo) GetTopic() [][]byte {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *ContractLoginfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//*
// The result returned by a call to a smart contract function. This is part of the response to a
// ContractCallLocal query, and is in the record for a ContractCall or ContractCreateInstance
// transaction. The ContractCreateInstance transaction record has the results of the call to the
// constructor.
type ContractFunctionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// the smart contract instance whose function was called
	ContractID *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	//*
	// the result returned by the function
	ContractCallResult []byte `protobuf:"bytes,2,opt,name=contractCallResult,proto3" json:"contractCallResult,omitempty"`
	//*
	// message In case there was an error during smart contract execution
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	//*
	// bloom filter for record
	Bloom []byte `protobuf:"bytes,4,opt,name=bloom,proto3" json:"bloom,omitempty"`
	//*
	// units of gas used to execute contract
	GasUsed uint64 `protobuf:"varint,5,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	//*
	// the log info for events returned by the function
	LogInfo []*ContractLoginfo `protobuf:"bytes,6,rep,name=logInfo,proto3" json:"logInfo,omitempty"`
	//*
	// the list of smart contracts that were created by the function call
	CreatedContractIDs []*ContractID `protobuf:"bytes,7,rep,name=createdContractIDs,proto3" json:"createdContractIDs,omitempty"`
}

func (x *ContractFunctionResult) Reset() {
	*x = ContractFunctionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_call_local_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractFunctionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractFunctionResult) ProtoMessage() {}

func (x *ContractFunctionResult) ProtoReflect() protoreflect.Message {
	mi := &file_contract_call_local_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractFunctionResult.ProtoReflect.Descriptor instead.
func (*ContractFunctionResult) Descriptor() ([]byte, []int) {
	return file_contract_call_local_proto_rawDescGZIP(), []int{1}
}

func (x *ContractFunctionResult) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractFunctionResult) GetContractCallResult() []byte {
	if x != nil {
		return x.ContractCallResult
	}
	return nil
}

func (x *ContractFunctionResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ContractFunctionResult) GetBloom() []byte {
	if x != nil {
		return x.Bloom
	}
	return nil
}

func (x *ContractFunctionResult) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *ContractFunctionResult) GetLogInfo() []*ContractLoginfo {
	if x != nil {
		return x.LogInfo
	}
	return nil
}

func (x *ContractFunctionResult) GetCreatedContractIDs() []*ContractID {
	if x != nil {
		return x.CreatedContractIDs
	}
	return nil
}

//*
// Call a function of the given smart contract instance, giving it functionParameters as its inputs.
// This is performed locally on the particular node that the client is communicating with.
// It cannot change the state of the contract instance (and so, cannot spend anything from the instance's cryptocurrency account).
// It will not have a consensus timestamp. It cannot generate a record or a receipt. The response will contain the output
// returned by the function call.  This is useful for calling getter functions, which purely read the state and don't change it.
// It is faster and cheaper than a normal call, because it is purely local to a single  node.
//
// Unlike a ContractCall transaction, the node will consume the entire amount of provided gas in determining
// the fee for this query.
type ContractCallLocalQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither). The payment must cover the fees and all of the gas offered.
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// the contract instance to call, in the format used in transactions
	ContractID *ContractID `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	//*
	// The amount of gas to use for the call; all of the gas offered will be used and charged a corresponding fee
	Gas int64 `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	//*
	// which function to call, and the parameters to pass to the function
	FunctionParameters []byte `protobuf:"bytes,4,opt,name=functionParameters,proto3" json:"functionParameters,omitempty"`
	//*
	// max number of bytes that the result might include. The run will fail if it would have returned more than this number of bytes.
	//
	// Deprecated: Do not use.
	MaxResultSize int64 `protobuf:"varint,5,opt,name=maxResultSize,proto3" json:"maxResultSize,omitempty"`
}

func (x *ContractCallLocalQuery) Reset() {
	*x = ContractCallLocalQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_call_local_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractCallLocalQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractCallLocalQuery) ProtoMessage() {}

func (x *ContractCallLocalQuery) ProtoReflect() protoreflect.Message {
	mi := &file_contract_call_local_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractCallLocalQuery.ProtoReflect.Descriptor instead.
func (*ContractCallLocalQuery) Descriptor() ([]byte, []int) {
	return file_contract_call_local_proto_rawDescGZIP(), []int{2}
}

func (x *ContractCallLocalQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractCallLocalQuery) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractCallLocalQuery) GetGas() int64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *ContractCallLocalQuery) GetFunctionParameters() []byte {
	if x != nil {
		return x.FunctionParameters
	}
	return nil
}

// Deprecated: Do not use.
func (x *ContractCallLocalQuery) GetMaxResultSize() int64 {
	if x != nil {
		return x.MaxResultSize
	}
	return 0
}

//*
// Response when the client sends the node ContractCallLocalQuery
type ContractCallLocalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// the value returned by the function (if it completed and didn't fail)
	FunctionResult *ContractFunctionResult `protobuf:"bytes,2,opt,name=functionResult,proto3" json:"functionResult,omitempty"`
}

func (x *ContractCallLocalResponse) Reset() {
	*x = ContractCallLocalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_call_local_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractCallLocalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractCallLocalResponse) ProtoMessage() {}

func (x *ContractCallLocalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_call_local_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractCallLocalResponse.ProtoReflect.Descriptor instead.
func (*ContractCallLocalResponse) Descriptor() ([]byte, []int) {
	return file_contract_call_local_proto_rawDescGZIP(), []int{3}
}

func (x *ContractCallLocalResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractCallLocalResponse) GetFunctionResult() *ContractFunctionResult {
	if x != nil {
		return x.FunctionResult
	}
	return nil
}

var File_contract_call_local_proto protoreflect.FileDescriptor

var file_contract_call_local_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc4, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f,
	0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x6c, 0x6f, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x73, 0x22, 0xe3,
	0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_call_local_proto_rawDescOnce sync.Once
	file_contract_call_local_proto_rawDescData = file_contract_call_local_proto_rawDesc
)

func file_contract_call_local_proto_rawDescGZIP() []byte {
	file_contract_call_local_proto_rawDescOnce.Do(func() {
		file_contract_call_local_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_call_local_proto_rawDescData)
	})
	return file_contract_call_local_proto_rawDescData
}

var file_contract_call_local_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contract_call_local_proto_goTypes = []interface{}{
	(*ContractLoginfo)(nil),           // 0: proto.ContractLoginfo
	(*ContractFunctionResult)(nil),    // 1: proto.ContractFunctionResult
	(*ContractCallLocalQuery)(nil),    // 2: proto.ContractCallLocalQuery
	(*ContractCallLocalResponse)(nil), // 3: proto.ContractCallLocalResponse
	(*ContractID)(nil),                // 4: proto.ContractID
	(*QueryHeader)(nil),               // 5: proto.QueryHeader
	(*ResponseHeader)(nil),            // 6: proto.ResponseHeader
}
var file_contract_call_local_proto_depIdxs = []int32{
	4, // 0: proto.ContractLoginfo.contractID:type_name -> proto.ContractID
	4, // 1: proto.ContractFunctionResult.contractID:type_name -> proto.ContractID
	0, // 2: proto.ContractFunctionResult.logInfo:type_name -> proto.ContractLoginfo
	4, // 3: proto.ContractFunctionResult.createdContractIDs:type_name -> proto.ContractID
	5, // 4: proto.ContractCallLocalQuery.header:type_name -> proto.QueryHeader
	4, // 5: proto.ContractCallLocalQuery.contractID:type_name -> proto.ContractID
	6, // 6: proto.ContractCallLocalResponse.header:type_name -> proto.ResponseHeader
	1, // 7: proto.ContractCallLocalResponse.functionResult:type_name -> proto.ContractFunctionResult
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_contract_call_local_proto_init() }
func file_contract_call_local_proto_init() {
	if File_contract_call_local_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_call_local_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractLoginfo); i {
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
		file_contract_call_local_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractFunctionResult); i {
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
		file_contract_call_local_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractCallLocalQuery); i {
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
		file_contract_call_local_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractCallLocalResponse); i {
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
			RawDescriptor: file_contract_call_local_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_call_local_proto_goTypes,
		DependencyIndexes: file_contract_call_local_proto_depIdxs,
		MessageInfos:      file_contract_call_local_proto_msgTypes,
	}.Build()
	File_contract_call_local_proto = out.File
	file_contract_call_local_proto_rawDesc = nil
	file_contract_call_local_proto_goTypes = nil
	file_contract_call_local_proto_depIdxs = nil
}
