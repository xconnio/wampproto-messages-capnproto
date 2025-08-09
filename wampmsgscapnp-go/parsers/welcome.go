package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Welcome struct {
	gen *gen.Welcome
}

func NewWelcomeFields(g *gen.Welcome) messages.WelcomeFields {
	return &Welcome{gen: g}
}

func (w *Welcome) SessionID() uint64 {
	return w.gen.SessionID()
}

func (w *Welcome) Details() map[string]any {
	return map[string]any{}
}

func WelcomeToCapnproto(w *messages.Welcome) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	welcome, err := gen.NewWelcome(seg)
	if err != nil {
		return nil, err
	}

	welcome.SetSessionID(w.SessionID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeWelcome, &data), nil
}

func CapnprotoToWelcome(data []byte) (*messages.Welcome, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	welcome, err := gen.ReadRootWelcome(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewWelcomeWithFields(NewWelcomeFields(&welcome)), nil
}
