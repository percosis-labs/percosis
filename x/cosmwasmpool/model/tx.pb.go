// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: percosis/cosmwasmpool/v1beta1/model/tx.proto

package model

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ===================== MsgCreateCosmwasmPool
type MsgCreateCosmWasmPool struct {
	CodeId         uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	InstantiateMsg []byte `protobuf:"bytes,2,opt,name=instantiate_msg,json=instantiateMsg,proto3" json:"instantiate_msg,omitempty" yaml:"instantiate_msg"`
	Sender         string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *MsgCreateCosmWasmPool) Reset()         { *m = MsgCreateCosmWasmPool{} }
func (m *MsgCreateCosmWasmPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPool) ProtoMessage()    {}
func (*MsgCreateCosmWasmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{0}
}
func (m *MsgCreateCosmWasmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPool.Merge(m, src)
}
func (m *MsgCreateCosmWasmPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPool proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPool) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *MsgCreateCosmWasmPool) GetInstantiateMsg() []byte {
	if m != nil {
		return m.InstantiateMsg
	}
	return nil
}

func (m *MsgCreateCosmWasmPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// Returns a unique poolID to identify the pool with.
type MsgCreateCosmWasmPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateCosmWasmPoolResponse) Reset()         { *m = MsgCreateCosmWasmPoolResponse{} }
func (m *MsgCreateCosmWasmPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPoolResponse) ProtoMessage()    {}
func (*MsgCreateCosmWasmPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{1}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Merge(m, src)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPoolResponse proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateCosmWasmPool)(nil), "percosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPool")
	proto.RegisterType((*MsgCreateCosmWasmPoolResponse)(nil), "percosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPoolResponse")
}

func init() {
	proto.RegisterFile("percosis/cosmwasmpool/v1beta1/model/tx.proto", fileDescriptor_2ff1ac8555d314d1)
}

var fileDescriptor_2ff1ac8555d314d1 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x8a, 0xd3, 0x50,
	0x14, 0xc6, 0x7b, 0x55, 0x52, 0xbc, 0x68, 0xc5, 0x8b, 0x4a, 0x09, 0x9a, 0x94, 0xb8, 0xa9, 0x14,
	0x73, 0xa9, 0x05, 0x11, 0xdd, 0xa5, 0xdd, 0x74, 0x51, 0x90, 0x6c, 0x04, 0x37, 0xe5, 0x26, 0xb9,
	0xc4, 0x40, 0x6e, 0x4e, 0xc9, 0xb9, 0xd6, 0xfa, 0x02, 0xae, 0xdd, 0xf8, 0x26, 0x3e, 0xc4, 0x2c,
	0xbb, 0x9c, 0x55, 0x18, 0xd2, 0x37, 0xe8, 0x13, 0x0c, 0xf9, 0xd3, 0x99, 0xce, 0x50, 0x66, 0x31,
	0xbb, 0x93, 0x73, 0x7e, 0xdf, 0xc9, 0xc7, 0x77, 0x0f, 0x1d, 0x01, 0x2a, 0xc0, 0x04, 0x79, 0x08,
	0xa8, 0x7e, 0x09, 0x54, 0x2b, 0x80, 0x94, 0xaf, 0xc7, 0x81, 0xd4, 0x62, 0xcc, 0x15, 0x44, 0x32,
	0xe5, 0x7a, 0xe3, 0xae, 0x72, 0xd0, 0xc0, 0x5e, 0xb7, 0xb0, 0x7b, 0x0c, 0xbb, 0x2d, 0x6c, 0xbe,
	0x88, 0x21, 0x86, 0x1a, 0xe4, 0x55, 0xd5, 0x68, 0x4c, 0x2b, 0xac, 0x45, 0x3c, 0x10, 0x28, 0xaf,
	0xf6, 0x86, 0x90, 0x64, 0xcd, 0xdc, 0xf9, 0x4f, 0xe8, 0xcb, 0x05, 0xc6, 0xd3, 0x5c, 0x0a, 0x2d,
	0xa7, 0x80, 0xea, 0x9b, 0x40, 0xf5, 0x15, 0x20, 0x65, 0x23, 0xda, 0x0d, 0x21, 0x92, 0xcb, 0x24,
	0xea, 0x93, 0x01, 0x19, 0x3e, 0xf2, 0xd8, 0xbe, 0xb0, 0x7b, 0xbf, 0x85, 0x4a, 0x3f, 0x3b, 0xed,
	0xc0, 0xf1, 0x8d, 0xaa, 0x9a, 0x47, 0x6c, 0x4a, 0x9f, 0x25, 0x19, 0x6a, 0x91, 0xe9, 0x44, 0x68,
	0xb9, 0x54, 0x18, 0xf7, 0x1f, 0x0c, 0xc8, 0xf0, 0x89, 0x67, 0xee, 0x0b, 0xfb, 0x55, 0x23, 0xba,
	0x05, 0x38, 0x7e, 0xef, 0xa8, 0xb3, 0xc0, 0x98, 0xbd, 0xa3, 0x06, 0xca, 0x2c, 0x92, 0x79, 0xff,
	0xe1, 0x80, 0x0c, 0x1f, 0x7b, 0xcf, 0xf7, 0x85, 0xfd, 0xb4, 0xd1, 0x36, 0x7d, 0xc7, 0x6f, 0x01,
	0x67, 0x46, 0xdf, 0x9c, 0x74, 0xed, 0x4b, 0x5c, 0x41, 0x86, 0x92, 0xbd, 0xa5, 0xdd, 0x2a, 0x9d,
	0x6b, 0xf7, 0xb4, 0x2c, 0x6c, 0xa3, 0x42, 0xe6, 0x33, 0xdf, 0xa8, 0x46, 0xf3, 0xe8, 0xc3, 0x3f,
	0x42, 0xe9, 0x61, 0x0d, 0xe4, 0xec, 0x0f, 0xa1, 0xec, 0x44, 0x10, 0x13, 0xf7, 0xae, 0xdc, 0xdd,
	0x93, 0x3e, 0xcc, 0x2f, 0xf7, 0x10, 0x1d, 0xcc, 0x7b, 0xfe, 0x59, 0x69, 0x91, 0x6d, 0x69, 0x91,
	0x8b, 0xd2, 0x22, 0x7f, 0x77, 0x56, 0x67, 0xbb, 0xb3, 0x3a, 0xe7, 0x3b, 0xab, 0xf3, 0xfd, 0x53,
	0x9c, 0xe8, 0x1f, 0x3f, 0x03, 0x37, 0x04, 0xc5, 0xdb, 0x1f, 0xbc, 0x4f, 0x45, 0x80, 0x87, 0x0f,
	0xbe, 0x1e, 0x7f, 0xe4, 0x9b, 0x9b, 0xd7, 0x54, 0x5f, 0x51, 0x60, 0xd4, 0xef, 0x3d, 0xb9, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0x00, 0x4d, 0x07, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgCreatorClient is the client API for MsgCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgCreatorClient interface {
	CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error)
}

type msgCreatorClient struct {
	cc grpc1.ClientConn
}

func NewMsgCreatorClient(cc grpc1.ClientConn) MsgCreatorClient {
	return &msgCreatorClient{cc}
}

func (c *msgCreatorClient) CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error) {
	out := new(MsgCreateCosmWasmPoolResponse)
	err := c.cc.Invoke(ctx, "/percosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgCreatorServer is the server API for MsgCreator service.
type MsgCreatorServer interface {
	CreateCosmWasmPool(context.Context, *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error)
}

// UnimplementedMsgCreatorServer can be embedded to have forward compatible implementations.
type UnimplementedMsgCreatorServer struct {
}

func (*UnimplementedMsgCreatorServer) CreateCosmWasmPool(ctx context.Context, req *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCosmWasmPool not implemented")
}

func RegisterMsgCreatorServer(s grpc1.Server, srv MsgCreatorServer) {
	s.RegisterService(&_MsgCreator_serviceDesc, srv)
}

func _MsgCreator_CreateCosmWasmPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateCosmWasmPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, req.(*MsgCreateCosmWasmPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "percosis.cosmwasmpool.v1beta1.MsgCreator",
	HandlerType: (*MsgCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCosmWasmPool",
			Handler:    _MsgCreator_CreateCosmWasmPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "percosis/cosmwasmpool/v1beta1/model/tx.proto",
}

func (m *MsgCreateCosmWasmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InstantiateMsg) > 0 {
		i -= len(m.InstantiateMsg)
		copy(dAtA[i:], m.InstantiateMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InstantiateMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCosmWasmPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateCosmWasmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovTx(uint64(m.CodeId))
	}
	l = len(m.InstantiateMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateCosmWasmPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateCosmWasmPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiateMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstantiateMsg = append(m.InstantiateMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InstantiateMsg == nil {
				m.InstantiateMsg = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateCosmWasmPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
