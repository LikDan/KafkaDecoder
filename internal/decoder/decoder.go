package decoder

import (
	"fmt"

	"github.com/likdan/kafka-decoder/internal/reflect"
	"github.com/segmentio/kafka-go"
)

type Decoder struct {
	msg *kafka.Message

	reflect *reflect.Reflect
}

func NewDecoder(msg *kafka.Message) *Decoder {
	return &Decoder{
		msg:     msg,
		reflect: reflect.NewReflect(),
	}
}

func (d *Decoder) Decode(v any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	if err = d.DecodeHeaders(d.msg.Headers, v); err != nil {
		return err
	}

	if err = d.DecodeValue(d.msg.Value, v); err != nil {
		return err
	}

	return nil
}
