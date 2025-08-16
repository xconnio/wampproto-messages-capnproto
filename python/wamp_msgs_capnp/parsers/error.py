import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import error as error_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "error.capnp")
error_capnp = capnp.load(str(module_file))


class Error(error_message.IErrorFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def message_type(self) -> int:
        return self._gen.messageType

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def uri(self) -> str:
        return self._gen.uri

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
    def payload_serializer_id(self) -> int:
        return self._gen.payloadSerializerID


def error_to_capnproto(e: error_message.Error) -> bytes:
    error = error_capnp.Error.new_message()

    error.messageType = e.message_type
    error.requestID = e.request_id
    error.uri = e.uri

    payload_serializer = helpers.select_payload_serializer(e.details)
    error.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, e.args, e.kwargs)
    packed_data = error.to_bytes_packed()

    return helpers.prepend_header(error_message.Error.TYPE, packed_data, payload)


def capnproto_to_error(data: bytes) -> error_message.Error:
    message_data, payload_data = helpers.extract_message(data)
    err_obj = error_capnp.Error.from_bytes_packed(message_data)

    return Error(err_obj, payload_data)
