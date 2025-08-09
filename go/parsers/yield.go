package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Yield struct {
	gen     *gen.Yield
	payload []byte
}

func NewYieldFields(g *gen.Yield, payload []byte) messages.YieldFields {
	return &Yield{gen: g, payload: payload}
}

func (y *Yield) RequestID() uint64 {
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

func (y *Yield) PayloadSerializer() uint64 {
	return 0
}

func YieldToCapnproto(m *messages.Yield) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	yield, err := gen.NewRootYield(seg)
	if err != nil {
		return nil, err
	}

	yield.SetRequestID(m.RequestID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeYield, data), nil
}

func CapnprotoToYield(data, payload []byte) (*messages.Yield, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	yield, err := gen.ReadRootYield(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewYieldWithFields(NewYieldFields(&yield, payload)), nil
}
