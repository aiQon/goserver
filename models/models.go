package models

type InternalApiError struct {
	HttpCode   int
	StatusCode int
}

type ExternalApiError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
