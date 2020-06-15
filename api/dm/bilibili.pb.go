// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.1
// source: bilibili.proto

package dm

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

type DmSegMobileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elems []*DanmakuElem `protobuf:"bytes,1,rep,name=elems,proto3" json:"elems,omitempty"`
}

func (x *DmSegMobileReply) Reset() {
	*x = DmSegMobileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DmSegMobileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DmSegMobileReply) ProtoMessage() {}

func (x *DmSegMobileReply) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DmSegMobileReply.ProtoReflect.Descriptor instead.
func (*DmSegMobileReply) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{0}
}

func (x *DmSegMobileReply) GetElems() []*DanmakuElem {
	if x != nil {
		return x.Elems
	}
	return nil
}

type DanmakuElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Progress int32  `protobuf:"varint,2,opt,name=progress,proto3" json:"progress,omitempty"`
	Mode     int32  `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Fontsize int32  `protobuf:"varint,4,opt,name=fontsize,proto3" json:"fontsize,omitempty"`
	Color    uint32 `protobuf:"varint,5,opt,name=color,proto3" json:"color,omitempty"`
	MidHash  string `protobuf:"bytes,6,opt,name=midHash,proto3" json:"midHash,omitempty"`
	Content  string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Ctime    int64  `protobuf:"varint,8,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Weight   int32  `protobuf:"varint,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Action   string `protobuf:"bytes,10,opt,name=action,proto3" json:"action,omitempty"`
	Pool     int32  `protobuf:"varint,11,opt,name=pool,proto3" json:"pool,omitempty"`
	IdStr    string `protobuf:"bytes,12,opt,name=idStr,proto3" json:"idStr,omitempty"`
}

func (x *DanmakuElem) Reset() {
	*x = DanmakuElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DanmakuElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DanmakuElem) ProtoMessage() {}

func (x *DanmakuElem) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DanmakuElem.ProtoReflect.Descriptor instead.
func (*DanmakuElem) Descriptor() ([]byte, []int) {
	return file_bilibili_proto_rawDescGZIP(), []int{1}
}

func (x *DanmakuElem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DanmakuElem) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *DanmakuElem) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *DanmakuElem) GetFontsize() int32 {
	if x != nil {
		return x.Fontsize
	}
	return 0
}

func (x *DanmakuElem) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *DanmakuElem) GetMidHash() string {
	if x != nil {
		return x.MidHash
	}
	return ""
}

func (x *DanmakuElem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DanmakuElem) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *DanmakuElem) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *DanmakuElem) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *DanmakuElem) GetPool() int32 {
	if x != nil {
		return x.Pool
	}
	return 0
}

func (x *DanmakuElem) GetIdStr() string {
	if x != nil {
		return x.IdStr
	}
	return ""
}

var File_bilibili_proto protoreflect.FileDescriptor

var file_bilibili_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x64, 0x6d, 0x22, 0x39, 0x0a, 0x10, 0x44, 0x6d, 0x53, 0x65, 0x67, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x6c, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6d, 0x2e, 0x44, 0x61, 0x6e,
	0x6d, 0x61, 0x6b, 0x75, 0x45, 0x6c, 0x65, 0x6d, 0x52, 0x05, 0x65, 0x6c, 0x65, 0x6d, 0x73, 0x22,
	0xa3, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x6e, 0x6d, 0x61, 0x6b, 0x75, 0x45, 0x6c, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x6f, 0x6e, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x64, 0x53, 0x74, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x64, 0x53, 0x74, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x64, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_proto_rawDescOnce sync.Once
	file_bilibili_proto_rawDescData = file_bilibili_proto_rawDesc
)

func file_bilibili_proto_rawDescGZIP() []byte {
	file_bilibili_proto_rawDescOnce.Do(func() {
		file_bilibili_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_proto_rawDescData)
	})
	return file_bilibili_proto_rawDescData
}

var file_bilibili_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bilibili_proto_goTypes = []interface{}{
	(*DmSegMobileReply)(nil), // 0: dm.DmSegMobileReply
	(*DanmakuElem)(nil),      // 1: dm.DanmakuElem
}
var file_bilibili_proto_depIdxs = []int32{
	1, // 0: dm.DmSegMobileReply.elems:type_name -> dm.DanmakuElem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bilibili_proto_init() }
func file_bilibili_proto_init() {
	if File_bilibili_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DmSegMobileReply); i {
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
		file_bilibili_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DanmakuElem); i {
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
			RawDescriptor: file_bilibili_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_proto_goTypes,
		DependencyIndexes: file_bilibili_proto_depIdxs,
		MessageInfos:      file_bilibili_proto_msgTypes,
	}.Build()
	File_bilibili_proto = out.File
	file_bilibili_proto_rawDesc = nil
	file_bilibili_proto_goTypes = nil
	file_bilibili_proto_depIdxs = nil
}
