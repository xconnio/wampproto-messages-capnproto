import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages import cancel as cancel_message

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "cancel.capnp")
cancel_capnp = capnp.load(str(module_file))


class Cancel(cancel_message.ICancelFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        return {}


def cancel_to_capnproto(c: cancel_message.Cancel) -> bytes:
    cancel = cancel_capnp.Cancel.new_message()
    cancel.requestID = c.request_id

    packed_data = cancel.to_bytes_packed()

    return helpers.prepend_header(cancel_message.Cancel.TYPE, packed_data, b"")


def capnproto_to_cancel(data: bytes) -> cancel_message.Cancel:
    message_data, _ = helpers.extract_message(data)
    cancel_obj = cancel_capnp.Cancel.from_bytes_packed(message_data)

    return Cancel(cancel_obj)
