import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import call as call_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "call.capnp")
call_capnp = capnp.load(str(schema_file))


class Call(call_message.ICallFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def procedure(self) -> str:
        return self._gen.procedure

    @property
    def args(self) -> list[Any]:
        return self._ex.args

    @property
    def kwargs(self) -> dict[str, Any]:
        return self._ex.kwargs

    @property
    def options(self) -> dict[str, Any]:
        return {}

    @property
    def payload_is_binary(self) -> bool:
        return True

    @property
    def payload(self) -> bytes:
        return self._ex.payload

    @property
    def payload_serializer(self) -> int:
        return self._gen.payloadSerializerID


def call_to_capnproto(c: call_message.Call) -> bytes:
    call = call_capnp.Call.new_message()

    call.requestID = c.request_id
    call.procedure = c.uri

    payload_serializer = helpers.select_payload_serializer(c.options)
    call.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, c.args, c.kwargs)
    packed_data = call.to_bytes_packed()

    return helpers.prepend_header(call_message.Call.TYPE, packed_data, payload)


def capnproto_to_call(data: bytes, payload: bytes) -> call_message.Call:
    message_data, _ = helpers.extract_message(data)
    call_obj = call_capnp.Call.from_bytes_packed(message_data)

    return Call(call_obj, payload)
