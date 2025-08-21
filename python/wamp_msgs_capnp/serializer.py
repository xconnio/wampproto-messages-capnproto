from wampproto import messages, serializers

from wamp_msgs_capnp.parsers.hello import hello_to_capnproto, capnproto_to_hello
from wamp_msgs_capnp.parsers.welcome import welcome_to_capnproto, capnproto_to_welcome
from wamp_msgs_capnp.parsers.challenge import challenge_to_capnproto, capnproto_to_challenge
from wamp_msgs_capnp.parsers.authenticate import authenticate_to_capnproto, capnproto_to_authenticate
from wamp_msgs_capnp.parsers.abort import abort_to_capnproto, capnproto_to_abort
from wamp_msgs_capnp.parsers.error import error_to_capnproto, capnproto_to_error
from wamp_msgs_capnp.parsers.cancel import cancel_to_capnproto, capnproto_to_cancel
from wamp_msgs_capnp.parsers.interrupt import interrupt_to_capnproto, capnproto_to_interrupt
from wamp_msgs_capnp.parsers.goodbye import goodbye_to_capnproto, capnproto_to_goodbye
from wamp_msgs_capnp.parsers.register import register_to_capnproto, capnproto_to_register
from wamp_msgs_capnp.parsers.registered import registered_to_capnproto, capnproto_to_registered
from wamp_msgs_capnp.parsers.unregister import unregister_to_capnproto, capnproto_to_unregister
from wamp_msgs_capnp.parsers.unregistered import unregistered_to_capnproto, capnproto_to_unregistered
from wamp_msgs_capnp.parsers.call import call_to_capnproto, capnproto_to_call
from wamp_msgs_capnp.parsers.invocation import invocation_to_capnproto, capnproto_to_invocation
from wamp_msgs_capnp.parsers.yield_ import yield_to_capnproto, capnproto_to_yield
from wamp_msgs_capnp.parsers.result import result_to_capnproto, capnproto_to_result
from wamp_msgs_capnp.parsers.subscribe import subscribe_to_capnproto, capnproto_to_subscribe
from wamp_msgs_capnp.parsers.subscribed import subscribed_to_capnproto, capnproto_to_subscribed
from wamp_msgs_capnp.parsers.unsubscribe import unsubscribe_to_capnproto, capnproto_to_unsubscribe
from wamp_msgs_capnp.parsers.unsubscribed import unsubscribed_to_capnproto, capnproto_to_unsubscribed
from wamp_msgs_capnp.parsers.publish import publish_to_capnproto, capnproto_to_publish
from wamp_msgs_capnp.parsers.event import event_to_capnproto, capnproto_to_event


class CapnProtoSerializer(serializers.Serializer):
    def serialize(self, message: messages.Message) -> bytes:
        if isinstance(message, messages.Hello):
            return hello_to_capnproto(message)
        elif isinstance(message, messages.Welcome):
            return welcome_to_capnproto(message)
        elif isinstance(message, messages.Challenge):
            return challenge_to_capnproto(message)
        elif isinstance(message, messages.Authenticate):
            return authenticate_to_capnproto(message)
        elif isinstance(message, messages.Abort):
            return abort_to_capnproto(message)
        elif isinstance(message, messages.Error):
            return error_to_capnproto(message)
        elif isinstance(message, messages.Cancel):
            return cancel_to_capnproto(message)
        elif isinstance(message, messages.Interrupt):
            return interrupt_to_capnproto(message)
        elif isinstance(message, messages.Goodbye):
            return goodbye_to_capnproto(message)
        elif isinstance(message, messages.Register):
            return register_to_capnproto(message)
        elif isinstance(message, messages.Registered):
            return registered_to_capnproto(message)
        elif isinstance(message, messages.Unregister):
            return unregister_to_capnproto(message)
        elif isinstance(message, messages.Unregistered):
            return unregistered_to_capnproto(message)
        elif isinstance(message, messages.Call):
            return call_to_capnproto(message)
        elif isinstance(message, messages.Invocation):
            return invocation_to_capnproto(message)
        elif isinstance(message, messages.Yield):
            return yield_to_capnproto(message)
        elif isinstance(message, messages.Result):
            return result_to_capnproto(message)
        elif isinstance(message, messages.Subscribe):
            return subscribe_to_capnproto(message)
        elif isinstance(message, messages.Subscribed):
            return subscribed_to_capnproto(message)
        elif isinstance(message, messages.Unsubscribe):
            return unsubscribe_to_capnproto(message)
        elif isinstance(message, messages.Unsubscribed):
            return unsubscribed_to_capnproto(message)
        elif isinstance(message, messages.Publish):
            return publish_to_capnproto(message)
        elif isinstance(message, messages.Event):
            return event_to_capnproto(message)
        else:
            raise TypeError(f"unknown message type {message.TYPE}")

    def deserialize(self, data: bytes) -> messages.Message:
        msg_type = data[0]
        match msg_type:
            case messages.Hello.TYPE:
                return capnproto_to_hello(data)
            case messages.Welcome.TYPE:
                return capnproto_to_welcome(data)
            case messages.Challenge.TYPE:
                return capnproto_to_challenge(data)
            case messages.Authenticate.TYPE:
                return capnproto_to_authenticate(data)
            case messages.Abort.TYPE:
                return capnproto_to_abort(data)
            case messages.Error.TYPE:
                return capnproto_to_error(data)
            case messages.Cancel.TYPE:
                return capnproto_to_cancel(data)
            case messages.Interrupt.TYPE:
                return capnproto_to_interrupt(data)
            case messages.Goodbye.TYPE:
                return capnproto_to_goodbye(data)
            case messages.Register.TYPE:
                return capnproto_to_register(data)
            case messages.Registered.TYPE:
                return capnproto_to_registered(data)
            case messages.Unregister.TYPE:
                return capnproto_to_unregister(data)
            case messages.Unregistered.TYPE:
                return capnproto_to_unregistered(data)
            case messages.Call.TYPE:
                return capnproto_to_call(data)
            case messages.Invocation.TYPE:
                return capnproto_to_invocation(data)
            case messages.Yield.TYPE:
                return capnproto_to_yield(data)
            case messages.Result.TYPE:
                return capnproto_to_result(data)
            case messages.Subscribe.TYPE:
                return capnproto_to_subscribe(data)
            case messages.Subscribed.TYPE:
                return capnproto_to_subscribed(data)
            case messages.Unsubscribe.TYPE:
                return capnproto_to_unsubscribe(data)
            case messages.Unsubscribed.TYPE:
                return capnproto_to_unsubscribed(data)
            case messages.Publish.TYPE:
                return capnproto_to_publish(data)
            case messages.Event.TYPE:
                return capnproto_to_event(data)
            case _:
                raise ValueError(f"unknown message type {msg_type}")

    def static(self) -> bool:
        return True
