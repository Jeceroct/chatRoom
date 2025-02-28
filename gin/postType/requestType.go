package postType

import (
	"chatroom/myUser"
	"encoding/json"
)

type PostRequest struct {
	Type    string      `json:"type"`
	Context string      `json:"context"`
	Time    string      `json:"time"`
	From    myUser.User `json:"from"`
	Quote   QuoteType   `json:"quote"`
}

type QuoteType struct {
	From    myUser.User `json:"from"`
	Type    string      `json:"type"`
	Context string      `json:"context"`
	Time    string      `json:"time"`
}

type FileType struct {
	Title   string
	Context string
}

type Error struct {
	Type    string
	Code    string
	Context string
}

func ErrorMsg(code string, context string) Error {
	return Error{
		Type:    "error",
		Code:    code,
		Context: context,
	}
}

func ParseError(err Error) PostRequest {
	return PostRequest{
		Type:    err.Type,
		Context: err.Context,
		Time:    err.Code,
		From:    myUser.User{},
		Quote:   QuoteType{},
	}
}

func ParseFileContext(msg PostRequest) FileType {
	var file FileType
	json.Unmarshal([]byte(msg.Context), &file)
	return file
}
