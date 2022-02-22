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
	InflationRate              string                                  `protobuf:"bytes,10,opt,name=inflationRate,proto3" json:"inflationRate,omitempty"`
	RaresPerPack               uint64                                  `protobuf:"varint,11,opt,name=raresPerPack,proto3" json:"raresPerPack,omitempty"`
	CommonsPerPack             uint64                                  `protobuf:"varint,12,opt,name=commonsPerPack,proto3" json:"commonsPerPack,omitempty"`
	UnCommonsPerPack           uint64                                  `protobuf:"varint,13,opt,name=unCommonsPerPack,proto3" json:"unCommonsPerPack,omitempty"`
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

func (m *Params) GetInflationRate() string {
	if m != nil {
		return m.InflationRate
	}
	return ""
}

func (m *Params) GetRaresPerPack() uint64 {
	if m != nil {
		return m.RaresPerPack
	}
	return 0
}

func (m *Params) GetCommonsPerPack() uint64 {
	if m != nil {
		return m.CommonsPerPack
	}
	return 0
}

func (m *Params) GetUnCommonsPerPack() uint64 {
	if m != nil {
		return m.UnCommonsPerPack
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "DecentralCardGame.cardchain.cardchain.Params")
}

func init() { proto.RegisterFile("cardchain/params.proto", fileDescriptor_c4f04567b0d8a99a) }

var fileDescriptor_c4f04567b0d8a99a = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x13, 0x56, 0x0a, 0xf3, 0x3a, 0x5e, 0x2c, 0x5e, 0xac, 0x1d, 0xd2, 0x6a, 0xe2, 0xa5,
	0x42, 0xa2, 0x39, 0x70, 0x99, 0x38, 0x20, 0xb1, 0x8c, 0x71, 0x2d, 0x19, 0x17, 0x90, 0x38, 0x78,
	0xee, 0x43, 0x6a, 0x2d, 0xf1, 0x13, 0xd9, 0x4e, 0xb7, 0xf1, 0x29, 0x38, 0x72, 0xe4, 0xe3, 0xec,
	0xb8, 0x23, 0xe2, 0x30, 0xa1, 0xf6, 0xca, 0x87, 0x40, 0xf1, 0x20, 0x6d, 0x5a, 0xe0, 0x90, 0x53,
	0xac, 0xff, 0xf3, 0xf7, 0x2f, 0x7f, 0x5b, 0xcf, 0x63, 0x72, 0x4f, 0x70, 0x3d, 0x12, 0x63, 0x2e,
	0x55, 0x98, 0x73, 0xcd, 0x33, 0x33, 0xc8, 0x35, 0x5a, 0xa4, 0x0f, 0xf7, 0x40, 0x80, 0xb2, 0x9a,
	0xa7, 0x11, 0xd7, 0xa3, 0xd7, 0x3c, 0x83, 0x41, 0xe5, 0x9c, 0xaf, 0xb6, 0xee, 0x24, 0x98, 0xa0,
	0xdb, 0x11, 0x96, 0xab, 0xcb, 0xcd, 0xdb, 0x3f, 0xdb, 0xa4, 0x3d, 0x74, 0x34, 0xfa, 0x82, 0x6c,
	0x4d, 0xd0, 0x4a, 0x95, 0xc4, 0x32, 0x19, 0x5b, 0xf3, 0xea, 0x24, 0x97, 0x9a, 0x5b, 0x89, 0xea,
	0xad, 0xcc, 0x80, 0xf9, 0x3d, 0xbf, 0xbf, 0x16, 0xff, 0xc7, 0x41, 0x1f, 0x91, 0x1b, 0x02, 0xd3,
	0x14, 0x44, 0xa9, 0x1c, 0xc8, 0x4f, 0xc0, 0xae, 0xf4, 0xfc, 0x7e, 0x2b, 0x5e, 0x52, 0xe9, 0x3b,
	0x72, 0x73, 0xae, 0x0c, 0xb5, 0x14, 0xc0, 0xd6, 0x7a, 0x7e, 0x7f, 0x7d, 0x37, 0x3c, 0xbb, 0xe8,
	0x7a, 0xdf, 0x2f, 0xba, 0x8f, 0x13, 0x69, 0xc7, 0xc5, 0xe1, 0x40, 0x60, 0x16, 0x0a, 0x34, 0x19,
	0x9a, 0xdf, 0x9f, 0xa7, 0x66, 0x74, 0x14, 0xda, 0xd3, 0x1c, 0xcc, 0x20, 0x42, 0xa9, 0xe2, 0x65,
	0x0e, 0xdd, 0x21, 0xf7, 0xb9, 0xb0, 0x72, 0x02, 0x51, 0x55, 0x30, 0x2f, 0x33, 0x2c, 0x94, 0x65,
	0x2d, 0x97, 0xe5, 0x5f, 0x65, 0x0a, 0xe4, 0xee, 0x1c, 0x16, 0x69, 0x70, 0xc7, 0xda, 0x07, 0x60,
	0x57, 0x9b, 0x45, 0xfb, 0x3b, 0x8d, 0x7e, 0x20, 0xb7, 0xcb, 0x02, 0xb7, 0xa0, 0x79, 0xba, 0x07,
	0x39, 0x1a, 0x69, 0x59, 0xbb, 0xd9, 0x2f, 0x56, 0x49, 0xf4, 0x80, 0x74, 0x8e, 0xa5, 0x52, 0xa0,
	0x63, 0x38, 0xe6, 0x7a, 0xc4, 0xae, 0x35, 0x23, 0xd7, 0x20, 0xf4, 0x0d, 0xd9, 0x98, 0xa0, 0xad,
	0x98, 0xd7, 0x9b, 0x31, 0x17, 0x19, 0x65, 0xce, 0x31, 0x16, 0x3a, 0x3d, 0xdd, 0xe7, 0x85, 0x00,
	0xcb, 0xd6, 0x1b, 0xe6, 0x5c, 0x84, 0xd0, 0x07, 0x64, 0x53, 0xaa, 0x8f, 0xa9, 0xbb, 0xeb, 0x98,
	0x5b, 0x60, 0xa4, 0xa4, 0xc6, 0x75, 0x91, 0x6e, 0x93, 0x8e, 0xe6, 0x1a, 0xcc, 0x10, 0xf4, 0x90,
	0x8b, 0x23, 0xb6, 0xe1, 0xfa, 0xa2, 0xa6, 0x5d, 0x76, 0x72, 0x96, 0xa1, 0xaa, 0x5c, 0x9d, 0x3f,
	0x9d, 0xbc, 0xa8, 0xd2, 0x27, 0xe4, 0x56, 0xa1, 0xa2, 0xba, 0x73, 0xd3, 0x39, 0x57, 0xf4, 0xe7,
	0xad, 0x2f, 0x5f, 0xbb, 0xde, 0x6e, 0x7c, 0x36, 0x0d, 0xfc, 0xf3, 0x69, 0xe0, 0xff, 0x98, 0x06,
	0xfe, 0xe7, 0x59, 0xe0, 0x9d, 0xcf, 0x02, 0xef, 0xdb, 0x2c, 0xf0, 0xde, 0xef, 0x2c, 0x1c, 0x7a,
	0x65, 0xa0, 0xc3, 0xa8, 0x1a, 0xfd, 0x93, 0x70, 0xfe, 0x0c, 0xb8, 0xab, 0x38, 0x6c, 0xbb, 0x49,
	0x7e, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xb9, 0x6a, 0xf0, 0x20, 0x04, 0x00, 0x00,
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
	if m.UnCommonsPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnCommonsPerPack))
		i--
		dAtA[i] = 0x68
	}
	if m.CommonsPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommonsPerPack))
		i--
		dAtA[i] = 0x60
	}
	if m.RaresPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RaresPerPack))
		i--
		dAtA[i] = 0x58
	}
	if len(m.InflationRate) > 0 {
		i -= len(m.InflationRate)
		copy(dAtA[i:], m.InflationRate)
		i = encodeVarintParams(dAtA, i, uint64(len(m.InflationRate)))
		i--
		dAtA[i] = 0x52
	}
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
	l = len(m.InflationRate)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.RaresPerPack != 0 {
		n += 1 + sovParams(uint64(m.RaresPerPack))
	}
	if m.CommonsPerPack != 0 {
		n += 1 + sovParams(uint64(m.CommonsPerPack))
	}
	if m.UnCommonsPerPack != 0 {
		n += 1 + sovParams(uint64(m.UnCommonsPerPack))
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
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRate", wireType)
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
			m.InflationRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaresPerPack", wireType)
			}
			m.RaresPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaresPerPack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonsPerPack", wireType)
			}
			m.CommonsPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommonsPerPack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnCommonsPerPack", wireType)
			}
			m.UnCommonsPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnCommonsPerPack |= uint64(b&0x7F) << shift
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
