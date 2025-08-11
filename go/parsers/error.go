package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-go/serializers"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Error struct {
	gen *gen.Error
	ex  *PayloadExpander
}

func NewErrorFields(g *gen.Error, payload []byte) messages.ErrorFields {
	return &Error{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (e *Error) MessageType() uint64 {
	return e.gen.MessageType()
}

func (e *Error) RequestID() uint64 {
	return e.gen.RequestID()
}

func (e *Error) URI() string {
	val, _ := e.gen.Uri()
	return val
}

func (e *Error) Details() map[string]any {
	return map[string]any{}
}

func (e *Error) Args() []any {
	return e.ex.Args()
}

func (e *Error) KwArgs() map[string]any {
	return e.ex.Kwargs()
}

func (e *Error) PayloadIsBinary() bool {
	return true
}

func (e *Error) Payload() []byte {
	return e.ex.Payload()
}

func (e *Error) PayloadSerializer() uint64 {
	return e.gen.PayloadSerializerID()
}

func ErrorToCapnproto(m *messages.Error) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	e, err := gen.NewRootError(seg)
	if err != nil {
		return nil, err
	}

	e.SetMessageType(m.MessageType())
	e.SetRequestID(m.RequestID())
	e.SetPayloadSerializerID(m.PayloadSerializer())
	if err = e.SetUri(m.URI()); err != nil {
		return nil, err
	}

	e.SetPayloadSerializerID(serializers.MsgPackSerializerID)

	payload, err := Encode(serializers.MsgPackSerializerID, m.Args(), m.KwArgs())
	if err != nil {
		return nil, err
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeError, data, payload), nil
}

func CapnprotoToError(data, payload []byte) (*messages.Error, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	e, err := gen.ReadRootError(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewErrorWithFields(NewErrorFields(&e, payload)), nil
}
