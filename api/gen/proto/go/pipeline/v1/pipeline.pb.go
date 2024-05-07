// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: pipeline/v1/pipeline.proto

package pipelinev1

import (
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

// A Pipeline is a self-contained snippet of configuration that can be assigned to
// collectors based on matchers.
type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the pipeline. This is the unique identifier for the pipeline.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The configuration contents of the pipeline to be used by collectors.
	Contents string `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	// Matchers are used to match against collectors and assign pipelines to them.
	// They follow the syntax of the Prometheus [Alertmanager matchers](https://prometheus.io/docs/alerting/latest/configuration/#matcher)
	Matchers []string `protobuf:"bytes,3,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *Pipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pipeline) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

func (x *Pipeline) GetMatchers() []string {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type Pipelines struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipelines []*Pipeline `protobuf:"bytes,1,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
}

func (x *Pipelines) Reset() {
	*x = Pipelines{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipelines) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipelines) ProtoMessage() {}

func (x *Pipelines) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipelines.ProtoReflect.Descriptor instead.
func (*Pipelines) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *Pipelines) GetPipelines() []*Pipeline {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

type GetPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPipelineRequest) Reset() {
	*x = GetPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPipelineRequest) ProtoMessage() {}

func (x *GetPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPipelineRequest.ProtoReflect.Descriptor instead.
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *GetPipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListPipelinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPipelinesRequest) Reset() {
	*x = ListPipelinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPipelinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPipelinesRequest) ProtoMessage() {}

func (x *ListPipelinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPipelinesRequest.ProtoReflect.Descriptor instead.
func (*ListPipelinesRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{3}
}

type CreatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *CreatePipelineRequest) Reset() {
	*x = CreatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineRequest) ProtoMessage() {}

func (x *CreatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineRequest.ProtoReflect.Descriptor instead.
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePipelineRequest) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type UpdatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *UpdatePipelineRequest) Reset() {
	*x = UpdatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePipelineRequest) ProtoMessage() {}

func (x *UpdatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePipelineRequest.ProtoReflect.Descriptor instead.
func (*UpdatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePipelineRequest) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type DeletePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePipelineRequest) Reset() {
	*x = DeletePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePipelineRequest) ProtoMessage() {}

func (x *DeletePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePipelineRequest.ProtoReflect.Descriptor instead.
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeletePipelineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePipelineResponse) Reset() {
	*x = DeletePipelineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_v1_pipeline_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePipelineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePipelineResponse) ProtoMessage() {}

func (x *DeletePipelineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_v1_pipeline_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePipelineResponse.ProtoReflect.Descriptor instead.
func (*DeletePipelineResponse) Descriptor() ([]byte, []int) {
	return file_pipeline_v1_pipeline_proto_rawDescGZIP(), []int{7}
}

var File_pipeline_v1_pipeline_proto protoreflect.FileDescriptor

var file_pipeline_v1_pipeline_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x08, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x22, 0x40, 0x0a, 0x09, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x33,
	0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31,
	0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x22, 0x4a, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x2b, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xee, 0x03, 0x0a, 0x0f, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x4a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x21, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x53, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x22, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x79, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipeline_v1_pipeline_proto_rawDescOnce sync.Once
	file_pipeline_v1_pipeline_proto_rawDescData = file_pipeline_v1_pipeline_proto_rawDesc
)

func file_pipeline_v1_pipeline_proto_rawDescGZIP() []byte {
	file_pipeline_v1_pipeline_proto_rawDescOnce.Do(func() {
		file_pipeline_v1_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipeline_v1_pipeline_proto_rawDescData)
	})
	return file_pipeline_v1_pipeline_proto_rawDescData
}

var file_pipeline_v1_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pipeline_v1_pipeline_proto_goTypes = []interface{}{
	(*Pipeline)(nil),               // 0: pipeline.v1.Pipeline
	(*Pipelines)(nil),              // 1: pipeline.v1.Pipelines
	(*GetPipelineRequest)(nil),     // 2: pipeline.v1.GetPipelineRequest
	(*ListPipelinesRequest)(nil),   // 3: pipeline.v1.ListPipelinesRequest
	(*CreatePipelineRequest)(nil),  // 4: pipeline.v1.CreatePipelineRequest
	(*UpdatePipelineRequest)(nil),  // 5: pipeline.v1.UpdatePipelineRequest
	(*DeletePipelineRequest)(nil),  // 6: pipeline.v1.DeletePipelineRequest
	(*DeletePipelineResponse)(nil), // 7: pipeline.v1.DeletePipelineResponse
}
var file_pipeline_v1_pipeline_proto_depIdxs = []int32{
	0, // 0: pipeline.v1.Pipelines.pipelines:type_name -> pipeline.v1.Pipeline
	0, // 1: pipeline.v1.CreatePipelineRequest.pipeline:type_name -> pipeline.v1.Pipeline
	0, // 2: pipeline.v1.UpdatePipelineRequest.pipeline:type_name -> pipeline.v1.Pipeline
	2, // 3: pipeline.v1.PipelineService.GetPipeline:input_type -> pipeline.v1.GetPipelineRequest
	3, // 4: pipeline.v1.PipelineService.ListPipelines:input_type -> pipeline.v1.ListPipelinesRequest
	4, // 5: pipeline.v1.PipelineService.CreatePipeline:input_type -> pipeline.v1.CreatePipelineRequest
	5, // 6: pipeline.v1.PipelineService.UpdatePipeline:input_type -> pipeline.v1.UpdatePipelineRequest
	4, // 7: pipeline.v1.PipelineService.CreateOrUpdatePipeline:input_type -> pipeline.v1.CreatePipelineRequest
	6, // 8: pipeline.v1.PipelineService.DeletePipeline:input_type -> pipeline.v1.DeletePipelineRequest
	0, // 9: pipeline.v1.PipelineService.GetPipeline:output_type -> pipeline.v1.Pipeline
	1, // 10: pipeline.v1.PipelineService.ListPipelines:output_type -> pipeline.v1.Pipelines
	0, // 11: pipeline.v1.PipelineService.CreatePipeline:output_type -> pipeline.v1.Pipeline
	0, // 12: pipeline.v1.PipelineService.UpdatePipeline:output_type -> pipeline.v1.Pipeline
	0, // 13: pipeline.v1.PipelineService.CreateOrUpdatePipeline:output_type -> pipeline.v1.Pipeline
	7, // 14: pipeline.v1.PipelineService.DeletePipeline:output_type -> pipeline.v1.DeletePipelineResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pipeline_v1_pipeline_proto_init() }
func file_pipeline_v1_pipeline_proto_init() {
	if File_pipeline_v1_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipeline_v1_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipelines); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPipelineRequest); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPipelinesRequest); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineRequest); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePipelineRequest); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePipelineRequest); i {
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
		file_pipeline_v1_pipeline_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePipelineResponse); i {
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
			RawDescriptor: file_pipeline_v1_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pipeline_v1_pipeline_proto_goTypes,
		DependencyIndexes: file_pipeline_v1_pipeline_proto_depIdxs,
		MessageInfos:      file_pipeline_v1_pipeline_proto_msgTypes,
	}.Build()
	File_pipeline_v1_pipeline_proto = out.File
	file_pipeline_v1_pipeline_proto_rawDesc = nil
	file_pipeline_v1_pipeline_proto_goTypes = nil
	file_pipeline_v1_pipeline_proto_depIdxs = nil
}
