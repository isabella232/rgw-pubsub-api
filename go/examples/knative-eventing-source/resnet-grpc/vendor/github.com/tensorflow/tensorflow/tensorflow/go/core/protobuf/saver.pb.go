// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/saver.proto

package protobuf

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

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Deprecated format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Current format: more efficient.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

var SaverDef_CheckpointFormatVersion_name = map[int32]string{
	0: "LEGACY",
	1: "V1",
	2: "V2",
}

var SaverDef_CheckpointFormatVersion_value = map[string]int32{
	"LEGACY": 0,
	"V1":     1,
	"V2":     2,
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return proto.EnumName(SaverDef_CheckpointFormatVersion_name, int32(x))
}

func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5551ea1a7581c104, []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName,proto3" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName,proto3" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName,proto3" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep,proto3" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours,proto3" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,proto3,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                         `json:"-"`
	XXX_unrecognized          []byte                           `json:"-"`
	XXX_sizecache             int32                            `json:"-"`
}

func (m *SaverDef) Reset()         { *m = SaverDef{} }
func (m *SaverDef) String() string { return proto.CompactTextString(m) }
func (*SaverDef) ProtoMessage()    {}
func (*SaverDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_5551ea1a7581c104, []int{0}
}

func (m *SaverDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaverDef.Unmarshal(m, b)
}
func (m *SaverDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaverDef.Marshal(b, m, deterministic)
}
func (m *SaverDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaverDef.Merge(m, src)
}
func (m *SaverDef) XXX_Size() int {
	return xxx_messageInfo_SaverDef.Size(m)
}
func (m *SaverDef) XXX_DiscardUnknown() {
	xxx_messageInfo_SaverDef.DiscardUnknown(m)
}

var xxx_messageInfo_SaverDef proto.InternalMessageInfo

func (m *SaverDef) GetFilenameTensorName() string {
	if m != nil {
		return m.FilenameTensorName
	}
	return ""
}

func (m *SaverDef) GetSaveTensorName() string {
	if m != nil {
		return m.SaveTensorName
	}
	return ""
}

func (m *SaverDef) GetRestoreOpName() string {
	if m != nil {
		return m.RestoreOpName
	}
	return ""
}

func (m *SaverDef) GetMaxToKeep() int32 {
	if m != nil {
		return m.MaxToKeep
	}
	return 0
}

func (m *SaverDef) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if m != nil {
		return m.KeepCheckpointEveryNHours
	}
	return 0
}

func (m *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if m != nil {
		return m.Version
	}
	return SaverDef_LEGACY
}

func init() {
	proto.RegisterEnum("tensorflow.SaverDef_CheckpointFormatVersion", SaverDef_CheckpointFormatVersion_name, SaverDef_CheckpointFormatVersion_value)
	proto.RegisterType((*SaverDef)(nil), "tensorflow.SaverDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/saver.proto", fileDescriptor_5551ea1a7581c104)
}

var fileDescriptor_5551ea1a7581c104 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x8f, 0xd2, 0x40,
	0x14, 0xc6, 0x9d, 0x22, 0x05, 0x86, 0x88, 0xcd, 0x68, 0x62, 0x3d, 0x68, 0x1a, 0x62, 0x4c, 0x0f,
	0xa6, 0x55, 0x8c, 0x37, 0x0f, 0x0a, 0x82, 0x26, 0x1a, 0x24, 0x95, 0x90, 0xb8, 0x97, 0x49, 0x29,
	0xaf, 0xb4, 0x81, 0xf6, 0x35, 0xd3, 0x29, 0xcb, 0xfe, 0x09, 0xfb, 0x1f, 0xef, 0x71, 0x33, 0xc3,
	0x76, 0x81, 0x64, 0xf7, 0xd4, 0xf7, 0xbd, 0xef, 0xf7, 0xf5, 0x25, 0xef, 0x0d, 0x7d, 0x27, 0x21,
	0x2f, 0x51, 0xc4, 0x5b, 0xbc, 0xf4, 0x23, 0x14, 0xe0, 0x17, 0x02, 0x25, 0x2e, 0xab, 0xd8, 0x2f,
	0xc3, 0x1d, 0x08, 0x4f, 0x4b, 0x46, 0x8f, 0x54, 0xff, 0xba, 0x41, 0xdb, 0xff, 0x94, 0xf7, 0x03,
	0x62, 0xf6, 0x91, 0xbe, 0x8c, 0xd3, 0x2d, 0xe4, 0x61, 0x06, 0xfc, 0xc0, 0x70, 0x55, 0xdb, 0xc4,
	0x21, 0x6e, 0x27, 0x60, 0xb5, 0x37, 0xd7, 0xd6, 0x34, 0xcc, 0x80, 0xb9, 0xd4, 0x52, 0x7f, 0x3e,
	0xa3, 0x0d, 0x4d, 0xf7, 0x54, 0xff, 0x84, 0x7c, 0x4f, 0x9f, 0x0b, 0x28, 0x25, 0x0a, 0xe0, 0x58,
	0x1c, 0xc0, 0x86, 0x06, 0x9f, 0xdd, 0xb5, 0xff, 0x16, 0x9a, 0x7b, 0x4b, 0xbb, 0x59, 0xb8, 0xe7,
	0x12, 0xf9, 0x06, 0xa0, 0xb0, 0x9f, 0x3a, 0xc4, 0x6d, 0x06, 0x9d, 0x2c, 0xdc, 0xcf, 0xf1, 0x37,
	0x40, 0xc1, 0x6c, 0xda, 0x2a, 0x93, 0x50, 0xac, 0x60, 0x65, 0x37, 0x1d, 0xe2, 0xb6, 0x83, 0x5a,
	0xb2, 0x6f, 0xf4, 0x8d, 0x8a, 0xf0, 0x28, 0x81, 0x68, 0x53, 0x60, 0x9a, 0x4b, 0x0e, 0x3b, 0x10,
	0x57, 0x3c, 0xe7, 0x09, 0x56, 0xa2, 0xb4, 0x4d, 0x87, 0xb8, 0x46, 0xf0, 0x5a, 0x41, 0xa3, 0x7b,
	0x66, 0xac, 0x90, 0xe9, 0x2f, 0x05, 0xb0, 0x09, 0x6d, 0xed, 0x40, 0x94, 0x29, 0xe6, 0x76, 0xcb,
	0x21, 0x6e, 0x6f, 0xf0, 0xc1, 0x3b, 0xae, 0xca, 0xab, 0xd7, 0xe4, 0x1d, 0xc3, 0x13, 0x14, 0x59,
	0x28, 0x17, 0x87, 0x4c, 0x50, 0x87, 0xfb, 0x5f, 0xe8, 0xab, 0x47, 0x18, 0x46, 0xa9, 0xf9, 0x67,
	0xfc, 0xf3, 0xfb, 0xe8, 0xbf, 0xf5, 0x84, 0x99, 0xd4, 0x58, 0x7c, 0xb2, 0x88, 0xfe, 0x0e, 0x2c,
	0x63, 0x08, 0xf4, 0x05, 0x8a, 0xf5, 0xe9, 0xc8, 0x4a, 0xa6, 0xdb, 0x61, 0x57, 0x0f, 0x9e, 0xa9,
	0xd3, 0x95, 0x33, 0x72, 0xf1, 0x75, 0x9d, 0xca, 0xa4, 0x5a, 0x7a, 0x11, 0x66, 0xfe, 0xc9, 0xb9,
	0x1f, 0x2e, 0xd7, 0x78, 0xfe, 0x0e, 0x6e, 0x08, 0x59, 0x9a, 0x5a, 0x7c, 0xbe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x52, 0xc9, 0x64, 0xa0, 0x2d, 0x02, 0x00, 0x00,
}
