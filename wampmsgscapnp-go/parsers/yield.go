package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Yield struct {
	gen *gen.Yield
}

func NewYieldFields(g *gen.Yield) messages.YieldFields {
	return &Yield{gen: g}
}

func (y *Yield) RequestID() int64 {
	return y.gen.RequestID()
}

func (y *Yield) Options() map[string]any {
	return map[string]any{}
}

func (y *Yield) Args() []any {
	return nil
}

func (y *Yield) KwArgs() map[string]any {
	return nil
}

func (y *Yield) PayloadIsBinary() bool {
	return true
}

func (y *Yield) Payload() []byte {
	return nil
}

func (y *Yield) PayloadSerializer() int {
	return 0
}

func YieldToCapnproto(m *messages.Yield) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	yield, err := gen.NewYield(seg)
	if err != nil {
		return nil, err
	}

	yield.SetRequestID(m.RequestID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeYield)}, data.Bytes()...), nil
}

func CapnprotoToYield(data []byte) (*messages.Yield, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	yield, err := gen.ReadRootYield(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewYieldWithFields(NewYieldFields(&yield)), nil
}
