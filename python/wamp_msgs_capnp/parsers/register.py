import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import register as register_message

from wamp_msgs_capnp.parsers import helpers

# Load the Cap'n Proto schema
root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "register.capnp")
register_capnp = capnp.load(str(schema_file))


class Register(register_message.IRegisterFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        return {}

    @property
    def procedure(self) -> str:
        return self._gen.procedure


def register_to_capnproto(r: register_message.Register) -> bytes:
    register = register_capnp.Register.new_message()

    register.requestID = r.request_id
    register.procedure = r.uri

    packed_data = register.to_bytes_packed()

    return helpers.prepend_header(register_message.Register.TYPE, packed_data, b"")


def capnproto_to_register(data: bytes) -> register_message.Register:
    message_data, _ = helpers.extract_message(data)
    register_obj = register_capnp.Register.from_bytes_packed(message_data)

    return Register(register_obj)
