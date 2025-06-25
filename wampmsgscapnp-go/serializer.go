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
	case messages.MessageTypePublished:
		msg := message.(*messages.Published)
		return parsers.PublishedToCapnproto(msg)
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
