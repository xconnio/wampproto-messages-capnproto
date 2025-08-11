package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-go/serializers"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Call struct {
	gen *gen.Call
	ex  *PayloadExpander
}

func NewCallFields(g *gen.Call, payload []byte) messages.CallFields {
	return &Call{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (c *Call) RequestID() uint64 {
	return c.gen.RequestID()
}

func (c *Call) Procedure() string {
	proc, _ := c.gen.Procedure()
	return proc
}

func (c *Call) Args() []any {
	return c.ex.Args()
}

func (c *Call) KwArgs() map[string]any {
	return c.ex.Kwargs()
}

func (c *Call) Options() map[string]any {
	return map[string]any{}
}

func (c *Call) PayloadIsBinary() bool {
	return true
}

func (c *Call) Payload() []byte {
	return c.ex.Payload()
}

func (c *Call) PayloadSerializer() uint64 {
	return 0
}

func CallToCapnproto(m *messages.Call) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	call, err := gen.NewRootCall(seg)
	if err != nil {
		return nil, err
	}

	call.SetRequestID(m.RequestID())
	if err = call.SetProcedure(m.Procedure()); err != nil {
		return nil, err
	}

	payloadSerializer := selectPayloadSerializer(m.Options())
	call.SetPayloadSerializerID(payloadSerializer)
	payload, err := serializers.SerializePayload(payloadSerializer, m.Args(), m.KwArgs())
	if err != nil {
		return nil, err
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeCall, data, payload), nil
}

func CapnprotoToCall(data, payload []byte) (*messages.Call, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	call, err := gen.ReadRootCall(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewCallWithFields(NewCallFields(&call, payload)), nil
}
