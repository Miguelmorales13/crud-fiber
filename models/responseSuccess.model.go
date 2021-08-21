package models

type ResponseSuccess struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
