package utils

// Response is a response format template for SUCCESS or FAIL response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateResponseOk return response success format with default value
func CreateResponseOk(message string, data interface{}) Response {
	return Response{"OK", message, data}
}

// CreateResponseFail return response success format with default value
func CreateResponseFail(message string, data interface{}) (res Response) {
	return Response{"FAIL", message, data}
}

// ResponseError is a response format template for ERROR response
type ResponseError struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// CreateResponseError return response error template format with default value
func CreateResponseError(message string, errors interface{}) (res ResponseError) {
	return ResponseError{"ERROR", message, errors}
}
