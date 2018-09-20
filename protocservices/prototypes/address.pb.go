// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prototypes/address.proto

/*
	Package prototypes is a generated protocol buffer package.

	It is generated from these files:
		prototypes/address.proto

	It has these top-level messages:
		AddressData
		AddressFilter
*/
package prototypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressData struct {
	OrgID               int32   `protobuf:"varint,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	AddressID           int64   `protobuf:"varint,2,opt,name=AddressID,proto3" json:"AddressID,omitempty"`
	GroupID             int32   `protobuf:"varint,3,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	HostAddress         string  `protobuf:"bytes,4,opt,name=HostAddress,proto3" json:"HostAddress,omitempty"`
	IPAddress           string  `protobuf:"bytes,5,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	DiscoveryTime       int64   `protobuf:"varint,6,opt,name=DiscoveryTime,proto3" json:"DiscoveryTime,omitempty"`
	DiscoveredBy        string  `protobuf:"bytes,7,opt,name=DiscoveredBy,proto3" json:"DiscoveredBy,omitempty"`
	LastScannedTime     int64   `protobuf:"varint,8,opt,name=LastScannedTime,proto3" json:"LastScannedTime,omitempty"`
	LastSeenTime        int64   `protobuf:"varint,9,opt,name=LastSeenTime,proto3" json:"LastSeenTime,omitempty"`
	ConfidenceScore     float32 `protobuf:"fixed32,10,opt,name=ConfidenceScore,proto3" json:"ConfidenceScore,omitempty"`
	UserConfidenceScore float32 `protobuf:"fixed32,11,opt,name=UserConfidenceScore,proto3" json:"UserConfidenceScore,omitempty"`
	IsSOA               bool    `protobuf:"varint,12,opt,name=IsSOA,proto3" json:"IsSOA,omitempty"`
	IsWildcardZone      bool    `protobuf:"varint,13,opt,name=IsWildcardZone,proto3" json:"IsWildcardZone,omitempty"`
	IsHostedService     bool    `protobuf:"varint,14,opt,name=IsHostedService,proto3" json:"IsHostedService,omitempty"`
	Ignored             bool    `protobuf:"varint,15,opt,name=Ignored,proto3" json:"Ignored,omitempty"`
	Deleted             bool    `protobuf:"varint,16,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	FoundFrom           int64   `protobuf:"varint,17,opt,name=FoundFrom,proto3" json:"FoundFrom,omitempty"`
	NSRecord            int32   `protobuf:"varint,18,opt,name=NSRecord,proto3" json:"NSRecord,omitempty"`
	AddressHash         string  `protobuf:"bytes,19,opt,name=AddressHash,proto3" json:"AddressHash,omitempty"`
}

func (m *AddressData) Reset()                    { *m = AddressData{} }
func (m *AddressData) String() string            { return proto.CompactTextString(m) }
func (*AddressData) ProtoMessage()               {}
func (*AddressData) Descriptor() ([]byte, []int) { return fileDescriptorAddress, []int{0} }

func (m *AddressData) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *AddressData) GetAddressID() int64 {
	if m != nil {
		return m.AddressID
	}
	return 0
}

func (m *AddressData) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *AddressData) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

func (m *AddressData) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *AddressData) GetDiscoveryTime() int64 {
	if m != nil {
		return m.DiscoveryTime
	}
	return 0
}

func (m *AddressData) GetDiscoveredBy() string {
	if m != nil {
		return m.DiscoveredBy
	}
	return ""
}

func (m *AddressData) GetLastScannedTime() int64 {
	if m != nil {
		return m.LastScannedTime
	}
	return 0
}

func (m *AddressData) GetLastSeenTime() int64 {
	if m != nil {
		return m.LastSeenTime
	}
	return 0
}

func (m *AddressData) GetConfidenceScore() float32 {
	if m != nil {
		return m.ConfidenceScore
	}
	return 0
}

func (m *AddressData) GetUserConfidenceScore() float32 {
	if m != nil {
		return m.UserConfidenceScore
	}
	return 0
}

func (m *AddressData) GetIsSOA() bool {
	if m != nil {
		return m.IsSOA
	}
	return false
}

func (m *AddressData) GetIsWildcardZone() bool {
	if m != nil {
		return m.IsWildcardZone
	}
	return false
}

func (m *AddressData) GetIsHostedService() bool {
	if m != nil {
		return m.IsHostedService
	}
	return false
}

func (m *AddressData) GetIgnored() bool {
	if m != nil {
		return m.Ignored
	}
	return false
}

func (m *AddressData) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *AddressData) GetFoundFrom() int64 {
	if m != nil {
		return m.FoundFrom
	}
	return 0
}

func (m *AddressData) GetNSRecord() int32 {
	if m != nil {
		return m.NSRecord
	}
	return 0
}

func (m *AddressData) GetAddressHash() string {
	if m != nil {
		return m.AddressHash
	}
	return ""
}

type AddressFilter struct {
	OrgID               int32 `protobuf:"varint,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	GroupID             int32 `protobuf:"varint,2,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	Start               int64 `protobuf:"varint,3,opt,name=Start,proto3" json:"Start,omitempty"`
	Limit               int32 `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
	WithIgnored         bool  `protobuf:"varint,5,opt,name=WithIgnored,proto3" json:"WithIgnored,omitempty"`
	IgnoredValue        bool  `protobuf:"varint,6,opt,name=IgnoredValue,proto3" json:"IgnoredValue,omitempty"`
	WithLastScannedTime bool  `protobuf:"varint,7,opt,name=WithLastScannedTime,proto3" json:"WithLastScannedTime,omitempty"`
	SinceScannedTime    int64 `protobuf:"varint,8,opt,name=SinceScannedTime,proto3" json:"SinceScannedTime,omitempty"`
}

func (m *AddressFilter) Reset()                    { *m = AddressFilter{} }
func (m *AddressFilter) String() string            { return proto.CompactTextString(m) }
func (*AddressFilter) ProtoMessage()               {}
func (*AddressFilter) Descriptor() ([]byte, []int) { return fileDescriptorAddress, []int{1} }

func (m *AddressFilter) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *AddressFilter) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *AddressFilter) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *AddressFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AddressFilter) GetWithIgnored() bool {
	if m != nil {
		return m.WithIgnored
	}
	return false
}

func (m *AddressFilter) GetIgnoredValue() bool {
	if m != nil {
		return m.IgnoredValue
	}
	return false
}

func (m *AddressFilter) GetWithLastScannedTime() bool {
	if m != nil {
		return m.WithLastScannedTime
	}
	return false
}

func (m *AddressFilter) GetSinceScannedTime() int64 {
	if m != nil {
		return m.SinceScannedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*AddressData)(nil), "AddressData")
	proto.RegisterType((*AddressFilter)(nil), "AddressFilter")
}
func (m *AddressData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OrgID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.OrgID))
	}
	if m.AddressID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.AddressID))
	}
	if m.GroupID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.GroupID))
	}
	if len(m.HostAddress) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAddress(dAtA, i, uint64(len(m.HostAddress)))
		i += copy(dAtA[i:], m.HostAddress)
	}
	if len(m.IPAddress) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAddress(dAtA, i, uint64(len(m.IPAddress)))
		i += copy(dAtA[i:], m.IPAddress)
	}
	if m.DiscoveryTime != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.DiscoveryTime))
	}
	if len(m.DiscoveredBy) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintAddress(dAtA, i, uint64(len(m.DiscoveredBy)))
		i += copy(dAtA[i:], m.DiscoveredBy)
	}
	if m.LastScannedTime != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.LastScannedTime))
	}
	if m.LastSeenTime != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.LastSeenTime))
	}
	if m.ConfidenceScore != 0 {
		dAtA[i] = 0x55
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.ConfidenceScore))))
		i += 4
	}
	if m.UserConfidenceScore != 0 {
		dAtA[i] = 0x5d
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.UserConfidenceScore))))
		i += 4
	}
	if m.IsSOA {
		dAtA[i] = 0x60
		i++
		if m.IsSOA {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IsWildcardZone {
		dAtA[i] = 0x68
		i++
		if m.IsWildcardZone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IsHostedService {
		dAtA[i] = 0x70
		i++
		if m.IsHostedService {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Ignored {
		dAtA[i] = 0x78
		i++
		if m.Ignored {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Deleted {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.FoundFrom != 0 {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.FoundFrom))
	}
	if m.NSRecord != 0 {
		dAtA[i] = 0x90
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.NSRecord))
	}
	if len(m.AddressHash) > 0 {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintAddress(dAtA, i, uint64(len(m.AddressHash)))
		i += copy(dAtA[i:], m.AddressHash)
	}
	return i, nil
}

func (m *AddressFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OrgID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.OrgID))
	}
	if m.GroupID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.GroupID))
	}
	if m.Start != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.Start))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.Limit))
	}
	if m.WithIgnored {
		dAtA[i] = 0x28
		i++
		if m.WithIgnored {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IgnoredValue {
		dAtA[i] = 0x30
		i++
		if m.IgnoredValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WithLastScannedTime {
		dAtA[i] = 0x38
		i++
		if m.WithLastScannedTime {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SinceScannedTime != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintAddress(dAtA, i, uint64(m.SinceScannedTime))
	}
	return i, nil
}

func encodeVarintAddress(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AddressData) Size() (n int) {
	var l int
	_ = l
	if m.OrgID != 0 {
		n += 1 + sovAddress(uint64(m.OrgID))
	}
	if m.AddressID != 0 {
		n += 1 + sovAddress(uint64(m.AddressID))
	}
	if m.GroupID != 0 {
		n += 1 + sovAddress(uint64(m.GroupID))
	}
	l = len(m.HostAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	l = len(m.IPAddress)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	if m.DiscoveryTime != 0 {
		n += 1 + sovAddress(uint64(m.DiscoveryTime))
	}
	l = len(m.DiscoveredBy)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	if m.LastScannedTime != 0 {
		n += 1 + sovAddress(uint64(m.LastScannedTime))
	}
	if m.LastSeenTime != 0 {
		n += 1 + sovAddress(uint64(m.LastSeenTime))
	}
	if m.ConfidenceScore != 0 {
		n += 5
	}
	if m.UserConfidenceScore != 0 {
		n += 5
	}
	if m.IsSOA {
		n += 2
	}
	if m.IsWildcardZone {
		n += 2
	}
	if m.IsHostedService {
		n += 2
	}
	if m.Ignored {
		n += 2
	}
	if m.Deleted {
		n += 3
	}
	if m.FoundFrom != 0 {
		n += 2 + sovAddress(uint64(m.FoundFrom))
	}
	if m.NSRecord != 0 {
		n += 2 + sovAddress(uint64(m.NSRecord))
	}
	l = len(m.AddressHash)
	if l > 0 {
		n += 2 + l + sovAddress(uint64(l))
	}
	return n
}

func (m *AddressFilter) Size() (n int) {
	var l int
	_ = l
	if m.OrgID != 0 {
		n += 1 + sovAddress(uint64(m.OrgID))
	}
	if m.GroupID != 0 {
		n += 1 + sovAddress(uint64(m.GroupID))
	}
	if m.Start != 0 {
		n += 1 + sovAddress(uint64(m.Start))
	}
	if m.Limit != 0 {
		n += 1 + sovAddress(uint64(m.Limit))
	}
	if m.WithIgnored {
		n += 2
	}
	if m.IgnoredValue {
		n += 2
	}
	if m.WithLastScannedTime {
		n += 2
	}
	if m.SinceScannedTime != 0 {
		n += 1 + sovAddress(uint64(m.SinceScannedTime))
	}
	return n
}

func sovAddress(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAddress(x uint64) (n int) {
	return sovAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddressData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddressData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressID", wireType)
			}
			m.AddressID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddressID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscoveryTime", wireType)
			}
			m.DiscoveryTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DiscoveryTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscoveredBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DiscoveredBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastScannedTime", wireType)
			}
			m.LastScannedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastScannedTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSeenTime", wireType)
			}
			m.LastSeenTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSeenTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfidenceScore", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.ConfidenceScore = float32(math.Float32frombits(v))
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserConfidenceScore", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.UserConfidenceScore = float32(math.Float32frombits(v))
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSOA", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsSOA = bool(v != 0)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsWildcardZone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsWildcardZone = bool(v != 0)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsHostedService", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsHostedService = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ignored", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ignored = bool(v != 0)
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FoundFrom", wireType)
			}
			m.FoundFrom = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FoundFrom |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NSRecord", wireType)
			}
			m.NSRecord = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NSRecord |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddress
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
func (m *AddressFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddressFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithIgnored", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithIgnored = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IgnoredValue = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithLastScannedTime", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithLastScannedTime = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SinceScannedTime", wireType)
			}
			m.SinceScannedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SinceScannedTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAddress
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
func skipAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAddress
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAddress
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAddress
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAddress(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAddress = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddress   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prototypes/address.proto", fileDescriptorAddress) }

var fileDescriptorAddress = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0x47, 0xd6, 0xd6, 0x5d, 0xb7, 0xe1, 0xed, 0xc2, 0x42, 0xa8, 0x8a, 0x2a, 0x84,
	0x22, 0x24, 0x56, 0x10, 0x4f, 0xb0, 0x2d, 0x2a, 0x8b, 0x34, 0x31, 0x94, 0x00, 0x93, 0x76, 0xe7,
	0xc5, 0x87, 0xd6, 0xa2, 0x8d, 0x2b, 0xdb, 0x9d, 0xd4, 0x27, 0xe0, 0x92, 0x5b, 0x1e, 0x89, 0x4b,
	0x1e, 0x01, 0x95, 0x17, 0x41, 0x3e, 0x4e, 0xd7, 0x36, 0x8c, 0x3b, 0xff, 0xdf, 0xf9, 0x7d, 0xe2,
	0xf8, 0xfc, 0x26, 0x6c, 0xa6, 0x95, 0x55, 0x76, 0x31, 0x03, 0x33, 0xe0, 0x42, 0x68, 0x30, 0xe6,
	0x04, 0x51, 0xff, 0x5b, 0x48, 0x3a, 0xa7, 0x9e, 0x24, 0xdc, 0x72, 0x7a, 0x4c, 0xc2, 0x2b, 0x3d,
	0x4a, 0x13, 0x16, 0x44, 0x41, 0x1c, 0x66, 0x5e, 0xd0, 0x67, 0xa4, 0x5d, 0x99, 0xd2, 0x84, 0x35,
	0xa2, 0x20, 0xde, 0xc9, 0xd6, 0x80, 0x32, 0xd2, 0x7c, 0xa7, 0xd5, 0x7c, 0x96, 0x26, 0x6c, 0x07,
	0x77, 0xad, 0x24, 0x8d, 0x48, 0xe7, 0x42, 0x19, 0x5b, 0x59, 0xd9, 0xe3, 0x28, 0x88, 0xdb, 0xd9,
	0x26, 0x72, 0x9d, 0xd3, 0x0f, 0xab, 0x7a, 0x88, 0xf5, 0x35, 0xa0, 0xcf, 0x49, 0x37, 0x91, 0xa6,
	0x50, 0x77, 0xa0, 0x17, 0x1f, 0xe5, 0x14, 0xd8, 0x2e, 0x7e, 0x7b, 0x1b, 0xd2, 0x3e, 0xd9, 0x5b,
	0x01, 0x10, 0x67, 0x0b, 0xd6, 0xc4, 0x36, 0x5b, 0x8c, 0xc6, 0xe4, 0xe0, 0x92, 0x1b, 0x9b, 0x17,
	0xbc, 0x2c, 0x41, 0x60, 0xaf, 0x16, 0xf6, 0xaa, 0x63, 0xd7, 0x0d, 0x11, 0x40, 0x89, 0xb6, 0x36,
	0xda, 0xb6, 0x98, 0xeb, 0x76, 0xae, 0xca, 0x2f, 0x52, 0x40, 0x59, 0x40, 0x5e, 0x28, 0x0d, 0x8c,
	0x44, 0x41, 0xdc, 0xc8, 0xea, 0x98, 0xbe, 0x26, 0x47, 0x9f, 0x0c, 0xe8, 0xba, 0xbb, 0x83, 0xee,
	0x87, 0x4a, 0x6e, 0x02, 0xa9, 0xc9, 0xaf, 0x4e, 0xd9, 0x5e, 0x14, 0xc4, 0xad, 0xcc, 0x0b, 0xfa,
	0x82, 0xec, 0xa7, 0xe6, 0x5a, 0x4e, 0x44, 0xc1, 0xb5, 0xb8, 0x51, 0x25, 0xb0, 0x2e, 0x96, 0x6b,
	0xd4, 0x9d, 0x2c, 0x35, 0xee, 0x82, 0x41, 0xe4, 0xa0, 0xef, 0x64, 0x01, 0x6c, 0x1f, 0x8d, 0x75,
	0xec, 0xa6, 0x96, 0x8e, 0x4a, 0xa5, 0x41, 0xb0, 0x03, 0x74, 0xac, 0xa4, 0xab, 0x24, 0x30, 0x01,
	0x0b, 0x82, 0x1d, 0xfa, 0x4a, 0x25, 0xdd, 0xb4, 0x86, 0x6a, 0x5e, 0x8a, 0xa1, 0x56, 0x53, 0xf6,
	0xc4, 0xe7, 0xe0, 0x1e, 0xd0, 0xa7, 0xa4, 0xf5, 0x3e, 0xcf, 0xa0, 0x50, 0x5a, 0x30, 0x8a, 0x41,
	0xb8, 0xd7, 0x2e, 0x09, 0xd5, 0x50, 0x2f, 0xb8, 0x19, 0xb3, 0x23, 0x9f, 0x84, 0x0d, 0xd4, 0xff,
	0xde, 0x20, 0xdd, 0x4a, 0x0f, 0xe5, 0xc4, 0x82, 0xfe, 0x4f, 0x16, 0x37, 0xd2, 0xd6, 0xd8, 0x4e,
	0xdb, 0x31, 0x09, 0x73, 0xcb, 0xb5, 0xc5, 0x14, 0xee, 0x64, 0x5e, 0x38, 0x7a, 0x29, 0xa7, 0xd2,
	0x62, 0xfa, 0xc2, 0xcc, 0x0b, 0x77, 0x9e, 0x6b, 0x69, 0xc7, 0xab, 0x1b, 0x08, 0xf1, 0x3f, 0x37,
	0x91, 0xcb, 0x41, 0xb5, 0xfc, 0xcc, 0x27, 0x73, 0x1f, 0xbd, 0x56, 0xb6, 0xc5, 0xdc, 0x74, 0xdd,
	0x96, 0x7a, 0xb2, 0x9a, 0x68, 0x7d, 0xa8, 0x44, 0x5f, 0x92, 0xc3, 0x5c, 0xe2, 0xac, 0xeb, 0x41,
	0xfc, 0x87, 0x9f, 0x9d, 0xff, 0x5c, 0xf6, 0x82, 0x5f, 0xcb, 0x5e, 0xf0, 0x7b, 0xd9, 0x0b, 0x7e,
	0xfc, 0xe9, 0x3d, 0xba, 0x79, 0x33, 0x92, 0x76, 0x3c, 0xbf, 0x3d, 0x29, 0xd4, 0x74, 0x30, 0x91,
	0xe5, 0x57, 0x2e, 0x5f, 0x49, 0x35, 0xe0, 0xd3, 0x01, 0x3e, 0xe6, 0xc2, 0xf8, 0xd1, 0x9a, 0xc1,
	0xfa, 0xb9, 0xdf, 0xee, 0xe2, 0xfa, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5f, 0x9e,
	0x7e, 0x03, 0x04, 0x00, 0x00,
}
