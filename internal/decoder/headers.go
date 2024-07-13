package decoder

import (
	"github.com/segmentio/kafka-go"
)

func (d *Decoder) DecodeHeaders(headers []kafka.Header, v any) error {
	for _, header := range headers {
		if err := d.DecodeHeader(header, v); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) DecodeHeader(header kafka.Header, v any) error {
	if err := d.reflect.EnsureIsPointer(v); err != nil {
		return err
	}

	raw, err := d.unmarshal(header.Value)
	if err != nil {
		raw = string(header.Value)
	}

	hMap := map[string]any{
		header.Key: raw,
	}
	return d.reflect.CorrelateMapToValue(hMap, v, "header")
}
