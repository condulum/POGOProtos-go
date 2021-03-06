// Code generated by protoc-gen-go.
// source: data_logs.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type BuddyPokemonLogEntry_Result int32

const (
	BuddyPokemonLogEntry_UNSET       BuddyPokemonLogEntry_Result = 0
	BuddyPokemonLogEntry_CANDY_FOUND BuddyPokemonLogEntry_Result = 1
)

var BuddyPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "CANDY_FOUND",
}
var BuddyPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":       0,
	"CANDY_FOUND": 1,
}

func (x BuddyPokemonLogEntry_Result) String() string {
	return proto.EnumName(BuddyPokemonLogEntry_Result_name, int32(x))
}
func (BuddyPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{1, 0}
}

type CatchPokemonLogEntry_Result int32

const (
	CatchPokemonLogEntry_UNSET            CatchPokemonLogEntry_Result = 0
	CatchPokemonLogEntry_POKEMON_CAPTURED CatchPokemonLogEntry_Result = 1
	CatchPokemonLogEntry_POKEMON_FLED     CatchPokemonLogEntry_Result = 2
	CatchPokemonLogEntry_POKEMON_HATCHED  CatchPokemonLogEntry_Result = 3
)

var CatchPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "POKEMON_CAPTURED",
	2: "POKEMON_FLED",
	3: "POKEMON_HATCHED",
}
var CatchPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":            0,
	"POKEMON_CAPTURED": 1,
	"POKEMON_FLED":     2,
	"POKEMON_HATCHED":  3,
}

func (x CatchPokemonLogEntry_Result) String() string {
	return proto.EnumName(CatchPokemonLogEntry_Result_name, int32(x))
}
func (CatchPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{2, 0}
}

type FortSearchLogEntry_Result int32

const (
	FortSearchLogEntry_UNSET   FortSearchLogEntry_Result = 0
	FortSearchLogEntry_SUCCESS FortSearchLogEntry_Result = 1
)

var FortSearchLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
}
var FortSearchLogEntry_Result_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
}

func (x FortSearchLogEntry_Result) String() string {
	return proto.EnumName(FortSearchLogEntry_Result_name, int32(x))
}
func (FortSearchLogEntry_Result) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3, 0} }

type ActionLogEntry struct {
	TimestampMs int64 `protobuf:"varint,1,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	Sfida       bool  `protobuf:"varint,2,opt,name=sfida" json:"sfida,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*ActionLogEntry_CatchPokemon
	//	*ActionLogEntry_FortSearch
	//	*ActionLogEntry_BuddyPokemon
	Action isActionLogEntry_Action `protobuf_oneof:"Action"`
}

func (m *ActionLogEntry) Reset()                    { *m = ActionLogEntry{} }
func (m *ActionLogEntry) String() string            { return proto.CompactTextString(m) }
func (*ActionLogEntry) ProtoMessage()               {}
func (*ActionLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isActionLogEntry_Action interface {
	isActionLogEntry_Action()
}

type ActionLogEntry_CatchPokemon struct {
	CatchPokemon *CatchPokemonLogEntry `protobuf:"bytes,3,opt,name=catch_pokemon,json=catchPokemon,oneof"`
}
type ActionLogEntry_FortSearch struct {
	FortSearch *FortSearchLogEntry `protobuf:"bytes,4,opt,name=fort_search,json=fortSearch,oneof"`
}
type ActionLogEntry_BuddyPokemon struct {
	BuddyPokemon *BuddyPokemonLogEntry `protobuf:"bytes,5,opt,name=buddy_pokemon,json=buddyPokemon,oneof"`
}

func (*ActionLogEntry_CatchPokemon) isActionLogEntry_Action() {}
func (*ActionLogEntry_FortSearch) isActionLogEntry_Action()   {}
func (*ActionLogEntry_BuddyPokemon) isActionLogEntry_Action() {}

func (m *ActionLogEntry) GetAction() isActionLogEntry_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ActionLogEntry) GetCatchPokemon() *CatchPokemonLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_CatchPokemon); ok {
		return x.CatchPokemon
	}
	return nil
}

func (m *ActionLogEntry) GetFortSearch() *FortSearchLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_FortSearch); ok {
		return x.FortSearch
	}
	return nil
}

func (m *ActionLogEntry) GetBuddyPokemon() *BuddyPokemonLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_BuddyPokemon); ok {
		return x.BuddyPokemon
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionLogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionLogEntry_OneofMarshaler, _ActionLogEntry_OneofUnmarshaler, _ActionLogEntry_OneofSizer, []interface{}{
		(*ActionLogEntry_CatchPokemon)(nil),
		(*ActionLogEntry_FortSearch)(nil),
		(*ActionLogEntry_BuddyPokemon)(nil),
	}
}

func _ActionLogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CatchPokemon); err != nil {
			return err
		}
	case *ActionLogEntry_FortSearch:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FortSearch); err != nil {
			return err
		}
	case *ActionLogEntry_BuddyPokemon:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuddyPokemon); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ActionLogEntry.Action has unexpected type %T", x)
	}
	return nil
}

func _ActionLogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActionLogEntry)
	switch tag {
	case 3: // Action.catch_pokemon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CatchPokemonLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_CatchPokemon{msg}
		return true, err
	case 4: // Action.fort_search
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FortSearchLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_FortSearch{msg}
		return true, err
	case 5: // Action.buddy_pokemon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuddyPokemonLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_BuddyPokemon{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ActionLogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		s := proto.Size(x.CatchPokemon)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionLogEntry_FortSearch:
		s := proto.Size(x.FortSearch)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionLogEntry_BuddyPokemon:
		s := proto.Size(x.BuddyPokemon)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BuddyPokemonLogEntry struct {
	Result    BuddyPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.BuddyPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Amount    int32                       `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *BuddyPokemonLogEntry) Reset()                    { *m = BuddyPokemonLogEntry{} }
func (m *BuddyPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*BuddyPokemonLogEntry) ProtoMessage()               {}
func (*BuddyPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type CatchPokemonLogEntry struct {
	Result        CatchPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.CatchPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId     PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	CombatPoints  int32                       `protobuf:"varint,3,opt,name=combat_points,json=combatPoints" json:"combat_points,omitempty"`
	PokemonDataId uint64                      `protobuf:"fixed64,4,opt,name=pokemon_data_id,json=pokemonDataId" json:"pokemon_data_id,omitempty"`
}

func (m *CatchPokemonLogEntry) Reset()                    { *m = CatchPokemonLogEntry{} }
func (m *CatchPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*CatchPokemonLogEntry) ProtoMessage()               {}
func (*CatchPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type FortSearchLogEntry struct {
	Result FortSearchLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.FortSearchLogEntry_Result" json:"result,omitempty"`
	FortId string                    `protobuf:"bytes,2,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	Items  []*ItemData               `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Eggs   int32                     `protobuf:"varint,4,opt,name=eggs" json:"eggs,omitempty"`
}

func (m *FortSearchLogEntry) Reset()                    { *m = FortSearchLogEntry{} }
func (m *FortSearchLogEntry) String() string            { return proto.CompactTextString(m) }
func (*FortSearchLogEntry) ProtoMessage()               {}
func (*FortSearchLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *FortSearchLogEntry) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionLogEntry)(nil), "POGOProtos.Data.Logs.ActionLogEntry")
	proto.RegisterType((*BuddyPokemonLogEntry)(nil), "POGOProtos.Data.Logs.BuddyPokemonLogEntry")
	proto.RegisterType((*CatchPokemonLogEntry)(nil), "POGOProtos.Data.Logs.CatchPokemonLogEntry")
	proto.RegisterType((*FortSearchLogEntry)(nil), "POGOProtos.Data.Logs.FortSearchLogEntry")
	proto.RegisterEnum("POGOProtos.Data.Logs.BuddyPokemonLogEntry_Result", BuddyPokemonLogEntry_Result_name, BuddyPokemonLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.CatchPokemonLogEntry_Result", CatchPokemonLogEntry_Result_name, CatchPokemonLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.FortSearchLogEntry_Result", FortSearchLogEntry_Result_name, FortSearchLogEntry_Result_value)
}

func init() { proto.RegisterFile("data_logs.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0x9b, 0x40,
	0x10, 0x0e, 0x76, 0x20, 0xc9, 0xe0, 0x1f, 0xb4, 0x45, 0xad, 0x95, 0x5e, 0x5c, 0x52, 0x55, 0x56,
	0x0f, 0x54, 0x75, 0x4f, 0xed, 0xcd, 0x06, 0x1c, 0xa3, 0x38, 0x36, 0x5d, 0xdb, 0x87, 0xf6, 0x82,
	0x30, 0x60, 0x07, 0x35, 0xb0, 0x16, 0xbb, 0xae, 0xe4, 0x77, 0xeb, 0xb9, 0x2f, 0xd1, 0x07, 0xe8,
	0x6b, 0x54, 0x2c, 0x60, 0x59, 0x29, 0xaa, 0x7c, 0xe8, 0x65, 0x35, 0xfb, 0x69, 0x66, 0xbe, 0x6f,
	0xbf, 0xdd, 0x1d, 0x68, 0x07, 0x1e, 0xf3, 0xdc, 0x47, 0xb2, 0xa1, 0xfa, 0x36, 0x25, 0x8c, 0x20,
	0xd5, 0x99, 0xdd, 0xce, 0x9c, 0x2c, 0xa4, 0xba, 0xe9, 0x31, 0x4f, 0x9f, 0x90, 0x0d, 0xbd, 0x96,
	0xc3, 0x64, 0x17, 0x17, 0x29, 0xd7, 0x6a, 0x94, 0x7c, 0x0f, 0x13, 0x46, 0xd2, 0xbd, 0x1b, 0xb1,
	0x30, 0xce, 0x51, 0xed, 0x67, 0x0d, 0x5a, 0x03, 0x9f, 0x45, 0x24, 0x99, 0x90, 0x8d, 0x95, 0xb0,
	0x74, 0x8f, 0x5e, 0x41, 0x83, 0x45, 0x71, 0x48, 0x99, 0x17, 0x6f, 0xdd, 0x98, 0x76, 0x84, 0xae,
	0xd0, 0xab, 0x63, 0xf9, 0x80, 0xdd, 0x53, 0xa4, 0x82, 0x48, 0xd7, 0x51, 0xe0, 0x75, 0x6a, 0x5d,
	0xa1, 0x77, 0x89, 0xf3, 0x0d, 0xfa, 0x0c, 0x4d, 0xdf, 0x63, 0xfe, 0x83, 0xbb, 0x25, 0xdf, 0xc2,
	0x98, 0x24, 0x9d, 0x7a, 0x57, 0xe8, 0xc9, 0xfd, 0xb7, 0x7a, 0x95, 0x38, 0xdd, 0xc8, 0x52, 0x9d,
	0x3c, 0xb3, 0xe4, 0x1e, 0x9f, 0xe1, 0x86, 0x7f, 0x84, 0xa3, 0x3b, 0x90, 0xd7, 0x24, 0x65, 0x2e,
	0x0d, 0xbd, 0xd4, 0x7f, 0xe8, 0x9c, 0xf3, 0x86, 0xbd, 0xea, 0x86, 0x23, 0x92, 0xb2, 0x39, 0xcf,
	0x3b, 0x6a, 0x07, 0xeb, 0x03, 0x9a, 0xe9, 0x5b, 0xed, 0x82, 0x60, 0x7f, 0xd0, 0x27, 0xfe, 0x4b,
	0xdf, 0x30, 0x4b, 0xad, 0xd0, 0xb7, 0x3a, 0xc2, 0x87, 0x97, 0x20, 0xe5, 0xee, 0x69, 0xbf, 0x04,
	0x50, 0xab, 0x4a, 0x90, 0x0d, 0x52, 0x1a, 0xd2, 0xdd, 0x23, 0xe3, 0x46, 0xb6, 0xfa, 0xef, 0x4f,
	0xa7, 0xd3, 0x31, 0x2f, 0xc4, 0x45, 0x03, 0xf4, 0x09, 0xa0, 0x90, 0xee, 0x46, 0x01, 0xf7, 0xbe,
	0xd5, 0x7f, 0x79, 0xdc, 0xce, 0xe2, 0xf7, 0x5d, 0x74, 0xb1, 0x03, 0x7c, 0xb5, 0x2d, 0x43, 0xf4,
	0x1c, 0x24, 0x2f, 0x26, 0xbb, 0x84, 0xf1, 0x5b, 0x11, 0x71, 0xb1, 0xd3, 0x5e, 0x83, 0x94, 0xb3,
	0xa0, 0x2b, 0x10, 0x97, 0xd3, 0xb9, 0xb5, 0x50, 0xce, 0x50, 0x1b, 0x64, 0x63, 0x30, 0x35, 0xbf,
	0xb8, 0xa3, 0xd9, 0x72, 0x6a, 0x2a, 0x82, 0xf6, 0xa3, 0x06, 0x6a, 0xd5, 0x85, 0x9d, 0x7a, 0xba,
	0xaa, 0xda, 0xff, 0x79, 0xba, 0x1b, 0x68, 0xfa, 0x24, 0x5e, 0x79, 0xcc, 0xdd, 0x92, 0x28, 0x61,
	0xb4, 0x38, 0x64, 0x23, 0x07, 0x1d, 0x8e, 0xa1, 0x37, 0xd0, 0x2e, 0x09, 0xf8, 0xff, 0x89, 0x02,
	0xfe, 0xa0, 0x24, 0xdc, 0x2c, 0xe0, 0x4c, 0xaf, 0x1d, 0x68, 0x4e, 0x95, 0x25, 0x2a, 0x28, 0xce,
	0xec, 0xce, 0xba, 0x9f, 0x4d, 0x5d, 0x63, 0xe0, 0x2c, 0x96, 0xd8, 0x32, 0x15, 0x01, 0x29, 0xd0,
	0x28, 0xd1, 0xd1, 0xc4, 0x32, 0x95, 0x1a, 0x7a, 0x06, 0xed, 0x12, 0x19, 0x0f, 0x16, 0xc6, 0xd8,
	0x32, 0x95, 0xba, 0xf6, 0x5b, 0x00, 0xf4, 0xf7, 0xf3, 0x44, 0xb7, 0x4f, 0xcc, 0x7b, 0x77, 0xea,
	0xc3, 0x7e, 0x6a, 0xdd, 0x0b, 0xb8, 0xe0, 0xdf, 0xa4, 0xf0, 0xed, 0x0a, 0x4b, 0xd9, 0xd6, 0x0e,
	0xd0, 0x47, 0x10, 0xb3, 0xcf, 0x9e, 0xf9, 0x51, 0xef, 0xc9, 0xfd, 0x9b, 0x63, 0x02, 0xbb, 0x9c,
	0x07, 0xba, 0x9d, 0xcd, 0x83, 0x6c, 0xc9, 0x38, 0x71, 0x5e, 0x81, 0x10, 0x9c, 0x87, 0x9b, 0x0d,
	0xe5, 0x16, 0x89, 0x98, 0xc7, 0x5a, 0xb7, 0xca, 0x19, 0x19, 0x2e, 0xe6, 0x4b, 0xc3, 0xb0, 0xe6,
	0x73, 0x45, 0x18, 0x5e, 0x7e, 0x95, 0xf8, 0x60, 0xa1, 0xce, 0x99, 0x23, 0xac, 0xf2, 0xf8, 0xc3,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x70, 0x30, 0x64, 0xb0, 0x04, 0x00, 0x00,
}
