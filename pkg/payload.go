package pkg

type Payload struct {
	Service string                 `json:"service"`
	Data    map[string]interface{} `json:"data"`
}
