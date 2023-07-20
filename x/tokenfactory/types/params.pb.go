// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: percosis/tokenfactory/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the tokenfactory module.
type Params struct {
	// DenomCreationFee defines the fee to be charged on the creation of a new
	// denom. The fee is drawn from the MsgCreateDenom's sender account, and
	// transferred to the community pool.
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
	// DenomCreationGasConsume defines the gas cost for creating a new denom.
	// This is intended as a spam deterrence mechanism.
	//
	// See: https://github.com/CosmWasm/token-factory/issues/11
	DenomCreationGasConsume uint64 `protobuf:"varint,2,opt,name=denom_creation_gas_consume,json=denomCreationGasConsume,proto3" json:"denom_creation_gas_consume,omitempty" yaml:"denom_creation_gas_consume"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad789f846e084053, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func (m *Params) GetDenomCreationGasConsume() uint64 {
	if m != nil {
		return m.DenomCreationGasConsume
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "percosis.tokenfactory.v1beta1.Params")
}

func init() {
	proto.RegisterFile("percosis/tokenfactory/v1beta1/params.proto", fileDescriptor_ad789f846e084053)
}

var fileDescriptor_ad789f846e084053 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5b, 0x34, 0x1c, 0xea, 0xc5, 0x34, 0x26, 0x02, 0x89, 0x5b, 0xec, 0x09, 0x4d, 0xe8,
	0x06, 0x8d, 0x1e, 0x3c, 0x42, 0xa2, 0x27, 0x12, 0xc2, 0xd1, 0x4b, 0x33, 0x5d, 0x96, 0xd2, 0x40,
	0x3b, 0x4d, 0x77, 0x31, 0xf6, 0x11, 0xbc, 0x79, 0xf2, 0x21, 0x7c, 0x12, 0x8e, 0x1c, 0x3d, 0x55,
	0x03, 0x6f, 0xc0, 0x13, 0x18, 0xda, 0x85, 0x80, 0x1a, 0x4f, 0xed, 0x64, 0xfe, 0xff, 0x9b, 0x7f,
	0x76, 0x8c, 0xcb, 0x98, 0x27, 0x0c, 0x45, 0x20, 0xa8, 0xc4, 0x31, 0x8f, 0x86, 0xc0, 0x24, 0x26,
	0x29, 0x7d, 0x6a, 0x79, 0x5c, 0x42, 0x8b, 0xc6, 0x90, 0x40, 0x28, 0x9c, 0x38, 0x41, 0x89, 0xe6,
	0xd9, 0x46, 0xeb, 0xec, 0x6a, 0x1d, 0xa5, 0xad, 0x9d, 0xf8, 0xe8, 0x63, 0xae, 0xa4, 0xeb, 0xbf,
	0xc2, 0x54, 0xbb, 0xf9, 0x7f, 0x00, 0x4c, 0xe5, 0x08, 0x93, 0x40, 0xa6, 0x5d, 0x2e, 0x61, 0x00,
	0x12, 0x94, 0xad, 0xca, 0x50, 0x84, 0x28, 0xdc, 0x82, 0x57, 0x14, 0xaa, 0x45, 0x8a, 0x8a, 0x7a,
	0x20, 0xf8, 0x96, 0xc3, 0x30, 0x88, 0x8a, 0xbe, 0xfd, 0x52, 0x32, 0xca, 0xbd, 0x3c, 0xb7, 0xf9,
	0xa6, 0x1b, 0xe6, 0x80, 0x47, 0x18, 0xba, 0x2c, 0xe1, 0x20, 0x03, 0x8c, 0xdc, 0x21, 0xe7, 0x15,
	0xbd, 0x7e, 0xd0, 0x38, 0xba, 0xaa, 0x3a, 0x0a, 0xbb, 0x06, 0x6d, 0xb6, 0x70, 0x3a, 0x18, 0x44,
	0xed, 0xee, 0x2c, 0xb3, 0xb4, 0x55, 0x66, 0x55, 0x53, 0x08, 0x27, 0x77, 0xf6, 0x6f, 0x84, 0xfd,
	0xfe, 0x69, 0x35, 0xfc, 0x40, 0x8e, 0xa6, 0x9e, 0xc3, 0x30, 0x54, 0x01, 0xd5, 0xa7, 0x29, 0x06,
	0x63, 0x2a, 0xd3, 0x98, 0x8b, 0x9c, 0x26, 0xfa, 0xc7, 0x39, 0xa0, 0xa3, 0xfc, 0xf7, 0x9c, 0x9b,
	0x43, 0xa3, 0xf6, 0x03, 0xea, 0x83, 0x70, 0x19, 0x46, 0x62, 0x1a, 0xf2, 0x4a, 0xa9, 0xae, 0x37,
	0x0e, 0xdb, 0x17, 0xb3, 0xcc, 0xd2, 0x57, 0x99, 0x75, 0xfe, 0x67, 0x88, 0x1d, 0xbd, 0xdd, 0x3f,
	0xdd, 0x1b, 0xf0, 0x00, 0xa2, 0x53, 0x74, 0xda, 0xbd, 0xd9, 0x82, 0xe8, 0xf3, 0x05, 0xd1, 0xbf,
	0x16, 0x44, 0x7f, 0x5d, 0x12, 0x6d, 0xbe, 0x24, 0xda, 0xc7, 0x92, 0x68, 0x8f, 0xb7, 0x3b, 0xe9,
	0x37, 0x27, 0x6a, 0x4e, 0xc0, 0x13, 0xdb, 0x8a, 0x3e, 0xef, 0x9f, 0x2c, 0xdf, 0xc8, 0x2b, 0xe7,
	0x8f, 0x7c, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x68, 0xd5, 0x04, 0x73, 0x39, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomCreationGasConsume != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DenomCreationGasConsume))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DenomCreationGasConsume != 0 {
		n += 1 + sovParams(uint64(m.DenomCreationGasConsume))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationGasConsume", wireType)
			}
			m.DenomCreationGasConsume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomCreationGasConsume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
