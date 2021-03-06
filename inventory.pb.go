// Code generated by protoc-gen-go.
// source: inventory.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

type EggIncubatorType int32

const (
	EggIncubatorType_INCUBATOR_UNSET    EggIncubatorType = 0
	EggIncubatorType_INCUBATOR_DISTANCE EggIncubatorType = 1
)

var EggIncubatorType_name = map[int32]string{
	0: "INCUBATOR_UNSET",
	1: "INCUBATOR_DISTANCE",
}
var EggIncubatorType_value = map[string]int32{
	"INCUBATOR_UNSET":    0,
	"INCUBATOR_DISTANCE": 1,
}

func (x EggIncubatorType) String() string {
	return proto.EnumName(EggIncubatorType_name, int32(x))
}
func (EggIncubatorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type InventoryUpgradeType int32

const (
	InventoryUpgradeType_UPGRADE_UNSET            InventoryUpgradeType = 0
	InventoryUpgradeType_INCREASE_ITEM_STORAGE    InventoryUpgradeType = 1
	InventoryUpgradeType_INCREASE_POKEMON_STORAGE InventoryUpgradeType = 2
)

var InventoryUpgradeType_name = map[int32]string{
	0: "UPGRADE_UNSET",
	1: "INCREASE_ITEM_STORAGE",
	2: "INCREASE_POKEMON_STORAGE",
}
var InventoryUpgradeType_value = map[string]int32{
	"UPGRADE_UNSET":            0,
	"INCREASE_ITEM_STORAGE":    1,
	"INCREASE_POKEMON_STORAGE": 2,
}

func (x InventoryUpgradeType) String() string {
	return proto.EnumName(InventoryUpgradeType_name, int32(x))
}
func (InventoryUpgradeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type AppliedItem struct {
	ItemId    ItemId   `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemType  ItemType `protobuf:"varint,2,opt,name=item_type,json=itemType,enum=POGOProtos.Inventory.Item.ItemType" json:"item_type,omitempty"`
	ExpireMs  int64    `protobuf:"varint,3,opt,name=expire_ms,json=expireMs" json:"expire_ms,omitempty"`
	AppliedMs int64    `protobuf:"varint,4,opt,name=applied_ms,json=appliedMs" json:"applied_ms,omitempty"`
}

func (m *AppliedItem) Reset()                    { *m = AppliedItem{} }
func (m *AppliedItem) String() string            { return proto.CompactTextString(m) }
func (*AppliedItem) ProtoMessage()               {}
func (*AppliedItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type AppliedItems struct {
	Item []*AppliedItem `protobuf:"bytes,4,rep,name=item" json:"item,omitempty"`
}

func (m *AppliedItems) Reset()                    { *m = AppliedItems{} }
func (m *AppliedItems) String() string            { return proto.CompactTextString(m) }
func (*AppliedItems) ProtoMessage()               {}
func (*AppliedItems) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *AppliedItems) GetItem() []*AppliedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type Candy struct {
	FamilyId PokemonFamilyId `protobuf:"varint,1,opt,name=family_id,json=familyId,enum=POGOProtos.Enums.PokemonFamilyId" json:"family_id,omitempty"`
	Candy    int32           `protobuf:"varint,2,opt,name=candy" json:"candy,omitempty"`
}

func (m *Candy) Reset()                    { *m = Candy{} }
func (m *Candy) String() string            { return proto.CompactTextString(m) }
func (*Candy) ProtoMessage()               {}
func (*Candy) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

type EggIncubator struct {
	Id             string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ItemId         ItemId           `protobuf:"varint,2,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	IncubatorType  EggIncubatorType `protobuf:"varint,3,opt,name=incubator_type,json=incubatorType,enum=POGOProtos.Inventory.EggIncubatorType" json:"incubator_type,omitempty"`
	UsesRemaining  int32            `protobuf:"varint,4,opt,name=uses_remaining,json=usesRemaining" json:"uses_remaining,omitempty"`
	PokemonId      uint64           `protobuf:"varint,5,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
	StartKmWalked  float64          `protobuf:"fixed64,6,opt,name=start_km_walked,json=startKmWalked" json:"start_km_walked,omitempty"`
	TargetKmWalked float64          `protobuf:"fixed64,7,opt,name=target_km_walked,json=targetKmWalked" json:"target_km_walked,omitempty"`
}

func (m *EggIncubator) Reset()                    { *m = EggIncubator{} }
func (m *EggIncubator) String() string            { return proto.CompactTextString(m) }
func (*EggIncubator) ProtoMessage()               {}
func (*EggIncubator) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

type EggIncubators struct {
	EggIncubator []*EggIncubator `protobuf:"bytes,1,rep,name=egg_incubator,json=eggIncubator" json:"egg_incubator,omitempty"`
}

func (m *EggIncubators) Reset()                    { *m = EggIncubators{} }
func (m *EggIncubators) String() string            { return proto.CompactTextString(m) }
func (*EggIncubators) ProtoMessage()               {}
func (*EggIncubators) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *EggIncubators) GetEggIncubator() []*EggIncubator {
	if m != nil {
		return m.EggIncubator
	}
	return nil
}

type InventoryDelta struct {
	OriginalTimestampMs int64            `protobuf:"varint,1,opt,name=original_timestamp_ms,json=originalTimestampMs" json:"original_timestamp_ms,omitempty"`
	NewTimestampMs      int64            `protobuf:"varint,2,opt,name=new_timestamp_ms,json=newTimestampMs" json:"new_timestamp_ms,omitempty"`
	InventoryItems      []*InventoryItem `protobuf:"bytes,3,rep,name=inventory_items,json=inventoryItems" json:"inventory_items,omitempty"`
}

func (m *InventoryDelta) Reset()                    { *m = InventoryDelta{} }
func (m *InventoryDelta) String() string            { return proto.CompactTextString(m) }
func (*InventoryDelta) ProtoMessage()               {}
func (*InventoryDelta) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *InventoryDelta) GetInventoryItems() []*InventoryItem {
	if m != nil {
		return m.InventoryItems
	}
	return nil
}

type InventoryItem struct {
	ModifiedTimestampMs int64                      `protobuf:"varint,1,opt,name=modified_timestamp_ms,json=modifiedTimestampMs" json:"modified_timestamp_ms,omitempty"`
	DeletedItem         *InventoryItem_DeletedItem `protobuf:"bytes,2,opt,name=deleted_item,json=deletedItem" json:"deleted_item,omitempty"`
	InventoryItemData   *InventoryItemData         `protobuf:"bytes,3,opt,name=inventory_item_data,json=inventoryItemData" json:"inventory_item_data,omitempty"`
}

func (m *InventoryItem) Reset()                    { *m = InventoryItem{} }
func (m *InventoryItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem) ProtoMessage()               {}
func (*InventoryItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *InventoryItem) GetDeletedItem() *InventoryItem_DeletedItem {
	if m != nil {
		return m.DeletedItem
	}
	return nil
}

func (m *InventoryItem) GetInventoryItemData() *InventoryItemData {
	if m != nil {
		return m.InventoryItemData
	}
	return nil
}

type InventoryItem_DeletedItem struct {
	PokemonId uint64 `protobuf:"fixed64,1,opt,name=pokemon_id,json=pokemonId" json:"pokemon_id,omitempty"`
}

func (m *InventoryItem_DeletedItem) Reset()                    { *m = InventoryItem_DeletedItem{} }
func (m *InventoryItem_DeletedItem) String() string            { return proto.CompactTextString(m) }
func (*InventoryItem_DeletedItem) ProtoMessage()               {}
func (*InventoryItem_DeletedItem) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6, 0} }

type InventoryItemData struct {
	PokemonData       *PokemonData       `protobuf:"bytes,1,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	Item              *ItemData          `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	PokedexEntry      *PokedexEntry      `protobuf:"bytes,3,opt,name=pokedex_entry,json=pokedexEntry" json:"pokedex_entry,omitempty"`
	PlayerStats       *PlayerStats       `protobuf:"bytes,4,opt,name=player_stats,json=playerStats" json:"player_stats,omitempty"`
	PlayerCurrency    *PlayerCurrency    `protobuf:"bytes,5,opt,name=player_currency,json=playerCurrency" json:"player_currency,omitempty"`
	PlayerCamera      *PlayerCamera      `protobuf:"bytes,6,opt,name=player_camera,json=playerCamera" json:"player_camera,omitempty"`
	InventoryUpgrades *InventoryUpgrades `protobuf:"bytes,7,opt,name=inventory_upgrades,json=inventoryUpgrades" json:"inventory_upgrades,omitempty"`
	AppliedItems      *AppliedItems      `protobuf:"bytes,8,opt,name=applied_items,json=appliedItems" json:"applied_items,omitempty"`
	EggIncubators     *EggIncubators     `protobuf:"bytes,9,opt,name=egg_incubators,json=eggIncubators" json:"egg_incubators,omitempty"`
	Candy             *Candy             `protobuf:"bytes,10,opt,name=candy" json:"candy,omitempty"`
}

func (m *InventoryItemData) Reset()                    { *m = InventoryItemData{} }
func (m *InventoryItemData) String() string            { return proto.CompactTextString(m) }
func (*InventoryItemData) ProtoMessage()               {}
func (*InventoryItemData) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *InventoryItemData) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *InventoryItemData) GetItem() *ItemData {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *InventoryItemData) GetPokedexEntry() *PokedexEntry {
	if m != nil {
		return m.PokedexEntry
	}
	return nil
}

func (m *InventoryItemData) GetPlayerStats() *PlayerStats {
	if m != nil {
		return m.PlayerStats
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCurrency() *PlayerCurrency {
	if m != nil {
		return m.PlayerCurrency
	}
	return nil
}

func (m *InventoryItemData) GetPlayerCamera() *PlayerCamera {
	if m != nil {
		return m.PlayerCamera
	}
	return nil
}

func (m *InventoryItemData) GetInventoryUpgrades() *InventoryUpgrades {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func (m *InventoryItemData) GetAppliedItems() *AppliedItems {
	if m != nil {
		return m.AppliedItems
	}
	return nil
}

func (m *InventoryItemData) GetEggIncubators() *EggIncubators {
	if m != nil {
		return m.EggIncubators
	}
	return nil
}

func (m *InventoryItemData) GetCandy() *Candy {
	if m != nil {
		return m.Candy
	}
	return nil
}

type InventoryUpgrade struct {
	ItemId            ItemId               `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	UpgradeType       InventoryUpgradeType `protobuf:"varint,2,opt,name=upgrade_type,json=upgradeType,enum=POGOProtos.Inventory.InventoryUpgradeType" json:"upgrade_type,omitempty"`
	AdditionalStorage int32                `protobuf:"varint,3,opt,name=additional_storage,json=additionalStorage" json:"additional_storage,omitempty"`
}

func (m *InventoryUpgrade) Reset()                    { *m = InventoryUpgrade{} }
func (m *InventoryUpgrade) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrade) ProtoMessage()               {}
func (*InventoryUpgrade) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

type InventoryUpgrades struct {
	InventoryUpgrades []*InventoryUpgrade `protobuf:"bytes,1,rep,name=inventory_upgrades,json=inventoryUpgrades" json:"inventory_upgrades,omitempty"`
}

func (m *InventoryUpgrades) Reset()                    { *m = InventoryUpgrades{} }
func (m *InventoryUpgrades) String() string            { return proto.CompactTextString(m) }
func (*InventoryUpgrades) ProtoMessage()               {}
func (*InventoryUpgrades) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *InventoryUpgrades) GetInventoryUpgrades() []*InventoryUpgrade {
	if m != nil {
		return m.InventoryUpgrades
	}
	return nil
}

func init() {
	proto.RegisterType((*AppliedItem)(nil), "POGOProtos.Inventory.AppliedItem")
	proto.RegisterType((*AppliedItems)(nil), "POGOProtos.Inventory.AppliedItems")
	proto.RegisterType((*Candy)(nil), "POGOProtos.Inventory.Candy")
	proto.RegisterType((*EggIncubator)(nil), "POGOProtos.Inventory.EggIncubator")
	proto.RegisterType((*EggIncubators)(nil), "POGOProtos.Inventory.EggIncubators")
	proto.RegisterType((*InventoryDelta)(nil), "POGOProtos.Inventory.InventoryDelta")
	proto.RegisterType((*InventoryItem)(nil), "POGOProtos.Inventory.InventoryItem")
	proto.RegisterType((*InventoryItem_DeletedItem)(nil), "POGOProtos.Inventory.InventoryItem.DeletedItem")
	proto.RegisterType((*InventoryItemData)(nil), "POGOProtos.Inventory.InventoryItemData")
	proto.RegisterType((*InventoryUpgrade)(nil), "POGOProtos.Inventory.InventoryUpgrade")
	proto.RegisterType((*InventoryUpgrades)(nil), "POGOProtos.Inventory.InventoryUpgrades")
	proto.RegisterEnum("POGOProtos.Inventory.EggIncubatorType", EggIncubatorType_name, EggIncubatorType_value)
	proto.RegisterEnum("POGOProtos.Inventory.InventoryUpgradeType", InventoryUpgradeType_name, InventoryUpgradeType_value)
}

func init() { proto.RegisterFile("inventory.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0x5e, 0xa7, 0x4d, 0xb7, 0x39, 0x8e, 0xdd, 0x74, 0xda, 0x7d, 0x95, 0xb7, 0xbb, 0x2b, 0x15,
	0x03, 0x25, 0xaa, 0xa0, 0x88, 0x22, 0x84, 0xc4, 0x05, 0x4b, 0x9a, 0x78, 0x2b, 0xb3, 0xe4, 0x43,
	0x93, 0x84, 0x45, 0x48, 0xc8, 0x9a, 0xd6, 0xd3, 0x68, 0x68, 0xfc, 0x21, 0x7b, 0x42, 0x9b, 0x3f,
	0xc3, 0x2f, 0xe1, 0x86, 0x2b, 0x6e, 0xb8, 0xe5, 0xff, 0xa0, 0x99, 0xb1, 0x93, 0x49, 0x36, 0x25,
	0x45, 0xdc, 0x34, 0x73, 0xce, 0x79, 0xe6, 0xf1, 0x33, 0x67, 0xce, 0x9c, 0x53, 0xd8, 0x63, 0xd1,
	0x2f, 0x34, 0xe2, 0x71, 0x3a, 0x3b, 0x4b, 0xd2, 0x98, 0xc7, 0xe8, 0xb0, 0xdf, 0xbb, 0xec, 0xf5,
	0xc5, 0x32, 0x3b, 0xf3, 0x8a, 0xd8, 0xd1, 0xe1, 0x1c, 0xe6, 0x33, 0x4e, 0x43, 0x85, 0x3d, 0x32,
	0x69, 0x34, 0x0d, 0xb3, 0xdc, 0x80, 0x80, 0x70, 0x92, 0xaf, 0xf7, 0xc5, 0xda, 0x4f, 0x26, 0x64,
	0x46, 0x53, 0xe5, 0x72, 0xfe, 0x30, 0xc0, 0x6c, 0x26, 0xc9, 0x84, 0xd1, 0xc0, 0xe3, 0x34, 0x44,
	0x5f, 0xc1, 0x53, 0xc1, 0xe4, 0xb3, 0xa0, 0x6e, 0x1c, 0x1b, 0x0d, 0xfb, 0xfc, 0xbd, 0xb3, 0x75,
	0x5f, 0x3e, 0x13, 0x60, 0xf9, 0xc7, 0x0b, 0xf0, 0x0e, 0x93, 0xbf, 0xe8, 0x1b, 0xa8, 0xc8, 0xbd,
	0x7c, 0x96, 0xd0, 0x7a, 0x49, 0xee, 0x7e, 0x7f, 0xc3, 0xee, 0xe1, 0x2c, 0xa1, 0x78, 0x97, 0xe5,
	0x2b, 0xf4, 0x1c, 0x2a, 0xf4, 0x3e, 0x61, 0x29, 0xf5, 0xc3, 0xac, 0xbe, 0x75, 0x6c, 0x34, 0xb6,
	0xf0, 0xae, 0x72, 0x74, 0x32, 0xf4, 0x12, 0x80, 0x28, 0xa5, 0x22, 0xba, 0x2d, 0xa3, 0x95, 0xdc,
	0xd3, 0xc9, 0x1c, 0x17, 0xaa, 0xda, 0x41, 0x32, 0xf4, 0x05, 0x6c, 0x0b, 0xde, 0xfa, 0xf6, 0xf1,
	0x56, 0xc3, 0x7c, 0xe8, 0x18, 0xda, 0x0e, 0x2c, 0xe1, 0xce, 0x4f, 0x50, 0x6e, 0x91, 0x28, 0x98,
	0xa1, 0xaf, 0xa1, 0x72, 0x43, 0x42, 0x36, 0x99, 0x3d, 0x90, 0x0b, 0x57, 0x26, 0xb9, 0x1f, 0xdf,
	0xd2, 0x30, 0x8e, 0x5e, 0x4b, 0xa4, 0x17, 0xe0, 0xdd, 0x9b, 0x7c, 0x85, 0x0e, 0xa1, 0x7c, 0x2d,
	0x88, 0x64, 0x26, 0xca, 0x58, 0x19, 0xce, 0xef, 0x25, 0xa8, 0xba, 0xe3, 0xb1, 0x17, 0x5d, 0x4f,
	0xaf, 0x08, 0x8f, 0x53, 0x64, 0x43, 0x29, 0xe7, 0xaf, 0xe0, 0x12, 0x0b, 0xf4, 0x0b, 0x28, 0xfd,
	0xdb, 0x0b, 0xe8, 0x80, 0xcd, 0x0a, 0x62, 0x75, 0x0b, 0x5b, 0x92, 0xe2, 0x64, 0x3d, 0x85, 0xae,
	0x43, 0x5e, 0x84, 0xc5, 0x74, 0x13, 0x7d, 0x08, 0xf6, 0x34, 0xa3, 0x99, 0x9f, 0xd2, 0x90, 0xb0,
	0x88, 0x45, 0x63, 0x99, 0xf4, 0x32, 0xb6, 0x84, 0x17, 0x17, 0x4e, 0x71, 0x2f, 0x89, 0xca, 0x82,
	0x10, 0x5d, 0x3e, 0x36, 0x1a, 0xdb, 0xb8, 0x92, 0x7b, 0xbc, 0x00, 0x9d, 0xc0, 0x5e, 0xc6, 0x49,
	0xca, 0xfd, 0xdb, 0xd0, 0xbf, 0x23, 0x93, 0x5b, 0x1a, 0xd4, 0x77, 0x8e, 0x8d, 0x86, 0x81, 0x2d,
	0xe9, 0x7e, 0x13, 0xbe, 0x95, 0x4e, 0xd4, 0x80, 0x1a, 0x27, 0xe9, 0x98, 0xea, 0xc0, 0xa7, 0x12,
	0x68, 0x2b, 0x7f, 0x81, 0x74, 0x7e, 0x00, 0x4b, 0x97, 0x9e, 0xa1, 0x4b, 0xb0, 0xe8, 0x78, 0xec,
	0xcf, 0xd5, 0xd7, 0x0d, 0x79, 0xe7, 0xce, 0xe6, 0x63, 0xe3, 0x2a, 0xd5, 0x2c, 0xe7, 0x37, 0x03,
	0xec, 0x39, 0xb0, 0x4d, 0x27, 0x9c, 0xa0, 0x73, 0x78, 0x16, 0xa7, 0x6c, 0xcc, 0x22, 0x32, 0xf1,
	0x39, 0x0b, 0x69, 0xc6, 0x49, 0x98, 0x88, 0x02, 0x34, 0x64, 0x01, 0x1e, 0x14, 0xc1, 0x61, 0x11,
	0xeb, 0x64, 0xe2, 0x28, 0x11, 0xbd, 0x5b, 0x86, 0x97, 0x24, 0xdc, 0x8e, 0xe8, 0x9d, 0x8e, 0xfc,
	0x4e, 0x7b, 0xe9, 0xf2, 0x09, 0x8b, 0xb2, 0x17, 0xda, 0x1f, 0x7a, 0x38, 0xc5, 0x4a, 0x56, 0xac,
	0xcd, 0x74, 0x33, 0x73, 0x7e, 0x2d, 0x81, 0xb5, 0x84, 0x10, 0xea, 0xc3, 0x38, 0x60, 0x37, 0xe2,
	0xd1, 0xac, 0x53, 0x5f, 0x04, 0x75, 0x4d, 0x18, 0xaa, 0x01, 0x9d, 0x50, 0x4e, 0x03, 0xa9, 0x48,
	0x2a, 0x37, 0xcf, 0x3f, 0x7d, 0x84, 0xa0, 0xb3, 0xb6, 0xda, 0x27, 0xc5, 0x99, 0xc1, 0xc2, 0x40,
	0x6f, 0xe1, 0x60, 0xf9, 0x9c, 0xbe, 0x68, 0x45, 0xb2, 0x3c, 0xcd, 0xf3, 0x8f, 0x1e, 0x41, 0xdd,
	0x26, 0x9c, 0xe0, 0x7d, 0xb6, 0xea, 0x3a, 0xfa, 0x18, 0x4c, 0xed, 0xa3, 0x2b, 0xb5, 0x28, 0x0e,
	0xb9, 0xa3, 0xd5, 0xa2, 0xf3, 0x57, 0x19, 0xf6, 0xdf, 0xa1, 0x45, 0xaf, 0xa0, 0x5a, 0x6c, 0x92,
	0xaa, 0x0c, 0xa9, 0xea, 0x85, 0xae, 0x4a, 0xe0, 0x8a, 0xb7, 0x2e, 0xa5, 0x98, 0xc9, 0xc2, 0x40,
	0x5f, 0xe6, 0xad, 0x46, 0x65, 0x6a, 0x53, 0xcf, 0x93, 0xfb, 0xe5, 0x06, 0x74, 0x01, 0x96, 0xe0,
	0x09, 0xe8, 0xbd, 0x4f, 0x23, 0x9e, 0xce, 0xf2, 0x84, 0xbc, 0x5c, 0xfb, 0xe9, 0x80, 0xde, 0xbb,
	0x02, 0x84, 0xab, 0x89, 0x66, 0xa1, 0xd7, 0x50, 0x55, 0x1d, 0xdd, 0xcf, 0x38, 0xe1, 0xaa, 0x31,
	0xae, 0x88, 0x50, 0x14, 0xaa, 0xed, 0xab, 0x9f, 0x81, 0x80, 0x62, 0x33, 0x59, 0x18, 0xa8, 0x07,
	0x7b, 0x39, 0xcf, 0xf5, 0x34, 0x4d, 0x69, 0x74, 0x3d, 0x93, 0x6f, 0xd9, 0x5c, 0xee, 0x1e, 0xef,
	0x52, 0xb5, 0x72, 0x34, 0xb6, 0x93, 0x25, 0x1b, 0x79, 0x60, 0x15, 0x84, 0x24, 0xa4, 0x29, 0x91,
	0xcf, 0xde, 0x3c, 0xff, 0x60, 0x03, 0x9d, 0xc4, 0xe2, 0xfc, 0x4c, 0xca, 0x42, 0xdf, 0x03, 0x5a,
	0x94, 0xcf, 0x34, 0x19, 0xa7, 0x24, 0xa0, 0x99, 0xec, 0x0e, 0x9b, 0xab, 0x67, 0x94, 0xc3, 0xb5,
	0xea, 0x29, 0x5c, 0xa2, 0x71, 0x14, 0x23, 0x45, 0x3d, 0xbe, 0x5d, 0x49, 0xe9, 0x6c, 0x1c, 0x16,
	0x19, 0xae, 0x12, 0x7d, 0xd8, 0x7c, 0x0b, 0xf6, 0x52, 0x07, 0xca, 0xea, 0x95, 0x7f, 0xaa, 0x85,
	0xa5, 0xf6, 0x85, 0x2d, 0xba, 0xd4, 0xcd, 0x3e, 0x2b, 0x06, 0x07, 0x48, 0x8a, 0xe7, 0xeb, 0x29,
	0xe4, 0x90, 0x2a, 0xa6, 0xca, 0x9f, 0x06, 0xd4, 0x56, 0x0f, 0xfc, 0x9f, 0x46, 0x79, 0x07, 0xaa,
	0x79, 0x9a, 0xf5, 0x69, 0x7e, 0xfa, 0xb8, 0x54, 0xcb, 0x59, 0x62, 0x4e, 0x17, 0x06, 0xfa, 0x04,
	0x10, 0x09, 0x02, 0xc6, 0x59, 0x2c, 0xda, 0x68, 0xc6, 0xe3, 0x94, 0x8c, 0xd5, 0x70, 0x2a, 0xe3,
	0xfd, 0x45, 0x64, 0xa0, 0x02, 0xce, 0xcf, 0xda, 0x2b, 0x9d, 0xdf, 0xd5, 0x68, 0x6d, 0x0d, 0xa8,
	0x4e, 0x7f, 0xf2, 0x38, 0x61, 0x6b, 0x4a, 0xe0, 0xf4, 0x15, 0xd4, 0x56, 0xe7, 0x20, 0x3a, 0x80,
	0x3d, 0xaf, 0xdb, 0x1a, 0x5d, 0x34, 0x87, 0x3d, 0xec, 0x8f, 0xba, 0x03, 0x77, 0x58, 0x7b, 0x82,
	0xfe, 0x07, 0x68, 0xe1, 0x6c, 0x7b, 0x83, 0x61, 0xb3, 0xdb, 0x72, 0x6b, 0xc6, 0xe9, 0x15, 0x1c,
	0xae, 0x4b, 0x00, 0xda, 0x07, 0x6b, 0xd4, 0xbf, 0xc4, 0xcd, 0xb6, 0x3b, 0xa7, 0xf8, 0x3f, 0x3c,
	0xf3, 0xba, 0x2d, 0xec, 0x36, 0x07, 0xae, 0xef, 0x0d, 0xdd, 0x8e, 0x3f, 0x18, 0xf6, 0x70, 0xf3,
	0xd2, 0xad, 0x19, 0xe8, 0x05, 0xd4, 0xe7, 0xa1, 0x7e, 0xef, 0x8d, 0xdb, 0xe9, 0x75, 0xe7, 0xd1,
	0xd2, 0xc5, 0xee, 0x8f, 0x3b, 0xf2, 0xdf, 0xb5, 0xac, 0xff, 0xa4, 0x6f, 0xf4, 0x4b, 0xfd, 0xad,
	0x2b, 0x65, 0x7f, 0xfe, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xf9, 0xf2, 0x83, 0x29, 0x0a,
	0x00, 0x00,
}
