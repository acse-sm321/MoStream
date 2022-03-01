package main

type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"Method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized  = Err{Error: "API not recognized, bad request", ErrorCode: "001"}
	ErrorRquestBodyParseFailed = Err{Error: "Request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults        = Err{Error: "Internal service error", ErrorCode: "003"}
)
