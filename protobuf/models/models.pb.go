// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/models.proto

package models // import "github.com/ubclaunchpad/pinpoint/protobuf/models"

/*
package model defines model datatypes
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Verified             bool     `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *User) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

type Club struct {
	ClubID               string   `protobuf:"bytes,1,opt,name=clubID,proto3" json:"clubID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Club) Reset()         { *m = Club{} }
func (m *Club) String() string { return proto.CompactTextString(m) }
func (*Club) ProtoMessage()    {}
func (*Club) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{1}
}
func (m *Club) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Club.Unmarshal(m, b)
}
func (m *Club) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Club.Marshal(b, m, deterministic)
}
func (dst *Club) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Club.Merge(dst, src)
}
func (m *Club) XXX_Size() int {
	return xxx_messageInfo_Club.Size(m)
}
func (m *Club) XXX_DiscardUnknown() {
	xxx_messageInfo_Club.DiscardUnknown(m)
}

var xxx_messageInfo_Club proto.InternalMessageInfo

func (m *Club) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func (m *Club) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Club) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ClubUser struct {
	ClubID               string   `protobuf:"bytes,1,opt,name=clubID,proto3" json:"clubID,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClubUser) Reset()         { *m = ClubUser{} }
func (m *ClubUser) String() string { return proto.CompactTextString(m) }
func (*ClubUser) ProtoMessage()    {}
func (*ClubUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{2}
}
func (m *ClubUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubUser.Unmarshal(m, b)
}
func (m *ClubUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubUser.Marshal(b, m, deterministic)
}
func (dst *ClubUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubUser.Merge(dst, src)
}
func (m *ClubUser) XXX_Size() int {
	return xxx_messageInfo_ClubUser.Size(m)
}
func (m *ClubUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubUser.DiscardUnknown(m)
}

var xxx_messageInfo_ClubUser proto.InternalMessageInfo

func (m *ClubUser) GetClubID() string {
	if m != nil {
		return m.ClubID
	}
	return ""
}

func (m *ClubUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ClubUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClubUser) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type EmailVerification struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Expiry               int64    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailVerification) Reset()         { *m = EmailVerification{} }
func (m *EmailVerification) String() string { return proto.CompactTextString(m) }
func (*EmailVerification) ProtoMessage()    {}
func (*EmailVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{3}
}
func (m *EmailVerification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailVerification.Unmarshal(m, b)
}
func (m *EmailVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailVerification.Marshal(b, m, deterministic)
}
func (dst *EmailVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailVerification.Merge(dst, src)
}
func (m *EmailVerification) XXX_Size() int {
	return xxx_messageInfo_EmailVerification.Size(m)
}
func (m *EmailVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailVerification.DiscardUnknown(m)
}

var xxx_messageInfo_EmailVerification proto.InternalMessageInfo

func (m *EmailVerification) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *EmailVerification) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailVerification) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type EventProps struct {
	Period               string                   `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	EventID              string                   `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Name                 string                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Club                 string                   `protobuf:"bytes,4,opt,name=club,proto3" json:"club,omitempty"`
	Description          string                   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Fields               []*EventProps_FieldProps `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EventProps) Reset()         { *m = EventProps{} }
func (m *EventProps) String() string { return proto.CompactTextString(m) }
func (*EventProps) ProtoMessage()    {}
func (*EventProps) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{4}
}
func (m *EventProps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProps.Unmarshal(m, b)
}
func (m *EventProps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProps.Marshal(b, m, deterministic)
}
func (dst *EventProps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProps.Merge(dst, src)
}
func (m *EventProps) XXX_Size() int {
	return xxx_messageInfo_EventProps.Size(m)
}
func (m *EventProps) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProps.DiscardUnknown(m)
}

var xxx_messageInfo_EventProps proto.InternalMessageInfo

func (m *EventProps) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *EventProps) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *EventProps) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventProps) GetClub() string {
	if m != nil {
		return m.Club
	}
	return ""
}

func (m *EventProps) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventProps) GetFields() []*EventProps_FieldProps {
	if m != nil {
		return m.Fields
	}
	return nil
}

type EventProps_FieldProps struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required    bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Blurb       string `protobuf:"bytes,3,opt,name=blurb,proto3" json:"blurb,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are valid to be assigned to Prsoperties:
	//	*EventProps_FieldProps_LongText_
	//	*EventProps_FieldProps_ShortText_
	Prsoperties          isEventProps_FieldProps_Prsoperties `protobuf_oneof:"prsoperties"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *EventProps_FieldProps) Reset()         { *m = EventProps_FieldProps{} }
func (m *EventProps_FieldProps) String() string { return proto.CompactTextString(m) }
func (*EventProps_FieldProps) ProtoMessage()    {}
func (*EventProps_FieldProps) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{4, 0}
}
func (m *EventProps_FieldProps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProps_FieldProps.Unmarshal(m, b)
}
func (m *EventProps_FieldProps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProps_FieldProps.Marshal(b, m, deterministic)
}
func (dst *EventProps_FieldProps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProps_FieldProps.Merge(dst, src)
}
func (m *EventProps_FieldProps) XXX_Size() int {
	return xxx_messageInfo_EventProps_FieldProps.Size(m)
}
func (m *EventProps_FieldProps) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProps_FieldProps.DiscardUnknown(m)
}

var xxx_messageInfo_EventProps_FieldProps proto.InternalMessageInfo

func (m *EventProps_FieldProps) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventProps_FieldProps) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *EventProps_FieldProps) GetBlurb() string {
	if m != nil {
		return m.Blurb
	}
	return ""
}

func (m *EventProps_FieldProps) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type isEventProps_FieldProps_Prsoperties interface {
	isEventProps_FieldProps_Prsoperties()
}

type EventProps_FieldProps_LongText_ struct {
	LongText *EventProps_FieldProps_LongText `protobuf:"bytes,5,opt,name=long_text,json=longText,proto3,oneof"`
}

type EventProps_FieldProps_ShortText_ struct {
	ShortText *EventProps_FieldProps_ShortText `protobuf:"bytes,6,opt,name=short_text,json=shortText,proto3,oneof"`
}

func (*EventProps_FieldProps_LongText_) isEventProps_FieldProps_Prsoperties() {}

func (*EventProps_FieldProps_ShortText_) isEventProps_FieldProps_Prsoperties() {}

func (m *EventProps_FieldProps) GetPrsoperties() isEventProps_FieldProps_Prsoperties {
	if m != nil {
		return m.Prsoperties
	}
	return nil
}

func (m *EventProps_FieldProps) GetLongText() *EventProps_FieldProps_LongText {
	if x, ok := m.GetPrsoperties().(*EventProps_FieldProps_LongText_); ok {
		return x.LongText
	}
	return nil
}

func (m *EventProps_FieldProps) GetShortText() *EventProps_FieldProps_ShortText {
	if x, ok := m.GetPrsoperties().(*EventProps_FieldProps_ShortText_); ok {
		return x.ShortText
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventProps_FieldProps) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventProps_FieldProps_OneofMarshaler, _EventProps_FieldProps_OneofUnmarshaler, _EventProps_FieldProps_OneofSizer, []interface{}{
		(*EventProps_FieldProps_LongText_)(nil),
		(*EventProps_FieldProps_ShortText_)(nil),
	}
}

func _EventProps_FieldProps_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventProps_FieldProps)
	// prsoperties
	switch x := m.Prsoperties.(type) {
	case *EventProps_FieldProps_LongText_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LongText); err != nil {
			return err
		}
	case *EventProps_FieldProps_ShortText_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShortText); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventProps_FieldProps.Prsoperties has unexpected type %T", x)
	}
	return nil
}

func _EventProps_FieldProps_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventProps_FieldProps)
	switch tag {
	case 5: // prsoperties.long_text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventProps_FieldProps_LongText)
		err := b.DecodeMessage(msg)
		m.Prsoperties = &EventProps_FieldProps_LongText_{msg}
		return true, err
	case 6: // prsoperties.short_text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventProps_FieldProps_ShortText)
		err := b.DecodeMessage(msg)
		m.Prsoperties = &EventProps_FieldProps_ShortText_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventProps_FieldProps_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventProps_FieldProps)
	// prsoperties
	switch x := m.Prsoperties.(type) {
	case *EventProps_FieldProps_LongText_:
		s := proto.Size(x.LongText)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventProps_FieldProps_ShortText_:
		s := proto.Size(x.ShortText)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EventProps_FieldProps_LongText struct {
	MaxLen               string   `protobuf:"bytes,1,opt,name=maxLen,proto3" json:"maxLen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventProps_FieldProps_LongText) Reset()         { *m = EventProps_FieldProps_LongText{} }
func (m *EventProps_FieldProps_LongText) String() string { return proto.CompactTextString(m) }
func (*EventProps_FieldProps_LongText) ProtoMessage()    {}
func (*EventProps_FieldProps_LongText) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{4, 0, 0}
}
func (m *EventProps_FieldProps_LongText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProps_FieldProps_LongText.Unmarshal(m, b)
}
func (m *EventProps_FieldProps_LongText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProps_FieldProps_LongText.Marshal(b, m, deterministic)
}
func (dst *EventProps_FieldProps_LongText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProps_FieldProps_LongText.Merge(dst, src)
}
func (m *EventProps_FieldProps_LongText) XXX_Size() int {
	return xxx_messageInfo_EventProps_FieldProps_LongText.Size(m)
}
func (m *EventProps_FieldProps_LongText) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProps_FieldProps_LongText.DiscardUnknown(m)
}

var xxx_messageInfo_EventProps_FieldProps_LongText proto.InternalMessageInfo

func (m *EventProps_FieldProps_LongText) GetMaxLen() string {
	if m != nil {
		return m.MaxLen
	}
	return ""
}

type EventProps_FieldProps_ShortText struct {
	MaxLen               string   `protobuf:"bytes,1,opt,name=maxLen,proto3" json:"maxLen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventProps_FieldProps_ShortText) Reset()         { *m = EventProps_FieldProps_ShortText{} }
func (m *EventProps_FieldProps_ShortText) String() string { return proto.CompactTextString(m) }
func (*EventProps_FieldProps_ShortText) ProtoMessage()    {}
func (*EventProps_FieldProps_ShortText) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{4, 0, 1}
}
func (m *EventProps_FieldProps_ShortText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProps_FieldProps_ShortText.Unmarshal(m, b)
}
func (m *EventProps_FieldProps_ShortText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProps_FieldProps_ShortText.Marshal(b, m, deterministic)
}
func (dst *EventProps_FieldProps_ShortText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProps_FieldProps_ShortText.Merge(dst, src)
}
func (m *EventProps_FieldProps_ShortText) XXX_Size() int {
	return xxx_messageInfo_EventProps_FieldProps_ShortText.Size(m)
}
func (m *EventProps_FieldProps_ShortText) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProps_FieldProps_ShortText.DiscardUnknown(m)
}

var xxx_messageInfo_EventProps_FieldProps_ShortText proto.InternalMessageInfo

func (m *EventProps_FieldProps_ShortText) GetMaxLen() string {
	if m != nil {
		return m.MaxLen
	}
	return ""
}

type Applicant struct {
	Period               string   `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Applicant) Reset()         { *m = Applicant{} }
func (m *Applicant) String() string { return proto.CompactTextString(m) }
func (*Applicant) ProtoMessage()    {}
func (*Applicant) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{5}
}
func (m *Applicant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Applicant.Unmarshal(m, b)
}
func (m *Applicant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Applicant.Marshal(b, m, deterministic)
}
func (dst *Applicant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Applicant.Merge(dst, src)
}
func (m *Applicant) XXX_Size() int {
	return xxx_messageInfo_Applicant.Size(m)
}
func (m *Applicant) XXX_DiscardUnknown() {
	xxx_messageInfo_Applicant.DiscardUnknown(m)
}

var xxx_messageInfo_Applicant proto.InternalMessageInfo

func (m *Applicant) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *Applicant) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Applicant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Applicant) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Application struct {
	Period               string                 `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	EventID              string                 `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Email                string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Entries              map[string]*FieldEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{6}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *Application) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *Application) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetEntries() map[string]*FieldEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type FieldEntry struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldEntry) Reset()         { *m = FieldEntry{} }
func (m *FieldEntry) String() string { return proto.CompactTextString(m) }
func (*FieldEntry) ProtoMessage()    {}
func (*FieldEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{7}
}
func (m *FieldEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldEntry.Unmarshal(m, b)
}
func (m *FieldEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldEntry.Marshal(b, m, deterministic)
}
func (dst *FieldEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldEntry.Merge(dst, src)
}
func (m *FieldEntry) XXX_Size() int {
	return xxx_messageInfo_FieldEntry.Size(m)
}
func (m *FieldEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FieldEntry proto.InternalMessageInfo

func (m *FieldEntry) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Tag struct {
	Period               string   `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	TagName              string   `protobuf:"bytes,2,opt,name=tagName,proto3" json:"tagName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_2add85134884b5c1, []int{8}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *Tag) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "models.User")
	proto.RegisterType((*Club)(nil), "models.Club")
	proto.RegisterType((*ClubUser)(nil), "models.ClubUser")
	proto.RegisterType((*EmailVerification)(nil), "models.EmailVerification")
	proto.RegisterType((*EventProps)(nil), "models.EventProps")
	proto.RegisterType((*EventProps_FieldProps)(nil), "models.EventProps.FieldProps")
	proto.RegisterType((*EventProps_FieldProps_LongText)(nil), "models.EventProps.FieldProps.LongText")
	proto.RegisterType((*EventProps_FieldProps_ShortText)(nil), "models.EventProps.FieldProps.ShortText")
	proto.RegisterType((*Applicant)(nil), "models.Applicant")
	proto.RegisterType((*Application)(nil), "models.Application")
	proto.RegisterMapType((map[string]*FieldEntry)(nil), "models.Application.EntriesEntry")
	proto.RegisterType((*FieldEntry)(nil), "models.FieldEntry")
	proto.RegisterType((*Tag)(nil), "models.Tag")
}

func init() { proto.RegisterFile("models/models.proto", fileDescriptor_models_2add85134884b5c1) }

var fileDescriptor_models_2add85134884b5c1 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6b, 0xdb, 0x3e,
	0x14, 0xfd, 0x39, 0x76, 0x5c, 0xe7, 0xa6, 0x3f, 0xd8, 0xbc, 0x52, 0x4c, 0x60, 0x10, 0x3c, 0xd8,
	0xf2, 0x94, 0x8c, 0x8c, 0xb1, 0xd1, 0xb7, 0x75, 0xcd, 0x68, 0xa1, 0x94, 0xe1, 0xb5, 0x7b, 0xd8,
	0xcb, 0x2a, 0xdb, 0x6a, 0x22, 0xa6, 0x58, 0x9a, 0x24, 0x97, 0xf4, 0xd3, 0xec, 0x93, 0xed, 0x63,
	0xec, 0x7d, 0xe8, 0x8f, 0xdd, 0xac, 0x4d, 0x0a, 0x7d, 0xf2, 0x3d, 0xd7, 0x57, 0xe7, 0x9c, 0x7b,
	0x75, 0x11, 0x3c, 0x5b, 0xb2, 0x12, 0x53, 0x39, 0xb1, 0x9f, 0x31, 0x17, 0x4c, 0xb1, 0x38, 0xb4,
	0x28, 0xbd, 0x84, 0xe0, 0x42, 0x62, 0x11, 0xef, 0x41, 0x17, 0x2f, 0x11, 0xa1, 0x89, 0x37, 0xf4,
	0x46, 0xbd, 0xcc, 0x82, 0x38, 0x86, 0xa0, 0x42, 0x4b, 0x9c, 0x74, 0x4c, 0xd2, 0xc4, 0x3a, 0xb7,
	0x40, 0x72, 0x91, 0xf8, 0x36, 0xa7, 0xe3, 0x78, 0x00, 0xd1, 0x35, 0x16, 0xe4, 0x8a, 0xe0, 0x32,
	0x09, 0x86, 0xde, 0x28, 0xca, 0x5a, 0x9c, 0x9e, 0x43, 0xf0, 0x91, 0xd6, 0x79, 0xbc, 0x0f, 0x61,
	0x41, 0xeb, 0xfc, 0xe4, 0xc8, 0x49, 0x38, 0xb4, 0x51, 0x63, 0x08, 0xfd, 0x12, 0xcb, 0x42, 0x10,
	0xae, 0x08, 0xab, 0x9c, 0xd4, 0x7a, 0x2a, 0xbd, 0x84, 0x48, 0xb3, 0x1a, 0xef, 0xdb, 0x98, 0xdb,
	0x9e, 0x3a, 0x9b, 0x7a, 0xf2, 0xff, 0xed, 0x49, 0x30, 0x8a, 0x8d, 0xf7, 0x5e, 0x66, 0xe2, 0xf4,
	0x02, 0x9e, 0xce, 0xf4, 0x81, 0xaf, 0xa6, 0x91, 0x02, 0x69, 0xd9, 0xb6, 0x79, 0x6f, 0xad, 0xf9,
	0xcd, 0x32, 0xfb, 0x10, 0xe2, 0x15, 0x27, 0xe2, 0xc6, 0x08, 0xf9, 0x99, 0x43, 0xe9, 0xaf, 0x00,
	0x60, 0x76, 0x8d, 0x2b, 0xf5, 0x59, 0x30, 0x2e, 0x75, 0x19, 0xc7, 0x82, 0xb0, 0xb2, 0xf1, 0x6e,
	0x51, 0x9c, 0xc0, 0x0e, 0xd6, 0x55, 0x27, 0x47, 0x8e, 0xb6, 0x81, 0xdb, 0xfc, 0xeb, 0x9e, 0x1b,
	0xff, 0x3a, 0xbe, 0x3b, 0xc3, 0xee, 0xbd, 0x19, 0xc6, 0x6f, 0x21, 0xbc, 0x22, 0x98, 0x96, 0x32,
	0x09, 0x87, 0xfe, 0xa8, 0x3f, 0x7d, 0x3e, 0x76, 0x2b, 0x72, 0xeb, 0x6f, 0xfc, 0x49, 0x17, 0x98,
	0x30, 0x73, 0xc5, 0x83, 0xdf, 0x1d, 0x80, 0xdb, 0x74, 0xeb, 0xc7, 0x5b, 0xf3, 0x33, 0x80, 0x48,
	0xe0, 0x9f, 0x35, 0x11, 0xb8, 0x34, 0xf6, 0xa3, 0xac, 0xc5, 0x7a, 0x5c, 0x39, 0xad, 0x45, 0xee,
	0x1a, 0xb0, 0xe0, 0xae, 0xdb, 0xe0, 0xbe, 0xdb, 0x19, 0xf4, 0x28, 0xab, 0xe6, 0xdf, 0x15, 0x5e,
	0x29, 0xd3, 0x4d, 0x7f, 0xfa, 0xf2, 0x41, 0xc3, 0xe3, 0x53, 0x56, 0xcd, 0xcf, 0xf1, 0x4a, 0x1d,
	0xff, 0x97, 0x45, 0xd4, 0xc5, 0xf1, 0x31, 0x80, 0x5c, 0x30, 0xa1, 0x2c, 0x4f, 0x68, 0x78, 0x5e,
	0x3d, 0xcc, 0xf3, 0x45, 0xd7, 0x3b, 0xa2, 0x9e, 0x6c, 0xc0, 0x20, 0x85, 0xa8, 0x51, 0xd0, 0xd7,
	0xb8, 0x44, 0xab, 0x53, 0x5c, 0x35, 0xd7, 0x68, 0xd1, 0xe0, 0x05, 0xf4, 0xda, 0xd3, 0xdb, 0x8a,
	0x0e, 0xff, 0x87, 0x3e, 0x17, 0x92, 0x71, 0x2c, 0x14, 0xc1, 0x32, 0x45, 0xd0, 0xfb, 0xc0, 0x39,
	0x25, 0x05, 0xaa, 0xd4, 0xd6, 0xfd, 0x78, 0xd4, 0x6e, 0x2b, 0x34, 0x97, 0x49, 0x30, 0xf4, 0x75,
	0x4e, 0xc7, 0xe9, 0x1f, 0x0f, 0xfa, 0x4e, 0xc3, 0xcc, 0xf6, 0xf1, 0x5b, 0xd8, 0xea, 0xfb, 0x9b,
	0xf4, 0x83, 0x35, 0xfd, 0x03, 0xd8, 0xc1, 0x95, 0x12, 0x04, 0xcb, 0xa4, 0x6b, 0xd6, 0x6c, 0xd8,
	0x4c, 0x7b, 0xcd, 0xc1, 0x78, 0x66, 0x4b, 0xf4, 0xe7, 0x26, 0x6b, 0x0e, 0x0c, 0xce, 0x60, 0x77,
	0xfd, 0x47, 0xfc, 0x04, 0xfc, 0x1f, 0xf8, 0xc6, 0x99, 0xd4, 0x61, 0x3c, 0x82, 0xee, 0x35, 0xa2,
	0xb5, 0x7d, 0x3e, 0xfa, 0xd3, 0xb8, 0xe1, 0x36, 0xd7, 0x67, 0xd9, 0x6c, 0xc1, 0x41, 0xe7, 0xbd,
	0x97, 0xa6, 0x6e, 0x73, 0x2d, 0xdb, 0x5e, 0x73, 0x56, 0xf3, 0xed, 0xba, 0xba, 0xf4, 0x1d, 0xf8,
	0xe7, 0x68, 0xfe, 0xd0, 0x48, 0x14, 0x9a, 0x9f, 0xdd, 0xbe, 0x58, 0x0d, 0x3c, 0x9c, 0x7e, 0x7b,
	0x3d, 0x27, 0x6a, 0x51, 0xe7, 0xe3, 0x82, 0x2d, 0x27, 0x75, 0x5e, 0x50, 0x54, 0x57, 0xc5, 0x82,
	0xa3, 0x72, 0xc2, 0x49, 0xc5, 0x19, 0xa9, 0xd4, 0xc4, 0x3c, 0xbe, 0x79, 0x7d, 0xe5, 0x1e, 0xe3,
	0x3c, 0x34, 0x89, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe1, 0xb6, 0xd0, 0xa4, 0x05,
	0x00, 0x00,
}
