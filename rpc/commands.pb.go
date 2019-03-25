// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration contains information to instantiate an Arduino Platform Service
type Configuration struct {
	// dataDir represents the current root of the arduino tree (defaulted to
	// `$HOME/.arduino15` on linux).
	DataDir string `protobuf:"bytes,1,opt,name=dataDir,proto3" json:"dataDir,omitempty"`
	// sketchbookDir represents the current root of the sketchbooks tree
	// (defaulted to `$HOME/Arduino`).
	SketchbookDir string `protobuf:"bytes,2,opt,name=sketchbookDir,proto3" json:"sketchbookDir,omitempty"`
	// ArduinoIDEDirectory is the directory of the Arduino IDE if the CLI runs
	// together with it.
	DownloadsDir string `protobuf:"bytes,3,opt,name=downloadsDir,proto3" json:"downloadsDir,omitempty"`
	// BoardManagerAdditionalUrls contains the additional URL for 3rd party
	// packages
	BoardManagerAdditionalUrls []string `protobuf:"bytes,4,rep,name=boardManagerAdditionalUrls,proto3" json:"boardManagerAdditionalUrls,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_ecd59741834cc893, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

func (m *Configuration) GetSketchbookDir() string {
	if m != nil {
		return m.SketchbookDir
	}
	return ""
}

func (m *Configuration) GetDownloadsDir() string {
	if m != nil {
		return m.DownloadsDir
	}
	return ""
}

func (m *Configuration) GetBoardManagerAdditionalUrls() []string {
	if m != nil {
		return m.BoardManagerAdditionalUrls
	}
	return nil
}

type InitReq struct {
	Configuration        *Configuration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	LibraryManagerOnly   bool           `protobuf:"varint,2,opt,name=library_manager_only,json=libraryManagerOnly,proto3" json:"library_manager_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitReq) Reset()         { *m = InitReq{} }
func (m *InitReq) String() string { return proto.CompactTextString(m) }
func (*InitReq) ProtoMessage()    {}
func (*InitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_ecd59741834cc893, []int{1}
}
func (m *InitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitReq.Unmarshal(m, b)
}
func (m *InitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitReq.Marshal(b, m, deterministic)
}
func (dst *InitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitReq.Merge(dst, src)
}
func (m *InitReq) XXX_Size() int {
	return xxx_messageInfo_InitReq.Size(m)
}
func (m *InitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitReq proto.InternalMessageInfo

func (m *InitReq) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *InitReq) GetLibraryManagerOnly() bool {
	if m != nil {
		return m.LibraryManagerOnly
	}
	return false
}

type InitResp struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Result               *Result   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InitResp) Reset()         { *m = InitResp{} }
func (m *InitResp) String() string { return proto.CompactTextString(m) }
func (*InitResp) ProtoMessage()    {}
func (*InitResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_ecd59741834cc893, []int{2}
}
func (m *InitResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResp.Unmarshal(m, b)
}
func (m *InitResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResp.Marshal(b, m, deterministic)
}
func (dst *InitResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResp.Merge(dst, src)
}
func (m *InitResp) XXX_Size() int {
	return xxx_messageInfo_InitResp.Size(m)
}
func (m *InitResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResp.DiscardUnknown(m)
}

var xxx_messageInfo_InitResp proto.InternalMessageInfo

func (m *InitResp) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *InitResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type DestroyReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DestroyReq) Reset()         { *m = DestroyReq{} }
func (m *DestroyReq) String() string { return proto.CompactTextString(m) }
func (*DestroyReq) ProtoMessage()    {}
func (*DestroyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_ecd59741834cc893, []int{3}
}
func (m *DestroyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyReq.Unmarshal(m, b)
}
func (m *DestroyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyReq.Marshal(b, m, deterministic)
}
func (dst *DestroyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyReq.Merge(dst, src)
}
func (m *DestroyReq) XXX_Size() int {
	return xxx_messageInfo_DestroyReq.Size(m)
}
func (m *DestroyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyReq proto.InternalMessageInfo

func (m *DestroyReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type DestroyResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyResp) Reset()         { *m = DestroyResp{} }
func (m *DestroyResp) String() string { return proto.CompactTextString(m) }
func (*DestroyResp) ProtoMessage()    {}
func (*DestroyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_ecd59741834cc893, []int{4}
}
func (m *DestroyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyResp.Unmarshal(m, b)
}
func (m *DestroyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyResp.Marshal(b, m, deterministic)
}
func (dst *DestroyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyResp.Merge(dst, src)
}
func (m *DestroyResp) XXX_Size() int {
	return xxx_messageInfo_DestroyResp.Size(m)
}
func (m *DestroyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyResp proto.InternalMessageInfo

func (m *DestroyResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Configuration)(nil), "arduino.Configuration")
	proto.RegisterType((*InitReq)(nil), "arduino.InitReq")
	proto.RegisterType((*InitResp)(nil), "arduino.InitResp")
	proto.RegisterType((*DestroyReq)(nil), "arduino.DestroyReq")
	proto.RegisterType((*DestroyResp)(nil), "arduino.DestroyResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArduinoCoreClient is the client API for ArduinoCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArduinoCoreClient interface {
	// Start a new instance of the Arduino Core Service
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitResp, error)
	// Destroy an instance of the Arduino Core Service
	Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyResp, error)
	// Requests details about a board
	BoardDetails(ctx context.Context, in *BoardDetailsReq, opts ...grpc.CallOption) (*BoardDetailsResp, error)
	Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (*CompileResp, error)
}

type arduinoCoreClient struct {
	cc *grpc.ClientConn
}

func NewArduinoCoreClient(cc *grpc.ClientConn) ArduinoCoreClient {
	return &arduinoCoreClient{cc}
}

func (c *arduinoCoreClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitResp, error) {
	out := new(InitResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyResp, error) {
	out := new(DestroyResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) BoardDetails(ctx context.Context, in *BoardDetailsReq, opts ...grpc.CallOption) (*BoardDetailsResp, error) {
	out := new(BoardDetailsResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/BoardDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (*CompileResp, error) {
	out := new(CompileResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArduinoCoreServer is the server API for ArduinoCore service.
type ArduinoCoreServer interface {
	// Start a new instance of the Arduino Core Service
	Init(context.Context, *InitReq) (*InitResp, error)
	// Destroy an instance of the Arduino Core Service
	Destroy(context.Context, *DestroyReq) (*DestroyResp, error)
	// Requests details about a board
	BoardDetails(context.Context, *BoardDetailsReq) (*BoardDetailsResp, error)
	Compile(context.Context, *CompileReq) (*CompileResp, error)
}

func RegisterArduinoCoreServer(s *grpc.Server, srv ArduinoCoreServer) {
	s.RegisterService(&_ArduinoCore_serviceDesc, srv)
}

func _ArduinoCore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Destroy(ctx, req.(*DestroyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_BoardDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).BoardDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/BoardDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).BoardDetails(ctx, req.(*BoardDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Compile(ctx, req.(*CompileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArduinoCore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arduino.ArduinoCore",
	HandlerType: (*ArduinoCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ArduinoCore_Init_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _ArduinoCore_Destroy_Handler,
		},
		{
			MethodName: "BoardDetails",
			Handler:    _ArduinoCore_BoardDetails_Handler,
		},
		{
			MethodName: "Compile",
			Handler:    _ArduinoCore_Compile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands.proto",
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_commands_ecd59741834cc893) }

var fileDescriptor_commands_ecd59741834cc893 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x09, 0x9b, 0x96, 0xf2, 0x4f, 0x0b, 0xcc, 0x4c, 0xa8, 0xe4, 0x34, 0x59, 0x93, 0xe0,
	0xb2, 0x0c, 0x15, 0xb4, 0x0b, 0x08, 0x69, 0x4b, 0x2f, 0x3b, 0x20, 0xa4, 0x48, 0x5c, 0xb8, 0x4c,
	0x4e, 0x62, 0x3a, 0x6b, 0xae, 0x1d, 0x6c, 0x47, 0x28, 0x9f, 0x8b, 0x0f, 0xc6, 0x57, 0xc0, 0x76,
	0xbc, 0x2c, 0x01, 0x51, 0x69, 0xa7, 0xca, 0xef, 0x3d, 0xfb, 0xf7, 0x7f, 0x6e, 0x0c, 0x4f, 0x2b,
	0xb9, 0xdd, 0x12, 0x51, 0xeb, 0xac, 0x51, 0xd2, 0x48, 0x14, 0x13, 0x55, 0xb7, 0x4c, 0xc8, 0x74,
	0xee, 0x0c, 0x29, 0x7a, 0x39, 0x4d, 0x4a, 0x69, 0x8d, 0xb0, 0x58, 0x58, 0xab, 0x61, 0x9c, 0xf6,
	0x4b, 0xfc, 0x2b, 0x82, 0x45, 0x2e, 0xc5, 0x77, 0xb6, 0x69, 0x15, 0x31, 0x4c, 0x0a, 0xb4, 0x84,
	0xb8, 0x26, 0x86, 0xac, 0x99, 0x5a, 0x46, 0xc7, 0xd1, 0x9b, 0x27, 0xc5, 0xdd, 0x12, 0x9d, 0xc0,
	0x42, 0xdf, 0x52, 0x53, 0xdd, 0x94, 0x52, 0xde, 0x3a, 0xff, 0xb1, 0xf7, 0xa7, 0x22, 0xc2, 0x30,
	0xaf, 0xe5, 0x4f, 0xc1, 0x25, 0xa9, 0xb5, 0x0b, 0xed, 0xf9, 0xd0, 0x44, 0x43, 0x9f, 0x20, 0xf5,
	0x33, 0x7d, 0x26, 0x82, 0x6c, 0xa8, 0xba, 0xa8, 0x6b, 0xe6, 0xd8, 0x84, 0x7f, 0x55, 0x5c, 0x2f,
	0xf7, 0x8f, 0xf7, 0xec, 0x8e, 0x1d, 0x09, 0xdc, 0x41, 0x7c, 0x25, 0x98, 0x29, 0xe8, 0x0f, 0xf4,
	0x11, 0x6c, 0xa3, 0xd1, 0xfc, 0x7e, 0xe8, 0x64, 0xf5, 0x32, 0x0b, 0x77, 0x91, 0x4d, 0xda, 0x15,
	0xd3, 0x30, 0x7a, 0x0b, 0x47, 0x9c, 0x95, 0x8a, 0xa8, 0xee, 0x7a, 0xdb, 0x93, 0xae, 0xa5, 0xe0,
	0x9d, 0x6f, 0x36, 0x2b, 0x50, 0xf0, 0xc2, 0x10, 0x5f, 0xac, 0x83, 0x4b, 0x98, 0xf5, 0x68, 0xdd,
	0xa0, 0x53, 0x98, 0x31, 0xa1, 0x0d, 0x11, 0x15, 0x0d, 0xd8, 0xc3, 0x01, 0x7b, 0x15, 0x8c, 0x62,
	0x88, 0xa0, 0xd7, 0x70, 0xa0, 0xa8, 0x6e, 0xb9, 0xf1, 0xc7, 0x27, 0xab, 0x67, 0x43, 0xb8, 0xf0,
	0x72, 0x11, 0x6c, 0xfc, 0x01, 0x60, 0x4d, 0xb5, 0x51, 0xb2, 0x73, 0x0d, 0x1f, 0x46, 0xc1, 0xe7,
	0x90, 0x0c, 0x9b, 0xed, 0x8c, 0xf7, 0xd0, 0x68, 0x27, 0x74, 0xf5, 0x3b, 0x82, 0xe4, 0xa2, 0xb7,
	0x72, 0xa9, 0xa8, 0xc5, 0xee, 0xbb, 0xa2, 0xe8, 0xf9, 0x08, 0xe6, 0xaf, 0x3c, 0x3d, 0xfc, 0x4b,
	0xd1, 0x0d, 0x7e, 0x84, 0xce, 0x21, 0x0e, 0x58, 0xf4, 0x62, 0xf0, 0xef, 0x5b, 0xa4, 0x47, 0xff,
	0x8a, 0x7e, 0x5f, 0x0e, 0xf3, 0x4b, 0xf7, 0x47, 0xaf, 0xa9, 0x21, 0x8c, 0x6b, 0xb4, 0x1c, 0x72,
	0x63, 0xd9, 0x9d, 0xf0, 0xea, 0x3f, 0x8e, 0x2d, 0xf9, 0x1e, 0xe2, 0xbc, 0xff, 0xac, 0x47, 0xf0,
	0xa0, 0x4c, 0xe1, 0x83, 0xa8, 0x9b, 0xcb, 0x93, 0x6f, 0x78, 0xc3, 0xcc, 0x4d, 0x5b, 0x66, 0xf6,
	0x4d, 0x9c, 0x85, 0xc4, 0xdd, 0xef, 0x69, 0xc5, 0xd9, 0x99, 0x6a, 0xaa, 0xf2, 0xc0, 0x3f, 0x94,
	0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x12, 0x35, 0x77, 0x6d, 0x03, 0x00, 0x00,
}