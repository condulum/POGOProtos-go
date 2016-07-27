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

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

type CatchPokemonLogEntry_Result int32

const (
	CatchPokemonLogEntry_UNSET            CatchPokemonLogEntry_Result = 0
	CatchPokemonLogEntry_POKEMON_CAPTURED CatchPokemonLogEntry_Result = 1
	CatchPokemonLogEntry_POKEMON_FLED     CatchPokemonLogEntry_Result = 2
)

var CatchPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "POKEMON_CAPTURED",
	2: "POKEMON_FLED",
}
var CatchPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":            0,
	"POKEMON_CAPTURED": 1,
	"POKEMON_FLED":     2,
}

func (x CatchPokemonLogEntry_Result) String() string {
	return proto.EnumName(CatchPokemonLogEntry_Result_name, int32(x))
}
func (CatchPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor12, []int{1, 0}
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
func (FortSearchLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor12, []int{2, 0}
}

type ActionLogEntry struct {
	TimestampMs int64 `protobuf:"varint,1,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	Sfida       bool  `protobuf:"varint,2,opt,name=sfida" json:"sfida,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*ActionLogEntry_CatchPokemon
	//	*ActionLogEntry_FortSearch
	Action isActionLogEntry_Action `protobuf_oneof:"Action"`
}

func (m *ActionLogEntry) Reset()                    { *m = ActionLogEntry{} }
func (m *ActionLogEntry) String() string            { return proto.CompactTextString(m) }
func (*ActionLogEntry) ProtoMessage()               {}
func (*ActionLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

type isActionLogEntry_Action interface {
	isActionLogEntry_Action()
}

type ActionLogEntry_CatchPokemon struct {
	CatchPokemon *CatchPokemonLogEntry `protobuf:"bytes,3,opt,name=catch_pokemon,json=catchPokemon,oneof"`
}
type ActionLogEntry_FortSearch struct {
	FortSearch *FortSearchLogEntry `protobuf:"bytes,4,opt,name=fort_search,json=fortSearch,oneof"`
}

func (*ActionLogEntry_CatchPokemon) isActionLogEntry_Action() {}
func (*ActionLogEntry_FortSearch) isActionLogEntry_Action()   {}

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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionLogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionLogEntry_OneofMarshaler, _ActionLogEntry_OneofUnmarshaler, _ActionLogEntry_OneofSizer, []interface{}{
		(*ActionLogEntry_CatchPokemon)(nil),
		(*ActionLogEntry_FortSearch)(nil),
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
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CatchPokemonLogEntry struct {
	Result        CatchPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.CatchPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId     PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	CombatPoints  int32                       `protobuf:"varint,3,opt,name=combat_points,json=combatPoints" json:"combat_points,omitempty"`
	PokemonDataId uint64                      `protobuf:"varint,4,opt,name=pokemon_data_id,json=pokemonDataId" json:"pokemon_data_id,omitempty"`
}

func (m *CatchPokemonLogEntry) Reset()                    { *m = CatchPokemonLogEntry{} }
func (m *CatchPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*CatchPokemonLogEntry) ProtoMessage()               {}
func (*CatchPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

type FortSearchLogEntry struct {
	Result FortSearchLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=POGOProtos.Data.Logs.FortSearchLogEntry_Result" json:"result,omitempty"`
	FortId string                    `protobuf:"bytes,2,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	Items  []*ItemData               `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Eggs   int32                     `protobuf:"varint,4,opt,name=eggs" json:"eggs,omitempty"`
}

func (m *FortSearchLogEntry) Reset()                    { *m = FortSearchLogEntry{} }
func (m *FortSearchLogEntry) String() string            { return proto.CompactTextString(m) }
func (*FortSearchLogEntry) ProtoMessage()               {}
func (*FortSearchLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *FortSearchLogEntry) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionLogEntry)(nil), "POGOProtos.Data.Logs.ActionLogEntry")
	proto.RegisterType((*CatchPokemonLogEntry)(nil), "POGOProtos.Data.Logs.CatchPokemonLogEntry")
	proto.RegisterType((*FortSearchLogEntry)(nil), "POGOProtos.Data.Logs.FortSearchLogEntry")
	proto.RegisterEnum("POGOProtos.Data.Logs.CatchPokemonLogEntry_Result", CatchPokemonLogEntry_Result_name, CatchPokemonLogEntry_Result_value)
	proto.RegisterEnum("POGOProtos.Data.Logs.FortSearchLogEntry_Result", FortSearchLogEntry_Result_name, FortSearchLogEntry_Result_value)
}

func init() { proto.RegisterFile("data_logs.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0x9b, 0x40,
	0x10, 0xc7, 0x8d, 0x6d, 0x48, 0x3c, 0x38, 0x0e, 0x5a, 0x59, 0xaa, 0x95, 0x56, 0x55, 0x4a, 0xa4,
	0x2a, 0xaa, 0x54, 0xac, 0xba, 0xa7, 0xb6, 0xa7, 0xc4, 0x21, 0x29, 0xca, 0x87, 0xe9, 0x12, 0x5f,
	0x7a, 0x41, 0x18, 0x30, 0x46, 0x09, 0x2c, 0x62, 0xd7, 0x95, 0xf2, 0x44, 0x7d, 0xb4, 0x3e, 0x40,
	0x5f, 0xa0, 0xbb, 0x0b, 0x44, 0x56, 0xea, 0x83, 0x2f, 0x68, 0xf6, 0x3f, 0xb3, 0x33, 0xb3, 0xbf,
	0x19, 0xe0, 0x30, 0x0a, 0x58, 0xe0, 0x3f, 0x92, 0x84, 0x5a, 0x45, 0x49, 0x18, 0x41, 0x43, 0x77,
	0x76, 0x35, 0x73, 0x85, 0x49, 0xad, 0x0b, 0xee, 0xb3, 0x6e, 0xb8, 0xef, 0x48, 0x8f, 0xf3, 0x75,
	0x56, 0x87, 0x1c, 0x0d, 0xd3, 0xfc, 0x57, 0x9c, 0x33, 0x52, 0x3e, 0xf9, 0x29, 0x8b, 0xb3, 0x4a,
	0x35, 0xff, 0x2a, 0x30, 0x38, 0x0b, 0x59, 0x4a, 0x72, 0x7e, 0xc3, 0xce, 0x59, 0xf9, 0x84, 0xde,
	0x41, 0x9f, 0xa5, 0x59, 0x4c, 0x59, 0x90, 0x15, 0x7e, 0x46, 0x47, 0xca, 0xb1, 0x72, 0xda, 0xc1,
	0xfa, 0xb3, 0x76, 0x4b, 0xd1, 0x10, 0x54, 0xba, 0x4c, 0xa3, 0x60, 0xd4, 0xe6, 0xbe, 0x7d, 0x5c,
	0x1d, 0xd0, 0x0f, 0x38, 0x08, 0x03, 0x16, 0xae, 0xfc, 0x82, 0x3c, 0xc4, 0x19, 0xc9, 0x47, 0x1d,
	0xee, 0xd5, 0x27, 0x1f, 0xac, 0x6d, 0xcd, 0x59, 0x53, 0x11, 0xea, 0x56, 0x91, 0x4d, 0xed, 0xef,
	0x2d, 0xdc, 0x0f, 0x37, 0x74, 0x74, 0x0d, 0xfa, 0x92, 0x94, 0xcc, 0xa7, 0x71, 0x50, 0x86, 0xab,
	0x51, 0x57, 0x26, 0x3c, 0xdd, 0x9e, 0xf0, 0x92, 0x07, 0x7a, 0x32, 0x6e, 0x23, 0x1d, 0x2c, 0x9f,
	0xd5, 0xf3, 0x7d, 0xd0, 0xaa, 0xa7, 0x9a, 0xbf, 0xdb, 0x30, 0xdc, 0x56, 0x1f, 0x39, 0xa0, 0x95,
	0x31, 0x5d, 0x3f, 0x32, 0xf9, 0xea, 0xc1, 0xe4, 0xd3, 0xee, 0xbd, 0x5b, 0x58, 0x5e, 0xc4, 0x75,
	0x02, 0xf4, 0x15, 0xa0, 0xe6, 0xe0, 0xa7, 0x91, 0x04, 0x35, 0x98, 0xbc, 0xde, 0x4c, 0x67, 0xcb,
	0xe1, 0xd4, 0x59, 0x9c, 0x08, 0xf7, 0x8a, 0xc6, 0x44, 0x27, 0x9c, 0x24, 0xc9, 0x16, 0x01, 0xe3,
	0x28, 0xd3, 0x9c, 0x51, 0x49, 0x52, 0xe5, 0x6c, 0xa4, 0xe8, 0x4a, 0x0d, 0xbd, 0x87, 0xc3, 0xa6,
	0x80, 0x5c, 0x07, 0x5e, 0x45, 0xf0, 0xe9, 0xe2, 0x83, 0x5a, 0x16, 0xfd, 0x3a, 0x91, 0xf9, 0x0d,
	0xb4, 0xaa, 0x35, 0xd4, 0x03, 0x75, 0x7e, 0xe7, 0xd9, 0xf7, 0x46, 0x8b, 0x4f, 0xd0, 0x70, 0x67,
	0xd7, 0xf6, 0xed, 0xec, 0xce, 0x9f, 0x9e, 0xb9, 0xf7, 0x73, 0x6c, 0x5f, 0x18, 0x0a, 0x32, 0xa0,
	0xdf, 0xa8, 0x97, 0x37, 0x5c, 0x69, 0x9b, 0x7f, 0x14, 0x40, 0xff, 0x83, 0x45, 0x57, 0x2f, 0x38,
	0x8d, 0x77, 0x1d, 0xc9, 0x4b, 0x4a, 0xaf, 0x60, 0x4f, 0x0e, 0xb8, 0x46, 0xd4, 0xc3, 0x9a, 0x38,
	0x72, 0x04, 0x5f, 0x40, 0x15, 0x6b, 0x2a, 0x9e, 0xde, 0xe1, 0x33, 0x3f, 0xd9, 0x2c, 0xe0, 0x34,
	0x9b, 0x6c, 0x39, 0x62, 0x93, 0xc5, 0x47, 0xd4, 0xc4, 0xd5, 0x0d, 0x84, 0xa0, 0x1b, 0x27, 0x09,
	0x95, 0x34, 0x54, 0x2c, 0x6d, 0xf3, 0x78, 0x1b, 0x04, 0x1d, 0xf6, 0xbc, 0xf9, 0x74, 0x6a, 0x7b,
	0x9e, 0xa1, 0x9c, 0xbf, 0xfd, 0xf9, 0x26, 0x49, 0xd9, 0x6a, 0xbd, 0xb0, 0x38, 0xe5, 0x71, 0xf1,
	0x90, 0xe5, 0x09, 0xf9, 0x48, 0xa2, 0x74, 0x2c, 0x7f, 0x14, 0xea, 0xb6, 0x5c, 0x65, 0xa1, 0x49,
	0xfb, 0xf3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x82, 0x9c, 0x05, 0x80, 0x03, 0x00, 0x00,
}