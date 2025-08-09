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
