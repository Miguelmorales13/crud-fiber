package models

import "time"

type ResponseError struct {
	Code       int       `json:"code"`
	Path       string    `json:"path"`
	Method     string    `json:"method"`
	TimeStamps time.Time `json:"timestamps"`
	Message    string    `json:"message"`
}
