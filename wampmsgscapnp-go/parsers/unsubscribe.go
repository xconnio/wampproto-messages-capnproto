package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Unsubscribe struct {
	gen *gen.Unsubscribe
}

func NewUnsubscribeFields(g *gen.Unsubscribe) messages.UnsubscribeFields {
	return &Unsubscribe{gen: g}
}

func (u *Unsubscribe) RequestID() int64 {
	return u.gen.RequestID()
}

func (u *Unsubscribe) SubscriptionID() int64 {
	return u.gen.SubscriptionID()
}

func UnsubscribeToCapnproto(m *messages.Unsubscribe) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	unsubscribe, err := gen.NewUnsubscribe(seg)
	if err != nil {
		return nil, err
	}

	unsubscribe.SetRequestID(m.RequestID())
	unsubscribe.SetSubscriptionID(m.SubscriptionID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeUnsubscribe)}, data.Bytes()...), nil
}

func CapnprotoToUnsubscribe(data []byte) (*messages.Unsubscribe, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	unsubscribe, err := gen.ReadRootUnsubscribe(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewUnsubscribeWithFields(NewUnsubscribeFields(&unsubscribe)), nil
}
