// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: SqrtService.proto

package api

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

type Args struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	C int64 `protobuf:"varint,3,opt,name=C,proto3" json:"C,omitempty"`
}

func (x *Args) Reset() {
	*x = Args{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SqrtService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Args) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Args) ProtoMessage() {}

func (x *Args) ProtoReflect() protoreflect.Message {
	mi := &file_SqrtService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Args.ProtoReflect.Descriptor instead.
func (*Args) Descriptor() ([]byte, []int) {
	return file_SqrtService_proto_rawDescGZIP(), []int{0}
}

func (x *Args) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Args) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Args) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []float64 `protobuf:"fixed64,1,rep,packed,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SqrtService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_SqrtService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_SqrtService_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetResult() []float64 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_SqrtService_proto protoreflect.FileDescriptor

var file_SqrtService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x71, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x30, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x41, 0x12, 0x0c,
	0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x42, 0x12, 0x0c, 0x0a, 0x01,
	0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x43, 0x22, 0x1f, 0x0a, 0x05, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x2e, 0x0a, 0x0b, 0x53,
	0x71, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x53, 0x71,
	0x72, 0x74, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SqrtService_proto_rawDescOnce sync.Once
	file_SqrtService_proto_rawDescData = file_SqrtService_proto_rawDesc
)

func file_SqrtService_proto_rawDescGZIP() []byte {
	file_SqrtService_proto_rawDescOnce.Do(func() {
		file_SqrtService_proto_rawDescData = protoimpl.X.CompressGZIP(file_SqrtService_proto_rawDescData)
	})
	return file_SqrtService_proto_rawDescData
}

var file_SqrtService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SqrtService_proto_goTypes = []interface{}{
	(*Args)(nil),  // 0: api.Args
	(*Reply)(nil), // 1: api.Reply
}
var file_SqrtService_proto_depIdxs = []int32{
	0, // 0: api.SqrtService.Sqrt:input_type -> api.Args
	1, // 1: api.SqrtService.Sqrt:output_type -> api.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SqrtService_proto_init() }
func file_SqrtService_proto_init() {
	if File_SqrtService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SqrtService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Args); i {
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
		file_SqrtService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_SqrtService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_SqrtService_proto_goTypes,
		DependencyIndexes: file_SqrtService_proto_depIdxs,
		MessageInfos:      file_SqrtService_proto_msgTypes,
	}.Build()
	File_SqrtService_proto = out.File
	file_SqrtService_proto_rawDesc = nil
	file_SqrtService_proto_goTypes = nil
	file_SqrtService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SqrtServiceClient is the client API for SqrtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqrtServiceClient interface {
	Sqrt(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error)
}

type sqrtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqrtServiceClient(cc grpc.ClientConnInterface) SqrtServiceClient {
	return &sqrtServiceClient{cc}
}

func (c *sqrtServiceClient) Sqrt(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/api.SqrtService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqrtServiceServer is the server API for SqrtService service.
type SqrtServiceServer interface {
	Sqrt(context.Context, *Args) (*Reply, error)
}

// UnimplementedSqrtServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqrtServiceServer struct {
}

func (*UnimplementedSqrtServiceServer) Sqrt(context.Context, *Args) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}

func RegisterSqrtServiceServer(s *grpc.Server, srv SqrtServiceServer) {
	s.RegisterService(&_SqrtService_serviceDesc, srv)
}

func _SqrtService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Args)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqrtServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SqrtService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqrtServiceServer).Sqrt(ctx, req.(*Args))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqrtService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SqrtService",
	HandlerType: (*SqrtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sqrt",
			Handler:    _SqrtService_Sqrt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SqrtService.proto",
}
