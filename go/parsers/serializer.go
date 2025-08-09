package parsers

import (
	"encoding/json"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/xconnio/wampproto-go/serializers"
)

func EncodeToCBOR(args []any, kwargs map[string]any) ([]byte, error) {
	data := []any{args, kwargs}
	return cbor.Marshal(data)
}

func decode(arr []any) ([]any, map[string]any, error) {
	if len(arr) != 2 {
		return nil, nil, fmt.Errorf("expected 2 elements [args, kwargs], got %d", len(arr))
	}

	args, ok := arr[0].([]any)
	if !ok {
		return nil, nil, fmt.Errorf("args element is not []any")
	}

	kwargs, ok := arr[1].(map[string]any)
	if !ok {
		return nil, nil, fmt.Errorf("kwargs element is not map[string]any")
	}

	return args, kwargs, nil
}

func DecodeFromCBOR(b []byte) ([]any, map[string]any, error) {
	var arr []any
	if err := cbor.Unmarshal(b, &arr); err != nil {
		return nil, nil, err
	}

	return decode(arr)
}

func EncodeToMsgPack(args []any, kwargs map[string]any) ([]byte, error) {
	data := []any{args, kwargs}
	return msgpack.Marshal(data)
}

func DecodeFromMsgPack(b []byte) ([]any, map[string]any, error) {
	var arr []any
	if err := msgpack.Unmarshal(b, &arr); err != nil {
		return nil, nil, err
	}

	return decode(arr)
}

func EncodeToJSON(args []any, kwargs map[string]any) ([]byte, error) {
	data := []any{args, kwargs}
	return json.Marshal(data)
}

func DecodeFromJSON(b []byte) ([]any, map[string]any, error) {
	var arr []any
	if err := json.Unmarshal(b, &arr); err != nil {
		return nil, nil, err
	}

	return decode(arr)
}

func Decode(serializerID uint64, payload []byte) ([]any, map[string]any, error) {
	switch serializerID {
	case serializers.JSONSerializerID:
		return DecodeFromJSON(payload)
	case serializers.CBORSerializerID:
		return DecodeFromCBOR(payload)
	case serializers.MsgPackSerializerID:
		return DecodeFromMsgPack(payload)
	default:
		return nil, nil, fmt.Errorf("serializer %d not recognized", serializerID)
	}
}
