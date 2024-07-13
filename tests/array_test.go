package tests

import (
	"reflect"
	"testing"

	"github.com/likdan/kafka-decoder"
	"github.com/segmentio/kafka-go"
)

func TestDecodeArray(t *testing.T) {
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
				Primitives   []int      `value:"array_primitives"`
				Structs      []nested   `value:"array_structs"`
				Matrix       [][]int    `value:"matrix"`
				StructMatrix [][]nested `value:"struct_matrix"`
			}

			return test{
				name: "array",
				args: args{
					msg: &kafka.Message{
						Value: []byte(`
							{
								"array_primitives": [1,2,3,4,5],
								"array_structs": [
									{
										"value": "value1"
									},
									{
										"value": "value2"
									},
									{
										"value": "value3"
									}
								],
								"matrix": [
									[11,22,33,44,55],
									[21,22,23,24,25]
								],
								"struct_matrix": [
									[
										{
											"value": "value11"
										},
										{
											"value": "value12"
										}
									],
									[
										{
											"value": "value21"
										},
										{
											"value": "value22"
										}
									]
								]
							}
						`),
					},
					v: &result{},
				},
				want: &result{
					Primitives: []int{1, 2, 3, 4, 5},
					Structs: []nested{
						{
							Value: "value1",
						},
						{
							Value: "value2",
						},
						{
							Value: "value3",
						},
					},
					Matrix: [][]int{{11, 22, 33, 44, 55}, {21, 22, 23, 24, 25}},
					StructMatrix: [][]nested{
						{{Value: "value11"}, {Value: "value12"}},
						{{Value: "value21"}, {Value: "value22"}},
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
