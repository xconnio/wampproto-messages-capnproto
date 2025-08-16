import os
from pathlib import Path
from typing import Any

import capnp
from wampproto.messages import result as result_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "result.capnp")
result_capnp = capnp.load(str(schema_file))


class Result(result_message.IResultFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def details(self) -> dict[str, Any]:
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


def result_to_capnproto(r: result_message.Result) -> bytes:
    result = result_capnp.Result.new_message()

    result.requestID = r.request_id

    payload_serializer = helpers.select_payload_serializer(r.options)
    result.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, r.args, r.kwargs)
    packed_data = result.to_bytes_packed()

    return helpers.prepend_header(result_message.Result.TYPE, packed_data, payload)


def capnproto_to_result(data: bytes, payload: bytes) -> result_message.Result:
    message_data, _ = helpers.extract_message(data)
    result_obj = result_capnp.Result.from_bytes_packed(message_data)

    return Result(result_obj, payload)
