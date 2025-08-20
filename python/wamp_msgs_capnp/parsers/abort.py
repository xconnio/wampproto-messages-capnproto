import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.abort import Abort, IAbortFields
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "abort.capnp")
abort_capnp = capnp.load(str(module_file))


class AbortFields(IAbortFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def reason(self) -> str:
        return self._gen.reason

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
    def payload_serializer_id(self) -> bytes:
        return self._gen.payloadSerializerID

    @property
    def payload(self) -> bytes:
        return self._ex.payload


def abort_to_capnproto(a: Abort) -> bytes:
    abort = abort_capnp.Abort.new_message()
    abort.reason = a.reason
    payload_serializer = helpers.select_payload_serializer(a.details)
    abort.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, a.args, a.kwargs)
    packed_data = abort.to_bytes_packed()

    return helpers.prepend_header(Abort.TYPE, packed_data, payload)


def capnproto_to_abort(data: bytes) -> Abort:
    message_data, payload_data = helpers.extract_message(data)
    abort_obj = abort_capnp.Abort.from_bytes_packed(message_data)

    return Abort(AbortFields(abort_obj, payload_data))
