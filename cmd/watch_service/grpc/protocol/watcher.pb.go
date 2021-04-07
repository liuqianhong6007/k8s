// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: watcher.proto

package protocol

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

type WatchResponse_ResponseStatus int32

const (
	WatchResponse_Ok           WatchResponse_ResponseStatus = 0 // 成功
	WatchResponse_ParamError   WatchResponse_ResponseStatus = 1 // 请求参数错误
	WatchResponse_UnknownError WatchResponse_ResponseStatus = 2 // 未知错误
)

// Enum value maps for WatchResponse_ResponseStatus.
var (
	WatchResponse_ResponseStatus_name = map[int32]string{
		0: "Ok",
		1: "ParamError",
		2: "UnknownError",
	}
	WatchResponse_ResponseStatus_value = map[string]int32{
		"Ok":           0,
		"ParamError":   1,
		"UnknownError": 2,
	}
)

func (x WatchResponse_ResponseStatus) Enum() *WatchResponse_ResponseStatus {
	p := new(WatchResponse_ResponseStatus)
	*p = x
	return p
}

func (x WatchResponse_ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchResponse_ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_watcher_proto_enumTypes[0].Descriptor()
}

func (WatchResponse_ResponseStatus) Type() protoreflect.EnumType {
	return &file_watcher_proto_enumTypes[0]
}

func (x WatchResponse_ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchResponse_ResponseStatus.Descriptor instead.
func (WatchResponse_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_watcher_proto_rawDescGZIP(), []int{1, 0}
}

type WatchResponse_Action int32

const (
	WatchResponse_Add    WatchResponse_Action = 0 // 添加
	WatchResponse_Delete WatchResponse_Action = 1 // 删除
	WatchResponse_Modify WatchResponse_Action = 2 // 修改
)

// Enum value maps for WatchResponse_Action.
var (
	WatchResponse_Action_name = map[int32]string{
		0: "Add",
		1: "Delete",
		2: "Modify",
	}
	WatchResponse_Action_value = map[string]int32{
		"Add":    0,
		"Delete": 1,
		"Modify": 2,
	}
)

func (x WatchResponse_Action) Enum() *WatchResponse_Action {
	p := new(WatchResponse_Action)
	*p = x
	return p
}

func (x WatchResponse_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WatchResponse_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_watcher_proto_enumTypes[1].Descriptor()
}

func (WatchResponse_Action) Type() protoreflect.EnumType {
	return &file_watcher_proto_enumTypes[1]
}

func (x WatchResponse_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WatchResponse_Action.Descriptor instead.
func (WatchResponse_Action) EnumDescriptor() ([]byte, []int) {
	return file_watcher_proto_rawDescGZIP(), []int{1, 1}
}

type MatchCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace     string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`                              // 命名空间
	LabelSelector string `protobuf:"bytes,2,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"` // 标签匹配
}

func (x *MatchCondition) Reset() {
	*x = MatchCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCondition) ProtoMessage() {}

func (x *MatchCondition) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCondition.ProtoReflect.Descriptor instead.
func (*MatchCondition) Descriptor() ([]byte, []int) {
	return file_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *MatchCondition) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MatchCondition) GetLabelSelector() string {
	if x != nil {
		return x.LabelSelector
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status WatchResponse_ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=WatchResponse_ResponseStatus" json:"status,omitempty"` // 返回状态
	Name   string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                        // pod 名称
	Addr   string                       `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`                                        // pod 详情
	Action WatchResponse_Action         `protobuf:"varint,4,opt,name=action,proto3,enum=WatchResponse_Action" json:"action,omitempty"`         // 行为
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *WatchResponse) GetStatus() WatchResponse_ResponseStatus {
	if x != nil {
		return x.Status
	}
	return WatchResponse_Ok
}

func (x *WatchResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WatchResponse) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *WatchResponse) GetAction() WatchResponse_Action {
	if x != nil {
		return x.Action
	}
	return WatchResponse_Add
}

var File_watcher_proto protoreflect.FileDescriptor

var file_watcher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x84, 0x02, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x02, 0x22, 0x29, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0x02, 0x32, 0x3f, 0x0a,
	0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x64, 0x12, 0x0f, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0f,
	0x5a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watcher_proto_rawDescOnce sync.Once
	file_watcher_proto_rawDescData = file_watcher_proto_rawDesc
)

func file_watcher_proto_rawDescGZIP() []byte {
	file_watcher_proto_rawDescOnce.Do(func() {
		file_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_watcher_proto_rawDescData)
	})
	return file_watcher_proto_rawDescData
}

var file_watcher_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_watcher_proto_goTypes = []interface{}{
	(WatchResponse_ResponseStatus)(0), // 0: WatchResponse.ResponseStatus
	(WatchResponse_Action)(0),         // 1: WatchResponse.Action
	(*MatchCondition)(nil),            // 2: MatchCondition
	(*WatchResponse)(nil),             // 3: WatchResponse
}
var file_watcher_proto_depIdxs = []int32{
	0, // 0: WatchResponse.status:type_name -> WatchResponse.ResponseStatus
	1, // 1: WatchResponse.action:type_name -> WatchResponse.Action
	2, // 2: WatchService.WatchPod:input_type -> MatchCondition
	3, // 3: WatchService.WatchPod:output_type -> WatchResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_watcher_proto_init() }
func file_watcher_proto_init() {
	if File_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchCondition); i {
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
		file_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
			RawDescriptor: file_watcher_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watcher_proto_goTypes,
		DependencyIndexes: file_watcher_proto_depIdxs,
		EnumInfos:         file_watcher_proto_enumTypes,
		MessageInfos:      file_watcher_proto_msgTypes,
	}.Build()
	File_watcher_proto = out.File
	file_watcher_proto_rawDesc = nil
	file_watcher_proto_goTypes = nil
	file_watcher_proto_depIdxs = nil
}
