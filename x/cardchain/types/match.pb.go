// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/match.proto

package types

import (
	fmt "fmt"
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

type Match struct {
	Timestamp uint64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Reporter  string  `protobuf:"bytes,2,opt,name=reporter,proto3" json:"reporter,omitempty"`
	PlayerA   string  `protobuf:"bytes,3,opt,name=playerA,proto3" json:"playerA,omitempty"`
	PlayerB   string  `protobuf:"bytes,4,opt,name=playerB,proto3" json:"playerB,omitempty"`
	Outcome   Outcome `protobuf:"varint,5,opt,name=outcome,proto3,enum=DecentralCardGame.cardchain.cardchain.Outcome" json:"outcome,omitempty"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcab7767e7f2388f, []int{0}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Match.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return m.Size()
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Match) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *Match) GetPlayerA() string {
	if m != nil {
		return m.PlayerA
	}
	return ""
}

func (m *Match) GetPlayerB() string {
	if m != nil {
		return m.PlayerB
	}
	return ""
}

func (m *Match) GetOutcome() Outcome {
	if m != nil {
		return m.Outcome
	}
	return Outcome_AWon
}

func init() {
	proto.RegisterType((*Match)(nil), "DecentralCardGame.cardchain.cardchain.Match")
}

func init() { proto.RegisterFile("cardchain/match.proto", fileDescriptor_bcab7767e7f2388f) }

var fileDescriptor_bcab7767e7f2388f = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4e, 0x2c, 0x4a,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x75, 0x49, 0x4d, 0x4e, 0xcd, 0x2b, 0x29, 0x4a, 0xcc, 0x71, 0x4e, 0x2c, 0x4a,
	0x71, 0x4f, 0xcc, 0x4d, 0xd5, 0x83, 0x2b, 0x44, 0xb0, 0xa4, 0x84, 0x10, 0xba, 0x4b, 0x2a, 0x20,
	0x5a, 0x95, 0xf6, 0x33, 0x72, 0xb1, 0xfa, 0x82, 0x8c, 0x12, 0x92, 0xe1, 0xe2, 0x2c, 0xc9, 0xcc,
	0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x42, 0x08, 0x08,
	0x49, 0x71, 0x71, 0x14, 0xa5, 0x16, 0xe4, 0x17, 0x95, 0xa4, 0x16, 0x49, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x06, 0xc1, 0xf9, 0x42, 0x12, 0x5c, 0xec, 0x05, 0x39, 0x89, 0x95, 0xa9, 0x45, 0x8e, 0x12,
	0xcc, 0x60, 0x29, 0x18, 0x17, 0x21, 0xe3, 0x24, 0xc1, 0x82, 0x2c, 0xe3, 0x24, 0xe4, 0xc1, 0xc5,
	0x9e, 0x5f, 0x5a, 0x92, 0x9c, 0x9f, 0x9b, 0x2a, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x67, 0xa4, 0xa7,
	0x47, 0x94, 0x27, 0xf4, 0xfc, 0x21, 0xba, 0x82, 0x60, 0xda, 0x9d, 0x82, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x22, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0xc3, 0x70, 0x7d, 0x67, 0x78, 0x60, 0x54, 0xe8, 0x23, 0x05, 0x4c, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0x38, 0x70, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x86, 0xa5,
	0xc6, 0x70, 0x01, 0x00, 0x00,
}

func (m *Match) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Match) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Match) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Outcome != 0 {
		i = encodeVarintMatch(dAtA, i, uint64(m.Outcome))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PlayerB) > 0 {
		i -= len(m.PlayerB)
		copy(dAtA[i:], m.PlayerB)
		i = encodeVarintMatch(dAtA, i, uint64(len(m.PlayerB)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PlayerA) > 0 {
		i -= len(m.PlayerA)
		copy(dAtA[i:], m.PlayerA)
		i = encodeVarintMatch(dAtA, i, uint64(len(m.PlayerA)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintMatch(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintMatch(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMatch(dAtA []byte, offset int, v uint64) int {
	offset -= sovMatch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Match) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovMatch(uint64(m.Timestamp))
	}
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	l = len(m.PlayerA)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	l = len(m.PlayerB)
	if l > 0 {
		n += 1 + l + sovMatch(uint64(l))
	}
	if m.Outcome != 0 {
		n += 1 + sovMatch(uint64(m.Outcome))
	}
	return n
}

func sovMatch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMatch(x uint64) (n int) {
	return sovMatch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Match) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatch
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
			return fmt.Errorf("proto: Match: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Match: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
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
				return ErrInvalidLengthMatch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outcome", wireType)
			}
			m.Outcome = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Outcome |= Outcome(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMatch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMatch
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
func skipMatch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
					return 0, ErrIntOverflowMatch
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
				return 0, ErrInvalidLengthMatch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMatch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMatch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMatch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMatch = fmt.Errorf("proto: unexpected end of group")
)
