package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Subscribed struct {
	gen *gen.Subscribed
}

func NewSubscribedFields(g *gen.Subscribed) messages.SubscribedFields {
	return &Subscribed{gen: g}
}

func (s *Subscribed) RequestID() uint64 {
	return s.gen.RequestID()
}

func (s *Subscribed) SubscriptionID() uint64 {
	return s.gen.SubscriptionID()
}

func SubscribedToCapnproto(m *messages.Subscribed) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	subscribed, err := gen.NewRootSubscribed(seg)
	if err != nil {
		return nil, err
	}

	subscribed.SetRequestID(m.RequestID())
	subscribed.SetSubscriptionID(m.SubscriptionID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeSubscribed, data, nil), nil
}

func CapnprotoToSubscribed(data []byte) (*messages.Subscribed, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	subscribed, err := gen.ReadRootSubscribed(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewSubscribedWithFields(NewSubscribedFields(&subscribed)), nil
}
