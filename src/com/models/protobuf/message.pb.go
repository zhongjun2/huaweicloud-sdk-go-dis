// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protobuf

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutRecordsRequest struct {
	StreamName           *string                   `protobuf:"bytes,1,req,name=streamName" json:"streamName,omitempty"`
	Records              []*PutRecordsRequestEntry `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PutRecordsRequest) Reset()         { *m = PutRecordsRequest{} }
func (m *PutRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*PutRecordsRequest) ProtoMessage()    {}
func (*PutRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *PutRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsRequest.Unmarshal(m, b)
}
func (m *PutRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsRequest.Marshal(b, m, deterministic)
}
func (m *PutRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsRequest.Merge(m, src)
}
func (m *PutRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_PutRecordsRequest.Size(m)
}
func (m *PutRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsRequest proto.InternalMessageInfo

func (m *PutRecordsRequest) GetStreamName() string {
	if m != nil && m.StreamName != nil {
		return *m.StreamName
	}
	return ""
}

func (m *PutRecordsRequest) GetRecords() []*PutRecordsRequestEntry {
	if m != nil {
		return m.Records
	}
	return nil
}

type PutRecordsRequestEntry struct {
	Data                 []byte                              `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	ExplicitHashKey      *string                             `protobuf:"bytes,4,opt,name=explicitHashKey" json:"explicitHashKey,omitempty"`
	PartitionKey         *string                             `protobuf:"bytes,5,opt,name=partitionKey" json:"partitionKey,omitempty"`
	PartitionId          *string                             `protobuf:"bytes,6,opt,name=partitionId" json:"partitionId,omitempty"`
	ExtendedInfo         *PutRecordsRequestEntryExtendedInfo `protobuf:"bytes,7,opt,name=extendedInfo" json:"extendedInfo,omitempty"`
	Timestamp            *int64                              `protobuf:"varint,29,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *PutRecordsRequestEntry) Reset()         { *m = PutRecordsRequestEntry{} }
func (m *PutRecordsRequestEntry) String() string { return proto.CompactTextString(m) }
func (*PutRecordsRequestEntry) ProtoMessage()    {}
func (*PutRecordsRequestEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *PutRecordsRequestEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsRequestEntry.Unmarshal(m, b)
}
func (m *PutRecordsRequestEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsRequestEntry.Marshal(b, m, deterministic)
}
func (m *PutRecordsRequestEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsRequestEntry.Merge(m, src)
}
func (m *PutRecordsRequestEntry) XXX_Size() int {
	return xxx_messageInfo_PutRecordsRequestEntry.Size(m)
}
func (m *PutRecordsRequestEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsRequestEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsRequestEntry proto.InternalMessageInfo

func (m *PutRecordsRequestEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PutRecordsRequestEntry) GetExplicitHashKey() string {
	if m != nil && m.ExplicitHashKey != nil {
		return *m.ExplicitHashKey
	}
	return ""
}

func (m *PutRecordsRequestEntry) GetPartitionKey() string {
	if m != nil && m.PartitionKey != nil {
		return *m.PartitionKey
	}
	return ""
}

func (m *PutRecordsRequestEntry) GetPartitionId() string {
	if m != nil && m.PartitionId != nil {
		return *m.PartitionId
	}
	return ""
}

func (m *PutRecordsRequestEntry) GetExtendedInfo() *PutRecordsRequestEntryExtendedInfo {
	if m != nil {
		return m.ExtendedInfo
	}
	return nil
}

func (m *PutRecordsRequestEntry) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type PutRecordsRequestEntryExtendedInfo struct {
	FileName             *string  `protobuf:"bytes,8,req,name=fileName" json:"fileName,omitempty"`
	DeliverDataId        *string  `protobuf:"bytes,9,req,name=deliverDataId" json:"deliverDataId,omitempty"`
	EndFlag              *bool    `protobuf:"varint,10,opt,name=endFlag,def=0" json:"endFlag,omitempty"`
	SeqNum               *int64   `protobuf:"varint,11,req,name=seqNum" json:"seqNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRecordsRequestEntryExtendedInfo) Reset()         { *m = PutRecordsRequestEntryExtendedInfo{} }
func (m *PutRecordsRequestEntryExtendedInfo) String() string { return proto.CompactTextString(m) }
func (*PutRecordsRequestEntryExtendedInfo) ProtoMessage()    {}
func (*PutRecordsRequestEntryExtendedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *PutRecordsRequestEntryExtendedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsRequestEntryExtendedInfo.Unmarshal(m, b)
}
func (m *PutRecordsRequestEntryExtendedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsRequestEntryExtendedInfo.Marshal(b, m, deterministic)
}
func (m *PutRecordsRequestEntryExtendedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsRequestEntryExtendedInfo.Merge(m, src)
}
func (m *PutRecordsRequestEntryExtendedInfo) XXX_Size() int {
	return xxx_messageInfo_PutRecordsRequestEntryExtendedInfo.Size(m)
}
func (m *PutRecordsRequestEntryExtendedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsRequestEntryExtendedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsRequestEntryExtendedInfo proto.InternalMessageInfo

const Default_PutRecordsRequestEntryExtendedInfo_EndFlag bool = false

func (m *PutRecordsRequestEntryExtendedInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *PutRecordsRequestEntryExtendedInfo) GetDeliverDataId() string {
	if m != nil && m.DeliverDataId != nil {
		return *m.DeliverDataId
	}
	return ""
}

func (m *PutRecordsRequestEntryExtendedInfo) GetEndFlag() bool {
	if m != nil && m.EndFlag != nil {
		return *m.EndFlag
	}
	return Default_PutRecordsRequestEntryExtendedInfo_EndFlag
}

func (m *PutRecordsRequestEntryExtendedInfo) GetSeqNum() int64 {
	if m != nil && m.SeqNum != nil {
		return *m.SeqNum
	}
	return 0
}

type PutRecordsResult struct {
	FailedRecordCount    *int32                   `protobuf:"varint,21,req,name=failedRecordCount" json:"failedRecordCount,omitempty"`
	Records              []*PutRecordsResultEntry `protobuf:"bytes,22,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PutRecordsResult) Reset()         { *m = PutRecordsResult{} }
func (m *PutRecordsResult) String() string { return proto.CompactTextString(m) }
func (*PutRecordsResult) ProtoMessage()    {}
func (*PutRecordsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *PutRecordsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsResult.Unmarshal(m, b)
}
func (m *PutRecordsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsResult.Marshal(b, m, deterministic)
}
func (m *PutRecordsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsResult.Merge(m, src)
}
func (m *PutRecordsResult) XXX_Size() int {
	return xxx_messageInfo_PutRecordsResult.Size(m)
}
func (m *PutRecordsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsResult.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsResult proto.InternalMessageInfo

func (m *PutRecordsResult) GetFailedRecordCount() int32 {
	if m != nil && m.FailedRecordCount != nil {
		return *m.FailedRecordCount
	}
	return 0
}

func (m *PutRecordsResult) GetRecords() []*PutRecordsResultEntry {
	if m != nil {
		return m.Records
	}
	return nil
}

type PutRecordsResultEntry struct {
	ShardId              *string  `protobuf:"bytes,23,opt,name=shardId" json:"shardId,omitempty"`
	SequenceNumber       *string  `protobuf:"bytes,24,opt,name=sequenceNumber" json:"sequenceNumber,omitempty"`
	ErrorCode            *string  `protobuf:"bytes,25,opt,name=errorCode" json:"errorCode,omitempty"`
	ErrorMessage         *string  `protobuf:"bytes,26,opt,name=errorMessage" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRecordsResultEntry) Reset()         { *m = PutRecordsResultEntry{} }
func (m *PutRecordsResultEntry) String() string { return proto.CompactTextString(m) }
func (*PutRecordsResultEntry) ProtoMessage()    {}
func (*PutRecordsResultEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *PutRecordsResultEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsResultEntry.Unmarshal(m, b)
}
func (m *PutRecordsResultEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsResultEntry.Marshal(b, m, deterministic)
}
func (m *PutRecordsResultEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsResultEntry.Merge(m, src)
}
func (m *PutRecordsResultEntry) XXX_Size() int {
	return xxx_messageInfo_PutRecordsResultEntry.Size(m)
}
func (m *PutRecordsResultEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsResultEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsResultEntry proto.InternalMessageInfo

func (m *PutRecordsResultEntry) GetShardId() string {
	if m != nil && m.ShardId != nil {
		return *m.ShardId
	}
	return ""
}

func (m *PutRecordsResultEntry) GetSequenceNumber() string {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return ""
}

func (m *PutRecordsResultEntry) GetErrorCode() string {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ""
}

func (m *PutRecordsResultEntry) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

type GetRecordsRequest struct {
	ShardIterator        *string  `protobuf:"bytes,12,req,name=shardIterator" json:"shardIterator,omitempty"`
	Limit                *int32   `protobuf:"varint,13,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecordsRequest) Reset()         { *m = GetRecordsRequest{} }
func (m *GetRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecordsRequest) ProtoMessage()    {}
func (*GetRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *GetRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordsRequest.Unmarshal(m, b)
}
func (m *GetRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordsRequest.Marshal(b, m, deterministic)
}
func (m *GetRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordsRequest.Merge(m, src)
}
func (m *GetRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecordsRequest.Size(m)
}
func (m *GetRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordsRequest proto.InternalMessageInfo

func (m *GetRecordsRequest) GetShardIterator() string {
	if m != nil && m.ShardIterator != nil {
		return *m.ShardIterator
	}
	return ""
}

func (m *GetRecordsRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type GetRecordsResult struct {
	NextShardIterator *string `protobuf:"bytes,14,req,name=nextShardIterator" json:"nextShardIterator,omitempty"`
	//    int64 millisBehindLatest =15;
	Records              []*Record `protobuf:"bytes,16,rep,name=records" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetRecordsResult) Reset()         { *m = GetRecordsResult{} }
func (m *GetRecordsResult) String() string { return proto.CompactTextString(m) }
func (*GetRecordsResult) ProtoMessage()    {}
func (*GetRecordsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *GetRecordsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordsResult.Unmarshal(m, b)
}
func (m *GetRecordsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordsResult.Marshal(b, m, deterministic)
}
func (m *GetRecordsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordsResult.Merge(m, src)
}
func (m *GetRecordsResult) XXX_Size() int {
	return xxx_messageInfo_GetRecordsResult.Size(m)
}
func (m *GetRecordsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordsResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordsResult proto.InternalMessageInfo

func (m *GetRecordsResult) GetNextShardIterator() string {
	if m != nil && m.NextShardIterator != nil {
		return *m.NextShardIterator
	}
	return ""
}

func (m *GetRecordsResult) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Record struct {
	PartitionKey   *string `protobuf:"bytes,17,opt,name=partitionKey" json:"partitionKey,omitempty"`
	SequenceNumber *string `protobuf:"bytes,18,req,name=sequenceNumber" json:"sequenceNumber,omitempty"`
	Data           []byte  `protobuf:"bytes,19,opt,name=data" json:"data,omitempty"`
	// 	approximateArrivalTimestamp = 20;
	Timestamp            *int64   `protobuf:"varint,27,opt,name=timestamp" json:"timestamp,omitempty"`
	TimestampType        *string  `protobuf:"bytes,28,opt,name=timestampType" json:"timestampType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetPartitionKey() string {
	if m != nil && m.PartitionKey != nil {
		return *m.PartitionKey
	}
	return ""
}

func (m *Record) GetSequenceNumber() string {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return ""
}

func (m *Record) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Record) GetTimestampType() string {
	if m != nil && m.TimestampType != nil {
		return *m.TimestampType
	}
	return ""
}

func init() {
	proto.RegisterType((*PutRecordsRequest)(nil), "protobuf.PutRecordsRequest")
	proto.RegisterType((*PutRecordsRequestEntry)(nil), "protobuf.PutRecordsRequestEntry")
	proto.RegisterType((*PutRecordsRequestEntryExtendedInfo)(nil), "protobuf.PutRecordsRequestEntryExtendedInfo")
	proto.RegisterType((*PutRecordsResult)(nil), "protobuf.PutRecordsResult")
	proto.RegisterType((*PutRecordsResultEntry)(nil), "protobuf.PutRecordsResultEntry")
	proto.RegisterType((*GetRecordsRequest)(nil), "protobuf.GetRecordsRequest")
	proto.RegisterType((*GetRecordsResult)(nil), "protobuf.GetRecordsResult")
	proto.RegisterType((*Record)(nil), "protobuf.Record")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdb, 0x6e, 0x13, 0x31,
	0x14, 0xd4, 0x26, 0xcd, 0xed, 0x24, 0x29, 0x89, 0xa1, 0xc1, 0x94, 0x42, 0xad, 0x55, 0x85, 0x56,
	0xa8, 0xca, 0x43, 0xdf, 0xe8, 0x6b, 0x29, 0x10, 0x21, 0x42, 0x65, 0xf8, 0x01, 0x37, 0x3e, 0x69,
	0x57, 0xda, 0x4b, 0x6a, 0x7b, 0x51, 0x22, 0x7e, 0x81, 0x3f, 0xe0, 0x81, 0x1f, 0xe0, 0x23, 0xd1,
	0x7a, 0xd9, 0xec, 0x25, 0x91, 0xfa, 0x94, 0xcc, 0x78, 0x3c, 0x71, 0xe6, 0x9c, 0x81, 0x61, 0x88,
	0x5a, 0x8b, 0x3b, 0x9c, 0xae, 0x54, 0x6c, 0x62, 0xd2, 0xb5, 0x1f, 0xb7, 0xc9, 0xd2, 0x8d, 0x61,
	0x7c, 0x93, 0x18, 0x8e, 0x8b, 0x58, 0x49, 0xcd, 0xf1, 0x21, 0x41, 0x6d, 0xc8, 0x6b, 0x00, 0x6d,
	0x14, 0x8a, 0x70, 0x2e, 0x42, 0xa4, 0x0e, 0x6b, 0x78, 0x3d, 0x5e, 0x62, 0xc8, 0x25, 0x74, 0x54,
	0x76, 0x83, 0x36, 0x58, 0xd3, 0xeb, 0x5f, 0xb0, 0x69, 0x6e, 0x38, 0xdd, 0x71, 0xbb, 0x8e, 0x8c,
	0xda, 0xf0, 0xfc, 0x82, 0xfb, 0xab, 0x01, 0x93, 0xfd, 0x1a, 0x42, 0xe0, 0x40, 0x0a, 0x23, 0x68,
	0x93, 0x39, 0xde, 0x80, 0xdb, 0xef, 0xc4, 0x83, 0x27, 0xb8, 0x5e, 0x05, 0xfe, 0xc2, 0x37, 0x9f,
	0x84, 0xbe, 0xff, 0x8c, 0x1b, 0x7a, 0xc0, 0x1c, 0xaf, 0xc7, 0xeb, 0x34, 0x71, 0x61, 0xb0, 0x12,
	0xca, 0xf8, 0xc6, 0x8f, 0xa3, 0x54, 0xd6, 0xb2, 0xb2, 0x0a, 0x47, 0x18, 0xf4, 0xb7, 0x78, 0x26,
	0x69, 0xdb, 0x4a, 0xca, 0x14, 0xb9, 0x81, 0x01, 0xae, 0x0d, 0x46, 0x12, 0xe5, 0x2c, 0x5a, 0xc6,
	0xb4, 0xc3, 0x1c, 0xaf, 0x7f, 0x71, 0xfe, 0xd8, 0xff, 0xbb, 0x2e, 0xdd, 0xe1, 0x15, 0x07, 0x72,
	0x02, 0x3d, 0xe3, 0x87, 0xa8, 0x8d, 0x08, 0x57, 0xf4, 0x15, 0x73, 0xbc, 0x26, 0x2f, 0x08, 0xf7,
	0x8f, 0x03, 0xee, 0xe3, 0x96, 0xe4, 0x18, 0xba, 0x4b, 0x3f, 0x40, 0x3b, 0x8f, 0xae, 0x9d, 0xc7,
	0x16, 0x93, 0x33, 0x18, 0x4a, 0x0c, 0xfc, 0x1f, 0xa8, 0xde, 0x0b, 0x23, 0x66, 0x92, 0xf6, 0xac,
	0xa0, 0x4a, 0x92, 0x53, 0xe8, 0x60, 0x24, 0x3f, 0x04, 0xe2, 0x8e, 0x02, 0x73, 0xbc, 0xee, 0x65,
	0x6b, 0x29, 0x02, 0x8d, 0x3c, 0x67, 0xc9, 0x04, 0xda, 0x1a, 0x1f, 0xe6, 0x49, 0x48, 0xfb, 0xac,
	0xe1, 0x35, 0xf9, 0x7f, 0xe4, 0xfe, 0x84, 0x51, 0xf9, 0x81, 0x3a, 0x09, 0x0c, 0x39, 0x87, 0xf1,
	0x52, 0xf8, 0x01, 0xca, 0x8c, 0xbe, 0x8a, 0x93, 0xc8, 0xd0, 0x23, 0xd6, 0xf0, 0x5a, 0x7c, 0xf7,
	0x80, 0xbc, 0x2b, 0xd6, 0x65, 0x62, 0xd7, 0xe5, 0x74, 0x7f, 0x9c, 0xa9, 0x75, 0x6d, 0x5b, 0x7e,
	0x3b, 0x70, 0xb4, 0x57, 0x42, 0x28, 0x74, 0xf4, 0xbd, 0x50, 0x72, 0x26, 0xe9, 0x73, 0x3b, 0xc6,
	0x1c, 0x92, 0x37, 0x70, 0xa8, 0xd3, 0x1c, 0xa3, 0x05, 0xce, 0x93, 0xf0, 0x16, 0x15, 0xa5, 0x56,
	0x50, 0x63, 0xd3, 0xc1, 0xa0, 0x52, 0xb1, 0xba, 0x8a, 0x25, 0xd2, 0x17, 0x56, 0x52, 0x10, 0xe9,
	0x3a, 0x59, 0xf0, 0x25, 0x2b, 0x0e, 0x3d, 0xce, 0xd6, 0xa9, 0xcc, 0xb9, 0x5f, 0x61, 0xfc, 0x11,
	0xeb, 0xe5, 0x39, 0x83, 0x61, 0xf6, 0x12, 0x83, 0x4a, 0x98, 0x58, 0xd1, 0x41, 0x36, 0x8e, 0x0a,
	0x49, 0x9e, 0x41, 0x2b, 0xf0, 0x43, 0xdf, 0xd0, 0x21, 0x73, 0xbc, 0x16, 0xcf, 0x80, 0x1b, 0xc0,
	0xa8, 0x6c, 0x98, 0x67, 0x1d, 0xe1, 0xda, 0x7c, 0xab, 0x78, 0x1e, 0x5a, 0xcf, 0xdd, 0x03, 0xf2,
	0xb6, 0xc8, 0x7a, 0x64, 0xb3, 0x1e, 0x15, 0x59, 0x67, 0xbe, 0x45, 0xb8, 0x7f, 0x1d, 0x68, 0x67,
	0xdc, 0x4e, 0x79, 0xc6, 0x7b, 0xca, 0xb3, 0x9b, 0x2b, 0xb1, 0xaf, 0xa8, 0xe7, 0x9a, 0xd7, 0xf8,
	0x69, 0xa9, 0xc6, 0x95, 0x12, 0xbc, 0xac, 0x95, 0x20, 0x8d, 0x6c, 0x0b, 0xbe, 0x6f, 0x56, 0x48,
	0x4f, 0xec, 0xcf, 0x57, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0xf3, 0x11, 0x62, 0xc4,
	0x04, 0x00, 0x00,
}