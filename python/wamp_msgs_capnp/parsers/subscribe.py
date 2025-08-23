import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.subscribe import Subscribe, ISubscribeFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "subscribe.capnp")
subscribe_capnp = capnp.load(str(module_file))


class SubscribeFields(ISubscribeFields):
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

        return options

    @property
    def topic(self) -> str:
        return self._gen.topic


def subscribe_to_capnproto(s: Subscribe) -> bytes:
    subscribe = subscribe_capnp.Subscribe.new_message()
    subscribe.requestID = s.request_id
    subscribe.topic = s.topic

    match = s.options.get("match", None)
    if match is not None:
        subscribe.match = match

    packed_data = subscribe.to_bytes_packed()

    return helpers.prepend_header(Subscribe.TYPE, packed_data, b"")


def capnproto_to_subscribe(data: bytes) -> Subscribe:
    message_data, _ = helpers.extract_message(data)
    subscribe_obj = subscribe_capnp.Subscribe.from_bytes_packed(message_data)

    return Subscribe(SubscribeFields(subscribe_obj))
