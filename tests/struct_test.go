package tests

import (
	"reflect"
	"testing"

	"github.com/likdan/kafka-decoder"
	"github.com/segmentio/kafka-go"
)

func TestDecodeStruct(t *testing.T) {
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
			type nested struct {
				Value string `value:"value"`
			}

			type result struct {
				Nested nested `value:"nested"`
			}

			return test{
				name: "struct_value",
				args: args{
					msg: &kafka.Message{
						Value: []byte(`
							{
								"nested": {
									"value": "value"
								}
							}
						`),
					},
					v: &result{},
				},
				want: &result{
					Nested: nested{
						Value: "value",
					},
				},
			}
		},
		func() test {
			type nested struct {
				Value string `header:"value"`
			}

			type result struct {
				Nested nested `header:"nested"`
			}

			return test{
				name: "struct_headers",
				args: args{
					msg: &kafka.Message{
						Value: make([]byte, 0),
						Headers: []kafka.Header{
							{
								Key: "nested",
								Value: []byte(`
									{
										"value": "value"
									}
								`),
							},
						},
					},
					v: &result{},
				},
				want: &result{
					Nested: nested{
						Value: "value",
					},
				},
			}
		},
		func() test {
			type valueNested struct {
				Value string `value:"value"`
			}
			type headerNested struct {
				Value string `header:"value"`
			}

			type result struct {
				ValueNested  valueNested  `value:"nested"`
				HeaderNested headerNested `header:"nested"`
			}

			return test{
				name: "struct",
				args: args{
					msg: &kafka.Message{
						Value: []byte(`
							{
								"nested": {
									"value": "value"
								}
							}
						`),
						Headers: []kafka.Header{
							{
								Key: "nested",
								Value: []byte(`
									{
										"value": "value"
									}
								`),
							},
						},
					},
					v: &result{},
				},
				want: &result{
					ValueNested: valueNested{
						Value: "value",
					},
					HeaderNested: headerNested{
						Value: "value",
					},
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
