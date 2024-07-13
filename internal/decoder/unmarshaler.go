package decoder

import (
	"encoding/json"
)

func (d *Decoder) unmarshal(bytes []byte) (v any, err error) {
	err = json.Unmarshal(bytes, &v)
	return
}
