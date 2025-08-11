package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-go/serializers"
	"github.com/xconnio/wampproto-go/util"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Event struct {
	gen *gen.Event
	ex  *PayloadExpander
}

func NewEventFields(g *gen.Event, payload []byte) messages.EventFields {
	return &Event{
		gen: g,
		ex:  &PayloadExpander{payload: payload, serializer: g.PayloadSerializerID()},
	}
}

func (e *Event) SubscriptionID() uint64 {
	return e.gen.SubscriptionID()
}

func (e *Event) PublicationID() uint64 {
	return e.gen.PublicationID()
}

func setDetail(details *map[string]any, key string, value any) {
	if *details == nil {
		*details = make(map[string]any)
	}

	(*details)[key] = value
}

func (e *Event) Details() map[string]any {
	var details map[string]any

	if e.gen.Publisher() > 0 {
		setDetail(&details, "publisher", e.gen.Publisher())
	}

	if e.gen.HasPublisherAuthID() {
		authID, _ := e.gen.PublisherAuthID()
		setDetail(&details, "publisher_authid", authID)
	}

	if e.gen.HasPublisherAuthRole() {
		authRole, _ := e.gen.PublisherAuthRole()
		setDetail(&details, "publisher_authrole", authRole)
	}

	if e.gen.HasTopic() {
		topic, _ := e.gen.Topic()
		setDetail(&details, "topic", topic)
	}

	return details
}

func (e *Event) Args() []any {
	return e.ex.Args()
}

func (e *Event) KwArgs() map[string]any {
	return e.ex.Kwargs()
}

func (e *Event) PayloadIsBinary() bool {
	return true
}

func (e *Event) Payload() []byte {
	return e.ex.Payload()
}

func (e *Event) PayloadSerializer() uint64 {
	return e.gen.PayloadSerializerID()
}

func EventToCapnproto(m *messages.Event) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	event, err := gen.NewRootEvent(seg)
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

	payloadSerializer := selectPayloadSerializer(m.Details())
	event.SetPayloadSerializerID(payloadSerializer)
	payload, err := serializers.SerializePayload(payloadSerializer, m.Args(), m.KwArgs())
	if err != nil {
		return nil, err
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeEvent, data, payload), nil
}

func CapnprotoToEvent(data, payload []byte) (*messages.Event, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	event, err := gen.ReadRootEvent(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewEventWithFields(NewEventFields(&event, payload)), nil
}
