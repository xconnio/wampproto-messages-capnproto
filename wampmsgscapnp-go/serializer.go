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
	case messages.MessageTypePublished:
		msg := message.(*messages.Published)
		return parsers.PublishedToCapnproto(msg)
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
	default:
		return nil, fmt.Errorf("unknown message type: %v", message.Type())
	}
}

func (c *CapnprotoSerializer) Deserialize(data []byte) (messages.Message, error) {
	switch data[0] {
	case messages.MessageTypePublished:
		return parsers.CapnprotoToPublished(data[1:])
	default:
		return nil, fmt.Errorf("unknown message type: %v", data[0])
	}
}

func (c *CapnprotoSerializer) Static() bool {
	return true
}
