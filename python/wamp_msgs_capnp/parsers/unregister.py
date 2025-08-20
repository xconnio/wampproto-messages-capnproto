import os
from pathlib import Path

import capnp
from wampproto.messages.unregister import Unregister, IUnregisterFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "unregister.capnp")
unregister_capnp = capnp.load(str(schema_file))


class UnregisterFields(IUnregisterFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def registration_id(self) -> int:
        return self._gen.registrationID


def unregister_to_capnproto(u: Unregister) -> bytes:
    unregister = unregister_capnp.Unregister.new_message()
    unregister.requestID = u.request_id
    unregister.registrationID = u.registration_id

    packed_data = unregister.to_bytes_packed()
    return helpers.prepend_header(Unregister.TYPE, packed_data, b"")


def capnproto_to_unregister(data: bytes) -> Unregister:
    message_data, _ = helpers.extract_message(data)
    unregister_obj = unregister_capnp.Unregister.from_bytes_packed(message_data)

    return Unregister(UnregisterFields(unregister_obj))
