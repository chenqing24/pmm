// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/object_details.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MetricsRequest defines filtering of metrics for specific value of dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
type MetricsRequest struct {
	PeriodStartFrom *timestamp.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `protobuf:"bytes,3,opt,name=filter_by,json=filterBy,proto3" json:"filter_by,omitempty"`
	// one of dimension: queryid | host ...
	GroupBy           string           `protobuf:"bytes,4,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Labels            []*MapFieldEntry `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	IncludeOnlyFields []string         `protobuf:"bytes,6,rep,name=include_only_fields,json=includeOnlyFields,proto3" json:"include_only_fields,omitempty"`
	// retrieve only values for totals, excluding N/A values
	Totals               bool     `protobuf:"varint,7,opt,name=totals,proto3" json:"totals,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{0}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

func (m *MetricsRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *MetricsRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

func (m *MetricsRequest) GetFilterBy() string {
	if m != nil {
		return m.FilterBy
	}
	return ""
}

func (m *MetricsRequest) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

func (m *MetricsRequest) GetLabels() []*MapFieldEntry {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MetricsRequest) GetIncludeOnlyFields() []string {
	if m != nil {
		return m.IncludeOnlyFields
	}
	return nil
}

func (m *MetricsRequest) GetTotals() bool {
	if m != nil {
		return m.Totals
	}
	return false
}

// MetricsReply defines metrics for specific value of dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
type MetricsReply struct {
	Metrics              map[string]*MetricValues `protobuf:"bytes,3,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Sparkline            []*Point                 `protobuf:"bytes,4,rep,name=sparkline,proto3" json:"sparkline,omitempty"`
	Totals               map[string]*MetricValues `protobuf:"bytes,5,rep,name=totals,proto3" json:"totals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Fingerprint          string                   `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MetricsReply) Reset()         { *m = MetricsReply{} }
func (m *MetricsReply) String() string { return proto.CompactTextString(m) }
func (*MetricsReply) ProtoMessage()    {}
func (*MetricsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{1}
}

func (m *MetricsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsReply.Unmarshal(m, b)
}
func (m *MetricsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsReply.Marshal(b, m, deterministic)
}
func (m *MetricsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsReply.Merge(m, src)
}
func (m *MetricsReply) XXX_Size() int {
	return xxx_messageInfo_MetricsReply.Size(m)
}
func (m *MetricsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsReply.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsReply proto.InternalMessageInfo

func (m *MetricsReply) GetMetrics() map[string]*MetricValues {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *MetricsReply) GetSparkline() []*Point {
	if m != nil {
		return m.Sparkline
	}
	return nil
}

func (m *MetricsReply) GetTotals() map[string]*MetricValues {
	if m != nil {
		return m.Totals
	}
	return nil
}

func (m *MetricsReply) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

// MetricValues is statistics of specific metric.
type MetricValues struct {
	Rate                 float32  `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Cnt                  float32  `protobuf:"fixed32,2,opt,name=cnt,proto3" json:"cnt,omitempty"`
	Sum                  float32  `protobuf:"fixed32,3,opt,name=sum,proto3" json:"sum,omitempty"`
	Min                  float32  `protobuf:"fixed32,4,opt,name=min,proto3" json:"min,omitempty"`
	Max                  float32  `protobuf:"fixed32,5,opt,name=max,proto3" json:"max,omitempty"`
	Avg                  float32  `protobuf:"fixed32,6,opt,name=avg,proto3" json:"avg,omitempty"`
	P99                  float32  `protobuf:"fixed32,7,opt,name=p99,proto3" json:"p99,omitempty"`
	PercentOfTotal       float32  `protobuf:"fixed32,8,opt,name=percent_of_total,json=percentOfTotal,proto3" json:"percent_of_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricValues) Reset()         { *m = MetricValues{} }
func (m *MetricValues) String() string { return proto.CompactTextString(m) }
func (*MetricValues) ProtoMessage()    {}
func (*MetricValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{2}
}

func (m *MetricValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValues.Unmarshal(m, b)
}
func (m *MetricValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValues.Marshal(b, m, deterministic)
}
func (m *MetricValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValues.Merge(m, src)
}
func (m *MetricValues) XXX_Size() int {
	return xxx_messageInfo_MetricValues.Size(m)
}
func (m *MetricValues) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValues.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValues proto.InternalMessageInfo

func (m *MetricValues) GetRate() float32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *MetricValues) GetCnt() float32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func (m *MetricValues) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *MetricValues) GetMin() float32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *MetricValues) GetMax() float32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MetricValues) GetAvg() float32 {
	if m != nil {
		return m.Avg
	}
	return 0
}

func (m *MetricValues) GetP99() float32 {
	if m != nil {
		return m.P99
	}
	return 0
}

func (m *MetricValues) GetPercentOfTotal() float32 {
	if m != nil {
		return m.PercentOfTotal
	}
	return 0
}

// Labels are list of labels or dimensions values.
type Labels struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Labels) Reset()         { *m = Labels{} }
func (m *Labels) String() string { return proto.CompactTextString(m) }
func (*Labels) ProtoMessage()    {}
func (*Labels) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{3}
}

func (m *Labels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Labels.Unmarshal(m, b)
}
func (m *Labels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Labels.Marshal(b, m, deterministic)
}
func (m *Labels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Labels.Merge(m, src)
}
func (m *Labels) XXX_Size() int {
	return xxx_messageInfo_Labels.Size(m)
}
func (m *Labels) XXX_DiscardUnknown() {
	xxx_messageInfo_Labels.DiscardUnknown(m)
}

var xxx_messageInfo_Labels proto.InternalMessageInfo

func (m *Labels) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// QueryExampleRequest defines filtering of query examples for specific value of
// dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
type QueryExampleRequest struct {
	PeriodStartFrom *timestamp.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `protobuf:"bytes,3,opt,name=filter_by,json=filterBy,proto3" json:"filter_by,omitempty"`
	// one of dimension: queryid | host ...
	GroupBy              string           `protobuf:"bytes,4,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Labels               []*MapFieldEntry `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	Limit                uint32           `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *QueryExampleRequest) Reset()         { *m = QueryExampleRequest{} }
func (m *QueryExampleRequest) String() string { return proto.CompactTextString(m) }
func (*QueryExampleRequest) ProtoMessage()    {}
func (*QueryExampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{4}
}

func (m *QueryExampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryExampleRequest.Unmarshal(m, b)
}
func (m *QueryExampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryExampleRequest.Marshal(b, m, deterministic)
}
func (m *QueryExampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExampleRequest.Merge(m, src)
}
func (m *QueryExampleRequest) XXX_Size() int {
	return xxx_messageInfo_QueryExampleRequest.Size(m)
}
func (m *QueryExampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExampleRequest proto.InternalMessageInfo

func (m *QueryExampleRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *QueryExampleRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

func (m *QueryExampleRequest) GetFilterBy() string {
	if m != nil {
		return m.FilterBy
	}
	return ""
}

func (m *QueryExampleRequest) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

func (m *QueryExampleRequest) GetLabels() []*MapFieldEntry {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *QueryExampleRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// QueryExampleReply list of query examples.
type QueryExampleReply struct {
	QueryExamples        []*QueryExample `protobuf:"bytes,1,rep,name=query_examples,json=queryExamples,proto3" json:"query_examples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *QueryExampleReply) Reset()         { *m = QueryExampleReply{} }
func (m *QueryExampleReply) String() string { return proto.CompactTextString(m) }
func (*QueryExampleReply) ProtoMessage()    {}
func (*QueryExampleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{5}
}

func (m *QueryExampleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryExampleReply.Unmarshal(m, b)
}
func (m *QueryExampleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryExampleReply.Marshal(b, m, deterministic)
}
func (m *QueryExampleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExampleReply.Merge(m, src)
}
func (m *QueryExampleReply) XXX_Size() int {
	return xxx_messageInfo_QueryExampleReply.Size(m)
}
func (m *QueryExampleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExampleReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExampleReply proto.InternalMessageInfo

func (m *QueryExampleReply) GetQueryExamples() []*QueryExample {
	if m != nil {
		return m.QueryExamples
	}
	return nil
}

// QueryExample shows query examples and their metrics.
type QueryExample struct {
	Example              string        `protobuf:"bytes,1,opt,name=example,proto3" json:"example,omitempty"`
	ExampleFormat        ExampleFormat `protobuf:"varint,2,opt,name=example_format,json=exampleFormat,proto3,enum=qan.v1beta1.ExampleFormat" json:"example_format,omitempty"` // Deprecated: Do not use.
	ExampleType          ExampleType   `protobuf:"varint,3,opt,name=example_type,json=exampleType,proto3,enum=qan.v1beta1.ExampleType" json:"example_type,omitempty"`
	IsTruncated          uint32        `protobuf:"varint,4,opt,name=is_truncated,json=isTruncated,proto3" json:"is_truncated,omitempty"`
	ExampleMetrics       string        `protobuf:"bytes,5,opt,name=example_metrics,json=exampleMetrics,proto3" json:"example_metrics,omitempty"`
	ServiceId            string        `protobuf:"bytes,6,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceType          string        `protobuf:"bytes,7,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	Schema               string        `protobuf:"bytes,8,opt,name=schema,proto3" json:"schema,omitempty"`
	Tables               []string      `protobuf:"bytes,9,rep,name=tables,proto3" json:"tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueryExample) Reset()         { *m = QueryExample{} }
func (m *QueryExample) String() string { return proto.CompactTextString(m) }
func (*QueryExample) ProtoMessage()    {}
func (*QueryExample) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{6}
}

func (m *QueryExample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryExample.Unmarshal(m, b)
}
func (m *QueryExample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryExample.Marshal(b, m, deterministic)
}
func (m *QueryExample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryExample.Merge(m, src)
}
func (m *QueryExample) XXX_Size() int {
	return xxx_messageInfo_QueryExample.Size(m)
}
func (m *QueryExample) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryExample.DiscardUnknown(m)
}

var xxx_messageInfo_QueryExample proto.InternalMessageInfo

func (m *QueryExample) GetExample() string {
	if m != nil {
		return m.Example
	}
	return ""
}

// Deprecated: Do not use.
func (m *QueryExample) GetExampleFormat() ExampleFormat {
	if m != nil {
		return m.ExampleFormat
	}
	return ExampleFormat_EXAMPLE_FORMAT_INVALID
}

func (m *QueryExample) GetExampleType() ExampleType {
	if m != nil {
		return m.ExampleType
	}
	return ExampleType_EXAMPLE_TYPE_INVALID
}

func (m *QueryExample) GetIsTruncated() uint32 {
	if m != nil {
		return m.IsTruncated
	}
	return 0
}

func (m *QueryExample) GetExampleMetrics() string {
	if m != nil {
		return m.ExampleMetrics
	}
	return ""
}

func (m *QueryExample) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *QueryExample) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *QueryExample) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *QueryExample) GetTables() []string {
	if m != nil {
		return m.Tables
	}
	return nil
}

// ObjectDetailsLabelsRequest defines filtering of object detail's labels for specific value of
// dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
type ObjectDetailsLabelsRequest struct {
	PeriodStartFrom *timestamp.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	// dimension value: ex: queryid - 1D410B4BE5060972.
	FilterBy string `protobuf:"bytes,3,opt,name=filter_by,json=filterBy,proto3" json:"filter_by,omitempty"`
	// one of dimension: queryid | host ...
	GroupBy              string   `protobuf:"bytes,4,opt,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectDetailsLabelsRequest) Reset()         { *m = ObjectDetailsLabelsRequest{} }
func (m *ObjectDetailsLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectDetailsLabelsRequest) ProtoMessage()    {}
func (*ObjectDetailsLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{7}
}

func (m *ObjectDetailsLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDetailsLabelsRequest.Unmarshal(m, b)
}
func (m *ObjectDetailsLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDetailsLabelsRequest.Marshal(b, m, deterministic)
}
func (m *ObjectDetailsLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDetailsLabelsRequest.Merge(m, src)
}
func (m *ObjectDetailsLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectDetailsLabelsRequest.Size(m)
}
func (m *ObjectDetailsLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDetailsLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDetailsLabelsRequest proto.InternalMessageInfo

func (m *ObjectDetailsLabelsRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *ObjectDetailsLabelsRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

func (m *ObjectDetailsLabelsRequest) GetFilterBy() string {
	if m != nil {
		return m.FilterBy
	}
	return ""
}

func (m *ObjectDetailsLabelsRequest) GetGroupBy() string {
	if m != nil {
		return m.GroupBy
	}
	return ""
}

// ObjectDetailsLabelsReply is a map of labels names as keys and labels values as a list.
type ObjectDetailsLabelsReply struct {
	Labels               map[string]*ListLabelValues `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ObjectDetailsLabelsReply) Reset()         { *m = ObjectDetailsLabelsReply{} }
func (m *ObjectDetailsLabelsReply) String() string { return proto.CompactTextString(m) }
func (*ObjectDetailsLabelsReply) ProtoMessage()    {}
func (*ObjectDetailsLabelsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{8}
}

func (m *ObjectDetailsLabelsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDetailsLabelsReply.Unmarshal(m, b)
}
func (m *ObjectDetailsLabelsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDetailsLabelsReply.Marshal(b, m, deterministic)
}
func (m *ObjectDetailsLabelsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDetailsLabelsReply.Merge(m, src)
}
func (m *ObjectDetailsLabelsReply) XXX_Size() int {
	return xxx_messageInfo_ObjectDetailsLabelsReply.Size(m)
}
func (m *ObjectDetailsLabelsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDetailsLabelsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDetailsLabelsReply proto.InternalMessageInfo

func (m *ObjectDetailsLabelsReply) GetLabels() map[string]*ListLabelValues {
	if m != nil {
		return m.Labels
	}
	return nil
}

// ListLabelValues is list of label's values.
type ListLabelValues struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLabelValues) Reset()         { *m = ListLabelValues{} }
func (m *ListLabelValues) String() string { return proto.CompactTextString(m) }
func (*ListLabelValues) ProtoMessage()    {}
func (*ListLabelValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_173f4868b35b4ff7, []int{9}
}

func (m *ListLabelValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabelValues.Unmarshal(m, b)
}
func (m *ListLabelValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabelValues.Marshal(b, m, deterministic)
}
func (m *ListLabelValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabelValues.Merge(m, src)
}
func (m *ListLabelValues) XXX_Size() int {
	return xxx_messageInfo_ListLabelValues.Size(m)
}
func (m *ListLabelValues) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabelValues.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabelValues proto.InternalMessageInfo

func (m *ListLabelValues) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsRequest)(nil), "qan.v1beta1.MetricsRequest")
	proto.RegisterType((*MetricsReply)(nil), "qan.v1beta1.MetricsReply")
	proto.RegisterMapType((map[string]*MetricValues)(nil), "qan.v1beta1.MetricsReply.MetricsEntry")
	proto.RegisterMapType((map[string]*MetricValues)(nil), "qan.v1beta1.MetricsReply.TotalsEntry")
	proto.RegisterType((*MetricValues)(nil), "qan.v1beta1.MetricValues")
	proto.RegisterType((*Labels)(nil), "qan.v1beta1.Labels")
	proto.RegisterType((*QueryExampleRequest)(nil), "qan.v1beta1.QueryExampleRequest")
	proto.RegisterType((*QueryExampleReply)(nil), "qan.v1beta1.QueryExampleReply")
	proto.RegisterType((*QueryExample)(nil), "qan.v1beta1.QueryExample")
	proto.RegisterType((*ObjectDetailsLabelsRequest)(nil), "qan.v1beta1.ObjectDetailsLabelsRequest")
	proto.RegisterType((*ObjectDetailsLabelsReply)(nil), "qan.v1beta1.ObjectDetailsLabelsReply")
	proto.RegisterMapType((map[string]*ListLabelValues)(nil), "qan.v1beta1.ObjectDetailsLabelsReply.LabelsEntry")
	proto.RegisterType((*ListLabelValues)(nil), "qan.v1beta1.ListLabelValues")
}

func init() {
	proto.RegisterFile("qanpb/object_details.proto", fileDescriptor_173f4868b35b4ff7)
}

var fileDescriptor_173f4868b35b4ff7 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x56, 0xcf, 0x6e, 0x1b, 0x45,
	0x18, 0xd7, 0xae, 0x6b, 0x3b, 0x1e, 0xc7, 0x76, 0x33, 0xad, 0xaa, 0x8d, 0x5b, 0x5a, 0xb3, 0x22,
	0x34, 0x04, 0xe2, 0x4d, 0x82, 0x84, 0x68, 0x10, 0x12, 0x8a, 0x48, 0xa2, 0x4a, 0xad, 0x02, 0x8b,
	0x0b, 0x12, 0x97, 0xd5, 0xd8, 0x1e, 0xbb, 0x43, 0x76, 0x77, 0xd6, 0x3b, 0x63, 0x27, 0x7b, 0x45,
	0x42, 0x5c, 0x91, 0xe0, 0x0d, 0x78, 0x09, 0x4e, 0x3c, 0x03, 0x12, 0x2f, 0xc0, 0x81, 0x13, 0x4f,
	0x81, 0xe6, 0x9b, 0xd9, 0x64, 0x37, 0xd8, 0x94, 0x03, 0x27, 0x38, 0xf9, 0xfb, 0xff, 0xfd, 0xe6,
	0x9b, 0xdf, 0x7c, 0x6b, 0xd4, 0x9d, 0x91, 0x38, 0x19, 0x7a, 0x7c, 0xf8, 0x15, 0x1d, 0xc9, 0x60,
	0x4c, 0x25, 0x61, 0xa1, 0xe8, 0x27, 0x29, 0x97, 0x1c, 0x37, 0x67, 0x24, 0xee, 0x2f, 0xf6, 0x87,
	0x54, 0x92, 0xfd, 0xee, 0x7b, 0x53, 0x26, 0x5f, 0xce, 0x87, 0xfd, 0x11, 0x8f, 0xbc, 0xe8, 0x82,
	0xc9, 0x73, 0x7e, 0xe1, 0x4d, 0xf9, 0x2e, 0x44, 0xee, 0x2e, 0x48, 0xc8, 0xc6, 0x44, 0xf2, 0x54,
	0x78, 0x57, 0xa2, 0x2e, 0xd2, 0x7d, 0x30, 0xe5, 0x7c, 0x1a, 0x52, 0x8f, 0x24, 0xcc, 0x23, 0x71,
	0xcc, 0x25, 0x91, 0x8c, 0xc7, 0xa6, 0x45, 0xf7, 0x91, 0xf1, 0x82, 0x36, 0x9c, 0x4f, 0x3c, 0xc9,
	0x22, 0x2a, 0x24, 0x89, 0x12, 0x13, 0xf0, 0x0e, 0xfc, 0x8c, 0x76, 0xa7, 0x34, 0xde, 0x15, 0x17,
	0x64, 0x3a, 0xa5, 0xa9, 0xc7, 0x13, 0x28, 0xb1, 0xa4, 0x5c, 0x47, 0x9f, 0x46, 0xe1, 0x06, 0x83,
	0xfb, 0x8b, 0x8d, 0xda, 0xcf, 0xa9, 0x4c, 0xd9, 0x48, 0xf8, 0x74, 0x36, 0xa7, 0x42, 0xe2, 0x13,
	0xb4, 0x91, 0xd0, 0x94, 0xf1, 0x71, 0x20, 0x24, 0x49, 0x65, 0x30, 0x49, 0x79, 0xe4, 0x58, 0x3d,
	0x6b, 0xbb, 0x79, 0xd0, 0xed, 0x6b, 0x38, 0xfd, 0x1c, 0x4e, 0x7f, 0x90, 0xc3, 0xf1, 0x3b, 0x3a,
	0xe9, 0x33, 0x95, 0x73, 0x92, 0xf2, 0x08, 0x1f, 0xa1, 0x4e, 0xa9, 0x8e, 0xe4, 0x8e, 0xfd, 0xca,
	0x2a, 0xad, 0x42, 0x95, 0x01, 0xc7, 0xf7, 0x51, 0x63, 0xc2, 0x42, 0x49, 0xd3, 0x60, 0x98, 0x39,
	0x95, 0x9e, 0xb5, 0xdd, 0xf0, 0xd7, 0xb4, 0xe1, 0x28, 0xc3, 0x9b, 0x68, 0x6d, 0x9a, 0xf2, 0x79,
	0xa2, 0x7c, 0xb7, 0xc0, 0x57, 0x07, 0xfd, 0x28, 0xc3, 0x07, 0xa8, 0x16, 0x92, 0x21, 0x0d, 0x85,
	0x53, 0xed, 0x55, 0xa0, 0x65, 0xe1, 0xaa, 0xfa, 0xcf, 0x49, 0x72, 0xc2, 0x68, 0x38, 0x3e, 0x8e,
	0x65, 0x9a, 0xf9, 0x26, 0x12, 0xf7, 0xd1, 0x1d, 0x16, 0x8f, 0xc2, 0xf9, 0x98, 0x06, 0x3c, 0x0e,
	0xb3, 0x60, 0xa2, 0x42, 0x84, 0x53, 0xeb, 0x55, 0xb6, 0x1b, 0xfe, 0x86, 0x71, 0x9d, 0xc5, 0x61,
	0x06, 0xb9, 0x02, 0xdf, 0x43, 0x35, 0xc9, 0x25, 0x09, 0x85, 0x53, 0xef, 0x59, 0xdb, 0x6b, 0xbe,
	0xd1, 0xdc, 0xef, 0x2a, 0x68, 0xfd, 0x6a, 0xa4, 0x49, 0x98, 0xe1, 0x8f, 0x50, 0x3d, 0xd2, 0xba,
	0x53, 0x01, 0x34, 0x6f, 0x96, 0xd1, 0x14, 0x62, 0x73, 0x45, 0x23, 0xcb, 0xd3, 0xf0, 0x1e, 0x6a,
	0x88, 0x84, 0xa4, 0xe7, 0x21, 0x8b, 0xa9, 0x73, 0x0b, 0x6a, 0xe0, 0x52, 0x8d, 0x4f, 0x38, 0x8b,
	0xa5, 0x7f, 0x1d, 0x84, 0x3f, 0xbc, 0x02, 0xa7, 0x07, 0xb0, 0xb5, 0xba, 0xe5, 0x00, 0xe2, 0xcc,
	0x2c, 0x74, 0x12, 0xee, 0xa1, 0xe6, 0x84, 0xc5, 0x53, 0x9a, 0x26, 0x29, 0x8b, 0xa5, 0x53, 0x83,
	0xe9, 0x16, 0x4d, 0xdd, 0x17, 0x57, 0x87, 0x84, 0x4c, 0x7c, 0x1b, 0x55, 0xce, 0x69, 0x06, 0x3c,
	0x69, 0xf8, 0x4a, 0xc4, 0x1e, 0xaa, 0x2e, 0x48, 0x38, 0xa7, 0xe6, 0xd6, 0x37, 0x97, 0x20, 0xf8,
	0x5c, 0xf9, 0x85, 0xaf, 0xe3, 0x0e, 0xed, 0xf7, 0xad, 0xee, 0x00, 0x35, 0x0b, 0x78, 0xfe, 0xa5,
	0xaa, 0xee, 0x4f, 0x56, 0x8e, 0x56, 0xfb, 0x30, 0x46, 0xb7, 0x52, 0x22, 0x29, 0x14, 0xb6, 0x7d,
	0x90, 0x55, 0xaf, 0x51, 0x2c, 0xa1, 0xae, 0xed, 0x2b, 0x51, 0x59, 0xc4, 0x3c, 0x02, 0xde, 0xd9,
	0xbe, 0x12, 0x95, 0x25, 0x62, 0x31, 0xb0, 0xcd, 0xf6, 0x95, 0x08, 0x16, 0x72, 0xe9, 0x54, 0x8d,
	0x85, 0x5c, 0x2a, 0x0b, 0x59, 0x4c, 0x61, 0x66, 0xb6, 0xaf, 0x44, 0x65, 0x49, 0x9e, 0x3c, 0x01,
	0x9a, 0xd8, 0xbe, 0x12, 0xf1, 0x36, 0xba, 0x9d, 0xd0, 0x74, 0x44, 0x63, 0x19, 0xf0, 0x49, 0x00,
	0x43, 0x77, 0xd6, 0xc0, 0xdd, 0x36, 0xf6, 0xb3, 0x09, 0xcc, 0xc1, 0x7d, 0x88, 0x6a, 0xcf, 0x34,
	0x3f, 0xef, 0xe6, 0x27, 0xb7, 0x80, 0x91, 0x5a, 0x71, 0x7f, 0xb4, 0xd1, 0x9d, 0x4f, 0xe7, 0x34,
	0xcd, 0x8e, 0x2f, 0x49, 0x94, 0x84, 0xf4, 0xff, 0xfe, 0x8a, 0xef, 0xa2, 0x6a, 0xc8, 0x22, 0xa6,
	0x39, 0xdb, 0xf2, 0xb5, 0xe2, 0xbe, 0x40, 0x1b, 0xe5, 0x21, 0xe9, 0x77, 0xd9, 0x9e, 0x29, 0x63,
	0x40, 0xb5, 0x55, 0xc0, 0x64, 0x6f, 0x72, 0xaa, 0x94, 0xd7, 0x9a, 0x15, 0x34, 0xe1, 0xfe, 0x61,
	0xa3, 0xf5, 0xa2, 0x1f, 0x3b, 0xa8, 0x6e, 0x8a, 0x19, 0xce, 0xe6, 0x2a, 0x3e, 0x46, 0x6d, 0x23,
	0x06, 0x13, 0x9e, 0x46, 0x44, 0x13, 0xad, 0x7d, 0xe3, 0x4c, 0xa6, 0xce, 0x09, 0x44, 0x1c, 0xd9,
	0x8e, 0xe5, 0xb7, 0x68, 0xd1, 0x84, 0x3f, 0x40, 0xeb, 0x79, 0x19, 0x99, 0x25, 0x14, 0xa6, 0xd9,
	0x3e, 0x70, 0x96, 0x15, 0x19, 0x64, 0x09, 0xf5, 0x9b, 0xf4, 0x5a, 0xc1, 0xaf, 0xa3, 0x75, 0x26,
	0x02, 0x99, 0xce, 0xe3, 0x11, 0x91, 0x74, 0x0c, 0xe3, 0x6e, 0xf9, 0x4d, 0x26, 0x06, 0xb9, 0x09,
	0x3f, 0x46, 0x9d, 0xbc, 0x7e, 0xbe, 0xb3, 0xaa, 0x70, 0x90, 0x1c, 0xbd, 0x79, 0xf4, 0xf8, 0x35,
	0x84, 0x04, 0x4d, 0x17, 0x6c, 0x44, 0x03, 0x36, 0x36, 0x0b, 0xa2, 0x61, 0x2c, 0x4f, 0xc7, 0xaa,
	0x55, 0xee, 0x06, 0x9c, 0x75, 0xbd, 0x41, 0x8c, 0x0d, 0xd0, 0xdc, 0x43, 0x35, 0x31, 0x7a, 0x49,
	0x23, 0x02, 0xcc, 0x6f, 0xf8, 0x46, 0x83, 0xbd, 0x4a, 0x86, 0xea, 0x3a, 0x1a, 0x40, 0x74, 0xa3,
	0xb9, 0xbf, 0x59, 0xa8, 0x7b, 0x06, 0x9f, 0xe1, 0x8f, 0xf5, 0x57, 0x58, 0xbf, 0x8b, 0xff, 0x10,
	0xe1, 0xdd, 0x9f, 0x2d, 0xe4, 0x2c, 0x3d, 0xa2, 0xa2, 0xeb, 0xd3, 0xab, 0xd7, 0xa0, 0x69, 0xba,
	0x5f, 0xba, 0xf4, 0x55, 0x69, 0x7d, 0x2d, 0x97, 0x1e, 0x49, 0xf7, 0x0b, 0xd4, 0x2c, 0x98, 0x97,
	0x6c, 0xd9, 0x83, 0xf2, 0x96, 0x7d, 0x50, 0x6a, 0xf5, 0x8c, 0x09, 0x09, 0xe9, 0x7f, 0x5d, 0xb4,
	0x6f, 0xa1, 0xce, 0x0d, 0xaf, 0xba, 0x4e, 0xf0, 0x0b, 0xb3, 0xb7, 0x8c, 0x76, 0xf0, 0x43, 0x05,
	0xb5, 0x4a, 0xa0, 0xf1, 0x0c, 0xa1, 0x53, 0x2a, 0x73, 0x82, 0xdd, 0x5f, 0xfe, 0xc5, 0x82, 0xcb,
	0xee, 0x6e, 0xae, 0xfc, 0x9c, 0xb9, 0x6f, 0x7f, 0xfd, 0xeb, 0xef, 0xdf, 0xdb, 0x5b, 0x6e, 0xcf,
	0x5b, 0xec, 0xa9, 0x3f, 0x3a, 0x5e, 0xa9, 0x89, 0x77, 0xdd, 0xe1, 0xd0, 0xda, 0xc1, 0xdf, 0x58,
	0xa8, 0x73, 0x4a, 0x65, 0xe9, 0x0d, 0xf7, 0x56, 0x3f, 0x7f, 0xd3, 0xfd, 0xe1, 0xdf, 0x44, 0x28,
	0x08, 0x7b, 0x00, 0x61, 0xc7, 0xdd, 0x5a, 0x09, 0xa1, 0x98, 0xa3, 0x70, 0x7c, 0x6b, 0xa1, 0xc6,
	0x29, 0x95, 0x66, 0xd3, 0x3f, 0x7e, 0xf5, 0xcd, 0x6a, 0x20, 0x5b, 0xff, 0x88, 0x02, 0xee, 0x0e,
	0xe0, 0x79, 0xc3, 0x7d, 0xb4, 0x12, 0x8f, 0x8e, 0x3e, 0xb4, 0x76, 0x8e, 0xea, 0x5f, 0x56, 0xe1,
	0x3f, 0xe2, 0xb0, 0x06, 0x34, 0x7f, 0xf7, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x12, 0x49,
	0x34, 0x01, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ObjectDetailsClient is the client API for ObjectDetails service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObjectDetailsClient interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsReply, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(ctx context.Context, in *QueryExampleRequest, opts ...grpc.CallOption) (*QueryExampleReply, error)
	// GetLabels gets list of labels for object details.
	GetLabels(ctx context.Context, in *ObjectDetailsLabelsRequest, opts ...grpc.CallOption) (*ObjectDetailsLabelsReply, error)
}

type objectDetailsClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectDetailsClient(cc grpc.ClientConnInterface) ObjectDetailsClient {
	return &objectDetailsClient{cc}
}

func (c *objectDetailsClient) GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsReply, error) {
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetQueryExample(ctx context.Context, in *QueryExampleRequest, opts ...grpc.CallOption) (*QueryExampleReply, error) {
	out := new(QueryExampleReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetQueryExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetLabels(ctx context.Context, in *ObjectDetailsLabelsRequest, opts ...grpc.CallOption) (*ObjectDetailsLabelsReply, error) {
	out := new(ObjectDetailsLabelsReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectDetailsServer is the server API for ObjectDetails service.
type ObjectDetailsServer interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(context.Context, *MetricsRequest) (*MetricsReply, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(context.Context, *QueryExampleRequest) (*QueryExampleReply, error)
	// GetLabels gets list of labels for object details.
	GetLabels(context.Context, *ObjectDetailsLabelsRequest) (*ObjectDetailsLabelsReply, error)
}

// UnimplementedObjectDetailsServer can be embedded to have forward compatible implementations.
type UnimplementedObjectDetailsServer struct {
}

func (*UnimplementedObjectDetailsServer) GetMetrics(ctx context.Context, req *MetricsRequest) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (*UnimplementedObjectDetailsServer) GetQueryExample(ctx context.Context, req *QueryExampleRequest) (*QueryExampleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryExample not implemented")
}
func (*UnimplementedObjectDetailsServer) GetLabels(ctx context.Context, req *ObjectDetailsLabelsRequest) (*ObjectDetailsLabelsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}

func RegisterObjectDetailsServer(s *grpc.Server, srv ObjectDetailsServer) {
	s.RegisterService(&_ObjectDetails_serviceDesc, srv)
}

func _ObjectDetails_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetMetrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetQueryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetQueryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetQueryExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetQueryExample(ctx, req.(*QueryExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectDetailsLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetLabels(ctx, req.(*ObjectDetailsLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectDetails_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.ObjectDetails",
	HandlerType: (*ObjectDetailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _ObjectDetails_GetMetrics_Handler,
		},
		{
			MethodName: "GetQueryExample",
			Handler:    _ObjectDetails_GetQueryExample_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _ObjectDetails_GetLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/object_details.proto",
}
