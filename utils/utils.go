package utils

import (
	"strconv"
)

func GetLSPContentLengthFromHeader(header []byte) (int, error) {
	contentLenBytes := header[len("Content-Length: "):]

	contentLen, err := strconv.Atoi(string(contentLenBytes))
	if err != nil {
		return 0, err
	}

	return contentLen, nil
}
