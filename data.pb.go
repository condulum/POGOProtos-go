// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	data.proto
	data_battle.proto
	data_capture.proto
	data_gym.proto
	data_logs.proto
	data_player.proto
	enums.proto
	inventory.proto
	inventory_item.proto
	map_fort.proto
	map_pokemon.proto
	maps.proto
	networking_envelopes.proto
	networking_platform.proto
	networking_platform_requests.proto
	networking_platform_responses.proto
	networking_requests.proto
	networking_requests_messages.proto
	networking_responses.proto
	settings.proto
	settings_master.proto
	settings_master_item.proto
	settings_master_pokemon.proto

It has these top-level messages:
	AssetDigestEntry
	BuddyPokemon
	DownloadUrlEntry
	PlayerBadge
	PlayerData
	PokedexEntry
	PokemonData
	BattleAction
	BattleLog
	BattleParticipant
	BattlePokemonInfo
	BattleResults
	CaptureAward
	CaptureProbability
	GymMembership
	GymState
	ActionLogEntry
	BuddyPokemonLogEntry
	CatchPokemonLogEntry
	FortSearchLogEntry
	ContactSettings
	Currency
	DailyBonus
	EquippedBadge
	PlayerAvatar
	PlayerCamera
	PlayerCurrency
	PlayerPublicProfile
	PlayerStats
	AppliedItem
	AppliedItems
	Candy
	EggIncubator
	EggIncubators
	InventoryDelta
	InventoryItem
	InventoryItemData
	InventoryUpgrade
	InventoryUpgrades
	ItemAward
	ItemData
	FortData
	FortLureInfo
	FortModifier
	FortSummary
	MapPokemon
	NearbyPokemon
	WildPokemon
	MapCell
	SpawnPoint
	AuthTicket
	RequestEnvelope
	ResponseEnvelope
	Signature
	UnknownMessage
	BuyItemAndroidRequest
	BuyItemPokeCoinsRequest
	SendEncryptedSignatureRequest
	BuyItemAndroidResponse
	BuyItemPokeCoinsResponse
	GetStoreItemsResponse
	SendEncryptedSignatureResponse
	Request
	AddFortModifierMessage
	AttackGymMessage
	CatchPokemonMessage
	CheckAwardedBadgesMessage
	CheckChallengeMessage
	CheckCodenameAvailableMessage
	ClaimCodenameMessage
	CollectDailyBonusMessage
	CollectDailyDefenderBonusMessage
	DiskEncounterMessage
	DownloadItemTemplatesMessage
	DownloadRemoteConfigVersionMessage
	DownloadSettingsMessage
	EchoMessage
	EncounterMessage
	EncounterTutorialCompleteMessage
	EquipBadgeMessage
	EvolvePokemonMessage
	FortDeployPokemonMessage
	FortDetailsMessage
	FortRecallPokemonMessage
	FortSearchMessage
	GetAssetDigestMessage
	GetBuddyWalkedMessage
	GetDownloadUrlsMessage
	GetGymDetailsMessage
	GetHatchedEggsMessage
	GetIncensePokemonMessage
	GetInventoryMessage
	GetMapObjectsMessage
	GetPlayerMessage
	GetPlayerProfileMessage
	GetSuggestedCodenamesMessage
	IncenseEncounterMessage
	LevelUpRewardsMessage
	MarkTutorialCompleteMessage
	NicknamePokemonMessage
	PlayerUpdateMessage
	RecycleInventoryItemMessage
	ReleasePokemonMessage
	SetAvatarMessage
	SetBuddyPokemonMessage
	SetContactSettingsMessage
	SetFavoritePokemonMessage
	SetPlayerTeamMessage
	SfidaActionLogMessage
	StartGymBattleMessage
	UpgradePokemonMessage
	UseIncenseMessage
	UseItemCaptureMessage
	UseItemEggIncubatorMessage
	UseItemGymMessage
	UseItemPotionMessage
	UseItemReviveMessage
	UseItemXpBoostMessage
	VerifyChallengeMessage
	AddFortModifierResponse
	AttackGymResponse
	CatchPokemonResponse
	CheckAwardedBadgesResponse
	CheckChallengeResponse
	CheckCodenameAvailableResponse
	ClaimCodenameResponse
	CollectDailyBonusResponse
	CollectDailyDefenderBonusResponse
	DiskEncounterResponse
	DownloadItemTemplatesResponse
	DownloadRemoteConfigVersionResponse
	DownloadSettingsResponse
	EchoResponse
	EncounterResponse
	EncounterTutorialCompleteResponse
	EquipBadgeResponse
	EvolvePokemonResponse
	FortDeployPokemonResponse
	FortDetailsResponse
	FortRecallPokemonResponse
	FortSearchResponse
	GetAssetDigestResponse
	GetBuddyWalkedResponse
	GetDownloadUrlsResponse
	GetGymDetailsResponse
	GetHatchedEggsResponse
	GetIncensePokemonResponse
	GetInventoryResponse
	GetMapObjectsResponse
	GetPlayerProfileResponse
	GetPlayerResponse
	GetSuggestedCodenamesResponse
	IncenseEncounterResponse
	LevelUpRewardsResponse
	MarkTutorialCompleteResponse
	NicknamePokemonResponse
	PlayerUpdateResponse
	RecycleInventoryItemResponse
	ReleasePokemonResponse
	SetAvatarResponse
	SetBuddyPokemonResponse
	SetContactSettingsResponse
	SetFavoritePokemonResponse
	SetPlayerTeamResponse
	SfidaActionLogResponse
	StartGymBattleResponse
	UpgradePokemonResponse
	UseIncenseResponse
	UseItemCaptureResponse
	UseItemEggIncubatorResponse
	UseItemGymResponse
	UseItemPotionResponse
	UseItemReviveResponse
	UseItemXpBoostResponse
	VerifyChallengeResponse
	DownloadSettingsAction
	FortSettings
	GlobalSettings
	GpsSettings
	InventorySettings
	LevelSettings
	MapSettings
	BadgeSettings
	CameraSettings
	EncounterSettings
	EquippedBadgeSettings
	GymBattleSettings
	GymLevelSettings
	IapItemDisplay
	IapSettings
	ItemSettings
	MoveSequenceSettings
	MoveSettings
	PlayerLevelSettings
	PokemonSettings
	PokemonUpgradeSettings
	TypeEffectiveSettings
	BattleAttributes
	EggIncubatorAttributes
	ExperienceBoostAttributes
	FoodAttributes
	FortModifierAttributes
	IncenseAttributes
	InventoryUpgradeAttributes
	PokeballAttributes
	PotionAttributes
	ReviveAttributes
	CameraAttributes
	EncounterAttributes
	StatsAttributes
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type AssetDigestEntry struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	BundleName string `protobuf:"bytes,2,opt,name=bundle_name,json=bundleName" json:"bundle_name,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Checksum   uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
	Size       int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Key        []byte `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AssetDigestEntry) Reset()                    { *m = AssetDigestEntry{} }
func (m *AssetDigestEntry) String() string            { return proto.CompactTextString(m) }
func (*AssetDigestEntry) ProtoMessage()               {}
func (*AssetDigestEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BuddyPokemon struct {
	Id            uint64  `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	StartKmWalked float64 `protobuf:"fixed64,2,opt,name=start_km_walked,json=startKmWalked" json:"start_km_walked,omitempty"`
	LastKmAwarded float64 `protobuf:"fixed64,3,opt,name=last_km_awarded,json=lastKmAwarded" json:"last_km_awarded,omitempty"`
}

func (m *BuddyPokemon) Reset()                    { *m = BuddyPokemon{} }
func (m *BuddyPokemon) String() string            { return proto.CompactTextString(m) }
func (*BuddyPokemon) ProtoMessage()               {}
func (*BuddyPokemon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DownloadUrlEntry struct {
	AssetId  string `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Checksum uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *DownloadUrlEntry) Reset()                    { *m = DownloadUrlEntry{} }
func (m *DownloadUrlEntry) String() string            { return proto.CompactTextString(m) }
func (*DownloadUrlEntry) ProtoMessage()               {}
func (*DownloadUrlEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PlayerBadge struct {
	BadgeType    BadgeType `protobuf:"varint,1,opt,name=badge_type,json=badgeType,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	Rank         int32     `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"`
	StartValue   int32     `protobuf:"varint,3,opt,name=start_value,json=startValue" json:"start_value,omitempty"`
	EndValue     int32     `protobuf:"varint,4,opt,name=end_value,json=endValue" json:"end_value,omitempty"`
	CurrentValue float64   `protobuf:"fixed64,5,opt,name=current_value,json=currentValue" json:"current_value,omitempty"`
}

func (m *PlayerBadge) Reset()                    { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string            { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()               {}
func (*PlayerBadge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PlayerData struct {
	CreationTimestampMs     int64            `protobuf:"varint,1,opt,name=creation_timestamp_ms,json=creationTimestampMs" json:"creation_timestamp_ms,omitempty"`
	Username                string           `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team                    TeamColor        `protobuf:"varint,5,opt,name=team,enum=POGOProtos.Enums.TeamColor" json:"team,omitempty"`
	TutorialState           []TutorialState  `protobuf:"varint,7,rep,packed,name=tutorial_state,json=tutorialState,enum=POGOProtos.Enums.TutorialState" json:"tutorial_state,omitempty"`
	Avatar                  *PlayerAvatar    `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	MaxPokemonStorage       int32            `protobuf:"varint,9,opt,name=max_pokemon_storage,json=maxPokemonStorage" json:"max_pokemon_storage,omitempty"`
	MaxItemStorage          int32            `protobuf:"varint,10,opt,name=max_item_storage,json=maxItemStorage" json:"max_item_storage,omitempty"`
	DailyBonus              *DailyBonus      `protobuf:"bytes,11,opt,name=daily_bonus,json=dailyBonus" json:"daily_bonus,omitempty"`
	EquippedBadge           *EquippedBadge   `protobuf:"bytes,12,opt,name=equipped_badge,json=equippedBadge" json:"equipped_badge,omitempty"`
	ContactSettings         *ContactSettings `protobuf:"bytes,13,opt,name=contact_settings,json=contactSettings" json:"contact_settings,omitempty"`
	Currencies              []*Currency      `protobuf:"bytes,14,rep,name=currencies" json:"currencies,omitempty"`
	RemainingCodenameClaims int32            `protobuf:"varint,15,opt,name=remaining_codename_claims,json=remainingCodenameClaims" json:"remaining_codename_claims,omitempty"`
	BuddyPokemon            *BuddyPokemon    `protobuf:"bytes,16,opt,name=buddy_pokemon,json=buddyPokemon" json:"buddy_pokemon,omitempty"`
}

func (m *PlayerData) Reset()                    { *m = PlayerData{} }
func (m *PlayerData) String() string            { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()               {}
func (*PlayerData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlayerData) GetAvatar() *PlayerAvatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *PlayerData) GetDailyBonus() *DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *PlayerData) GetEquippedBadge() *EquippedBadge {
	if m != nil {
		return m.EquippedBadge
	}
	return nil
}

func (m *PlayerData) GetContactSettings() *ContactSettings {
	if m != nil {
		return m.ContactSettings
	}
	return nil
}

func (m *PlayerData) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *PlayerData) GetBuddyPokemon() *BuddyPokemon {
	if m != nil {
		return m.BuddyPokemon
	}
	return nil
}

type PokedexEntry struct {
	PokemonId            PokemonId `protobuf:"varint,1,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	TimesEncountered     int32     `protobuf:"varint,2,opt,name=times_encountered,json=timesEncountered" json:"times_encountered,omitempty"`
	TimesCaptured        int32     `protobuf:"varint,3,opt,name=times_captured,json=timesCaptured" json:"times_captured,omitempty"`
	EvolutionStonePieces int32     `protobuf:"varint,4,opt,name=evolution_stone_pieces,json=evolutionStonePieces" json:"evolution_stone_pieces,omitempty"`
	EvolutionStones      int32     `protobuf:"varint,5,opt,name=evolution_stones,json=evolutionStones" json:"evolution_stones,omitempty"`
}

func (m *PokedexEntry) Reset()                    { *m = PokedexEntry{} }
func (m *PokedexEntry) String() string            { return proto.CompactTextString(m) }
func (*PokedexEntry) ProtoMessage()               {}
func (*PokedexEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PokemonData struct {
	Id                     uint64      `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	PokemonId              PokemonId   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Cp                     int32       `protobuf:"varint,3,opt,name=cp" json:"cp,omitempty"`
	Stamina                int32       `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	StaminaMax             int32       `protobuf:"varint,5,opt,name=stamina_max,json=staminaMax" json:"stamina_max,omitempty"`
	Move_1                 PokemonMove `protobuf:"varint,6,opt,name=move_1,json=move1,enum=POGOProtos.Enums.PokemonMove" json:"move_1,omitempty"`
	Move_2                 PokemonMove `protobuf:"varint,7,opt,name=move_2,json=move2,enum=POGOProtos.Enums.PokemonMove" json:"move_2,omitempty"`
	DeployedFortId         string      `protobuf:"bytes,8,opt,name=deployed_fort_id,json=deployedFortId" json:"deployed_fort_id,omitempty"`
	OwnerName              string      `protobuf:"bytes,9,opt,name=owner_name,json=ownerName" json:"owner_name,omitempty"`
	IsEgg                  bool        `protobuf:"varint,10,opt,name=is_egg,json=isEgg" json:"is_egg,omitempty"`
	EggKmWalkedTarget      float64     `protobuf:"fixed64,11,opt,name=egg_km_walked_target,json=eggKmWalkedTarget" json:"egg_km_walked_target,omitempty"`
	EggKmWalkedStart       float64     `protobuf:"fixed64,12,opt,name=egg_km_walked_start,json=eggKmWalkedStart" json:"egg_km_walked_start,omitempty"`
	Origin                 int32       `protobuf:"varint,14,opt,name=origin" json:"origin,omitempty"`
	HeightM                float32     `protobuf:"fixed32,15,opt,name=height_m,json=heightM" json:"height_m,omitempty"`
	WeightKg               float32     `protobuf:"fixed32,16,opt,name=weight_kg,json=weightKg" json:"weight_kg,omitempty"`
	IndividualAttack       int32       `protobuf:"varint,17,opt,name=individual_attack,json=individualAttack" json:"individual_attack,omitempty"`
	IndividualDefense      int32       `protobuf:"varint,18,opt,name=individual_defense,json=individualDefense" json:"individual_defense,omitempty"`
	IndividualStamina      int32       `protobuf:"varint,19,opt,name=individual_stamina,json=individualStamina" json:"individual_stamina,omitempty"`
	CpMultiplier           float32     `protobuf:"fixed32,20,opt,name=cp_multiplier,json=cpMultiplier" json:"cp_multiplier,omitempty"`
	Pokeball               ItemId      `protobuf:"varint,21,opt,name=pokeball,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball,omitempty"`
	CapturedCellId         uint64      `protobuf:"varint,22,opt,name=captured_cell_id,json=capturedCellId" json:"captured_cell_id,omitempty"`
	BattlesAttacked        int32       `protobuf:"varint,23,opt,name=battles_attacked,json=battlesAttacked" json:"battles_attacked,omitempty"`
	BattlesDefended        int32       `protobuf:"varint,24,opt,name=battles_defended,json=battlesDefended" json:"battles_defended,omitempty"`
	EggIncubatorId         string      `protobuf:"bytes,25,opt,name=egg_incubator_id,json=eggIncubatorId" json:"egg_incubator_id,omitempty"`
	CreationTimeMs         uint64      `protobuf:"varint,26,opt,name=creation_time_ms,json=creationTimeMs" json:"creation_time_ms,omitempty"`
	NumUpgrades            int32       `protobuf:"varint,27,opt,name=num_upgrades,json=numUpgrades" json:"num_upgrades,omitempty"`
	AdditionalCpMultiplier float32     `protobuf:"fixed32,28,opt,name=additional_cp_multiplier,json=additionalCpMultiplier" json:"additional_cp_multiplier,omitempty"`
	Favorite               int32       `protobuf:"varint,29,opt,name=favorite" json:"favorite,omitempty"`
	Nickname               string      `protobuf:"bytes,30,opt,name=nickname" json:"nickname,omitempty"`
	FromFort               int32       `protobuf:"varint,31,opt,name=from_fort,json=fromFort" json:"from_fort,omitempty"`
	BuddyCandyAwarded      int32       `protobuf:"varint,32,opt,name=buddy_candy_awarded,json=buddyCandyAwarded" json:"buddy_candy_awarded,omitempty"`
}

func (m *PokemonData) Reset()                    { *m = PokemonData{} }
func (m *PokemonData) String() string            { return proto.CompactTextString(m) }
func (*PokemonData) ProtoMessage()               {}
func (*PokemonData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*AssetDigestEntry)(nil), "POGOProtos.Data.AssetDigestEntry")
	proto.RegisterType((*BuddyPokemon)(nil), "POGOProtos.Data.BuddyPokemon")
	proto.RegisterType((*DownloadUrlEntry)(nil), "POGOProtos.Data.DownloadUrlEntry")
	proto.RegisterType((*PlayerBadge)(nil), "POGOProtos.Data.PlayerBadge")
	proto.RegisterType((*PlayerData)(nil), "POGOProtos.Data.PlayerData")
	proto.RegisterType((*PokedexEntry)(nil), "POGOProtos.Data.PokedexEntry")
	proto.RegisterType((*PokemonData)(nil), "POGOProtos.Data.PokemonData")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x6f, 0x73, 0x1b, 0x3d,
	0x11, 0xe7, 0xec, 0xd8, 0xb1, 0xe5, 0x3f, 0x71, 0x94, 0x34, 0xcf, 0x35, 0xa5, 0xd4, 0x8f, 0xa1,
	0x60, 0x86, 0x69, 0x3a, 0x0d, 0x7d, 0xc1, 0x74, 0x60, 0x86, 0xc4, 0x09, 0x8c, 0x29, 0xa1, 0x9e,
	0x4b, 0x0a, 0x33, 0xbc, 0xb9, 0x91, 0x4f, 0x9b, 0xab, 0xc6, 0x77, 0xba, 0x43, 0xd2, 0x39, 0x31,
	0x1f, 0x85, 0x8f, 0xc0, 0xa7, 0xe0, 0x0d, 0x9f, 0x0b, 0x46, 0x7b, 0x3a, 0xc7, 0x0e, 0x0d, 0xf0,
	0xbc, 0xb1, 0xb5, 0xbf, 0xfd, 0xed, 0x69, 0xb5, 0xda, 0x3f, 0x22, 0x84, 0x33, 0xc3, 0x4e, 0x72,
	0x95, 0x99, 0x8c, 0xee, 0xcd, 0x3e, 0xfd, 0xf6, 0xd3, 0xcc, 0x2e, 0xf5, 0xc9, 0x05, 0x33, 0xec,
	0xb8, 0x03, 0xb2, 0x48, 0x75, 0xa9, 0x3d, 0xde, 0xb7, 0xcc, 0x30, 0x4f, 0xd8, 0x0a, 0x94, 0x83,
	0x0e, 0x85, 0x5c, 0x82, 0x34, 0x99, 0x5a, 0x85, 0xc2, 0x40, 0x5a, 0xa2, 0xa3, 0xbf, 0x7b, 0x64,
	0x70, 0xa6, 0x35, 0x98, 0x0b, 0x11, 0x83, 0x36, 0x97, 0xd2, 0xa8, 0x15, 0x7d, 0x4e, 0x5a, 0xcc,
	0x62, 0xa1, 0xe0, 0xbe, 0x37, 0xf4, 0xc6, 0xed, 0x60, 0x17, 0xe5, 0x29, 0xa7, 0xaf, 0x48, 0x67,
	0x5e, 0x48, 0x9e, 0x40, 0x28, 0x59, 0x0a, 0x7e, 0x0d, 0xb5, 0xa4, 0x84, 0xfe, 0xc0, 0x52, 0xa0,
	0x3e, 0xd9, 0x5d, 0x82, 0xd2, 0x22, 0x93, 0x7e, 0x7d, 0xe8, 0x8d, 0xeb, 0x41, 0x25, 0xd2, 0x63,
	0xd2, 0x8a, 0xbe, 0x40, 0xb4, 0xd0, 0x45, 0xea, 0xef, 0x0c, 0xbd, 0xf1, 0x6e, 0xb0, 0x96, 0x29,
	0x25, 0x3b, 0x5a, 0xfc, 0x15, 0xfc, 0xc6, 0xd0, 0x1b, 0x37, 0x02, 0x5c, 0xd3, 0x01, 0xa9, 0x2f,
	0x60, 0xe5, 0x37, 0x87, 0xde, 0xb8, 0x1b, 0xd8, 0xe5, 0x48, 0x92, 0xee, 0x79, 0xc1, 0xf9, 0x6a,
	0x96, 0x2d, 0x20, 0xcd, 0x24, 0xed, 0x93, 0x9a, 0xf3, 0xb0, 0x19, 0xd4, 0x04, 0xa7, 0x3f, 0x26,
	0x7b, 0xda, 0x30, 0x65, 0xc2, 0x45, 0x1a, 0xde, 0xb1, 0x64, 0x01, 0x1c, 0x1d, 0xf4, 0x82, 0x1e,
	0xc2, 0x1f, 0xd3, 0x3f, 0x21, 0x68, 0x79, 0x09, 0xd3, 0x48, 0x63, 0x77, 0x4c, 0x71, 0xe0, 0xe8,
	0xab, 0x17, 0xf4, 0x2c, 0xfc, 0x31, 0x3d, 0x2b, 0xc1, 0x51, 0x46, 0x06, 0x17, 0xd9, 0x9d, 0x4c,
	0x32, 0xc6, 0x3f, 0xab, 0xe4, 0x7f, 0xc6, 0x66, 0x40, 0xea, 0x85, 0x4a, 0x5c, 0x4c, 0xec, 0x72,
	0x7d, 0xac, 0xfa, 0xc6, 0xb1, 0xfe, 0x4b, 0x18, 0x46, 0xff, 0xf0, 0x48, 0x67, 0x86, 0x97, 0x76,
	0xce, 0x78, 0x0c, 0xf4, 0x03, 0x21, 0x73, 0xbb, 0x08, 0xcd, 0x2a, 0x07, 0xdc, 0xae, 0x7f, 0xfa,
	0xe2, 0x64, 0xe3, 0xe6, 0x2f, 0xf1, 0xce, 0x91, 0x7c, 0xb3, 0xca, 0x21, 0x68, 0xcf, 0xab, 0xa5,
	0xdd, 0x5b, 0x31, 0xb9, 0x40, 0x77, 0x1a, 0x01, 0xae, 0xed, 0xed, 0x95, 0x01, 0x5a, 0xb2, 0xa4,
	0xa8, 0xdc, 0x22, 0x08, 0xfd, 0xd1, 0x22, 0xf4, 0x05, 0x69, 0x83, 0xe4, 0x4e, 0xbd, 0x83, 0xea,
	0x16, 0x48, 0x5e, 0x2a, 0x7f, 0x48, 0x7a, 0x51, 0xa1, 0x14, 0xc8, 0xca, 0xbe, 0x81, 0x41, 0xeb,
	0x3a, 0x10, 0x49, 0xa3, 0xbf, 0x35, 0x09, 0x29, 0x8f, 0x60, 0xb3, 0x92, 0x9e, 0x92, 0x67, 0x91,
	0x02, 0x66, 0x44, 0x26, 0x43, 0x23, 0x52, 0xd0, 0x86, 0xa5, 0x79, 0x98, 0x6a, 0x3c, 0x4c, 0x3d,
	0x38, 0xa8, 0x94, 0x37, 0x95, 0xee, 0x4a, 0xdb, 0x08, 0x15, 0x1a, 0xd4, 0x46, 0x82, 0xad, 0x65,
	0xfa, 0x96, 0xec, 0x18, 0x60, 0x29, 0x6e, 0xfd, 0xd5, 0x58, 0xdc, 0x00, 0x4b, 0x27, 0x59, 0x92,
	0xa9, 0x00, 0x89, 0xf4, 0x77, 0xa4, 0x6f, 0x0a, 0x93, 0x29, 0xc1, 0x92, 0x50, 0x1b, 0x66, 0xc0,
	0xdf, 0x1d, 0xd6, 0xc7, 0xfd, 0xd3, 0x57, 0x5f, 0x31, 0x75, 0xbc, 0x6b, 0x4b, 0x3b, 0xaf, 0x0d,
	0xbc, 0xa0, 0x67, 0x36, 0x21, 0xfa, 0x4b, 0xd2, 0x64, 0x4b, 0x66, 0x98, 0xf2, 0x5b, 0x43, 0x6f,
	0xdc, 0x39, 0xfd, 0xd1, 0xc9, 0xa3, 0x22, 0x3c, 0x29, 0x4f, 0xee, 0xfe, 0xce, 0x90, 0x1b, 0x38,
	0x1b, 0x7a, 0x42, 0x0e, 0x52, 0x76, 0x1f, 0xe6, 0x65, 0xf2, 0x86, 0xda, 0x64, 0x8a, 0xc5, 0xe0,
	0xb7, 0x31, 0xca, 0xfb, 0x29, 0xbb, 0x77, 0x69, 0x7d, 0x5d, 0x2a, 0xe8, 0x98, 0x0c, 0x2c, 0xdf,
	0x16, 0xeb, 0x9a, 0x4c, 0x90, 0xdc, 0x4f, 0xd9, 0xfd, 0xd4, 0x40, 0x5a, 0x31, 0x27, 0xa4, 0xc3,
	0x99, 0x48, 0x56, 0xe1, 0x3c, 0x93, 0x85, 0xf6, 0x3b, 0xe8, 0xdc, 0xe8, 0x29, 0xe7, 0x2e, 0x2c,
	0xf5, 0xdc, 0x32, 0x03, 0xc2, 0xd7, 0x6b, 0xfa, 0x7b, 0xd2, 0x87, 0xbf, 0x14, 0x22, 0xcf, 0x81,
	0x87, 0x98, 0x45, 0x7e, 0x17, 0xbf, 0xf3, 0xfa, 0xa9, 0xef, 0x5c, 0x3a, 0x36, 0x66, 0x5f, 0xd0,
	0x83, 0x4d, 0x91, 0x06, 0x64, 0x10, 0x65, 0xd2, 0xb0, 0xc8, 0x84, 0x1a, 0x8c, 0x11, 0x32, 0xd6,
	0x7e, 0x0f, 0xbf, 0xf7, 0x93, 0xa7, 0xbe, 0x37, 0x29, 0xf9, 0xd7, 0x8e, 0x1e, 0xec, 0x45, 0xdb,
	0x00, 0xfd, 0x35, 0x21, 0x65, 0xaa, 0x45, 0x02, 0xb4, 0xdf, 0x1f, 0xd6, 0xc7, 0x9d, 0xd3, 0xe1,
	0x93, 0x5f, 0x2b, 0x99, 0xab, 0x60, 0xc3, 0x86, 0x7e, 0x20, 0xcf, 0x15, 0xa4, 0x4c, 0x48, 0x21,
	0xe3, 0x30, 0xca, 0x38, 0xd8, 0x9c, 0x0a, 0xa3, 0x84, 0x89, 0x54, 0xfb, 0x7b, 0x18, 0xdb, 0x6f,
	0xd6, 0x84, 0x89, 0xd3, 0x4f, 0x50, 0x4d, 0xcf, 0x49, 0x6f, 0x6e, 0x9b, 0x4f, 0x75, 0x81, 0xfe,
	0x00, 0x8f, 0xf3, 0xf2, 0x3f, 0x1c, 0xd8, 0x6c, 0x51, 0x41, 0x77, 0xbe, 0x21, 0x8d, 0xfe, 0xe5,
	0x91, 0xae, 0x5d, 0x73, 0xb8, 0x2f, 0xbb, 0xc9, 0x07, 0x42, 0xaa, 0x7c, 0x70, 0xfd, 0xe4, 0xab,
	0x49, 0xed, 0xec, 0xa7, 0x3c, 0x68, 0xe7, 0xd5, 0x92, 0xfe, 0x8c, 0xec, 0x63, 0x45, 0x85, 0x20,
	0xa3, 0xac, 0x90, 0x06, 0x94, 0xeb, 0x77, 0x8d, 0x60, 0x80, 0x8a, 0xcb, 0x07, 0x9c, 0xbe, 0x26,
	0xfd, 0x92, 0x1c, 0xb1, 0xdc, 0x14, 0xca, 0x75, 0xbc, 0x46, 0xd0, 0x43, 0x74, 0xe2, 0x40, 0xfa,
	0x9e, 0x1c, 0xc1, 0x32, 0x4b, 0x0a, 0xac, 0x57, 0x6d, 0x32, 0x09, 0x61, 0x2e, 0x20, 0x02, 0xed,
	0x9a, 0xc1, 0xe1, 0x5a, 0x7b, 0x6d, 0x95, 0x33, 0xd4, 0xd1, 0x9f, 0x92, 0xc1, 0x23, 0x2b, 0xed,
	0x3a, 0xf9, 0xde, 0x36, 0x5f, 0x8f, 0xfe, 0xd9, 0x26, 0x1d, 0x77, 0x1a, 0xec, 0x0f, 0x8f, 0x5b,
	0xf8, 0x76, 0x40, 0x6a, 0xdf, 0x29, 0x20, 0x7d, 0x52, 0x8b, 0x72, 0x77, 0xae, 0x5a, 0x94, 0xdb,
	0x51, 0x64, 0x5b, 0x8a, 0x90, 0xcc, 0x79, 0x5f, 0x89, 0xae, 0x0f, 0xda, 0x65, 0x98, 0xb2, 0x7b,
	0xe7, 0x2b, 0x71, 0xd0, 0x15, 0xbb, 0xa7, 0xef, 0x49, 0x33, 0xcd, 0x96, 0x10, 0xbe, 0xc3, 0xf1,
	0xd3, 0xdf, 0xbe, 0xe5, 0x2d, 0x17, 0xae, 0xb2, 0x25, 0x04, 0x0d, 0x4b, 0x7e, 0xb7, 0xb6, 0x3a,
	0xf5, 0x77, 0xff, 0x6f, 0xab, 0x53, 0x5b, 0xe7, 0x1c, 0xf2, 0x24, 0x5b, 0x01, 0x0f, 0x6f, 0x33,
	0x85, 0x93, 0xa5, 0x85, 0x6d, 0xaf, 0x5f, 0xe1, 0xbf, 0xc9, 0x94, 0x1d, 0x30, 0x2f, 0x09, 0xc9,
	0xee, 0x24, 0xa8, 0x72, 0xf6, 0xb6, 0x91, 0xd3, 0x46, 0x04, 0x47, 0xef, 0x33, 0xd2, 0x14, 0x3a,
	0x84, 0x38, 0xc6, 0x36, 0xd1, 0x0a, 0x1a, 0x42, 0x5f, 0xc6, 0x31, 0x7d, 0x4b, 0x0e, 0x21, 0x8e,
	0x1f, 0x66, 0x62, 0x68, 0x98, 0x8a, 0xc1, 0x60, 0x9b, 0xf0, 0x82, 0x7d, 0x88, 0xe3, 0x6a, 0x30,
	0xde, 0xa0, 0x82, 0xbe, 0x21, 0x07, 0xdb, 0x06, 0x38, 0x20, 0xb0, 0x1d, 0x78, 0xc1, 0x60, 0x83,
	0x7f, 0x6d, 0x71, 0x7a, 0x44, 0x9a, 0x99, 0x12, 0xb1, 0x90, 0x7e, 0x1f, 0xe3, 0xe8, 0x24, 0x3b,
	0x29, 0xbf, 0x80, 0x88, 0xbf, 0x98, 0x30, 0xc5, 0xda, 0xaa, 0x05, 0xbb, 0xa5, 0x7c, 0x65, 0xc7,
	0xcc, 0x5d, 0xa9, 0x5a, 0xc4, 0x58, 0x47, 0xb5, 0xa0, 0x55, 0x02, 0x1f, 0x63, 0x9b, 0xd7, 0x42,
	0x72, 0xb1, 0x14, 0xbc, 0x60, 0x49, 0xc8, 0x8c, 0x61, 0xd1, 0xc2, 0xdf, 0x2f, 0xf3, 0xfa, 0x41,
	0x71, 0x86, 0x38, 0x7d, 0x43, 0xe8, 0x06, 0x99, 0xc3, 0x2d, 0x48, 0x0d, 0x3e, 0x2d, 0x7b, 0xea,
	0x83, 0xe6, 0xa2, 0x54, 0x3c, 0xa2, 0x57, 0xd9, 0x71, 0xf0, 0x98, 0x7e, 0xed, 0xf2, 0xc4, 0x4e,
	0xbc, 0x3c, 0x4c, 0x8b, 0xc4, 0x88, 0x3c, 0x11, 0xa0, 0xfc, 0x43, 0xf4, 0xb5, 0x1b, 0xe5, 0x57,
	0x6b, 0x8c, 0xfe, 0x8a, 0xb4, 0x6c, 0x0e, 0xce, 0x59, 0x92, 0xf8, 0xcf, 0xf0, 0xde, 0xbf, 0xdd,
	0xbc, 0xf7, 0x69, 0xf5, 0xec, 0x3a, 0xb1, 0x2d, 0x1b, 0x7f, 0xa6, 0x3c, 0x58, 0x9b, 0xd8, 0xeb,
	0xaf, 0x6a, 0x32, 0x8c, 0x20, 0x49, 0xec, 0xf5, 0x1f, 0x0d, 0xbd, 0xf1, 0x4e, 0xd0, 0xaf, 0xf0,
	0x09, 0x24, 0xc9, 0x94, 0xdb, 0x32, 0x9b, 0x33, 0x63, 0x12, 0xd0, 0x2e, 0x2a, 0xc0, 0xfd, 0x6f,
	0xca, 0x32, 0x73, 0xf8, 0x99, 0x83, 0x37, 0xa9, 0x18, 0x13, 0xfb, 0xc4, 0xf1, 0xb7, 0xa8, 0x17,
	0x0e, 0xb6, 0xfb, 0xdb, 0xdb, 0x16, 0x32, 0x2a, 0xe6, 0xcc, 0x64, 0xca, 0xee, 0xff, 0xbc, 0x4c,
	0x3f, 0x88, 0xe3, 0x69, 0x05, 0x4f, 0x91, 0xb9, 0x35, 0xcb, 0xed, 0x18, 0x3f, 0x76, 0x9e, 0x6e,
	0x8c, 0xf1, 0x2b, 0x4d, 0xbf, 0x25, 0x5d, 0x59, 0xa4, 0x61, 0x91, 0xc7, 0x8a, 0x71, 0xd0, 0xfe,
	0x0b, 0xdc, 0xba, 0x23, 0x8b, 0xf4, 0xb3, 0x83, 0xe8, 0x2f, 0x88, 0xcf, 0x38, 0x17, 0xd6, 0x88,
	0x25, 0xe1, 0x76, 0x94, 0xbf, 0x8f, 0x51, 0x3e, 0x7a, 0xd0, 0x4f, 0x36, 0xe3, 0x7d, 0x4c, 0x5a,
	0xb7, 0x6c, 0x99, 0x29, 0x61, 0xc0, 0x7f, 0x59, 0x3e, 0x51, 0x2a, 0xd9, 0xea, 0xa4, 0x88, 0x16,
	0x58, 0x1f, 0x3f, 0x28, 0x9f, 0x0e, 0x95, 0x6c, 0x93, 0xee, 0x56, 0x65, 0x29, 0xd6, 0x98, 0xff,
	0xca, 0x19, 0xaa, 0x2c, 0xb5, 0xc5, 0x65, 0x87, 0x73, 0xd9, 0xdd, 0x23, 0x26, 0xf9, 0x6a, 0xfd,
	0x2c, 0x1c, 0x96, 0x99, 0x81, 0xaa, 0x89, 0xd5, 0xb8, 0xa7, 0xe1, 0x79, 0xeb, 0xcf, 0x4d, 0x7c,
	0x40, 0xeb, 0xd9, 0xf7, 0x66, 0xde, 0xac, 0x36, 0x2f, 0xa5, 0x9f, 0xff, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x0d, 0x71, 0x4c, 0xf1, 0xa3, 0x0b, 0x00, 0x00,
}
