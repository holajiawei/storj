// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: overlay.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}
var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}
func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{11, 0}
}

type Restriction_Operand int32

const (
	Restriction_FREE_BANDWIDTH Restriction_Operand = 0
	Restriction_FREE_DISK      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "FREE_BANDWIDTH",
	1: "FREE_DISK",
}
var Restriction_Operand_value = map[string]int32{
	"FREE_BANDWIDTH": 0,
	"FREE_DISK":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}
func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{11, 1}
}

// LookupRequest is is request message for the lookup rpc call
type LookupRequest struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (dst *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(dst, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

// LookupResponse is is response message for the lookup rpc call
type LookupResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{1}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (dst *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(dst, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// LookupRequests is a list of LookupRequest
type LookupRequests struct {
	LookupRequest        []*LookupRequest `protobuf:"bytes,1,rep,name=lookup_request,json=lookupRequest" json:"lookup_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LookupRequests) Reset()         { *m = LookupRequests{} }
func (m *LookupRequests) String() string { return proto.CompactTextString(m) }
func (*LookupRequests) ProtoMessage()    {}
func (*LookupRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{2}
}
func (m *LookupRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequests.Unmarshal(m, b)
}
func (m *LookupRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequests.Marshal(b, m, deterministic)
}
func (dst *LookupRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequests.Merge(dst, src)
}
func (m *LookupRequests) XXX_Size() int {
	return xxx_messageInfo_LookupRequests.Size(m)
}
func (m *LookupRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequests.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequests proto.InternalMessageInfo

func (m *LookupRequests) GetLookupRequest() []*LookupRequest {
	if m != nil {
		return m.LookupRequest
	}
	return nil
}

// LookupResponse is a list of LookupResponse
type LookupResponses struct {
	LookupResponse       []*LookupResponse `protobuf:"bytes,1,rep,name=lookup_response,json=lookupResponse" json:"lookup_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LookupResponses) Reset()         { *m = LookupResponses{} }
func (m *LookupResponses) String() string { return proto.CompactTextString(m) }
func (*LookupResponses) ProtoMessage()    {}
func (*LookupResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{3}
}
func (m *LookupResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponses.Unmarshal(m, b)
}
func (m *LookupResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponses.Marshal(b, m, deterministic)
}
func (dst *LookupResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponses.Merge(dst, src)
}
func (m *LookupResponses) XXX_Size() int {
	return xxx_messageInfo_LookupResponses.Size(m)
}
func (m *LookupResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponses.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponses proto.InternalMessageInfo

func (m *LookupResponses) GetLookupResponse() []*LookupResponse {
	if m != nil {
		return m.LookupResponse
	}
	return nil
}

// FindStorageNodesResponse is is response message for the FindStorageNodes rpc call
type FindStorageNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStorageNodesResponse) Reset()         { *m = FindStorageNodesResponse{} }
func (m *FindStorageNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesResponse) ProtoMessage()    {}
func (*FindStorageNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{4}
}
func (m *FindStorageNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesResponse.Unmarshal(m, b)
}
func (m *FindStorageNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesResponse.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesResponse.Merge(dst, src)
}
func (m *FindStorageNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesResponse.Size(m)
}
func (m *FindStorageNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesResponse proto.InternalMessageInfo

func (m *FindStorageNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// FindStorageNodesRequest is is request message for the FindStorageNodes rpc call
type FindStorageNodesRequest struct {
	ObjectSize           int64              `protobuf:"varint,1,opt,name=object_size,json=objectSize,proto3" json:"object_size,omitempty"`
	ContractLength       *duration.Duration `protobuf:"bytes,2,opt,name=contract_length,json=contractLength" json:"contract_length,omitempty"`
	Opts                 *OverlayOptions    `protobuf:"bytes,3,opt,name=opts" json:"opts,omitempty"`
	Start                NodeID             `protobuf:"bytes,4,opt,name=start,proto3,customtype=NodeID" json:"start"`
	MaxNodes             int64              `protobuf:"varint,5,opt,name=max_nodes,json=maxNodes,proto3" json:"max_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindStorageNodesRequest) Reset()         { *m = FindStorageNodesRequest{} }
func (m *FindStorageNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesRequest) ProtoMessage()    {}
func (*FindStorageNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{5}
}
func (m *FindStorageNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesRequest.Unmarshal(m, b)
}
func (m *FindStorageNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesRequest.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesRequest.Merge(dst, src)
}
func (m *FindStorageNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesRequest.Size(m)
}
func (m *FindStorageNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesRequest proto.InternalMessageInfo

func (m *FindStorageNodesRequest) GetObjectSize() int64 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *FindStorageNodesRequest) GetContractLength() *duration.Duration {
	if m != nil {
		return m.ContractLength
	}
	return nil
}

func (m *FindStorageNodesRequest) GetOpts() *OverlayOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

func (m *FindStorageNodesRequest) GetMaxNodes() int64 {
	if m != nil {
		return m.MaxNodes
	}
	return 0
}

// OverlayOptions is a set of criteria that a node must meet to be considered for a storage opportunity
type OverlayOptions struct {
	MaxLatency           *duration.Duration `protobuf:"bytes,1,opt,name=max_latency,json=maxLatency" json:"max_latency,omitempty"`
	MinStats             *NodeStats         `protobuf:"bytes,2,opt,name=min_stats,json=minStats" json:"min_stats,omitempty"`
	MinSpeedKbps         int64              `protobuf:"varint,3,opt,name=min_speed_kbps,json=minSpeedKbps,proto3" json:"min_speed_kbps,omitempty"`
	Amount               int64              `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Restrictions         *NodeRestrictions  `protobuf:"bytes,5,opt,name=restrictions" json:"restrictions,omitempty"`
	ExcludedNodes        []NodeID           `protobuf:"bytes,6,rep,name=excluded_nodes,json=excludedNodes,customtype=NodeID" json:"excluded_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverlayOptions) Reset()         { *m = OverlayOptions{} }
func (m *OverlayOptions) String() string { return proto.CompactTextString(m) }
func (*OverlayOptions) ProtoMessage()    {}
func (*OverlayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{6}
}
func (m *OverlayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverlayOptions.Unmarshal(m, b)
}
func (m *OverlayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverlayOptions.Marshal(b, m, deterministic)
}
func (dst *OverlayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverlayOptions.Merge(dst, src)
}
func (m *OverlayOptions) XXX_Size() int {
	return xxx_messageInfo_OverlayOptions.Size(m)
}
func (m *OverlayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_OverlayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_OverlayOptions proto.InternalMessageInfo

func (m *OverlayOptions) GetMaxLatency() *duration.Duration {
	if m != nil {
		return m.MaxLatency
	}
	return nil
}

func (m *OverlayOptions) GetMinStats() *NodeStats {
	if m != nil {
		return m.MinStats
	}
	return nil
}

func (m *OverlayOptions) GetMinSpeedKbps() int64 {
	if m != nil {
		return m.MinSpeedKbps
	}
	return 0
}

func (m *OverlayOptions) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OverlayOptions) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

type QueryRequest struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Target               *Node    `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Pingback             bool     `protobuf:"varint,4,opt,name=pingback,proto3" json:"pingback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{7}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (dst *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(dst, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetPingback() bool {
	if m != nil {
		return m.Pingback
	}
	return false
}

type QueryResponse struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Response             []*Node  `protobuf:"bytes,2,rep,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{8}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (dst *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(dst, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryResponse) GetResponse() []*Node {
	if m != nil {
		return m.Response
	}
	return nil
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{9}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{10}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

type Restriction struct {
	Operator             Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand              Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value                int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Restriction) Reset()         { *m = Restriction{} }
func (m *Restriction) String() string { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()    {}
func (*Restriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_99a51b72c26b0776, []int{11}
}
func (m *Restriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restriction.Unmarshal(m, b)
}
func (m *Restriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restriction.Marshal(b, m, deterministic)
}
func (dst *Restriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restriction.Merge(dst, src)
}
func (m *Restriction) XXX_Size() int {
	return xxx_messageInfo_Restriction.Size(m)
}
func (m *Restriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Restriction.DiscardUnknown(m)
}

var xxx_messageInfo_Restriction proto.InternalMessageInfo

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_FREE_BANDWIDTH
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "overlay.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "overlay.LookupResponse")
	proto.RegisterType((*LookupRequests)(nil), "overlay.LookupRequests")
	proto.RegisterType((*LookupResponses)(nil), "overlay.LookupResponses")
	proto.RegisterType((*FindStorageNodesResponse)(nil), "overlay.FindStorageNodesResponse")
	proto.RegisterType((*FindStorageNodesRequest)(nil), "overlay.FindStorageNodesRequest")
	proto.RegisterType((*OverlayOptions)(nil), "overlay.OverlayOptions")
	proto.RegisterType((*QueryRequest)(nil), "overlay.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "overlay.QueryResponse")
	proto.RegisterType((*PingRequest)(nil), "overlay.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "overlay.PingResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Overlay service

type OverlayClient interface {
	// Lookup finds a nodes address from the network
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error)
}

type overlayClient struct {
	cc *grpc.ClientConn
}

func NewOverlayClient(cc *grpc.ClientConn) OverlayClient {
	return &overlayClient{cc}
}

func (c *overlayClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) BulkLookup(ctx context.Context, in *LookupRequests, opts ...grpc.CallOption) (*LookupResponses, error) {
	out := new(LookupResponses)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/BulkLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error) {
	out := new(FindStorageNodesResponse)
	err := c.cc.Invoke(ctx, "/overlay.Overlay/FindStorageNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Overlay service

type OverlayServer interface {
	// Lookup finds a nodes address from the network
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	// BulkLookup finds nodes addresses from the network
	BulkLookup(context.Context, *LookupRequests) (*LookupResponses, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(context.Context, *FindStorageNodesRequest) (*FindStorageNodesResponse, error)
}

func RegisterOverlayServer(s *grpc.Server, srv OverlayServer) {
	s.RegisterService(&_Overlay_serviceDesc, srv)
}

func _Overlay_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_BulkLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).BulkLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/BulkLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).BulkLookup(ctx, req.(*LookupRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_FindStorageNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).FindStorageNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Overlay/FindStorageNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).FindStorageNodes(ctx, req.(*FindStorageNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Overlay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Overlay",
	HandlerType: (*OverlayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Overlay_Lookup_Handler,
		},
		{
			MethodName: "BulkLookup",
			Handler:    _Overlay_BulkLookup_Handler,
		},
		{
			MethodName: "FindStorageNodes",
			Handler:    _Overlay_FindStorageNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

// Client API for Nodes service

type NodesClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nodes service

type NodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Nodes_Query_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Nodes_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_overlay_99a51b72c26b0776) }

var fileDescriptor_overlay_99a51b72c26b0776 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0xe7, 0xc7, 0xc9, 0x9e, 0x24, 0x5e, 0x6b, 0xd4, 0xee, 0x06, 0x03, 0xdd, 0x60, 0x55,
	0xb0, 0x12, 0x55, 0x0a, 0x29, 0xaa, 0x68, 0x05, 0x02, 0xa2, 0xa4, 0x65, 0xd5, 0xa8, 0x4b, 0x27,
	0x91, 0x2a, 0xc1, 0x85, 0xe5, 0xc4, 0x83, 0x31, 0xeb, 0x78, 0x8c, 0x67, 0x5c, 0xed, 0xf6, 0x09,
	0x78, 0x13, 0x5e, 0x85, 0x67, 0xe0, 0x62, 0x1f, 0x81, 0x07, 0xe0, 0x0a, 0xcd, 0x8f, 0x1d, 0x67,
	0x77, 0x53, 0xf5, 0x6a, 0xe6, 0x9c, 0xf3, 0x7d, 0x67, 0xe6, 0x3b, 0x73, 0xe6, 0x40, 0x8f, 0xbe,
	0x21, 0x59, 0xec, 0x5f, 0x0e, 0xd3, 0x8c, 0x72, 0x8a, 0x5a, 0xda, 0x74, 0xee, 0x85, 0x94, 0x86,
	0x31, 0x79, 0x28, 0xdd, 0xcb, 0xfc, 0xd7, 0x87, 0x41, 0x9e, 0xf9, 0x3c, 0xa2, 0x89, 0x02, 0x3a,
	0x10, 0xd2, 0x90, 0x16, 0xfb, 0x84, 0x06, 0x44, 0xed, 0xdd, 0xaf, 0xa1, 0x37, 0xa3, 0xf4, 0x3c,
	0x4f, 0x31, 0xf9, 0x23, 0x27, 0x8c, 0xa3, 0xcf, 0xa0, 0x25, 0xc2, 0x5e, 0x14, 0xf4, 0x8d, 0x81,
	0x71, 0xd2, 0x1d, 0x5b, 0x7f, 0x5f, 0x1d, 0xef, 0xfd, 0x73, 0x75, 0x6c, 0xbe, 0xa4, 0x01, 0x39,
	0x9d, 0x60, 0x53, 0x84, 0x4f, 0x03, 0xf7, 0x0b, 0xb0, 0x0a, 0x26, 0x4b, 0x69, 0xc2, 0x08, 0xba,
	0x07, 0x0d, 0x11, 0x93, 0xbc, 0xce, 0x08, 0x86, 0xf2, 0x18, 0xc1, 0xc2, 0xd2, 0xef, 0x9e, 0x6d,
	0x18, 0xf2, 0x2c, 0x86, 0xbe, 0x05, 0x2b, 0x96, 0x1e, 0x2f, 0x53, 0xae, 0xbe, 0x31, 0xa8, 0x9f,
	0x74, 0x46, 0x87, 0xc3, 0x42, 0xe6, 0x16, 0x01, 0xf7, 0xe2, 0xaa, 0xe9, 0xce, 0xe1, 0x60, 0xfb,
	0x0a, 0x0c, 0x7d, 0x0f, 0x07, 0x65, 0x46, 0xe5, 0xd3, 0x29, 0x8f, 0x6e, 0xa4, 0x54, 0x61, 0x6c,
	0xc5, 0x5b, 0xb6, 0xfb, 0x0d, 0xf4, 0x9f, 0x45, 0x49, 0x30, 0xe7, 0x34, 0xf3, 0x43, 0x22, 0xae,
	0xcf, 0x4a, 0x85, 0x03, 0x68, 0x0a, 0x25, 0x4c, 0xe7, 0xac, 0x4a, 0x54, 0x01, 0xf7, 0x5f, 0x03,
	0x8e, 0x6e, 0xd2, 0x55, 0x69, 0x8f, 0xa1, 0x43, 0x97, 0xbf, 0x93, 0x15, 0xf7, 0x58, 0xf4, 0x56,
	0x95, 0xa9, 0x8e, 0x41, 0xb9, 0xe6, 0xd1, 0x5b, 0x82, 0xc6, 0x70, 0xb0, 0xa2, 0x09, 0xcf, 0xfc,
	0x15, 0xf7, 0x62, 0x92, 0x84, 0xfc, 0xb7, 0x7e, 0x4d, 0xd6, 0xf2, 0x83, 0xa1, 0x7a, 0xde, 0x61,
	0xf1, 0xbc, 0xc3, 0x89, 0x7e, 0x5e, 0x6c, 0x15, 0x8c, 0x99, 0x24, 0xa0, 0xcf, 0xa1, 0x41, 0x53,
	0xce, 0xfa, 0x75, 0x49, 0xdc, 0xa8, 0x3e, 0x53, 0xeb, 0x59, 0x2a, 0x58, 0x0c, 0x4b, 0x10, 0xba,
	0x0f, 0x4d, 0xc6, 0xfd, 0x8c, 0xf7, 0x1b, 0xb7, 0x3e, 0xb5, 0x0a, 0xa2, 0x0f, 0x61, 0x7f, 0xed,
	0x5f, 0x78, 0x4a, 0x79, 0x53, 0xde, 0xba, 0xbd, 0xf6, 0x2f, 0xa4, 0x36, 0xf7, 0xaf, 0x1a, 0x58,
	0xdb, 0xb9, 0xd1, 0x53, 0xe8, 0x08, 0x7c, 0xec, 0x73, 0x92, 0xac, 0x2e, 0x75, 0x3b, 0xbc, 0x43,
	0x02, 0xac, 0xfd, 0x8b, 0x99, 0x02, 0xa3, 0x07, 0xb0, 0xbf, 0x8e, 0x12, 0x8f, 0x71, 0x9f, 0x33,
	0x2d, 0xfe, 0x60, 0x53, 0xe5, 0xb9, 0x70, 0xe3, 0xf6, 0x3a, 0x4a, 0xe4, 0x0e, 0xdd, 0x07, 0x4b,
	0xa2, 0x53, 0x42, 0x02, 0xef, 0x7c, 0x99, 0x2a, 0xd9, 0x75, 0xdc, 0x15, 0x08, 0xe1, 0x7c, 0xb1,
	0x4c, 0x19, 0x3a, 0x04, 0xd3, 0x5f, 0xd3, 0x3c, 0x51, 0x32, 0xeb, 0x58, 0x5b, 0xe8, 0x29, 0x74,
	0x33, 0xc2, 0x78, 0x16, 0xad, 0xe4, 0xbd, 0xa5, 0x34, 0xd1, 0x7b, 0x9b, 0x47, 0xad, 0x44, 0xf1,
	0x16, 0x16, 0x7d, 0x09, 0x16, 0xb9, 0x58, 0xc5, 0x79, 0x40, 0x02, 0x5d, 0x18, 0x73, 0x50, 0x3f,
	0xe9, 0x8e, 0xa1, 0x52, 0xbe, 0x5e, 0x81, 0x50, 0x95, 0xfa, 0xd3, 0x80, 0xee, 0xab, 0x9c, 0x64,
	0x97, 0x45, 0x3f, 0xb8, 0x60, 0x32, 0x92, 0x04, 0x24, 0xbb, 0xe5, 0xc7, 0xe8, 0x88, 0xc0, 0x70,
	0x3f, 0x0b, 0x09, 0xd7, 0xc5, 0xd8, 0xc2, 0xa8, 0x08, 0xba, 0x03, 0xcd, 0x38, 0x5a, 0x47, 0x5c,
	0x8b, 0x57, 0x06, 0x72, 0xa0, 0x9d, 0x46, 0x49, 0xb8, 0xf4, 0x57, 0xe7, 0x52, 0x77, 0x1b, 0x97,
	0xb6, 0xfb, 0x0b, 0xf4, 0xf4, 0x4d, 0x74, 0x63, 0xbf, 0xcf, 0x55, 0x3e, 0x85, 0x76, 0xf9, 0xa7,
	0x6a, 0x37, 0xfa, 0xbf, 0x8c, 0xb9, 0x3d, 0xe8, 0xfc, 0x14, 0x25, 0x61, 0xf1, 0x49, 0x2d, 0xe8,
	0x2a, 0x53, 0x87, 0xff, 0x33, 0xa0, 0x53, 0x29, 0x2c, 0x7a, 0x02, 0x6d, 0x9a, 0x92, 0xcc, 0xe7,
	0x54, 0x1d, 0x6e, 0x8d, 0x3e, 0x2e, 0x9b, 0xb6, 0x82, 0x1b, 0x9e, 0x69, 0x10, 0x2e, 0xe1, 0xe8,
	0x31, 0xb4, 0xe4, 0x3e, 0x09, 0x64, 0x75, 0xac, 0xd1, 0x47, 0xbb, 0x99, 0x49, 0x80, 0x0b, 0xb0,
	0x28, 0xd8, 0x1b, 0x3f, 0xce, 0x49, 0x51, 0x30, 0x69, 0xb8, 0x5f, 0x41, 0xbb, 0x38, 0x03, 0x99,
	0x50, 0x9b, 0x2d, 0xec, 0x3d, 0xb1, 0x4e, 0x5f, 0xd9, 0x86, 0x58, 0x9f, 0x2f, 0xec, 0x1a, 0x6a,
	0x41, 0x7d, 0xb6, 0x98, 0xda, 0x75, 0xb1, 0x79, 0xbe, 0x98, 0xda, 0x0d, 0xf7, 0x01, 0xb4, 0x74,
	0x7e, 0x84, 0xc0, 0x7a, 0x86, 0xa7, 0x53, 0x6f, 0xfc, 0xc3, 0xcb, 0xc9, 0xeb, 0xd3, 0xc9, 0xe2,
	0x47, 0x7b, 0x0f, 0xf5, 0x60, 0x5f, 0xfa, 0x26, 0xa7, 0xf3, 0x17, 0xb6, 0x31, 0xba, 0x32, 0xa0,
	0xa5, 0x7f, 0x0b, 0x7a, 0x02, 0xa6, 0x1a, 0x45, 0x68, 0xc7, 0xb8, 0x73, 0x76, 0xcd, 0x2c, 0xf4,
	0x1d, 0xc0, 0x38, 0x8f, 0xcf, 0x35, 0xfd, 0xe8, 0x76, 0x3a, 0x73, 0xfa, 0x3b, 0xf8, 0x0c, 0xbd,
	0x06, 0xfb, 0xfa, 0x94, 0x42, 0x83, 0x12, 0xbd, 0x63, 0x80, 0x39, 0x9f, 0xbc, 0x03, 0xa1, 0x32,
	0x8f, 0x38, 0x34, 0x55, 0xb6, 0xc7, 0xd0, 0x94, 0x2d, 0x86, 0xee, 0x96, 0xa4, 0x6a, 0xf3, 0x3b,
	0x87, 0xd7, 0xdd, 0x5a, 0xda, 0x23, 0x68, 0x88, 0x76, 0x41, 0x77, 0xca, 0x78, 0xa5, 0x99, 0x9c,
	0xbb, 0xd7, 0xbc, 0x8a, 0x34, 0x6e, 0xfc, 0x5c, 0x4b, 0x97, 0x4b, 0x53, 0x8e, 0x96, 0x47, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xae, 0x57, 0x27, 0x55, 0x24, 0x07, 0x00, 0x00,
}
