package models

import "encoding/json"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewResponse(status int, data interface{}) Response {
	return Response{status, data}
}

func GenerateResponseJson(status int, data interface{}) string {
	response := NewResponse(status, data)
	responseJson, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	return string(responseJson)
}
