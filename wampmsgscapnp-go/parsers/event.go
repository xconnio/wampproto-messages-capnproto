package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Event struct {
	gen *gen.Event
}

func NewEventFields(g *gen.Event) messages.EventFields {
	return &Event{gen: g}
}

func (e *Event) SubscriptionID() int64 {
	return e.gen.SubscriptionID()
}

func (e *Event) PublicationID() int64 {
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

func (e *Event) PayloadSerializer() int {
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

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeEvent)}, data.Bytes()...), nil
}

func CapnprotoToEvent(data []byte) (*messages.Event, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	event, err := gen.ReadRootEvent(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewEventWithFields(NewEventFields(&event)), nil
}
