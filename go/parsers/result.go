package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Result struct {
	gen *gen.Result
	ex  *PayloadExpander
}

func NewResultFields(g *gen.Result, payload []byte) messages.ResultFields {
	return &Result{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (r *Result) RequestID() uint64 {
	return r.gen.RequestID()
}

func (r *Result) Details() map[string]any {
	return map[string]any{}
}

func (r *Result) Args() []any {
	return r.ex.Args()
}

func (r *Result) KwArgs() map[string]any {
	return r.ex.Kwargs()
}

func (r *Result) PayloadIsBinary() bool {
	return true
}

func (r *Result) Payload() []byte {
	return r.ex.Payload()
}

func (r *Result) PayloadSerializer() uint64 {
	return r.gen.PayloadSerializerID()
}

func ResultToCapnproto(m *messages.Result) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	result, err := gen.NewRootResult(seg)
	if err != nil {
		return nil, err
	}

	result.SetRequestID(m.RequestID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeResult, data), nil
}

func CapnprotoToResult(data, payload []byte) (*messages.Result, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	result, err := gen.ReadRootResult(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewResultWithFields(NewResultFields(&result, payload)), nil
}
