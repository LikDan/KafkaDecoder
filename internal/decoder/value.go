package decoder

func (d *Decoder) DecodeValue(bytes []byte, v any) error {
	if err := d.reflect.EnsureIsPointer(v); err != nil {
		return err
	}

	if len(bytes) == 0 {
		return nil
	}

	rawValue, err := d.unmarshal(bytes)
	if err != nil {
		return err
	}

	return d.reflect.CorrelateRawToValue(rawValue, v, "value")
}
