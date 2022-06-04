package models

const (
	StatusCustomError = 321
)

//ResponseJsonApi Type
type ResponseJsonApi struct {
	Success      bool        `json:"success"`
	ResponseCode int64       `json:"responseCode"`
	Data         interface{} `json:"data"`
	Error        interface{} `json:"error"`
}
