import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.interrupt import Interrupt, IInterruptFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "interrupt.capnp")
interrupt_capnp = capnp.load(str(module_file))


class InterruptFields(IInterruptFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def options(self) -> dict[str, Any]:
        return {}


def interrupt_to_capnproto(i: Interrupt) -> bytes:
    interrupt = interrupt_capnp.Interrupt.new_message()
    interrupt.requestID = i.request_id

    packed_data = interrupt.to_bytes_packed()

    return helpers.prepend_header(Interrupt.TYPE, packed_data, b"")


def capnproto_to_interrupt(data: bytes) -> Interrupt:
    message_data, _ = helpers.extract_message(data)
    interrupt_obj = interrupt_capnp.Interrupt.from_bytes_packed(message_data)

    return Interrupt(InterruptFields(interrupt_obj))
