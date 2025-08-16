import os
from pathlib import Path

import capnp
from wampproto.messages import unsubscribed as unsubscribed_message

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "unsubscribed.capnp")
unsubscribed_capnp = capnp.load(str(schema_file))


class Unsubscribed(unsubscribed_message.IUnsubscribedFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID


def unsubscribed_to_capnproto(unsub: unsubscribed_message.Unsubscribed) -> bytes:
    unsubscribed = unsubscribed_capnp.Unsubscribed.new_message()
    unsubscribed.requestID = unsub.request_id

    packed_data = unsubscribed.to_bytes_packed()
    return helpers.prepend_header(unsubscribed_message.Unsubscribed.TYPE, packed_data, b"")


def capnproto_to_unsubscribed(data: bytes) -> unsubscribed_message.Unsubscribed:
    message_data, _ = helpers.extract_message(data)
    unsubscribed_obj = unsubscribed_capnp.Unsubscribed.from_bytes_packed(message_data)

    return Unsubscribed(unsubscribed_obj)
