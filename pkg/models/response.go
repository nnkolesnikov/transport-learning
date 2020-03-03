package models

// Response ...
type Response struct {
	Error       bool              `json:error`
	ErrorText   string            `json:error_text`
	Data        *Data             `json:data`
	customError map[string]string `json:custom_error`
}

// Data ...
type Data struct {
	Res int
}
