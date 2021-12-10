// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/proposals.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/provenance-io/provenance/x/marker/types"
	_ "github.com/regen-network/cosmos-proto"
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

// AddMsgBasedFeeProposal defines a governance proposal to add additional msg based fee
type AddMsgBasedFeeProposal struct {
	Title         string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl    string     `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	AdditionalFee types.Coin `protobuf:"bytes,4,opt,name=additional_fee,json=additionalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"additional_fee" yaml:"additional_fee"`
}

func (m *AddMsgBasedFeeProposal) Reset()      { *m = AddMsgBasedFeeProposal{} }
func (*AddMsgBasedFeeProposal) ProtoMessage() {}
func (*AddMsgBasedFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{0}
}
func (m *AddMsgBasedFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddMsgBasedFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddMsgBasedFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddMsgBasedFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMsgBasedFeeProposal.Merge(m, src)
}
func (m *AddMsgBasedFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddMsgBasedFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMsgBasedFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddMsgBasedFeeProposal proto.InternalMessageInfo

func (m *AddMsgBasedFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddMsgBasedFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddMsgBasedFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *AddMsgBasedFeeProposal) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

// UpdateMsgBasedFeeProposal defines a governance proposal to update a current msg based fee
type UpdateMsgBasedFeeProposal struct {
	Title         string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl    string     `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	AdditionalFee types.Coin `protobuf:"bytes,4,opt,name=additional_fee,json=additionalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"additional_fee" yaml:"additional_fee"`
}

func (m *UpdateMsgBasedFeeProposal) Reset()      { *m = UpdateMsgBasedFeeProposal{} }
func (*UpdateMsgBasedFeeProposal) ProtoMessage() {}
func (*UpdateMsgBasedFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{1}
}
func (m *UpdateMsgBasedFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMsgBasedFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMsgBasedFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateMsgBasedFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMsgBasedFeeProposal.Merge(m, src)
}
func (m *UpdateMsgBasedFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMsgBasedFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMsgBasedFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMsgBasedFeeProposal proto.InternalMessageInfo

func (m *UpdateMsgBasedFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateMsgBasedFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateMsgBasedFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *UpdateMsgBasedFeeProposal) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

// RemoveMsgBasedFeeProposal defines a governance proposal to delete a current msg based fee
type RemoveMsgBasedFeeProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl  string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (m *RemoveMsgBasedFeeProposal) Reset()      { *m = RemoveMsgBasedFeeProposal{} }
func (*RemoveMsgBasedFeeProposal) ProtoMessage() {}
func (*RemoveMsgBasedFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{2}
}
func (m *RemoveMsgBasedFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveMsgBasedFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveMsgBasedFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveMsgBasedFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMsgBasedFeeProposal.Merge(m, src)
}
func (m *RemoveMsgBasedFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveMsgBasedFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMsgBasedFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMsgBasedFeeProposal proto.InternalMessageInfo

func (m *RemoveMsgBasedFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveMsgBasedFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveMsgBasedFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*AddMsgBasedFeeProposal)(nil), "provenance.msgfees.v1.AddMsgBasedFeeProposal")
	proto.RegisterType((*UpdateMsgBasedFeeProposal)(nil), "provenance.msgfees.v1.UpdateMsgBasedFeeProposal")
	proto.RegisterType((*RemoveMsgBasedFeeProposal)(nil), "provenance.msgfees.v1.RemoveMsgBasedFeeProposal")
}

func init() {
	proto.RegisterFile("provenance/msgfees/v1/proposals.proto", fileDescriptor_a2e168825d6c34a4)
}

var fileDescriptor_a2e168825d6c34a4 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0xe1, 0x8f, 0xe0, 0x0a, 0x0c, 0x56, 0x8b, 0x9c, 0x0c, 0x76, 0x88, 0x04, 0xca,
	0x52, 0x9f, 0x42, 0xb7, 0x6e, 0x04, 0xa9, 0x12, 0x03, 0x52, 0x15, 0xd1, 0x85, 0x25, 0x3a, 0xdb,
	0x6f, 0x8e, 0x53, 0x6c, 0x9f, 0x75, 0xef, 0xc5, 0x4a, 0x26, 0x76, 0x10, 0x12, 0x23, 0x63, 0x67,
	0x3e, 0x49, 0xc7, 0x8e, 0x4c, 0x05, 0x25, 0x0b, 0x33, 0x9f, 0x00, 0xf9, 0x7c, 0x24, 0x41, 0x65,
	0x66, 0xea, 0xe4, 0x7b, 0xfc, 0xfc, 0xce, 0x7a, 0xde, 0x47, 0x7a, 0x4d, 0x9f, 0x96, 0x5a, 0x55,
	0x50, 0xf0, 0x22, 0x01, 0x96, 0xa3, 0x98, 0x02, 0x20, 0xab, 0x86, 0xac, 0xd4, 0xaa, 0x54, 0xc8,
	0x33, 0x8c, 0x4a, 0xad, 0x8c, 0xf2, 0x0e, 0xb6, 0x58, 0xe4, 0xb0, 0xa8, 0x1a, 0x76, 0xf7, 0x85,
	0x12, 0xca, 0x12, 0xac, 0x3e, 0x35, 0x70, 0x37, 0x48, 0x14, 0xe6, 0x0a, 0x59, 0xcc, 0x8b, 0x19,
	0xab, 0x86, 0x31, 0x18, 0x3e, 0xb4, 0xe2, 0x9a, 0x8f, 0xb0, 0xf1, 0x13, 0x25, 0x0b, 0xe7, 0x3f,
	0xd9, 0xcd, 0xc4, 0xf5, 0x0c, 0x74, 0x1d, 0xa9, 0x39, 0x39, 0xe4, 0xd9, 0x3f, 0x11, 0x9e, 0x24,
	0x80, 0x28, 0x34, 0x2f, 0x8c, 0xe3, 0x3a, 0x42, 0x29, 0x91, 0x01, 0xb3, 0x2a, 0x9e, 0x4f, 0x19,
	0x2f, 0x96, 0xce, 0xea, 0xba, 0x14, 0x66, 0xb1, 0xc9, 0x60, 0x16, 0x7f, 0xae, 0x35, 0xde, 0xa4,
	0x19, 0xad, 0x11, 0x8d, 0xd5, 0xff, 0xd0, 0xa6, 0x8f, 0x5f, 0xa4, 0xe9, 0x6b, 0x14, 0x23, 0x8e,
	0x90, 0x9e, 0x00, 0x9c, 0xba, 0xae, 0xbc, 0x7d, 0x7a, 0xc7, 0x48, 0x93, 0x81, 0x4f, 0x7a, 0x64,
	0x70, 0x7f, 0xdc, 0x08, 0xaf, 0x47, 0xf7, 0x52, 0xc0, 0x44, 0xcb, 0xd2, 0x48, 0x55, 0xf8, 0x6d,
	0xeb, 0xed, 0xbe, 0xf2, 0x7a, 0xf4, 0x41, 0x8e, 0x62, 0x62, 0x96, 0x25, 0x4c, 0xe6, 0x3a, 0xf3,
	0x6f, 0x59, 0x84, 0xe6, 0x28, 0xde, 0x2c, 0x4b, 0x38, 0xd3, 0x99, 0xf7, 0x91, 0xd0, 0x47, 0x3c,
	0x4d, 0x65, 0x8d, 0xf3, 0x6c, 0x32, 0x05, 0xf0, 0x6f, 0xf7, 0xc8, 0x60, 0xef, 0x79, 0x27, 0x72,
	0xe1, 0xea, 0x2e, 0x23, 0x37, 0x47, 0xf4, 0x52, 0xc9, 0x62, 0xf4, 0xea, 0xe2, 0x2a, 0x6c, 0xfd,
	0xba, 0x0a, 0x0f, 0x96, 0x3c, 0xcf, 0x8e, 0xfb, 0x7f, 0x5f, 0xef, 0x7f, 0xfd, 0x1e, 0x0e, 0x84,
	0x34, 0xef, 0xe6, 0x71, 0x94, 0xa8, 0xdc, 0x8d, 0xe8, 0x1e, 0x87, 0x98, 0xce, 0x58, 0x9d, 0x06,
	0xed, 0x97, 0x70, 0xfc, 0x70, 0x7b, 0xf9, 0x04, 0xe0, 0xf8, 0xde, 0x97, 0xf3, 0xb0, 0xf5, 0xf3,
	0x3c, 0x24, 0xfd, 0x4f, 0x6d, 0xda, 0x39, 0x2b, 0x53, 0x6e, 0xe0, 0xa6, 0x8f, 0xba, 0x8f, 0xf7,
	0xb4, 0x33, 0x86, 0x5c, 0x55, 0xff, 0xb7, 0x8e, 0x6d, 0x80, 0x91, 0xbc, 0x58, 0x05, 0xe4, 0x72,
	0x15, 0x90, 0x1f, 0xab, 0x80, 0x7c, 0x5e, 0x07, 0xad, 0xcb, 0x75, 0xd0, 0xfa, 0xb6, 0x0e, 0x5a,
	0xd4, 0x97, 0x76, 0x3d, 0xaf, 0x2f, 0xf1, 0x29, 0x79, 0x7b, 0xb4, 0x33, 0xf9, 0x96, 0x39, 0x94,
	0x6a, 0x47, 0xb1, 0xc5, 0xe6, 0xff, 0x60, 0xab, 0x88, 0xef, 0xda, 0x7d, 0x38, 0xfa, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x00, 0x1c, 0xf1, 0x9b, 0x42, 0x04, 0x00, 0x00,
}

func (this *AddMsgBasedFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddMsgBasedFeeProposal)
	if !ok {
		that2, ok := that.(AddMsgBasedFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	if !this.AdditionalFee.Equal(&that1.AdditionalFee) {
		return false
	}
	return true
}
func (this *UpdateMsgBasedFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateMsgBasedFeeProposal)
	if !ok {
		that2, ok := that.(UpdateMsgBasedFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	if !this.AdditionalFee.Equal(&that1.AdditionalFee) {
		return false
	}
	return true
}
func (this *RemoveMsgBasedFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveMsgBasedFeeProposal)
	if !ok {
		that2, ok := that.(RemoveMsgBasedFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	return true
}
func (m *AddMsgBasedFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddMsgBasedFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddMsgBasedFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateMsgBasedFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMsgBasedFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateMsgBasedFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveMsgBasedFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveMsgBasedFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveMsgBasedFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddMsgBasedFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *UpdateMsgBasedFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *RemoveMsgBasedFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddMsgBasedFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddMsgBasedFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddMsgBasedFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *UpdateMsgBasedFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: UpdateMsgBasedFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMsgBasedFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *RemoveMsgBasedFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: RemoveMsgBasedFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveMsgBasedFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)
