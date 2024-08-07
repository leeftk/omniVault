// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: halo/registry/keeper/registry.proto

package keeper

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
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

// Network defines an instance of the cross-chain network configuration including supported chains, portals, and shards.
type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Auto-incremented ID
	CreatedHeight uint64    `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"` // Height this network was created at
	Portals       []*Portal `protobuf:"bytes,3,rep,name=portals,proto3" json:"portals,omitempty"`                                   // Supported portals by source chain.
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_registry_keeper_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_halo_registry_keeper_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_halo_registry_keeper_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Network) GetCreatedHeight() uint64 {
	if x != nil {
		return x.CreatedHeight
	}
	return 0
}

func (x *Network) GetPortals() []*Portal {
	if x != nil {
		return x.Portals
	}
	return nil
}

type Portal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId      uint64   `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`                // Chain ID as per https://chainlist.org/
	Address      []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                                // Portal contract address
	DeployHeight uint64   `protobuf:"varint,3,opt,name=deploy_height,json=deployHeight,proto3" json:"deploy_height,omitempty"` // Height this portal contract was deployed at
	ShardIds     []uint64 `protobuf:"varint,4,rep,packed,name=shard_ids,json=shardIds,proto3" json:"shard_ids,omitempty"`      // Shard IDs supported by this portal
}

func (x *Portal) Reset() {
	*x = Portal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halo_registry_keeper_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Portal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Portal) ProtoMessage() {}

func (x *Portal) ProtoReflect() protoreflect.Message {
	mi := &file_halo_registry_keeper_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Portal.ProtoReflect.Descriptor instead.
func (*Portal) Descriptor() ([]byte, []int) {
	return file_halo_registry_keeper_registry_proto_rawDescGZIP(), []int{1}
}

func (x *Portal) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Portal) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Portal) GetDeployHeight() uint64 {
	if x != nil {
		return x.DeployHeight
	}
	return 0
}

func (x *Portal) GetShardIds() []uint64 {
	if x != nil {
		return x.ShardIds
	}
	return nil
}

var File_halo_registry_keeper_registry_proto protoreflect.FileDescriptor

var file_halo_registry_keeper_registry_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x61, 0x6c, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x61, 0x6c, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x52, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x73, 0x3a,
	0x10, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x0a, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x18,
	0x04, 0x22, 0x7f, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x73, 0x42, 0xce, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x6c, 0x6f, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x42,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x6e,
	0x69, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x68,
	0x61, 0x6c, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x48, 0x52, 0x4b, 0xaa, 0x02, 0x14, 0x48, 0x61, 0x6c, 0x6f,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0xca, 0x02, 0x14, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0xe2, 0x02, 0x20, 0x48, 0x61, 0x6c, 0x6f, 0x5c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x48, 0x61, 0x6c,
	0x6f, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x4b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_halo_registry_keeper_registry_proto_rawDescOnce sync.Once
	file_halo_registry_keeper_registry_proto_rawDescData = file_halo_registry_keeper_registry_proto_rawDesc
)

func file_halo_registry_keeper_registry_proto_rawDescGZIP() []byte {
	file_halo_registry_keeper_registry_proto_rawDescOnce.Do(func() {
		file_halo_registry_keeper_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_halo_registry_keeper_registry_proto_rawDescData)
	})
	return file_halo_registry_keeper_registry_proto_rawDescData
}

var file_halo_registry_keeper_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_halo_registry_keeper_registry_proto_goTypes = []any{
	(*Network)(nil), // 0: halo.registry.keeper.Network
	(*Portal)(nil),  // 1: halo.registry.keeper.Portal
}
var file_halo_registry_keeper_registry_proto_depIdxs = []int32{
	1, // 0: halo.registry.keeper.Network.portals:type_name -> halo.registry.keeper.Portal
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_halo_registry_keeper_registry_proto_init() }
func file_halo_registry_keeper_registry_proto_init() {
	if File_halo_registry_keeper_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_halo_registry_keeper_registry_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Network); i {
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
		file_halo_registry_keeper_registry_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Portal); i {
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
			RawDescriptor: file_halo_registry_keeper_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halo_registry_keeper_registry_proto_goTypes,
		DependencyIndexes: file_halo_registry_keeper_registry_proto_depIdxs,
		MessageInfos:      file_halo_registry_keeper_registry_proto_msgTypes,
	}.Build()
	File_halo_registry_keeper_registry_proto = out.File
	file_halo_registry_keeper_registry_proto_rawDesc = nil
	file_halo_registry_keeper_registry_proto_goTypes = nil
	file_halo_registry_keeper_registry_proto_depIdxs = nil
}