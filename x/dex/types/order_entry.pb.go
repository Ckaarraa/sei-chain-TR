// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/order_entry.proto

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

type OrderEntry struct {
	Price             github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
	Quantity          github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,2,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quantity" yaml:"quantity"`
	AllocationCreator []string                                 `protobuf:"bytes,3,rep,name=allocationCreator,proto3" json:"allocationCreator,omitempty"`
	Allocation        []github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,rep,name=allocation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"allocation" yaml:"allocation"`
	PriceDenom        Denom                                    `protobuf:"varint,5,opt,name=priceDenom,proto3,enum=seiprotocol.seichain.dex.Denom" json:"priceDenom,omitempty"`
	AssetDenom        Denom                                    `protobuf:"varint,6,opt,name=assetDenom,proto3,enum=seiprotocol.seichain.dex.Denom" json:"assetDenom,omitempty"`
}

func (m *OrderEntry) Reset()         { *m = OrderEntry{} }
func (m *OrderEntry) String() string { return proto.CompactTextString(m) }
func (*OrderEntry) ProtoMessage()    {}
func (*OrderEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_25878922effe12c2, []int{0}
}
func (m *OrderEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderEntry.Merge(m, src)
}
func (m *OrderEntry) XXX_Size() int {
	return m.Size()
}
func (m *OrderEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OrderEntry proto.InternalMessageInfo

func (m *OrderEntry) GetAllocationCreator() []string {
	if m != nil {
		return m.AllocationCreator
	}
	return nil
}

func (m *OrderEntry) GetPriceDenom() Denom {
	if m != nil {
		return m.PriceDenom
	}
	return Denom_SEI
}

func (m *OrderEntry) GetAssetDenom() Denom {
	if m != nil {
		return m.AssetDenom
	}
	return Denom_SEI
}

func init() {
	proto.RegisterType((*OrderEntry)(nil), "seiprotocol.seichain.dex.OrderEntry")
}

func init() { proto.RegisterFile("dex/order_entry.proto", fileDescriptor_25878922effe12c2) }

var fileDescriptor_25878922effe12c2 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd2, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0x00, 0xf0, 0xc4, 0xda, 0x62, 0x0f, 0xb1, 0x34, 0x28, 0x84, 0x0e, 0x49, 0xc9, 0x20, 0x1d,
	0xec, 0x05, 0x74, 0x73, 0x50, 0x6c, 0x2b, 0x8e, 0x42, 0x70, 0x12, 0x44, 0xae, 0x97, 0x47, 0x7b,
	0x98, 0xe4, 0xea, 0xdd, 0x15, 0x9a, 0x6f, 0xe1, 0xc7, 0xea, 0xd8, 0x49, 0xc4, 0xa1, 0x48, 0xfb,
	0x0d, 0xfc, 0x04, 0x72, 0x17, 0x6b, 0x0b, 0x22, 0xd8, 0x29, 0x2f, 0xef, 0xcf, 0x8f, 0xf7, 0x48,
	0xd0, 0x51, 0x0c, 0x93, 0x90, 0x8b, 0x18, 0xc4, 0x23, 0x64, 0x4a, 0xe4, 0x78, 0x24, 0xb8, 0xe2,
	0x8e, 0x2b, 0x81, 0x99, 0x88, 0xf2, 0x04, 0x4b, 0x60, 0x74, 0x48, 0x58, 0x86, 0x63, 0x98, 0x34,
	0x0e, 0x07, 0x7c, 0xc0, 0x4d, 0x29, 0xd4, 0x51, 0xd1, 0xdf, 0xa8, 0x69, 0x06, 0xb2, 0x71, 0x2a,
	0x8b, 0x44, 0xf0, 0x5a, 0x42, 0xe8, 0x56, 0xb3, 0xd7, 0x5a, 0x75, 0xee, 0x50, 0x79, 0x24, 0x18,
	0x05, 0xd7, 0x6e, 0xda, 0xad, 0x6a, 0xe7, 0x62, 0x3a, 0xf7, 0xad, 0xf7, 0xb9, 0x7f, 0x3c, 0x60,
	0x6a, 0x38, 0xee, 0x63, 0xca, 0xd3, 0x90, 0x72, 0x99, 0x72, 0xf9, 0xfd, 0x68, 0xcb, 0xf8, 0x29,
	0x54, 0xf9, 0x08, 0x24, 0xee, 0x01, 0xfd, 0x9c, 0xfb, 0xfb, 0x39, 0x49, 0x93, 0xf3, 0xc0, 0x20,
	0x41, 0x54, 0x60, 0xce, 0x03, 0xda, 0x7b, 0x1e, 0x93, 0x4c, 0x31, 0x95, 0xbb, 0x3b, 0x06, 0xbe,
	0xda, 0x1a, 0xae, 0x15, 0xf0, 0xca, 0x09, 0xa2, 0x1f, 0xd2, 0x39, 0x41, 0x75, 0x92, 0x24, 0x9c,
	0x12, 0xc5, 0x78, 0xd6, 0x15, 0x40, 0x14, 0x17, 0x6e, 0xa9, 0x59, 0x6a, 0x55, 0xa3, 0xdf, 0x05,
	0x87, 0x22, 0xb4, 0x4e, 0xba, 0xbb, 0xba, 0xad, 0xd3, 0xdd, 0x7a, 0x9d, 0x7a, 0xb1, 0xce, 0x5a,
	0x0a, 0xa2, 0x0d, 0xd6, 0xb9, 0x44, 0xc8, 0x9c, 0xde, 0x83, 0x8c, 0xa7, 0x6e, 0xb9, 0x69, 0xb7,
	0x0e, 0x4e, 0x7d, 0xfc, 0xd7, 0xc7, 0xc2, 0xa6, 0x2d, 0xda, 0x18, 0xd1, 0x00, 0x91, 0x12, 0x54,
	0x01, 0x54, 0xfe, 0x09, 0xac, 0x47, 0x3a, 0x37, 0xd3, 0x85, 0x67, 0xcf, 0x16, 0x9e, 0xfd, 0xb1,
	0xf0, 0xec, 0x97, 0xa5, 0x67, 0xcd, 0x96, 0x9e, 0xf5, 0xb6, 0xf4, 0xac, 0xfb, 0xf6, 0xc6, 0x91,
	0x12, 0x58, 0x7b, 0x25, 0x9a, 0x17, 0x43, 0x86, 0x93, 0x50, 0xff, 0x27, 0xe6, 0xde, 0x7e, 0xc5,
	0xd4, 0xcf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x06, 0x8b, 0x03, 0x66, 0x82, 0x02, 0x00, 0x00,
}

func (m *OrderEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetDenom != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.AssetDenom))
		i--
		dAtA[i] = 0x30
	}
	if m.PriceDenom != 0 {
		i = encodeVarintOrderEntry(dAtA, i, uint64(m.PriceDenom))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Allocation) > 0 {
		for iNdEx := len(m.Allocation) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Allocation[iNdEx].Size()
				i -= size
				if _, err := m.Allocation[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintOrderEntry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AllocationCreator) > 0 {
		for iNdEx := len(m.AllocationCreator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllocationCreator[iNdEx])
			copy(dAtA[i:], m.AllocationCreator[iNdEx])
			i = encodeVarintOrderEntry(dAtA, i, uint64(len(m.AllocationCreator[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrderEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintOrderEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrderEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OrderEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovOrderEntry(uint64(l))
	l = m.Quantity.Size()
	n += 1 + l + sovOrderEntry(uint64(l))
	if len(m.AllocationCreator) > 0 {
		for _, s := range m.AllocationCreator {
			l = len(s)
			n += 1 + l + sovOrderEntry(uint64(l))
		}
	}
	if len(m.Allocation) > 0 {
		for _, e := range m.Allocation {
			l = e.Size()
			n += 1 + l + sovOrderEntry(uint64(l))
		}
	}
	if m.PriceDenom != 0 {
		n += 1 + sovOrderEntry(uint64(m.PriceDenom))
	}
	if m.AssetDenom != 0 {
		n += 1 + sovOrderEntry(uint64(m.AssetDenom))
	}
	return n
}

func sovOrderEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrderEntry(x uint64) (n int) {
	return sovOrderEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OrderEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrderEntry
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
			return fmt.Errorf("proto: OrderEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocationCreator = append(m.AllocationCreator, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
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
				return ErrInvalidLengthOrderEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrderEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Allocation = append(m.Allocation, v)
			if err := m.Allocation[len(m.Allocation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			m.PriceDenom = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceDenom |= Denom(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			m.AssetDenom = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrderEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetDenom |= Denom(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrderEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrderEntry
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
func skipOrderEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
					return 0, ErrIntOverflowOrderEntry
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
				return 0, ErrInvalidLengthOrderEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrderEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrderEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrderEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrderEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrderEntry = fmt.Errorf("proto: unexpected end of group")
)
