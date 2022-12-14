// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course_category.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Blank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{0}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

type Category struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateCategoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{2}
}

func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(m, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCategoryRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CategoryResponse struct {
	Category             *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CategoryResponse) Reset()         { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()    {}
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{3}
}

func (m *CategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryResponse.Unmarshal(m, b)
}
func (m *CategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryResponse.Marshal(b, m, deterministic)
}
func (m *CategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryResponse.Merge(m, src)
}
func (m *CategoryResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryResponse.Size(m)
}
func (m *CategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryResponse proto.InternalMessageInfo

func (m *CategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type CategoryList struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CategoryList) Reset()         { *m = CategoryList{} }
func (m *CategoryList) String() string { return proto.CompactTextString(m) }
func (*CategoryList) ProtoMessage()    {}
func (*CategoryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{4}
}

func (m *CategoryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryList.Unmarshal(m, b)
}
func (m *CategoryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryList.Marshal(b, m, deterministic)
}
func (m *CategoryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryList.Merge(m, src)
}
func (m *CategoryList) XXX_Size() int {
	return xxx_messageInfo_CategoryList.Size(m)
}
func (m *CategoryList) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryList.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryList proto.InternalMessageInfo

func (m *CategoryList) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CategoryGetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryGetRequest) Reset()         { *m = CategoryGetRequest{} }
func (m *CategoryGetRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryGetRequest) ProtoMessage()    {}
func (*CategoryGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{5}
}

func (m *CategoryGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryGetRequest.Unmarshal(m, b)
}
func (m *CategoryGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryGetRequest.Marshal(b, m, deterministic)
}
func (m *CategoryGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryGetRequest.Merge(m, src)
}
func (m *CategoryGetRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryGetRequest.Size(m)
}
func (m *CategoryGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryGetRequest proto.InternalMessageInfo

func (m *CategoryGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Blank)(nil), "pb.blank")
	proto.RegisterType((*Category)(nil), "pb.Category")
	proto.RegisterType((*CreateCategoryRequest)(nil), "pb.CreateCategoryRequest")
	proto.RegisterType((*CategoryResponse)(nil), "pb.CategoryResponse")
	proto.RegisterType((*CategoryList)(nil), "pb.CategoryList")
	proto.RegisterType((*CategoryGetRequest)(nil), "pb.CategoryGetRequest")
}

func init() { proto.RegisterFile("proto/course_category.proto", fileDescriptor_fde79cb73866aac8) }

var fileDescriptor_fde79cb73866aac8 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0x49, 0xbf, 0x6a, 0x3b, 0xa9, 0xb5, 0x0c, 0x2a, 0x51, 0x2f, 0x21, 0x78, 0xe8,
	0x41, 0x5a, 0xa9, 0x78, 0x10, 0x7a, 0x6a, 0x0f, 0xbd, 0x28, 0x68, 0x7b, 0xf3, 0x22, 0x9b, 0x64,
	0x90, 0xc5, 0x76, 0x37, 0xee, 0x6e, 0x05, 0xff, 0x4b, 0xff, 0x24, 0xc9, 0xd2, 0x4d, 0xd3, 0x1f,
	0x8a, 0xb7, 0xf0, 0xde, 0xcb, 0x67, 0x67, 0x1e, 0x03, 0x17, 0xb9, 0x92, 0x46, 0xf6, 0x52, 0xb9,
	0x50, 0x9a, 0x5e, 0x52, 0x66, 0xe8, 0x55, 0xaa, 0xcf, 0xae, 0x55, 0xd1, 0xcf, 0x93, 0xf8, 0x00,
	0xf6, 0x92, 0x19, 0x13, 0x6f, 0xf1, 0x23, 0xd4, 0x47, 0x4b, 0x1b, 0x5b, 0xe0, 0xf3, 0x2c, 0xf4,
	0x22, 0xaf, 0xd3, 0x98, 0xf8, 0x3c, 0x43, 0x84, 0xff, 0x82, 0xcd, 0x29, 0xf4, 0xad, 0x62, 0xbf,
	0x31, 0x82, 0x20, 0x23, 0x9d, 0x2a, 0x9e, 0x1b, 0x2e, 0x45, 0x58, 0xb3, 0x56, 0x55, 0x8a, 0x1f,
	0xe0, 0x64, 0xa4, 0x88, 0x19, 0x72, 0xdc, 0x09, 0xbd, 0x2f, 0x48, 0x9b, 0x12, 0xe7, 0xfd, 0x8c,
	0xf3, 0xb7, 0x71, 0x03, 0x68, 0xaf, 0x40, 0x3a, 0x97, 0x42, 0x13, 0x76, 0xa0, 0xee, 0x76, 0xb2,
	0xb4, 0xa0, 0xdf, 0xec, 0xe6, 0x49, 0xb7, 0xcc, 0x95, 0x6e, 0x3c, 0x80, 0xa6, 0x53, 0xef, 0xb9,
	0x36, 0x78, 0x05, 0xb0, 0xf4, 0x38, 0xe9, 0xd0, 0x8b, 0x6a, 0x5b, 0xff, 0x56, 0xfc, 0xf8, 0x12,
	0xd0, 0xe9, 0x63, 0x32, 0x6e, 0x8f, 0x8d, 0x9a, 0xfa, 0x5f, 0x3e, 0x1c, 0xb9, 0xd8, 0x94, 0xd4,
	0x07, 0x4f, 0x09, 0xef, 0xa0, 0xb5, 0x5e, 0x02, 0x9e, 0xd9, 0x57, 0x76, 0x15, 0x73, 0xbe, 0x36,
	0x00, 0x8e, 0xe1, 0x78, 0x3d, 0x36, 0x35, 0x8a, 0xd8, 0xfc, 0x37, 0x40, 0xbb, 0x0a, 0x28, 0xf6,
	0x8c, 0xff, 0x75, 0x3c, 0x7c, 0x82, 0x68, 0x17, 0x68, 0xc8, 0x33, 0xae, 0x28, 0xe5, 0x52, 0xb0,
	0xd9, 0x9f, 0xa7, 0x2a, 0x80, 0xd7, 0x1e, 0xf6, 0xa0, 0x55, 0xe0, 0x47, 0x65, 0x45, 0xd8, 0x28,
	0x52, 0xf6, 0x94, 0x76, 0x4d, 0x81, 0xb7, 0x10, 0x8c, 0xc9, 0x94, 0xbb, 0x9d, 0x56, 0x23, 0xab,
	0x4a, 0x37, 0xdf, 0x1a, 0x1e, 0x3e, 0x07, 0x5c, 0x18, 0x52, 0x82, 0xcd, 0x7a, 0x79, 0x92, 0xec,
	0xdb, 0xc3, 0xbd, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x27, 0xf7, 0x85, 0xd7, 0x02, 0x00,
	0x00,
}
