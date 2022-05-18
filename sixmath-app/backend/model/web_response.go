package model

type WebResponse struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"errors"`
}
