// Code generated by protoc-gen-go. DO NOT EDIT.
// source: visit.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type VisitRegisterRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsMobile             bool     `protobuf:"varint,3,opt,name=is_mobile,json=isMobile,proto3" json:"is_mobile,omitempty"`
	Browser              string   `protobuf:"bytes,4,opt,name=browser,proto3" json:"browser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitRegisterRequest) Reset()         { *m = VisitRegisterRequest{} }
func (m *VisitRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*VisitRegisterRequest) ProtoMessage()    {}
func (*VisitRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{0}
}

func (m *VisitRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitRegisterRequest.Unmarshal(m, b)
}
func (m *VisitRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitRegisterRequest.Marshal(b, m, deterministic)
}
func (m *VisitRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitRegisterRequest.Merge(m, src)
}
func (m *VisitRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_VisitRegisterRequest.Size(m)
}
func (m *VisitRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VisitRegisterRequest proto.InternalMessageInfo

func (m *VisitRegisterRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitRegisterRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VisitRegisterRequest) GetIsMobile() bool {
	if m != nil {
		return m.IsMobile
	}
	return false
}

func (m *VisitRegisterRequest) GetBrowser() string {
	if m != nil {
		return m.Browser
	}
	return ""
}

type VisitReportByUrlResponse struct {
	Today                *VisitReportByUrlResponse_Report `protobuf:"bytes,1,opt,name=today,proto3" json:"today,omitempty"`
	LastDay              *VisitReportByUrlResponse_Report `protobuf:"bytes,2,opt,name=last_day,json=lastDay,proto3" json:"last_day,omitempty"`
	LastWeek             *VisitReportByUrlResponse_Report `protobuf:"bytes,3,opt,name=last_week,json=lastWeek,proto3" json:"last_week,omitempty"`
	LastMonth            *VisitReportByUrlResponse_Report `protobuf:"bytes,4,opt,name=last_month,json=lastMonth,proto3" json:"last_month,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *VisitReportByUrlResponse) Reset()         { *m = VisitReportByUrlResponse{} }
func (m *VisitReportByUrlResponse) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUrlResponse) ProtoMessage()    {}
func (*VisitReportByUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{1}
}

func (m *VisitReportByUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUrlResponse.Unmarshal(m, b)
}
func (m *VisitReportByUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUrlResponse.Marshal(b, m, deterministic)
}
func (m *VisitReportByUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUrlResponse.Merge(m, src)
}
func (m *VisitReportByUrlResponse) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUrlResponse.Size(m)
}
func (m *VisitReportByUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUrlResponse proto.InternalMessageInfo

func (m *VisitReportByUrlResponse) GetToday() *VisitReportByUrlResponse_Report {
	if m != nil {
		return m.Today
	}
	return nil
}

func (m *VisitReportByUrlResponse) GetLastDay() *VisitReportByUrlResponse_Report {
	if m != nil {
		return m.LastDay
	}
	return nil
}

func (m *VisitReportByUrlResponse) GetLastWeek() *VisitReportByUrlResponse_Report {
	if m != nil {
		return m.LastWeek
	}
	return nil
}

func (m *VisitReportByUrlResponse) GetLastMonth() *VisitReportByUrlResponse_Report {
	if m != nil {
		return m.LastMonth
	}
	return nil
}

type VisitReportByUrlResponse_Total struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUrlResponse_Total) Reset()         { *m = VisitReportByUrlResponse_Total{} }
func (m *VisitReportByUrlResponse_Total) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUrlResponse_Total) ProtoMessage()    {}
func (*VisitReportByUrlResponse_Total) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{1, 0}
}

func (m *VisitReportByUrlResponse_Total) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUrlResponse_Total.Unmarshal(m, b)
}
func (m *VisitReportByUrlResponse_Total) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUrlResponse_Total.Marshal(b, m, deterministic)
}
func (m *VisitReportByUrlResponse_Total) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUrlResponse_Total.Merge(m, src)
}
func (m *VisitReportByUrlResponse_Total) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUrlResponse_Total.Size(m)
}
func (m *VisitReportByUrlResponse_Total) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUrlResponse_Total.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUrlResponse_Total proto.InternalMessageInfo

func (m *VisitReportByUrlResponse_Total) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUrlResponse_Total) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUrlResponse_ByMobile struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	IsMobile             bool     `protobuf:"varint,2,opt,name=is_mobile,json=isMobile,proto3" json:"is_mobile,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUrlResponse_ByMobile) Reset()         { *m = VisitReportByUrlResponse_ByMobile{} }
func (m *VisitReportByUrlResponse_ByMobile) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUrlResponse_ByMobile) ProtoMessage()    {}
func (*VisitReportByUrlResponse_ByMobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{1, 1}
}

func (m *VisitReportByUrlResponse_ByMobile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUrlResponse_ByMobile.Unmarshal(m, b)
}
func (m *VisitReportByUrlResponse_ByMobile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUrlResponse_ByMobile.Marshal(b, m, deterministic)
}
func (m *VisitReportByUrlResponse_ByMobile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUrlResponse_ByMobile.Merge(m, src)
}
func (m *VisitReportByUrlResponse_ByMobile) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUrlResponse_ByMobile.Size(m)
}
func (m *VisitReportByUrlResponse_ByMobile) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUrlResponse_ByMobile.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUrlResponse_ByMobile proto.InternalMessageInfo

func (m *VisitReportByUrlResponse_ByMobile) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUrlResponse_ByMobile) GetIsMobile() bool {
	if m != nil {
		return m.IsMobile
	}
	return false
}

func (m *VisitReportByUrlResponse_ByMobile) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUrlResponse_ByBrowser struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Browser              string   `protobuf:"bytes,2,opt,name=browser,proto3" json:"browser,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUrlResponse_ByBrowser) Reset()         { *m = VisitReportByUrlResponse_ByBrowser{} }
func (m *VisitReportByUrlResponse_ByBrowser) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUrlResponse_ByBrowser) ProtoMessage()    {}
func (*VisitReportByUrlResponse_ByBrowser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{1, 2}
}

func (m *VisitReportByUrlResponse_ByBrowser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUrlResponse_ByBrowser.Unmarshal(m, b)
}
func (m *VisitReportByUrlResponse_ByBrowser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUrlResponse_ByBrowser.Marshal(b, m, deterministic)
}
func (m *VisitReportByUrlResponse_ByBrowser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUrlResponse_ByBrowser.Merge(m, src)
}
func (m *VisitReportByUrlResponse_ByBrowser) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUrlResponse_ByBrowser.Size(m)
}
func (m *VisitReportByUrlResponse_ByBrowser) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUrlResponse_ByBrowser.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUrlResponse_ByBrowser proto.InternalMessageInfo

func (m *VisitReportByUrlResponse_ByBrowser) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUrlResponse_ByBrowser) GetBrowser() string {
	if m != nil {
		return m.Browser
	}
	return ""
}

func (m *VisitReportByUrlResponse_ByBrowser) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUrlResponse_Report struct {
	Total                []*VisitReportByUrlResponse_Total     `protobuf:"bytes,1,rep,name=total,proto3" json:"total,omitempty"`
	ByMobile             []*VisitReportByUrlResponse_ByMobile  `protobuf:"bytes,2,rep,name=by_mobile,json=byMobile,proto3" json:"by_mobile,omitempty"`
	ByBrowser            []*VisitReportByUrlResponse_ByBrowser `protobuf:"bytes,3,rep,name=by_browser,json=byBrowser,proto3" json:"by_browser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *VisitReportByUrlResponse_Report) Reset()         { *m = VisitReportByUrlResponse_Report{} }
func (m *VisitReportByUrlResponse_Report) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUrlResponse_Report) ProtoMessage()    {}
func (*VisitReportByUrlResponse_Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{1, 3}
}

func (m *VisitReportByUrlResponse_Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUrlResponse_Report.Unmarshal(m, b)
}
func (m *VisitReportByUrlResponse_Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUrlResponse_Report.Marshal(b, m, deterministic)
}
func (m *VisitReportByUrlResponse_Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUrlResponse_Report.Merge(m, src)
}
func (m *VisitReportByUrlResponse_Report) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUrlResponse_Report.Size(m)
}
func (m *VisitReportByUrlResponse_Report) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUrlResponse_Report.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUrlResponse_Report proto.InternalMessageInfo

func (m *VisitReportByUrlResponse_Report) GetTotal() []*VisitReportByUrlResponse_Total {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *VisitReportByUrlResponse_Report) GetByMobile() []*VisitReportByUrlResponse_ByMobile {
	if m != nil {
		return m.ByMobile
	}
	return nil
}

func (m *VisitReportByUrlResponse_Report) GetByBrowser() []*VisitReportByUrlResponse_ByBrowser {
	if m != nil {
		return m.ByBrowser
	}
	return nil
}

type VisitReportByUserResponse struct {
	Today                *VisitReportByUserResponse_Report `protobuf:"bytes,1,opt,name=today,proto3" json:"today,omitempty"`
	LastDay              *VisitReportByUserResponse_Report `protobuf:"bytes,2,opt,name=last_day,json=lastDay,proto3" json:"last_day,omitempty"`
	LastWeek             *VisitReportByUserResponse_Report `protobuf:"bytes,3,opt,name=last_week,json=lastWeek,proto3" json:"last_week,omitempty"`
	LastMonth            *VisitReportByUserResponse_Report `protobuf:"bytes,4,opt,name=last_month,json=lastMonth,proto3" json:"last_month,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *VisitReportByUserResponse) Reset()         { *m = VisitReportByUserResponse{} }
func (m *VisitReportByUserResponse) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUserResponse) ProtoMessage()    {}
func (*VisitReportByUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{2}
}

func (m *VisitReportByUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUserResponse.Unmarshal(m, b)
}
func (m *VisitReportByUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUserResponse.Marshal(b, m, deterministic)
}
func (m *VisitReportByUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUserResponse.Merge(m, src)
}
func (m *VisitReportByUserResponse) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUserResponse.Size(m)
}
func (m *VisitReportByUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUserResponse proto.InternalMessageInfo

func (m *VisitReportByUserResponse) GetToday() *VisitReportByUserResponse_Report {
	if m != nil {
		return m.Today
	}
	return nil
}

func (m *VisitReportByUserResponse) GetLastDay() *VisitReportByUserResponse_Report {
	if m != nil {
		return m.LastDay
	}
	return nil
}

func (m *VisitReportByUserResponse) GetLastWeek() *VisitReportByUserResponse_Report {
	if m != nil {
		return m.LastWeek
	}
	return nil
}

func (m *VisitReportByUserResponse) GetLastMonth() *VisitReportByUserResponse_Report {
	if m != nil {
		return m.LastMonth
	}
	return nil
}

type VisitReportByUserResponse_Total struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUserResponse_Total) Reset()         { *m = VisitReportByUserResponse_Total{} }
func (m *VisitReportByUserResponse_Total) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUserResponse_Total) ProtoMessage()    {}
func (*VisitReportByUserResponse_Total) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{2, 0}
}

func (m *VisitReportByUserResponse_Total) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUserResponse_Total.Unmarshal(m, b)
}
func (m *VisitReportByUserResponse_Total) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUserResponse_Total.Marshal(b, m, deterministic)
}
func (m *VisitReportByUserResponse_Total) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUserResponse_Total.Merge(m, src)
}
func (m *VisitReportByUserResponse_Total) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUserResponse_Total.Size(m)
}
func (m *VisitReportByUserResponse_Total) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUserResponse_Total.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUserResponse_Total proto.InternalMessageInfo

func (m *VisitReportByUserResponse_Total) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUserResponse_Total) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VisitReportByUserResponse_Total) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUserResponse_ByMobile struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsMobile             bool     `protobuf:"varint,3,opt,name=is_mobile,json=isMobile,proto3" json:"is_mobile,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUserResponse_ByMobile) Reset()         { *m = VisitReportByUserResponse_ByMobile{} }
func (m *VisitReportByUserResponse_ByMobile) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUserResponse_ByMobile) ProtoMessage()    {}
func (*VisitReportByUserResponse_ByMobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{2, 1}
}

func (m *VisitReportByUserResponse_ByMobile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUserResponse_ByMobile.Unmarshal(m, b)
}
func (m *VisitReportByUserResponse_ByMobile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUserResponse_ByMobile.Marshal(b, m, deterministic)
}
func (m *VisitReportByUserResponse_ByMobile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUserResponse_ByMobile.Merge(m, src)
}
func (m *VisitReportByUserResponse_ByMobile) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUserResponse_ByMobile.Size(m)
}
func (m *VisitReportByUserResponse_ByMobile) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUserResponse_ByMobile.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUserResponse_ByMobile proto.InternalMessageInfo

func (m *VisitReportByUserResponse_ByMobile) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUserResponse_ByMobile) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VisitReportByUserResponse_ByMobile) GetIsMobile() bool {
	if m != nil {
		return m.IsMobile
	}
	return false
}

func (m *VisitReportByUserResponse_ByMobile) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUserResponse_ByBrowser struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Browser              string   `protobuf:"bytes,3,opt,name=browser,proto3" json:"browser,omitempty"`
	Count                int64    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VisitReportByUserResponse_ByBrowser) Reset()         { *m = VisitReportByUserResponse_ByBrowser{} }
func (m *VisitReportByUserResponse_ByBrowser) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUserResponse_ByBrowser) ProtoMessage()    {}
func (*VisitReportByUserResponse_ByBrowser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{2, 2}
}

func (m *VisitReportByUserResponse_ByBrowser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUserResponse_ByBrowser.Unmarshal(m, b)
}
func (m *VisitReportByUserResponse_ByBrowser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUserResponse_ByBrowser.Marshal(b, m, deterministic)
}
func (m *VisitReportByUserResponse_ByBrowser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUserResponse_ByBrowser.Merge(m, src)
}
func (m *VisitReportByUserResponse_ByBrowser) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUserResponse_ByBrowser.Size(m)
}
func (m *VisitReportByUserResponse_ByBrowser) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUserResponse_ByBrowser.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUserResponse_ByBrowser proto.InternalMessageInfo

func (m *VisitReportByUserResponse_ByBrowser) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *VisitReportByUserResponse_ByBrowser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VisitReportByUserResponse_ByBrowser) GetBrowser() string {
	if m != nil {
		return m.Browser
	}
	return ""
}

func (m *VisitReportByUserResponse_ByBrowser) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VisitReportByUserResponse_Report struct {
	Total                []*VisitReportByUserResponse_Total     `protobuf:"bytes,1,rep,name=total,proto3" json:"total,omitempty"`
	ByMobile             []*VisitReportByUserResponse_ByMobile  `protobuf:"bytes,2,rep,name=by_mobile,json=byMobile,proto3" json:"by_mobile,omitempty"`
	ByBrowser            []*VisitReportByUserResponse_ByBrowser `protobuf:"bytes,3,rep,name=by_browser,json=byBrowser,proto3" json:"by_browser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *VisitReportByUserResponse_Report) Reset()         { *m = VisitReportByUserResponse_Report{} }
func (m *VisitReportByUserResponse_Report) String() string { return proto.CompactTextString(m) }
func (*VisitReportByUserResponse_Report) ProtoMessage()    {}
func (*VisitReportByUserResponse_Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_a498f0e5194d943b, []int{2, 3}
}

func (m *VisitReportByUserResponse_Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VisitReportByUserResponse_Report.Unmarshal(m, b)
}
func (m *VisitReportByUserResponse_Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VisitReportByUserResponse_Report.Marshal(b, m, deterministic)
}
func (m *VisitReportByUserResponse_Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VisitReportByUserResponse_Report.Merge(m, src)
}
func (m *VisitReportByUserResponse_Report) XXX_Size() int {
	return xxx_messageInfo_VisitReportByUserResponse_Report.Size(m)
}
func (m *VisitReportByUserResponse_Report) XXX_DiscardUnknown() {
	xxx_messageInfo_VisitReportByUserResponse_Report.DiscardUnknown(m)
}

var xxx_messageInfo_VisitReportByUserResponse_Report proto.InternalMessageInfo

func (m *VisitReportByUserResponse_Report) GetTotal() []*VisitReportByUserResponse_Total {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *VisitReportByUserResponse_Report) GetByMobile() []*VisitReportByUserResponse_ByMobile {
	if m != nil {
		return m.ByMobile
	}
	return nil
}

func (m *VisitReportByUserResponse_Report) GetByBrowser() []*VisitReportByUserResponse_ByBrowser {
	if m != nil {
		return m.ByBrowser
	}
	return nil
}

func init() {
	proto.RegisterType((*VisitRegisterRequest)(nil), "api.VisitRegisterRequest")
	proto.RegisterType((*VisitReportByUrlResponse)(nil), "api.VisitReportByUrlResponse")
	proto.RegisterType((*VisitReportByUrlResponse_Total)(nil), "api.VisitReportByUrlResponse.Total")
	proto.RegisterType((*VisitReportByUrlResponse_ByMobile)(nil), "api.VisitReportByUrlResponse.ByMobile")
	proto.RegisterType((*VisitReportByUrlResponse_ByBrowser)(nil), "api.VisitReportByUrlResponse.ByBrowser")
	proto.RegisterType((*VisitReportByUrlResponse_Report)(nil), "api.VisitReportByUrlResponse.Report")
	proto.RegisterType((*VisitReportByUserResponse)(nil), "api.VisitReportByUserResponse")
	proto.RegisterType((*VisitReportByUserResponse_Total)(nil), "api.VisitReportByUserResponse.Total")
	proto.RegisterType((*VisitReportByUserResponse_ByMobile)(nil), "api.VisitReportByUserResponse.ByMobile")
	proto.RegisterType((*VisitReportByUserResponse_ByBrowser)(nil), "api.VisitReportByUserResponse.ByBrowser")
	proto.RegisterType((*VisitReportByUserResponse_Report)(nil), "api.VisitReportByUserResponse.Report")
}

func init() {
	proto.RegisterFile("visit.proto", fileDescriptor_a498f0e5194d943b)
}

var fileDescriptor_a498f0e5194d943b = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x9a, 0xa5, 0x4d, 0x4e, 0xb9, 0x18, 0xd6, 0x24, 0xb2, 0x22, 0xd0, 0x54, 0x06, 0xeb,
	0x55, 0x91, 0xca, 0x0d, 0x0c, 0x09, 0x46, 0x56, 0xf1, 0x73, 0x51, 0x21, 0x59, 0xfc, 0x5c, 0x56,
	0xc9, 0x6a, 0x98, 0xb5, 0xb4, 0x0e, 0xb1, 0xcb, 0xc8, 0x8b, 0xf0, 0x42, 0x7b, 0x01, 0x24, 0x5e,
	0x08, 0xd9, 0x69, 0x5a, 0xa7, 0xf9, 0x69, 0xd8, 0x5d, 0xec, 0x9c, 0xef, 0x3b, 0x3e, 0xe7, 0x7c,
	0x9f, 0x0d, 0xdd, 0x9f, 0x94, 0x53, 0x31, 0x8c, 0x62, 0x26, 0x18, 0x32, 0xfd, 0x88, 0xf6, 0x60,
	0x4e, 0xf9, 0x45, 0xba, 0xd1, 0xff, 0x05, 0x07, 0x5f, 0xe4, 0x7f, 0x4c, 0xbe, 0x53, 0x2e, 0x48,
	0x8c, 0xc9, 0x8f, 0x25, 0xe1, 0x02, 0xed, 0x83, 0xb9, 0x8c, 0x43, 0xd7, 0x38, 0x32, 0x06, 0x0e,
	0x96, 0x9f, 0xe8, 0x1e, 0x74, 0x96, 0x9c, 0xc4, 0x53, 0x3a, 0x73, 0x5b, 0x6a, 0xb7, 0x2d, 0x97,
	0x1f, 0x66, 0xe8, 0x3e, 0x38, 0x94, 0x4f, 0xe7, 0x2c, 0xa0, 0x21, 0x71, 0xcd, 0x23, 0x63, 0x60,
	0x63, 0x9b, 0xf2, 0x89, 0x5a, 0x23, 0x17, 0x3a, 0x41, 0xcc, 0xae, 0x39, 0x89, 0xdd, 0x3d, 0x85,
	0xca, 0x96, 0xfd, 0x1b, 0x0b, 0xdc, 0x55, 0xea, 0x88, 0xc5, 0xc2, 0x4b, 0x3e, 0xc7, 0x21, 0x26,
	0x3c, 0x62, 0x0b, 0x4e, 0xd0, 0x29, 0x58, 0x82, 0xcd, 0xfc, 0x44, 0x1d, 0xa0, 0x3b, 0x3a, 0x1e,
	0xfa, 0x11, 0x1d, 0x56, 0x45, 0x0f, 0xd3, 0x3d, 0x9c, 0x42, 0xd0, 0x6b, 0xb0, 0x43, 0x9f, 0x8b,
	0xa9, 0x84, 0xb7, 0xfe, 0x03, 0xde, 0x91, 0xa8, 0xb1, 0x9f, 0xa0, 0x37, 0xe0, 0x28, 0x82, 0x6b,
	0x42, 0xae, 0x54, 0x41, 0x4d, 0x19, 0x54, 0xde, 0xaf, 0x84, 0x5c, 0xa1, 0x73, 0x00, 0x45, 0x31,
	0x67, 0x0b, 0x71, 0xa9, 0x2a, 0x6f, 0xca, 0xa1, 0x52, 0x4f, 0x24, 0xac, 0xf7, 0x14, 0xac, 0x4f,
	0x4c, 0xf8, 0x61, 0xc9, 0x30, 0x0e, 0xc0, 0xba, 0x60, 0xcb, 0x85, 0x50, 0x05, 0x9a, 0x38, 0x5d,
	0xf4, 0x3e, 0x82, 0xed, 0x25, 0xab, 0xc6, 0x17, 0x31, 0xb9, 0x39, 0xb5, 0xb6, 0xe6, 0xb4, 0x26,
	0x34, 0x75, 0xc2, 0x09, 0x38, 0x5e, 0xe2, 0xa5, 0x03, 0x2b, 0x61, 0xd4, 0x86, 0xdb, 0xca, 0x0d,
	0xb7, 0x82, 0xee, 0x8f, 0x01, 0xed, 0xb4, 0x4c, 0xf4, 0x42, 0x0e, 0x58, 0xf8, 0x92, 0xce, 0x1c,
	0x74, 0x47, 0x8f, 0xea, 0x7b, 0xa3, 0xda, 0x80, 0x53, 0x04, 0x3a, 0x07, 0x27, 0x48, 0x36, 0x75,
	0x48, 0xf8, 0x93, 0x7a, 0x78, 0xd6, 0x14, 0x6c, 0x07, 0x59, 0x7b, 0xde, 0x02, 0x04, 0xc9, 0x34,
	0x3b, 0xbd, 0xa9, 0x58, 0x4e, 0x76, 0xb1, 0xac, 0x3a, 0x81, 0x9d, 0x20, 0xfb, 0xec, 0xff, 0x6e,
	0xc3, 0x61, 0x1e, 0x21, 0x03, 0x32, 0x19, 0xbf, 0xcc, 0xcb, 0xf8, 0x71, 0x49, 0x02, 0x2d, 0x7c,
	0x4b, 0xc7, 0x67, 0x05, 0x1d, 0x37, 0xc4, 0xaf, 0x85, 0xec, 0x15, 0x85, 0xdc, 0x90, 0x62, 0xa3,
	0xe4, 0x71, 0x89, 0x92, 0x1b, 0x92, 0x68, 0x52, 0x7e, 0x5f, 0x2d, 0xe5, 0xca, 0x7b, 0xa5, 0x5c,
	0x43, 0x97, 0xb5, 0x1a, 0xbf, 0xdd, 0x25, 0xb5, 0xce, 0xb4, 0xa7, 0x67, 0xfa, 0x56, 0x2f, 0xfe,
	0xca, 0x54, 0x9a, 0x2b, 0xcc, 0x0a, 0x57, 0xe4, 0xf2, 0xfc, 0xdd, 0xb8, 0xe2, 0x34, 0xef, 0x8a,
	0xe3, 0x1d, 0x7d, 0xce, 0xd9, 0x62, 0x5c, 0xb4, 0xc5, 0xc9, 0x0e, 0x7c, 0x89, 0x2f, 0xde, 0x95,
	0xf8, 0x62, 0xb0, 0x93, 0xa6, 0x68, 0x8c, 0xd1, 0x8d, 0x01, 0x96, 0x82, 0xa0, 0xe7, 0x60, 0x67,
	0xaf, 0x0b, 0x3a, 0xd4, 0xa9, 0x72, 0x2f, 0x4e, 0xef, 0x6e, 0xfa, 0x8b, 0xd1, 0xd9, 0xda, 0x3e,
	0xaf, 0xa0, 0xab, 0x19, 0x11, 0xed, 0x6b, 0x11, 0x29, 0xe6, 0x41, 0xad, 0x63, 0xd1, 0x19, 0xdc,
	0xd1, 0x0f, 0x5c, 0x42, 0xf0, 0xb0, 0xbe, 0xb4, 0xa0, 0xad, 0x5e, 0xc9, 0x67, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0xc0, 0x49, 0xbd, 0x45, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VisitClient is the client API for Visit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisitClient interface {
	Register(ctx context.Context, in *VisitRegisterRequest, opts ...grpc.CallOption) (*VoidResponse, error)
	ReportByUrl(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*VisitReportByUrlResponse, error)
	ReportByUser(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*VisitReportByUserResponse, error)
}

type visitClient struct {
	cc grpc.ClientConnInterface
}

func NewVisitClient(cc grpc.ClientConnInterface) VisitClient {
	return &visitClient{cc}
}

func (c *visitClient) Register(ctx context.Context, in *VisitRegisterRequest, opts ...grpc.CallOption) (*VoidResponse, error) {
	out := new(VoidResponse)
	err := c.cc.Invoke(ctx, "/api.Visit/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visitClient) ReportByUrl(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*VisitReportByUrlResponse, error) {
	out := new(VisitReportByUrlResponse)
	err := c.cc.Invoke(ctx, "/api.Visit/ReportByUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visitClient) ReportByUser(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*VisitReportByUserResponse, error) {
	out := new(VisitReportByUserResponse)
	err := c.cc.Invoke(ctx, "/api.Visit/ReportByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisitServer is the server API for Visit service.
type VisitServer interface {
	Register(context.Context, *VisitRegisterRequest) (*VoidResponse, error)
	ReportByUrl(context.Context, *VoidRequest) (*VisitReportByUrlResponse, error)
	ReportByUser(context.Context, *VoidRequest) (*VisitReportByUserResponse, error)
}

// UnimplementedVisitServer can be embedded to have forward compatible implementations.
type UnimplementedVisitServer struct {
}

func (*UnimplementedVisitServer) Register(ctx context.Context, req *VisitRegisterRequest) (*VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedVisitServer) ReportByUrl(ctx context.Context, req *VoidRequest) (*VisitReportByUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportByUrl not implemented")
}
func (*UnimplementedVisitServer) ReportByUser(ctx context.Context, req *VoidRequest) (*VisitReportByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportByUser not implemented")
}

func RegisterVisitServer(s *grpc.Server, srv VisitServer) {
	s.RegisterService(&_Visit_serviceDesc, srv)
}

func _Visit_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Visit/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServer).Register(ctx, req.(*VisitRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Visit_ReportByUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServer).ReportByUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Visit/ReportByUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServer).ReportByUrl(ctx, req.(*VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Visit_ReportByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServer).ReportByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Visit/ReportByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServer).ReportByUser(ctx, req.(*VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Visit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Visit",
	HandlerType: (*VisitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Visit_Register_Handler,
		},
		{
			MethodName: "ReportByUrl",
			Handler:    _Visit_ReportByUrl_Handler,
		},
		{
			MethodName: "ReportByUser",
			Handler:    _Visit_ReportByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "visit.proto",
}
