// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incentive/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the incentive module's genesis state.
type GenesisState struct {
	Params                Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	EurxAccumulationTimes []GenesisAccumulationTime `protobuf:"bytes,2,rep,name=eurx_accumulation_times,json=eurxAccumulationTimes,proto3" json:"eurx_accumulation_times" yaml:"eurx_accumulation_times"`
	EurxMintingClaims     []EURXMintingClaim        `protobuf:"bytes,3,rep,name=eurx_minting_claims,json=eurxMintingClaims,proto3" json:"eurx_minting_claims" yaml:"eurx_minting_claims"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEurxAccumulationTimes() []GenesisAccumulationTime {
	if m != nil {
		return m.EurxAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetEurxMintingClaims() []EURXMintingClaim {
	if m != nil {
		return m.EurxMintingClaims
	}
	return nil
}

type GenesisAccumulationTime struct {
	CollateralType           string    `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "eurx.incentive.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "eurx.incentive.GenesisAccumulationTime")
}

func init() { proto.RegisterFile("incentive/genesis.proto", fileDescriptor_b5ea08f29a85d2f6) }

var fileDescriptor_b5ea08f29a85d2f6 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xbd, 0x6e, 0xd4, 0x30,
	0x1c, 0x3f, 0x5f, 0xa5, 0x4a, 0xa4, 0x50, 0x44, 0xa0, 0xbd, 0x90, 0x21, 0x39, 0x3c, 0xd0, 0x2e,
	0xd8, 0x52, 0xd9, 0xd8, 0x9a, 0x0a, 0x21, 0x21, 0x21, 0x55, 0xe1, 0x06, 0x60, 0x89, 0x7c, 0x91,
	0x09, 0x46, 0x76, 0x1c, 0xc5, 0x4e, 0xd5, 0x3c, 0x01, 0x6b, 0x1f, 0xab, 0x63, 0x17, 0x24, 0xa6,
	0x03, 0xdd, 0x2d, 0xcc, 0x7d, 0x82, 0xca, 0x76, 0xae, 0x77, 0xcd, 0xe9, 0xb6, 0xc4, 0xff, 0xdf,
	0xd7, 0xff, 0xc3, 0x1b, 0xb1, 0x32, 0xa7, 0xa5, 0x66, 0x17, 0x14, 0x17, 0xb4, 0xa4, 0x8a, 0x29,
	0x54, 0xd5, 0x52, 0x4b, 0x7f, 0xff, 0x67, 0xd5, 0x5e, 0xa2, 0xfb, 0x6a, 0xf8, 0xa2, 0x90, 0x85,
	0xb4, 0x25, 0x6c, 0xbe, 0x1c, 0x2a, 0x8c, 0x0b, 0x29, 0x0b, 0x4e, 0xb1, 0xfd, 0x9b, 0x36, 0xdf,
	0xb1, 0x66, 0x82, 0x2a, 0x4d, 0x44, 0xd5, 0x01, 0x5e, 0xae, 0xf4, 0xef, 0xbf, 0x5c, 0x09, 0xfe,
	0x1e, 0x7a, 0x8f, 0x3f, 0x38, 0xcf, 0xcf, 0x9a, 0x68, 0xea, 0xbf, 0xf7, 0x76, 0x2b, 0x52, 0x13,
	0xa1, 0x02, 0x30, 0x06, 0xc7, 0x7b, 0x27, 0x87, 0xe8, 0x61, 0x06, 0x74, 0x6e, 0xab, 0xc9, 0xc1,
	0xf5, 0x2c, 0x1e, 0xdc, 0xce, 0xe2, 0x27, 0x2d, 0x11, 0xfc, 0x1d, 0x74, 0x1c, 0x98, 0x76, 0x64,
	0xff, 0x17, 0xf0, 0x46, 0x86, 0x98, 0x91, 0x3c, 0x6f, 0x44, 0xc3, 0x89, 0x66, 0xb2, 0xcc, 0x6c,
	0xb0, 0x60, 0x38, 0xde, 0x39, 0xde, 0x3b, 0x39, 0xea, 0x0b, 0x77, 0x31, 0x4e, 0xd7, 0x08, 0x13,
	0x26, 0x68, 0xf2, 0xba, 0x73, 0x8a, 0x9c, 0xd3, 0x16, 0x55, 0x98, 0x1e, 0x98, 0x4a, 0x9f, 0xad,
	0x7c, 0xed, 0x3d, 0xb7, 0x14, 0xc1, 0x4a, 0xcd, 0xca, 0x22, 0xcb, 0x39, 0x61, 0x42, 0x05, 0x3b,
	0x36, 0xc4, 0xb8, 0x1f, 0xe2, 0xe3, 0xf9, 0xd7, 0x2f, 0x9f, 0x1c, 0xf2, 0xcc, 0x00, 0x13, 0xd8,
	0xb9, 0x87, 0x6b, 0xee, 0x0f, 0xa5, 0x60, 0xfa, 0xcc, 0xbc, 0xae, 0xb3, 0x14, 0xfc, 0x0f, 0xbc,
	0xd1, 0x96, 0x86, 0xfc, 0x33, 0xef, 0x69, 0x2e, 0x39, 0x27, 0x9a, 0xd6, 0x84, 0x67, 0xba, 0xad,
	0xa8, 0x9d, 0xf5, 0xa3, 0x24, 0xbc, 0x9d, 0xc5, 0x87, 0xce, 0xa7, 0x07, 0x80, 0xe9, 0xfe, 0xea,
	0x65, 0xd2, 0x56, 0xd4, 0x0c, 0x38, 0xac, 0x6a, 0x7a, 0xc1, 0x64, 0xa3, 0x36, 0xc7, 0x11, 0x0c,
	0xed, 0xf2, 0x42, 0xe4, 0x4e, 0x03, 0x2d, 0x4f, 0x03, 0x4d, 0x96, 0xa7, 0x91, 0xbc, 0xe9, 0x1a,
	0x7b, 0xd5, 0x2d, 0x70, 0xab, 0x16, 0xbc, 0xfa, 0x1b, 0x83, 0x34, 0x58, 0x02, 0x36, 0xf6, 0x73,
	0x7a, 0x3d, 0x8f, 0xc0, 0xcd, 0x3c, 0x02, 0xff, 0xe6, 0x11, 0xb8, 0x5a, 0x44, 0x83, 0x9b, 0x45,
	0x34, 0xf8, 0xb3, 0x88, 0x06, 0xdf, 0x8e, 0x0a, 0xa6, 0x7f, 0x34, 0x53, 0x94, 0x4b, 0x81, 0x79,
	0x5e, 0x52, 0x81, 0xcd, 0xa0, 0xf0, 0xe5, 0xea, 0x0a, 0xb1, 0x69, 0x4e, 0x4d, 0x77, 0x6d, 0xbe,
	0xb7, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x5d, 0xde, 0x73, 0x09, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EurxMintingClaims) > 0 {
		for iNdEx := len(m.EurxMintingClaims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EurxMintingClaims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.EurxAccumulationTimes) > 0 {
		for iNdEx := len(m.EurxAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EurxAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.EurxAccumulationTimes) > 0 {
		for _, e := range m.EurxAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EurxMintingClaims) > 0 {
		for _, e := range m.EurxMintingClaims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EurxAccumulationTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EurxAccumulationTimes = append(m.EurxAccumulationTimes, GenesisAccumulationTime{})
			if err := m.EurxAccumulationTimes[len(m.EurxAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EurxMintingClaims", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EurxMintingClaims = append(m.EurxMintingClaims, EURXMintingClaim{})
			if err := m.EurxMintingClaims[len(m.EurxMintingClaims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
