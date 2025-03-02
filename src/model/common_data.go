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

type ConfigParams struct {
	Name      string `json:"name"`
	ConfigKey string `json:"config_key"`
}

type LogParams struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
	Level     int    `json:"level"`
	Tag       string `json:"tag"`
	Log       string `json:"log"`
}

type RegisterParams struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Prefix string `json:"prefix"`
}
