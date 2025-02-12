package postType

import (
	"chatroom/myUser"
	"encoding/json"
)

type PostRequest struct {
	Type    string
	Context string
	Time    string
	From    myUser.User
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
	}
}

func ParseFileContext(msg PostRequest) FileType {
	var file FileType
	json.Unmarshal([]byte(msg.Context), &file)
	return file
}
