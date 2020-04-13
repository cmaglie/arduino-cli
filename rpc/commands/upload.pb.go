// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/upload.proto

package commands

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UploadReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"` // Deprecated: Do not use.
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketch_path,json=sketchPath,proto3" json:"sketch_path,omitempty"`
	Port                 string    `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Verbose              bool      `protobuf:"varint,5,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Verify               bool      `protobuf:"varint,6,opt,name=verify,proto3" json:"verify,omitempty"`
	ImportFile           string    `protobuf:"bytes,7,opt,name=import_file,json=importFile,proto3" json:"import_file,omitempty"`
	Board                string    `protobuf:"bytes,8,opt,name=board,proto3" json:"board,omitempty"`
	BoardConfig          []string  `protobuf:"bytes,9,rep,name=board_config,json=boardConfig,proto3" json:"board_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UploadReq) Reset()         { *m = UploadReq{} }
func (m *UploadReq) String() string { return proto.CompactTextString(m) }
func (*UploadReq) ProtoMessage()    {}
func (*UploadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd642cc079f8acdb, []int{0}
}

func (m *UploadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadReq.Unmarshal(m, b)
}
func (m *UploadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadReq.Marshal(b, m, deterministic)
}
func (m *UploadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadReq.Merge(m, src)
}
func (m *UploadReq) XXX_Size() int {
	return xxx_messageInfo_UploadReq.Size(m)
}
func (m *UploadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadReq proto.InternalMessageInfo

func (m *UploadReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

// Deprecated: Do not use.
func (m *UploadReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *UploadReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *UploadReq) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *UploadReq) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *UploadReq) GetVerify() bool {
	if m != nil {
		return m.Verify
	}
	return false
}

func (m *UploadReq) GetImportFile() string {
	if m != nil {
		return m.ImportFile
	}
	return ""
}

func (m *UploadReq) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *UploadReq) GetBoardConfig() []string {
	if m != nil {
		return m.BoardConfig
	}
	return nil
}

type UploadResp struct {
	OutStream            []byte   `protobuf:"bytes,1,opt,name=out_stream,json=outStream,proto3" json:"out_stream,omitempty"`
	ErrStream            []byte   `protobuf:"bytes,2,opt,name=err_stream,json=errStream,proto3" json:"err_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResp) Reset()         { *m = UploadResp{} }
func (m *UploadResp) String() string { return proto.CompactTextString(m) }
func (*UploadResp) ProtoMessage()    {}
func (*UploadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd642cc079f8acdb, []int{1}
}

func (m *UploadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResp.Unmarshal(m, b)
}
func (m *UploadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResp.Marshal(b, m, deterministic)
}
func (m *UploadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResp.Merge(m, src)
}
func (m *UploadResp) XXX_Size() int {
	return xxx_messageInfo_UploadResp.Size(m)
}
func (m *UploadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResp.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResp proto.InternalMessageInfo

func (m *UploadResp) GetOutStream() []byte {
	if m != nil {
		return m.OutStream
	}
	return nil
}

func (m *UploadResp) GetErrStream() []byte {
	if m != nil {
		return m.ErrStream
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadReq)(nil), "cc.arduino.cli.commands.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "cc.arduino.cli.commands.UploadResp")
}

func init() { proto.RegisterFile("commands/upload.proto", fileDescriptor_cd642cc079f8acdb) }

var fileDescriptor_cd642cc079f8acdb = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0x69, 0xe7, 0x6f, 0xef, 0xcc, 0x2a, 0x7c, 0xdf, 0x18, 0x04, 0xb1, 0x33, 0xab, 0x82,
	0x4c, 0x0b, 0xba, 0x76, 0x33, 0x82, 0xa0, 0x2b, 0x89, 0xb8, 0x71, 0x53, 0xd2, 0x34, 0x9d, 0x06,
	0xdb, 0xa4, 0x93, 0xa6, 0x03, 0x3e, 0x8f, 0x2f, 0x2a, 0x4d, 0x5a, 0x5d, 0xb9, 0xea, 0x3d, 0xe7,
	0x77, 0x72, 0x2f, 0x9c, 0xc2, 0x7f, 0xa6, 0xea, 0x9a, 0xca, 0xbc, 0x4d, 0xba, 0xa6, 0x52, 0x34,
	0x8f, 0x1b, 0xad, 0x8c, 0x42, 0x17, 0x8c, 0xc5, 0x54, 0xe7, 0x9d, 0x90, 0x2a, 0x66, 0x95, 0x88,
	0xc7, 0xd4, 0xe5, 0x6f, 0xbe, 0x1f, 0x94, 0x74, 0xf9, 0xdd, 0x97, 0x0f, 0xc1, 0x9b, 0x5d, 0x40,
	0xf8, 0x09, 0xdd, 0xc3, 0x52, 0xc8, 0xd6, 0x50, 0xc9, 0x38, 0xf6, 0x42, 0x2f, 0x5a, 0xdd, 0x6e,
	0xe3, 0x3f, 0x16, 0xc6, 0x4f, 0x43, 0x90, 0xfc, 0x3c, 0x41, 0x1b, 0x98, 0x16, 0xa7, 0x4c, 0x62,
	0x3f, 0xf4, 0xa2, 0xe0, 0xe0, 0x63, 0x8f, 0x58, 0x8d, 0xae, 0x61, 0xd5, 0x7e, 0x70, 0xc3, 0xca,
	0xb4, 0xa1, 0xa6, 0xc4, 0x93, 0x1e, 0x13, 0x70, 0xd6, 0x0b, 0x35, 0x25, 0x42, 0x30, 0x6d, 0x94,
	0x36, 0x78, 0x6a, 0x89, 0x9d, 0x11, 0x86, 0xc5, 0x99, 0xeb, 0x4c, 0xb5, 0x1c, 0xcf, 0x42, 0x2f,
	0x5a, 0x92, 0x51, 0xa2, 0x0d, 0xcc, 0xcf, 0x5c, 0x8b, 0xe2, 0x13, 0xcf, 0x2d, 0x18, 0x54, 0x7f,
	0x46, 0xd4, 0xfd, 0xdb, 0xb4, 0x10, 0x15, 0xc7, 0x0b, 0x77, 0xc6, 0x59, 0x8f, 0xa2, 0xe2, 0xe8,
	0x1f, 0xcc, 0x32, 0x45, 0x75, 0x8e, 0x97, 0x16, 0x39, 0x81, 0xb6, 0xb0, 0xb6, 0x43, 0xca, 0x94,
	0x2c, 0xc4, 0x11, 0x07, 0xe1, 0x24, 0x0a, 0xc8, 0xca, 0x7a, 0x0f, 0xd6, 0xda, 0x3d, 0x03, 0x8c,
	0x25, 0xb5, 0x0d, 0xba, 0x02, 0x50, 0x9d, 0x49, 0x5b, 0xa3, 0x39, 0xad, 0x6d, 0x4f, 0x6b, 0x12,
	0xa8, 0xce, 0xbc, 0x5a, 0xa3, 0xc7, 0x5c, 0xeb, 0x11, 0xfb, 0x0e, 0x73, 0xad, 0x1d, 0x3e, 0xec,
	0xdf, 0x6f, 0x8e, 0xc2, 0x94, 0x5d, 0xd6, 0x57, 0x99, 0x0c, 0xd5, 0x8e, 0xdf, 0x3d, 0xab, 0x44,
	0xa2, 0x1b, 0x96, 0x8c, 0x35, 0x67, 0x73, 0xfb, 0x9f, 0xee, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x22, 0x5f, 0xe2, 0xf0, 0x01, 0x00, 0x00,
}
