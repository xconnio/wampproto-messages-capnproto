import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.authenticate import Authenticate, IAuthenticateFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "authenticate.capnp")
authenticate_capnp = capnp.load(str(module_file))


class AuthenticateFields(IAuthenticateFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def signature(self) -> str:
        return self._gen.signature

    @property
    def extra(self) -> dict[str, Any]:
        return {}


def authenticate_to_capnproto(a: Authenticate) -> bytes:
    authenticate = authenticate_capnp.Authenticate.new_message()
    authenticate.signature = a.signature

    data = authenticate.to_bytes_packed()

    return helpers.prepend_header(Authenticate.TYPE, data, b"")


def capnproto_to_authenticate(data: bytes) -> Authenticate:
    message_data, _ = helpers.extract_message(data)
    authenticate_obj = authenticate_capnp.Authenticate.from_bytes_packed(message_data)

    return Authenticate(AuthenticateFields(authenticate_obj))
