package parsers

import (
	"capnproto.org/go/capnp/v3"
	"github.com/xconnio/wampproto-go/serializers"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Yield struct {
	gen *gen.Yield
	ex  *PayloadExpander
}

func NewYieldFields(g *gen.Yield, payload []byte) messages.YieldFields {
	return &Yield{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (y *Yield) RequestID() uint64 {
	return y.gen.RequestID()
}

func (y *Yield) Options() map[string]any {
	return map[string]any{}
}

func (y *Yield) Args() []any {
	return y.ex.Args()
}

func (y *Yield) KwArgs() map[string]any {
	return y.ex.Kwargs()
}

func (y *Yield) PayloadIsBinary() bool {
	return true
}

func (y *Yield) Payload() []byte {
	return y.ex.Payload()
}

func (y *Yield) PayloadSerializer() uint64 {
	return y.gen.PayloadSerializerID()
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
	payloadSerializer := selectPayloadSerializer(m.Options())
	yield.SetPayloadSerializerID(payloadSerializer)

	var payload []byte
	if m.PayloadIsBinary() {
		payload = m.Payload()
	} else {
		payload, err = serializers.SerializePayload(payloadSerializer, m.Args(), m.KwArgs())
		if err != nil {
			return nil, err
		}
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeYield, data, payload), nil
}

func CapnprotoToYield(data, payload []byte) (*messages.Yield, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	yield, err := gen.ReadRootYield(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewYieldWithFields(NewYieldFields(&yield, payload)), nil
}
