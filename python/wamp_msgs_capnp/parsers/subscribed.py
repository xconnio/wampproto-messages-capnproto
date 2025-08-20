import os
from pathlib import Path

import capnp
from wampproto.messages.subscribed import Subscribed, ISubscribedFields

from wamp_msgs_capnp.parsers import helpers

# Load schema
root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "subscribed.capnp")
subscribed_capnp = capnp.load(str(module_file))


class SubscribedFields(ISubscribedFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def subscription_id(self) -> int:
        return self._gen.subscriptionID


def subscribed_to_capnproto(s: Subscribed) -> bytes:
    subscribed = subscribed_capnp.Subscribed.new_message()
    subscribed.requestID = s.request_id
    subscribed.subscriptionID = s.subscription_id

    packed_data = subscribed.to_bytes_packed()

    return helpers.prepend_header(Subscribed.TYPE, packed_data, b"")


def capnproto_to_subscribed(data: bytes) -> Subscribed:
    message_data, _ = helpers.extract_message(data)
    subscribed_obj = subscribed_capnp.Subscribed.from_bytes_packed(message_data)

    return Subscribed(SubscribedFields(subscribed_obj))
