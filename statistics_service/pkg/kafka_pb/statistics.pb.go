// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: statistics_service/pkg/kafka_pb/statistics.proto

package kafka_pb

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

type ViewEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PostId   int64 `protobuf:"varint,2,opt,name=postId,proto3" json:"postId,omitempty"`
	AuthorId int64 `protobuf:"varint,3,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *ViewEvent) Reset() {
	*x = ViewEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewEvent) ProtoMessage() {}

func (x *ViewEvent) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewEvent.ProtoReflect.Descriptor instead.
func (*ViewEvent) Descriptor() ([]byte, []int) {
	return file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *ViewEvent) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ViewEvent) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *ViewEvent) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type LikeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	PostId   int64 `protobuf:"varint,2,opt,name=postId,proto3" json:"postId,omitempty"`
	AuthorId int64 `protobuf:"varint,3,opt,name=authorId,proto3" json:"authorId,omitempty"`
}

func (x *LikeEvent) Reset() {
	*x = LikeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeEvent) ProtoMessage() {}

func (x *LikeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeEvent.ProtoReflect.Descriptor instead.
func (*LikeEvent) Descriptor() ([]byte, []int) {
	return file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *LikeEvent) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeEvent) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *LikeEvent) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EventType:
	//
	//	*Event_ViewEvent
	//	*Event_LikeEvent
	EventType isEvent_EventType `protobuf_oneof:"event_type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[2]
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
	return file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescGZIP(), []int{2}
}

func (m *Event) GetEventType() isEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (x *Event) GetViewEvent() *ViewEvent {
	if x, ok := x.GetEventType().(*Event_ViewEvent); ok {
		return x.ViewEvent
	}
	return nil
}

func (x *Event) GetLikeEvent() *LikeEvent {
	if x, ok := x.GetEventType().(*Event_LikeEvent); ok {
		return x.LikeEvent
	}
	return nil
}

type isEvent_EventType interface {
	isEvent_EventType()
}

type Event_ViewEvent struct {
	ViewEvent *ViewEvent `protobuf:"bytes,1,opt,name=view_event,json=viewEvent,proto3,oneof"`
}

type Event_LikeEvent struct {
	LikeEvent *LikeEvent `protobuf:"bytes,2,opt,name=like_event,json=likeEvent,proto3,oneof"`
}

func (*Event_ViewEvent) isEvent_EventType() {}

func (*Event_LikeEvent) isEvent_EventType() {}

var File_statistics_service_pkg_kafka_pb_statistics_proto protoreflect.FileDescriptor

var file_statistics_service_pkg_kafka_pb_statistics_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x70,
	0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x57,
	0x0a, 0x09, 0x56, 0x69, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x85, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x6c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescOnce sync.Once
	file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescData = file_statistics_service_pkg_kafka_pb_statistics_proto_rawDesc
)

func file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescGZIP() []byte {
	file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescOnce.Do(func() {
		file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescData)
	})
	return file_statistics_service_pkg_kafka_pb_statistics_proto_rawDescData
}

var file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_statistics_service_pkg_kafka_pb_statistics_proto_goTypes = []interface{}{
	(*ViewEvent)(nil), // 0: statistics.ViewEvent
	(*LikeEvent)(nil), // 1: statistics.LikeEvent
	(*Event)(nil),     // 2: statistics.Event
}
var file_statistics_service_pkg_kafka_pb_statistics_proto_depIdxs = []int32{
	0, // 0: statistics.Event.view_event:type_name -> statistics.ViewEvent
	1, // 1: statistics.Event.like_event:type_name -> statistics.LikeEvent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_statistics_service_pkg_kafka_pb_statistics_proto_init() }
func file_statistics_service_pkg_kafka_pb_statistics_proto_init() {
	if File_statistics_service_pkg_kafka_pb_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewEvent); i {
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
		file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeEvent); i {
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
		file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Event_ViewEvent)(nil),
		(*Event_LikeEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_statistics_service_pkg_kafka_pb_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_statistics_service_pkg_kafka_pb_statistics_proto_goTypes,
		DependencyIndexes: file_statistics_service_pkg_kafka_pb_statistics_proto_depIdxs,
		MessageInfos:      file_statistics_service_pkg_kafka_pb_statistics_proto_msgTypes,
	}.Build()
	File_statistics_service_pkg_kafka_pb_statistics_proto = out.File
	file_statistics_service_pkg_kafka_pb_statistics_proto_rawDesc = nil
	file_statistics_service_pkg_kafka_pb_statistics_proto_goTypes = nil
	file_statistics_service_pkg_kafka_pb_statistics_proto_depIdxs = nil
}
