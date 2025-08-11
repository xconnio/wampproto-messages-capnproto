package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
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
	authID, _ := w.gen.Authid()
	authRole, _ := w.gen.Authrole()
	authMethod, _ := w.gen.Authmethod()
	authProvider, _ := w.gen.Authprovider()

	return map[string]any{
		"authid":       authID,
		"authrole":     authRole,
		"authmethod":   authMethod,
		"authprovider": authProvider,
	}
}

func WelcomeToCapnproto(w *messages.Welcome) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	welcome, err := gen.NewRootWelcome(seg)
	if err != nil {
		return nil, err
	}

	welcome.SetSessionID(w.SessionID())

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeWelcome, data, nil), nil
}

func CapnprotoToWelcome(data []byte) (*messages.Welcome, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	welcome, err := gen.ReadRootWelcome(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewWelcomeWithFields(NewWelcomeFields(&welcome)), nil
}
