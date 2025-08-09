package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Abort struct {
	gen *gen.Abort
	ex  *PayloadExpander
}

func NewAbortFields(g *gen.Abort, payload []byte) messages.AbortFields {
	return &Abort{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (a *Abort) Reason() string {
	reason, _ := a.gen.Reason()
	return reason
}

func (a *Abort) Details() map[string]any {
	return map[string]any{}
}

func (a *Abort) Args() []any {
	return a.ex.Args()
}

func (a *Abort) KwArgs() map[string]any {
	return a.ex.Kwargs()
}

func AbortToCapnproto(m *messages.Abort) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	abort, err := gen.NewRootAbort(seg)
	if err != nil {
		return nil, err
	}

	if err = abort.SetReason(m.Reason()); err != nil {
		return nil, err
	}

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeAbort, data), nil
}

func CapnprotoToAbort(data, payload []byte) (*messages.Abort, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	abort, err := gen.ReadRootAbort(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewAbortWithFields(NewAbortFields(&abort, payload)), nil
}
