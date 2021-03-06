// Code generated by protoc-gen-go.
// source: networking_platform.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlatformRequestType int32

const (
	PlatformRequestType_METHOD_UNSET             PlatformRequestType = 0
	PlatformRequestType_BUY_ITEM_POKECOINS       PlatformRequestType = 2
	PlatformRequestType_BUY_ITEM_ANDROID         PlatformRequestType = 3
	PlatformRequestType_BUY_ITEM_IOS             PlatformRequestType = 4
	PlatformRequestType_GET_STORE_ITEMS          PlatformRequestType = 5
	PlatformRequestType_SEND_ENCRYPTED_SIGNATURE PlatformRequestType = 6
)

var PlatformRequestType_name = map[int32]string{
	0: "METHOD_UNSET",
	2: "BUY_ITEM_POKECOINS",
	3: "BUY_ITEM_ANDROID",
	4: "BUY_ITEM_IOS",
	5: "GET_STORE_ITEMS",
	6: "SEND_ENCRYPTED_SIGNATURE",
}
var PlatformRequestType_value = map[string]int32{
	"METHOD_UNSET":             0,
	"BUY_ITEM_POKECOINS":       2,
	"BUY_ITEM_ANDROID":         3,
	"BUY_ITEM_IOS":             4,
	"GET_STORE_ITEMS":          5,
	"SEND_ENCRYPTED_SIGNATURE": 6,
}

func (x PlatformRequestType) String() string {
	return proto.EnumName(PlatformRequestType_name, int32(x))
}
func (PlatformRequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func init() {
	proto.RegisterEnum("POGOProtos.Networking.Platform.PlatformRequestType", PlatformRequestType_name, PlatformRequestType_value)
}

func init() { proto.RegisterFile("networking_platform.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xce, 0xcc, 0x4b, 0x8f, 0x2f, 0xc8, 0x49, 0x2c, 0x49, 0xcb, 0x2f, 0xca, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0b, 0xf0, 0x77, 0xf7, 0x0f, 0x00, 0x31, 0x8b, 0xf5,
	0xfc, 0xe0, 0xaa, 0xf4, 0x02, 0xa0, 0xaa, 0xb4, 0x66, 0x31, 0x72, 0x09, 0xc3, 0x38, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x21, 0x95, 0x05, 0xa9, 0x42, 0x02, 0x5c, 0x3c, 0xbe, 0xae, 0x21,
	0x1e, 0xfe, 0x2e, 0xf1, 0xa1, 0x7e, 0xc1, 0xae, 0x21, 0x02, 0x0c, 0x42, 0x62, 0x5c, 0x42, 0x4e,
	0xa1, 0x91, 0xf1, 0x9e, 0x21, 0xae, 0xbe, 0xf1, 0x01, 0xfe, 0xde, 0xae, 0xce, 0xfe, 0x9e, 0x7e,
	0xc1, 0x02, 0x4c, 0x42, 0x22, 0x5c, 0x02, 0x70, 0x71, 0x47, 0x3f, 0x97, 0x20, 0x7f, 0x4f, 0x17,
	0x01, 0x66, 0x90, 0x7e, 0xb8, 0xa8, 0xa7, 0x7f, 0xb0, 0x00, 0x8b, 0x90, 0x30, 0x17, 0xbf, 0xbb,
	0x6b, 0x48, 0x7c, 0x70, 0x88, 0x7f, 0x90, 0x2b, 0x58, 0x3c, 0x58, 0x80, 0x55, 0x48, 0x86, 0x4b,
	0x22, 0xd8, 0xd5, 0xcf, 0x25, 0xde, 0xd5, 0xcf, 0x39, 0x28, 0x32, 0x20, 0xc4, 0xd5, 0x25, 0x3e,
	0xd8, 0xd3, 0xdd, 0xcf, 0x31, 0x24, 0x34, 0xc8, 0x55, 0x80, 0xcd, 0x89, 0x23, 0x8a, 0x0d, 0xec,
	0x8b, 0xe2, 0x24, 0x08, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x25, 0x54, 0xf2, 0x7d, 0xea,
	0x00, 0x00, 0x00,
}
