import os
from pathlib import Path

import capnp
from wampproto.messages.registered import Registered, IRegisteredFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "registered.capnp")
registered_capnp = capnp.load(str(schema_file))


class RegisteredFields(IRegisteredFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def registration_id(self) -> int:
        return self._gen.registrationID


def registered_to_capnproto(r: Registered) -> bytes:
    registered = registered_capnp.Registered.new_message()
    registered.requestID = r.request_id
    registered.registrationID = r.registration_id

    packed_data = registered.to_bytes_packed()
    return helpers.prepend_header(Registered.TYPE, packed_data, b"")


def capnproto_to_registered(data: bytes) -> Registered:
    message_data, _ = helpers.extract_message(data)
    registered_obj = registered_capnp.Registered.from_bytes_packed(message_data)

    return Registered(RegisteredFields(registered_obj))
