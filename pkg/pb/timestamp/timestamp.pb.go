// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timestamp.proto

package timestamp

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Timestamp is a HLC time value. All its field should never be accessed
// directly by its users.
type Timestamp struct {
	// PhysicalTime is the physical component of the HLC, it is read from a node's
	// wall clock time as Unix epoch time in nanoseconds. HLC requires this field
	// to be monotonically increase on each node.
	PhysicalTime int64 `protobuf:"varint,1,opt,name=PhysicalTime,proto3" json:"PhysicalTime,omitempty"`
	// LogicalTime is the logical component of the HLC, its value is maintained
	// according to the HLC algorithm. The HLC paper further establishes that its
	// value will not overflow in a real production environment.
	LogicalTime uint32 `protobuf:"varint,2,opt,name=LogicalTime,proto3" json:"LogicalTime,omitempty"`
	// NodeID just used to compatible with TAE some constraint to guaranteed unique
	// timestamp. uint16 is not defined, so use uint32, but only 2 bytes. The field
	// is not used for comparing the timestamps.
	NodeID               uint32   `protobuf:"varint,3,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_edac929d8ae1e24f, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return m.ProtoSize()
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetPhysicalTime() int64 {
	if m != nil {
		return m.PhysicalTime
	}
	return 0
}

func (m *Timestamp) GetLogicalTime() uint32 {
	if m != nil {
		return m.LogicalTime
	}
	return 0
}

func (m *Timestamp) GetNodeID() uint32 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "timestamp.Timestamp")
}

func init() { proto.RegisterFile("timestamp.proto", fileDescriptor_edac929d8ae1e24f) }

var fileDescriptor_edac929d8ae1e24f = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0x55, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa9, 0x94, 0xc9,
	0xc5, 0x19, 0x02, 0xd3, 0x2b, 0xa4, 0xc4, 0xc5, 0x13, 0x90, 0x51, 0x59, 0x9c, 0x99, 0x9c, 0x98,
	0x03, 0x12, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x42, 0x11, 0x13, 0x52, 0xe0, 0xe2, 0xf6,
	0xc9, 0x4f, 0x87, 0x2b, 0x61, 0x52, 0x60, 0xd4, 0xe0, 0x0d, 0x42, 0x16, 0x12, 0x12, 0xe3, 0x62,
	0xf3, 0xcb, 0x4f, 0x49, 0xf5, 0x74, 0x91, 0x60, 0x06, 0x4b, 0x42, 0x79, 0x4e, 0x2e, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xc3, 0x82, 0xc7, 0x72, 0x8c, 0x51,
	0x46, 0x48, 0x6e, 0xcd, 0x4d, 0x2c, 0x29, 0xca, 0xac, 0xc8, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x83,
	0x71, 0xf2, 0x52, 0xf5, 0x0b, 0xb2, 0xd3, 0xf5, 0x0b, 0x92, 0xf4, 0xe1, 0xfe, 0x4b, 0x62, 0x03,
	0xbb, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x38, 0x07, 0xe0, 0xd4, 0x04, 0x01, 0x00, 0x00,
}

func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Timestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NodeID != 0 {
		i = encodeVarintTimestamp(dAtA, i, uint64(m.NodeID))
		i--
		dAtA[i] = 0x18
	}
	if m.LogicalTime != 0 {
		i = encodeVarintTimestamp(dAtA, i, uint64(m.LogicalTime))
		i--
		dAtA[i] = 0x10
	}
	if m.PhysicalTime != 0 {
		i = encodeVarintTimestamp(dAtA, i, uint64(m.PhysicalTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimestamp(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimestamp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Timestamp) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PhysicalTime != 0 {
		n += 1 + sovTimestamp(uint64(m.PhysicalTime))
	}
	if m.LogicalTime != 0 {
		n += 1 + sovTimestamp(uint64(m.LogicalTime))
	}
	if m.NodeID != 0 {
		n += 1 + sovTimestamp(uint64(m.NodeID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTimestamp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimestamp(x uint64) (n int) {
	return sovTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Timestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimestamp
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
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PhysicalTime", wireType)
			}
			m.PhysicalTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PhysicalTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogicalTime", wireType)
			}
			m.LogicalTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogicalTime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimestamp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimestamp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimestamp
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
					return 0, ErrIntOverflowTimestamp
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
					return 0, ErrIntOverflowTimestamp
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
				return 0, ErrInvalidLengthTimestamp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimestamp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimestamp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimestamp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimestamp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimestamp = fmt.Errorf("proto: unexpected end of group")
)
