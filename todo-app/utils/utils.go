package utils

import "encoding/json"

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func FormattedResponse(message string, data any) APIResponse {
	return APIResponse{
		Message: message,
		Data:    data,
	}
}

func PrettyJson(data any) string {
	result, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	return string(result)
}
