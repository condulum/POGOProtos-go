// Code generated by protoc-gen-go.
// source: maps.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of MapPokemon from map_pokemon.proto

// Ignoring public import of NearbyPokemon from map_pokemon.proto

// Ignoring public import of WildPokemon from map_pokemon.proto

type MapObjectsStatus int32

const (
	MapObjectsStatus_UNSET_STATUS   MapObjectsStatus = 0
	MapObjectsStatus_SUCCESS        MapObjectsStatus = 1
	MapObjectsStatus_LOCATION_UNSET MapObjectsStatus = 2
)

var MapObjectsStatus_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "SUCCESS",
	2: "LOCATION_UNSET",
}
var MapObjectsStatus_value = map[string]int32{
	"UNSET_STATUS":   0,
	"SUCCESS":        1,
	"LOCATION_UNSET": 2,
}

func (x MapObjectsStatus) String() string {
	return proto.EnumName(MapObjectsStatus_name, int32(x))
}
func (MapObjectsStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type MapCell struct {
	// S2 geographic area that the cell covers (http://s2map.com/) (https://code.google.com/archive/p/s2-geometry-library/)
	S2CellId             uint64         `protobuf:"varint,1,opt,name=s2_cell_id,json=s2CellId" json:"s2_cell_id,omitempty"`
	CurrentTimestampMs   int64          `protobuf:"varint,2,opt,name=current_timestamp_ms,json=currentTimestampMs" json:"current_timestamp_ms,omitempty"`
	Forts                []*FortData    `protobuf:"bytes,3,rep,name=forts" json:"forts,omitempty"`
	SpawnPoints          []*SpawnPoint  `protobuf:"bytes,4,rep,name=spawn_points,json=spawnPoints" json:"spawn_points,omitempty"`
	DeletedObjects       []string       `protobuf:"bytes,6,rep,name=deleted_objects,json=deletedObjects" json:"deleted_objects,omitempty"`
	IsTruncatedList      bool           `protobuf:"varint,7,opt,name=is_truncated_list,json=isTruncatedList" json:"is_truncated_list,omitempty"`
	FortSummaries        []*FortSummary `protobuf:"bytes,8,rep,name=fort_summaries,json=fortSummaries" json:"fort_summaries,omitempty"`
	DecimatedSpawnPoints []*SpawnPoint  `protobuf:"bytes,9,rep,name=decimated_spawn_points,json=decimatedSpawnPoints" json:"decimated_spawn_points,omitempty"`
	// Pokemon within 2 steps or less.
	WildPokemons []*WildPokemon `protobuf:"bytes,5,rep,name=wild_pokemons,json=wildPokemons" json:"wild_pokemons,omitempty"`
	// Pokemon within 1 step or none.
	CatchablePokemons []*MapPokemon `protobuf:"bytes,10,rep,name=catchable_pokemons,json=catchablePokemons" json:"catchable_pokemons,omitempty"`
	// Pokemon farther away than 2 steps, but still in the area.
	NearbyPokemons []*NearbyPokemon `protobuf:"bytes,11,rep,name=nearby_pokemons,json=nearbyPokemons" json:"nearby_pokemons,omitempty"`
}

func (m *MapCell) Reset()                    { *m = MapCell{} }
func (m *MapCell) String() string            { return proto.CompactTextString(m) }
func (*MapCell) ProtoMessage()               {}
func (*MapCell) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *MapCell) GetForts() []*FortData {
	if m != nil {
		return m.Forts
	}
	return nil
}

func (m *MapCell) GetSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.SpawnPoints
	}
	return nil
}

func (m *MapCell) GetFortSummaries() []*FortSummary {
	if m != nil {
		return m.FortSummaries
	}
	return nil
}

func (m *MapCell) GetDecimatedSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.DecimatedSpawnPoints
	}
	return nil
}

func (m *MapCell) GetWildPokemons() []*WildPokemon {
	if m != nil {
		return m.WildPokemons
	}
	return nil
}

func (m *MapCell) GetCatchablePokemons() []*MapPokemon {
	if m != nil {
		return m.CatchablePokemons
	}
	return nil
}

func (m *MapCell) GetNearbyPokemons() []*NearbyPokemon {
	if m != nil {
		return m.NearbyPokemons
	}
	return nil
}

type SpawnPoint struct {
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *SpawnPoint) Reset()                    { *m = SpawnPoint{} }
func (m *SpawnPoint) String() string            { return proto.CompactTextString(m) }
func (*SpawnPoint) ProtoMessage()               {}
func (*SpawnPoint) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func init() {
	proto.RegisterType((*MapCell)(nil), "POGOProtos.Map.MapCell")
	proto.RegisterType((*SpawnPoint)(nil), "POGOProtos.Map.SpawnPoint")
	proto.RegisterEnum("POGOProtos.Map.MapObjectsStatus", MapObjectsStatus_name, MapObjectsStatus_value)
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0x51, 0x6b, 0xdb, 0x30,
	0x10, 0xc7, 0xeb, 0xa6, 0x6d, 0x92, 0x4b, 0xea, 0x24, 0xa2, 0x0c, 0x13, 0x3a, 0x30, 0x19, 0x63,
	0xa1, 0x0f, 0x61, 0xa4, 0xcf, 0x7b, 0xe8, 0xb2, 0xb4, 0x2b, 0x34, 0x89, 0x67, 0x3b, 0x0c, 0xf6,
	0x22, 0x14, 0x5b, 0xdd, 0xb4, 0xc9, 0xb6, 0xf0, 0x29, 0x84, 0x7e, 0xa5, 0x7d, 0xca, 0x61, 0x39,
	0x71, 0xda, 0x42, 0xd8, 0x8b, 0xb9, 0xbb, 0xff, 0xff, 0x7e, 0xba, 0x43, 0x32, 0x40, 0xc2, 0x14,
	0x8e, 0x54, 0x9e, 0xe9, 0x8c, 0xd8, 0xde, 0xe2, 0x6e, 0xe1, 0x15, 0x21, 0x8e, 0x66, 0x4c, 0xf5,
	0xed, 0x84, 0x29, 0xfa, 0x98, 0xe5, 0xba, 0xd4, 0xfb, 0xbd, 0x22, 0x57, 0xd9, 0x1f, 0x9e, 0x64,
	0x69, 0x59, 0x1a, 0xfc, 0x3d, 0x85, 0xfa, 0x8c, 0xa9, 0x09, 0x97, 0x92, 0x5c, 0x02, 0xe0, 0x98,
	0x46, 0x5c, 0x4a, 0x2a, 0x62, 0xc7, 0x72, 0xad, 0xe1, 0x89, 0xdf, 0xc0, 0x71, 0xa1, 0xdd, 0xc7,
	0xe4, 0x23, 0x5c, 0x44, 0xeb, 0x3c, 0xe7, 0xa9, 0xa6, 0x5a, 0x24, 0x1c, 0x35, 0x4b, 0x14, 0x4d,
	0xd0, 0x39, 0x76, 0xad, 0x61, 0xcd, 0x27, 0x5b, 0x2d, 0xdc, 0x49, 0x33, 0x24, 0xd7, 0x70, 0x5a,
	0x1c, 0x8e, 0x4e, 0xcd, 0xad, 0x0d, 0x5b, 0xe3, 0xb7, 0xa3, 0x97, 0xe3, 0x8d, 0x6e, 0x8b, 0xc9,
	0x8a, 0xcf, 0x17, 0xa6, 0x99, 0x5f, 0x7a, 0xc9, 0x27, 0x68, 0xa3, 0x62, 0x9b, 0x94, 0xaa, 0x4c,
	0xa4, 0x1a, 0x9d, 0x13, 0xd3, 0xdb, 0x7f, 0xdd, 0x1b, 0x14, 0x1e, 0xaf, 0xb0, 0xf8, 0x2d, 0xac,
	0x62, 0x24, 0x1f, 0xa0, 0x13, 0x73, 0xc9, 0x35, 0x8f, 0x69, 0xb6, 0xfa, 0xcd, 0x23, 0x8d, 0xce,
	0x99, 0x5b, 0x1b, 0x36, 0x7d, 0x7b, 0x5b, 0x5e, 0x94, 0x55, 0x72, 0x05, 0x3d, 0x81, 0x54, 0xe7,
	0xeb, 0x34, 0x62, 0x85, 0x5b, 0x0a, 0xd4, 0x4e, 0xdd, 0xb5, 0x86, 0x0d, 0xbf, 0x23, 0x30, 0xdc,
	0xd5, 0x1f, 0x04, 0x6a, 0x72, 0x07, 0x76, 0x31, 0x1c, 0xc5, 0x75, 0x92, 0xb0, 0x5c, 0x70, 0x74,
	0x1a, 0x66, 0x2a, 0xf7, 0xe0, 0x46, 0x81, 0x71, 0x3e, 0xf9, 0xe7, 0x8f, 0x55, 0x22, 0x38, 0x12,
	0x0f, 0xde, 0xc4, 0x3c, 0x12, 0x89, 0x39, 0xf1, 0xc5, 0x9a, 0xcd, 0xff, 0xae, 0x79, 0x51, 0x75,
	0x06, 0xcf, 0xf6, 0xfd, 0x0a, 0xe7, 0x1b, 0x21, 0xe3, 0xdd, 0xad, 0xa2, 0x73, 0x6a, 0x40, 0xef,
	0x5e, 0x83, 0xbc, 0xed, 0xad, 0x7f, 0x17, 0x32, 0xde, 0xc6, 0x7e, 0x7b, 0xb3, 0x4f, 0x90, 0x7c,
	0x03, 0x12, 0x31, 0x1d, 0xfd, 0x62, 0x2b, 0xc9, 0xf7, 0x38, 0x30, 0xb8, 0xc1, 0x21, 0xdc, 0x8c,
	0xa9, 0x1d, 0xad, 0x57, 0x75, 0x57, 0xc8, 0x39, 0x74, 0x52, 0xce, 0xf2, 0xd5, 0xd3, 0x9e, 0xd7,
	0x32, 0xbc, 0xf7, 0x87, 0x78, 0x73, 0x63, 0xdf, 0x21, 0xed, 0xf4, 0x79, 0x8a, 0x83, 0x5b, 0x80,
	0xfd, 0xee, 0xa4, 0x0f, 0x0d, 0xc9, 0xb4, 0xd0, 0xeb, 0x98, 0x9b, 0x47, 0x68, 0xf9, 0x55, 0x4e,
	0x2e, 0xa1, 0x29, 0xb3, 0xf4, 0x67, 0x29, 0xd6, 0x8c, 0xb8, 0x2f, 0x5c, 0x4d, 0xa1, 0x3b, 0x63,
	0x6a, 0xfb, 0x12, 0x02, 0xcd, 0xf4, 0x1a, 0x49, 0x17, 0xda, 0xcb, 0x79, 0x30, 0x0d, 0x69, 0x10,
	0xde, 0x84, 0xcb, 0xa0, 0x7b, 0x44, 0x5a, 0x50, 0x0f, 0x96, 0x93, 0xc9, 0x34, 0x08, 0xba, 0x16,
	0x21, 0x60, 0x3f, 0x2c, 0x26, 0x37, 0xe1, 0xfd, 0x62, 0x4e, 0x8d, 0xaf, 0x7b, 0xfc, 0xb9, 0xf1,
	0xe3, 0xcc, 0xfc, 0x44, 0xe8, 0x1d, 0x79, 0xd6, 0xaa, 0x8c, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0x27, 0x2c, 0xed, 0x91, 0x03, 0x00, 0x00,
}
