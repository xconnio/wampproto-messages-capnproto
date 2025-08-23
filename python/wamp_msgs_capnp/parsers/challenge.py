import os
from typing import Any
from pathlib import Path

import capnp
from wampproto import auth
from wampproto.messages.challenge import Challenge, IChallengeFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "challenge.capnp")
challenge_capnp = capnp.load(str(module_file))


class ChallengeFields(IChallengeFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def authmethod(self) -> str:
        return self._gen.authMethod

    @property
    def extra(self) -> dict[str, Any]:
        extra = {}
        if self._gen.challenge:
            extra["challenge"] = self._gen.challenge

        if self._gen.salt:
            extra["salt"] = self._gen.salt
            extra["iterations"] = self._gen.iterations
            extra["keylen"] = self._gen.keylen

        return extra


def challenge_to_capnproto(c: Challenge) -> bytes:
    challenge = challenge_capnp.Challenge.new_message()

    challenge_string = c.extra.get("challenge", "")
    challenge.challenge = challenge_string
    challenge.authMethod = c.authmethod

    if c.authmethod == auth.WAMPCRAAuthenticator.TYPE:
        if c.extra.get("salt", None) is not None:
            challenge.salt = c.extra.get("salt")
        if c.extra.get("iterations", None) is not None:
            challenge.iterations = c.extra.get("iterations")
        if c.extra.get("keylen", None) is not None:
            challenge.keylen = c.extra.get("keylen")

    data = challenge.to_bytes_packed()

    return helpers.prepend_header(Challenge.TYPE, data, b"")


def capnproto_to_challenge(data: bytes) -> Challenge:
    message_data, _ = helpers.extract_message(data)
    challenge_obj = challenge_capnp.Challenge.from_bytes_packed(message_data)

    return Challenge(ChallengeFields(challenge_obj))
