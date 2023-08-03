package models

type ErrorMsg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
