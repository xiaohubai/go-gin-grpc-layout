// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.26.0
// source: http/v1/http.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 用户名
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 验证码
	Captcha string `protobuf:"bytes,3,opt,name=captcha,proto3" json:"captcha,omitempty"`
	// 校验ID
	CaptchaID     string `protobuf:"bytes,4,opt,name=captchaID,proto3" json:"captchaID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_http_v1_http_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1_http_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_http_v1_http_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetCaptcha() string {
	if x != nil {
		return x.Captcha
	}
	return ""
}

func (x *LoginRequest) GetCaptchaID() string {
	if x != nil {
		return x.CaptchaID
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UID           string                 `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID,omitempty"`
	UserName      string                 `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	NickName      string                 `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Birth         string                 `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	Avatar        string                 `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	RoleID        string                 `protobuf:"bytes,7,opt,name=roleID,proto3" json:"roleID,omitempty"`
	RoleName      string                 `protobuf:"bytes,8,opt,name=roleName,proto3" json:"roleName,omitempty"`
	Phone         string                 `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Wechat        string                 `protobuf:"bytes,10,opt,name=wechat,proto3" json:"wechat,omitempty"`
	Email         string                 `protobuf:"bytes,11,opt,name=email,proto3" json:"email,omitempty"`
	State         int32                  `protobuf:"varint,12,opt,name=state,proto3" json:"state,omitempty"`
	Motto         string                 `protobuf:"bytes,13,opt,name=motto,proto3" json:"motto,omitempty"`
	Token         string                 `protobuf:"bytes,18,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,19,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_http_v1_http_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_v1_http_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_http_v1_http_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *LoginResponse) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *LoginResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *LoginResponse) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *LoginResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *LoginResponse) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *LoginResponse) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *LoginResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginResponse) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *LoginResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginResponse) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *LoginResponse) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_http_v1_http_proto protoreflect.FileDescriptor

var file_http_v1_http_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x72, 0x04, 0x10, 0x03, 0x18, 0x0b, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x06, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98,
	0x01, 0x06, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49, 0x44, 0x22, 0xf5, 0x02, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x47, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x69, 0x61, 0x6f, 0x68, 0x75, 0x62, 0x61, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x69, 0x6e, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_http_v1_http_proto_rawDescOnce sync.Once
	file_http_v1_http_proto_rawDescData = file_http_v1_http_proto_rawDesc
)

func file_http_v1_http_proto_rawDescGZIP() []byte {
	file_http_v1_http_proto_rawDescOnce.Do(func() {
		file_http_v1_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_v1_http_proto_rawDescData)
	})
	return file_http_v1_http_proto_rawDescData
}

var file_http_v1_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_http_v1_http_proto_goTypes = []any{
	(*LoginRequest)(nil),  // 0: api.http.v1.LoginRequest
	(*LoginResponse)(nil), // 1: api.http.v1.LoginResponse
}
var file_http_v1_http_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_v1_http_proto_init() }
func file_http_v1_http_proto_init() {
	if File_http_v1_http_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_v1_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_v1_http_proto_goTypes,
		DependencyIndexes: file_http_v1_http_proto_depIdxs,
		MessageInfos:      file_http_v1_http_proto_msgTypes,
	}.Build()
	File_http_v1_http_proto = out.File
	file_http_v1_http_proto_rawDesc = nil
	file_http_v1_http_proto_goTypes = nil
	file_http_v1_http_proto_depIdxs = nil
}
