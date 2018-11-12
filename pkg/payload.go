package pkg

type Payload struct {
	Service string                 `json:"service"`
	Data    map[string]interface{} `json:"data"`
}

type ReportPayload struct {
	Service string `json:"service"`
	Status  int    `json:"status"`
}
