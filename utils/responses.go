package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RespOk struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type RespBad struct {
	StatusCode int
	Message    string `json:"message"`
}

type Response struct {
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
	statusCode    int
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateResponseDefault(w http.ResponseWriter) Response {
	return Response{
		Status:        "ok",
		statusCode:    200,
		responseWrite: w,
		contentType:   "application/json",
	}
}

func (res *Response) Send() {
	res.responseWrite.Header().Set("Content-Type", res.contentType)
	res.responseWrite.WriteHeader(res.statusCode)
	dataOutPut, _ := json.Marshal(&res)
	fmt.Fprintln(res.responseWrite, string(dataOutPut))
}

func SendResponse(w http.ResponseWriter, dataSend RespOk) {
	response := CreateResponseDefault(w)
	response.Message = dataSend.Message
	response.Data = dataSend.Data
	response.Send()
}
func SendResponseVoid(w http.ResponseWriter) {
	response := CreateResponseDefault(w)
	response.statusCode = 200
	response.Message = "No existen los recursos requeridos"
	response.Data = make([]string, 0)
	response.Send()
}

func CreatedResponse(w http.ResponseWriter, dataSend RespOk) {
	response := CreateResponseDefault(w)
	response.statusCode = 201
	response.Message = dataSend.Message
	response.Data = dataSend.Data
	response.Send()
}

func BadResponse(w http.ResponseWriter, dataSend RespBad) {
	response := CreateResponseDefault(w)
	response.Status = "error"
	response.Message = dataSend.Message
	response.statusCode = dataSend.StatusCode
	response.Send()
}
