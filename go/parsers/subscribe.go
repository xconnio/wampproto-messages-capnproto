package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Subscribe struct {
	gen *gen.Subscribe
}

func NewSubscribeFields(g *gen.Subscribe) messages.SubscribeFields {
	return &Subscribe{gen: g}
}

func (s *Subscribe) RequestID() uint64 {
	return s.gen.RequestID()
}

func (s *Subscribe) Options() map[string]any {
	return map[string]any{}
}

func (s *Subscribe) Topic() string {
	topic, _ := s.gen.Topic()
	return topic
}

func SubscribeToCapnproto(m *messages.Subscribe) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	subscribe, err := gen.NewRootSubscribe(seg)
	if err != nil {
		return nil, err
	}

	subscribe.SetRequestID(m.RequestID())
	if err := subscribe.SetTopic(m.Topic()); err != nil {
		return nil, err
	}

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeSubscribe, data), nil
}

func CapnprotoToSubscribe(data []byte) (*messages.Subscribe, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	subscribe, err := gen.ReadRootSubscribe(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewSubscribeWithFields(NewSubscribeFields(&subscribe)), nil
}
