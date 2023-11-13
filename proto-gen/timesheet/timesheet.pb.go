// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: proto/timesheet/timesheet.proto

package timesheet

import (
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

type Timesheet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContractorId     string  `protobuf:"bytes,2,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
	ClientId         string  `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Date             string  `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	DayOfWeek        string  `protobuf:"bytes,5,opt,name=dayOfWeek,proto3" json:"dayOfWeek,omitempty"`
	Task             string  `protobuf:"bytes,6,opt,name=task,proto3" json:"task,omitempty"`
	WorkDetails      string  `protobuf:"bytes,7,opt,name=workDetails,proto3" json:"workDetails,omitempty"`
	TotalHours       float32 `protobuf:"fixed32,8,opt,name=totalHours,proto3" json:"totalHours,omitempty"`
	NonBillableHours float32 `protobuf:"fixed32,9,opt,name=nonBillableHours,proto3" json:"nonBillableHours,omitempty"`
	TotalWorkHours   float32 `protobuf:"fixed32,10,opt,name=totalWorkHours,proto3" json:"totalWorkHours,omitempty"`
	HourlyRate       float32 `protobuf:"fixed32,11,opt,name=hourlyRate,proto3" json:"hourlyRate,omitempty"`
	Pay              float32 `protobuf:"fixed32,12,opt,name=pay,proto3" json:"pay,omitempty"`
}

func (x *Timesheet) Reset() {
	*x = Timesheet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesheet_timesheet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timesheet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timesheet) ProtoMessage() {}

func (x *Timesheet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesheet_timesheet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timesheet.ProtoReflect.Descriptor instead.
func (*Timesheet) Descriptor() ([]byte, []int) {
	return file_proto_timesheet_timesheet_proto_rawDescGZIP(), []int{0}
}

func (x *Timesheet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Timesheet) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

func (x *Timesheet) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Timesheet) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Timesheet) GetDayOfWeek() string {
	if x != nil {
		return x.DayOfWeek
	}
	return ""
}

func (x *Timesheet) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *Timesheet) GetWorkDetails() string {
	if x != nil {
		return x.WorkDetails
	}
	return ""
}

func (x *Timesheet) GetTotalHours() float32 {
	if x != nil {
		return x.TotalHours
	}
	return 0
}

func (x *Timesheet) GetNonBillableHours() float32 {
	if x != nil {
		return x.NonBillableHours
	}
	return 0
}

func (x *Timesheet) GetTotalWorkHours() float32 {
	if x != nil {
		return x.TotalWorkHours
	}
	return 0
}

func (x *Timesheet) GetHourlyRate() float32 {
	if x != nil {
		return x.HourlyRate
	}
	return 0
}

func (x *Timesheet) GetPay() float32 {
	if x != nil {
		return x.Pay
	}
	return 0
}

type CreateTimesheetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractorId string  `protobuf:"bytes,1,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
	ClientId     string  `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Date         string  `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Task         string  `protobuf:"bytes,4,opt,name=task,proto3" json:"task,omitempty"`
	WorkDetails  string  `protobuf:"bytes,5,opt,name=workDetails,proto3" json:"workDetails,omitempty"`
	TotalHours   float32 `protobuf:"fixed32,6,opt,name=totalHours,proto3" json:"totalHours,omitempty"`
}

func (x *CreateTimesheetRequest) Reset() {
	*x = CreateTimesheetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesheet_timesheet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTimesheetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTimesheetRequest) ProtoMessage() {}

func (x *CreateTimesheetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesheet_timesheet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTimesheetRequest.ProtoReflect.Descriptor instead.
func (*CreateTimesheetRequest) Descriptor() ([]byte, []int) {
	return file_proto_timesheet_timesheet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTimesheetRequest) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

func (x *CreateTimesheetRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateTimesheetRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *CreateTimesheetRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *CreateTimesheetRequest) GetWorkDetails() string {
	if x != nil {
		return x.WorkDetails
	}
	return ""
}

func (x *CreateTimesheetRequest) GetTotalHours() float32 {
	if x != nil {
		return x.TotalHours
	}
	return 0
}

type GetTimesheetsByContractorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractorId string `protobuf:"bytes,1,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
}

func (x *GetTimesheetsByContractorRequest) Reset() {
	*x = GetTimesheetsByContractorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesheet_timesheet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimesheetsByContractorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimesheetsByContractorRequest) ProtoMessage() {}

func (x *GetTimesheetsByContractorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesheet_timesheet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimesheetsByContractorRequest.ProtoReflect.Descriptor instead.
func (*GetTimesheetsByContractorRequest) Descriptor() ([]byte, []int) {
	return file_proto_timesheet_timesheet_proto_rawDescGZIP(), []int{2}
}

func (x *GetTimesheetsByContractorRequest) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

type GetTimesheetsByContractorAndClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractorId string `protobuf:"bytes,1,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
	ClientId     string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *GetTimesheetsByContractorAndClientRequest) Reset() {
	*x = GetTimesheetsByContractorAndClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesheet_timesheet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimesheetsByContractorAndClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimesheetsByContractorAndClientRequest) ProtoMessage() {}

func (x *GetTimesheetsByContractorAndClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesheet_timesheet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimesheetsByContractorAndClientRequest.ProtoReflect.Descriptor instead.
func (*GetTimesheetsByContractorAndClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_timesheet_timesheet_proto_rawDescGZIP(), []int{3}
}

func (x *GetTimesheetsByContractorAndClientRequest) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

func (x *GetTimesheetsByContractorAndClientRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type GetTimesheetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timesheets []*Timesheet `protobuf:"bytes,1,rep,name=timesheets,proto3" json:"timesheets,omitempty"`
}

func (x *GetTimesheetsResponse) Reset() {
	*x = GetTimesheetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_timesheet_timesheet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimesheetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimesheetsResponse) ProtoMessage() {}

func (x *GetTimesheetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_timesheet_timesheet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimesheetsResponse.ProtoReflect.Descriptor instead.
func (*GetTimesheetsResponse) Descriptor() ([]byte, []int) {
	return file_proto_timesheet_timesheet_proto_rawDescGZIP(), []int{4}
}

func (x *GetTimesheetsResponse) GetTimesheets() []*Timesheet {
	if x != nil {
		return x.Timesheets
	}
	return nil
}

var File_proto_timesheet_timesheet_proto protoreflect.FileDescriptor

var file_proto_timesheet_timesheet_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70,
	0x22, 0xeb, 0x02, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65,
	0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x42,
	0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x6e, 0x6f, 0x6e, 0x42, 0x69, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x48,
	0x6f, 0x75, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x6f, 0x72,
	0x6b, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x70, 0x61, 0x79, 0x22, 0xc4,
	0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x47, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x6d,
	0x0a, 0x29, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x51, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73,
	0x32, 0xe1, 0x02, 0x0a, 0x10, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x72, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x5f, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x6e,
	0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_timesheet_timesheet_proto_rawDescOnce sync.Once
	file_proto_timesheet_timesheet_proto_rawDescData = file_proto_timesheet_timesheet_proto_rawDesc
)

func file_proto_timesheet_timesheet_proto_rawDescGZIP() []byte {
	file_proto_timesheet_timesheet_proto_rawDescOnce.Do(func() {
		file_proto_timesheet_timesheet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_timesheet_timesheet_proto_rawDescData)
	})
	return file_proto_timesheet_timesheet_proto_rawDescData
}

var file_proto_timesheet_timesheet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_timesheet_timesheet_proto_goTypes = []interface{}{
	(*Timesheet)(nil),                                 // 0: timesheet_app.Timesheet
	(*CreateTimesheetRequest)(nil),                    // 1: timesheet_app.CreateTimesheetRequest
	(*GetTimesheetsByContractorRequest)(nil),          // 2: timesheet_app.GetTimesheetsByContractorRequest
	(*GetTimesheetsByContractorAndClientRequest)(nil), // 3: timesheet_app.GetTimesheetsByContractorAndClientRequest
	(*GetTimesheetsResponse)(nil),                     // 4: timesheet_app.GetTimesheetsResponse
}
var file_proto_timesheet_timesheet_proto_depIdxs = []int32{
	0, // 0: timesheet_app.GetTimesheetsResponse.timesheets:type_name -> timesheet_app.Timesheet
	1, // 1: timesheet_app.TimesheetService.CreateTimesheet:input_type -> timesheet_app.CreateTimesheetRequest
	2, // 2: timesheet_app.TimesheetService.GetTimesheetsByContractor:input_type -> timesheet_app.GetTimesheetsByContractorRequest
	3, // 3: timesheet_app.TimesheetService.GetTimesheetsByContractorAndClient:input_type -> timesheet_app.GetTimesheetsByContractorAndClientRequest
	0, // 4: timesheet_app.TimesheetService.CreateTimesheet:output_type -> timesheet_app.Timesheet
	4, // 5: timesheet_app.TimesheetService.GetTimesheetsByContractor:output_type -> timesheet_app.GetTimesheetsResponse
	4, // 6: timesheet_app.TimesheetService.GetTimesheetsByContractorAndClient:output_type -> timesheet_app.GetTimesheetsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_timesheet_timesheet_proto_init() }
func file_proto_timesheet_timesheet_proto_init() {
	if File_proto_timesheet_timesheet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_timesheet_timesheet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timesheet); i {
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
		file_proto_timesheet_timesheet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTimesheetRequest); i {
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
		file_proto_timesheet_timesheet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimesheetsByContractorRequest); i {
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
		file_proto_timesheet_timesheet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimesheetsByContractorAndClientRequest); i {
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
		file_proto_timesheet_timesheet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimesheetsResponse); i {
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
			RawDescriptor: file_proto_timesheet_timesheet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_timesheet_timesheet_proto_goTypes,
		DependencyIndexes: file_proto_timesheet_timesheet_proto_depIdxs,
		MessageInfos:      file_proto_timesheet_timesheet_proto_msgTypes,
	}.Build()
	File_proto_timesheet_timesheet_proto = out.File
	file_proto_timesheet_timesheet_proto_rawDesc = nil
	file_proto_timesheet_timesheet_proto_goTypes = nil
	file_proto_timesheet_timesheet_proto_depIdxs = nil
}
