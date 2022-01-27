// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: robots.proto

package robots

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	groups "github.com/slavayssiere-spoon/groups"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Robot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"-"
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"-"`
	// @inject_tag: gorm:"unique;not null; index:email"
	Email   string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" gorm:"unique;not null; index:email"`
	Gps     string `protobuf:"bytes,3,opt,name=gps,proto3" json:"gps,omitempty"`
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Mac     string `protobuf:"bytes,5,opt,name=mac,proto3" json:"mac,omitempty"`
	// @inject_tag: gorm:"-"
	Groups []*groups.Group `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty" gorm:"-"`
	// @inject_tag: gorm:"-"
	Paths            []string `protobuf:"bytes,7,rep,name=paths,proto3" json:"paths,omitempty" gorm:"-"`
	Name             string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Key              string   `protobuf:"bytes,9,opt,name=key,proto3" json:"key,omitempty"`
	DirectusUser     string   `protobuf:"bytes,10,opt,name=directusUser,proto3" json:"directusUser,omitempty"`
	DirectusPassword string   `protobuf:"bytes,11,opt,name=directusPassword,proto3" json:"directusPassword,omitempty"`
	PubsubTopic      string   `protobuf:"bytes,12,opt,name=pubsubTopic,proto3" json:"pubsubTopic,omitempty"`
}

func (x *Robot) Reset() {
	*x = Robot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robots_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Robot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Robot) ProtoMessage() {}

func (x *Robot) ProtoReflect() protoreflect.Message {
	mi := &file_robots_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Robot.ProtoReflect.Descriptor instead.
func (*Robot) Descriptor() ([]byte, []int) {
	return file_robots_proto_rawDescGZIP(), []int{0}
}

func (x *Robot) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Robot) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Robot) GetGps() string {
	if x != nil {
		return x.Gps
	}
	return ""
}

func (x *Robot) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Robot) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Robot) GetGroups() []*groups.Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Robot) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Robot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Robot) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Robot) GetDirectusUser() string {
	if x != nil {
		return x.DirectusUser
	}
	return ""
}

func (x *Robot) GetDirectusPassword() string {
	if x != nil {
		return x.DirectusPassword
	}
	return ""
}

func (x *Robot) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

type Robots struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List   []*Robot `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Limit  uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Max    uint64   `protobuf:"varint,4,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Robots) Reset() {
	*x = Robots{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robots_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Robots) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Robots) ProtoMessage() {}

func (x *Robots) ProtoReflect() protoreflect.Message {
	mi := &file_robots_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Robots.ProtoReflect.Descriptor instead.
func (*Robots) Descriptor() ([]byte, []int) {
	return file_robots_proto_rawDescGZIP(), []int{1}
}

func (x *Robots) GetList() []*Robot {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Robots) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Robots) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Robots) GetMax() uint64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type SaFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SaFile) Reset() {
	*x = SaFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robots_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaFile) ProtoMessage() {}

func (x *SaFile) ProtoReflect() protoreflect.Message {
	mi := &file_robots_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaFile.ProtoReflect.Descriptor instead.
func (*SaFile) Descriptor() ([]byte, []int) {
	return file_robots_proto_rawDescGZIP(), []int{2}
}

func (x *SaFile) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DirectusToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	// @inject_tag: json:"expires"
	Expires int64 `protobuf:"varint,2,opt,name=expires,proto3" json:"expires"`
	// @inject_tag: json:"refresh_token"
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
}

func (x *DirectusToken) Reset() {
	*x = DirectusToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_robots_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectusToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectusToken) ProtoMessage() {}

func (x *DirectusToken) ProtoReflect() protoreflect.Message {
	mi := &file_robots_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectusToken.ProtoReflect.Descriptor instead.
func (*DirectusToken) Descriptor() ([]byte, []int) {
	return file_robots_proto_rawDescGZIP(), []int{3}
}

func (x *DirectusToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *DirectusToken) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

func (x *DirectusToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_robots_proto protoreflect.FileDescriptor

var file_robots_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc0, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x70,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x25, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x22, 0x6b, 0x0a, 0x06, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22,
	0x22, 0x0a, 0x06, 0x53, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x71, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa5, 0x05, 0x0a, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x73, 0x12, 0x3c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x1a, 0x0e, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x12,
	0x4c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x3f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62,
	0x6f, 0x74, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5f,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x75, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f,
	0x74, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x75, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x75, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x4a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0d, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0e, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x53, 0x41, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x1a, 0x0e, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73,
	0x2e, 0x53, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x73, 0x61, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x41,
	0x64, 0x64, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f,
	0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74,
	0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x32, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x73, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x1a, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x61, 0x63, 0x2f, 0x7b, 0x6d, 0x61, 0x63, 0x7d, 0x42, 0xb9,
	0x04, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c,
	0x61, 0x76, 0x61, 0x79, 0x73, 0x73, 0x69, 0x65, 0x72, 0x65, 0x2d, 0x73, 0x70, 0x6f, 0x6f, 0x6e,
	0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x92, 0x41, 0x8f, 0x04, 0x12, 0x81, 0x01, 0x0a, 0x11,
	0x53, 0x70, 0x6f, 0x6f, 0x6e, 0x20, 0x2d, 0x20, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x20, 0x57,
	0x53, 0x22, 0x65, 0x0a, 0x10, 0x53, 0x70, 0x6f, 0x6f, 0x6e, 0x20, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x20, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x31, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x6f, 0x6f, 0x6e, 0x51,
	0x49, 0x52, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x1a, 0x1e, 0x73, 0x65, 0x62, 0x61, 0x73, 0x74,
	0x69, 0x65, 0x6e, 0x2e, 0x6c, 0x61, 0x76, 0x61, 0x79, 0x73, 0x73, 0x69, 0x65, 0x72, 0x65, 0x40,
	0x73, 0x70, 0x6f, 0x6f, 0x6e, 0x2e, 0x61, 0x69, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x2a,
	0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49,
	0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64,
	0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06,
	0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x57, 0x0a, 0x03, 0x34, 0x31, 0x38, 0x12, 0x50, 0x0a,
	0x0d, 0x49, 0x27, 0x6d, 0x20, 0x61, 0x20, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x2e, 0x12, 0x3f,
	0x0a, 0x3d, 0x1a, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x75, 0x6d, 0x5a,
	0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x72, 0x42, 0x0a, 0x0d, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x12, 0x31, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x70, 0x6f, 0x6f,
	0x6e, 0x51, 0x49, 0x52, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_robots_proto_rawDescOnce sync.Once
	file_robots_proto_rawDescData = file_robots_proto_rawDesc
)

func file_robots_proto_rawDescGZIP() []byte {
	file_robots_proto_rawDescOnce.Do(func() {
		file_robots_proto_rawDescData = protoimpl.X.CompressGZIP(file_robots_proto_rawDescData)
	})
	return file_robots_proto_rawDescData
}

var file_robots_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_robots_proto_goTypes = []interface{}{
	(*Robot)(nil),         // 0: robots.Robot
	(*Robots)(nil),        // 1: robots.Robots
	(*SaFile)(nil),        // 2: robots.SaFile
	(*DirectusToken)(nil), // 3: robots.DirectusToken
	(*groups.Group)(nil),  // 4: groups.Group
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
	(*groups.Groups)(nil), // 6: groups.Groups
}
var file_robots_proto_depIdxs = []int32{
	4,  // 0: robots.Robot.groups:type_name -> groups.Group
	0,  // 1: robots.Robots.list:type_name -> robots.Robot
	1,  // 2: robots.robots.GetAll:input_type -> robots.Robots
	5,  // 3: robots.robots.GetGraph:input_type -> google.protobuf.Empty
	0,  // 4: robots.robots.Get:input_type -> robots.Robot
	0,  // 5: robots.robots.GetDirectusToken:input_type -> robots.Robot
	4,  // 6: robots.robots.GetByGroup:input_type -> groups.Group
	0,  // 7: robots.robots.GetSAFile:input_type -> robots.Robot
	0,  // 8: robots.robots.Add:input_type -> robots.Robot
	0,  // 9: robots.robots.Update:input_type -> robots.Robot
	0,  // 10: robots.robots.UpdateMacAddress:input_type -> robots.Robot
	1,  // 11: robots.robots.GetAll:output_type -> robots.Robots
	6,  // 12: robots.robots.GetGraph:output_type -> groups.Groups
	0,  // 13: robots.robots.Get:output_type -> robots.Robot
	3,  // 14: robots.robots.GetDirectusToken:output_type -> robots.DirectusToken
	1,  // 15: robots.robots.GetByGroup:output_type -> robots.Robots
	2,  // 16: robots.robots.GetSAFile:output_type -> robots.SaFile
	0,  // 17: robots.robots.Add:output_type -> robots.Robot
	0,  // 18: robots.robots.Update:output_type -> robots.Robot
	0,  // 19: robots.robots.UpdateMacAddress:output_type -> robots.Robot
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_robots_proto_init() }
func file_robots_proto_init() {
	if File_robots_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_robots_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Robot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_robots_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Robots); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_robots_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_robots_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectusToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_robots_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_robots_proto_goTypes,
		DependencyIndexes: file_robots_proto_depIdxs,
		MessageInfos:      file_robots_proto_msgTypes,
	}.Build()
	File_robots_proto = out.File
	file_robots_proto_rawDesc = nil
	file_robots_proto_goTypes = nil
	file_robots_proto_depIdxs = nil
}
