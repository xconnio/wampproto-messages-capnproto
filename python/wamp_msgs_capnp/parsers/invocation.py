import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import invocation as invocation_message
from wampproto.serializers.payload import serialize_payload

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "invocation.capnp")
invocation_capnp = capnp.load(str(schema_file))


class Invocation(invocation_message.IInvocationFields):
    def __init__(self, gen, payload: bytes):
        self._gen = gen
        self._ex = helpers.PayloadExpander(payload, gen.payloadSerializerID)

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def registration_id(self) -> str:
        return self._gen.registration_id

    @property
    def args(self) -> list[Any]:
        return self._ex.args

    @property
    def kwargs(self) -> dict[str, Any]:
        return self._ex.kwargs

    @property
    def details(self) -> dict[str, Any]:
        details = {}

        if self._gen.caller:
            details["caller"] = self._gen.caller

        if self._gen.callerAuthID:
            details["caller_authid"] = self._gen.callerAuthID

        if self._gen.callerAuthRole:
            details["caller_authrole"] = self._gen.callerAuthRole

        if self._gen.procedure:
            details["procedure"] = self._gen.procedure

        return details

    @property
    def payload_is_binary(self) -> bool:
        return True

    @property
    def payload(self) -> bytes:
        return self._ex.payload

    @property
    def payload_serializer(self) -> int:
        return self._gen.payloadSerializerID


def invocation_to_capnproto(c: invocation_message.Invocation) -> bytes:
    invocation = invocation_capnp.invocation.new_message()

    invocation.requestID = c.request_id
    invocation.registrationID = c.registration_id

    payload_serializer = helpers.select_payload_serializer(c.details)
    invocation.payloadSerializerID = payload_serializer

    payload = serialize_payload(payload_serializer, c.args, c.kwargs)
    packed_data = invocation.to_bytes_packed()

    return helpers.prepend_header(invocation_message.Invocation.TYPE, packed_data, payload)


def capnproto_to_invocation(data: bytes, payload: bytes) -> invocation_message.Invocation:
    message_data, _ = helpers.extract_message(data)
    invocation_obj = invocation_capnp.invocation.from_bytes_packed(message_data)

    return Invocation(invocation_obj, payload)
