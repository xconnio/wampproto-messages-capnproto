import os
from pathlib import Path

import capnp
from wampproto.messages import registered as registered_message

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "registered.capnp")
registered_capnp = capnp.load(str(schema_file))


class Registered(registered_message.IRegisteredFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def registration_id(self) -> int:
        return self._gen.registrationID


def registered_to_capnproto(r: registered_message.Registered) -> bytes:
    registered = registered_capnp.Registered.new_message()
    registered.requestID = r.request_id
    registered.registrationID = r.registration_id

    packed_data = registered.to_bytes_packed()
    return helpers.prepend_header(registered_message.Registered.TYPE, packed_data, b"")


def capnproto_to_registered(data: bytes) -> registered_message.Registered:
    message_data, _ = helpers.extract_message(data)
    registered_obj = registered_capnp.Registered.from_bytes_packed(message_data)

    return Registered(registered_obj)
