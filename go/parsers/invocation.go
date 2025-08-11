package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-go/serializers"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Invocation struct {
	gen *gen.Invocation
	ex  *PayloadExpander
}

func NewInvocationFields(g *gen.Invocation, payload []byte) messages.InvocationFields {
	return &Invocation{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (i *Invocation) RequestID() uint64 {
	return i.gen.RequestID()
}

func (i *Invocation) RegistrationID() uint64 {
	return i.gen.RegistrationID()
}

func (i *Invocation) Args() []any {
	return i.ex.Args()
}

func (i *Invocation) KwArgs() map[string]any {
	return i.ex.Kwargs()
}

func (i *Invocation) Details() map[string]any {
	var details map[string]any

	if i.gen.Caller() > 0 {
		setDetail(&details, "caller", i.gen.Caller())
	}

	if i.gen.HasCallerAuthID() {
		authID, _ := i.gen.CallerAuthID()
		setDetail(&details, "caller_authid", authID)
	}

	if i.gen.HasCallerAuthRole() {
		authRole, _ := i.gen.CallerAuthRole()
		setDetail(&details, "caller_authrole", authRole)
	}

	if i.gen.HasProcedure() {
		topic, _ := i.gen.Procedure()
		setDetail(&details, "procedure", topic)
	}

	return details
}

func (i *Invocation) PayloadIsBinary() bool {
	return true
}

func (i *Invocation) Payload() []byte {
	return i.ex.Payload()
}

func (i *Invocation) PayloadSerializer() uint64 {
	return i.gen.PayloadSerializerID()
}

func InvocationToCapnproto(m *messages.Invocation) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	invocation, err := gen.NewRootInvocation(seg)
	if err != nil {
		return nil, err
	}

	invocation.SetRequestID(m.RequestID())
	invocation.SetRegistrationID(m.RegistrationID())
	invocation.SetPayloadSerializerID(serializers.MsgPackSerializerID)

	payload, err := Encode(serializers.MsgPackSerializerID, m.Args(), m.KwArgs())
	if err != nil {
		return nil, err
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeInvocation, data, payload), nil
}

func CapnprotoToInvocation(data, payload []byte) (*messages.Invocation, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	invocation, err := gen.ReadRootInvocation(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewInvocationWithFields(NewInvocationFields(&invocation, payload)), nil
}
