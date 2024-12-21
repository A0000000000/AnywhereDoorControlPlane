package model

type Result struct {
	Name   string `json:"name"`
	Target string `json:"target"`
	Data   string `json:"data"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}
