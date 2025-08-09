package wampmsgscapnp

import (
	"fmt"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-go/serializers"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/parsers"
)

type CapnprotoSerializer struct{}

var _ serializers.Serializer = &CapnprotoSerializer{}

func (c *CapnprotoSerializer) Serialize(message messages.Message) ([]byte, error) {
	switch message.Type() {
	case messages.MessageTypeHello:
		msg := message.(*messages.Hello)
		return parsers.HelloToCapnproto(msg)
	case messages.MessageTypeWelcome:
		msg := message.(*messages.Welcome)
		return parsers.WelcomeToCapnproto(msg)
	case messages.MessageTypeChallenge:
		msg := message.(*messages.Challenge)
		return parsers.ChallengeToCapnproto(msg)
	case messages.MessageTypeAuthenticate:
		msg := message.(*messages.Authenticate)
		return parsers.AuthenticateToCapnproto(msg)
	case messages.MessageTypeAbort:
		msg := message.(*messages.Abort)
		return parsers.AbortToCapnproto(msg)
	case messages.MessageTypeError:
		msg := message.(*messages.Error)
		return parsers.ErrorToCapnproto(msg)
	case messages.MessageTypeCancel:
		msg := message.(*messages.Cancel)
		return parsers.CancelToCapnproto(msg)
	case messages.MessageTypeInterrupt:
		msg := message.(*messages.Interrupt)
		return parsers.InterruptToCapnproto(msg)
	case messages.MessageTypeGoodbye:
		msg := message.(*messages.GoodBye)
		return parsers.GoodbyeToCapnproto(msg)
	case messages.MessageTypeRegister:
		msg := message.(*messages.Register)
		return parsers.RegisterToCapnproto(msg)
	case messages.MessageTypeRegistered:
		msg := message.(*messages.Registered)
		return parsers.RegisteredToCapnproto(msg)
	case messages.MessageTypeUnregister:
		msg := message.(*messages.Unregister)
		return parsers.UnregisterToCapnproto(msg)
	case messages.MessageTypeUnregistered:
		msg := message.(*messages.Unregistered)
		return parsers.UnregisteredToCapnproto(msg)
	case messages.MessageTypeCall:
		msg := message.(*messages.Call)
		return parsers.CallToCapnproto(msg)
	case messages.MessageTypeInvocation:
		msg := message.(*messages.Invocation)
		return parsers.InvocationToCapnproto(msg)
	case messages.MessageTypeYield:
		msg := message.(*messages.Yield)
		return parsers.YieldToCapnproto(msg)
	case messages.MessageTypeResult:
		msg := message.(*messages.Result)
		return parsers.ResultToCapnproto(msg)
	case messages.MessageTypeSubscribe:
		msg := message.(*messages.Subscribe)
		return parsers.SubscribeToCapnproto(msg)
	case messages.MessageTypeSubscribed:
		msg := message.(*messages.Subscribed)
		return parsers.SubscribedToCapnproto(msg)
	case messages.MessageTypeUnsubscribe:
		msg := message.(*messages.Unsubscribe)
		return parsers.UnsubscribeToCapnproto(msg)
	case messages.MessageTypeUnsubscribed:
		msg := message.(*messages.Unsubscribed)
		return parsers.UnsubscribedToCapnproto(msg)
	case messages.MessageTypePublish:
		msg := message.(*messages.Publish)
		return parsers.PublishToCapnproto(msg)
	case messages.MessageTypePublished:
		msg := message.(*messages.Published)
		return parsers.PublishedToCapnproto(msg)
	case messages.MessageTypeEvent:
		msg := message.(*messages.Event)
		return parsers.EventToCapnproto(msg)
	default:
		return nil, fmt.Errorf("unknown message type: %v", message.Type())
	}
}

func (c *CapnprotoSerializer) Deserialize(data []byte) (messages.Message, error) {
	messageData, payloadData, err := parsers.ExtractMessage(data)
	if err != nil {
		return nil, err
	}

	switch uint64(data[0]) {
	case messages.MessageTypeHello:
		return parsers.CapnprotoToHello(messageData)
	case messages.MessageTypeWelcome:
		return parsers.CapnprotoToWelcome(messageData)
	case messages.MessageTypeChallenge:
		return parsers.CapnprotoToChallenge(messageData)
	case messages.MessageTypeAuthenticate:
		return parsers.CapnprotoToAuthenticate(messageData)
	case messages.MessageTypeAbort:
		return parsers.CapnprotoToAbort(messageData, payloadData)
	case messages.MessageTypeError:
		return parsers.CapnprotoToError(messageData, payloadData)
	case messages.MessageTypeCancel:
		return parsers.CapnprotoToCancel(messageData)
	case messages.MessageTypeInterrupt:
		return parsers.CapnprotoToInterrupt(messageData)
	case messages.MessageTypeGoodbye:
		return parsers.CapnprotoToGoodbye(messageData)
	case messages.MessageTypeRegister:
		return parsers.CapnprotoToRegister(messageData)
	case messages.MessageTypeRegistered:
		return parsers.CapnprotoToRegistered(messageData)
	case messages.MessageTypeUnregister:
		return parsers.CapnprotoToUnregister(messageData)
	case messages.MessageTypeUnregistered:
		return parsers.CapnprotoToUnregistered(messageData)
	case messages.MessageTypeCall:
		return parsers.CapnprotoToCall(messageData, payloadData)
	case messages.MessageTypeInvocation:
		return parsers.CapnprotoToInvocation(messageData, payloadData)
	case messages.MessageTypeYield:
		return parsers.CapnprotoToYield(messageData, payloadData)
	case messages.MessageTypeResult:
		return parsers.CapnprotoToResult(messageData, payloadData)
	case messages.MessageTypeSubscribe:
		return parsers.CapnprotoToSubscribe(messageData)
	case messages.MessageTypeSubscribed:
		return parsers.CapnprotoToSubscribed(messageData)
	case messages.MessageTypeUnsubscribe:
		return parsers.CapnprotoToUnsubscribe(messageData)
	case messages.MessageTypeUnsubscribed:
		return parsers.CapnprotoToUnsubscribed(messageData)
	case messages.MessageTypePublish:
		return parsers.CapnprotoToPublish(messageData, payloadData)
	case messages.MessageTypePublished:
		return parsers.CapnprotoToPublished(messageData)
	case messages.MessageTypeEvent:
		return parsers.CapnprotoToEvent(messageData, payloadData)
	default:
		return nil, fmt.Errorf("unknown message type: %v", data[0])
	}
}

func (c *CapnprotoSerializer) Static() bool {
	return true
}
