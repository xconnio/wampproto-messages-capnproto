package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Call struct {
	gen *gen.Call
}

func NewCallFields(g *gen.Call) messages.CallFields {
	return &Call{gen: g}
}

func (c *Call) RequestID() int64 {
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

func (c *Call) PayloadSerializer() int {
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
	if err := call.SetProcedure(m.Procedure()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeCall)}, data.Bytes()...), nil
}

func CapnprotoToCall(data []byte) (*messages.Call, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	call, err := gen.ReadRootCall(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewCallWithFields(NewCallFields(&call)), nil
}
