// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: recorder/v1/services.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Requests & Responses
type StartRecordingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resolution    string                 `protobuf:"bytes,1,opt,name=resolution,proto3" json:"resolution,omitempty"` // "720p" | "1080p" | "4k"
	Fps           int32                  `protobuf:"varint,2,opt,name=fps,proto3" json:"fps,omitempty"`
	CameraId      string                 `protobuf:"bytes,4,opt,name=camera_id,json=cameraId,proto3" json:"camera_id,omitempty"`
	MicId         string                 `protobuf:"bytes,5,opt,name=mic_id,json=micId,proto3" json:"mic_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartRecordingRequest) Reset() {
	*x = StartRecordingRequest{}
	mi := &file_recorder_v1_services_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecordingRequest) ProtoMessage() {}

func (x *StartRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecordingRequest.ProtoReflect.Descriptor instead.
func (*StartRecordingRequest) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{0}
}

func (x *StartRecordingRequest) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *StartRecordingRequest) GetFps() int32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *StartRecordingRequest) GetCameraId() string {
	if x != nil {
		return x.CameraId
	}
	return ""
}

func (x *StartRecordingRequest) GetMicId() string {
	if x != nil {
		return x.MicId
	}
	return ""
}

type StartRecordingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordingId   string                 `protobuf:"bytes,1,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	Status        RecordingStatus        `protobuf:"varint,2,opt,name=status,proto3,enum=recorder.v1.RecordingStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartRecordingResponse) Reset() {
	*x = StartRecordingResponse{}
	mi := &file_recorder_v1_services_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRecordingResponse) ProtoMessage() {}

func (x *StartRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRecordingResponse.ProtoReflect.Descriptor instead.
func (*StartRecordingResponse) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{1}
}

func (x *StartRecordingResponse) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *StartRecordingResponse) GetStatus() RecordingStatus {
	if x != nil {
		return x.Status
	}
	return RecordingStatus_RECORDING_STATUS_UNSPECIFIED
}

type StopRecordingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RecordingId   string                 `protobuf:"bytes,1,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRecordingRequest) Reset() {
	*x = StopRecordingRequest{}
	mi := &file_recorder_v1_services_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRecordingRequest) ProtoMessage() {}

func (x *StopRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRecordingRequest.ProtoReflect.Descriptor instead.
func (*StopRecordingRequest) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{2}
}

func (x *StopRecordingRequest) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

type StopRecordingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FilePath      string                 `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopRecordingResponse) Reset() {
	*x = StopRecordingResponse{}
	mi := &file_recorder_v1_services_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRecordingResponse) ProtoMessage() {}

func (x *StopRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRecordingResponse.ProtoReflect.Descriptor instead.
func (*StopRecordingResponse) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{3}
}

func (x *StopRecordingResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type StreamStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamStatusRequest) Reset() {
	*x = StreamStatusRequest{}
	mi := &file_recorder_v1_services_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStatusRequest) ProtoMessage() {}

func (x *StreamStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStatusRequest.ProtoReflect.Descriptor instead.
func (*StreamStatusRequest) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{4}
}

type StreamStatusResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	MicLevel       float32                `protobuf:"fixed32,1,opt,name=mic_level,json=micLevel,proto3" json:"mic_level,omitempty"` // 0-100
	RecordingState RecordingStatus        `protobuf:"varint,2,opt,name=recording_state,json=recordingState,proto3,enum=recorder.v1.RecordingStatus" json:"recording_state,omitempty"`
	Fps            int32                  `protobuf:"varint,3,opt,name=fps,proto3" json:"fps,omitempty"`
	FileSize       int64                  `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"` // Bytes
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *StreamStatusResponse) Reset() {
	*x = StreamStatusResponse{}
	mi := &file_recorder_v1_services_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStatusResponse) ProtoMessage() {}

func (x *StreamStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_v1_services_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStatusResponse.ProtoReflect.Descriptor instead.
func (*StreamStatusResponse) Descriptor() ([]byte, []int) {
	return file_recorder_v1_services_proto_rawDescGZIP(), []int{5}
}

func (x *StreamStatusResponse) GetMicLevel() float32 {
	if x != nil {
		return x.MicLevel
	}
	return 0
}

func (x *StreamStatusResponse) GetRecordingState() RecordingStatus {
	if x != nil {
		return x.RecordingState
	}
	return RecordingStatus_RECORDING_STATUS_UNSPECIFIED
}

func (x *StreamStatusResponse) GetFps() int32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *StreamStatusResponse) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

var File_recorder_v1_services_proto protoreflect.FileDescriptor

var file_recorder_v1_services_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x69,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x63, 0x49,
	0x64, 0x22, 0x71, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x34,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x39, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22,
	0x34, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa9, 0x01, 0x0a,
	0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x63, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x63, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x45, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x70, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x9b, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x20, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x33, 0x79, 0x6d, 0x61, 0x6e, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_recorder_v1_services_proto_rawDescOnce sync.Once
	file_recorder_v1_services_proto_rawDescData []byte
)

func file_recorder_v1_services_proto_rawDescGZIP() []byte {
	file_recorder_v1_services_proto_rawDescOnce.Do(func() {
		file_recorder_v1_services_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_recorder_v1_services_proto_rawDesc), len(file_recorder_v1_services_proto_rawDesc)))
	})
	return file_recorder_v1_services_proto_rawDescData
}

var file_recorder_v1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_recorder_v1_services_proto_goTypes = []any{
	(*StartRecordingRequest)(nil),  // 0: recorder.v1.StartRecordingRequest
	(*StartRecordingResponse)(nil), // 1: recorder.v1.StartRecordingResponse
	(*StopRecordingRequest)(nil),   // 2: recorder.v1.StopRecordingRequest
	(*StopRecordingResponse)(nil),  // 3: recorder.v1.StopRecordingResponse
	(*StreamStatusRequest)(nil),    // 4: recorder.v1.StreamStatusRequest
	(*StreamStatusResponse)(nil),   // 5: recorder.v1.StreamStatusResponse
	(RecordingStatus)(0),           // 6: recorder.v1.RecordingStatus
}
var file_recorder_v1_services_proto_depIdxs = []int32{
	6, // 0: recorder.v1.StartRecordingResponse.status:type_name -> recorder.v1.RecordingStatus
	6, // 1: recorder.v1.StreamStatusResponse.recording_state:type_name -> recorder.v1.RecordingStatus
	0, // 2: recorder.v1.RecorderService.StartRecording:input_type -> recorder.v1.StartRecordingRequest
	2, // 3: recorder.v1.RecorderService.StopRecording:input_type -> recorder.v1.StopRecordingRequest
	4, // 4: recorder.v1.RecorderService.StreamStatus:input_type -> recorder.v1.StreamStatusRequest
	1, // 5: recorder.v1.RecorderService.StartRecording:output_type -> recorder.v1.StartRecordingResponse
	3, // 6: recorder.v1.RecorderService.StopRecording:output_type -> recorder.v1.StopRecordingResponse
	5, // 7: recorder.v1.RecorderService.StreamStatus:output_type -> recorder.v1.StreamStatusResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_recorder_v1_services_proto_init() }
func file_recorder_v1_services_proto_init() {
	if File_recorder_v1_services_proto != nil {
		return
	}
	file_recorder_v1_enums_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_recorder_v1_services_proto_rawDesc), len(file_recorder_v1_services_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recorder_v1_services_proto_goTypes,
		DependencyIndexes: file_recorder_v1_services_proto_depIdxs,
		MessageInfos:      file_recorder_v1_services_proto_msgTypes,
	}.Build()
	File_recorder_v1_services_proto = out.File
	file_recorder_v1_services_proto_goTypes = nil
	file_recorder_v1_services_proto_depIdxs = nil
}
