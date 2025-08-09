package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"
	"github.com/xconnio/wampproto-go/util"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Event struct {
	gen     *gen.Event
	payload []byte
}

func NewEventFields(g *gen.Event, payload []byte) messages.EventFields {
	return &Event{gen: g, payload: payload}
}

func (e *Event) SubscriptionID() uint64 {
	return e.gen.SubscriptionID()
}

func (e *Event) PublicationID() uint64 {
	return e.gen.PublicationID()
}

func (e *Event) Details() map[string]any {
	return map[string]any{}
}

func (e *Event) Args() []any {
	return nil
}

func (e *Event) KwArgs() map[string]any {
	return nil
}

func (e *Event) PayloadIsBinary() bool {
	return true
}

func (e *Event) Payload() []byte {
	return nil
}

func (e *Event) PayloadSerializer() uint64 {
	return 0
}

func EventToCapnproto(m *messages.Event) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	event, err := gen.NewEvent(seg)
	if err != nil {
		return nil, err
	}

	event.SetSubscriptionID(m.SubscriptionID())
	event.SetPublicationID(m.PublicationID())
	event.SetPayloadSerializerID(m.PayloadSerializer())

	if publisher, ok := util.AsUInt64(m.Details()["publisher"]); ok {
		event.SetPublicationID(publisher)

		authID, ok := util.AsString(m.Details()["publisher_authid"])
		if err = event.SetPublisherAuthID(authID); err != nil && ok {
			return nil, err
		}

		authRole, ok := util.AsString(m.Details()["publisher_authrole"])
		if err = event.SetPublisherAuthRole(authRole); err != nil && ok {
			return nil, err
		}

		topic, ok := util.AsString(m.Details()["topic"])
		if err = event.SetTopic(topic); err != nil && ok {
			return nil, err
		}
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeEvent, &data), nil
}

func CapnprotoToEvent(data, payload []byte) (*messages.Event, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	event, err := gen.ReadRootEvent(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewEventWithFields(NewEventFields(&event, payload)), nil
}
