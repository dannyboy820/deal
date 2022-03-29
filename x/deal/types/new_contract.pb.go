// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deal/new_contract.proto

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

type NewContract struct {
	DealId     string `protobuf:"bytes,1,opt,name=dealId,proto3" json:"dealId,omitempty"`
	ContractId string `protobuf:"bytes,2,opt,name=contractId,proto3" json:"contractId,omitempty"`
	Consumer   string `protobuf:"bytes,3,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Desc       string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	OwnerETA   string `protobuf:"bytes,5,opt,name=ownerETA,proto3" json:"ownerETA,omitempty"`
	VendorETA  string `protobuf:"bytes,6,opt,name=vendorETA,proto3" json:"vendorETA,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Expiry     string `protobuf:"bytes,8,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (m *NewContract) Reset()         { *m = NewContract{} }
func (m *NewContract) String() string { return proto.CompactTextString(m) }
func (*NewContract) ProtoMessage()    {}
func (*NewContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_22b4a21c76f0584b, []int{0}
}
func (m *NewContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NewContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NewContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewContract.Merge(m, src)
}
func (m *NewContract) XXX_Size() int {
	return m.Size()
}
func (m *NewContract) XXX_DiscardUnknown() {
	xxx_messageInfo_NewContract.DiscardUnknown(m)
}

var xxx_messageInfo_NewContract proto.InternalMessageInfo

func (m *NewContract) GetDealId() string {
	if m != nil {
		return m.DealId
	}
	return ""
}

func (m *NewContract) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *NewContract) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *NewContract) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *NewContract) GetOwnerETA() string {
	if m != nil {
		return m.OwnerETA
	}
	return ""
}

func (m *NewContract) GetVendorETA() string {
	if m != nil {
		return m.VendorETA
	}
	return ""
}

func (m *NewContract) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *NewContract) GetExpiry() string {
	if m != nil {
		return m.Expiry
	}
	return ""
}

func init() {
	proto.RegisterType((*NewContract)(nil), "Harry027.deal.deal.NewContract")
}

func init() { proto.RegisterFile("deal/new_contract.proto", fileDescriptor_22b4a21c76f0584b) }

var fileDescriptor_22b4a21c76f0584b = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x49, 0x4d, 0xcc,
	0xd1, 0xcf, 0x4b, 0x2d, 0x8f, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xf2, 0x48, 0x2c, 0x2a, 0xaa, 0x34, 0x30, 0x32, 0xd7, 0x03, 0xa9,
	0x00, 0x13, 0x4a, 0xf7, 0x19, 0xb9, 0xb8, 0xfd, 0x52, 0xcb, 0x9d, 0xa1, 0x2a, 0x85, 0xc4, 0xb8,
	0xd8, 0x40, 0xe2, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x90, 0x1c,
	0x17, 0x17, 0xcc, 0x34, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0x1c, 0x92, 0x88, 0x90, 0x14, 0x17, 0x47,
	0x72, 0x7e, 0x5e, 0x71, 0x69, 0x6e, 0x6a, 0x91, 0x04, 0x33, 0x58, 0x16, 0xce, 0x17, 0x12, 0xe2,
	0x62, 0x49, 0x49, 0x2d, 0x4e, 0x96, 0x60, 0x01, 0x8b, 0x83, 0xd9, 0x20, 0xf5, 0xf9, 0xe5, 0x79,
	0xa9, 0x45, 0xae, 0x21, 0x8e, 0x12, 0xac, 0x10, 0xf5, 0x30, 0xbe, 0x90, 0x0c, 0x17, 0x67, 0x59,
	0x6a, 0x5e, 0x4a, 0x3e, 0x58, 0x92, 0x0d, 0x2c, 0x89, 0x10, 0x00, 0xb9, 0xb0, 0xb8, 0x24, 0xb1,
	0xa4, 0xb4, 0x58, 0x82, 0x1d, 0xe2, 0x42, 0x08, 0x0f, 0x24, 0x9e, 0x5a, 0x51, 0x90, 0x59, 0x54,
	0x29, 0xc1, 0x01, 0x11, 0x87, 0xf0, 0x9c, 0x1c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x1c,
	0x34, 0xba, 0x06, 0x46, 0xe6, 0xfa, 0xe0, 0xd0, 0xab, 0x80, 0x50, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0xe0, 0xe0, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x29, 0xe3, 0x95, 0x59,
	0x01, 0x00, 0x00,
}

func (m *NewContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NewContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Expiry) > 0 {
		i -= len(m.Expiry)
		copy(dAtA[i:], m.Expiry)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.Expiry)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.VendorETA) > 0 {
		i -= len(m.VendorETA)
		copy(dAtA[i:], m.VendorETA)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.VendorETA)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.OwnerETA) > 0 {
		i -= len(m.OwnerETA)
		copy(dAtA[i:], m.OwnerETA)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.OwnerETA)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractId) > 0 {
		i -= len(m.ContractId)
		copy(dAtA[i:], m.ContractId)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.ContractId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DealId) > 0 {
		i -= len(m.DealId)
		copy(dAtA[i:], m.DealId)
		i = encodeVarintNewContract(dAtA, i, uint64(len(m.DealId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNewContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovNewContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NewContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DealId)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.ContractId)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.OwnerETA)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.VendorETA)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	l = len(m.Expiry)
	if l > 0 {
		n += 1 + l + sovNewContract(uint64(l))
	}
	return n
}

func sovNewContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNewContract(x uint64) (n int) {
	return sovNewContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NewContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNewContract
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
			return fmt.Errorf("proto: NewContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DealId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DealId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerETA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerETA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorETA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorETA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewContract
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
				return ErrInvalidLengthNewContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expiry = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNewContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNewContract
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
func skipNewContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNewContract
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
					return 0, ErrIntOverflowNewContract
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
					return 0, ErrIntOverflowNewContract
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
				return 0, ErrInvalidLengthNewContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNewContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNewContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNewContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNewContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNewContract = fmt.Errorf("proto: unexpected end of group")
)
