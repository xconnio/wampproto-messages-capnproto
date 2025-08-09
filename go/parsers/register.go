package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Register struct {
	gen *gen.Register
}

func NewRegisterFields(g *gen.Register) messages.RegisterFields {
	return &Register{gen: g}
}

func (r *Register) RequestID() uint64 {
	return r.gen.RequestID()
}

func (r *Register) Options() map[string]any {
	return map[string]any{}
}

func (r *Register) Procedure() string {
	proc, _ := r.gen.Procedure()
	return proc
}

func RegisterToCapnproto(m *messages.Register) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	register, err := gen.NewRegister(seg)
	if err != nil {
		return nil, err
	}

	register.SetRequestID(m.RequestID())

	if err := register.SetProcedure(m.Procedure()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeRegister, &data), nil
}

func CapnprotoToRegister(data []byte) (*messages.Register, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	register, err := gen.ReadRootRegister(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewRegisterWithFields(NewRegisterFields(&register)), nil
}
