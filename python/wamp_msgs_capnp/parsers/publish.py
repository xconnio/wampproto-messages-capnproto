import os
from typing import Any, Optional
from pathlib import Path

import capnp
from wampproto.messages.publish import Publish, IPublishFields
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "publish.capnp")
publish_capnp = capnp.load(str(module_file))


class PublishFields(IPublishFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def topic(self) -> str:
        return self._gen.topic

    @property
    def options(self) -> Optional[dict[str, Any]]:
        details = {}

        if not self._gen.excludeMe:
            details["exclude_me"] = False

        return details

    @property
    def args(self) -> list[Any]:
        return self._ex.args

    @property
    def kwargs(self) -> dict[str, Any]:
        return self._ex.kwargs

    @property
    def payload_is_binary(self) -> bool:
        return True

    @property
    def payload(self) -> bytes:
        return self._ex.payload

    @property
    def payload_serializer_id(self) -> int:
        return self._gen.payloadSerializerID


def publish_to_capnproto(p: Publish) -> bytes:
    publish = publish_capnp.Publish.new_message()

    publish.requestID = p.request_id
    publish.topic = p.topic

    payload_serializer = helpers.select_payload_serializer(p.options)
    publish.payloadSerializerID = payload_serializer

    if p.payload_is_binary():
        payload = p.payload
    else:
        payload = serialize_payload(payload_serializer, p.args, p.kwargs)

    packed_data = publish.to_bytes_packed()

    return helpers.prepend_header(Publish.TYPE, packed_data, payload)


def capnproto_to_publish(data: bytes) -> Publish:
    message_data, payload_data = helpers.extract_message(data)
    publish_obj = publish_capnp.Publish.from_bytes_packed(message_data)

    return Publish(PublishFields(publish_obj, payload_data))
