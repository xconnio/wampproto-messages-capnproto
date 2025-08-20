import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.welcome import Welcome, IWelcomeFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "welcome.capnp")
welcome_capnp = capnp.load(str(module_file))


class WelcomeFields(IWelcomeFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def session_id(self) -> int:
        return self._gen.sessionID

    @property
    def details(self) -> dict:
        return {
            "authid": self._gen.authid,
            "authrole": self._gen.authrole,
            "authmethod": self._gen.authmethod,
            "authprovider": self._gen.authprovider,
        }

    @property
    def roles(self) -> dict[str, Any]:
        return self._gen.roles

    @property
    def authid(self) -> str:
        return self._gen.authid

    @property
    def authrole(self) -> str:
        return self._gen.authrole

    @property
    def authmethod(self) -> str:
        return self._gen.authmethod

    @property
    def authextra(self) -> dict[str, Any]:
        return {}


def welcome_to_capnproto(w: Welcome) -> bytes:
    welcome = welcome_capnp.Welcome.new_message()

    welcome.sessionID = w.session_id
    welcome.authid = w.authid
    welcome.authrole = w.authrole
    welcome.authmethod = w.authmethod
    packed_data = welcome.to_bytes_packed()

    return helpers.prepend_header(Welcome.TYPE, packed_data, b"")


def capnproto_to_welcome(data: bytes) -> Welcome:
    message_data, _ = helpers.extract_message(data)
    welcome_obj = welcome_capnp.Welcome.from_bytes_packed(message_data)

    return Welcome(WelcomeFields(welcome_obj))
