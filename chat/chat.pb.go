// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: chat.proto

package chat

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Retail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda   string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *Retail) Reset() {
	*x = Retail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retail) ProtoMessage() {}

func (x *Retail) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retail.ProtoReflect.Descriptor instead.
func (*Retail) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Retail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Retail) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Retail) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Retail) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Retail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type Pyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor       string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioritario string `protobuf:"bytes,6,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
}

func (x *Pyme) Reset() {
	*x = Pyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pyme) ProtoMessage() {}

func (x *Pyme) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pyme.ProtoReflect.Descriptor instead.
func (*Pyme) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Pyme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pyme) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Pyme) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Pyme) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Pyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Pyme) GetPrioritario() string {
	if x != nil {
		return x.Prioritario
	}
	return ""
}

type Paquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idpaquete   string `protobuf:"bytes,1,opt,name=idpaquete,proto3" json:"idpaquete,omitempty"`
	Seguimiento string `protobuf:"bytes,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    string `protobuf:"bytes,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Paquete) Reset() {
	*x = Paquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paquete) ProtoMessage() {}

func (x *Paquete) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paquete.ProtoReflect.Descriptor instead.
func (*Paquete) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Paquete) GetIdpaquete() string {
	if x != nil {
		return x.Idpaquete
	}
	return ""
}

func (x *Paquete) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *Paquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Paquete) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Paquete) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

func (x *Paquete) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Confirmation) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Tipo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo int32 `protobuf:"varint,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *Tipo) Reset() {
	*x = Tipo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tipo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tipo) ProtoMessage() {}

func (x *Tipo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tipo.ProtoReflect.Descriptor instead.
func (*Tipo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *Tipo) GetTipo() int32 {
	if x != nil {
		return x.Tipo
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x7c, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x22, 0x9c, 0x01, 0x0a, 0x04, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65,
	0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x22,
	0xa7, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67,
	0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1a, 0x0a,
	0x04, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x32, 0xcf, 0x01, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x0a, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x50, 0x79, 0x6d, 0x65, 0x1a, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x0a,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x69, 0x70, 0x6f, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x63, 0x68, 0x63, 0x6c, 0x2f,
	0x73, 0x64, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_proto_goTypes = []interface{}{
	(*Retail)(nil),       // 0: chat.Retail
	(*Pyme)(nil),         // 1: chat.Pyme
	(*Paquete)(nil),      // 2: chat.Paquete
	(*Confirmation)(nil), // 3: chat.Confirmation
	(*Tipo)(nil),         // 4: chat.Tipo
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.OrderRetail:input_type -> chat.Retail
	1, // 1: chat.ChatService.OrderPyme:input_type -> chat.Pyme
	3, // 2: chat.ChatService.Seguimiento:input_type -> chat.Confirmation
	4, // 3: chat.ChatService.Camion:input_type -> chat.Tipo
	3, // 4: chat.ChatService.OrderRetail:output_type -> chat.Confirmation
	3, // 5: chat.ChatService.OrderPyme:output_type -> chat.Confirmation
	3, // 6: chat.ChatService.Seguimiento:output_type -> chat.Confirmation
	2, // 7: chat.ChatService.Camion:output_type -> chat.Paquete
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retail); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pyme); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paquete); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tipo); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	OrderRetail(ctx context.Context, in *Retail, opts ...grpc.CallOption) (*Confirmation, error)
	OrderPyme(ctx context.Context, in *Pyme, opts ...grpc.CallOption) (*Confirmation, error)
	Seguimiento(ctx context.Context, in *Confirmation, opts ...grpc.CallOption) (*Confirmation, error)
	Camion(ctx context.Context, in *Tipo, opts ...grpc.CallOption) (*Paquete, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) OrderRetail(ctx context.Context, in *Retail, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/chat.ChatService/OrderRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) OrderPyme(ctx context.Context, in *Pyme, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/chat.ChatService/OrderPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Seguimiento(ctx context.Context, in *Confirmation, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Seguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Camion(ctx context.Context, in *Tipo, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Camion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	OrderRetail(context.Context, *Retail) (*Confirmation, error)
	OrderPyme(context.Context, *Pyme) (*Confirmation, error)
	Seguimiento(context.Context, *Confirmation) (*Confirmation, error)
	Camion(context.Context, *Tipo) (*Paquete, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) OrderRetail(context.Context, *Retail) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderRetail not implemented")
}
func (*UnimplementedChatServiceServer) OrderPyme(context.Context, *Pyme) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPyme not implemented")
}
func (*UnimplementedChatServiceServer) Seguimiento(context.Context, *Confirmation) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seguimiento not implemented")
}
func (*UnimplementedChatServiceServer) Camion(context.Context, *Tipo) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Camion not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_OrderRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Retail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).OrderRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/OrderRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).OrderRetail(ctx, req.(*Retail))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_OrderPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).OrderPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/OrderPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).OrderPyme(ctx, req.(*Pyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Seguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Confirmation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Seguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Seguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Seguimiento(ctx, req.(*Confirmation))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Camion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tipo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Camion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Camion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Camion(ctx, req.(*Tipo))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderRetail",
			Handler:    _ChatService_OrderRetail_Handler,
		},
		{
			MethodName: "OrderPyme",
			Handler:    _ChatService_OrderPyme_Handler,
		},
		{
			MethodName: "Seguimiento",
			Handler:    _ChatService_Seguimiento_Handler,
		},
		{
			MethodName: "Camion",
			Handler:    _ChatService_Camion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
