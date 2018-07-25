// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prototypes/user.proto

/*
	Package prototypes is a generated protocol buffer package.

	It is generated from these files:
		prototypes/user.proto

	It has these top-level messages:
		UserContext
		User
		UserFilter
*/
package prototypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type UserContext struct {
	TraceID   string   `protobuf:"bytes,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	OrgID     int32    `protobuf:"varint,2,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	UserID    int32    `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IPAddress string   `protobuf:"bytes,4,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	Roles     []string `protobuf:"bytes,5,rep,name=Roles" json:"Roles,omitempty"`
}

func (m *UserContext) Reset()                    { *m = UserContext{} }
func (m *UserContext) String() string            { return proto.CompactTextString(m) }
func (*UserContext) ProtoMessage()               {}
func (*UserContext) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{0} }

func (m *UserContext) GetTraceID() string {
	if m != nil {
		return m.TraceID
	}
	return ""
}

func (m *UserContext) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *UserContext) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserContext) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *UserContext) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type User struct {
	OrgID        int32  `protobuf:"varint,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	OrgCID       string `protobuf:"bytes,2,opt,name=OrgCID,proto3" json:"OrgCID,omitempty"`
	UserCID      string `protobuf:"bytes,3,opt,name=UserCID,proto3" json:"UserCID,omitempty"`
	UserID       int32  `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserEmail    string `protobuf:"bytes,5,opt,name=UserEmail,proto3" json:"UserEmail,omitempty"`
	FirstName    string `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName     string `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
	StatusID     int32  `protobuf:"varint,8,opt,name=StatusID,proto3" json:"StatusID,omitempty"`
	CreationTime int64  `protobuf:"varint,9,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	Deleted      bool   `protobuf:"varint,10,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{1} }

func (m *User) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *User) GetOrgCID() string {
	if m != nil {
		return m.OrgCID
	}
	return ""
}

func (m *User) GetUserCID() string {
	if m != nil {
		return m.UserCID
	}
	return ""
}

func (m *User) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *User) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetStatusID() int32 {
	if m != nil {
		return m.StatusID
	}
	return 0
}

func (m *User) GetCreationTime() int64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *User) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type UserFilter struct {
	Start             int32 `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
	Limit             int32 `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	OrgID             int32 `protobuf:"varint,3,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	WithStatus        bool  `protobuf:"varint,4,opt,name=WithStatus,proto3" json:"WithStatus,omitempty"`
	StatusValue       int32 `protobuf:"varint,5,opt,name=StatusValue,proto3" json:"StatusValue,omitempty"`
	WithDeleted       bool  `protobuf:"varint,6,opt,name=WithDeleted,proto3" json:"WithDeleted,omitempty"`
	DeletedValue      bool  `protobuf:"varint,7,opt,name=DeletedValue,proto3" json:"DeletedValue,omitempty"`
	SinceCreationTime int64 `protobuf:"varint,8,opt,name=SinceCreationTime,proto3" json:"SinceCreationTime,omitempty"`
}

func (m *UserFilter) Reset()                    { *m = UserFilter{} }
func (m *UserFilter) String() string            { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()               {}
func (*UserFilter) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{2} }

func (m *UserFilter) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *UserFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserFilter) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *UserFilter) GetWithStatus() bool {
	if m != nil {
		return m.WithStatus
	}
	return false
}

func (m *UserFilter) GetStatusValue() int32 {
	if m != nil {
		return m.StatusValue
	}
	return 0
}

func (m *UserFilter) GetWithDeleted() bool {
	if m != nil {
		return m.WithDeleted
	}
	return false
}

func (m *UserFilter) GetDeletedValue() bool {
	if m != nil {
		return m.DeletedValue
	}
	return false
}

func (m *UserFilter) GetSinceCreationTime() int64 {
	if m != nil {
		return m.SinceCreationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserContext)(nil), "UserContext")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*UserFilter)(nil), "UserFilter")
}
func (m *UserContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TraceID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.TraceID)))
		i += copy(dAtA[i:], m.TraceID)
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if m.UserID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserID))
	}
	if len(m.IPAddress) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.IPAddress)))
		i += copy(dAtA[i:], m.IPAddress)
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OrgID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if len(m.OrgCID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.OrgCID)))
		i += copy(dAtA[i:], m.OrgCID)
	}
	if len(m.UserCID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserCID)))
		i += copy(dAtA[i:], m.UserCID)
	}
	if m.UserID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserID))
	}
	if len(m.UserEmail) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserEmail)))
		i += copy(dAtA[i:], m.UserEmail)
	}
	if len(m.FirstName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i += copy(dAtA[i:], m.FirstName)
	}
	if len(m.LastName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i += copy(dAtA[i:], m.LastName)
	}
	if m.StatusID != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.StatusID))
	}
	if m.CreationTime != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.CreationTime))
	}
	if m.Deleted {
		dAtA[i] = 0x50
		i++
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *UserFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Start))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Limit))
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if m.WithStatus {
		dAtA[i] = 0x20
		i++
		if m.WithStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.StatusValue != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.StatusValue))
	}
	if m.WithDeleted {
		dAtA[i] = 0x30
		i++
		if m.WithDeleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DeletedValue {
		dAtA[i] = 0x38
		i++
		if m.DeletedValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SinceCreationTime != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.SinceCreationTime))
	}
	return i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserContext) Size() (n int) {
	var l int
	_ = l
	l = len(m.TraceID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	if m.UserID != 0 {
		n += 1 + sovUser(uint64(m.UserID))
	}
	l = len(m.IPAddress)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	return n
}

func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	l = len(m.OrgCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UserCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.UserID != 0 {
		n += 1 + sovUser(uint64(m.UserID))
	}
	l = len(m.UserEmail)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.StatusID != 0 {
		n += 1 + sovUser(uint64(m.StatusID))
	}
	if m.CreationTime != 0 {
		n += 1 + sovUser(uint64(m.CreationTime))
	}
	if m.Deleted {
		n += 2
	}
	return n
}

func (m *UserFilter) Size() (n int) {
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovUser(uint64(m.Start))
	}
	if m.Limit != 0 {
		n += 1 + sovUser(uint64(m.Limit))
	}
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	if m.WithStatus {
		n += 2
	}
	if m.StatusValue != 0 {
		n += 1 + sovUser(uint64(m.StatusValue))
	}
	if m.WithDeleted {
		n += 2
	}
	if m.DeletedValue {
		n += 2
	}
	if m.SinceCreationTime != 0 {
		n += 1 + sovUser(uint64(m.SinceCreationTime))
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusID", wireType)
			}
			m.StatusID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			m.CreationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *UserFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
			m.WithStatus = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusValue", wireType)
			}
			m.StatusValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusValue |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithDeleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
			m.WithDeleted = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
			m.DeletedValue = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SinceCreationTime", wireType)
			}
			m.SinceCreationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SinceCreationTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prototypes/user.proto", fileDescriptorUser) }

var fileDescriptorUser = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x71, 0xdb, 0xb4, 0xc9, 0x2c, 0x17, 0x2c, 0x40, 0x16, 0x42, 0x51, 0xd4, 0x53, 0x0e,
	0xa8, 0x11, 0x42, 0x3c, 0x00, 0x34, 0x2c, 0x8a, 0xb4, 0x62, 0x91, 0xbb, 0x80, 0xc4, 0xcd, 0xb4,
	0xa3, 0x62, 0x6d, 0x52, 0x57, 0xb6, 0xbb, 0x82, 0x3b, 0x17, 0xde, 0x60, 0x1f, 0x89, 0x23, 0x8f,
	0x80, 0xca, 0x8b, 0x20, 0xdb, 0x49, 0x9b, 0x2e, 0xb7, 0xf9, 0xfe, 0xb1, 0xc7, 0xff, 0xfc, 0x6d,
	0xe0, 0xd1, 0x56, 0x2b, 0xab, 0xec, 0xf7, 0x2d, 0x9a, 0x62, 0x67, 0x50, 0xcf, 0x3c, 0x4f, 0x7f,
	0x12, 0x38, 0xfb, 0x60, 0x50, 0xcf, 0xd5, 0xc6, 0xe2, 0x37, 0x4b, 0x19, 0x4c, 0xae, 0xb4, 0x58,
	0x62, 0x55, 0x32, 0x92, 0x91, 0x3c, 0xe1, 0x1d, 0xd2, 0x87, 0x10, 0x5d, 0xea, 0x75, 0x55, 0xb2,
	0x41, 0x46, 0xf2, 0x88, 0x07, 0xa0, 0x8f, 0x61, 0xec, 0xae, 0x57, 0x25, 0x1b, 0x7a, 0xb9, 0x25,
	0xfa, 0x14, 0x92, 0xea, 0xfd, 0xab, 0xd5, 0x4a, 0xa3, 0x31, 0x6c, 0xe4, 0x27, 0x1d, 0x05, 0x37,
	0x8b, 0xab, 0x1a, 0x0d, 0x8b, 0xb2, 0x61, 0x9e, 0xf0, 0x00, 0xd3, 0xdb, 0x01, 0x8c, 0xdc, 0xf5,
	0xe3, 0x53, 0xe4, 0xce, 0x53, 0x97, 0x7a, 0x3d, 0x6f, 0x1d, 0x24, 0xbc, 0x25, 0x67, 0xd9, 0x6f,
	0xd0, 0x7a, 0x48, 0x78, 0x87, 0x3d, 0x73, 0xa3, 0xbb, 0xe6, 0x5c, 0xf5, 0xa6, 0x11, 0xb2, 0x66,
	0x51, 0x30, 0x77, 0x10, 0x5c, 0xf7, 0x5c, 0x6a, 0x63, 0xdf, 0x89, 0x06, 0xd9, 0x38, 0x74, 0x0f,
	0x02, 0x7d, 0x02, 0xf1, 0x85, 0x68, 0x9b, 0x13, 0xdf, 0x3c, 0xb0, 0xeb, 0x2d, 0xac, 0xb0, 0x3b,
	0x53, 0x95, 0x2c, 0xf6, 0x2f, 0x1e, 0x98, 0x4e, 0xe1, 0xfe, 0x5c, 0xa3, 0xb0, 0x52, 0x6d, 0xae,
	0x64, 0x83, 0x2c, 0xc9, 0x48, 0x3e, 0xe4, 0x27, 0x9a, 0xdb, 0xa4, 0xc4, 0x1a, 0x2d, 0xae, 0x18,
	0x64, 0x24, 0x8f, 0x79, 0x87, 0xd3, 0x1f, 0x03, 0x00, 0xe7, 0xf0, 0x5c, 0xd6, 0x36, 0x04, 0xb4,
	0xb0, 0x42, 0xdb, 0x2e, 0x20, 0x0f, 0x4e, 0xbd, 0x90, 0x8d, 0xb4, 0xdd, 0x2f, 0xe4, 0xe1, 0x18,
	0xe6, 0xb0, 0x1f, 0x66, 0x0a, 0xf0, 0x49, 0xda, 0xaf, 0xc1, 0x9e, 0x8f, 0x27, 0xe6, 0x3d, 0x85,
	0x66, 0x70, 0x16, 0xaa, 0x8f, 0xa2, 0xde, 0xa1, 0x0f, 0x29, 0xe2, 0x7d, 0xc9, 0x9d, 0x70, 0xe7,
	0x3b, 0xc3, 0x63, 0x3f, 0xa2, 0x2f, 0xb9, 0x95, 0xdb, 0x32, 0x0c, 0x99, 0xf8, 0x23, 0x27, 0x1a,
	0x7d, 0x06, 0x0f, 0x16, 0x72, 0xb3, 0xc4, 0x93, 0x6c, 0x62, 0x9f, 0xcd, 0xff, 0x8d, 0xd7, 0x6f,
	0x7f, 0xed, 0x53, 0xf2, 0x7b, 0x9f, 0x92, 0x3f, 0xfb, 0x94, 0xdc, 0xfe, 0x4d, 0xef, 0x7d, 0x7e,
	0xb9, 0x56, 0xdb, 0xeb, 0xf5, 0xac, 0x96, 0x9b, 0x6b, 0x21, 0x67, 0x52, 0x15, 0x37, 0xcf, 0x0b,
	0x8d, 0x5b, 0x65, 0x0a, 0xd1, 0x14, 0xfe, 0x2f, 0xbe, 0x34, 0xa8, 0x6f, 0xe4, 0x12, 0x4d, 0x71,
	0xfc, 0x02, 0xbe, 0x8c, 0x7d, 0xfd, 0xe2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x43, 0x3b,
	0x3a, 0x16, 0x03, 0x00, 0x00,
}
