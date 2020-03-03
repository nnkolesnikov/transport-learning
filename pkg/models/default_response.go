package models

// DefaultResponse ...
type DefaultResponse struct {
	Error       bool              `json:error`
	ErrorText   string            `json:error_text`
	Data        *Data             `json:data`
	customError map[string]string `json:custom_error`
}

// Data ...
type Data struct {
	Res bool
}
