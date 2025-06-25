package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Invocation struct {
	gen *gen.Invocation
}

func NewInvocationFields(g *gen.Invocation) messages.InvocationFields {
	return &Invocation{gen: g}
}

func (i *Invocation) RequestID() int64 {
	return i.gen.RequestID()
}

func (i *Invocation) RegistrationID() int64 {
	return i.gen.RegistrationID()
}

func (i *Invocation) Args() []any {
	return nil
}

func (i *Invocation) KwArgs() map[string]any {
	return nil
}

func (i *Invocation) Details() map[string]any {
	return map[string]any{}
}

func (i *Invocation) PayloadIsBinary() bool {
	return true
}

func (i *Invocation) Payload() []byte {
	return nil
}

func (i *Invocation) PayloadSerializer() int {
	return 0
}

func InvocationToCapnproto(m *messages.Invocation) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	invocation, err := gen.NewInvocation(seg)
	if err != nil {
		return nil, err
	}

	invocation.SetRequestID(m.RequestID())
	invocation.SetRegistrationID(m.RegistrationID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeInvocation)}, data.Bytes()...), nil
}

func CapnprotoToInvocation(data []byte) (*messages.Invocation, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	invocation, err := gen.ReadRootInvocation(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewInvocationWithFields(NewInvocationFields(&invocation)), nil
}
