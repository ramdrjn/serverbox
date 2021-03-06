// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: state.proto

package sb_state_proto

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

type RegisterReq_Type int32

const (
	RegisterReq_SERVER RegisterReq_Type = 0
)

// Enum value maps for RegisterReq_Type.
var (
	RegisterReq_Type_name = map[int32]string{
		0: "SERVER",
	}
	RegisterReq_Type_value = map[string]int32{
		"SERVER": 0,
	}
)

func (x RegisterReq_Type) Enum() *RegisterReq_Type {
	p := new(RegisterReq_Type)
	*p = x
	return p
}

func (x RegisterReq_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterReq_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[0].Descriptor()
}

func (RegisterReq_Type) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[0]
}

func (x RegisterReq_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterReq_Type.Descriptor instead.
func (RegisterReq_Type) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0, 0}
}

type ReportReq_State int32

const (
	ReportReq_UP          ReportReq_State = 0
	ReportReq_DOWN        ReportReq_State = 1
	ReportReq_MAINTANENCE ReportReq_State = 2
)

// Enum value maps for ReportReq_State.
var (
	ReportReq_State_name = map[int32]string{
		0: "UP",
		1: "DOWN",
		2: "MAINTANENCE",
	}
	ReportReq_State_value = map[string]int32{
		"UP":          0,
		"DOWN":        1,
		"MAINTANENCE": 2,
	}
)

func (x ReportReq_State) Enum() *ReportReq_State {
	p := new(ReportReq_State)
	*p = x
	return p
}

func (x ReportReq_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportReq_State) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[1].Descriptor()
}

func (ReportReq_State) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[1]
}

func (x ReportReq_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportReq_State.Descriptor instead.
func (ReportReq_State) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2, 0}
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string           `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Type RegisterReq_Type `protobuf:"varint,2,opt,name=type,proto3,enum=sb_state_proto.RegisterReq_Type" json:"type,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RegisterReq) GetType() RegisterReq_Type {
	if x != nil {
		return x.Type
	}
	return RegisterReq_SERVER
}

type RegisterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterRes) Reset() {
	*x = RegisterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRes) ProtoMessage() {}

func (x *RegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRes.ProtoReflect.Descriptor instead.
func (*RegisterRes) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

type ReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUuid   string          `protobuf:"bytes,1,opt,name=target_uuid,json=targetUuid,proto3" json:"target_uuid,omitempty"`
	State        ReportReq_State `protobuf:"varint,2,opt,name=state,proto3,enum=sb_state_proto.ReportReq_State" json:"state,omitempty"`
	ReporteeUuid string          `protobuf:"bytes,3,opt,name=reportee_uuid,json=reporteeUuid,proto3" json:"reportee_uuid,omitempty"`
	Reason       string          `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ReportReq) Reset() {
	*x = ReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReq) ProtoMessage() {}

func (x *ReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReq.ProtoReflect.Descriptor instead.
func (*ReportReq) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{2}
}

func (x *ReportReq) GetTargetUuid() string {
	if x != nil {
		return x.TargetUuid
	}
	return ""
}

func (x *ReportReq) GetState() ReportReq_State {
	if x != nil {
		return x.State
	}
	return ReportReq_UP
}

func (x *ReportReq) GetReporteeUuid() string {
	if x != nil {
		return x.ReporteeUuid
	}
	return ""
}

func (x *ReportReq) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ReportRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportRes) Reset() {
	*x = ReportRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRes) ProtoMessage() {}

func (x *ReportRes) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRes.ProtoReflect.Descriptor instead.
func (*ReportRes) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{3}
}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73,
	0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x73, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x00, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x09, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x62, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x65,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x41, 0x49, 0x4e, 0x54,
	0x41, 0x4e, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x22, 0x0b, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x32, 0x9e, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x4e, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x73, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x73, 0x62, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x73, 0x62, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_state_proto_goTypes = []interface{}{
	(RegisterReq_Type)(0), // 0: sb_state_proto.RegisterReq.Type
	(ReportReq_State)(0),  // 1: sb_state_proto.ReportReq.State
	(*RegisterReq)(nil),   // 2: sb_state_proto.RegisterReq
	(*RegisterRes)(nil),   // 3: sb_state_proto.RegisterRes
	(*ReportReq)(nil),     // 4: sb_state_proto.ReportReq
	(*ReportRes)(nil),     // 5: sb_state_proto.ReportRes
}
var file_state_proto_depIdxs = []int32{
	0, // 0: sb_state_proto.RegisterReq.type:type_name -> sb_state_proto.RegisterReq.Type
	1, // 1: sb_state_proto.ReportReq.state:type_name -> sb_state_proto.ReportReq.State
	2, // 2: sb_state_proto.State.RegisterForState:input_type -> sb_state_proto.RegisterReq
	4, // 3: sb_state_proto.State.ReportState:input_type -> sb_state_proto.ReportReq
	3, // 4: sb_state_proto.State.RegisterForState:output_type -> sb_state_proto.RegisterRes
	5, // 5: sb_state_proto.State.ReportState:output_type -> sb_state_proto.ReportRes
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRes); i {
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
		file_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReq); i {
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
		file_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRes); i {
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
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		EnumInfos:         file_state_proto_enumTypes,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}
