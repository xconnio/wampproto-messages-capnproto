import os
from typing import Any, Optional
from pathlib import Path

import capnp
from wampproto.messages import publish as publish_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "publish.capnp")
publish_capnp = capnp.load(str(module_file))


class Publish(publish_message.IPublishFields):
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


def publish_to_capnproto(p: publish_message.Publish) -> bytes:
    publish = publish_capnp.Publish.new_message()

    publish.requestID = p.request_id
    publish.topic = p.uri

    payload_serializer = helpers.select_payload_serializer(p.options)
    publish.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, p.args, p.kwargs)
    packed_data = publish.to_bytes_packed()

    return helpers.prepend_header(publish_message.Publish.TYPE, packed_data, payload)


def capnproto_to_publish(data: bytes) -> publish_message.Publish:
    message_data, payload_data = helpers.extract_message(data)
    publish_obj = publish_capnp.Publish.from_bytes_packed(message_data)

    return Publish(publish_obj, payload_data)
