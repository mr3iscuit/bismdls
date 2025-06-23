package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mr3iscuit/bismdls/utils"
)

type BaseMessage struct {
	Method string `json:"method,omitempty"`
}

func EncodeMessage(msg any) []byte {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	reply := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
	return []byte(reply)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("did not find separator")
	}

	contentLen, err := utils.GetLSPContentLengthFromHeader(header)
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage

	if err := json.Unmarshal(content[:contentLen], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content, nil
}

func Split(data []byte, atEof bool) (int, []byte, error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		// not ready yet
		return 0, nil, nil
	}
	contentLen, err := utils.GetLSPContentLengthFromHeader(header)
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLen {
		// still not ready
		return 0, nil, nil
	}

	totalLen := len(header) + 4 + contentLen

	return totalLen, data[:totalLen], nil
}
