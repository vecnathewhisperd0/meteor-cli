// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: raystack/assets/v1beta2/model.proto

package assetsv1beta2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Schema of the model's inputs and outputs. Strongly inspired by
// https://mlflow.org/docs/latest/python_api/mlflow.models.html#mlflow.models.ModelSignature.
type ModelSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*ModelSignature_Parameter `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*ModelSignature_Parameter `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ModelSignature) Reset() {
	*x = ModelSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelSignature) ProtoMessage() {}

func (x *ModelSignature) ProtoReflect() protoreflect.Message {
	mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelSignature.ProtoReflect.Descriptor instead.
func (*ModelSignature) Descriptor() ([]byte, []int) {
	return file_raystack_assets_v1beta2_model_proto_rawDescGZIP(), []int{0}
}

func (x *ModelSignature) GetInputs() []*ModelSignature_Parameter {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ModelSignature) GetOutputs() []*ModelSignature_Parameter {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type ModelVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The schema of a model version’s inputs and outputs.
	Signature *ModelSignature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Status of the model version. ex: pending/ready/serving/terminated etc.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Version of the model
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// List of attributes the model version has. This could include the following:
	//   - mlflow_run_id[string]: MLFlow expriment run ID associated with the model
	//     version. Relevant for Merlin models.
	//   - mlflow_run_url[string]: URL of MLFlow experiment run associated with
	//     the model. Relevant for Merlin models.
	//   - endpoint_url[string]: Endpoint that the model is serving requests on.
	//     Ex: http://<model_name>-<version>.<project_name>.<merlin_base_url>.
	//   - version_endpoint_url[string]: Endpoint that the model is serving
	//     requests on for the specific version. Ex:
	//     http://<model_name>-<version>.<project_name>.<merlin_base_url>.
	//   - traffic[double]: Percentage of traffic being served by this version of
	//     the model.
	//   - weight[double]: Weightage for the model version endpoint.
	//   - params[map<string, string>]: Parameters for the Model's run.
	//   - metrics[map<string, double>]: Metrics for the model's run.
	Attributes *structpb.Struct `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Optional. List of labels the model version has.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The timestamp of the model's creation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The timestamp when the model was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *ModelVersion) Reset() {
	*x = ModelVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelVersion) ProtoMessage() {}

func (x *ModelVersion) ProtoReflect() protoreflect.Message {
	mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelVersion.ProtoReflect.Descriptor instead.
func (*ModelVersion) Descriptor() ([]byte, []int) {
	return file_raystack_assets_v1beta2_model_proto_rawDescGZIP(), []int{1}
}

func (x *ModelVersion) GetSignature() *ModelSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ModelVersion) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ModelVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ModelVersion) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ModelVersion) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ModelVersion) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ModelVersion) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Model represents a Data Science Model commonly used for Machine Learning
// (ML). Models are algorithms trained on data to find patterns or make
// predictions. Models typically consume ML features to generate a meaningful
// output. The inputs can also include contextual information that is made
// available in realtime as part of the request to the model server.
type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: Model's namespace or project.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Flavor of the ML Model. ex: pytorch, tensorflow etc.
	Flavor string `protobuf:"bytes,2,opt,name=flavor,proto3" json:"flavor,omitempty"`
	// Optional: Algorithm used to train the ML Model.
	Algorithm string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// Status of the model. ex: active/deleted etc.
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Versions of the model, similar to experiment runs in MLFlow and model
	// version in Merlin.
	Versions []*ModelVersion `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty"`
	// List of attributes the model version has. This could include the following:
	//   - project_id[double]: ID of project the model is present in.
	//   - project_name[string]: Name of project the model is present in.
	//   - mlflow_experiment_id[double]: MLFlow experiment ID associated with the
	//     model. Relevant for Merlin models.
	//   - mlflow_experiment_url[string]: URL of MLFlow experiment associated with
	//     the model. Relevant for Merlin models.
	//   - endpoint_urls[repeated string]: List of URLs associated with endpoints
	//     that are serving requests for the model. Relevant for Merlin models.
	Attributes *structpb.Struct `protobuf:"bytes,6,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The timestamp of the model's creation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The timestamp when the model was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_raystack_assets_v1beta2_model_proto_rawDescGZIP(), []int{2}
}

func (x *Model) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Model) GetFlavor() string {
	if x != nil {
		return x.Flavor
	}
	return ""
}

func (x *Model) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *Model) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Model) GetVersions() []*ModelVersion {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *Model) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Model) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Model) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Specification of name and type of a single column in a dataset.
type ModelSignature_Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: Name of the input or output parameter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data type of the parameter. Ex: boolean, double, numpy's dtypes etc.
	DataType string `protobuf:"bytes,2,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	// Optional: The tensor shape.
	Shape []int64 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
}

func (x *ModelSignature_Parameter) Reset() {
	*x = ModelSignature_Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelSignature_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelSignature_Parameter) ProtoMessage() {}

func (x *ModelSignature_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_raystack_assets_v1beta2_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelSignature_Parameter.ProtoReflect.Descriptor instead.
func (*ModelSignature_Parameter) Descriptor() ([]byte, []int) {
	return file_raystack_assets_v1beta2_model_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ModelSignature_Parameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelSignature_Parameter) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *ModelSignature_Parameter) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

var File_raystack_assets_v1beta2_model_proto protoreflect.FileDescriptor

var file_raystack_assets_v1beta2_model_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01,
	0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x49, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72,
	0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x52, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x22, 0xc0, 0x03, 0x0a,
	0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xe9, 0x02, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x32, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x59, 0x0a, 0x12, 0x69,
	0x6f, 0x2e, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x42, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x79, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raystack_assets_v1beta2_model_proto_rawDescOnce sync.Once
	file_raystack_assets_v1beta2_model_proto_rawDescData = file_raystack_assets_v1beta2_model_proto_rawDesc
)

func file_raystack_assets_v1beta2_model_proto_rawDescGZIP() []byte {
	file_raystack_assets_v1beta2_model_proto_rawDescOnce.Do(func() {
		file_raystack_assets_v1beta2_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_raystack_assets_v1beta2_model_proto_rawDescData)
	})
	return file_raystack_assets_v1beta2_model_proto_rawDescData
}

var file_raystack_assets_v1beta2_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_raystack_assets_v1beta2_model_proto_goTypes = []interface{}{
	(*ModelSignature)(nil),           // 0: raystack.assets.v1beta2.ModelSignature
	(*ModelVersion)(nil),             // 1: raystack.assets.v1beta2.ModelVersion
	(*Model)(nil),                    // 2: raystack.assets.v1beta2.Model
	(*ModelSignature_Parameter)(nil), // 3: raystack.assets.v1beta2.ModelSignature.Parameter
	nil,                              // 4: raystack.assets.v1beta2.ModelVersion.LabelsEntry
	(*structpb.Struct)(nil),          // 5: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_raystack_assets_v1beta2_model_proto_depIdxs = []int32{
	3,  // 0: raystack.assets.v1beta2.ModelSignature.inputs:type_name -> raystack.assets.v1beta2.ModelSignature.Parameter
	3,  // 1: raystack.assets.v1beta2.ModelSignature.outputs:type_name -> raystack.assets.v1beta2.ModelSignature.Parameter
	0,  // 2: raystack.assets.v1beta2.ModelVersion.signature:type_name -> raystack.assets.v1beta2.ModelSignature
	5,  // 3: raystack.assets.v1beta2.ModelVersion.attributes:type_name -> google.protobuf.Struct
	4,  // 4: raystack.assets.v1beta2.ModelVersion.labels:type_name -> raystack.assets.v1beta2.ModelVersion.LabelsEntry
	6,  // 5: raystack.assets.v1beta2.ModelVersion.create_time:type_name -> google.protobuf.Timestamp
	6,  // 6: raystack.assets.v1beta2.ModelVersion.update_time:type_name -> google.protobuf.Timestamp
	1,  // 7: raystack.assets.v1beta2.Model.versions:type_name -> raystack.assets.v1beta2.ModelVersion
	5,  // 8: raystack.assets.v1beta2.Model.attributes:type_name -> google.protobuf.Struct
	6,  // 9: raystack.assets.v1beta2.Model.create_time:type_name -> google.protobuf.Timestamp
	6,  // 10: raystack.assets.v1beta2.Model.update_time:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_raystack_assets_v1beta2_model_proto_init() }
func file_raystack_assets_v1beta2_model_proto_init() {
	if File_raystack_assets_v1beta2_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raystack_assets_v1beta2_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelSignature); i {
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
		file_raystack_assets_v1beta2_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelVersion); i {
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
		file_raystack_assets_v1beta2_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_raystack_assets_v1beta2_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelSignature_Parameter); i {
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
			RawDescriptor: file_raystack_assets_v1beta2_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_raystack_assets_v1beta2_model_proto_goTypes,
		DependencyIndexes: file_raystack_assets_v1beta2_model_proto_depIdxs,
		MessageInfos:      file_raystack_assets_v1beta2_model_proto_msgTypes,
	}.Build()
	File_raystack_assets_v1beta2_model_proto = out.File
	file_raystack_assets_v1beta2_model_proto_rawDesc = nil
	file_raystack_assets_v1beta2_model_proto_goTypes = nil
	file_raystack_assets_v1beta2_model_proto_depIdxs = nil
}
