// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: nsx/testapp/user.proto

package testapp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TournamentType int32

const (
	TournamentType_TOURNAMENT_TYPE_UNKNOWN TournamentType = 0
	TournamentType_TOURNAMENT_TYPE_DAILY   TournamentType = 1
	TournamentType_TOURNAMENT_TYPE_WEEKLY  TournamentType = 2
	TournamentType_TOURNAMENT_TYPE_MONTHLY TournamentType = 3
)

// Enum value maps for TournamentType.
var (
	TournamentType_name = map[int32]string{
		0: "TOURNAMENT_TYPE_UNKNOWN",
		1: "TOURNAMENT_TYPE_DAILY",
		2: "TOURNAMENT_TYPE_WEEKLY",
		3: "TOURNAMENT_TYPE_MONTHLY",
	}
	TournamentType_value = map[string]int32{
		"TOURNAMENT_TYPE_UNKNOWN": 0,
		"TOURNAMENT_TYPE_DAILY":   1,
		"TOURNAMENT_TYPE_WEEKLY":  2,
		"TOURNAMENT_TYPE_MONTHLY": 3,
	}
)

func (x TournamentType) Enum() *TournamentType {
	p := new(TournamentType)
	*p = x
	return p
}

func (x TournamentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TournamentType) Descriptor() protoreflect.EnumDescriptor {
	return file_nsx_testapp_user_proto_enumTypes[0].Descriptor()
}

func (TournamentType) Type() protoreflect.EnumType {
	return &file_nsx_testapp_user_proto_enumTypes[0]
}

func (x TournamentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TournamentType.Descriptor instead.
func (TournamentType) EnumDescriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{0}
}

// UserDetailsRequest contains the parameters that will be used to vary the cache with.
type UserDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserDetailsRequest) Reset() {
	*x = UserDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsRequest) ProtoMessage() {}

func (x *UserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsRequest.ProtoReflect.Descriptor instead.
func (*UserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetailsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// UserDetailsResponse contains the response that will be cached.
type UserDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserDetailsResponse) Reset() {
	*x = UserDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetailsResponse) ProtoMessage() {}

func (x *UserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetailsResponse.ProtoReflect.Descriptor instead.
func (*UserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetailsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type MainTournamentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *MainTournamentsRequest) Reset() {
	*x = MainTournamentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainTournamentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainTournamentsRequest) ProtoMessage() {}

func (x *MainTournamentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainTournamentsRequest.ProtoReflect.Descriptor instead.
func (*MainTournamentsRequest) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{3}
}

func (x *MainTournamentsRequest) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type MainTournamentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournaments []*Tournament `protobuf:"bytes,1,rep,name=tournaments,proto3" json:"tournaments,omitempty"`
}

func (x *MainTournamentsResponse) Reset() {
	*x = MainTournamentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainTournamentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainTournamentsResponse) ProtoMessage() {}

func (x *MainTournamentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainTournamentsResponse.ProtoReflect.Descriptor instead.
func (*MainTournamentsResponse) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{4}
}

func (x *MainTournamentsResponse) GetTournaments() []*Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

type Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string         `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Url      string         `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Dbl      float64        `protobuf:"fixed64,5,opt,name=dbl,proto3" json:"dbl,omitempty"`
	Flt      float32        `protobuf:"fixed32,6,opt,name=flt,proto3" json:"flt,omitempty"`
	Num32    int32          `protobuf:"varint,7,opt,name=num32,proto3" json:"num32,omitempty"`
	Num64    int64          `protobuf:"varint,8,opt,name=num64,proto3" json:"num64,omitempty"`
	Unum32   uint32         `protobuf:"varint,9,opt,name=unum32,proto3" json:"unum32,omitempty"`
	Unum64   uint64         `protobuf:"varint,10,opt,name=unum64,proto3" json:"unum64,omitempty"`
	Snum32   int32          `protobuf:"zigzag32,11,opt,name=snum32,proto3" json:"snum32,omitempty"`
	Snum64   int64          `protobuf:"zigzag64,12,opt,name=snum64,proto3" json:"snum64,omitempty"`
	Fnum32   uint32         `protobuf:"fixed32,13,opt,name=fnum32,proto3" json:"fnum32,omitempty"`
	Fnum64   uint64         `protobuf:"fixed64,14,opt,name=fnum64,proto3" json:"fnum64,omitempty"`
	Sfnum32  int32          `protobuf:"fixed32,15,opt,name=sfnum32,proto3" json:"sfnum32,omitempty"`
	Sfnum64  int64          `protobuf:"fixed64,16,opt,name=sfnum64,proto3" json:"sfnum64,omitempty"`
	IsActive bool           `protobuf:"varint,17,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Data     []byte         `protobuf:"bytes,18,opt,name=data,proto3" json:"data,omitempty"`
	Type     TournamentType `protobuf:"varint,19,opt,name=type,proto3,enum=nsx.testapp.TournamentType" json:"type,omitempty"`
	Events   []*Event       `protobuf:"bytes,20,rep,name=events,proto3" json:"events,omitempty"`
	// Types that are assignable to Prize:
	//
	//	*Tournament_PrizeVal
	//	*Tournament_PrizeNum
	Prize    isTournament_Prize `protobuf_oneof:"prize"`
	Metadata map[string]string  `protobuf:"bytes,23,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Tournament) Reset() {
	*x = Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament) ProtoMessage() {}

func (x *Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament.ProtoReflect.Descriptor instead.
func (*Tournament) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{5}
}

func (x *Tournament) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tournament) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tournament) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Tournament) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Tournament) GetDbl() float64 {
	if x != nil {
		return x.Dbl
	}
	return 0
}

func (x *Tournament) GetFlt() float32 {
	if x != nil {
		return x.Flt
	}
	return 0
}

func (x *Tournament) GetNum32() int32 {
	if x != nil {
		return x.Num32
	}
	return 0
}

func (x *Tournament) GetNum64() int64 {
	if x != nil {
		return x.Num64
	}
	return 0
}

func (x *Tournament) GetUnum32() uint32 {
	if x != nil {
		return x.Unum32
	}
	return 0
}

func (x *Tournament) GetUnum64() uint64 {
	if x != nil {
		return x.Unum64
	}
	return 0
}

func (x *Tournament) GetSnum32() int32 {
	if x != nil {
		return x.Snum32
	}
	return 0
}

func (x *Tournament) GetSnum64() int64 {
	if x != nil {
		return x.Snum64
	}
	return 0
}

func (x *Tournament) GetFnum32() uint32 {
	if x != nil {
		return x.Fnum32
	}
	return 0
}

func (x *Tournament) GetFnum64() uint64 {
	if x != nil {
		return x.Fnum64
	}
	return 0
}

func (x *Tournament) GetSfnum32() int32 {
	if x != nil {
		return x.Sfnum32
	}
	return 0
}

func (x *Tournament) GetSfnum64() int64 {
	if x != nil {
		return x.Sfnum64
	}
	return 0
}

func (x *Tournament) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Tournament) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Tournament) GetType() TournamentType {
	if x != nil {
		return x.Type
	}
	return TournamentType_TOURNAMENT_TYPE_UNKNOWN
}

func (x *Tournament) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (m *Tournament) GetPrize() isTournament_Prize {
	if m != nil {
		return m.Prize
	}
	return nil
}

func (x *Tournament) GetPrizeVal() string {
	if x, ok := x.GetPrize().(*Tournament_PrizeVal); ok {
		return x.PrizeVal
	}
	return ""
}

func (x *Tournament) GetPrizeNum() int32 {
	if x, ok := x.GetPrize().(*Tournament_PrizeNum); ok {
		return x.PrizeNum
	}
	return 0
}

func (x *Tournament) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isTournament_Prize interface {
	isTournament_Prize()
}

type Tournament_PrizeVal struct {
	PrizeVal string `protobuf:"bytes,21,opt,name=prize_val,json=prizeVal,proto3,oneof"`
}

type Tournament_PrizeNum struct {
	PrizeNum int32 `protobuf:"varint,22,opt,name=prize_num,json=prizeNum,proto3,oneof"`
}

func (*Tournament_PrizeVal) isTournament_Prize() {}

func (*Tournament_PrizeNum) isTournament_Prize() {}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players   []string               `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nsx_testapp_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_nsx_testapp_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_nsx_testapp_user_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetPlayers() []string {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_nsx_testapp_user_proto protoreflect.FileDescriptor

var file_nsx_testapp_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x73, 0x78, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x49, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x46, 0x0a, 0x16, 0x4d,
	0x61, 0x69, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x54, 0x0a, 0x17, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0b, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x70, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xc8, 0x05, 0x0a, 0x0a, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x62, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x64, 0x62, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6e, 0x75, 0x6d, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x6e, 0x75,
	0x6d, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x6e, 0x75,
	0x6d, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x07, 0x52, 0x06, 0x66, 0x6e, 0x75,
	0x6d, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x06, 0x52, 0x06, 0x66, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x66, 0x6e, 0x75, 0x6d, 0x33, 0x32, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x07, 0x73, 0x66,
	0x6e, 0x75, 0x6d, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x66, 0x6e, 0x75, 0x6d, 0x36, 0x34,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x10, 0x52, 0x07, 0x73, 0x66, 0x6e, 0x75, 0x6d, 0x36, 0x34, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x09, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x41, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x7a, 0x65, 0x22, 0x70, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2a, 0x81, 0x01, 0x0a, 0x0e, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4f, 0x55,
	0x52, 0x4e, 0x41, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x41,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x4f, 0x55, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x03, 0x32, 0x5f, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x71, 0x0a, 0x0f, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x5e,
	0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x23, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x70, 0x70, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xab,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x73, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x70, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x53, 0x58, 0x42,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x73, 0x78, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x70, 0xa2, 0x02, 0x03, 0x4e, 0x54, 0x58, 0xaa, 0x02, 0x0b, 0x4e, 0x73, 0x78, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x70, 0xca, 0x02, 0x0b, 0x4e, 0x73, 0x78, 0x5c, 0x54, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x70, 0xe2, 0x02, 0x17, 0x4e, 0x73, 0x78, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x4e, 0x73, 0x78, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nsx_testapp_user_proto_rawDescOnce sync.Once
	file_nsx_testapp_user_proto_rawDescData = file_nsx_testapp_user_proto_rawDesc
)

func file_nsx_testapp_user_proto_rawDescGZIP() []byte {
	file_nsx_testapp_user_proto_rawDescOnce.Do(func() {
		file_nsx_testapp_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_nsx_testapp_user_proto_rawDescData)
	})
	return file_nsx_testapp_user_proto_rawDescData
}

var file_nsx_testapp_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nsx_testapp_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nsx_testapp_user_proto_goTypes = []any{
	(TournamentType)(0),             // 0: nsx.testapp.TournamentType
	(*UserDetailsRequest)(nil),      // 1: nsx.testapp.UserDetailsRequest
	(*UserDetailsResponse)(nil),     // 2: nsx.testapp.UserDetailsResponse
	(*User)(nil),                    // 3: nsx.testapp.User
	(*MainTournamentsRequest)(nil),  // 4: nsx.testapp.MainTournamentsRequest
	(*MainTournamentsResponse)(nil), // 5: nsx.testapp.MainTournamentsResponse
	(*Tournament)(nil),              // 6: nsx.testapp.Tournament
	(*Event)(nil),                   // 7: nsx.testapp.Event
	nil,                             // 8: nsx.testapp.Tournament.MetadataEntry
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
	(*timestamppb.Timestamp)(nil),   // 10: google.protobuf.Timestamp
}
var file_nsx_testapp_user_proto_depIdxs = []int32{
	3,  // 0: nsx.testapp.UserDetailsResponse.user:type_name -> nsx.testapp.User
	9,  // 1: nsx.testapp.MainTournamentsRequest.empty:type_name -> google.protobuf.Empty
	6,  // 2: nsx.testapp.MainTournamentsResponse.tournaments:type_name -> nsx.testapp.Tournament
	0,  // 3: nsx.testapp.Tournament.type:type_name -> nsx.testapp.TournamentType
	7,  // 4: nsx.testapp.Tournament.events:type_name -> nsx.testapp.Event
	8,  // 5: nsx.testapp.Tournament.metadata:type_name -> nsx.testapp.Tournament.MetadataEntry
	10, // 6: nsx.testapp.Event.start_time:type_name -> google.protobuf.Timestamp
	1,  // 7: nsx.testapp.UserCache.UserDetails:input_type -> nsx.testapp.UserDetailsRequest
	4,  // 8: nsx.testapp.TournamentCache.MainTournaments:input_type -> nsx.testapp.MainTournamentsRequest
	2,  // 9: nsx.testapp.UserCache.UserDetails:output_type -> nsx.testapp.UserDetailsResponse
	5,  // 10: nsx.testapp.TournamentCache.MainTournaments:output_type -> nsx.testapp.MainTournamentsResponse
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_nsx_testapp_user_proto_init() }
func file_nsx_testapp_user_proto_init() {
	if File_nsx_testapp_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nsx_testapp_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserDetailsRequest); i {
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
		file_nsx_testapp_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserDetailsResponse); i {
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
		file_nsx_testapp_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_nsx_testapp_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MainTournamentsRequest); i {
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
		file_nsx_testapp_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MainTournamentsResponse); i {
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
		file_nsx_testapp_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Tournament); i {
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
		file_nsx_testapp_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
	file_nsx_testapp_user_proto_msgTypes[5].OneofWrappers = []any{
		(*Tournament_PrizeVal)(nil),
		(*Tournament_PrizeNum)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nsx_testapp_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_nsx_testapp_user_proto_goTypes,
		DependencyIndexes: file_nsx_testapp_user_proto_depIdxs,
		EnumInfos:         file_nsx_testapp_user_proto_enumTypes,
		MessageInfos:      file_nsx_testapp_user_proto_msgTypes,
	}.Build()
	File_nsx_testapp_user_proto = out.File
	file_nsx_testapp_user_proto_rawDesc = nil
	file_nsx_testapp_user_proto_goTypes = nil
	file_nsx_testapp_user_proto_depIdxs = nil
}
