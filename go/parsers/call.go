package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Call struct {
	gen     *gen.Call
	payload []byte
}

func NewCallFields(g *gen.Call, payload []byte) messages.CallFields {
	return &Call{gen: g, payload: payload}
}

func (c *Call) RequestID() uint64 {
	return c.gen.RequestID()
}

func (c *Call) Procedure() string {
	proc, _ := c.gen.Procedure()
	return proc
}

func (c *Call) Args() []any {
	return nil
}

func (c *Call) KwArgs() map[string]any {
	return nil
}

func (c *Call) Options() map[string]any {
	return map[string]any{}
}

func (c *Call) PayloadIsBinary() bool {
	return true
}

func (c *Call) Payload() []byte {
	return nil
}

func (c *Call) PayloadSerializer() uint64 {
	return 0
}

func CallToCapnproto(m *messages.Call) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	call, err := gen.NewCall(seg)
	if err != nil {
		return nil, err
	}

	call.SetRequestID(m.RequestID())
	call.SetPayloadSerializerID(m.PayloadSerializer())
	if err = call.SetProcedure(m.Procedure()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err = capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeCall, &data), nil
}

func CapnprotoToCall(data, payload []byte) (*messages.Call, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	call, err := gen.ReadRootCall(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewCallWithFields(NewCallFields(&call, payload)), nil
}
