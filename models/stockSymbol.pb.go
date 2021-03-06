// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/stockSymbol.proto

package models

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

type SymbolSearchResponse struct {
	Symbols              []*StockSymbol `protobuf:"bytes,1,rep,name=symbols,json=bestMatches,proto3" json:"symbols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SymbolSearchResponse) Reset()         { *m = SymbolSearchResponse{} }
func (m *SymbolSearchResponse) String() string { return proto.CompactTextString(m) }
func (*SymbolSearchResponse) ProtoMessage()    {}
func (*SymbolSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1678303e5b11f9f2, []int{0}
}

func (m *SymbolSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymbolSearchResponse.Unmarshal(m, b)
}
func (m *SymbolSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymbolSearchResponse.Marshal(b, m, deterministic)
}
func (m *SymbolSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymbolSearchResponse.Merge(m, src)
}
func (m *SymbolSearchResponse) XXX_Size() int {
	return xxx_messageInfo_SymbolSearchResponse.Size(m)
}
func (m *SymbolSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SymbolSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SymbolSearchResponse proto.InternalMessageInfo

func (m *SymbolSearchResponse) GetSymbols() []*StockSymbol {
	if m != nil {
		return m.Symbols
	}
	return nil
}

type StockSymbol struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,json=1. symbol,proto3" json:"symbol,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,json=2. name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,json=3. type,proto3" json:"type,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,json=4. region,proto3" json:"region,omitempty"`
	MarketOpen           string   `protobuf:"bytes,5,opt,name=marketOpen,json=5. marketOpen,proto3" json:"marketOpen,omitempty"`
	MarketClose          string   `protobuf:"bytes,6,opt,name=marketClose,json=6. marketClose,proto3" json:"marketClose,omitempty"`
	Timezone             string   `protobuf:"bytes,7,opt,name=timezone,json=7. timezone,proto3" json:"timezone,omitempty"`
	Currency             string   `protobuf:"bytes,8,opt,name=currency,json=8. currency,proto3" json:"currency,omitempty"`
	MatchScore           string   `protobuf:"bytes,9,opt,name=matchScore,json=9. matchScore,proto3" json:"matchScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockSymbol) Reset()         { *m = StockSymbol{} }
func (m *StockSymbol) String() string { return proto.CompactTextString(m) }
func (*StockSymbol) ProtoMessage()    {}
func (*StockSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_1678303e5b11f9f2, []int{1}
}

func (m *StockSymbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockSymbol.Unmarshal(m, b)
}
func (m *StockSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockSymbol.Marshal(b, m, deterministic)
}
func (m *StockSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockSymbol.Merge(m, src)
}
func (m *StockSymbol) XXX_Size() int {
	return xxx_messageInfo_StockSymbol.Size(m)
}
func (m *StockSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_StockSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_StockSymbol proto.InternalMessageInfo

func (m *StockSymbol) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *StockSymbol) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StockSymbol) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StockSymbol) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *StockSymbol) GetMarketOpen() string {
	if m != nil {
		return m.MarketOpen
	}
	return ""
}

func (m *StockSymbol) GetMarketClose() string {
	if m != nil {
		return m.MarketClose
	}
	return ""
}

func (m *StockSymbol) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *StockSymbol) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *StockSymbol) GetMatchScore() string {
	if m != nil {
		return m.MatchScore
	}
	return ""
}

func init() {
	proto.RegisterType((*SymbolSearchResponse)(nil), "models.SymbolSearchResponse")
	proto.RegisterType((*StockSymbol)(nil), "models.StockSymbol")
}

func init() { proto.RegisterFile("models/stockSymbol.proto", fileDescriptor_1678303e5b11f9f2) }

var fileDescriptor_1678303e5b11f9f2 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x49, 0x52, 0x92, 0xc6, 0x11, 0x0c, 0x06, 0x24, 0x33, 0x20, 0x99, 0xb2, 0x74, 0x32,
	0x6a, 0xcb, 0xef, 0x0a, 0x13, 0x03, 0x42, 0x4a, 0x9e, 0x20, 0x31, 0x57, 0xa4, 0x6a, 0x62, 0x47,
	0xb6, 0x19, 0xc2, 0xa3, 0xf0, 0x28, 0x3c, 0x1d, 0xb2, 0xaf, 0xa2, 0x76, 0x3c, 0xdf, 0x77, 0x72,
	0x72, 0x65, 0xc2, 0x7a, 0xfd, 0x09, 0x9d, 0xbd, 0xb5, 0x4e, 0xcb, 0x5d, 0x35, 0xf6, 0x8d, 0xee,
	0xc4, 0x60, 0xb4, 0xd3, 0x34, 0x45, 0xb3, 0x78, 0x23, 0xe7, 0xc8, 0x2b, 0xa8, 0x8d, 0x6c, 0x4b,
	0xb0, 0x83, 0x56, 0x16, 0xe8, 0x8a, 0x64, 0x36, 0x70, 0xcb, 0x22, 0x9e, 0x2c, 0x8b, 0xf5, 0x99,
	0xc0, 0x2f, 0x44, 0xb5, 0xdf, 0x2a, 0x8b, 0x06, 0xac, 0x7b, 0xaf, 0x9d, 0x6c, 0xc1, 0x2e, 0x7e,
	0x63, 0x52, 0x1c, 0x48, 0x7a, 0x49, 0x52, 0x9c, 0x60, 0x11, 0x8f, 0x96, 0x79, 0x99, 0xaf, 0x04,
	0x47, 0x40, 0x2f, 0xc8, 0x4c, 0xd5, 0x3d, 0xb0, 0x38, 0x88, 0x6c, 0x2d, 0xb8, 0x8f, 0x1e, 0xbb,
	0x71, 0x00, 0x96, 0x20, 0xde, 0x08, 0xee, 0xa3, 0x1f, 0x32, 0xf0, 0xb5, 0xd5, 0x8a, 0xcd, 0x70,
	0xe8, 0x4e, 0x70, 0x04, 0xf4, 0x9a, 0x90, 0xbe, 0x36, 0x3b, 0x70, 0x1f, 0x03, 0x28, 0x76, 0x1c,
	0xf4, 0xc9, 0xbd, 0xe0, 0x7b, 0x48, 0x6f, 0x48, 0x81, 0xe9, 0xb5, 0xd3, 0x16, 0x58, 0x1a, 0x3a,
	0xa7, 0x0f, 0x53, 0x27, 0x50, 0x7a, 0x45, 0xe6, 0x6e, 0xdb, 0xc3, 0x8f, 0x56, 0xc0, 0xb2, 0xd0,
	0x28, 0x1e, 0x05, 0x9f, 0x90, 0xd7, 0xf2, 0xdb, 0x18, 0x50, 0x72, 0x64, 0x73, 0xd4, 0x4f, 0x82,
	0x4f, 0x08, 0xaf, 0x70, 0xb2, 0xad, 0xa4, 0x36, 0xc0, 0x72, 0xbc, 0xe2, 0xd9, 0xff, 0x61, 0x82,
	0x2f, 0xc9, 0x5f, 0x7c, 0xd4, 0xa4, 0xe1, 0xed, 0x37, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85,
	0xf1, 0x7f, 0x07, 0x97, 0x01, 0x00, 0x00,
}
