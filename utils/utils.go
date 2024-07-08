package utils

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	return strings.TrimSpace(message)
}

func ParseHeaders(headers map[string]string) http.Header {
	headerMap := http.Header{}
	for key, value := range headers {
		headerMap.Add(key, value)
	}
	return headerMap
}
