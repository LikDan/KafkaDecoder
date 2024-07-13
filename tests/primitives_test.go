package tests

import (
	"reflect"
	"testing"

	"github.com/likdan/kafka-decoder"
	"github.com/segmentio/kafka-go"
)

func TestDecodePrimitives(t *testing.T) {
	type args struct {
		msg *kafka.Message
		v   any
	}

	type test struct {
		name    string
		args    args
		want    any
		wantErr bool
	}

	tests := []func() test{
		func() test {
			type result struct {
				Null    any     `value:"null"`
				Int     int     `value:"int"`
				Int32   int32   `value:"int32"`
				Int64   int64   `value:"int64"`
				Float32 float32 `value:"float32"`
				Float64 float64 `value:"float64"`
				Bool    bool    `value:"bool"`
				String  string  `value:"string"`
			}

			return test{
				name: "primitives_value",
				args: args{
					msg: &kafka.Message{
						Value: []byte(`
							{
								"null": null,
								"int": 1,
								"int32": 32,
								"int64": 64,
								"float32": 32.32,
								"float64": 64.64,
								"bool": true,
								"string": "string"
							}
						`),
					},
					v: &result{},
				},
				want: &result{
					Null:    nil,
					Int:     1,
					Int32:   32,
					Int64:   64,
					Float32: 32.32,
					Float64: 64.64,
					Bool:    true,
					String:  "string",
				},
			}
		},
		func() test {
			type result struct {
				Int            int     `header:"int"`
				Int32          int32   `header:"int32"`
				Int64          int64   `header:"int64"`
				Float32        float32 `header:"float32"`
				Float64        float64 `header:"float64"`
				Bool           bool    `header:"bool"`
				String         string  `header:"string"`
				StringNoQuotes string  `header:"stringNoQuotes"`
			}

			return test{
				name: "primitives_header",
				args: args{
					msg: &kafka.Message{
						Value: make([]byte, 0),
						Headers: []kafka.Header{
							{
								Key:   "int",
								Value: []byte("1"),
							},
							{
								Key:   "int32",
								Value: []byte("32"),
							},
							{
								Key:   "int64",
								Value: []byte("64"),
							},
							{
								Key:   "float32",
								Value: []byte("32.32"),
							},
							{
								Key:   "float64",
								Value: []byte("64.64"),
							},
							{
								Key:   "bool",
								Value: []byte("true"),
							},
							{
								Key:   "string",
								Value: []byte("\"string\""),
							},
							{
								Key:   "stringNoQuotes",
								Value: []byte("stringNoQuotes"),
							},
						},
					},
					v: &result{},
				},
				want: &result{
					Int:            1,
					Int32:          32,
					Int64:          64,
					Float32:        32.32,
					Float64:        64.64,
					Bool:           true,
					String:         "string",
					StringNoQuotes: "stringNoQuotes",
				},
			}
		},
		func() test {
			type result struct {
				ValueNull            any     `value:"null"`
				ValueInt             int     `value:"int"`
				ValueInt32           int32   `value:"int32"`
				ValueInt64           int64   `value:"int64"`
				ValueFloat32         float32 `value:"float32"`
				ValueFloat64         float64 `value:"float64"`
				ValueBool            bool    `value:"bool"`
				ValueString          string  `value:"string"`
				ValueStringNoQuotes  string  `value:"stringNoQuotes"`
				HeaderInt            int     `header:"int"`
				HeaderInt32          int32   `header:"int32"`
				HeaderInt64          int64   `header:"int64"`
				HeaderFloat32        float32 `header:"float32"`
				HeaderFloat64        float64 `header:"float64"`
				HeaderBool           bool    `header:"bool"`
				HeaderString         string  `header:"string"`
				HeaderStringNoQuotes string  `header:"stringNoQuotes"`
			}

			return test{
				name: "primitives",
				args: args{
					msg: &kafka.Message{
						Value: []byte(`
							{
								"null": null,
								"int": 1,
								"int32": 32,
								"int64": 64,
								"float32": 32.32,
								"float64": 64.64,
								"bool": true,
								"string": "string"
							}
						`),
						Headers: []kafka.Header{
							{
								Key:   "int",
								Value: []byte("1"),
							},
							{
								Key:   "int32",
								Value: []byte("32"),
							},
							{
								Key:   "int64",
								Value: []byte("64"),
							},
							{
								Key:   "float32",
								Value: []byte("32.32"),
							},
							{
								Key:   "float64",
								Value: []byte("64.64"),
							},
							{
								Key:   "bool",
								Value: []byte("true"),
							},
							{
								Key:   "string",
								Value: []byte("\"string\""),
							},
							{
								Key:   "stringNoQuotes",
								Value: []byte("stringNoQuotes"),
							},
						},
					},
					v: &result{},
				},
				want: &result{
					ValueNull:            nil,
					ValueInt:             1,
					ValueInt32:           32,
					ValueInt64:           64,
					ValueFloat32:         32.32,
					ValueFloat64:         64.64,
					ValueBool:            true,
					ValueString:          "string",
					HeaderInt:            1,
					HeaderInt32:          32,
					HeaderInt64:          64,
					HeaderFloat32:        32.32,
					HeaderFloat64:        64.64,
					HeaderBool:           true,
					HeaderString:         "string",
					HeaderStringNoQuotes: "stringNoQuotes",
				},
			}
		},
	}
	for _, ttf := range tests {
		tt := ttf()
		t.Run(tt.name, func(t *testing.T) {
			if err := dkafka.Unmarshal(tt.args.msg, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("Unmarshal() got = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
