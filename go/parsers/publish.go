package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Publish struct {
	gen     *gen.Publish
	payload []byte
}

func NewPublishFields(g *gen.Publish, payload []byte) messages.PublishFields {
	return &Publish{gen: g, payload: payload}
}

func (p *Publish) RequestID() uint64 {
	return p.gen.RequestID()
}

func (p *Publish) Topic() string {
	topic, _ := p.gen.Topic()
	return topic
}

func (p *Publish) Options() map[string]any {
	return map[string]any{}
}

func (p *Publish) Args() []any {
	return nil
}

func (p *Publish) KwArgs() map[string]any {
	return nil
}

func (p *Publish) PayloadIsBinary() bool {
	return true
}

func (p *Publish) Payload() []byte {
	return nil
}

func (p *Publish) PayloadSerializer() uint64 {
	return 0
}

func PublishToCapnproto(m *messages.Publish) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	publish, err := gen.NewRootPublish(seg)
	if err != nil {
		return nil, err
	}

	publish.SetRequestID(m.RequestID())
	if err := publish.SetTopic(m.Topic()); err != nil {
		return nil, err
	}

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypePublish, data), nil
}

func CapnprotoToPublish(data, payload []byte) (*messages.Publish, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	publish, err := gen.ReadRootPublish(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewPublishWithFields(NewPublishFields(&publish, payload)), nil
}
