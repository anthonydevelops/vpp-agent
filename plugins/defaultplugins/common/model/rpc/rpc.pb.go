// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	DataRequest
	PutResponse
	DelResponse
	ResyncResponse
	Statistics
	StatisticsResponse
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import acl "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/acl"
import bfd "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/bfd"
import interfaces "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/interfaces"
import l2 "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/l2"
import l3 "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/l3"
import l4 "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/l4"
import nat "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/nat"
import stn "github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/stn"
import interfaces1 "github.com/ligato/vpp-agent/plugins/linuxplugin/common/model/interfaces"
import l31 "github.com/ligato/vpp-agent/plugins/linuxplugin/common/model/l3"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Data request is an inventory of supported data types with one or multiple
// items of every type. Universal type for every data change/resync request
type DataRequest struct {
	// Defaultplugins
	AccessLists           []*acl.AccessLists_Acl                 `protobuf:"bytes,10,rep,name=AccessLists" json:"AccessLists,omitempty"`
	Interfaces            []*interfaces.Interfaces_Interface     `protobuf:"bytes,20,rep,name=Interfaces" json:"Interfaces,omitempty"`
	BfdSessions           []*bfd.SingleHopBFD_Session            `protobuf:"bytes,30,rep,name=BfdSessions" json:"BfdSessions,omitempty"`
	BfdAuthKeys           []*bfd.SingleHopBFD_Key                `protobuf:"bytes,31,rep,name=BfdAuthKeys" json:"BfdAuthKeys,omitempty"`
	BfdEchoFunction       *bfd.SingleHopBFD_EchoFunction         `protobuf:"bytes,32,opt,name=BfdEchoFunction" json:"BfdEchoFunction,omitempty"`
	BridgeDomains         []*l2.BridgeDomains_BridgeDomain       `protobuf:"bytes,40,rep,name=BridgeDomains" json:"BridgeDomains,omitempty"`
	FIBs                  []*l2.FibTable_FibEntry                `protobuf:"bytes,41,rep,name=FIBs" json:"FIBs,omitempty"`
	XCons                 []*l2.XConnectPairs_XConnectPair       `protobuf:"bytes,42,rep,name=XCons" json:"XCons,omitempty"`
	StaticRoutes          []*l3.StaticRoutes_Route               `protobuf:"bytes,50,rep,name=StaticRoutes" json:"StaticRoutes,omitempty"`
	ArpEntries            []*l3.ArpTable_ArpEntry                `protobuf:"bytes,51,rep,name=ArpEntries" json:"ArpEntries,omitempty"`
	ProxyArpInterfaces    []*l3.ProxyArpInterfaces_InterfaceList `protobuf:"bytes,52,rep,name=ProxyArpInterfaces" json:"ProxyArpInterfaces,omitempty"`
	ProxyArpRanges        []*l3.ProxyArpRanges_RangeList         `protobuf:"bytes,53,rep,name=ProxyArpRanges" json:"ProxyArpRanges,omitempty"`
	L4Feature             *l4.L4Features                         `protobuf:"bytes,60,opt,name=L4Feature" json:"L4Feature,omitempty"`
	ApplicationNamespaces []*l4.AppNamespaces_AppNamespace       `protobuf:"bytes,61,rep,name=ApplicationNamespaces" json:"ApplicationNamespaces,omitempty"`
	StnRules              []*stn.STN_Rule                        `protobuf:"bytes,70,rep,name=StnRules" json:"StnRules,omitempty"`
	NatGlobal             *nat.Nat44Global                       `protobuf:"bytes,71,opt,name=NatGlobal" json:"NatGlobal,omitempty"`
	DNATs                 []*nat.Nat44DNat_DNatConfig            `protobuf:"bytes,72,rep,name=DNATs" json:"DNATs,omitempty"`
	// Linuxplugin
	LinuxInterfaces []*interfaces1.LinuxInterfaces_Interface `protobuf:"bytes,80,rep,name=LinuxInterfaces" json:"LinuxInterfaces,omitempty"`
	LinuxArpEntries []*l31.LinuxStaticArpEntries_ArpEntry    `protobuf:"bytes,90,rep,name=LinuxArpEntries" json:"LinuxArpEntries,omitempty"`
	LinuxRoutes     []*l31.LinuxStaticRoutes_Route           `protobuf:"bytes,91,rep,name=LinuxRoutes" json:"LinuxRoutes,omitempty"`
}

func (m *DataRequest) Reset()                    { *m = DataRequest{} }
func (m *DataRequest) String() string            { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()               {}
func (*DataRequest) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

func (m *DataRequest) GetAccessLists() []*acl.AccessLists_Acl {
	if m != nil {
		return m.AccessLists
	}
	return nil
}

func (m *DataRequest) GetInterfaces() []*interfaces.Interfaces_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *DataRequest) GetBfdSessions() []*bfd.SingleHopBFD_Session {
	if m != nil {
		return m.BfdSessions
	}
	return nil
}

func (m *DataRequest) GetBfdAuthKeys() []*bfd.SingleHopBFD_Key {
	if m != nil {
		return m.BfdAuthKeys
	}
	return nil
}

func (m *DataRequest) GetBfdEchoFunction() *bfd.SingleHopBFD_EchoFunction {
	if m != nil {
		return m.BfdEchoFunction
	}
	return nil
}

func (m *DataRequest) GetBridgeDomains() []*l2.BridgeDomains_BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *DataRequest) GetFIBs() []*l2.FibTable_FibEntry {
	if m != nil {
		return m.FIBs
	}
	return nil
}

func (m *DataRequest) GetXCons() []*l2.XConnectPairs_XConnectPair {
	if m != nil {
		return m.XCons
	}
	return nil
}

func (m *DataRequest) GetStaticRoutes() []*l3.StaticRoutes_Route {
	if m != nil {
		return m.StaticRoutes
	}
	return nil
}

func (m *DataRequest) GetArpEntries() []*l3.ArpTable_ArpEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

func (m *DataRequest) GetProxyArpInterfaces() []*l3.ProxyArpInterfaces_InterfaceList {
	if m != nil {
		return m.ProxyArpInterfaces
	}
	return nil
}

func (m *DataRequest) GetProxyArpRanges() []*l3.ProxyArpRanges_RangeList {
	if m != nil {
		return m.ProxyArpRanges
	}
	return nil
}

func (m *DataRequest) GetL4Feature() *l4.L4Features {
	if m != nil {
		return m.L4Feature
	}
	return nil
}

func (m *DataRequest) GetApplicationNamespaces() []*l4.AppNamespaces_AppNamespace {
	if m != nil {
		return m.ApplicationNamespaces
	}
	return nil
}

func (m *DataRequest) GetStnRules() []*stn.STN_Rule {
	if m != nil {
		return m.StnRules
	}
	return nil
}

func (m *DataRequest) GetNatGlobal() *nat.Nat44Global {
	if m != nil {
		return m.NatGlobal
	}
	return nil
}

func (m *DataRequest) GetDNATs() []*nat.Nat44DNat_DNatConfig {
	if m != nil {
		return m.DNATs
	}
	return nil
}

func (m *DataRequest) GetLinuxInterfaces() []*interfaces1.LinuxInterfaces_Interface {
	if m != nil {
		return m.LinuxInterfaces
	}
	return nil
}

func (m *DataRequest) GetLinuxArpEntries() []*l31.LinuxStaticArpEntries_ArpEntry {
	if m != nil {
		return m.LinuxArpEntries
	}
	return nil
}

func (m *DataRequest) GetLinuxRoutes() []*l31.LinuxStaticRoutes_Route {
	if m != nil {
		return m.LinuxRoutes
	}
	return nil
}

// Response to data change 'put'
type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

// Response to data change 'del'
type DelResponse struct {
}

func (m *DelResponse) Reset()                    { *m = DelResponse{} }
func (m *DelResponse) String() string            { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()               {}
func (*DelResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{2} }

// Response to data resync
type ResyncResponse struct {
}

func (m *ResyncResponse) Reset()                    { *m = ResyncResponse{} }
func (m *ResyncResponse) String() string            { return proto.CompactTextString(m) }
func (*ResyncResponse) ProtoMessage()               {}
func (*ResyncResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{3} }

type Statistics struct {
	IfNotif *interfaces.InterfaceNotification `protobuf:"bytes,1,opt,name=ifNotif" json:"ifNotif,omitempty"`
}

func (m *Statistics) Reset()                    { *m = Statistics{} }
func (m *Statistics) String() string            { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()               {}
func (*Statistics) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{4} }

func (m *Statistics) GetIfNotif() *interfaces.InterfaceNotification {
	if m != nil {
		return m.IfNotif
	}
	return nil
}

type StatisticsResponse struct {
}

func (m *StatisticsResponse) Reset()                    { *m = StatisticsResponse{} }
func (m *StatisticsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatisticsResponse) ProtoMessage()               {}
func (*StatisticsResponse) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{5} }

func init() {
	proto.RegisterType((*DataRequest)(nil), "rpc.DataRequest")
	proto.RegisterType((*PutResponse)(nil), "rpc.PutResponse")
	proto.RegisterType((*DelResponse)(nil), "rpc.DelResponse")
	proto.RegisterType((*ResyncResponse)(nil), "rpc.ResyncResponse")
	proto.RegisterType((*Statistics)(nil), "rpc.Statistics")
	proto.RegisterType((*StatisticsResponse)(nil), "rpc.StatisticsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataChangeService service

type DataChangeServiceClient interface {
	// Creates or updates one or multiple configuration items
	Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error)
}

type dataChangeServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataChangeServiceClient(cc *grpc.ClientConn) DataChangeServiceClient {
	return &dataChangeServiceClient{cc}
}

func (c *dataChangeServiceClient) Put(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/rpc.DataChangeService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataChangeServiceClient) Del(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := grpc.Invoke(ctx, "/rpc.DataChangeService/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataChangeService service

type DataChangeServiceServer interface {
	// Creates or updates one or multiple configuration items
	Put(context.Context, *DataRequest) (*PutResponse, error)
	// Removes one or multiple configuration items
	Del(context.Context, *DataRequest) (*DelResponse, error)
}

func RegisterDataChangeServiceServer(s *grpc.Server, srv DataChangeServiceServer) {
	s.RegisterService(&_DataChangeService_serviceDesc, srv)
}

func _DataChangeService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Put(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataChangeService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataChangeServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataChangeService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataChangeServiceServer).Del(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataChangeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataChangeService",
	HandlerType: (*DataChangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _DataChangeService_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _DataChangeService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for DataResyncService service

type DataResyncServiceClient interface {
	// Calls vpp-agent resync
	Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error)
}

type dataResyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataResyncServiceClient(cc *grpc.ClientConn) DataResyncServiceClient {
	return &dataResyncServiceClient{cc}
}

func (c *dataResyncServiceClient) Resync(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*ResyncResponse, error) {
	out := new(ResyncResponse)
	err := grpc.Invoke(ctx, "/rpc.DataResyncService/Resync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataResyncService service

type DataResyncServiceServer interface {
	// Calls vpp-agent resync
	Resync(context.Context, *DataRequest) (*ResyncResponse, error)
}

func RegisterDataResyncServiceServer(s *grpc.Server, srv DataResyncServiceServer) {
	s.RegisterService(&_DataResyncService_serviceDesc, srv)
}

func _DataResyncService_Resync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataResyncServiceServer).Resync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DataResyncService/Resync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataResyncServiceServer).Resync(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataResyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DataResyncService",
	HandlerType: (*DataResyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resync",
			Handler:    _DataResyncService_Resync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for StatisticsService service

type StatisticsServiceClient interface {
	Send(ctx context.Context, in *Statistics, opts ...grpc.CallOption) (*StatisticsResponse, error)
}

type statisticsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatisticsServiceClient(cc *grpc.ClientConn) StatisticsServiceClient {
	return &statisticsServiceClient{cc}
}

func (c *statisticsServiceClient) Send(ctx context.Context, in *Statistics, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	out := new(StatisticsResponse)
	err := grpc.Invoke(ctx, "/rpc.StatisticsService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StatisticsService service

type StatisticsServiceServer interface {
	Send(context.Context, *Statistics) (*StatisticsResponse, error)
}

func RegisterStatisticsServiceServer(s *grpc.Server, srv StatisticsServiceServer) {
	s.RegisterService(&_StatisticsService_serviceDesc, srv)
}

func _StatisticsService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Statistics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StatisticsService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServiceServer).Send(ctx, req.(*Statistics))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatisticsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.StatisticsService",
	HandlerType: (*StatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _StatisticsService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xff, 0x6f, 0xe3, 0x34,
	0x14, 0xd7, 0x69, 0x77, 0x07, 0x73, 0xd9, 0x97, 0x33, 0x3b, 0x08, 0x03, 0x8d, 0x31, 0x81, 0xb4,
	0x21, 0x70, 0x50, 0x9a, 0x03, 0x09, 0x38, 0x89, 0x74, 0x59, 0x6f, 0xd3, 0x46, 0xa9, 0xdc, 0xfe,
	0x70, 0xe2, 0x7e, 0x72, 0x5d, 0xb7, 0xb3, 0xe4, 0xda, 0x21, 0x76, 0x4e, 0xd7, 0xbf, 0x99, 0x7f,
	0x02, 0xd9, 0x49, 0x1a, 0xa7, 0x2b, 0xd2, 0xa4, 0xed, 0x07, 0x2f, 0x79, 0xef, 0x7d, 0x3e, 0x1f,
	0x3f, 0xe7, 0x3d, 0xbf, 0x15, 0x6c, 0xe7, 0x19, 0x45, 0x59, 0xae, 0x8c, 0x82, 0x5b, 0x79, 0x46,
	0x0f, 0xff, 0x9c, 0x73, 0x73, 0x5b, 0x4c, 0x10, 0x55, 0x8b, 0x50, 0xf0, 0x39, 0x31, 0x2a, 0x7c,
	0x9f, 0x65, 0x3f, 0x92, 0x39, 0x93, 0x26, 0xcc, 0x44, 0x31, 0xe7, 0x52, 0x87, 0x53, 0x36, 0x23,
	0x85, 0x30, 0xb5, 0x49, 0xd5, 0x62, 0xa1, 0x64, 0xb8, 0x50, 0x53, 0x26, 0x42, 0x42, 0xdd, 0x2a,
	0x35, 0x1f, 0x2e, 0x37, 0x99, 0x4d, 0xed, 0xaa, 0xe4, 0xde, 0x3d, 0x54, 0x8e, 0x4b, 0xc3, 0xf2,
	0x19, 0xa1, 0x4c, 0x7b, 0xaf, 0x95, 0xf8, 0xf5, 0x43, 0xc5, 0x45, 0x14, 0x8a, 0xe8, 0xd1, 0xc4,
	0xba, 0xa1, 0xe8, 0x3e, 0x9a, 0x58, 0x1c, 0x8a, 0xf8, 0xb1, 0x4a, 0x22, 0x89, 0xb1, 0xeb, 0xb1,
	0xe4, 0xb4, 0x91, 0x76, 0x55, 0x72, 0x6f, 0xef, 0x23, 0x27, 0xb8, 0x2c, 0x3e, 0x94, 0xc6, 0xfd,
	0xcb, 0x7b, 0xf9, 0x20, 0x65, 0xaf, 0x1c, 0x27, 0xff, 0x6e, 0x83, 0x4e, 0x4a, 0x0c, 0xc1, 0xec,
	0x9f, 0x82, 0x69, 0x03, 0x7f, 0x06, 0x9d, 0x84, 0x52, 0xa6, 0xf5, 0x0d, 0xd7, 0x46, 0x07, 0xe0,
	0x78, 0xeb, 0xb4, 0x13, 0x1d, 0x20, 0x7b, 0x0b, 0x3c, 0x3f, 0x4a, 0xa8, 0xc0, 0x3e, 0x10, 0xfe,
	0x01, 0xc0, 0xd5, 0x2a, 0xcb, 0xe0, 0xc0, 0xd1, 0x8e, 0x91, 0x97, 0xf8, 0xd5, 0x86, 0x57, 0xec,
	0x71, 0xe0, 0x6f, 0xa0, 0xd3, 0x9b, 0x4d, 0x47, 0x4c, 0x6b, 0xae, 0xa4, 0x0e, 0x8e, 0x9c, 0xc4,
	0x17, 0xc8, 0x5e, 0x98, 0x11, 0x97, 0x73, 0xc1, 0x2e, 0x55, 0xd6, 0xeb, 0xa7, 0xa8, 0x42, 0x60,
	0x1f, 0x0d, 0x7f, 0x71, 0xe4, 0xa4, 0x30, 0xb7, 0xd7, 0x6c, 0xa9, 0x83, 0xaf, 0x1d, 0xf9, 0xe5,
	0x5d, 0xf2, 0x35, 0x5b, 0x62, 0x1f, 0x09, 0x2f, 0xc1, 0x5e, 0x6f, 0x36, 0xbd, 0xa0, 0xb7, 0xaa,
	0x5f, 0x48, 0x6a, 0xb8, 0x92, 0xc1, 0xf1, 0xf1, 0x93, 0xd3, 0x4e, 0x74, 0x74, 0x97, 0xec, 0xa3,
	0xf0, 0x3a, 0x0d, 0xa6, 0x60, 0xa7, 0x97, 0xf3, 0xe9, 0x9c, 0xa5, 0x6a, 0x41, 0xb8, 0xd4, 0xc1,
	0xa9, 0x4b, 0xe2, 0x08, 0x89, 0x08, 0xb5, 0x02, 0x2d, 0x0b, 0xb7, 0x49, 0xf0, 0x0c, 0x3c, 0xed,
	0x5f, 0xf5, 0x74, 0x70, 0x56, 0x9d, 0x40, 0x44, 0xa8, 0xcf, 0x27, 0x63, 0x32, 0x11, 0xcc, 0xbe,
	0x5c, 0x48, 0x93, 0x2f, 0xb1, 0x83, 0xc0, 0x18, 0x3c, 0x7b, 0x7b, 0x6e, 0x3f, 0xd5, 0xf7, 0xcd,
	0x46, 0xd6, 0x21, 0x19, 0x35, 0x43, 0xc2, 0x73, 0xdd, 0xb2, 0x70, 0x09, 0x86, 0xbf, 0x82, 0x4f,
	0x46, 0x86, 0x18, 0x4e, 0xb1, 0x2a, 0x0c, 0xd3, 0x41, 0xe4, 0xc8, 0x9f, 0x21, 0xd1, 0x45, 0xbe,
	0x1f, 0xb9, 0x07, 0x6e, 0x61, 0xe1, 0x2b, 0x00, 0x92, 0x3c, 0xb3, 0x39, 0x70, 0xa6, 0x83, 0x6e,
	0x9d, 0x62, 0x17, 0x25, 0x79, 0x56, 0xa6, 0x58, 0x85, 0x97, 0xd8, 0x03, 0xc2, 0x31, 0x80, 0xc3,
	0x5c, 0x7d, 0x58, 0x26, 0x79, 0xe6, 0xf5, 0x48, 0xec, 0xe8, 0xdf, 0x5a, 0xfa, 0xdd, 0x68, 0xd3,
	0x23, 0xb6, 0xbd, 0xf0, 0x06, 0x3e, 0x4c, 0xc1, 0x6e, 0xed, 0xc5, 0x44, 0xce, 0x99, 0x0e, 0x5e,
	0x39, 0xc5, 0xaf, 0x7c, 0xc5, 0x32, 0x82, 0xdc, 0xc3, 0x29, 0xad, 0x71, 0xe0, 0x0f, 0x60, 0xfb,
	0x26, 0xee, 0x33, 0x62, 0x8a, 0x9c, 0x05, 0xbf, 0xbb, 0xca, 0xef, 0x22, 0x11, 0xa3, 0x95, 0x53,
	0xe3, 0x06, 0x00, 0xc7, 0xe0, 0x65, 0x92, 0x65, 0x82, 0x53, 0x62, 0x4b, 0x3e, 0x20, 0x0b, 0xa6,
	0x33, 0x77, 0x98, 0xd7, 0x75, 0x09, 0x62, 0x94, 0x64, 0x59, 0x13, 0x68, 0x59, 0x78, 0x33, 0x19,
	0x9e, 0x81, 0x8f, 0x47, 0x46, 0xe2, 0x42, 0x30, 0x1d, 0xf4, 0x9d, 0xd0, 0x0e, 0xb2, 0x53, 0x64,
	0x34, 0x1e, 0x20, 0xeb, 0xc5, 0xab, 0x30, 0x44, 0x60, 0x7b, 0x40, 0xcc, 0x1b, 0xa1, 0x26, 0x44,
	0x04, 0x6f, 0x5c, 0xba, 0xfb, 0xc8, 0x0e, 0xb0, 0x01, 0x31, 0x71, 0x5c, 0xfa, 0x71, 0x03, 0x81,
	0x21, 0x78, 0x96, 0x0e, 0x92, 0xb1, 0x0e, 0x2e, 0xab, 0xeb, 0xb4, 0xc2, 0xa6, 0x03, 0x62, 0x90,
	0xfd, 0x73, 0xae, 0xe4, 0x8c, 0xcf, 0x71, 0x89, 0x83, 0x7f, 0x81, 0xbd, 0x1b, 0x3b, 0x37, 0xbc,
	0x42, 0x0d, 0x1d, 0xf5, 0x3b, 0xff, 0x32, 0xaf, 0x41, 0xbc, 0x1b, 0xbd, 0xce, 0x86, 0x37, 0x95,
	0xa0, 0xd7, 0x38, 0x7f, 0x3b, 0xc1, 0x13, 0x5b, 0x27, 0x17, 0x2a, 0x7b, 0xac, 0x01, 0x34, 0x5d,
	0xb4, 0x4e, 0x85, 0xaf, 0x41, 0xc7, 0xb9, 0xaa, 0xe6, 0x7d, 0xe7, 0x94, 0xbe, 0x5c, 0x53, 0x6a,
	0x75, 0xb0, 0x8f, 0x3f, 0xd9, 0x01, 0x9d, 0x61, 0x61, 0x30, 0xd3, 0x99, 0x92, 0x9a, 0x59, 0x33,
	0x65, 0x62, 0x65, 0xee, 0x83, 0x5d, 0xcc, 0xf4, 0x52, 0xd2, 0x95, 0xe7, 0x0a, 0x00, 0x27, 0xa9,
	0x0d, 0xa7, 0x76, 0x42, 0x7d, 0xc4, 0x67, 0x03, 0x65, 0xf8, 0x2c, 0x78, 0xe2, 0x3e, 0xfd, 0x37,
	0x1b, 0x07, 0x9c, 0x43, 0x54, 0x55, 0xc6, 0x35, 0xe3, 0xe4, 0x00, 0xc0, 0x46, 0xaa, 0xde, 0x20,
	0xe2, 0xe0, 0x85, 0x9d, 0xbe, 0xe7, 0xb7, 0xb6, 0x1b, 0x47, 0x2c, 0x7f, 0xcf, 0x29, 0x83, 0x67,
	0x60, 0x6b, 0x58, 0x18, 0xb8, 0x8f, 0xec, 0xef, 0x19, 0x6f, 0x38, 0x1f, 0x96, 0x1e, 0xef, 0x04,
	0x16, 0x9a, 0x32, 0xf1, 0xbf, 0x50, 0xef, 0x74, 0x51, 0x5a, 0x6e, 0x55, 0x9e, 0xb0, 0xde, 0x2a,
	0x04, 0xcf, 0x4b, 0xc7, 0x06, 0x89, 0x4f, 0x9d, 0xa7, 0xfd, 0x45, 0xa2, 0x0b, 0xf0, 0xa2, 0x39,
	0x46, 0xad, 0xf2, 0x13, 0x78, 0x3a, 0x62, 0x72, 0x0a, 0xf7, 0x1c, 0xa3, 0x89, 0x1f, 0x7e, 0xbe,
	0xe6, 0xa8, 0x65, 0x26, 0xcf, 0xdd, 0x7f, 0x9f, 0xee, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76,
	0x56, 0xab, 0x83, 0xb3, 0x09, 0x00, 0x00,
}
