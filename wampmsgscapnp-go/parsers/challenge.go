package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
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
	return map[string]any{}
}

func ChallengeToCapnproto(m *messages.Challenge) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	challenge, err := gen.NewChallenge(seg)
	if err != nil {
		return nil, err
	}

	if err := challenge.SetAuthMethod(m.AuthMethod()); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	return append([]byte{byte(messages.MessageTypeChallenge)}, data.Bytes()...), nil
}

func CapnprotoToChallenge(data []byte) (*messages.Challenge, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	challenge, err := gen.ReadRootChallenge(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewChallengeWithFields(NewChallengeFields(&challenge)), nil
}
