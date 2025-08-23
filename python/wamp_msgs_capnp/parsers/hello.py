import os
from typing import Any
from pathlib import Path

import capnp
from wampproto.messages.hello import Hello, IHelloFields

from wamp_msgs_capnp.parsers import helpers

root_dir = Path(__file__).resolve().parents[1]
module_file = os.path.join(root_dir, "schemas", "hello.capnp")
hello_capnp = capnp.load(str(module_file))


class HelloFields(IHelloFields):
    def __init__(self, gen):
        self._gen = gen

    @property
    def realm(self) -> str:
        return self._gen.realm

    @property
    def authid(self) -> str:
        return self._gen.authid

    @property
    def authmethods(self) -> list[str]:
        return self._gen.authmethods

    @property
    def authextra(self) -> dict[str, Any]:
        extra = {}
        if self._gen.publicKey:
            extra["pubkey"] = self._gen.publicKey

        return extra

    @property
    def roles(self) -> dict[str, Any]:
        roles: dict[str, Any] = {}

        r = self._gen.roles.caller
        roles["caller"] = {
            "caller_identification": r.callerIdentification,
            "call_timeout": r.callTimeout,
            "call_canceling": r.callCanceling,
            "progressive_call_results": r.progressiveCallResults,
        }

        r = self._gen.roles.callee
        roles["callee"] = {
            "caller_identification": r.callerIdentification,
            "call_timeout": r.callTimeout,
            "call_canceling": r.callCanceling,
            "progressive_call_results": r.progressiveCallResults,
            "pattern_based_registration": r.patternBasedRegistration,
            "shared_registration": r.sharedRegistration,
        }

        r = self._gen.roles.publisher
        roles["publisher"] = {
            "publisher_identification": r.publisherIdentification,
            "publisher_exclusion": r.publisherExclusion,
            "acknowledge_event_received": r.acknowledgeEventReceived,
        }

        r = self._gen.roles.subscriber
        roles["subscriber"] = {
            "publisher_identification": r.publisherIdentification,
            "pattern_based_subscription": r.patternBasedSubscription,
        }

        return roles


def hello_to_capnproto(h: Hello) -> bytes:
    hello = hello_capnp.Hello.new_message()
    hello.realm = h.realm
    hello.authid = h.authid

    hello.authmethods = h.authmethods
    public_key = h.authextra.get("pubkey", None)
    if public_key is not None:
        hello.publicKey = public_key

    roles = hello.init("roles")

    if "caller" in h.roles:
        caller = roles.init("caller")
        caller.callerIdentification = h.roles["caller"].get("caller_identification", False)
        caller.callTimeout = h.roles["caller"].get("call_timeout", False)
        caller.callCanceling = h.roles["caller"].get("call_canceling", False)
        caller.progressiveCallResults = h.roles["caller"].get("progressive_call_results", False)

    if "callee" in h.roles:
        callee = roles.init("callee")
        callee.callerIdentification = h.roles["callee"].get("caller_identification", False)
        callee.callTimeout = h.roles["callee"].get("call_timeout", False)
        callee.callCanceling = h.roles["callee"].get("call_canceling", False)
        callee.progressiveCallResults = h.roles["callee"].get("progressive_call_results", False)
        callee.patternBasedRegistration = h.roles["callee"].get("pattern_based_registration", False)
        callee.sharedRegistration = h.roles["callee"].get("shared_registration", False)

    if "publisher" in h.roles:
        publisher = roles.init("publisher")
        publisher.publisherIdentification = h.roles["publisher"].get("publisher_identification", False)
        publisher.publisherExclusion = h.roles["publisher"].get("publisher_exclusion", False)
        publisher.acknowledgeEventReceived = h.roles["publisher"].get("acknowledge_event_received", False)

    if "subscriber" in h.roles:
        subscriber = roles.init("subscriber")
        subscriber.publisherIdentification = h.roles["subscriber"].get("publisher_identification", False)
        subscriber.patternBasedSubscription = h.roles["subscriber"].get("pattern_based_subscription", False)

    hello.roles = roles
    packed_data = hello.to_bytes_packed()

    return helpers.prepend_header(Hello.TYPE, packed_data, b"")


def capnproto_to_hello(data: bytes) -> Hello:
    message_data, _ = helpers.extract_message(data)
    hello_obj = hello_capnp.Hello.from_bytes_packed(message_data)

    return Hello(HelloFields(hello_obj))
