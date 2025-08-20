import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.cancel import Cancel, ICancelFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "cancel.capnp")
cancel_capnp = capnp.load(str(module_file))


class CancelFields(ICancelFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        return {}


def cancel_to_capnproto(c: Cancel) -> bytes:
    cancel = cancel_capnp.Cancel.new_message()
    cancel.requestID = c.request_id

    packed_data = cancel.to_bytes_packed()

    return helpers.prepend_header(Cancel.TYPE, packed_data, b"")


def capnproto_to_cancel(data: bytes) -> Cancel:
    message_data, _ = helpers.extract_message(data)
    cancel_obj = cancel_capnp.Cancel.from_bytes_packed(message_data)

    return Cancel(CancelFields(cancel_obj))
