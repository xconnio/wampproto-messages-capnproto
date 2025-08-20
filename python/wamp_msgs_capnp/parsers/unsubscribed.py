import os
from pathlib import Path

import capnp
from wampproto.messages.unsubscribed import Unsubscribed, IUnsubscribedFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "unsubscribed.capnp")
unsubscribed_capnp = capnp.load(str(schema_file))


class UnsubscribedFields(IUnsubscribedFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID


def unsubscribed_to_capnproto(u: Unsubscribed) -> bytes:
    unsubscribed = unsubscribed_capnp.Unsubscribed.new_message()
    unsubscribed.requestID = u.request_id

    packed_data = unsubscribed.to_bytes_packed()
    return helpers.prepend_header(Unsubscribed.TYPE, packed_data, b"")


def capnproto_to_unsubscribed(data: bytes) -> Unsubscribed:
    message_data, _ = helpers.extract_message(data)
    unsubscribed_obj = unsubscribed_capnp.Unsubscribed.from_bytes_packed(message_data)

    return Unsubscribed(UnsubscribedFields(unsubscribed_obj))
