import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import yield_ as yield_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "yield.capnp")
yield_capnp = capnp.load(str(schema_file))


class Yield(yield_message.IYieldFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        return {}

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
    def payload_serializer(self) -> int:
        return self._gen.payloadSerializerID


def yield_to_capnproto(y: yield_message.Yield) -> bytes:
    yield_obj = yield_capnp.Yield.new_message()

    yield_obj.requestID = y.request_id

    payload_serializer = helpers.select_payload_serializer(y.options)
    yield_obj.payloadSerializerID = payload_serializer

    if y.payload_is_binary():
        payload = y.payload
    else:
        payload = serialize_payload(payload_serializer, y.args, y.kwargs)

    packed_data = yield_obj.to_bytes_packed()

    return helpers.prepend_header(yield_message.Yield.TYPE, packed_data, payload)


def capnproto_to_yield(data: bytes) -> yield_message.Yield:
    message_data, payload_data = helpers.extract_message(data)
    yield_obj = yield_capnp.Yield.from_bytes_packed(message_data)

    return Yield(yield_obj, payload_data)
