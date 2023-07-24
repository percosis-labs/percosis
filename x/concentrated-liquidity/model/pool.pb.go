// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: percosis/concentrated-liquidity/pool.proto

// This is a legacy package that requires additional migration logic
// in order to use the correct package. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	github_com_osmosis_labs_osmosis_osmomath "github.com/percosis-labs/percosis/osmomath v0.0.3-dev.fury"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Pool struct {
	// pool's address holding all liquidity tokens.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// address holding the incentives liquidity.
	IncentivesAddress string `protobuf:"bytes,2,opt,name=incentives_address,json=incentivesAddress,proto3" json:"incentives_address,omitempty" yaml:"incentives_address"`
	// address holding spread rewards from swaps.
	SpreadRewardsAddress string `protobuf:"bytes,3,opt,name=spread_rewards_address,json=spreadRewardsAddress,proto3" json:"spread_rewards_address,omitempty" yaml:"spread_rewards_address"`
	Id                   uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	// Amount of total liquidity
	CurrentTickLiquidity github_com_cosmos_cosmos_sdk_types.Dec          `protobuf:"bytes,5,opt,name=current_tick_liquidity,json=currentTickLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_tick_liquidity" yaml:"current_tick_liquidity"`
	Token0               string                                          `protobuf:"bytes,6,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1               string                                          `protobuf:"bytes,7,opt,name=token1,proto3" json:"token1,omitempty"`
	CurrentSqrtPrice     github_com_osmosis_labs_osmosis_osmomath.BigDec `protobuf:"bytes,8,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=github.com/percosis-labs/percosis/osmomath v0.0.3-dev.fury.BigDec" json:"current_sqrt_price" yaml:"spot_price"`
	CurrentTick          int64                                           `protobuf:"varint,9,opt,name=current_tick,json=currentTick,proto3" json:"current_tick,omitempty" yaml:"current_tick"`
	// tick_spacing must be one of the authorized_tick_spacing values set in the
	// concentrated-liquidity parameters
	TickSpacing        uint64 `protobuf:"varint,10,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty" yaml:"tick_spacing"`
	ExponentAtPriceOne int64  `protobuf:"varint,11,opt,name=exponent_at_price_one,json=exponentAtPriceOne,proto3" json:"exponent_at_price_one,omitempty" yaml:"exponent_at_price_one"`
	// spread_factor is the ratio that is charged on the amount of token in.
	SpreadFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=spread_factor,json=spreadFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"spread_factor" yaml:"spread_factor"`
	// last_liquidity_update is the last time either the pool liquidity or the
	// active tick changed
	LastLiquidityUpdate time.Time `protobuf:"bytes,13,opt,name=last_liquidity_update,json=lastLiquidityUpdate,proto3,stdtime" json:"last_liquidity_update" yaml:"last_liquidity_update"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe62f04765626a6, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pool)(nil), "percosis.concentratedliquidity.v1beta1.Pool")
}

func init() {
	proto.RegisterFile("percosis/concentrated-liquidity/pool.proto", fileDescriptor_fbe62f04765626a6)
}

var fileDescriptor_fbe62f04765626a6 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0x8e, 0xf9, 0xfb, 0x63, 0x03, 0xe8, 0xc7, 0x12, 0xa8, 0x41, 0x25, 0x4e, 0x7d, 0x40, 0x51,
	0xd5, 0xd8, 0x4d, 0x7b, 0xe3, 0x54, 0xa2, 0x0a, 0xa9, 0x12, 0x12, 0xc8, 0x50, 0x55, 0xe2, 0x62,
	0x6d, 0xec, 0xc5, 0xac, 0x62, 0x7b, 0xcd, 0xee, 0x86, 0x42, 0xef, 0x95, 0x7a, 0xe4, 0xd8, 0x23,
	0x0f, 0xd1, 0x87, 0x40, 0x3d, 0x71, 0xac, 0x7a, 0x70, 0x2b, 0x22, 0xf5, 0x01, 0xf2, 0x04, 0x95,
	0x77, 0xed, 0xc4, 0x15, 0xe9, 0xa1, 0x27, 0x7b, 0xbe, 0x99, 0xf9, 0xf6, 0x9b, 0xd9, 0x99, 0x05,
	0x4f, 0x13, 0xcc, 0x3c, 0xca, 0x09, 0xb7, 0x3d, 0x1a, 0x7b, 0x38, 0x16, 0x0c, 0x09, 0xec, 0xb7,
	0x42, 0x72, 0xde, 0x27, 0x3e, 0x11, 0x57, 0x76, 0x42, 0x69, 0x68, 0x25, 0x8c, 0x0a, 0x0a, 0xb7,
	0x8b, 0x58, 0xab, 0x1c, 0x3b, 0x0a, 0xb5, 0x2e, 0xda, 0x5d, 0x2c, 0x50, 0x7b, 0x73, 0xc3, 0xa3,
	0x3c, 0xa2, 0xdc, 0x95, 0x59, 0xb6, 0x32, 0x14, 0xc5, 0x66, 0x2d, 0xa0, 0x01, 0x55, 0x78, 0xf6,
	0x97, 0xa3, 0x46, 0x40, 0x69, 0x10, 0x62, 0x5b, 0x5a, 0xdd, 0xfe, 0xa9, 0x2d, 0x48, 0x84, 0xb9,
	0x40, 0x51, 0xa2, 0x02, 0xcc, 0x5f, 0xf3, 0x60, 0xe6, 0x90, 0xd2, 0x10, 0x3e, 0x03, 0xf3, 0xc8,
	0xf7, 0x19, 0xe6, 0x5c, 0xd7, 0x1a, 0x5a, 0x73, 0xa1, 0x03, 0x87, 0xa9, 0xb1, 0x7c, 0x85, 0xa2,
	0x70, 0xc7, 0xcc, 0x1d, 0xa6, 0x53, 0x84, 0xc0, 0x7d, 0x00, 0x89, 0x14, 0x4a, 0x2e, 0x30, 0x77,
	0x8b, 0xc4, 0x29, 0x99, 0xb8, 0x35, 0x4c, 0x8d, 0x0d, 0x95, 0xf8, 0x30, 0xc6, 0x74, 0x56, 0xc6,
	0xe0, 0x6e, 0xce, 0xf6, 0x0e, 0xac, 0xf3, 0x84, 0x61, 0xe4, 0xbb, 0x0c, 0xbf, 0x47, 0xcc, 0x1f,
	0x33, 0x4e, 0x4b, 0xc6, 0x27, 0xc3, 0xd4, 0xd8, 0x52, 0x8c, 0x93, 0xe3, 0x4c, 0xa7, 0xa6, 0x1c,
	0x8e, 0xc2, 0x0b, 0xe2, 0x65, 0x30, 0x45, 0x7c, 0x7d, 0xa6, 0xa1, 0x35, 0x67, 0x9c, 0x29, 0xe2,
	0xc3, 0x8f, 0x1a, 0x58, 0xf7, 0xfa, 0x8c, 0xe1, 0x58, 0xb8, 0x82, 0x78, 0x3d, 0x77, 0xd4, 0x62,
	0x7d, 0x56, 0x9e, 0x74, 0x70, 0x9b, 0x1a, 0x95, 0xef, 0xa9, 0xb1, 0x1d, 0x10, 0x71, 0xd6, 0xef,
	0x5a, 0x1e, 0x8d, 0xf2, 0x36, 0xe7, 0x9f, 0x16, 0xf7, 0x7b, 0xb6, 0xb8, 0x4a, 0x30, 0xb7, 0x5e,
	0x63, 0x6f, 0xac, 0x6b, 0x32, 0xab, 0xe9, 0xd4, 0x72, 0xc7, 0x31, 0xf1, 0x7a, 0xfb, 0x05, 0x0c,
	0xd7, 0xc1, 0x9c, 0xa0, 0x3d, 0x1c, 0x3f, 0xd7, 0xe7, 0xb2, 0x63, 0x9d, 0xdc, 0x1a, 0xe1, 0x6d,
	0x7d, 0xbe, 0x84, 0xb7, 0xe1, 0x07, 0x00, 0x8b, 0x03, 0xf8, 0x39, 0x13, 0x6e, 0xc2, 0x88, 0x87,
	0xf5, 0xff, 0xa4, 0xe4, 0xfd, 0x5c, 0xb2, 0x5d, 0x92, 0x2c, 0xa5, 0x12, 0xde, 0x0a, 0x51, 0x97,
	0x17, 0x86, 0xfc, 0x46, 0x48, 0x9c, 0x59, 0x1d, 0x12, 0x28, 0xed, 0x2b, 0x45, 0x4f, 0x69, 0x4e,
	0x69, 0x3a, 0xff, 0xe7, 0xe7, 0x1c, 0x9d, 0x33, 0x71, 0x98, 0x41, 0x70, 0x07, 0x2c, 0x96, 0x8b,
	0xd3, 0x17, 0x1a, 0x5a, 0x73, 0xba, 0xf3, 0x68, 0x98, 0x1a, 0xab, 0x0f, 0x4b, 0x37, 0x9d, 0x6a,
	0xa9, 0xe0, 0x2c, 0x57, 0x36, 0x84, 0x27, 0xc8, 0x23, 0x71, 0xa0, 0x83, 0xec, 0x26, 0xca, 0xb9,
	0x65, 0xaf, 0xe9, 0x54, 0x33, 0xf3, 0x48, 0x59, 0xf0, 0x08, 0xac, 0xe1, 0xcb, 0x84, 0xc6, 0x19,
	0x35, 0xca, 0xf5, 0xb9, 0x34, 0xc6, 0x7a, 0x55, 0x0a, 0x68, 0x0c, 0x53, 0xe3, 0xb1, 0x22, 0x99,
	0x18, 0x66, 0x3a, 0xb0, 0xc0, 0x77, 0x55, 0x25, 0x07, 0x31, 0x86, 0x3d, 0xb0, 0x94, 0x4f, 0xd0,
	0x29, 0xf2, 0x04, 0x65, 0xfa, 0xa2, 0xec, 0xe1, 0xde, 0x3f, 0x5f, 0x7b, 0xed, 0x8f, 0x71, 0x54,
	0x64, 0xa6, 0xb3, 0xa8, 0xec, 0x3d, 0x69, 0xc2, 0x4b, 0xb0, 0x16, 0x22, 0x2e, 0xc6, 0xe3, 0xe0,
	0xf6, 0x13, 0x1f, 0x09, 0xac, 0x2f, 0x35, 0xb4, 0x66, 0xf5, 0xc5, 0xa6, 0xa5, 0x96, 0xd3, 0x2a,
	0x96, 0xd3, 0x3a, 0x2e, 0x96, 0xb3, 0xd3, 0xcc, 0x04, 0x8d, 0x2b, 0x9c, 0x48, 0x63, 0x5e, 0xff,
	0x30, 0x34, 0x67, 0x35, 0xf3, 0x8d, 0x26, 0xeb, 0xad, 0xf4, 0xec, 0xac, 0x7c, 0xba, 0x31, 0x2a,
	0x9f, 0x6f, 0x8c, 0xca, 0xd7, 0x2f, 0xad, 0xd9, 0x6c, 0xbd, 0xdf, 0x74, 0x4e, 0x6e, 0xef, 0xeb,
	0xda, 0xdd, 0x7d, 0x5d, 0xfb, 0x79, 0x5f, 0xd7, 0xae, 0x07, 0xf5, 0xca, 0xdd, 0xa0, 0x5e, 0xf9,
	0x36, 0xa8, 0x57, 0x4e, 0x5e, 0x95, 0x8a, 0x2e, 0xde, 0x21, 0x35, 0x39, 0xa3, 0x17, 0xec, 0xf2,
	0x6f, 0x6f, 0x58, 0x44, 0x7d, 0x1c, 0x76, 0xe7, 0x64, 0x05, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xec, 0xe4, 0xcc, 0x42, 0xf3, 0x04, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastLiquidityUpdate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x6a
	{
		size := m.SpreadFactor.Size()
		i -= size
		if _, err := m.SpreadFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.ExponentAtPriceOne != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.ExponentAtPriceOne))
		i--
		dAtA[i] = 0x58
	}
	if m.TickSpacing != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x50
	}
	if m.CurrentTick != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CurrentTick))
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.CurrentTickLiquidity.Size()
		i -= size
		if _, err := m.CurrentTickLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SpreadRewardsAddress) > 0 {
		i -= len(m.SpreadRewardsAddress)
		copy(dAtA[i:], m.SpreadRewardsAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.SpreadRewardsAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IncentivesAddress) > 0 {
		i -= len(m.IncentivesAddress)
		copy(dAtA[i:], m.IncentivesAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.IncentivesAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.IncentivesAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.SpreadRewardsAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = m.CurrentTickLiquidity.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.CurrentTick != 0 {
		n += 1 + sovPool(uint64(m.CurrentTick))
	}
	if m.TickSpacing != 0 {
		n += 1 + sovPool(uint64(m.TickSpacing))
	}
	if m.ExponentAtPriceOne != 0 {
		n += 1 + sovPool(uint64(m.ExponentAtPriceOne))
	}
	l = m.SpreadFactor.Size()
	n += 1 + l + sovPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate)
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentivesAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadRewardsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpreadRewardsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTickLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTickLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			m.CurrentTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentAtPriceOne", wireType)
			}
			m.ExponentAtPriceOne = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExponentAtPriceOne |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpreadFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLiquidityUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastLiquidityUpdate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPool
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPool
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
