/**
*   由于开发机过老，无法支持可编译go的protoc，无法根据proto自动生成go文件，下面的内容是从demo上手工编辑的，语法只是一个示意性的，有可能会编译不通过
*/
package hellomall_userbase

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request struct {
	UserId int64 `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	errno int32 `protobuf:"bytes,1,rep,name=errno" json:"errno,omitempty"`
	errmsg string `protobuf:"bytes,2,rep,name=errmsg" json:"errmsg,omitempty"`
	userBaseInfo *UserBaseInfo `protobuf:"bytes,3,rep,name=user_base_info" json:"user_base_info,omitempty"`
}

type UserBaseInfo struct {
	UserId int64 `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name" json:"user_name,omitempty"`
	Sex int32 `protobuf:"bytes,3,opt,name=sex" json:"sex,omitempty"`
	RealUserName string `protobuf:"bytes,4,opt,name=real_user_name" json:"real_user_name,omitempty"`
	Phone string `protobuf:"bytes,5,opt,name=phone" json:"phone,omitempty"`
	Email string `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }


func (m *Response) GetUesrBaseInfo() *UserBaseInfo {
	if m != nil {
		return m.user_base_info
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "hellomall.userbase.Request")
	proto.RegisterType((*Response)(nil), "hellomall.userbase.Response")
	proto.RegisterType((*UserBaseInfo)(nil), "hellomall.userbase.UserBaseInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option



// Client API for Base service

type UserBaseClient interface {
	GetUserBaseInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Respone, error)
}

type userBaseClient struct {
	c           client.Client
	serviceName string
}

func NewUserBaseClient(serviceName string, c client.Client) UserBaseClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "userbase"
	}
	return &userBaseClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userBaseClient) GetUserBaseInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserBase.GetUserBase", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UseBase service

type UserBaseHandler interface {
	GetUserBaseInfo(context.Context, *Request, *Response) error
}

func RegisterUserBaseHandler(s server.Server, hdlr UserBaseHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserBase{hdlr}, opts...))
}

type UserBase struct {
	UserBaseHandler
}

func (h *UserBase) GetUserBaseInfo(ctx context.Context, in *Request, out *Result) error {
	return h.UserBaseHandler.GetUserBaseInfo(ctx, in, out)
}

func init() { proto.RegisterFile("srv/proto/userbase.proto", fileDescriptor0) }


// fileDescriptor0 是以前例子的，没有换，也不知道该好何手工换
var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0x87, 0x92, 0xa9, 0x15,
	0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0xfa, 0xe9, 0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0xfa, 0xc5,
	0xa9, 0x45, 0x65, 0x40, 0xaa, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x3f, 0x23, 0x35, 0x27, 0x07, 0x4a,
	0xea, 0x81, 0x45, 0x84, 0x44, 0xd2, 0xf3, 0xf5, 0xc0, 0x3a, 0xf5, 0x8a, 0x8b, 0xca, 0xf4, 0xa0,
	0x9a, 0x94, 0x64, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x19, 0x2e,
	0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e, 0xe6, 0xdc, 0xe2, 0x74,
	0xa8, 0x34, 0x88, 0x69, 0xe4, 0xcf, 0xc5, 0x1c, 0x9c, 0x58, 0x29, 0xe4, 0xc1, 0xc5, 0xea, 0x01,
	0xb2, 0x48, 0x48, 0x56, 0x0f, 0x9b, 0x1d, 0x7a, 0x50, 0x0b, 0xa4, 0xe4, 0x70, 0x49, 0x43, 0x2c,
	0x50, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x05, 0xb2,
	0xd3, 0xf4, 0x00, 0x00, 0x00,
}
