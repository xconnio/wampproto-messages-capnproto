package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Authenticate struct {
	gen *gen.Authenticate
}

func NewAuthenticateFields(g *gen.Authenticate) messages.AuthenticateFields {
	return &Authenticate{gen: g}
}

func (a *Authenticate) Signature() string {
	val, _ := a.gen.Signature()
	return val
}

func (a *Authenticate) Extra() map[string]any {
	return map[string]any{}
}

func AuthenticateToCapnproto(m *messages.Authenticate) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	auth, err := gen.NewRootAuthenticate(seg)
	if err != nil {
		return nil, err
	}

	if err = auth.SetSignature(m.Signature()); err != nil {
		return nil, err
	}

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeAuthenticate, data), nil
}

func CapnprotoToAuthenticate(data []byte) (*messages.Authenticate, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	auth, err := gen.ReadRootAuthenticate(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewAuthenticateWithFields(NewAuthenticateFields(&auth)), nil
}
