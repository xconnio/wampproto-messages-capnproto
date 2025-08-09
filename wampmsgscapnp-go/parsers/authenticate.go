package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
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

	auth, err := gen.NewAuthenticate(seg)
	if err != nil {
		return nil, err
	}

	if err = auth.SetSignature(m.Signature()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeAuthenticate, &data), nil
}

func CapnprotoToAuthenticate(data []byte) (*messages.Authenticate, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	auth, err := gen.ReadRootAuthenticate(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewAuthenticateWithFields(NewAuthenticateFields(&auth)), nil
}
