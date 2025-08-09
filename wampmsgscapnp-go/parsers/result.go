package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Result struct {
	gen     *gen.Result
	payload []byte
}

func NewResultFields(g *gen.Result, payload []byte) messages.ResultFields {
	return &Result{gen: g, payload: payload}
}

func (r *Result) RequestID() uint64 {
	return r.gen.RequestID()
}

func (r *Result) Details() map[string]any {
	return map[string]any{}
}

func (r *Result) Args() []any {
	return nil
}

func (r *Result) KwArgs() map[string]any {
	return nil
}

func (r *Result) PayloadIsBinary() bool {
	return true
}

func (r *Result) Payload() []byte {
	return nil
}

func (r *Result) PayloadSerializer() uint64 {
	return 0
}

func ResultToCapnproto(m *messages.Result) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	result, err := gen.NewResult(seg)
	if err != nil {
		return nil, err
	}

	result.SetRequestID(m.RequestID())

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeResult, &data), nil
}

func CapnprotoToResult(data, payload []byte) (*messages.Result, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	result, err := gen.ReadRootResult(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewResultWithFields(NewResultFields(&result, payload)), nil
}
