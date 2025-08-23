package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/auth"
	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Challenge struct {
	gen *gen.Challenge
}

func NewChallengeFields(g *gen.Challenge) messages.ChallengeFields {
	return &Challenge{gen: g}
}

func (c *Challenge) AuthMethod() string {
	val, _ := c.gen.AuthMethod()
	return val
}

func (c *Challenge) Extra() map[string]any {
	extra := make(map[string]any)
	challenge, err := c.gen.Challenge()
	if err == nil {
		extra["challenge"] = challenge
	}

	salt, err := c.gen.Salt()
	if err == nil {
		extra["salt"] = salt
		extra["iterations"] = c.gen.Iterations()
		extra["keylen"] = c.gen.Keylen()
	}
	return extra
}

func ChallengeToCapnproto(m *messages.Challenge) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	challenge, err := gen.NewRootChallenge(seg)
	if err != nil {
		return nil, err
	}

	challengeString, _ := m.Extra()["challenge"].(string)
	if err = challenge.SetChallenge(challengeString); err != nil {
		return nil, err
	}

	if err := challenge.SetAuthMethod(m.AuthMethod()); err != nil {
		return nil, err
	}

	if m.AuthMethod() == auth.MethodCRA {
		saltString, ok := m.Extra()["salt"].(string)
		if ok {
			if err = challenge.SetSalt(saltString); err != nil {
				return nil, err
			}
		}

		iterations, ok := m.Extra()["iterations"].(uint32)
		if ok {
			challenge.SetIterations(iterations)
		}

		keylen, ok := m.Extra()["keylen"].(uint16)
		if ok {
			challenge.SetKeylen(keylen)
		}
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeChallenge, data, nil), nil
}

func CapnprotoToChallenge(data []byte) (*messages.Challenge, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	challenge, err := gen.ReadRootChallenge(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewChallengeWithFields(NewChallengeFields(&challenge)), nil
}
