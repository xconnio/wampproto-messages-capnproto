import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.goodbye import Goodbye, IGoodbyeFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "goodbye.capnp")
goodbye_capnp = capnp.load(str(module_file))


class GoodbyeFields(IGoodbyeFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def reason(self) -> str:
        return self._gen.reason

    @property
    def details(self) -> dict[str, Any]:
        return {}


def goodbye_to_capnproto(g: Goodbye) -> bytes:
    goodbye = goodbye_capnp.Goodbye.new_message()
    goodbye.reason = g.reason

    packed_data = goodbye.to_bytes_packed()

    return helpers.prepend_header(Goodbye.TYPE, packed_data, b"")


def capnproto_to_goodbye(data: bytes) -> Goodbye:
    message_data, _ = helpers.extract_message(data)
    goodbye_obj = goodbye_capnp.Goodbye.from_bytes_packed(message_data)

    return Goodbye(GoodbyeFields(goodbye_obj))
