// Code generated by capnpc-go. DO NOT EDIT.

package gen

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
)

type Published capnp.Struct

// Published_TypeID is the unique identifier for the type Published.
const Published_TypeID = 0x9bf8c21f751ebf99

func NewPublished(s *capnp.Segment) (Published, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Published(st), err
}

func NewRootPublished(s *capnp.Segment) (Published, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Published(st), err
}

func ReadRootPublished(msg *capnp.Message) (Published, error) {
	root, err := msg.Root()
	return Published(root.Struct()), err
}

func (s Published) String() string {
	str, _ := text.Marshal(0x9bf8c21f751ebf99, capnp.Struct(s))
	return str
}

func (s Published) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Published) DecodeFromPtr(p capnp.Ptr) Published {
	return Published(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Published) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Published) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Published) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Published) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Published) RequestID() int64 {
	return int64(capnp.Struct(s).Uint64(0))
}

func (s Published) SetRequestID(v int64) {
	capnp.Struct(s).SetUint64(0, uint64(v))
}

func (s Published) PublicationID() int64 {
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Published) SetPublicationID(v int64) {
	capnp.Struct(s).SetUint64(8, uint64(v))
}

// Published_List is a list of Published.
type Published_List = capnp.StructList[Published]

// NewPublished creates a new list of Published.
func NewPublished_List(s *capnp.Segment, sz int32) (Published_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Published](l), err
}

// Published_Future is a wrapper for a Published promised by a client call.
type Published_Future struct{ *capnp.Future }

func (f Published_Future) Struct() (Published, error) {
	p, err := f.Future.Ptr()
	return Published(p.Struct()), err
}
