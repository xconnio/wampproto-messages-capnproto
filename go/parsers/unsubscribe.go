package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Unsubscribe struct {
	gen *gen.Unsubscribe
}

func NewUnsubscribeFields(g *gen.Unsubscribe) messages.UnsubscribeFields {
	return &Unsubscribe{gen: g}
}

func (u *Unsubscribe) RequestID() uint64 {
	return u.gen.RequestID()
}

func (u *Unsubscribe) SubscriptionID() uint64 {
	return u.gen.SubscriptionID()
}

func UnsubscribeToCapnproto(m *messages.Unsubscribe) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unsubscribe, err := gen.NewRootUnsubscribe(seg)
	if err != nil {
		return nil, err
	}

	unsubscribe.SetRequestID(m.RequestID())
	unsubscribe.SetSubscriptionID(m.SubscriptionID())

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeUnsubscribe, data, nil), nil
}

func CapnprotoToUnsubscribe(data []byte) (*messages.Unsubscribe, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	unsubscribe, err := gen.ReadRootUnsubscribe(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnsubscribeWithFields(NewUnsubscribeFields(&unsubscribe)), nil
}
