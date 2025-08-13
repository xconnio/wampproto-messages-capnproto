import os
from pathlib import Path

import capnp
from wampproto.messages import published as published_message
from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
schema_file = os.path.join(root_dir, "schemas", "published.capnp")
published_capnp = capnp.load(str(schema_file))


class Published(published_message.IPublishedFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def request_id(self) -> int:
        return self._gen.requestID

    @property
    def publication_id(self) -> int:
        return self._gen.publicationID


def published_to_capnproto(p: published_message.Published) -> bytes:
    published = published_capnp.Published.new_message()

    published.requestID = p.request_id
    published.publicationID = p.publication_id

    packed_data = published.to_bytes_packed()

    return helpers.prepend_header(published_message.Published.TYPE, packed_data, b"")


def capnproto_to_published(data: bytes) -> published_message.Published:
    message_data, _ = helpers.extract_message(data)
    published_obj = published_capnp.Published.from_bytes_packed(message_data)

    return Published(published_obj)
