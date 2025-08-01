// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/blog.proto

package blog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blog struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId      string                 `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blog) Reset() {
	*x = Blog{}
	mi := &file_proto_blog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blog) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Blog) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Blog) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId      string                 `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	mi := &file_proto_blog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateBlogRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blog          *Blog                  `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	mi := &file_proto_blog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type GetBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBlogRequest) Reset() {
	*x = GetBlogRequest{}
	mi := &file_proto_blog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogRequest.ProtoReflect.Descriptor instead.
func (*GetBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blog          *Blog                  `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBlogResponse) Reset() {
	*x = GetBlogResponse{}
	mi := &file_proto_blog_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogResponse) ProtoMessage() {}

func (x *GetBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogResponse.ProtoReflect.Descriptor instead.
func (*GetBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type ListBlogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlogsRequest) Reset() {
	*x = ListBlogsRequest{}
	mi := &file_proto_blog_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsRequest) ProtoMessage() {}

func (x *ListBlogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsRequest.ProtoReflect.Descriptor instead.
func (*ListBlogsRequest) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{5}
}

func (x *ListBlogsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListBlogsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListBlogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blogs         []*Blog                `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlogsResponse) Reset() {
	*x = ListBlogsResponse{}
	mi := &file_proto_blog_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogsResponse) ProtoMessage() {}

func (x *ListBlogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogsResponse.ProtoReflect.Descriptor instead.
func (*ListBlogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{6}
}

func (x *ListBlogsResponse) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

func (x *ListBlogsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListBlogsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListBlogsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type UpdateBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBlogRequest) Reset() {
	*x = UpdateBlogRequest{}
	mi := &file_proto_blog_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogRequest) ProtoMessage() {}

func (x *UpdateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBlogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blog          *Blog                  `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBlogResponse) Reset() {
	*x = UpdateBlogResponse{}
	mi := &file_proto_blog_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogResponse) ProtoMessage() {}

func (x *UpdateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBlogRequest) Reset() {
	*x = DeleteBlogRequest{}
	mi := &file_proto_blog_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogRequest) ProtoMessage() {}

func (x *DeleteBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteBlogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBlogResponse) Reset() {
	*x = DeleteBlogResponse{}
	mi := &file_proto_blog_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogResponse) ProtoMessage() {}

func (x *DeleteBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blog_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogResponse.ProtoReflect.Descriptor instead.
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_blog_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteBlogResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_blog_proto protoreflect.FileDescriptor

const file_proto_blog_proto_rawDesc = "" +
	"\n" +
	"\x10proto/blog.proto\x12\x04blog\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd9\x01\n" +
	"\x04Blog\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x1b\n" +
	"\tauthor_id\x18\x04 \x01(\tR\bauthorId\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"`\n" +
	"\x11CreateBlogRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1b\n" +
	"\tauthor_id\x18\x03 \x01(\tR\bauthorId\"4\n" +
	"\x12CreateBlogResponse\x12\x1e\n" +
	"\x04blog\x18\x01 \x01(\v2\n" +
	".blog.BlogR\x04blog\" \n" +
	"\x0eGetBlogRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"1\n" +
	"\x0fGetBlogResponse\x12\x1e\n" +
	"\x04blog\x18\x01 \x01(\v2\n" +
	".blog.BlogR\x04blog\"<\n" +
	"\x10ListBlogsRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\"u\n" +
	"\x11ListBlogsResponse\x12 \n" +
	"\x05blogs\x18\x01 \x03(\v2\n" +
	".blog.BlogR\x05blogs\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\x12\x12\n" +
	"\x04page\x18\x03 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x04 \x01(\x05R\x05limit\"S\n" +
	"\x11UpdateBlogRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\"4\n" +
	"\x12UpdateBlogResponse\x12\x1e\n" +
	"\x04blog\x18\x01 \x01(\v2\n" +
	".blog.BlogR\x04blog\"#\n" +
	"\x11DeleteBlogRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\".\n" +
	"\x12DeleteBlogResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xc6\x02\n" +
	"\vBlogService\x12?\n" +
	"\n" +
	"CreateBlog\x12\x17.blog.CreateBlogRequest\x1a\x18.blog.CreateBlogResponse\x126\n" +
	"\aGetBlog\x12\x14.blog.GetBlogRequest\x1a\x15.blog.GetBlogResponse\x12<\n" +
	"\tListBlogs\x12\x16.blog.ListBlogsRequest\x1a\x17.blog.ListBlogsResponse\x12?\n" +
	"\n" +
	"UpdateBlog\x12\x17.blog.UpdateBlogRequest\x1a\x18.blog.UpdateBlogResponse\x12?\n" +
	"\n" +
	"DeleteBlog\x12\x17.blog.DeleteBlogRequest\x1a\x18.blog.DeleteBlogResponseB2Z0example.com/internal-service/internal/proto/blogb\x06proto3"

var (
	file_proto_blog_proto_rawDescOnce sync.Once
	file_proto_blog_proto_rawDescData []byte
)

func file_proto_blog_proto_rawDescGZIP() []byte {
	file_proto_blog_proto_rawDescOnce.Do(func() {
		file_proto_blog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_blog_proto_rawDesc), len(file_proto_blog_proto_rawDesc)))
	})
	return file_proto_blog_proto_rawDescData
}

var file_proto_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_blog_proto_goTypes = []any{
	(*Blog)(nil),                  // 0: blog.Blog
	(*CreateBlogRequest)(nil),     // 1: blog.CreateBlogRequest
	(*CreateBlogResponse)(nil),    // 2: blog.CreateBlogResponse
	(*GetBlogRequest)(nil),        // 3: blog.GetBlogRequest
	(*GetBlogResponse)(nil),       // 4: blog.GetBlogResponse
	(*ListBlogsRequest)(nil),      // 5: blog.ListBlogsRequest
	(*ListBlogsResponse)(nil),     // 6: blog.ListBlogsResponse
	(*UpdateBlogRequest)(nil),     // 7: blog.UpdateBlogRequest
	(*UpdateBlogResponse)(nil),    // 8: blog.UpdateBlogResponse
	(*DeleteBlogRequest)(nil),     // 9: blog.DeleteBlogRequest
	(*DeleteBlogResponse)(nil),    // 10: blog.DeleteBlogResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_proto_blog_proto_depIdxs = []int32{
	11, // 0: blog.Blog.created_at:type_name -> google.protobuf.Timestamp
	11, // 1: blog.Blog.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: blog.CreateBlogResponse.blog:type_name -> blog.Blog
	0,  // 3: blog.GetBlogResponse.blog:type_name -> blog.Blog
	0,  // 4: blog.ListBlogsResponse.blogs:type_name -> blog.Blog
	0,  // 5: blog.UpdateBlogResponse.blog:type_name -> blog.Blog
	1,  // 6: blog.BlogService.CreateBlog:input_type -> blog.CreateBlogRequest
	3,  // 7: blog.BlogService.GetBlog:input_type -> blog.GetBlogRequest
	5,  // 8: blog.BlogService.ListBlogs:input_type -> blog.ListBlogsRequest
	7,  // 9: blog.BlogService.UpdateBlog:input_type -> blog.UpdateBlogRequest
	9,  // 10: blog.BlogService.DeleteBlog:input_type -> blog.DeleteBlogRequest
	2,  // 11: blog.BlogService.CreateBlog:output_type -> blog.CreateBlogResponse
	4,  // 12: blog.BlogService.GetBlog:output_type -> blog.GetBlogResponse
	6,  // 13: blog.BlogService.ListBlogs:output_type -> blog.ListBlogsResponse
	8,  // 14: blog.BlogService.UpdateBlog:output_type -> blog.UpdateBlogResponse
	10, // 15: blog.BlogService.DeleteBlog:output_type -> blog.DeleteBlogResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_blog_proto_init() }
func file_proto_blog_proto_init() {
	if File_proto_blog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_blog_proto_rawDesc), len(file_proto_blog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blog_proto_goTypes,
		DependencyIndexes: file_proto_blog_proto_depIdxs,
		MessageInfos:      file_proto_blog_proto_msgTypes,
	}.Build()
	File_proto_blog_proto = out.File
	file_proto_blog_proto_goTypes = nil
	file_proto_blog_proto_depIdxs = nil
}
