package main

import (
	"github.com/golang/protobuf/proto"
)

type Request struct {
	Hostname         *string    `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	ServerName       *string    `protobuf:"bytes,2,req,name=server_name" json:"server_name,omitempty"`
	ScriptName       *string    `protobuf:"bytes,3,req,name=script_name" json:"script_name,omitempty"`
	RequestCount     *uint32    `protobuf:"varint,4,req,name=request_count" json:"request_count,omitempty"`
	DocumentSize     *uint32    `protobuf:"varint,5,req,name=document_size" json:"document_size,omitempty"`
	MemoryPeak       *uint32    `protobuf:"varint,6,req,name=memory_peak" json:"memory_peak,omitempty"`
	RequestTime      *float32   `protobuf:"fixed32,7,req,name=request_time" json:"request_time,omitempty"`
	RuUtime          *float32   `protobuf:"fixed32,8,req,name=ru_utime" json:"ru_utime,omitempty"`
	RuStime          *float32   `protobuf:"fixed32,9,req,name=ru_stime" json:"ru_stime,omitempty"`
	TimerHitCount    []uint32   `protobuf:"varint,10,rep,name=timer_hit_count" json:"timer_hit_count,omitempty"`
	TimerValue       []float32  `protobuf:"fixed32,11,rep,name=timer_value" json:"timer_value,omitempty"`
	TimerTagCount    []uint32   `protobuf:"varint,12,rep,name=timer_tag_count" json:"timer_tag_count,omitempty"`
	TimerTagName     []uint32   `protobuf:"varint,13,rep,name=timer_tag_name" json:"timer_tag_name,omitempty"`
	TimerTagValue    []uint32   `protobuf:"varint,14,rep,name=timer_tag_value" json:"timer_tag_value,omitempty"`
	Dictionary       []string   `protobuf:"bytes,15,rep,name=dictionary" json:"dictionary,omitempty"`
	Status           *uint32    `protobuf:"varint,16,opt,name=status" json:"status,omitempty"`
	MemoryFootprint  *uint32    `protobuf:"varint,17,opt,name=memory_footprint" json:"memory_footprint,omitempty"`
	Requests         []*Request `protobuf:"bytes,18,rep,name=requests" json:"requests,omitempty"`
	Schema           *string    `protobuf:"bytes,19,opt,name=schema" json:"schema,omitempty"`
	TagName          []uint32   `protobuf:"varint,20,rep,name=tag_name" json:"tag_name,omitempty"`
	TagValue         []uint32   `protobuf:"varint,21,rep,name=tag_value" json:"tag_value,omitempty"`
	TimerUtime       []float32  `protobuf:"fixed32,22,rep,name=timer_ru_utime" json:"timer_ru_utime,omitempty"`
	TimerStime       []float32  `protobuf:"fixed32,23,rep,name=timer_ru_stime" json:"timer_ru_stime,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Request) Reset() {
	*m = Request{}
}

func (m *Request) String() string {
	return proto.CompactTextString(m)
}

func (*Request) ProtoMessage() {

}
