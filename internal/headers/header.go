package headers

import (
	"fmt"
	"regexp"
	"strings"
)

type Headers map[string]string

func NewHeader() Headers {
	h := Headers{}
	return h
}

func (h *Headers) SetHeaders(key, value string) {
	if v, ok := (*h)[key]; ok {
		(*h)[key] = fmt.Sprintf("%s,%s", v, value)
	} else {
		(*h)[key] = value
	}
}

func (h *Headers) Parse(data []byte) (n int, done bool, err error) {
	headerStr := string(data)
	crlfIndex := strings.Index(headerStr, "\r\n")
	if crlfIndex == -1 {
		return 0, false, nil
	}
	if crlfIndex == 0 {
		return 2, true, nil
	}
	headerLine := headerStr[:crlfIndex]
	colonIndex := strings.Index(headerLine, ":")
	if colonIndex == -1 {
		return 0, false, fmt.Errorf("invalid header format: missing colon")
	}

	// Split into key and value
	key := strings.ToLower(strings.TrimSpace(headerLine[:colonIndex]))
	value := strings.TrimSpace(headerLine[colonIndex+1:])

	// Validate key is not empty
	if key == "" {
		return 0, false, fmt.Errorf("invalid header format: empty key")
	}
	if *h == nil {
		*h = NewHeader()
	}
	if !validateHeaderKey(key) {
		return 0, false, fmt.Errorf("invalid header key")
	}
	h.SetHeaders(key, value)
	return crlfIndex + 2, false, nil
}

func validateHeaderKey(keyName string) bool {
	allowedPattern := regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_` + "`" + `|~]+$`)
	return len(keyName) > 0 && allowedPattern.MatchString(keyName)

}
