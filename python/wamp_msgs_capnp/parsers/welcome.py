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
        details = {
            "authid": self._gen.authid,
            "authrole": self._gen.authrole,
            "authmethod": self._gen.authmethod,
            "authprovider": self._gen.authprovider,
        }

        if self._gen.roles.dealer.callerIdentification:
            details["dealer"]["caller_identification"] = True

        if self._gen.roles.dealer.callTimeout:
            details["dealer"]["call_timeout"] = True

        if self._gen.roles.dealer.callCanceling:
            details["dealer"]["call_canceling"] = True

        if self._gen.roles.dealer.progressiveCallResults:
            details["dealer"]["progressive_call_results"] = True

        if self._gen.roles.dealer.patternBasedRegistration:
            details["dealer"]["pattern_based_registration"] = True

        if self._gen.roles.dealer.sharedRegistration:
            details["dealer"]["shared_registration"] = True

        r = self._gen.roles.callee
        details["broker"] = {
            "broker_identification": r.publisherIdentification,
            "publisher_exclusion": r.publisherExclusion,
            "acknowledge_event_received": r.acknowledgeEventReceived,
            "pattern_based_subscription": r.patternBasedSubscription,
        }

        if self._gen.role.broker.publisherIdentification:
            details["broker"]["publisher_identification"] = True

        if self._gen.role.broker.publisherExclusion:
            details["broker"]["publisher_exclusion"] = True

        if self._gen.role.broker.acknowledgeEventReceived:
            details["broker"]["acknowledge_event_received"] = True

        if self._gen.role.broker.patternBasedSubscription:
            details["broker"]["pattern_based_subscription"] = True

        return details

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

    roles = welcome.init("roles")
    if "dealer" in w.roles:
        dealer = roles.init("dealer")
        dealer.callerIdentification = w.roles["dealer"].get("caller_identification", False)
        dealer.callTimeout = w.roles["dealer"].get("call_timeout", False)
        dealer.callCanceling = w.roles["dealer"].get("call_canceling", False)
        dealer.progressiveCallResults = w.roles["dealer"].get("progressive_call_results", False)
        dealer.patternBasedSubscription = w.roles["dealer"].get("pattern_based_registration", False)
        dealer.sharedRegistration = w.roles["dealer"].get("shared_registration", False)

    if "broker" in w.roles:
        broker = roles.init("broker")
        broker.publisherIdentification = w.roles["broker"].get("publisher_identification", False)
        broker.publisherExclusion = w.roles["broker"].get("publisher_exclusion", False)
        broker.acknowledgeEventReceived = w.roles["broker"].get("acknowledge_event_received", False)
        broker.patternBasedSubscription = w.roles["broker"].get("pattern_based_subscription", False)

    welcome.roles = roles

    packed_data = welcome.to_bytes_packed()

    return helpers.prepend_header(Welcome.TYPE, packed_data, b"")


def capnproto_to_welcome(data: bytes) -> Welcome:
    message_data, _ = helpers.extract_message(data)
    welcome_obj = welcome_capnp.Welcome.from_bytes_packed(message_data)

    return Welcome(WelcomeFields(welcome_obj))
