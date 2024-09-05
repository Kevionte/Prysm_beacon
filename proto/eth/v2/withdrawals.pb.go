// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/eth/v2/withdrawals.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	github_com_prysmaticlabs_prysm_v5_consensus_types_primitives "github.com/Kevionte/prysm_beacon/v5/consensus-types/primitives"
	_ "github.com/Kevionte/prysm_beacon/v5/proto/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BLSToExecutionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorIndex     github_com_prysmaticlabs_prysm_v5_consensus_types_primitives.ValidatorIndex `protobuf:"varint,1,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty" cast-type:"github.com/Kevionte/prysm_beacon/v5/consensus-types/primitives.ValidatorIndex"`
	FromBlsPubkey      []byte                                                                      `protobuf:"bytes,2,opt,name=from_bls_pubkey,json=fromBlsPubkey,proto3" json:"from_bls_pubkey,omitempty" ssz-size:"48"`
	ToExecutionAddress []byte                                                                      `protobuf:"bytes,3,opt,name=to_execution_address,json=toExecutionAddress,proto3" json:"to_execution_address,omitempty" ssz-size:"20"`
}

func (x *BLSToExecutionChange) Reset() {
	*x = BLSToExecutionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_withdrawals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLSToExecutionChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLSToExecutionChange) ProtoMessage() {}

func (x *BLSToExecutionChange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_withdrawals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLSToExecutionChange.ProtoReflect.Descriptor instead.
func (*BLSToExecutionChange) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_withdrawals_proto_rawDescGZIP(), []int{0}
}

func (x *BLSToExecutionChange) GetValidatorIndex() github_com_prysmaticlabs_prysm_v5_consensus_types_primitives.ValidatorIndex {
	if x != nil {
		return x.ValidatorIndex
	}
	return github_com_prysmaticlabs_prysm_v5_consensus_types_primitives.ValidatorIndex(0)
}

func (x *BLSToExecutionChange) GetFromBlsPubkey() []byte {
	if x != nil {
		return x.FromBlsPubkey
	}
	return nil
}

func (x *BLSToExecutionChange) GetToExecutionAddress() []byte {
	if x != nil {
		return x.ToExecutionAddress
	}
	return nil
}

type SignedBLSToExecutionChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *BLSToExecutionChange `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte                `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *SignedBLSToExecutionChange) Reset() {
	*x = SignedBLSToExecutionChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_withdrawals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedBLSToExecutionChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedBLSToExecutionChange) ProtoMessage() {}

func (x *SignedBLSToExecutionChange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_withdrawals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedBLSToExecutionChange.ProtoReflect.Descriptor instead.
func (*SignedBLSToExecutionChange) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_withdrawals_proto_rawDescGZIP(), []int{1}
}

func (x *SignedBLSToExecutionChange) GetMessage() *BLSToExecutionChange {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedBLSToExecutionChange) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_proto_eth_v2_withdrawals_proto protoreflect.FileDescriptor

var file_proto_eth_v2_withdrawals_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x32, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa,
	0x01, 0x0a, 0x14, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x78, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x4f, 0x82, 0xb5, 0x18, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x35, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x2e, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x62, 0x6c, 0x73, 0x5f, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02,
	0x34, 0x38, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x73, 0x50, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x12, 0x38, 0x0a, 0x14, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0x8a, 0xb5, 0x18, 0x02, 0x32, 0x30, 0x52, 0x12, 0x74, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x1a,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x4c, 0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x4c,
	0x53, 0x54, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06,
	0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x42, 0x81, 0x01, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x35,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x3b, 0x65, 0x74,
	0x68, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68,
	0x2e, 0x56, 0x32, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45,
	0x74, 0x68, 0x5c, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eth_v2_withdrawals_proto_rawDescOnce sync.Once
	file_proto_eth_v2_withdrawals_proto_rawDescData = file_proto_eth_v2_withdrawals_proto_rawDesc
)

func file_proto_eth_v2_withdrawals_proto_rawDescGZIP() []byte {
	file_proto_eth_v2_withdrawals_proto_rawDescOnce.Do(func() {
		file_proto_eth_v2_withdrawals_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eth_v2_withdrawals_proto_rawDescData)
	})
	return file_proto_eth_v2_withdrawals_proto_rawDescData
}

var file_proto_eth_v2_withdrawals_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_eth_v2_withdrawals_proto_goTypes = []interface{}{
	(*BLSToExecutionChange)(nil),       // 0: ethereum.eth.v2.BLSToExecutionChange
	(*SignedBLSToExecutionChange)(nil), // 1: ethereum.eth.v2.SignedBLSToExecutionChange
}
var file_proto_eth_v2_withdrawals_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v2.SignedBLSToExecutionChange.message:type_name -> ethereum.eth.v2.BLSToExecutionChange
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_eth_v2_withdrawals_proto_init() }
func file_proto_eth_v2_withdrawals_proto_init() {
	if File_proto_eth_v2_withdrawals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eth_v2_withdrawals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLSToExecutionChange); i {
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
		file_proto_eth_v2_withdrawals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedBLSToExecutionChange); i {
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
			RawDescriptor: file_proto_eth_v2_withdrawals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eth_v2_withdrawals_proto_goTypes,
		DependencyIndexes: file_proto_eth_v2_withdrawals_proto_depIdxs,
		MessageInfos:      file_proto_eth_v2_withdrawals_proto_msgTypes,
	}.Build()
	File_proto_eth_v2_withdrawals_proto = out.File
	file_proto_eth_v2_withdrawals_proto_rawDesc = nil
	file_proto_eth_v2_withdrawals_proto_goTypes = nil
	file_proto_eth_v2_withdrawals_proto_depIdxs = nil
}
