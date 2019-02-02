// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/models.proto

package models // import "github.com/ubclaunchpad/pinpoint/protobuf/models"

/*
package model defines model datatypes
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{0}
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{1}
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{2}
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
	Hash                 string               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Expiry               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EmailVerification) Reset()         { *m = EmailVerification{} }
func (m *EmailVerification) String() string { return proto.CompactTextString(m) }
func (*EmailVerification) ProtoMessage()    {}
func (*EmailVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{3}
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

func (m *EmailVerification) GetExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type Event struct {
	Period               string   `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	EventID              string   `protobuf:"bytes,2,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Fields               []*Field `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{4}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

func (m *Event) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Field struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Types that are valid to be assigned to Properties:
	//	*Field_LongText
	//	*Field_ShortText
	Properties           isField_Properties `protobuf_oneof:"properties"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{5}
}
func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (dst *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(dst, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

type isField_Properties interface {
	isField_Properties()
}

type Field_LongText struct {
	LongText *LongText `protobuf:"bytes,3,opt,name=long_text,json=longText,proto3,oneof"`
}

type Field_ShortText struct {
	ShortText *ShortText `protobuf:"bytes,4,opt,name=short_text,json=shortText,proto3,oneof"`
}

func (*Field_LongText) isField_Properties() {}

func (*Field_ShortText) isField_Properties() {}

func (m *Field) GetProperties() isField_Properties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Field) GetLongText() *LongText {
	if x, ok := m.GetProperties().(*Field_LongText); ok {
		return x.LongText
	}
	return nil
}

func (m *Field) GetShortText() *ShortText {
	if x, ok := m.GetProperties().(*Field_ShortText); ok {
		return x.ShortText
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Field_OneofMarshaler, _Field_OneofUnmarshaler, _Field_OneofSizer, []interface{}{
		(*Field_LongText)(nil),
		(*Field_ShortText)(nil),
	}
}

func _Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Field)
	// properties
	switch x := m.Properties.(type) {
	case *Field_LongText:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LongText); err != nil {
			return err
		}
	case *Field_ShortText:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShortText); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Field.Properties has unexpected type %T", x)
	}
	return nil
}

func _Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Field)
	switch tag {
	case 3: // properties.long_text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LongText)
		err := b.DecodeMessage(msg)
		m.Properties = &Field_LongText{msg}
		return true, err
	case 4: // properties.short_text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ShortText)
		err := b.DecodeMessage(msg)
		m.Properties = &Field_ShortText{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Field)
	// properties
	switch x := m.Properties.(type) {
	case *Field_LongText:
		s := proto.Size(x.LongText)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Field_ShortText:
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

type LongText struct {
	ImLong               string   `protobuf:"bytes,1,opt,name=im_long,json=imLong,proto3" json:"im_long,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongText) Reset()         { *m = LongText{} }
func (m *LongText) String() string { return proto.CompactTextString(m) }
func (*LongText) ProtoMessage()    {}
func (*LongText) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{6}
}
func (m *LongText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongText.Unmarshal(m, b)
}
func (m *LongText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongText.Marshal(b, m, deterministic)
}
func (dst *LongText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongText.Merge(dst, src)
}
func (m *LongText) XXX_Size() int {
	return xxx_messageInfo_LongText.Size(m)
}
func (m *LongText) XXX_DiscardUnknown() {
	xxx_messageInfo_LongText.DiscardUnknown(m)
}

var xxx_messageInfo_LongText proto.InternalMessageInfo

func (m *LongText) GetImLong() string {
	if m != nil {
		return m.ImLong
	}
	return ""
}

type ShortText struct {
	ImShort              string   `protobuf:"bytes,1,opt,name=im_short,json=imShort,proto3" json:"im_short,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShortText) Reset()         { *m = ShortText{} }
func (m *ShortText) String() string { return proto.CompactTextString(m) }
func (*ShortText) ProtoMessage()    {}
func (*ShortText) Descriptor() ([]byte, []int) {
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{7}
}
func (m *ShortText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShortText.Unmarshal(m, b)
}
func (m *ShortText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShortText.Marshal(b, m, deterministic)
}
func (dst *ShortText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShortText.Merge(dst, src)
}
func (m *ShortText) XXX_Size() int {
	return xxx_messageInfo_ShortText.Size(m)
}
func (m *ShortText) XXX_DiscardUnknown() {
	xxx_messageInfo_ShortText.DiscardUnknown(m)
}

var xxx_messageInfo_ShortText proto.InternalMessageInfo

func (m *ShortText) GetImShort() string {
	if m != nil {
		return m.ImShort
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{8}
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{9}
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{10}
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
	return fileDescriptor_models_1c40b9ecd4f58ff2, []int{11}
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
	proto.RegisterType((*Event)(nil), "models.Event")
	proto.RegisterType((*Field)(nil), "models.Field")
	proto.RegisterType((*LongText)(nil), "models.LongText")
	proto.RegisterType((*ShortText)(nil), "models.ShortText")
	proto.RegisterType((*Applicant)(nil), "models.Applicant")
	proto.RegisterType((*Application)(nil), "models.Application")
	proto.RegisterMapType((map[string]*FieldEntry)(nil), "models.Application.EntriesEntry")
	proto.RegisterType((*FieldEntry)(nil), "models.FieldEntry")
	proto.RegisterType((*Tag)(nil), "models.Tag")
}

func init() { proto.RegisterFile("models/models.proto", fileDescriptor_models_1c40b9ecd4f58ff2) }

var fileDescriptor_models_1c40b9ecd4f58ff2 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0x93, 0x38, 0x93, 0x22, 0xb5, 0x0b, 0x02, 0x93, 0x0b, 0xd1, 0x22, 0x50, 0x4e,
	0x0e, 0x0a, 0x07, 0x50, 0x6f, 0x94, 0x16, 0x51, 0x09, 0xf5, 0x60, 0x02, 0x07, 0x2e, 0xed, 0xc6,
	0xde, 0x3a, 0x2b, 0x6c, 0xaf, 0xbb, 0x5e, 0x57, 0xed, 0x8b, 0xf0, 0x10, 0xbc, 0x1f, 0x77, 0xb4,
	0x7f, 0x4e, 0x0a, 0xa1, 0x52, 0x4f, 0x99, 0xf9, 0xe6, 0xe7, 0x9b, 0x6f, 0x66, 0x63, 0x78, 0x54,
	0xf2, 0x8c, 0x16, 0xcd, 0xcc, 0xfc, 0xc4, 0xb5, 0xe0, 0x92, 0xa3, 0xbe, 0xf1, 0xc6, 0xcf, 0x73,
	0xce, 0xf3, 0x82, 0xce, 0x34, 0xba, 0x6c, 0x2f, 0x66, 0x92, 0x95, 0xb4, 0x91, 0xa4, 0xac, 0x4d,
	0x22, 0x3e, 0x87, 0xe0, 0x6b, 0x43, 0x05, 0x7a, 0x0c, 0x3d, 0x5a, 0x12, 0x56, 0x44, 0xde, 0xc4,
	0x9b, 0x0e, 0x13, 0xe3, 0x20, 0x04, 0x41, 0x45, 0x4a, 0x1a, 0xed, 0x68, 0x50, 0xdb, 0x0a, 0x5b,
	0x91, 0x66, 0x15, 0xf9, 0x06, 0x53, 0x36, 0x1a, 0x43, 0x78, 0x45, 0x05, 0xbb, 0x60, 0x34, 0x8b,
	0x82, 0x89, 0x37, 0x0d, 0x93, 0xce, 0xc7, 0x0b, 0x08, 0x3e, 0x14, 0xed, 0x12, 0x3d, 0x81, 0x7e,
	0x5a, 0xb4, 0xcb, 0x93, 0x23, 0x4b, 0x61, 0xbd, 0xad, 0x1c, 0x13, 0x18, 0x65, 0xb4, 0x49, 0x05,
	0xab, 0x25, 0xe3, 0x95, 0xa5, 0xda, 0x84, 0xf0, 0x39, 0x84, 0xaa, 0xab, 0x9e, 0xfd, 0x7f, 0x9d,
	0x3b, 0x4d, 0x3b, 0xdb, 0x34, 0xf9, 0xb7, 0x35, 0x09, 0x5e, 0x50, 0x3d, 0xfb, 0x30, 0xd1, 0x36,
	0xbe, 0x84, 0xfd, 0x63, 0x55, 0xf0, 0x4d, 0x0b, 0x49, 0x89, 0xa2, 0xed, 0xc4, 0x7b, 0x1b, 0xe2,
	0xb7, 0xd3, 0xcc, 0xa1, 0x4f, 0xaf, 0x6b, 0x26, 0x6e, 0x34, 0xd1, 0x68, 0x3e, 0x8e, 0xcd, 0x29,
	0x62, 0x77, 0x8a, 0x78, 0xe1, 0x4e, 0x91, 0xd8, 0x4c, 0xfc, 0xd3, 0x83, 0xde, 0xf1, 0x15, 0xad,
	0xa4, 0x92, 0x54, 0x53, 0xc1, 0x78, 0xe6, 0x24, 0x19, 0x0f, 0x45, 0x30, 0xa0, 0x2a, 0xe1, 0xe4,
	0xc8, 0xb2, 0x39, 0x77, 0xab, 0xac, 0xbf, 0xd6, 0x18, 0xfc, 0xb3, 0x46, 0xf4, 0x12, 0xfa, 0x17,
	0x8c, 0x16, 0x59, 0x13, 0xf5, 0x26, 0xfe, 0x74, 0x34, 0x7f, 0x18, 0xdb, 0x67, 0xf4, 0x51, 0xa1,
	0x89, 0x0d, 0xe2, 0x5f, 0x1e, 0xf4, 0x34, 0xd2, 0xd1, 0x78, 0x1b, 0x34, 0x63, 0x08, 0x05, 0xbd,
	0x6c, 0x99, 0xa0, 0x99, 0x9e, 0x2a, 0x4c, 0x3a, 0x1f, 0xcd, 0x60, 0x58, 0xf0, 0x2a, 0x3f, 0x93,
	0xf4, 0x5a, 0xda, 0x4d, 0xec, 0x39, 0x8e, 0xcf, 0xbc, 0xca, 0x17, 0xf4, 0x5a, 0x7e, 0x7a, 0x90,
	0x84, 0x85, 0xb5, 0xd1, 0x1c, 0xa0, 0x59, 0x71, 0x21, 0x4d, 0x45, 0xa0, 0x2b, 0xf6, 0x5d, 0xc5,
	0x17, 0x15, 0xb1, 0x25, 0xc3, 0xc6, 0x39, 0x87, 0xbb, 0x00, 0xb5, 0xe0, 0x35, 0x15, 0x92, 0xd1,
	0x06, 0xbf, 0x80, 0xd0, 0x75, 0x46, 0x4f, 0x61, 0xc0, 0xca, 0x33, 0xd5, 0xdc, 0x2d, 0x92, 0x95,
	0x2a, 0x88, 0x5f, 0xc1, 0xb0, 0x6b, 0x86, 0x9e, 0x41, 0xc8, 0xca, 0x33, 0xdd, 0xcf, 0xa6, 0x0d,
	0x58, 0xa9, 0xc3, 0x98, 0xc0, 0xf0, 0x7d, 0x5d, 0x17, 0x2c, 0x25, 0x77, 0x5c, 0xe5, 0x5e, 0x0f,
	0x4d, 0x92, 0xbc, 0x89, 0x82, 0x89, 0xaf, 0x30, 0x65, 0xe3, 0xdf, 0x1e, 0x8c, 0x2c, 0x87, 0xbe,
	0xc9, 0xfd, 0x6f, 0xdf, 0xf1, 0xfb, 0xdb, 0xf8, 0x83, 0x0d, 0xfe, 0x03, 0x18, 0xd0, 0x4a, 0x0a,
	0x46, 0xdd, 0xc1, 0x27, 0x6e, 0xb5, 0x1b, 0x13, 0xc4, 0xc7, 0x26, 0x45, 0xfd, 0xdc, 0x24, 0xae,
	0x60, 0x7c, 0x0a, 0xbb, 0x9b, 0x01, 0xb4, 0x07, 0xfe, 0x0f, 0x7a, 0x63, 0x87, 0x54, 0x26, 0x9a,
	0x42, 0xef, 0x8a, 0x14, 0xad, 0xf9, 0x2f, 0x8f, 0xe6, 0xe8, 0xd6, 0x63, 0x32, 0xdd, 0x4c, 0xc2,
	0xc1, 0xce, 0x3b, 0x0f, 0x63, 0x80, 0x75, 0x40, 0x69, 0x30, 0xb5, 0xaa, 0xdf, 0xae, 0xcd, 0xc3,
	0x6f, 0xc1, 0x5f, 0x90, 0xfc, 0xae, 0x95, 0x48, 0x92, 0x9f, 0xae, 0x3f, 0x1f, 0xce, 0x3d, 0x9c,
	0x7f, 0x7f, 0x9d, 0x33, 0xb9, 0x6a, 0x97, 0x71, 0xca, 0xcb, 0x59, 0xbb, 0x4c, 0x0b, 0xd2, 0x56,
	0xe9, 0xaa, 0x26, 0xd9, 0xac, 0x66, 0x55, 0xcd, 0x59, 0x25, 0xd7, 0x1f, 0x45, 0x33, 0xe6, 0xb2,
	0xaf, 0x81, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x0b, 0xbe, 0x0a, 0x52, 0x05, 0x00,
	0x00,
}
