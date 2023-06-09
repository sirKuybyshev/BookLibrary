// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: protos/librarian.proto

package pb

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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_librarian_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_protos_librarian_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_protos_librarian_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Text     []byte  `protobuf:"bytes,2,opt,name=Text,proto3,oneof" json:"Text,omitempty"`         //sometimes we want just know book's authors by its name
	Encoding *string `protobuf:"bytes,3,opt,name=Encoding,proto3,oneof" json:"Encoding,omitempty"` //default -- no encoding, plain text
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_librarian_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_protos_librarian_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_protos_librarian_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *Book) GetEncoding() string {
	if x != nil && x.Encoding != nil {
		return *x.Encoding
	}
	return ""
}

var File_protos_librarian_proto protoreflect.FileDescriptor

var file_protos_librarian_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x22, 0x42, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x54, 0x65, 0x78, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x32, 0x72, 0x0a, 0x09, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x11, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a,
	0x0f, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x0f, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x1a, 0x11, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_librarian_proto_rawDescOnce sync.Once
	file_protos_librarian_proto_rawDescData = file_protos_librarian_proto_rawDesc
)

func file_protos_librarian_proto_rawDescGZIP() []byte {
	file_protos_librarian_proto_rawDescOnce.Do(func() {
		file_protos_librarian_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_librarian_proto_rawDescData)
	})
	return file_protos_librarian_proto_rawDescData
}

var file_protos_librarian_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_librarian_proto_goTypes = []interface{}{
	(*Author)(nil), // 0: librarian.Author
	(*Book)(nil),   // 1: librarian.Book
}
var file_protos_librarian_proto_depIdxs = []int32{
	0, // 0: librarian.Librarian.GetBooks:input_type -> librarian.Author
	1, // 1: librarian.Librarian.GetAuthor:input_type -> librarian.Book
	1, // 2: librarian.Librarian.GetBooks:output_type -> librarian.Book
	0, // 3: librarian.Librarian.GetAuthor:output_type -> librarian.Author
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_librarian_proto_init() }
func file_protos_librarian_proto_init() {
	if File_protos_librarian_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_librarian_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_protos_librarian_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
	file_protos_librarian_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_librarian_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_librarian_proto_goTypes,
		DependencyIndexes: file_protos_librarian_proto_depIdxs,
		MessageInfos:      file_protos_librarian_proto_msgTypes,
	}.Build()
	File_protos_librarian_proto = out.File
	file_protos_librarian_proto_rawDesc = nil
	file_protos_librarian_proto_goTypes = nil
	file_protos_librarian_proto_depIdxs = nil
}
