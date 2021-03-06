// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/reply-feed/api/api.proto

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		app/service/main/reply-feed/api/api.proto

	It has these top-level messages:
		HotReplyReq
		HotReplyRes
		ReplyReq
		ReplyRes
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HotReplyReq struct {
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	Oid int64 `protobuf:"varint,2,opt,name=oid,proto3" json:"oid,omitempty"`
	Tp  int32 `protobuf:"varint,3,opt,name=tp,proto3" json:"tp,omitempty"`
	Pn  int32 `protobuf:"varint,4,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps  int32 `protobuf:"varint,5,opt,name=ps,proto3" json:"ps,omitempty"`
}

func (m *HotReplyReq) Reset()                    { *m = HotReplyReq{} }
func (m *HotReplyReq) String() string            { return proto.CompactTextString(m) }
func (*HotReplyReq) ProtoMessage()               {}
func (*HotReplyReq) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

func (m *HotReplyReq) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *HotReplyReq) GetOid() int64 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *HotReplyReq) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *HotReplyReq) GetPn() int32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *HotReplyReq) GetPs() int32 {
	if m != nil {
		return m.Ps
	}
	return 0
}

type HotReplyRes struct {
	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RpIDs []int64 `protobuf:"varint,2,rep,packed,name=rpIDs" json:"rpIDs,omitempty"`
	Count int32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *HotReplyRes) Reset()                    { *m = HotReplyRes{} }
func (m *HotReplyRes) String() string            { return proto.CompactTextString(m) }
func (*HotReplyRes) ProtoMessage()               {}
func (*HotReplyRes) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *HotReplyRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HotReplyRes) GetRpIDs() []int64 {
	if m != nil {
		return m.RpIDs
	}
	return nil
}

func (m *HotReplyRes) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyReq struct {
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	Pn  int32 `protobuf:"varint,2,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps  int32 `protobuf:"varint,3,opt,name=ps,proto3" json:"ps,omitempty"`
}

func (m *ReplyReq) Reset()                    { *m = ReplyReq{} }
func (m *ReplyReq) String() string            { return proto.CompactTextString(m) }
func (*ReplyReq) ProtoMessage()               {}
func (*ReplyReq) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

func (m *ReplyReq) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *ReplyReq) GetPn() int32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *ReplyReq) GetPs() int32 {
	if m != nil {
		return m.Ps
	}
	return 0
}

type ReplyRes struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ReplyRes) Reset()                    { *m = ReplyRes{} }
func (m *ReplyRes) String() string            { return proto.CompactTextString(m) }
func (*ReplyRes) ProtoMessage()               {}
func (*ReplyRes) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{3} }

func (m *ReplyRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*HotReplyReq)(nil), "community.reply.feed.v1.HotReplyReq")
	proto.RegisterType((*HotReplyRes)(nil), "community.reply.feed.v1.HotReplyRes")
	proto.RegisterType((*ReplyReq)(nil), "community.reply.feed.v1.ReplyReq")
	proto.RegisterType((*ReplyRes)(nil), "community.reply.feed.v1.ReplyRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReplyFeed service

type ReplyFeedClient interface {
	HotReply(ctx context.Context, in *HotReplyReq, opts ...grpc.CallOption) (*HotReplyRes, error)
	Reply(ctx context.Context, in *ReplyReq, opts ...grpc.CallOption) (*ReplyRes, error)
}

type replyFeedClient struct {
	cc *grpc.ClientConn
}

func NewReplyFeedClient(cc *grpc.ClientConn) ReplyFeedClient {
	return &replyFeedClient{cc}
}

func (c *replyFeedClient) HotReply(ctx context.Context, in *HotReplyReq, opts ...grpc.CallOption) (*HotReplyRes, error) {
	out := new(HotReplyRes)
	err := grpc.Invoke(ctx, "/community.reply.feed.v1.ReplyFeed/HotReply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replyFeedClient) Reply(ctx context.Context, in *ReplyReq, opts ...grpc.CallOption) (*ReplyRes, error) {
	out := new(ReplyRes)
	err := grpc.Invoke(ctx, "/community.reply.feed.v1.ReplyFeed/Reply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReplyFeed service

type ReplyFeedServer interface {
	HotReply(context.Context, *HotReplyReq) (*HotReplyRes, error)
	Reply(context.Context, *ReplyReq) (*ReplyRes, error)
}

func RegisterReplyFeedServer(s *grpc.Server, srv ReplyFeedServer) {
	s.RegisterService(&_ReplyFeed_serviceDesc, srv)
}

func _ReplyFeed_HotReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplyFeedServer).HotReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.reply.feed.v1.ReplyFeed/HotReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplyFeedServer).HotReply(ctx, req.(*HotReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplyFeed_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplyFeedServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.reply.feed.v1.ReplyFeed/Reply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplyFeedServer).Reply(ctx, req.(*ReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReplyFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "community.reply.feed.v1.ReplyFeed",
	HandlerType: (*ReplyFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HotReply",
			Handler:    _ReplyFeed_HotReply_Handler,
		},
		{
			MethodName: "Reply",
			Handler:    _ReplyFeed_Reply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/main/reply-feed/api/api.proto",
}

func (m *HotReplyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotReplyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
	}
	if m.Oid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Oid))
	}
	if m.Tp != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Tp))
	}
	if m.Pn != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Pn))
	}
	if m.Ps != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Ps))
	}
	return i, nil
}

func (m *HotReplyRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HotReplyRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.RpIDs) > 0 {
		dAtA2 := make([]byte, len(m.RpIDs)*10)
		var j1 int
		for _, num1 := range m.RpIDs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Count != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *ReplyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
	}
	if m.Pn != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Pn))
	}
	if m.Ps != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Ps))
	}
	return i, nil
}

func (m *ReplyRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HotReplyReq) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	if m.Oid != 0 {
		n += 1 + sovApi(uint64(m.Oid))
	}
	if m.Tp != 0 {
		n += 1 + sovApi(uint64(m.Tp))
	}
	if m.Pn != 0 {
		n += 1 + sovApi(uint64(m.Pn))
	}
	if m.Ps != 0 {
		n += 1 + sovApi(uint64(m.Ps))
	}
	return n
}

func (m *HotReplyRes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.RpIDs) > 0 {
		l = 0
		for _, e := range m.RpIDs {
			l += sovApi(uint64(e))
		}
		n += 1 + sovApi(uint64(l)) + l
	}
	if m.Count != 0 {
		n += 1 + sovApi(uint64(m.Count))
	}
	return n
}

func (m *ReplyReq) Size() (n int) {
	var l int
	_ = l
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	if m.Pn != 0 {
		n += 1 + sovApi(uint64(m.Pn))
	}
	if m.Ps != 0 {
		n += 1 + sovApi(uint64(m.Ps))
	}
	return n
}

func (m *ReplyRes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HotReplyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: HotReplyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotReplyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oid", wireType)
			}
			m.Oid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Oid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pn", wireType)
			}
			m.Pn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pn |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ps", wireType)
			}
			m.Ps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ps |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *HotReplyRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: HotReplyRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HotReplyRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RpIDs = append(m.RpIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthApi
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowApi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RpIDs = append(m.RpIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RpIDs", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *ReplyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ReplyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pn", wireType)
			}
			m.Pn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pn |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ps", wireType)
			}
			m.Ps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ps |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *ReplyRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ReplyRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplyRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
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
				next, err := skipApi(dAtA[start:])
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
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("app/service/main/reply-feed/api/api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4a, 0xc4, 0x30,
	0x1c, 0xc6, 0x49, 0x72, 0x95, 0xbb, 0x08, 0x72, 0x04, 0xc1, 0xe0, 0x50, 0xce, 0xe2, 0x50, 0x07,
	0x53, 0x4e, 0x57, 0xa7, 0x43, 0x44, 0x07, 0x07, 0x33, 0x38, 0xb8, 0xd5, 0x36, 0x42, 0xc0, 0x26,
	0xb1, 0xc9, 0x15, 0xee, 0xb5, 0x7c, 0x0a, 0x47, 0x1f, 0x41, 0xfa, 0x24, 0x92, 0xe4, 0x2c, 0x87,
	0x78, 0xea, 0x50, 0xf8, 0x7f, 0x1f, 0x5f, 0xff, 0xdf, 0x8f, 0x7f, 0xf0, 0x49, 0x69, 0x4c, 0x61,
	0x45, 0xdb, 0xc9, 0x4a, 0x14, 0x4d, 0x29, 0x55, 0xd1, 0x0a, 0xf3, 0xbc, 0x3a, 0x7d, 0x12, 0xa2,
	0x2e, 0x4a, 0x23, 0xfd, 0xc7, 0x4c, 0xab, 0x9d, 0x26, 0x07, 0x95, 0x6e, 0x9a, 0xa5, 0x92, 0x6e,
	0xc5, 0x42, 0x86, 0xf9, 0x0c, 0xeb, 0xe6, 0x59, 0x89, 0x77, 0xaf, 0xb5, 0xe3, 0xde, 0xe3, 0xe2,
	0x85, 0x4c, 0x31, 0x6a, 0x64, 0x4d, 0xc1, 0x0c, 0xe4, 0x88, 0xfb, 0xd1, 0x3b, 0x5a, 0xd6, 0x14,
	0x46, 0x47, 0xcb, 0x9a, 0xec, 0x61, 0xe8, 0x0c, 0x45, 0x33, 0x90, 0x27, 0x1c, 0x3a, 0xe3, 0xb5,
	0x51, 0x74, 0x14, 0xb5, 0x51, 0x41, 0x5b, 0x9a, 0xac, 0xb5, 0xcd, 0xee, 0x36, 0x2b, 0x2c, 0x21,
	0x78, 0xa4, 0xca, 0x46, 0x84, 0x8e, 0x09, 0x0f, 0x33, 0xa1, 0x38, 0x69, 0xcd, 0xcd, 0xa5, 0xa5,
	0x70, 0x86, 0x72, 0xb4, 0x80, 0x53, 0xc0, 0xa3, 0x41, 0xf6, 0x71, 0x52, 0xe9, 0xa5, 0x72, 0xeb,
	0xbe, 0x28, 0xb2, 0x0b, 0x3c, 0xfe, 0x05, 0x39, 0x02, 0xc1, 0x6f, 0x40, 0x68, 0x00, 0x4a, 0x87,
	0xbf, 0x7f, 0xa4, 0x39, 0x7b, 0x05, 0x78, 0x12, 0x02, 0x57, 0x42, 0xd4, 0xe4, 0x1e, 0x8f, 0xbf,
	0xf0, 0xc9, 0x31, 0xdb, 0x72, 0x47, 0xb6, 0x71, 0xc4, 0xc3, 0xff, 0xa4, 0x2c, 0xb9, 0xc5, 0x49,
	0x5c, 0x7a, 0xb4, 0x35, 0x3e, 0x6c, 0xfc, 0x33, 0x62, 0x17, 0xd3, 0xb7, 0x3e, 0x05, 0xef, 0x7d,
	0x0a, 0x3e, 0xfa, 0x14, 0x3c, 0xc0, 0x6e, 0xfe, 0xb8, 0x13, 0x9e, 0xfe, 0xfc, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x68, 0x87, 0xe8, 0xf5, 0x27, 0x02, 0x00, 0x00,
}
