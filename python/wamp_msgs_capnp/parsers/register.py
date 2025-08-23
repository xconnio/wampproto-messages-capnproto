import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.register import Register, IRegisterFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "register.capnp")
register_capnp = capnp.load(str(schema_file))


class RegisterFields(IRegisterFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        options = {}
        if self._gen.match:
            options["match"] = self._gen.match

        if self._gen.invoke:
            options["invoke"] = self._gen.invoke

        return options

    @property
    def procedure(self) -> str:
        return self._gen.procedure


def register_to_capnproto(r: Register) -> bytes:
    register = register_capnp.Register.new_message()

    register.requestID = r.request_id
    register.procedure = r.procedure

    match = r.options.get("match", None)
    if match is not None:
        register.match = match

    invoke = r.options.get("invoke", None)
    if invoke is not None:
        register.invoke = invoke

    packed_data = register.to_bytes_packed()

    return helpers.prepend_header(Register.TYPE, packed_data, b"")


def capnproto_to_register(data: bytes) -> Register:
    message_data, _ = helpers.extract_message(data)
    register_obj = register_capnp.Register.from_bytes_packed(message_data)

    return Register(RegisterFields(register_obj))
