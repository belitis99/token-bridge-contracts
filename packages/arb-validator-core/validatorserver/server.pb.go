// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package validatorserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LogInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockHash            string   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber          string   `protobuf:"bytes,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	LogIndex             string   `protobuf:"bytes,5,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Topics               []string `protobuf:"bytes,6,rep,name=topics,proto3" json:"topics,omitempty"`
	TransactionIndex     string   `protobuf:"bytes,7,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionHash      string   `protobuf:"bytes,8,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LogInfo) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *LogInfo) GetBlockNumber() string {
	if m != nil {
		return m.BlockNumber
	}
	return ""
}

func (m *LogInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *LogInfo) GetLogIndex() string {
	if m != nil {
		return m.LogIndex
	}
	return ""
}

func (m *LogInfo) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *LogInfo) GetTransactionIndex() string {
	if m != nil {
		return m.TransactionIndex
	}
	return ""
}

func (m *LogInfo) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

type FindLogsArgs struct {
	FromHeight           string   `protobuf:"bytes,1,opt,name=fromHeight,proto3" json:"fromHeight,omitempty"`
	ToHeight             string   `protobuf:"bytes,2,opt,name=toHeight,proto3" json:"toHeight,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Topics               []string `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindLogsArgs) Reset()         { *m = FindLogsArgs{} }
func (m *FindLogsArgs) String() string { return proto.CompactTextString(m) }
func (*FindLogsArgs) ProtoMessage()    {}
func (*FindLogsArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *FindLogsArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLogsArgs.Unmarshal(m, b)
}
func (m *FindLogsArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLogsArgs.Marshal(b, m, deterministic)
}
func (m *FindLogsArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLogsArgs.Merge(m, src)
}
func (m *FindLogsArgs) XXX_Size() int {
	return xxx_messageInfo_FindLogsArgs.Size(m)
}
func (m *FindLogsArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLogsArgs.DiscardUnknown(m)
}

var xxx_messageInfo_FindLogsArgs proto.InternalMessageInfo

func (m *FindLogsArgs) GetFromHeight() string {
	if m != nil {
		return m.FromHeight
	}
	return ""
}

func (m *FindLogsArgs) GetToHeight() string {
	if m != nil {
		return m.ToHeight
	}
	return ""
}

func (m *FindLogsArgs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FindLogsArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type FindLogsReply struct {
	Logs                 []*LogInfo `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindLogsReply) Reset()         { *m = FindLogsReply{} }
func (m *FindLogsReply) String() string { return proto.CompactTextString(m) }
func (*FindLogsReply) ProtoMessage()    {}
func (*FindLogsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *FindLogsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLogsReply.Unmarshal(m, b)
}
func (m *FindLogsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLogsReply.Marshal(b, m, deterministic)
}
func (m *FindLogsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLogsReply.Merge(m, src)
}
func (m *FindLogsReply) XXX_Size() int {
	return xxx_messageInfo_FindLogsReply.Size(m)
}
func (m *FindLogsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLogsReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindLogsReply proto.InternalMessageInfo

func (m *FindLogsReply) GetLogs() []*LogInfo {
	if m != nil {
		return m.Logs
	}
	return nil
}

type GetOutputMessageArgs struct {
	AssertionNodeHash    string   `protobuf:"bytes,1,opt,name=AssertionNodeHash,proto3" json:"AssertionNodeHash,omitempty"`
	MsgIndex             string   `protobuf:"bytes,2,opt,name=MsgIndex,proto3" json:"MsgIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOutputMessageArgs) Reset()         { *m = GetOutputMessageArgs{} }
func (m *GetOutputMessageArgs) String() string { return proto.CompactTextString(m) }
func (*GetOutputMessageArgs) ProtoMessage()    {}
func (*GetOutputMessageArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *GetOutputMessageArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOutputMessageArgs.Unmarshal(m, b)
}
func (m *GetOutputMessageArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOutputMessageArgs.Marshal(b, m, deterministic)
}
func (m *GetOutputMessageArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOutputMessageArgs.Merge(m, src)
}
func (m *GetOutputMessageArgs) XXX_Size() int {
	return xxx_messageInfo_GetOutputMessageArgs.Size(m)
}
func (m *GetOutputMessageArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOutputMessageArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetOutputMessageArgs proto.InternalMessageInfo

func (m *GetOutputMessageArgs) GetAssertionNodeHash() string {
	if m != nil {
		return m.AssertionNodeHash
	}
	return ""
}

func (m *GetOutputMessageArgs) GetMsgIndex() string {
	if m != nil {
		return m.MsgIndex
	}
	return ""
}

type GetOutputMessageReply struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	RawVal               string   `protobuf:"bytes,2,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOutputMessageReply) Reset()         { *m = GetOutputMessageReply{} }
func (m *GetOutputMessageReply) String() string { return proto.CompactTextString(m) }
func (*GetOutputMessageReply) ProtoMessage()    {}
func (*GetOutputMessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *GetOutputMessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOutputMessageReply.Unmarshal(m, b)
}
func (m *GetOutputMessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOutputMessageReply.Marshal(b, m, deterministic)
}
func (m *GetOutputMessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOutputMessageReply.Merge(m, src)
}
func (m *GetOutputMessageReply) XXX_Size() int {
	return xxx_messageInfo_GetOutputMessageReply.Size(m)
}
func (m *GetOutputMessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOutputMessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetOutputMessageReply proto.InternalMessageInfo

func (m *GetOutputMessageReply) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *GetOutputMessageReply) GetRawVal() string {
	if m != nil {
		return m.RawVal
	}
	return ""
}

type GetMessageResultArgs struct {
	TxHash               string   `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResultArgs) Reset()         { *m = GetMessageResultArgs{} }
func (m *GetMessageResultArgs) String() string { return proto.CompactTextString(m) }
func (*GetMessageResultArgs) ProtoMessage()    {}
func (*GetMessageResultArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *GetMessageResultArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResultArgs.Unmarshal(m, b)
}
func (m *GetMessageResultArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResultArgs.Marshal(b, m, deterministic)
}
func (m *GetMessageResultArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResultArgs.Merge(m, src)
}
func (m *GetMessageResultArgs) XXX_Size() int {
	return xxx_messageInfo_GetMessageResultArgs.Size(m)
}
func (m *GetMessageResultArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResultArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResultArgs proto.InternalMessageInfo

func (m *GetMessageResultArgs) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type GetMessageResultReply struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	RawVal               string   `protobuf:"bytes,2,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
	LogPreHash           string   `protobuf:"bytes,3,opt,name=logPreHash,proto3" json:"logPreHash,omitempty"`
	LogPostHash          string   `protobuf:"bytes,4,opt,name=logPostHash,proto3" json:"logPostHash,omitempty"`
	LogValHashes         []string `protobuf:"bytes,5,rep,name=logValHashes,proto3" json:"logValHashes,omitempty"`
	OnChainTxHash        string   `protobuf:"bytes,6,opt,name=onChainTxHash,proto3" json:"onChainTxHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResultReply) Reset()         { *m = GetMessageResultReply{} }
func (m *GetMessageResultReply) String() string { return proto.CompactTextString(m) }
func (*GetMessageResultReply) ProtoMessage()    {}
func (*GetMessageResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *GetMessageResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResultReply.Unmarshal(m, b)
}
func (m *GetMessageResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResultReply.Marshal(b, m, deterministic)
}
func (m *GetMessageResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResultReply.Merge(m, src)
}
func (m *GetMessageResultReply) XXX_Size() int {
	return xxx_messageInfo_GetMessageResultReply.Size(m)
}
func (m *GetMessageResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResultReply proto.InternalMessageInfo

func (m *GetMessageResultReply) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *GetMessageResultReply) GetRawVal() string {
	if m != nil {
		return m.RawVal
	}
	return ""
}

func (m *GetMessageResultReply) GetLogPreHash() string {
	if m != nil {
		return m.LogPreHash
	}
	return ""
}

func (m *GetMessageResultReply) GetLogPostHash() string {
	if m != nil {
		return m.LogPostHash
	}
	return ""
}

func (m *GetMessageResultReply) GetLogValHashes() []string {
	if m != nil {
		return m.LogValHashes
	}
	return nil
}

func (m *GetMessageResultReply) GetOnChainTxHash() string {
	if m != nil {
		return m.OnChainTxHash
	}
	return ""
}

type GetAssertionCountArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssertionCountArgs) Reset()         { *m = GetAssertionCountArgs{} }
func (m *GetAssertionCountArgs) String() string { return proto.CompactTextString(m) }
func (*GetAssertionCountArgs) ProtoMessage()    {}
func (*GetAssertionCountArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *GetAssertionCountArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssertionCountArgs.Unmarshal(m, b)
}
func (m *GetAssertionCountArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssertionCountArgs.Marshal(b, m, deterministic)
}
func (m *GetAssertionCountArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssertionCountArgs.Merge(m, src)
}
func (m *GetAssertionCountArgs) XXX_Size() int {
	return xxx_messageInfo_GetAssertionCountArgs.Size(m)
}
func (m *GetAssertionCountArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssertionCountArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssertionCountArgs proto.InternalMessageInfo

type GetAssertionCountReply struct {
	AssertionCount       int32    `protobuf:"varint,1,opt,name=assertionCount,proto3" json:"assertionCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssertionCountReply) Reset()         { *m = GetAssertionCountReply{} }
func (m *GetAssertionCountReply) String() string { return proto.CompactTextString(m) }
func (*GetAssertionCountReply) ProtoMessage()    {}
func (*GetAssertionCountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{8}
}

func (m *GetAssertionCountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssertionCountReply.Unmarshal(m, b)
}
func (m *GetAssertionCountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssertionCountReply.Marshal(b, m, deterministic)
}
func (m *GetAssertionCountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssertionCountReply.Merge(m, src)
}
func (m *GetAssertionCountReply) XXX_Size() int {
	return xxx_messageInfo_GetAssertionCountReply.Size(m)
}
func (m *GetAssertionCountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssertionCountReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssertionCountReply proto.InternalMessageInfo

func (m *GetAssertionCountReply) GetAssertionCount() int32 {
	if m != nil {
		return m.AssertionCount
	}
	return 0
}

type GetVMInfoArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoArgs) Reset()         { *m = GetVMInfoArgs{} }
func (m *GetVMInfoArgs) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoArgs) ProtoMessage()    {}
func (*GetVMInfoArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{9}
}

func (m *GetVMInfoArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoArgs.Unmarshal(m, b)
}
func (m *GetVMInfoArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoArgs.Marshal(b, m, deterministic)
}
func (m *GetVMInfoArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoArgs.Merge(m, src)
}
func (m *GetVMInfoArgs) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoArgs.Size(m)
}
func (m *GetVMInfoArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoArgs proto.InternalMessageInfo

type GetVMInfoReply struct {
	VmID                 string   `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoReply) Reset()         { *m = GetVMInfoReply{} }
func (m *GetVMInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoReply) ProtoMessage()    {}
func (*GetVMInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{10}
}

func (m *GetVMInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoReply.Unmarshal(m, b)
}
func (m *GetVMInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoReply.Marshal(b, m, deterministic)
}
func (m *GetVMInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoReply.Merge(m, src)
}
func (m *GetVMInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoReply.Size(m)
}
func (m *GetVMInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoReply proto.InternalMessageInfo

func (m *GetVMInfoReply) GetVmID() string {
	if m != nil {
		return m.VmID
	}
	return ""
}

type CallMessageArgs struct {
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallMessageArgs) Reset()         { *m = CallMessageArgs{} }
func (m *CallMessageArgs) String() string { return proto.CompactTextString(m) }
func (*CallMessageArgs) ProtoMessage()    {}
func (*CallMessageArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{11}
}

func (m *CallMessageArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallMessageArgs.Unmarshal(m, b)
}
func (m *CallMessageArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallMessageArgs.Marshal(b, m, deterministic)
}
func (m *CallMessageArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallMessageArgs.Merge(m, src)
}
func (m *CallMessageArgs) XXX_Size() int {
	return xxx_messageInfo_CallMessageArgs.Size(m)
}
func (m *CallMessageArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_CallMessageArgs.DiscardUnknown(m)
}

var xxx_messageInfo_CallMessageArgs proto.InternalMessageInfo

func (m *CallMessageArgs) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *CallMessageArgs) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *CallMessageArgs) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CallMessageReply struct {
	RawVal               string   `protobuf:"bytes,1,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallMessageReply) Reset()         { *m = CallMessageReply{} }
func (m *CallMessageReply) String() string { return proto.CompactTextString(m) }
func (*CallMessageReply) ProtoMessage()    {}
func (*CallMessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{12}
}

func (m *CallMessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallMessageReply.Unmarshal(m, b)
}
func (m *CallMessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallMessageReply.Marshal(b, m, deterministic)
}
func (m *CallMessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallMessageReply.Merge(m, src)
}
func (m *CallMessageReply) XXX_Size() int {
	return xxx_messageInfo_CallMessageReply.Size(m)
}
func (m *CallMessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CallMessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_CallMessageReply proto.InternalMessageInfo

func (m *CallMessageReply) GetRawVal() string {
	if m != nil {
		return m.RawVal
	}
	return ""
}

func init() {
	proto.RegisterType((*LogInfo)(nil), "validatorserver.LogInfo")
	proto.RegisterType((*FindLogsArgs)(nil), "validatorserver.FindLogsArgs")
	proto.RegisterType((*FindLogsReply)(nil), "validatorserver.FindLogsReply")
	proto.RegisterType((*GetOutputMessageArgs)(nil), "validatorserver.GetOutputMessageArgs")
	proto.RegisterType((*GetOutputMessageReply)(nil), "validatorserver.GetOutputMessageReply")
	proto.RegisterType((*GetMessageResultArgs)(nil), "validatorserver.GetMessageResultArgs")
	proto.RegisterType((*GetMessageResultReply)(nil), "validatorserver.GetMessageResultReply")
	proto.RegisterType((*GetAssertionCountArgs)(nil), "validatorserver.GetAssertionCountArgs")
	proto.RegisterType((*GetAssertionCountReply)(nil), "validatorserver.GetAssertionCountReply")
	proto.RegisterType((*GetVMInfoArgs)(nil), "validatorserver.GetVMInfoArgs")
	proto.RegisterType((*GetVMInfoReply)(nil), "validatorserver.GetVMInfoReply")
	proto.RegisterType((*CallMessageArgs)(nil), "validatorserver.CallMessageArgs")
	proto.RegisterType((*CallMessageReply)(nil), "validatorserver.CallMessageReply")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x4f, 0xdb, 0x4a,
	0x10, 0x56, 0xc8, 0x85, 0x64, 0xb8, 0x04, 0x56, 0x1c, 0x8e, 0x15, 0x9d, 0xc3, 0xc9, 0xb1, 0x28,
	0x8d, 0x10, 0x24, 0x12, 0x7d, 0xae, 0x54, 0x4a, 0x5b, 0x40, 0x0a, 0xb4, 0xb2, 0xaa, 0x3c, 0xf4,
	0xa9, 0x1b, 0x7b, 0xe3, 0x58, 0x6c, 0xbc, 0xd1, 0xee, 0x9a, 0x52, 0xa9, 0x2f, 0xfd, 0x3b, 0xfd,
	0x37, 0xfd, 0x47, 0x95, 0x67, 0x1d, 0xc7, 0x97, 0x50, 0xa4, 0xbe, 0x79, 0xbe, 0x99, 0x9d, 0xcb,
	0x37, 0xdf, 0xae, 0x61, 0x53, 0x31, 0x79, 0xcf, 0x64, 0x7f, 0x2e, 0x85, 0x16, 0xa4, 0x7d, 0x4f,
	0x79, 0xe0, 0x51, 0x2d, 0xa4, 0x81, 0xed, 0xef, 0x6b, 0xb0, 0x3e, 0x14, 0xfe, 0x75, 0x38, 0x11,
	0xc4, 0x82, 0x75, 0xea, 0x79, 0x92, 0x29, 0x65, 0x55, 0xba, 0x95, 0x5e, 0xcb, 0x59, 0x98, 0xe4,
	0x1f, 0x68, 0x8d, 0xb9, 0x70, 0xef, 0xae, 0xa8, 0x9a, 0x5a, 0x6b, 0xe8, 0x5b, 0x02, 0xa4, 0x0b,
	0x1b, 0x68, 0xdc, 0x46, 0xb3, 0x31, 0x93, 0x56, 0x15, 0xfd, 0x59, 0x88, 0x10, 0xa8, 0x79, 0x54,
	0x53, 0xab, 0x86, 0x2e, 0xfc, 0x26, 0x1d, 0x68, 0xf2, 0xb8, 0xb0, 0xc7, 0x1e, 0xac, 0x3a, 0xe2,
	0xa9, 0x4d, 0xf6, 0xa1, 0xa1, 0xc5, 0x3c, 0x70, 0x95, 0xd5, 0xe8, 0x56, 0x7b, 0x2d, 0x27, 0xb1,
	0xc8, 0x31, 0xec, 0x68, 0x49, 0x43, 0x45, 0x5d, 0x1d, 0x88, 0xd0, 0x9c, 0x5d, 0xc7, 0xb3, 0x25,
	0x9c, 0xf4, 0xa0, 0x9d, 0xc1, 0xb0, 0xf3, 0x26, 0x86, 0x16, 0x61, 0xfb, 0x1b, 0x6c, 0xbe, 0x0b,
	0x42, 0x6f, 0x28, 0x7c, 0x75, 0x2e, 0x7d, 0x45, 0x0e, 0x00, 0x26, 0x52, 0xcc, 0xae, 0x58, 0xe0,
	0x4f, 0x75, 0x42, 0x45, 0x06, 0x89, 0x3b, 0xd7, 0x22, 0xf1, 0x1a, 0x32, 0x52, 0x3b, 0xcb, 0x61,
	0x35, 0xcf, 0xe1, 0x72, 0xa6, 0x5a, 0x76, 0x26, 0xfb, 0x25, 0x6c, 0x2d, 0xaa, 0x3b, 0x6c, 0xce,
	0xbf, 0x92, 0x13, 0xa8, 0x71, 0xe1, 0x9b, 0xb0, 0x8d, 0x33, 0xab, 0x5f, 0x58, 0x59, 0x3f, 0x59,
	0x97, 0x83, 0x51, 0xf6, 0x67, 0xd8, 0xbb, 0x64, 0xfa, 0x7d, 0xa4, 0xe7, 0x91, 0xbe, 0x61, 0x4a,
	0x51, 0x9f, 0xe1, 0x10, 0x27, 0xb0, 0x7b, 0xae, 0x14, 0x93, 0xf1, 0x94, 0xb7, 0xc2, 0x63, 0x48,
	0x80, 0x99, 0xa5, 0xec, 0x88, 0x47, 0xba, 0x51, 0xc9, 0x32, 0x92, 0x91, 0x16, 0xb6, 0xfd, 0x16,
	0xfe, 0x2a, 0x56, 0x30, 0x8d, 0xee, 0x41, 0x7d, 0x22, 0xa2, 0xd0, 0xc3, 0xb4, 0x4d, 0xc7, 0x18,
	0xf1, 0x9c, 0x92, 0x7e, 0x19, 0x51, 0x9e, 0x24, 0x4a, 0x2c, 0xbb, 0x8f, 0x8d, 0xa6, 0x09, 0x54,
	0xc4, 0x35, 0x36, 0x1a, 0xf3, 0xf2, 0x90, 0xe9, 0x2e, 0xb1, 0xec, 0x9f, 0x15, 0xac, 0x9b, 0x3b,
	0xf0, 0x07, 0x75, 0xe3, 0x6d, 0x72, 0xe1, 0x7f, 0x90, 0x86, 0x01, 0xb3, 0x94, 0x0c, 0x12, 0xab,
	0x37, 0xb6, 0x84, 0xd2, 0x18, 0x60, 0x24, 0x9a, 0x85, 0x88, 0x0d, 0x9b, 0x5c, 0xf8, 0x23, 0xca,
	0x63, 0x8b, 0x29, 0xab, 0x8e, 0xfb, 0xcb, 0x61, 0xe4, 0x10, 0xb6, 0x44, 0x78, 0x31, 0xa5, 0x41,
	0xf8, 0xd1, 0x0c, 0xd3, 0xc0, 0x3c, 0x79, 0xd0, 0xfe, 0x1b, 0x47, 0x4a, 0xe9, 0xbf, 0x10, 0x51,
	0x88, 0x24, 0xd8, 0xaf, 0x60, 0xbf, 0xe4, 0x30, 0xc3, 0x1e, 0xc1, 0x36, 0xcd, 0xc1, 0x38, 0x75,
	0xdd, 0x29, 0xa0, 0x76, 0x1b, 0xb6, 0x2e, 0x99, 0x1e, 0xdd, 0xc4, 0xd2, 0xc0, 0x94, 0x87, 0xb0,
	0x9d, 0x02, 0x26, 0x15, 0x81, 0xda, 0xfd, 0xec, 0xfa, 0x4d, 0xc2, 0x33, 0x7e, 0xdb, 0x3e, 0xb4,
	0x2f, 0x28, 0xe7, 0x59, 0xe5, 0xf4, 0xa0, 0xed, 0x8a, 0x50, 0x4b, 0xea, 0xea, 0xf3, 0xdc, 0x73,
	0x50, 0x84, 0x63, 0xca, 0x15, 0x0b, 0x3d, 0x26, 0x17, 0x94, 0x1b, 0x2b, 0xbd, 0xee, 0xd5, 0xe5,
	0x75, 0xb7, 0x8f, 0x61, 0x27, 0x53, 0xc8, 0x34, 0xb4, 0x5c, 0x59, 0x25, 0xbb, 0xb2, 0xb3, 0x1f,
	0x35, 0x68, 0x3b, 0x82, 0xf3, 0x68, 0x3e, 0x5a, 0x68, 0x9f, 0x50, 0xd8, 0x29, 0xaa, 0x90, 0x3c,
	0x2b, 0xdd, 0x8d, 0x55, 0x57, 0xa1, 0x73, 0xf4, 0x64, 0x98, 0x69, 0xc7, 0x94, 0xc8, 0x09, 0x6e,
	0x75, 0x89, 0x92, 0x88, 0x57, 0x97, 0x58, 0x21, 0x5d, 0x07, 0x36, 0x32, 0x2c, 0x90, 0x6e, 0xe9,
	0x58, 0x61, 0x19, 0x9d, 0xff, 0x7f, 0x17, 0x61, 0x72, 0x5e, 0x43, 0x73, 0xf1, 0x80, 0x90, 0x7f,
	0x4b, 0xe1, 0xd9, 0x97, 0xad, 0x73, 0xf0, 0xa8, 0xdb, 0xa4, 0xf2, 0x60, 0xb7, 0x24, 0x43, 0xb2,
	0x72, 0xb6, 0xb2, 0x86, 0x3b, 0xcf, 0x9f, 0x8e, 0x33, 0x55, 0x86, 0xd0, 0x4a, 0x95, 0x49, 0x0e,
	0x56, 0x9d, 0x5a, 0xca, 0xb8, 0xf3, 0xdf, 0xe3, 0x7e, 0xcc, 0xf6, 0xfa, 0xf6, 0xd3, 0xd0, 0x0f,
	0xf4, 0x34, 0x1a, 0xf7, 0x5d, 0x31, 0x1b, 0x88, 0xc9, 0xc4, 0x8d, 0x2f, 0x1c, 0xa7, 0x63, 0x35,
	0xa0, 0x72, 0x1c, 0x68, 0x19, 0xcd, 0x06, 0x73, 0xea, 0xde, 0x51, 0x9f, 0x21, 0x72, 0x9a, 0xe6,
	0x3b, 0x75, 0x85, 0x64, 0x83, 0x42, 0xfa, 0x71, 0x03, 0xff, 0x94, 0x2f, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0x04, 0xa3, 0xfe, 0x39, 0x07, 0x00, 0x00,
}
