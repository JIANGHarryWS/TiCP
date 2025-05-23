// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/job/scheduler.proto

package job

import (
	application "github.com/yuansuan/ticp/common/project-root-api/proto/job/application"
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

type SchedulerContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App         *application.Application                `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	AppRuntimes []*application.ServerApplicationRuntime `protobuf:"bytes,2,rep,name=app_runtimes,json=appRuntimes,proto3" json:"app_runtimes,omitempty"`
	MinCores    int64                                   `protobuf:"varint,3,opt,name=min_cores,json=minCores,proto3" json:"min_cores,omitempty"`
	HasLocalSc  bool                                    `protobuf:"varint,5,opt,name=has_local_sc,json=hasLocalSc,proto3" json:"has_local_sc,omitempty"`
	LocalScIds  []string                                `protobuf:"bytes,6,rep,name=local_sc_ids,json=localScIds,proto3" json:"local_sc_ids,omitempty"`
	// 0: upload_file_info,  1: volume_mounts
	VolumeType     int64           `protobuf:"varint,7,opt,name=volume_type,json=volumeType,proto3" json:"volume_type,omitempty"`
	UploadFileInfo *UploadFileInfo `protobuf:"bytes,4,opt,name=upload_file_info,json=uploadFileInfo,proto3" json:"upload_file_info,omitempty"`
	VolumeMounts   *VolumeMounts   `protobuf:"bytes,8,opt,name=volume_mounts,json=volumeMounts,proto3" json:"volume_mounts,omitempty"`
}

func (x *SchedulerContext) Reset() {
	*x = SchedulerContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_job_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerContext) ProtoMessage() {}

func (x *SchedulerContext) ProtoReflect() protoreflect.Message {
	mi := &file_proto_job_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerContext.ProtoReflect.Descriptor instead.
func (*SchedulerContext) Descriptor() ([]byte, []int) {
	return file_proto_job_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulerContext) GetApp() *application.Application {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *SchedulerContext) GetAppRuntimes() []*application.ServerApplicationRuntime {
	if x != nil {
		return x.AppRuntimes
	}
	return nil
}

func (x *SchedulerContext) GetMinCores() int64 {
	if x != nil {
		return x.MinCores
	}
	return 0
}

func (x *SchedulerContext) GetHasLocalSc() bool {
	if x != nil {
		return x.HasLocalSc
	}
	return false
}

func (x *SchedulerContext) GetLocalScIds() []string {
	if x != nil {
		return x.LocalScIds
	}
	return nil
}

func (x *SchedulerContext) GetVolumeType() int64 {
	if x != nil {
		return x.VolumeType
	}
	return 0
}

func (x *SchedulerContext) GetUploadFileInfo() *UploadFileInfo {
	if x != nil {
		return x.UploadFileInfo
	}
	return nil
}

func (x *SchedulerContext) GetVolumeMounts() *VolumeMounts {
	if x != nil {
		return x.VolumeMounts
	}
	return nil
}

var File_proto_job_scheduler_proto protoreflect.FileDescriptor

var file_proto_job_scheduler_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62,
	0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x03, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12,
	0x48, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x61, 0x70,
	0x70, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69,
	0x6e, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61,
	0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x73, 0x63, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0d, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x2e, 0x79, 0x75, 0x61, 0x6e, 0x73, 0x75,
	0x61, 0x6e, 0x2e, 0x63, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x72, 0x6f,
	0x6f, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_job_scheduler_proto_rawDescOnce sync.Once
	file_proto_job_scheduler_proto_rawDescData = file_proto_job_scheduler_proto_rawDesc
)

func file_proto_job_scheduler_proto_rawDescGZIP() []byte {
	file_proto_job_scheduler_proto_rawDescOnce.Do(func() {
		file_proto_job_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_job_scheduler_proto_rawDescData)
	})
	return file_proto_job_scheduler_proto_rawDescData
}

var file_proto_job_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_job_scheduler_proto_goTypes = []interface{}{
	(*SchedulerContext)(nil),                     // 0: job.SchedulerContext
	(*application.Application)(nil),              // 1: application.Application
	(*application.ServerApplicationRuntime)(nil), // 2: application.ServerApplicationRuntime
	(*UploadFileInfo)(nil),                       // 3: job.UploadFileInfo
	(*VolumeMounts)(nil),                         // 4: job.VolumeMounts
}
var file_proto_job_scheduler_proto_depIdxs = []int32{
	1, // 0: job.SchedulerContext.app:type_name -> application.Application
	2, // 1: job.SchedulerContext.app_runtimes:type_name -> application.ServerApplicationRuntime
	3, // 2: job.SchedulerContext.upload_file_info:type_name -> job.UploadFileInfo
	4, // 3: job.SchedulerContext.volume_mounts:type_name -> job.VolumeMounts
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_job_scheduler_proto_init() }
func file_proto_job_scheduler_proto_init() {
	if File_proto_job_scheduler_proto != nil {
		return
	}
	file_proto_job_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_job_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerContext); i {
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
			RawDescriptor: file_proto_job_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_job_scheduler_proto_goTypes,
		DependencyIndexes: file_proto_job_scheduler_proto_depIdxs,
		MessageInfos:      file_proto_job_scheduler_proto_msgTypes,
	}.Build()
	File_proto_job_scheduler_proto = out.File
	file_proto_job_scheduler_proto_rawDesc = nil
	file_proto_job_scheduler_proto_goTypes = nil
	file_proto_job_scheduler_proto_depIdxs = nil
}
