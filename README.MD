# KafkaDecoder

Decodes kafka message to struct using JSON and reflection.

___

## How to use

Struct
```go
package main

type Example struct {
	ValueField  string `value:"field"`
	HeaderField string `header:"headerName"`
}
```

KafkaValue
```json
{
  "field": "value"
}
```

KafkaHeaders
```
headerName -> "value" or value
```

___

Unmarshal

```go
package main

import (
	"github.com/likdan/kafka-decoder"
	"github.com/segmentio/kafka-go"
)

func decode(msg *kafka.Message) (out Example) {
	dkafka.Unmarshal(msg, &out)
}
```

Decoder
```go
package main

import (
	"github.com/likdan/kafka-decoder"
	"github.com/segmentio/kafka-go"
)

func decode(msg *kafka.Message) (out Example) {
	dkafka.NewDecoder(msg).Decode(&out)
}
```
