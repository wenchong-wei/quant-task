// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: quant-task.proto

package pub

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 在这里注册需要任务的服务，调用创建任务服务接口，然后读取redis的订阅，以NOT_USED服务为例，频道为task_0
type SourceService int32

const (
	SourceService_NOT_USED      SourceService = 0
	SourceService_ORDER         SourceService = 1
	SourceService_MOCKE_TRANDER SourceService = 2
)

// Enum value maps for SourceService.
var (
	SourceService_name = map[int32]string{
		0: "NOT_USED",
		1: "ORDER",
		2: "MOCKE_TRANDER",
	}
	SourceService_value = map[string]int32{
		"NOT_USED":      0,
		"ORDER":         1,
		"MOCKE_TRANDER": 2,
	}
)

func (x SourceService) Enum() *SourceService {
	p := new(SourceService)
	*p = x
	return p
}

func (x SourceService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceService) Descriptor() protoreflect.EnumDescriptor {
	return file_quant_task_proto_enumTypes[0].Descriptor()
}

func (SourceService) Type() protoreflect.EnumType {
	return &file_quant_task_proto_enumTypes[0]
}

func (x SourceService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceService.Descriptor instead.
func (SourceService) EnumDescriptor() ([]byte, []int) {
	return file_quant_task_proto_rawDescGZIP(), []int{0}
}

type PriceCondition int32

const (
	PriceCondition_LESS    PriceCondition = 0
	PriceCondition_GREATER PriceCondition = 1
)

// Enum value maps for PriceCondition.
var (
	PriceCondition_name = map[int32]string{
		0: "LESS",
		1: "GREATER",
	}
	PriceCondition_value = map[string]int32{
		"LESS":    0,
		"GREATER": 1,
	}
)

func (x PriceCondition) Enum() *PriceCondition {
	p := new(PriceCondition)
	*p = x
	return p
}

func (x PriceCondition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PriceCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_quant_task_proto_enumTypes[1].Descriptor()
}

func (PriceCondition) Type() protoreflect.EnumType {
	return &file_quant_task_proto_enumTypes[1]
}

func (x PriceCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PriceCondition.Descriptor instead.
func (PriceCondition) EnumDescriptor() ([]byte, []int) {
	return file_quant_task_proto_rawDescGZIP(), []int{1}
}

type CreatePriceTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source SourceService `protobuf:"varint,1,opt,name=source,proto3,enum=pub.SourceService" json:"source,omitempty"`
	Tasks  []*PriceTask  `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *CreatePriceTaskReq) Reset() {
	*x = CreatePriceTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceTaskReq) ProtoMessage() {}

func (x *CreatePriceTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_quant_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceTaskReq.ProtoReflect.Descriptor instead.
func (*CreatePriceTaskReq) Descriptor() ([]byte, []int) {
	return file_quant_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePriceTaskReq) GetSource() SourceService {
	if x != nil {
		return x.Source
	}
	return SourceService_NOT_USED
}

func (x *CreatePriceTaskReq) GetTasks() []*PriceTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type PriceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资产类型 1上证 2深证 3美股(包括道琼斯和纳斯达克) 4港股
	AssetType int32 `protobuf:"varint,1,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	// 资产代号
	AssetCode string         `protobuf:"bytes,2,opt,name=asset_code,json=assetCode,proto3" json:"asset_code,omitempty"`
	Condition PriceCondition `protobuf:"varint,3,opt,name=condition,proto3,enum=pub.PriceCondition" json:"condition,omitempty"`
	Price     int64          `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	StartTime int64          `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Message   string         `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PriceTask) Reset() {
	*x = PriceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceTask) ProtoMessage() {}

func (x *PriceTask) ProtoReflect() protoreflect.Message {
	mi := &file_quant_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceTask.ProtoReflect.Descriptor instead.
func (*PriceTask) Descriptor() ([]byte, []int) {
	return file_quant_task_proto_rawDescGZIP(), []int{1}
}

func (x *PriceTask) GetAssetType() int32 {
	if x != nil {
		return x.AssetType
	}
	return 0
}

func (x *PriceTask) GetAssetCode() string {
	if x != nil {
		return x.AssetCode
	}
	return ""
}

func (x *PriceTask) GetCondition() PriceCondition {
	if x != nil {
		return x.Condition
	}
	return PriceCondition_LESS
}

func (x *PriceTask) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PriceTask) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PriceTask) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreatePriceTaskRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePriceTaskRes) Reset() {
	*x = CreatePriceTaskRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quant_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceTaskRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceTaskRes) ProtoMessage() {}

func (x *CreatePriceTaskRes) ProtoReflect() protoreflect.Message {
	mi := &file_quant_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceTaskRes.ProtoReflect.Descriptor instead.
func (*CreatePriceTaskRes) Descriptor() ([]byte, []int) {
	return file_quant_task_proto_rawDescGZIP(), []int{2}
}

var File_quant_task_proto protoreflect.FileDescriptor

var file_quant_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x75, 0x62, 0x22, 0x66, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22,
	0xcb, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x2a, 0x3b, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x4f, 0x43, 0x4b, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x02,
	0x2a, 0x27, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x10, 0x01, 0x32, 0x4d, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x45, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x2f, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x77, 0x65, 0x69, 0x77, 0x65, 0x6e, 0x63, 0x68, 0x6f, 0x6e, 0x67, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2d, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x70, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quant_task_proto_rawDescOnce sync.Once
	file_quant_task_proto_rawDescData = file_quant_task_proto_rawDesc
)

func file_quant_task_proto_rawDescGZIP() []byte {
	file_quant_task_proto_rawDescOnce.Do(func() {
		file_quant_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_quant_task_proto_rawDescData)
	})
	return file_quant_task_proto_rawDescData
}

var file_quant_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_quant_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_quant_task_proto_goTypes = []interface{}{
	(SourceService)(0),         // 0: pub.SourceService
	(PriceCondition)(0),        // 1: pub.PriceCondition
	(*CreatePriceTaskReq)(nil), // 2: pub.CreatePriceTaskReq
	(*PriceTask)(nil),          // 3: pub.PriceTask
	(*CreatePriceTaskRes)(nil), // 4: pub.CreatePriceTaskRes
}
var file_quant_task_proto_depIdxs = []int32{
	0, // 0: pub.CreatePriceTaskReq.source:type_name -> pub.SourceService
	3, // 1: pub.CreatePriceTaskReq.tasks:type_name -> pub.PriceTask
	1, // 2: pub.PriceTask.condition:type_name -> pub.PriceCondition
	2, // 3: pub.Task.CreatePriceTask:input_type -> pub.CreatePriceTaskReq
	4, // 4: pub.Task.CreatePriceTask:output_type -> pub.CreatePriceTaskRes
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_quant_task_proto_init() }
func file_quant_task_proto_init() {
	if File_quant_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quant_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceTaskReq); i {
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
		file_quant_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceTask); i {
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
		file_quant_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceTaskRes); i {
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
			RawDescriptor: file_quant_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quant_task_proto_goTypes,
		DependencyIndexes: file_quant_task_proto_depIdxs,
		EnumInfos:         file_quant_task_proto_enumTypes,
		MessageInfos:      file_quant_task_proto_msgTypes,
	}.Build()
	File_quant_task_proto = out.File
	file_quant_task_proto_rawDesc = nil
	file_quant_task_proto_goTypes = nil
	file_quant_task_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	// 创建价格任务
	CreatePriceTask(ctx context.Context, in *CreatePriceTaskReq, opts ...grpc.CallOption) (*CreatePriceTaskRes, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) CreatePriceTask(ctx context.Context, in *CreatePriceTaskReq, opts ...grpc.CallOption) (*CreatePriceTaskRes, error) {
	out := new(CreatePriceTaskRes)
	err := c.cc.Invoke(ctx, "/pub.Task/CreatePriceTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	// 创建价格任务
	CreatePriceTask(context.Context, *CreatePriceTaskReq) (*CreatePriceTaskRes, error)
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) CreatePriceTask(context.Context, *CreatePriceTaskReq) (*CreatePriceTaskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePriceTask not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_CreatePriceTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePriceTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreatePriceTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pub.Task/CreatePriceTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreatePriceTask(ctx, req.(*CreatePriceTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pub.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePriceTask",
			Handler:    _Task_CreatePriceTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quant-task.proto",
}
