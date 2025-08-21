import os
from pathlib import Path

import capnp
from wampproto.messages.event import Event, IEventFields
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "event.capnp")
event_capnp = capnp.load(str(module_file))


class EventFields(IEventFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def subscription_id(self) -> int:
        return self._gen.subscriptionID

    @property
    def publication_id(self) -> int:
        return self._gen.publicationID

    @property
    def details(self) -> dict:
        details = {}

        if self._gen.publisher:
            details["publisher"] = self._gen.publisher

        if self._gen.publisherAuthID:
            details["publisher_authid"] = self._gen.publisherAuthID

        if self._gen.publisherAuthRole:
            details["publisher_authrole"] = self._gen.publisherAuthRole

        if self._gen.topic:
            details["topic"] = self._gen.topic

        return details

    @property
    def args(self):
        return self._ex.args

    @property
    def kwargs(self):
        return self._ex.kwargs

    @property
    def payload_is_binary(self) -> bool:
        return True

    @property
    def payload(self) -> bytes:
        return self._ex.payload

    @property
    def payload_serializer(self) -> int:
        return self._gen.payloadSerializerID


def event_to_capnproto(e: Event) -> bytes:
    event = event_capnp.Event.new_message()
    event.subscriptionID = e.subscription_id
    event.publicationID = e.publication_id
    event.payloadSerializerID = e.payload_serializer

    details = e.details
    if "publisher" in details:
        event.publicationID = details["publisher"]

        if "publisher_authid" in details:
            event.publisherAuthID = details["publisher_authid"]

        if "publisher_authrole" in details:
            event.publisherAuthRole = details["publisher_authrole"]

        if "topic" in details:
            event.topic = details["topic"]

    payload_serializer = helpers.select_payload_serializer(details)
    event.payloadSerializerID = payload_serializer
    payload = serialize_payload(payload_serializer, e.args, e.kwargs)

    packed_data = event.to_bytes_packed()

    return helpers.prepend_header(Event.TYPE, packed_data, payload)


def capnproto_to_event(data: bytes) -> Event:
    message_data, payload_data = helpers.extract_message(data)
    event_obj = event_capnp.Event.from_bytes_packed(message_data)

    return Event(EventFields(event_obj, payload_data))
