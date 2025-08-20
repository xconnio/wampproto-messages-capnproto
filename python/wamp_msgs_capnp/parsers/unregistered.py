import os
from pathlib import Path

import capnp
from wampproto.messages.unregistered import Unregistered, IUnregisteredFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "unregistered.capnp")
unregistered_capnp = capnp.load(str(schema_file))


class UnregisteredFields(IUnregisteredFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID


def unregistered_to_capnproto(u: Unregistered) -> bytes:
    unregistered = unregistered_capnp.Unregistered.new_message()
    unregistered.requestID = u.request_id

    packed_data = unregistered.to_bytes_packed()

    return helpers.prepend_header(Unregistered.TYPE, packed_data, b"")


def capnproto_to_unregistered(data: bytes) -> Unregistered:
    message_data, _ = helpers.extract_message(data)
    unregistered_obj = unregistered_capnp.Unregistered.from_bytes_packed(message_data)

    return Unregistered(UnregisteredFields(unregistered_obj))
