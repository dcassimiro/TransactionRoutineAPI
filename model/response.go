package model

type Response struct {
	Data interface{}            `json:"data,omitempty"`
	Meta map[string]interface{} `json:"metadata,omitempty"`
	Err  error                  `json:"error,omitempty"`
}
