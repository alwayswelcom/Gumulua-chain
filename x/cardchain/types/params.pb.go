// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	VotingRightsExpirationTime int64                                   `protobuf:"varint,1,opt,name=votingRightsExpirationTime,proto3" json:"votingRightsExpirationTime,omitempty"`
	CollectionSize             uint64                                  `protobuf:"varint,2,opt,name=collectionSize,proto3" json:"collectionSize,omitempty"`
	CollectionPrice            github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=collectionPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collectionPrice"`
	ActiveCollectionsAmount    uint64                                  `protobuf:"varint,4,opt,name=activeCollectionsAmount,proto3" json:"activeCollectionsAmount,omitempty"`
	CollectionCreationFee      github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=collectionCreationFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collectionCreationFee"`
	CollateralDeposit          github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=collateralDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateralDeposit"`
	WinnerReward               github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,7,opt,name=winnerReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"winnerReward"`
	VoterReward                github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,8,opt,name=voterReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"voterReward"`
	HourlyFaucet               github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,9,opt,name=hourlyFaucet,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"hourlyFaucet"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4f04567b0d8a99a, []int{0}
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

func (m *Params) GetVotingRightsExpirationTime() int64 {
	if m != nil {
		return m.VotingRightsExpirationTime
	}
	return 0
}

func (m *Params) GetCollectionSize() uint64 {
	if m != nil {
		return m.CollectionSize
	}
	return 0
}

func (m *Params) GetActiveCollectionsAmount() uint64 {
	if m != nil {
		return m.ActiveCollectionsAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "DecentralCardGame.cardchain.cardchain.Params")
}

func init() { proto.RegisterFile("cardchain/params.proto", fileDescriptor_c4f04567b0d8a99a) }

var fileDescriptor_c4f04567b0d8a99a = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3d, 0xef, 0xd3, 0x30,
	0x10, 0xc6, 0x13, 0xfe, 0xfd, 0x17, 0x6a, 0x10, 0x08, 0x8b, 0x17, 0xab, 0x43, 0x5a, 0x21, 0x01,
	0x5d, 0x48, 0x06, 0x96, 0x8a, 0x01, 0x89, 0xa6, 0x94, 0xb5, 0xa4, 0x2c, 0x20, 0x31, 0xb8, 0xce,
	0x29, 0xb1, 0x48, 0xec, 0xc8, 0x76, 0xfa, 0xc2, 0xa7, 0x60, 0x64, 0xe4, 0xe3, 0x74, 0x42, 0x1d,
	0x11, 0x43, 0x85, 0xda, 0x2f, 0x82, 0x12, 0x50, 0x52, 0x5a, 0x60, 0xc8, 0xe4, 0xd3, 0x3d, 0x77,
	0xbf, 0x7b, 0x3c, 0x3c, 0xe8, 0x1e, 0xa3, 0x2a, 0x64, 0x31, 0xe5, 0xc2, 0xcb, 0xa8, 0xa2, 0xa9,
	0x76, 0x33, 0x25, 0x8d, 0xc4, 0x0f, 0xc7, 0xc0, 0x40, 0x18, 0x45, 0x13, 0x9f, 0xaa, 0xf0, 0x15,
	0x4d, 0xc1, 0xad, 0x26, 0xeb, 0xaa, 0x7b, 0x27, 0x92, 0x91, 0x2c, 0x37, 0xbc, 0xa2, 0xfa, 0xb5,
	0xfc, 0xe0, 0xeb, 0x25, 0x6a, 0x4f, 0x4b, 0x1a, 0x7e, 0x8e, 0xba, 0x0b, 0x69, 0xb8, 0x88, 0x02,
	0x1e, 0xc5, 0x46, 0xbf, 0x5c, 0x65, 0x5c, 0x51, 0xc3, 0xa5, 0x78, 0xc3, 0x53, 0x20, 0x76, 0xdf,
	0x1e, 0x5c, 0x04, 0xff, 0x99, 0xc0, 0x8f, 0xd0, 0x4d, 0x26, 0x93, 0x04, 0x58, 0xd1, 0x99, 0xf1,
	0x8f, 0x40, 0xae, 0xf4, 0xed, 0x41, 0x2b, 0x38, 0xe9, 0xe2, 0xb7, 0xe8, 0x56, 0xdd, 0x99, 0x2a,
	0xce, 0x80, 0x5c, 0xf4, 0xed, 0x41, 0x67, 0xe4, 0x6d, 0x76, 0x3d, 0xeb, 0xfb, 0xae, 0xf7, 0x38,
	0xe2, 0x26, 0xce, 0xe7, 0x2e, 0x93, 0xa9, 0xc7, 0xa4, 0x4e, 0xa5, 0xfe, 0xfd, 0x3c, 0xd1, 0xe1,
	0x07, 0xcf, 0xac, 0x33, 0xd0, 0xae, 0x2f, 0xb9, 0x08, 0x4e, 0x39, 0x78, 0x88, 0xee, 0x53, 0x66,
	0xf8, 0x02, 0xfc, 0x4a, 0xd0, 0x2f, 0x52, 0x99, 0x0b, 0x43, 0x5a, 0xa5, 0x97, 0x7f, 0xc9, 0x18,
	0xd0, 0xdd, 0x1a, 0xe6, 0x2b, 0x28, 0xbf, 0x35, 0x01, 0x20, 0x97, 0xcd, 0xac, 0xfd, 0x9d, 0x86,
	0xdf, 0xa3, 0xdb, 0x85, 0x40, 0x0d, 0x28, 0x9a, 0x8c, 0x21, 0x93, 0x9a, 0x1b, 0xd2, 0x6e, 0x76,
	0xe2, 0x9c, 0x84, 0x67, 0xe8, 0xc6, 0x92, 0x0b, 0x01, 0x2a, 0x80, 0x25, 0x55, 0x21, 0xb9, 0xda,
	0x8c, 0xfc, 0x07, 0x04, 0xbf, 0x46, 0xd7, 0x17, 0xd2, 0x54, 0xcc, 0x6b, 0xcd, 0x98, 0xc7, 0x8c,
	0xc2, 0x67, 0x2c, 0x73, 0x95, 0xac, 0x27, 0x34, 0x67, 0x60, 0x48, 0xa7, 0xa1, 0xcf, 0x63, 0xc8,
	0xb3, 0xd6, 0xe7, 0x2f, 0x3d, 0x6b, 0x14, 0x6c, 0xf6, 0x8e, 0xbd, 0xdd, 0x3b, 0xf6, 0x8f, 0xbd,
	0x63, 0x7f, 0x3a, 0x38, 0xd6, 0xf6, 0xe0, 0x58, 0xdf, 0x0e, 0x8e, 0xf5, 0x6e, 0x78, 0x84, 0x3d,
	0x8b, 0x8c, 0xe7, 0x57, 0xe1, 0x5a, 0x79, 0x75, 0xd0, 0xca, 0x63, 0xf3, 0x76, 0x99, 0x95, 0xa7,
	0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xc0, 0xa0, 0xd5, 0x82, 0x03, 0x00, 0x00,
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
	{
		size := m.HourlyFaucet.Size()
		i -= size
		if _, err := m.HourlyFaucet.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.VoterReward.Size()
		i -= size
		if _, err := m.VoterReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.WinnerReward.Size()
		i -= size
		if _, err := m.WinnerReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.CollateralDeposit.Size()
		i -= size
		if _, err := m.CollateralDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.CollectionCreationFee.Size()
		i -= size
		if _, err := m.CollectionCreationFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ActiveCollectionsAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ActiveCollectionsAmount))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.CollectionPrice.Size()
		i -= size
		if _, err := m.CollectionPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.CollectionSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CollectionSize))
		i--
		dAtA[i] = 0x10
	}
	if m.VotingRightsExpirationTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingRightsExpirationTime))
		i--
		dAtA[i] = 0x8
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
	if m.VotingRightsExpirationTime != 0 {
		n += 1 + sovParams(uint64(m.VotingRightsExpirationTime))
	}
	if m.CollectionSize != 0 {
		n += 1 + sovParams(uint64(m.CollectionSize))
	}
	l = m.CollectionPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.ActiveCollectionsAmount != 0 {
		n += 1 + sovParams(uint64(m.ActiveCollectionsAmount))
	}
	l = m.CollectionCreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CollateralDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WinnerReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VoterReward.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.HourlyFaucet.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingRightsExpirationTime", wireType)
			}
			m.VotingRightsExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingRightsExpirationTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionSize", wireType)
			}
			m.CollectionSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollectionSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollectionPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveCollectionsAmount", wireType)
			}
			m.ActiveCollectionsAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveCollectionsAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionCreationFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollectionCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WinnerReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoterReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourlyFaucet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HourlyFaucet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
