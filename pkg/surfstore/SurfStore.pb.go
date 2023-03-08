// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/surfstore/SurfStore.proto

package surfstore

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BlockHash) Reset() {
	*x = BlockHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHash) ProtoMessage() {}

func (x *BlockHash) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHash.ProtoReflect.Descriptor instead.
func (*BlockHash) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHash) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type BlockHashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []string `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *BlockHashes) Reset() {
	*x = BlockHashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHashes) ProtoMessage() {}

func (x *BlockHashes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHashes.ProtoReflect.Descriptor instead.
func (*BlockHashes) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{1}
}

func (x *BlockHashes) GetHashes() []string {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockData []byte `protobuf:"bytes,1,opt,name=blockData,proto3" json:"blockData,omitempty"`
	BlockSize int32  `protobuf:"varint,2,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetBlockData() []byte {
	if x != nil {
		return x.BlockData
	}
	return nil
}

func (x *Block) GetBlockSize() int32 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

type FileMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename      string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Version       int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	BlockHashList []string `protobuf:"bytes,3,rep,name=blockHashList,proto3" json:"blockHashList,omitempty"`
}

func (x *FileMetaData) Reset() {
	*x = FileMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetaData) ProtoMessage() {}

func (x *FileMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetaData.ProtoReflect.Descriptor instead.
func (*FileMetaData) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{4}
}

func (x *FileMetaData) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileMetaData) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *FileMetaData) GetBlockHashList() []string {
	if x != nil {
		return x.BlockHashList
	}
	return nil
}

type FileInfoMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileInfoMap map[string]*FileMetaData `protobuf:"bytes,1,rep,name=fileInfoMap,proto3" json:"fileInfoMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileInfoMap) Reset() {
	*x = FileInfoMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoMap) ProtoMessage() {}

func (x *FileInfoMap) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoMap.ProtoReflect.Descriptor instead.
func (*FileInfoMap) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfoMap) GetFileInfoMap() map[string]*FileMetaData {
	if x != nil {
		return x.FileInfoMap
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{6}
}

func (x *Version) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type BlockStoreMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockStoreMap map[string]*BlockHashes `protobuf:"bytes,1,rep,name=blockStoreMap,proto3" json:"blockStoreMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BlockStoreMap) Reset() {
	*x = BlockStoreMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStoreMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStoreMap) ProtoMessage() {}

func (x *BlockStoreMap) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStoreMap.ProtoReflect.Descriptor instead.
func (*BlockStoreMap) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{7}
}

func (x *BlockStoreMap) GetBlockStoreMap() map[string]*BlockHashes {
	if x != nil {
		return x.BlockStoreMap
	}
	return nil
}

type BlockStoreAddrs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockStoreAddrs []string `protobuf:"bytes,1,rep,name=blockStoreAddrs,proto3" json:"blockStoreAddrs,omitempty"`
}

func (x *BlockStoreAddrs) Reset() {
	*x = BlockStoreAddrs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStoreAddrs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStoreAddrs) ProtoMessage() {}

func (x *BlockStoreAddrs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_surfstore_SurfStore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStoreAddrs.ProtoReflect.Descriptor instead.
func (*BlockStoreAddrs) Descriptor() ([]byte, []int) {
	return file_pkg_surfstore_SurfStore_proto_rawDescGZIP(), []int{8}
}

func (x *BlockStoreAddrs) GetBlockStoreAddrs() []string {
	if x != nil {
		return x.BlockStoreAddrs
	}
	return nil
}

var File_pkg_surfstore_SurfStore_proto protoreflect.FileDescriptor

var file_pkg_surfstore_SurfStore_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x53, 0x75, 0x72, 0x66, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x25, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22,
	0x43, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66,
	0x6c, 0x61, 0x67, 0x22, 0x6a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0xb1, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x12,
	0x49, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x57, 0x0a, 0x10, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x51, 0x0a, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x58, 0x0a,
	0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3b, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x73, 0x32, 0xf9, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x50, 0x75, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x12, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x09, 0x48, 0x61, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x73, 0x75, 0x72,
	0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x00,
	0x32, 0xa0, 0x02, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x42,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x73, 0x75, 0x72, 0x66,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x16, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x1a, 0x18, 0x2e, 0x73, 0x75,
	0x72, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x73, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x73, 0x65, 0x32, 0x32, 0x34, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x34, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_surfstore_SurfStore_proto_rawDescOnce sync.Once
	file_pkg_surfstore_SurfStore_proto_rawDescData = file_pkg_surfstore_SurfStore_proto_rawDesc
)

func file_pkg_surfstore_SurfStore_proto_rawDescGZIP() []byte {
	file_pkg_surfstore_SurfStore_proto_rawDescOnce.Do(func() {
		file_pkg_surfstore_SurfStore_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_surfstore_SurfStore_proto_rawDescData)
	})
	return file_pkg_surfstore_SurfStore_proto_rawDescData
}

var file_pkg_surfstore_SurfStore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_surfstore_SurfStore_proto_goTypes = []interface{}{
	(*BlockHash)(nil),       // 0: surfstore.BlockHash
	(*BlockHashes)(nil),     // 1: surfstore.BlockHashes
	(*Block)(nil),           // 2: surfstore.Block
	(*Success)(nil),         // 3: surfstore.Success
	(*FileMetaData)(nil),    // 4: surfstore.FileMetaData
	(*FileInfoMap)(nil),     // 5: surfstore.FileInfoMap
	(*Version)(nil),         // 6: surfstore.Version
	(*BlockStoreMap)(nil),   // 7: surfstore.BlockStoreMap
	(*BlockStoreAddrs)(nil), // 8: surfstore.BlockStoreAddrs
	nil,                     // 9: surfstore.FileInfoMap.FileInfoMapEntry
	nil,                     // 10: surfstore.BlockStoreMap.BlockStoreMapEntry
	(*emptypb.Empty)(nil),   // 11: google.protobuf.Empty
}
var file_pkg_surfstore_SurfStore_proto_depIdxs = []int32{
	9,  // 0: surfstore.FileInfoMap.fileInfoMap:type_name -> surfstore.FileInfoMap.FileInfoMapEntry
	10, // 1: surfstore.BlockStoreMap.blockStoreMap:type_name -> surfstore.BlockStoreMap.BlockStoreMapEntry
	4,  // 2: surfstore.FileInfoMap.FileInfoMapEntry.value:type_name -> surfstore.FileMetaData
	1,  // 3: surfstore.BlockStoreMap.BlockStoreMapEntry.value:type_name -> surfstore.BlockHashes
	0,  // 4: surfstore.BlockStore.GetBlock:input_type -> surfstore.BlockHash
	2,  // 5: surfstore.BlockStore.PutBlock:input_type -> surfstore.Block
	1,  // 6: surfstore.BlockStore.HasBlocks:input_type -> surfstore.BlockHashes
	11, // 7: surfstore.BlockStore.GetBlockHashes:input_type -> google.protobuf.Empty
	11, // 8: surfstore.MetaStore.GetFileInfoMap:input_type -> google.protobuf.Empty
	4,  // 9: surfstore.MetaStore.UpdateFile:input_type -> surfstore.FileMetaData
	1,  // 10: surfstore.MetaStore.GetBlockStoreMap:input_type -> surfstore.BlockHashes
	11, // 11: surfstore.MetaStore.GetBlockStoreAddrs:input_type -> google.protobuf.Empty
	2,  // 12: surfstore.BlockStore.GetBlock:output_type -> surfstore.Block
	3,  // 13: surfstore.BlockStore.PutBlock:output_type -> surfstore.Success
	1,  // 14: surfstore.BlockStore.HasBlocks:output_type -> surfstore.BlockHashes
	1,  // 15: surfstore.BlockStore.GetBlockHashes:output_type -> surfstore.BlockHashes
	5,  // 16: surfstore.MetaStore.GetFileInfoMap:output_type -> surfstore.FileInfoMap
	6,  // 17: surfstore.MetaStore.UpdateFile:output_type -> surfstore.Version
	7,  // 18: surfstore.MetaStore.GetBlockStoreMap:output_type -> surfstore.BlockStoreMap
	8,  // 19: surfstore.MetaStore.GetBlockStoreAddrs:output_type -> surfstore.BlockStoreAddrs
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_surfstore_SurfStore_proto_init() }
func file_pkg_surfstore_SurfStore_proto_init() {
	if File_pkg_surfstore_SurfStore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_surfstore_SurfStore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHash); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHashes); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetaData); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoMap); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStoreMap); i {
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
		file_pkg_surfstore_SurfStore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStoreAddrs); i {
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
			RawDescriptor: file_pkg_surfstore_SurfStore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_surfstore_SurfStore_proto_goTypes,
		DependencyIndexes: file_pkg_surfstore_SurfStore_proto_depIdxs,
		MessageInfos:      file_pkg_surfstore_SurfStore_proto_msgTypes,
	}.Build()
	File_pkg_surfstore_SurfStore_proto = out.File
	file_pkg_surfstore_SurfStore_proto_rawDesc = nil
	file_pkg_surfstore_SurfStore_proto_goTypes = nil
	file_pkg_surfstore_SurfStore_proto_depIdxs = nil
}
