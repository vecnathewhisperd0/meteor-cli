// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: odpf/assets/dashboard.proto

package assets

import (
	common "github.com/odpf/meteor/models/odpf/assets/common"
	facets "github.com/odpf/meteor/models/odpf/assets/facets"
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

// Dashboard is a resource that represents a dashboard.
type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Representation of the resource
	Resource *common.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The description of the dashboard.
	// Example: "This dashboard was created by the Metabase team."
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The list of the charts in the dashboard.
	// For an example, check the schema of the chart.
	Charts []*Chart `protobuf:"bytes,21,rep,name=charts,proto3" json:"charts,omitempty"`
	// The ownership of the topic.
	// For an example check out ownership.
	Ownership *facets.Ownership `protobuf:"bytes,31,opt,name=ownership,proto3" json:"ownership,omitempty"`
	// List of the user's custom properties.
	// Properties facet can be used to set custom properties, tags and labels for a user.
	Properties *facets.Properties `protobuf:"bytes,32,opt,name=properties,proto3" json:"properties,omitempty"`
	// The timestamp of the user's creation.
	// Timstamp facet can be used to set the creation and updation timestamp of a user.
	Timestamps *common.Timestamp `protobuf:"bytes,33,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	// The timestamp of the generated event.
	// Event schemas is defined in the common event schema.
	Event *common.Event `protobuf:"bytes,100,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_odpf_assets_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Dashboard) GetResource() *common.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Dashboard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dashboard) GetCharts() []*Chart {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *Dashboard) GetOwnership() *facets.Ownership {
	if x != nil {
		return x.Ownership
	}
	return nil
}

func (x *Dashboard) GetProperties() *facets.Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Dashboard) GetTimestamps() *common.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Dashboard) GetEvent() *common.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URN of the chart.
	// Example: `chart:1`.
	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	// The name of the chart.
	// Example: `My Chart`.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the chart.
	// Example: `line`.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// The source of the chart.
	// Example: `metabase`.
	Source string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// The description of the chart.
	// Example: `This is a chart for my dashboard.`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The url of the chart.
	// Example: `http://metabase.com/charts/mychart`.
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// The raw query of the chart.
	// Example: `SELECT * FROM my_table`.
	RawQuery string `protobuf:"bytes,7,opt,name=raw_query,json=rawQuery,proto3" json:"raw_query,omitempty"`
	// The source of the data.
	// Example: `bigquery,graphite`.
	DataSource string `protobuf:"bytes,8,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	// The dashboard ur of the chart.
	// Example: `dashboard:1`.
	DashboardUrn string `protobuf:"bytes,9,opt,name=dashboard_urn,json=dashboardUrn,proto3" json:"dashboard_urn,omitempty"`
	// The source of the dashboard of the chart.
	// Example: `metabase`.
	DashboardSource string `protobuf:"bytes,10,opt,name=dashboard_source,json=dashboardSource,proto3" json:"dashboard_source,omitempty"`
	// The ownership of the dashboard.
	// For an example check out ownership.
	Ownership *facets.Ownership `protobuf:"bytes,31,opt,name=ownership,proto3" json:"ownership,omitempty"`
	// The lineage of the dashboard.
	// For an example check out lineage schema.
	Lineage *facets.Lineage `protobuf:"bytes,32,opt,name=lineage,proto3" json:"lineage,omitempty"`
	// List of the user's custom properties.
	// Properties facet can be used to set custom properties, tags and labels for a dashboard.
	Properties *facets.Properties `protobuf:"bytes,33,opt,name=properties,proto3" json:"properties,omitempty"`
	// The timestamp of the user's creation.
	// Timstamp facet can be used to set the creation and updation timestamp of a dashboard.
	Timestamps *common.Timestamp `protobuf:"bytes,34,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	// The timestamp of the generated event.
	// Event schemas is defined in the common event schema.
	Event *common.Event `protobuf:"bytes,100,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Chart) Reset() {
	*x = Chart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_odpf_assets_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *Chart) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Chart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chart) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Chart) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Chart) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Chart) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Chart) GetRawQuery() string {
	if x != nil {
		return x.RawQuery
	}
	return ""
}

func (x *Chart) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

func (x *Chart) GetDashboardUrn() string {
	if x != nil {
		return x.DashboardUrn
	}
	return ""
}

func (x *Chart) GetDashboardSource() string {
	if x != nil {
		return x.DashboardSource
	}
	return ""
}

func (x *Chart) GetOwnership() *facets.Ownership {
	if x != nil {
		return x.Ownership
	}
	return nil
}

func (x *Chart) GetLineage() *facets.Lineage {
	if x != nil {
		return x.Lineage
	}
	return nil
}

func (x *Chart) GetProperties() *facets.Properties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Chart) GetTimestamps() *common.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Chart) GetEvent() *common.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_odpf_assets_dashboard_proto protoreflect.FileDescriptor

var file_odpf_assets_dashboard_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f,
	0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x1a, 0x22, 0x6f, 0x64, 0x70, 0x66,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x66, 0x61,
	0x63, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x64,
	0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a,
	0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f,
	0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x3e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x2f,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0xbf, 0x04, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x75, 0x72, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x55,
	0x72, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x66,
	0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x35, 0x0a, 0x07, 0x6c, 0x69,
	0x6e, 0x65, 0x61, 0x67, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64,
	0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x67,
	0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73,
	0x12, 0x2f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x3f, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x42, 0x0e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x64, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_assets_dashboard_proto_rawDescOnce sync.Once
	file_odpf_assets_dashboard_proto_rawDescData = file_odpf_assets_dashboard_proto_rawDesc
)

func file_odpf_assets_dashboard_proto_rawDescGZIP() []byte {
	file_odpf_assets_dashboard_proto_rawDescOnce.Do(func() {
		file_odpf_assets_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_assets_dashboard_proto_rawDescData)
	})
	return file_odpf_assets_dashboard_proto_rawDescData
}

var file_odpf_assets_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_odpf_assets_dashboard_proto_goTypes = []interface{}{
	(*Dashboard)(nil),         // 0: odpf.assets.Dashboard
	(*Chart)(nil),             // 1: odpf.assets.Chart
	(*common.Resource)(nil),   // 2: odpf.assets.common.Resource
	(*facets.Ownership)(nil),  // 3: odpf.assets.facets.Ownership
	(*facets.Properties)(nil), // 4: odpf.assets.facets.Properties
	(*common.Timestamp)(nil),  // 5: odpf.assets.common.Timestamp
	(*common.Event)(nil),      // 6: odpf.assets.common.Event
	(*facets.Lineage)(nil),    // 7: odpf.assets.facets.Lineage
}
var file_odpf_assets_dashboard_proto_depIdxs = []int32{
	2,  // 0: odpf.assets.Dashboard.resource:type_name -> odpf.assets.common.Resource
	1,  // 1: odpf.assets.Dashboard.charts:type_name -> odpf.assets.Chart
	3,  // 2: odpf.assets.Dashboard.ownership:type_name -> odpf.assets.facets.Ownership
	4,  // 3: odpf.assets.Dashboard.properties:type_name -> odpf.assets.facets.Properties
	5,  // 4: odpf.assets.Dashboard.timestamps:type_name -> odpf.assets.common.Timestamp
	6,  // 5: odpf.assets.Dashboard.event:type_name -> odpf.assets.common.Event
	3,  // 6: odpf.assets.Chart.ownership:type_name -> odpf.assets.facets.Ownership
	7,  // 7: odpf.assets.Chart.lineage:type_name -> odpf.assets.facets.Lineage
	4,  // 8: odpf.assets.Chart.properties:type_name -> odpf.assets.facets.Properties
	5,  // 9: odpf.assets.Chart.timestamps:type_name -> odpf.assets.common.Timestamp
	6,  // 10: odpf.assets.Chart.event:type_name -> odpf.assets.common.Event
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_odpf_assets_dashboard_proto_init() }
func file_odpf_assets_dashboard_proto_init() {
	if File_odpf_assets_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_assets_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
		file_odpf_assets_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chart); i {
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
			RawDescriptor: file_odpf_assets_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_assets_dashboard_proto_goTypes,
		DependencyIndexes: file_odpf_assets_dashboard_proto_depIdxs,
		MessageInfos:      file_odpf_assets_dashboard_proto_msgTypes,
	}.Build()
	File_odpf_assets_dashboard_proto = out.File
	file_odpf_assets_dashboard_proto_rawDesc = nil
	file_odpf_assets_dashboard_proto_goTypes = nil
	file_odpf_assets_dashboard_proto_depIdxs = nil
}
