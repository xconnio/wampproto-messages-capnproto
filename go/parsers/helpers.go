package parsers

import (
	"encoding/binary"
	"fmt"
)

func PrependHeader(messageType uint64, payload []byte) []byte {
	result := make([]byte, 3+len(payload))

	result[0] = uint8(messageType)
	binary.BigEndian.PutUint16(result[1:3], uint16(len(payload)))

	copy(result[3:], payload)

	return result
}

func ExtractMessage(data []byte) ([]byte, []byte, error) {
	if len(data) < 3 {
		return nil, nil, fmt.Errorf("invalid message length must be at least 3 bytes")
	}

	messageLength := binary.BigEndian.Uint16(data[1:3])
	if len(data) < 3+int(messageLength) {
		return nil, nil, fmt.Errorf("invalid message length")
	}

	messageData := data[3 : 3+int(messageLength)]
	payloadData := data[3+int(messageLength):]

	return messageData, payloadData, nil
}

type PayloadExpander struct {
	expanded   bool
	payload    []byte
	serializer uint64

	args   []any
	kwargs map[string]any
}

func (p *PayloadExpander) NewPayloadExpander(serializer uint64, payload []byte) *PayloadExpander {
	return &PayloadExpander{
		serializer: serializer,
		payload:    payload,
	}
}

func (p *PayloadExpander) Expand() error {
	args, kwargs, err := Decode(p.serializer, p.payload)
	if err != nil {
		return err
	}

	p.args = args
	p.kwargs = kwargs
	p.expanded = true
	return nil
}

func (p *PayloadExpander) Args() []any {
	if !p.expanded {
		if err := p.Expand(); err != nil {
			return nil
		}
	}

	return p.args
}

func (p *PayloadExpander) Kwargs() map[string]any {
	if !p.expanded {
		if err := p.Expand(); err != nil {
			return nil
		}
	}

	return p.kwargs
}

func (p *PayloadExpander) Payload() []byte {
	return p.payload
}
