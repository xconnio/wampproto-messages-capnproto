package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Abort struct {
	gen     *gen.Abort
	payload []byte
}

func NewAbortFields(g *gen.Abort, payload []byte) messages.AbortFields {
	return &Abort{gen: g, payload: payload}
}

func (a *Abort) Reason() string {
	reason, _ := a.gen.Reason()
	return reason
}

func (a *Abort) Details() map[string]any {
	return map[string]any{}
}

func (a *Abort) Args() []any {
	return nil
}

func (a *Abort) KwArgs() map[string]any {
	return nil
}

func AbortToCapnproto(m *messages.Abort) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	abort, err := gen.NewAbort(seg)
	if err != nil {
		return nil, err
	}

	if err := abort.SetReason(m.Reason()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeAbort, &data), nil
}

func CapnprotoToAbort(data, payload []byte) (*messages.Abort, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	abort, err := gen.ReadRootAbort(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewAbortWithFields(NewAbortFields(&abort, payload)), nil
}
