package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type GoodBye struct {
	gen *gen.Goodbye
}

func NewGoodByeFields(g *gen.Goodbye) messages.GoodByeFields {
	return &GoodBye{gen: g}
}

func (g *GoodBye) Reason() string {
	reason, _ := g.gen.Reason()
	return reason
}

func (g *GoodBye) Details() map[string]any {
	return map[string]any{}
}

func GoodbyeToCapnproto(m *messages.GoodBye) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	goodbye, err := gen.NewGoodbye(seg)
	if err != nil {
		return nil, err
	}

	if err := goodbye.SetReason(m.Reason()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeGoodbye)}, data.Bytes()...), nil
}

func CapnprotoToGoodbye(data []byte) (*messages.GoodBye, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	goodbye, err := gen.ReadRootGoodbye(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewGoodByeWithFields(NewGoodByeFields(&goodbye)), nil
}
