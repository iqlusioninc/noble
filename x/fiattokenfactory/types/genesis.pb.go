// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiattokenfactory/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the fiattokenfactory module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BlacklistedList      []Blacklisted      `protobuf:"bytes,2,rep,name=blacklistedList,proto3" json:"blacklistedList"`
	Paused               *Paused            `protobuf:"bytes,3,opt,name=paused,proto3" json:"paused,omitempty"`
	MasterMinter         *MasterMinter      `protobuf:"bytes,4,opt,name=masterMinter,proto3" json:"masterMinter,omitempty"`
	MintersList          []Minters          `protobuf:"bytes,5,rep,name=mintersList,proto3" json:"mintersList"`
	Pauser               *Pauser            `protobuf:"bytes,6,opt,name=pauser,proto3" json:"pauser,omitempty"`
	Blacklister          *Blacklister       `protobuf:"bytes,7,opt,name=blacklister,proto3" json:"blacklister,omitempty"`
	Owner                *Owner             `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	MinterControllerList []MinterController `protobuf:"bytes,9,rep,name=minterControllerList,proto3" json:"minterControllerList"`
	MintingDenom         *MintingDenom      `protobuf:"bytes,10,opt,name=mintingDenom,proto3" json:"mintingDenom,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_192e1a68f2f38605, []int{0}
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

func (m *GenesisState) GetBlacklistedList() []Blacklisted {
	if m != nil {
		return m.BlacklistedList
	}
	return nil
}

func (m *GenesisState) GetPaused() *Paused {
	if m != nil {
		return m.Paused
	}
	return nil
}

func (m *GenesisState) GetMasterMinter() *MasterMinter {
	if m != nil {
		return m.MasterMinter
	}
	return nil
}

func (m *GenesisState) GetMintersList() []Minters {
	if m != nil {
		return m.MintersList
	}
	return nil
}

func (m *GenesisState) GetPauser() *Pauser {
	if m != nil {
		return m.Pauser
	}
	return nil
}

func (m *GenesisState) GetBlacklister() *Blacklister {
	if m != nil {
		return m.Blacklister
	}
	return nil
}

func (m *GenesisState) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *GenesisState) GetMinterControllerList() []MinterController {
	if m != nil {
		return m.MinterControllerList
	}
	return nil
}

func (m *GenesisState) GetMintingDenom() *MintingDenom {
	if m != nil {
		return m.MintingDenom
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "noble.fiattokenfactory.GenesisState")
}

func init() { proto.RegisterFile("fiattokenfactory/genesis.proto", fileDescriptor_192e1a68f2f38605) }

var fileDescriptor_192e1a68f2f38605 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x13, 0x76, 0x5b, 0xc0, 0x59, 0x09, 0xc9, 0xea, 0xc1, 0xca, 0x2e, 0xa1, 0xe4, 0xb0,
	0xda, 0x0b, 0x89, 0x76, 0xe1, 0x84, 0x38, 0x75, 0xf9, 0x73, 0xa1, 0x02, 0x82, 0xc4, 0x81, 0x4b,
	0x48, 0x52, 0x13, 0xa2, 0x26, 0x76, 0x65, 0xbb, 0x85, 0xbe, 0x05, 0xcf, 0xc3, 0x13, 0xf4, 0xd8,
	0x23, 0x27, 0x84, 0xda, 0x17, 0x41, 0xb1, 0x4d, 0x9a, 0x46, 0x69, 0x7a, 0x8b, 0x34, 0xbf, 0x6f,
	0x32, 0xdf, 0x7c, 0x63, 0x70, 0x21, 0xe8, 0x14, 0x93, 0xaf, 0x51, 0x22, 0x28, 0x5b, 0x86, 0xd7,
	0x7e, 0x8a, 0x09, 0xe6, 0x19, 0xf7, 0x66, 0x8c, 0x0a, 0x0a, 0x07, 0x84, 0xc6, 0x39, 0xf6, 0xf6,
	0x19, 0x7b, 0x90, 0xd2, 0x94, 0x4a, 0xc0, 0x2f, 0xbf, 0x14, 0x6b, 0x9f, 0x37, 0x3a, 0xcd, 0x22,
	0x16, 0x15, 0xba, 0x91, 0x3d, 0x6c, 0x14, 0xe3, 0x3c, 0x4a, 0xa6, 0x79, 0xc6, 0x05, 0x9e, 0x1c,
	0x94, 0xcf, 0x79, 0x55, 0x74, 0x1b, 0xc5, 0x22, 0xe2, 0x02, 0xb3, 0xb0, 0xc8, 0x88, 0xc0, 0x4c,
	0x33, 0x4d, 0x27, 0xaa, 0xc8, 0xbb, 0xda, 0xb3, 0xa3, 0xd3, 0xfd, 0x27, 0xec, 0x06, 0x41, 0xbf,
	0x93, 0xaa, 0x76, 0xd9, 0xfa, 0xe3, 0x30, 0xa1, 0x44, 0x30, 0x9a, 0xe7, 0x15, 0xe7, 0xb6, 0x70,
	0x19, 0x49, 0xc3, 0x09, 0x26, 0xb4, 0x50, 0x8c, 0xfb, 0xab, 0x07, 0xce, 0xde, 0xa8, 0x08, 0x3e,
	0x8a, 0x48, 0x60, 0xf8, 0x1c, 0xf4, 0xd5, 0x22, 0x91, 0x39, 0x34, 0xaf, 0xac, 0x9b, 0x0b, 0xaf,
	0x2d, 0x12, 0xef, 0xbd, 0x64, 0x46, 0xa7, 0xab, 0x3f, 0x8f, 0x8c, 0x40, 0x2b, 0xe0, 0x07, 0xf0,
	0xa0, 0xb6, 0xe7, 0xb7, 0x19, 0x17, 0xe8, 0xce, 0xf0, 0xe4, 0xca, 0xba, 0x79, 0xdc, 0xde, 0x64,
	0xb4, 0x83, 0x75, 0xa7, 0xa6, 0x1e, 0x3e, 0x2b, 0xc7, 0x29, 0x83, 0x41, 0x27, 0xdd, 0xe3, 0x94,
	0x4c, 0xa0, 0x59, 0xf8, 0x1a, 0x9c, 0xa9, 0xc4, 0xc6, 0x72, 0x35, 0xe8, 0x54, 0x6a, 0xdd, 0x76,
	0xed, 0xb8, 0x46, 0x06, 0x7b, 0x3a, 0xf8, 0x0a, 0x58, 0x3a, 0x55, 0x69, 0xa6, 0x27, 0xcd, 0x3c,
	0x3c, 0xd0, 0x46, 0x81, 0xda, 0x48, 0x5d, 0x57, 0x99, 0x60, 0xa8, 0x7f, 0xd4, 0x04, 0xd3, 0x26,
	0x18, 0xbc, 0x05, 0x56, 0xed, 0x2e, 0xd0, 0x5d, 0x29, 0x3d, 0xba, 0x49, 0x16, 0xd4, 0x55, 0xf0,
	0x1a, 0xf4, 0xe4, 0xe9, 0xa0, 0x7b, 0x52, 0x7e, 0xde, 0x2e, 0x7f, 0x57, 0x22, 0x81, 0x22, 0xe1,
	0x17, 0x30, 0x50, 0xc3, 0xdf, 0x56, 0x07, 0x25, 0xdd, 0xdf, 0x97, 0xee, 0x2f, 0xbb, 0xdc, 0xef,
	0x14, 0x7a, 0x0d, 0xad, 0x9d, 0x64, 0x3c, 0xea, 0x16, 0x5f, 0x96, 0xa7, 0x88, 0x40, 0x67, 0x3c,
	0x35, 0x32, 0xd8, 0xd3, 0x8d, 0x3e, 0xad, 0x36, 0x8e, 0xb9, 0xde, 0x38, 0xe6, 0xdf, 0x8d, 0x63,
	0xfe, 0xdc, 0x3a, 0xc6, 0x7a, 0xeb, 0x18, 0xbf, 0xb7, 0x8e, 0xf1, 0xf9, 0x45, 0x9a, 0x89, 0x6f,
	0xf3, 0xd8, 0x4b, 0x68, 0xe1, 0x73, 0xc1, 0x22, 0x92, 0xe2, 0x9c, 0x2e, 0xf0, 0x93, 0x05, 0x26,
	0x62, 0xce, 0x30, 0xf7, 0xe5, 0xaf, 0xfc, 0x1f, 0x7e, 0xe3, 0x89, 0x88, 0xe5, 0x0c, 0xf3, 0xb8,
	0x2f, 0xdf, 0xc6, 0xd3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x6c, 0xeb, 0x64, 0xac, 0x04,
	0x00, 0x00,
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
	if m.MintingDenom != nil {
		{
			size, err := m.MintingDenom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.MinterControllerList) > 0 {
		for iNdEx := len(m.MinterControllerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinterControllerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.Owner != nil {
		{
			size, err := m.Owner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Blacklister != nil {
		{
			size, err := m.Blacklister.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Pauser != nil {
		{
			size, err := m.Pauser.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.MintersList) > 0 {
		for iNdEx := len(m.MintersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.MasterMinter != nil {
		{
			size, err := m.MasterMinter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Paused != nil {
		{
			size, err := m.Paused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BlacklistedList) > 0 {
		for iNdEx := len(m.BlacklistedList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlacklistedList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.BlacklistedList) > 0 {
		for _, e := range m.BlacklistedList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Paused != nil {
		l = m.Paused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MasterMinter != nil {
		l = m.MasterMinter.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MintersList) > 0 {
		for _, e := range m.MintersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Pauser != nil {
		l = m.Pauser.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Blacklister != nil {
		l = m.Blacklister.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Owner != nil {
		l = m.Owner.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MinterControllerList) > 0 {
		for _, e := range m.MinterControllerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MintingDenom != nil {
		l = m.MintingDenom.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedList", wireType)
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
			m.BlacklistedList = append(m.BlacklistedList, Blacklisted{})
			if err := m.BlacklistedList[len(m.BlacklistedList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
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
			if m.Paused == nil {
				m.Paused = &Paused{}
			}
			if err := m.Paused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterMinter", wireType)
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
			if m.MasterMinter == nil {
				m.MasterMinter = &MasterMinter{}
			}
			if err := m.MasterMinter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintersList", wireType)
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
			m.MintersList = append(m.MintersList, Minters{})
			if err := m.MintersList[len(m.MintersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pauser", wireType)
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
			if m.Pauser == nil {
				m.Pauser = &Pauser{}
			}
			if err := m.Pauser.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklister", wireType)
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
			if m.Blacklister == nil {
				m.Blacklister = &Blacklister{}
			}
			if err := m.Blacklister.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			if m.Owner == nil {
				m.Owner = &Owner{}
			}
			if err := m.Owner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinterControllerList", wireType)
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
			m.MinterControllerList = append(m.MinterControllerList, MinterController{})
			if err := m.MinterControllerList[len(m.MinterControllerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingDenom", wireType)
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
			if m.MintingDenom == nil {
				m.MintingDenom = &MintingDenom{}
			}
			if err := m.MintingDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
