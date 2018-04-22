// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sum.proto

/*
Package sum is a generated protocol buffer package.

It is generated from these files:
	proto/sum.proto

It has these top-level messages:
	Record
	RecordResponse
	Oracle
	OracleResponse
	Call
	Data
	CallResponse
	ById
	ByName
	ServerInfo
	Empty
*/
package sum

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Record struct {
	Id   uint64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data []float64         `protobuf:"fixed64,2,rep,packed,name=data" json:"data,omitempty"`
	Meta map[string]string `protobuf:"bytes,3,rep,name=meta" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetData() []float64 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Record) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type RecordResponse struct {
	Success bool    `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Record  *Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *RecordResponse) Reset()                    { *m = RecordResponse{} }
func (m *RecordResponse) String() string            { return proto.CompactTextString(m) }
func (*RecordResponse) ProtoMessage()               {}
func (*RecordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RecordResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RecordResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RecordResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Oracle struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *Oracle) Reset()                    { *m = Oracle{} }
func (m *Oracle) String() string            { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()               {}
func (*Oracle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Oracle) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Oracle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Oracle) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OracleResponse struct {
	Success bool      `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string    `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Oracles []*Oracle `protobuf:"bytes,3,rep,name=oracles" json:"oracles,omitempty"`
}

func (m *OracleResponse) Reset()                    { *m = OracleResponse{} }
func (m *OracleResponse) String() string            { return proto.CompactTextString(m) }
func (*OracleResponse) ProtoMessage()               {}
func (*OracleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OracleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *OracleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *OracleResponse) GetOracles() []*Oracle {
	if m != nil {
		return m.Oracles
	}
	return nil
}

type Call struct {
	OracleId uint64   `protobuf:"varint,1,opt,name=oracle_id,json=oracleId" json:"oracle_id,omitempty"`
	Args     []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Call) GetOracleId() uint64 {
	if m != nil {
		return m.OracleId
	}
	return 0
}

func (m *Call) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type Data struct {
	Compressed bool   `protobuf:"varint,1,opt,name=compressed" json:"compressed,omitempty"`
	Payload    []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Data) GetCompressed() bool {
	if m != nil {
		return m.Compressed
	}
	return false
}

func (m *Data) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CallResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data    *Data  `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *CallResponse) Reset()                    { *m = CallResponse{} }
func (m *CallResponse) String() string            { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()               {}
func (*CallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CallResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CallResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CallResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ById struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ById) Reset()                    { *m = ById{} }
func (m *ById) String() string            { return proto.CompactTextString(m) }
func (*ById) ProtoMessage()               {}
func (*ById) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ById) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ByName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ByName) Reset()                    { *m = ByName{} }
func (m *ByName) String() string            { return proto.CompactTextString(m) }
func (*ByName) ProtoMessage()               {}
func (*ByName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ByName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ServerInfo struct {
	Version string   `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Uptime  uint64   `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	Pid     uint64   `protobuf:"varint,3,opt,name=pid" json:"pid,omitempty"`
	Uid     uint64   `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	Argv    []string `protobuf:"bytes,5,rep,name=argv" json:"argv,omitempty"`
	Records uint64   `protobuf:"varint,6,opt,name=records" json:"records,omitempty"`
	Oracles uint64   `protobuf:"varint,7,opt,name=oracles" json:"oracles,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *ServerInfo) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ServerInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ServerInfo) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

func (m *ServerInfo) GetRecords() uint64 {
	if m != nil {
		return m.Records
	}
	return 0
}

func (m *ServerInfo) GetOracles() uint64 {
	if m != nil {
		return m.Oracles
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Record)(nil), "sum.Record")
	proto.RegisterType((*RecordResponse)(nil), "sum.RecordResponse")
	proto.RegisterType((*Oracle)(nil), "sum.Oracle")
	proto.RegisterType((*OracleResponse)(nil), "sum.OracleResponse")
	proto.RegisterType((*Call)(nil), "sum.Call")
	proto.RegisterType((*Data)(nil), "sum.Data")
	proto.RegisterType((*CallResponse)(nil), "sum.CallResponse")
	proto.RegisterType((*ById)(nil), "sum.ById")
	proto.RegisterType((*ByName)(nil), "sum.ByName")
	proto.RegisterType((*ServerInfo)(nil), "sum.ServerInfo")
	proto.RegisterType((*Empty)(nil), "sum.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SumService service

type SumServiceClient interface {
	// vectors CRUD
	CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error)
	UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error)
	ReadRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error)
	DeleteRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error)
	// oracles CRUD
	CreateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error)
	UpdateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error)
	ReadOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error)
	FindOracle(ctx context.Context, in *ByName, opts ...grpc.CallOption) (*OracleResponse, error)
	DeleteOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error)
	// execute a call to a oracle given its id
	Run(ctx context.Context, in *Call, opts ...grpc.CallOption) (*CallResponse, error)
	// get info about the service
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfo, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/CreateRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/UpdateRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) ReadRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/ReadRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) DeleteRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/DeleteRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) CreateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/CreateOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) UpdateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/UpdateOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) ReadOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/ReadOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) FindOracle(ctx context.Context, in *ByName, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/FindOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) DeleteOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/DeleteOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Run(ctx context.Context, in *Call, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := grpc.Invoke(ctx, "/sum.SumService/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SumService service

type SumServiceServer interface {
	// vectors CRUD
	CreateRecord(context.Context, *Record) (*RecordResponse, error)
	UpdateRecord(context.Context, *Record) (*RecordResponse, error)
	ReadRecord(context.Context, *ById) (*RecordResponse, error)
	DeleteRecord(context.Context, *ById) (*RecordResponse, error)
	// oracles CRUD
	CreateOracle(context.Context, *Oracle) (*OracleResponse, error)
	UpdateOracle(context.Context, *Oracle) (*OracleResponse, error)
	ReadOracle(context.Context, *ById) (*OracleResponse, error)
	FindOracle(context.Context, *ByName) (*OracleResponse, error)
	DeleteOracle(context.Context, *ById) (*OracleResponse, error)
	// execute a call to a oracle given its id
	Run(context.Context, *Call) (*CallResponse, error)
	// get info about the service
	Info(context.Context, *Empty) (*ServerInfo, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).CreateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).UpdateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_ReadRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).ReadRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/ReadRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).ReadRecord(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).DeleteRecord(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_CreateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).CreateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/CreateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).CreateOracle(ctx, req.(*Oracle))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_UpdateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).UpdateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/UpdateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).UpdateOracle(ctx, req.(*Oracle))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_ReadOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).ReadOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/ReadOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).ReadOracle(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_FindOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).FindOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/FindOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).FindOracle(ctx, req.(*ByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_DeleteOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).DeleteOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/DeleteOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).DeleteOracle(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Call)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Run(ctx, req.(*Call))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecord",
			Handler:    _SumService_CreateRecord_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _SumService_UpdateRecord_Handler,
		},
		{
			MethodName: "ReadRecord",
			Handler:    _SumService_ReadRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _SumService_DeleteRecord_Handler,
		},
		{
			MethodName: "CreateOracle",
			Handler:    _SumService_CreateOracle_Handler,
		},
		{
			MethodName: "UpdateOracle",
			Handler:    _SumService_UpdateOracle_Handler,
		},
		{
			MethodName: "ReadOracle",
			Handler:    _SumService_ReadOracle_Handler,
		},
		{
			MethodName: "FindOracle",
			Handler:    _SumService_FindOracle_Handler,
		},
		{
			MethodName: "DeleteOracle",
			Handler:    _SumService_DeleteOracle_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _SumService_Run_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _SumService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sum.proto",
}

func init() { proto.RegisterFile("proto/sum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x63, 0xc7, 0xa9, 0x27, 0x51, 0x0a, 0x0b, 0x54, 0x56, 0xf8, 0x51, 0xb4, 0xa8, 0x52,
	0xb8, 0x84, 0x2a, 0x1c, 0x8a, 0x38, 0x55, 0xfd, 0x41, 0xca, 0x01, 0x90, 0xb6, 0xe2, 0xc0, 0x09,
	0x2d, 0xf6, 0x52, 0x59, 0xf8, 0x4f, 0xbb, 0x76, 0x24, 0xbf, 0x05, 0x8f, 0xc0, 0x99, 0xa7, 0x44,
	0x3b, 0xbb, 0x76, 0x1d, 0x85, 0x0a, 0xda, 0xdb, 0xcc, 0xec, 0x7c, 0x33, 0xdf, 0xcc, 0x7c, 0x36,
	0x1c, 0x94, 0xb2, 0xa8, 0x8a, 0xd7, 0xaa, 0xce, 0x96, 0x68, 0x11, 0x57, 0xd5, 0x19, 0xfd, 0xe9,
	0x80, 0xcf, 0x44, 0x54, 0xc8, 0x98, 0x4c, 0x61, 0x90, 0xc4, 0xa1, 0x33, 0x77, 0x16, 0x1e, 0x1b,
	0x24, 0x31, 0x21, 0xe0, 0xc5, 0xbc, 0xe2, 0xe1, 0x60, 0xee, 0x2e, 0x1c, 0x86, 0x36, 0x79, 0x05,
	0x5e, 0x26, 0x2a, 0x1e, 0xba, 0x73, 0x77, 0x31, 0x5e, 0x3d, 0x59, 0xea, 0x6a, 0x06, 0xbe, 0xfc,
	0x20, 0x2a, 0x7e, 0x99, 0x57, 0xb2, 0x61, 0x98, 0x32, 0x3b, 0x81, 0xa0, 0x0b, 0x91, 0x07, 0xe0,
	0xfe, 0x10, 0x0d, 0x16, 0x0f, 0x98, 0x36, 0xc9, 0x63, 0x18, 0x6e, 0x78, 0x5a, 0x8b, 0x70, 0x80,
	0x31, 0xe3, 0xbc, 0x1b, 0xbc, 0x75, 0x28, 0x87, 0xa9, 0x29, 0xc9, 0x84, 0x2a, 0x8b, 0x5c, 0x09,
	0x12, 0xc2, 0x48, 0xd5, 0x51, 0x24, 0x94, 0xc2, 0x0a, 0xfb, 0xac, 0x75, 0x75, 0xdd, 0x4c, 0x5d,
	0xdb, 0x1a, 0xda, 0x24, 0x2f, 0xc1, 0x97, 0x88, 0x0e, 0xdd, 0xb9, 0xb3, 0x18, 0xaf, 0xc6, 0x3d,
	0x8e, 0xcc, 0x3e, 0xd1, 0x53, 0xf0, 0x3f, 0x49, 0x1e, 0xa5, 0xe2, 0x6f, 0x43, 0xe7, 0x3c, 0x6b,
	0x59, 0xa1, 0xad, 0x63, 0x51, 0x11, 0x0b, 0x2c, 0x18, 0x30, 0xb4, 0x69, 0x04, 0x53, 0x53, 0xe1,
	0x5e, 0x24, 0x8f, 0x60, 0x54, 0x20, 0x5a, 0xd9, 0x4d, 0x1a, 0x96, 0xb6, 0x62, 0xfb, 0x46, 0x4f,
	0xc0, 0x3b, 0xe7, 0x69, 0x4a, 0x9e, 0x42, 0x60, 0x42, 0x5f, 0x3b, 0xae, 0xfb, 0x26, 0xb0, 0x46,
	0xc6, 0x5c, 0x5e, 0x2b, 0x3c, 0x53, 0xc0, 0xd0, 0xa6, 0xa7, 0xe0, 0x5d, 0xe8, 0x73, 0xbd, 0x00,
	0x88, 0x8a, 0xac, 0x94, 0x42, 0x29, 0x11, 0x5b, 0x5a, 0xbd, 0x88, 0xe6, 0x5c, 0xf2, 0x26, 0x2d,
	0x78, 0x8c, 0xec, 0x26, 0xac, 0x75, 0xe9, 0x17, 0x98, 0xe8, 0xd6, 0xf7, 0x9a, 0xee, 0xb9, 0x15,
	0x8e, 0x39, 0x40, 0x80, 0xa3, 0x69, 0x3a, 0x46, 0x43, 0xf4, 0x10, 0xbc, 0xb3, 0x66, 0xbd, 0xa3,
	0x37, 0xfa, 0x0c, 0xfc, 0xb3, 0xe6, 0xa3, 0x5d, 0x38, 0x1e, 0xc1, 0xb9, 0x39, 0x02, 0xfd, 0xed,
	0x00, 0x5c, 0x09, 0xb9, 0x11, 0x72, 0x9d, 0x7f, 0x2f, 0x34, 0x9f, 0x8d, 0x90, 0x2a, 0x29, 0x72,
	0x9b, 0xd5, 0xba, 0xe4, 0x10, 0xfc, 0xba, 0xac, 0x12, 0x7b, 0x43, 0x8f, 0x59, 0x4f, 0xf3, 0x2c,
	0x13, 0xa3, 0x0a, 0x8f, 0x69, 0x53, 0x47, 0xea, 0x24, 0x0e, 0x3d, 0x13, 0xa9, 0x93, 0x76, 0x97,
	0x9b, 0x70, 0xd8, 0xed, 0x72, 0xa3, 0x3b, 0x19, 0xd5, 0xa8, 0xd0, 0xc7, 0xcc, 0xd6, 0xd5, 0x2f,
	0xed, 0x15, 0x47, 0xe6, 0xa5, 0x3d, 0xdc, 0x08, 0x86, 0x97, 0x59, 0x59, 0x35, 0xab, 0x5f, 0x1e,
	0xc0, 0x55, 0x9d, 0x69, 0xe2, 0x49, 0x24, 0xc8, 0x0a, 0x26, 0xe7, 0x52, 0xf0, 0x4a, 0xd8, 0x4f,
	0xae, 0x2f, 0xce, 0xd9, 0xa3, 0xbe, 0x52, 0xed, 0xde, 0xe9, 0x9e, 0xc6, 0x7c, 0x2e, 0xe3, 0xbb,
	0x61, 0x96, 0x00, 0x4c, 0xf0, 0xd8, 0x22, 0xcc, 0x05, 0xf4, 0xce, 0x6f, 0xcb, 0x3f, 0x86, 0xc9,
	0x85, 0x48, 0x45, 0xd7, 0xe3, 0xdf, 0x88, 0x6e, 0x12, 0xfb, 0x1d, 0xf5, 0x05, 0x6c, 0x31, 0xdb,
	0xdf, 0x47, 0x7f, 0x92, 0x3b, 0x60, 0xec, 0x24, 0x16, 0xb1, 0xc3, 0x6b, 0x27, 0xff, 0x18, 0xe0,
	0x7d, 0x92, 0xc7, 0x5b, 0x1d, 0x8c, 0xaa, 0x6e, 0x47, 0xd8, 0xd9, 0xff, 0xbb, 0xc7, 0x11, 0xb8,
	0xac, 0xce, 0x6d, 0xa2, 0xfe, 0x4a, 0x66, 0x0f, 0x3b, 0x73, 0x2b, 0xcd, 0x43, 0xa9, 0x02, 0x3e,
	0xa2, 0x1e, 0x66, 0x07, 0x68, 0xdf, 0xe8, 0x98, 0xee, 0x7d, 0xf3, 0xf1, 0x6f, 0xfc, 0xe6, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc7, 0xe1, 0xfb, 0xa0, 0x05, 0x00, 0x00,
}
