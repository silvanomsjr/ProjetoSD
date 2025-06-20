// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: proto/kvs.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Retorno de consultas
type Tupla struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// a chave passada na consulta
	Chave string `protobuf:"bytes,1,opt,name=chave,proto3" json:"chave,omitempty"`
	// valor encontrado
	Valor string `protobuf:"bytes,2,opt,name=valor,proto3" json:"valor,omitempty"`
	// versao do valor para a chave
	Versao        int32 `protobuf:"varint,3,opt,name=versao,proto3" json:"versao,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tupla) Reset() {
	*x = Tupla{}
	mi := &file_proto_kvs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tupla) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tupla) ProtoMessage() {}

func (x *Tupla) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tupla.ProtoReflect.Descriptor instead.
func (*Tupla) Descriptor() ([]byte, []int) {
	return file_proto_kvs_proto_rawDescGZIP(), []int{0}
}

func (x *Tupla) GetChave() string {
	if x != nil {
		return x.Chave
	}
	return ""
}

func (x *Tupla) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Tupla) GetVersao() int32 {
	if x != nil {
		return x.Versao
	}
	return 0
}

// Parametro de entrada para insercoes
type ChaveValor struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// a chave
	Chave string `protobuf:"bytes,1,opt,name=chave,proto3" json:"chave,omitempty"`
	// o valor
	Valor         string `protobuf:"bytes,2,opt,name=valor,proto3" json:"valor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChaveValor) Reset() {
	*x = ChaveValor{}
	mi := &file_proto_kvs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChaveValor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaveValor) ProtoMessage() {}

func (x *ChaveValor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaveValor.ProtoReflect.Descriptor instead.
func (*ChaveValor) Descriptor() ([]byte, []int) {
	return file_proto_kvs_proto_rawDescGZIP(), []int{1}
}

func (x *ChaveValor) GetChave() string {
	if x != nil {
		return x.Chave
	}
	return ""
}

func (x *ChaveValor) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

// Parametro de entrada para consultas
type ChaveVersao struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// a chave pesquisada
	Chave string `protobuf:"bytes,1,opt,name=chave,proto3" json:"chave,omitempty"`
	// a versao pesquisada (opcional)
	Versao        *int32 `protobuf:"varint,2,opt,name=versao,proto3,oneof" json:"versao,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChaveVersao) Reset() {
	*x = ChaveVersao{}
	mi := &file_proto_kvs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChaveVersao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaveVersao) ProtoMessage() {}

func (x *ChaveVersao) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaveVersao.ProtoReflect.Descriptor instead.
func (*ChaveVersao) Descriptor() ([]byte, []int) {
	return file_proto_kvs_proto_rawDescGZIP(), []int{2}
}

func (x *ChaveVersao) GetChave() string {
	if x != nil {
		return x.Chave
	}
	return ""
}

func (x *ChaveVersao) GetVersao() int32 {
	if x != nil && x.Versao != nil {
		return *x.Versao
	}
	return 0
}

// Versao para snapshot
type Versao struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// a versao pesquisada
	Versao        int32 `protobuf:"varint,1,opt,name=versao,proto3" json:"versao,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Versao) Reset() {
	*x = Versao{}
	mi := &file_proto_kvs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Versao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Versao) ProtoMessage() {}

func (x *Versao) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Versao.ProtoReflect.Descriptor instead.
func (*Versao) Descriptor() ([]byte, []int) {
	return file_proto_kvs_proto_rawDescGZIP(), []int{3}
}

func (x *Versao) GetVersao() int32 {
	if x != nil {
		return x.Versao
	}
	return 0
}

var File_proto_kvs_proto protoreflect.FileDescriptor

const file_proto_kvs_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/kvs.proto\x12\x03kvs\"K\n" +
	"\x05Tupla\x12\x14\n" +
	"\x05chave\x18\x01 \x01(\tR\x05chave\x12\x14\n" +
	"\x05valor\x18\x02 \x01(\tR\x05valor\x12\x16\n" +
	"\x06versao\x18\x03 \x01(\x05R\x06versao\"8\n" +
	"\n" +
	"ChaveValor\x12\x14\n" +
	"\x05chave\x18\x01 \x01(\tR\x05chave\x12\x14\n" +
	"\x05valor\x18\x02 \x01(\tR\x05valor\"K\n" +
	"\vChaveVersao\x12\x14\n" +
	"\x05chave\x18\x01 \x01(\tR\x05chave\x12\x1b\n" +
	"\x06versao\x18\x02 \x01(\x05H\x00R\x06versao\x88\x01\x01B\t\n" +
	"\a_versao\" \n" +
	"\x06Versao\x12\x16\n" +
	"\x06versao\x18\x01 \x01(\x05R\x06versao2\xce\x02\n" +
	"\x03KVS\x12(\n" +
	"\x06Insere\x12\x0f.kvs.ChaveValor\x1a\v.kvs.Versao\"\x00\x12*\n" +
	"\bConsulta\x12\x10.kvs.ChaveVersao\x1a\n" +
	".kvs.Tupla\"\x00\x12)\n" +
	"\x06Remove\x12\x10.kvs.ChaveVersao\x1a\v.kvs.Versao\"\x00\x122\n" +
	"\fInsereVarias\x12\x0f.kvs.ChaveValor\x1a\v.kvs.Versao\"\x00(\x010\x01\x124\n" +
	"\x0eConsultaVarias\x12\x10.kvs.ChaveVersao\x1a\n" +
	".kvs.Tupla\"\x00(\x010\x01\x123\n" +
	"\fRemoveVarias\x12\x10.kvs.ChaveVersao\x1a\v.kvs.Versao\"\x00(\x010\x01\x12'\n" +
	"\bSnapshot\x12\v.kvs.Versao\x1a\n" +
	".kvs.Tupla\"\x000\x01B\x1fZ\x1dgithub.com/username/kvs/protob\x06proto3"

var (
	file_proto_kvs_proto_rawDescOnce sync.Once
	file_proto_kvs_proto_rawDescData []byte
)

func file_proto_kvs_proto_rawDescGZIP() []byte {
	file_proto_kvs_proto_rawDescOnce.Do(func() {
		file_proto_kvs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_kvs_proto_rawDesc), len(file_proto_kvs_proto_rawDesc)))
	})
	return file_proto_kvs_proto_rawDescData
}

var file_proto_kvs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_kvs_proto_goTypes = []any{
	(*Tupla)(nil),       // 0: kvs.Tupla
	(*ChaveValor)(nil),  // 1: kvs.ChaveValor
	(*ChaveVersao)(nil), // 2: kvs.ChaveVersao
	(*Versao)(nil),      // 3: kvs.Versao
}
var file_proto_kvs_proto_depIdxs = []int32{
	1, // 0: kvs.KVS.Insere:input_type -> kvs.ChaveValor
	2, // 1: kvs.KVS.Consulta:input_type -> kvs.ChaveVersao
	2, // 2: kvs.KVS.Remove:input_type -> kvs.ChaveVersao
	1, // 3: kvs.KVS.InsereVarias:input_type -> kvs.ChaveValor
	2, // 4: kvs.KVS.ConsultaVarias:input_type -> kvs.ChaveVersao
	2, // 5: kvs.KVS.RemoveVarias:input_type -> kvs.ChaveVersao
	3, // 6: kvs.KVS.Snapshot:input_type -> kvs.Versao
	3, // 7: kvs.KVS.Insere:output_type -> kvs.Versao
	0, // 8: kvs.KVS.Consulta:output_type -> kvs.Tupla
	3, // 9: kvs.KVS.Remove:output_type -> kvs.Versao
	3, // 10: kvs.KVS.InsereVarias:output_type -> kvs.Versao
	0, // 11: kvs.KVS.ConsultaVarias:output_type -> kvs.Tupla
	3, // 12: kvs.KVS.RemoveVarias:output_type -> kvs.Versao
	0, // 13: kvs.KVS.Snapshot:output_type -> kvs.Tupla
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kvs_proto_init() }
func file_proto_kvs_proto_init() {
	if File_proto_kvs_proto != nil {
		return
	}
	file_proto_kvs_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_kvs_proto_rawDesc), len(file_proto_kvs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kvs_proto_goTypes,
		DependencyIndexes: file_proto_kvs_proto_depIdxs,
		MessageInfos:      file_proto_kvs_proto_msgTypes,
	}.Build()
	File_proto_kvs_proto = out.File
	file_proto_kvs_proto_goTypes = nil
	file_proto_kvs_proto_depIdxs = nil
}
