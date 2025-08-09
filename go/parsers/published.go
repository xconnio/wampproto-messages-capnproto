package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Published struct {
	gen *gen.Published
}

func NewPublishedFields(gen *gen.Published) messages.PublishedFields {
	return &Published{gen: gen}
}

func (p *Published) RequestID() uint64 {
	return p.gen.RequestID()
}

func (p *Published) PublicationID() uint64 {
	return p.gen.PublicationID()
}

func PublishedToCapnproto(published *messages.Published) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	pubed, err := gen.NewRootPublished(seg)
	if err != nil {
		return nil, err
	}

	pubed.SetRequestID(published.RequestID())
	pubed.SetPublicationID(published.PublicationID())

	data, err := msg.Marshal()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypePublished, data), nil
}

func CapnprotoToPublished(data []byte) (*messages.Published, error) {
	msg, err := capnp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	published, err := gen.ReadRootPublished(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewPublishedWithFields(NewPublishedFields(&published)), nil
}
