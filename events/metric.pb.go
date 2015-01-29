// Code generated by protoc-gen-gogo.
// source: metric.proto
// DO NOT EDIT!

package events

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// / A ValueMetric indicates the value of a metric at an instant in time.
type ValueMetric struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,req,name=value" json:"value,omitempty"`
	Unit             *string  `protobuf:"bytes,3,req,name=unit" json:"unit,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ValueMetric) Reset()         { *m = ValueMetric{} }
func (m *ValueMetric) String() string { return proto.CompactTextString(m) }
func (*ValueMetric) ProtoMessage()    {}

func (m *ValueMetric) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ValueMetric) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *ValueMetric) GetUnit() string {
	if m != nil && m.Unit != nil {
		return *m.Unit
	}
	return ""
}

// / A CounterEvent represents the increment of a counter. It contains only the change in the value; it is the responsibility of downstream consumers to maintain the value of the counter.
type CounterEvent struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Delta            *uint64 `protobuf:"varint,2,req,name=delta" json:"delta,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CounterEvent) Reset()         { *m = CounterEvent{} }
func (m *CounterEvent) String() string { return proto.CompactTextString(m) }
func (*CounterEvent) ProtoMessage()    {}

func (m *CounterEvent) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CounterEvent) GetDelta() uint64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

// / A ContainerMetric records resource usage of an app in a container.
type ContainerMetric struct {
	ApplicationId    *string  `protobuf:"bytes,1,req,name=applicationId" json:"applicationId,omitempty"`
	InstanceIndex    *int32   `protobuf:"varint,2,req,name=instanceIndex" json:"instanceIndex,omitempty"`
	CpuPercentage    *float64 `protobuf:"fixed64,3,req,name=cpuPercentage" json:"cpuPercentage,omitempty"`
	MemoryBytes      *uint64  `protobuf:"varint,4,req,name=memoryBytes" json:"memoryBytes,omitempty"`
	DiskBytes        *uint64  `protobuf:"varint,5,req,name=diskBytes" json:"diskBytes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ContainerMetric) Reset()         { *m = ContainerMetric{} }
func (m *ContainerMetric) String() string { return proto.CompactTextString(m) }
func (*ContainerMetric) ProtoMessage()    {}

func (m *ContainerMetric) GetApplicationId() string {
	if m != nil && m.ApplicationId != nil {
		return *m.ApplicationId
	}
	return ""
}

func (m *ContainerMetric) GetInstanceIndex() int32 {
	if m != nil && m.InstanceIndex != nil {
		return *m.InstanceIndex
	}
	return 0
}

func (m *ContainerMetric) GetCpuPercentage() float64 {
	if m != nil && m.CpuPercentage != nil {
		return *m.CpuPercentage
	}
	return 0
}

func (m *ContainerMetric) GetMemoryBytes() uint64 {
	if m != nil && m.MemoryBytes != nil {
		return *m.MemoryBytes
	}
	return 0
}

func (m *ContainerMetric) GetDiskBytes() uint64 {
	if m != nil && m.DiskBytes != nil {
		return *m.DiskBytes
	}
	return 0
}

func init() {
}
