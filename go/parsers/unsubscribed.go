package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Unsubscribed struct {
	gen *gen.Unsubscribed
}

func NewUnsubscribedFields(g *gen.Unsubscribed) messages.UnsubscribedFields {
	return &Unsubscribed{gen: g}
}

func (u *Unsubscribed) RequestID() uint64 {
	return u.gen.RequestID()
}

func UnsubscribedToCapnproto(m *messages.Unsubscribed) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unsubscribed, err := gen.NewRootUnsubscribed(seg)
	if err != nil {
		return nil, err
	}

	unsubscribed.SetRequestID(m.RequestID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeUnsubscribed, data), nil
}

func CapnprotoToUnsubscribed(data []byte) (*messages.Unsubscribed, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	unsubscribed, err := gen.ReadRootUnsubscribed(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnsubscribedWithFields(NewUnsubscribedFields(&unsubscribed)), nil
}
