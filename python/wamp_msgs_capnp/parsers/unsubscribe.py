import os
from pathlib import Path

import capnp
from wampproto.messages import unsubscribe as unsubscribe_message
from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "unsubscribe.capnp")
unsubscribe_capnp = capnp.load(str(schema_file))


class Unsubscribe(unsubscribe_message.IUnsubscribeFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def subscription_id(self) -> int:
        return self._gen.subscriptionID


def unsubscribe_to_capnproto(m: unsubscribe_message.Unsubscribe) -> bytes:
    unsubscribe = unsubscribe_capnp.Unsubscribe.new_message()

    unsubscribe.requestID = m.request_id
    unsubscribe.subscriptionID = m.subscription_id

    packed_data = unsubscribe.to_bytes_packed()

    return helpers.prepend_header(unsubscribe_message.Unsubscribe.TYPE, packed_data, b"")


def capnproto_to_unsubscribe(data: bytes) -> unsubscribe_message.Unsubscribe:
    message_data, _ = helpers.extract_message(data)
    unsubscribe_obj = unsubscribe_capnp.Unsubscribe.from_bytes_packed(message_data)

    return Unsubscribe(unsubscribe_obj)
