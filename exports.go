package dkafka

import (
	"github.com/likdan/kafka-decoder/internal/decoder"
	"github.com/segmentio/kafka-go"
)

type Decoder = decoder.Decoder

func NewDecoder(msg *kafka.Message) *decoder.Decoder {
	return decoder.NewDecoder(msg)
}

func Unmarshal(msg *kafka.Message, v any) error {
	return decoder.NewDecoder(msg).Decode(v)
}
