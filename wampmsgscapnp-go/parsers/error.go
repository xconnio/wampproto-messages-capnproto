package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Error struct {
	gen *gen.Error
}

func NewErrorFields(g *gen.Error) messages.ErrorFields {
	return &Error{gen: g}
}

func (e *Error) MessageType() int64 {
	return e.gen.MessageType()
}

func (e *Error) RequestID() int64 {
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
	return nil
}

func (e *Error) KwArgs() map[string]any {
	return nil
}

func (e *Error) PayloadIsBinary() bool {
	return true
}

func (e *Error) Payload() []byte {
	return nil
}

func (e *Error) PayloadSerializer() int {
	return 0
}

func ErrorToCapnproto(m *messages.Error) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	e, err := gen.NewError(seg)
	if err != nil {
		return nil, err
	}

	e.SetMessageType(m.MessageType())
	e.SetRequestID(m.RequestID())
	if err := e.SetUri(m.URI()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeError)}, data.Bytes()...), nil
}

func CapnprotoToError(data []byte) (*messages.Error, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	e, err := gen.ReadRootError(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewErrorWithFields(NewErrorFields(&e)), nil
}
