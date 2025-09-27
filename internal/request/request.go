package request

import (
	"fmt"
	"io"
	"strings"
)

type ParserState int

const (
	ParserInitialized ParserState = iota
	ParserDone
)

type Request struct {
	RequestLine RequestLine
	State       ParserState
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	buf := make([]byte, 1024)
	// GET /ckenf/jenf HTTP 1.1
	bufLen := 0
	req := Request{
		State: ParserInitialized,
	}
	for req.State != ParserDone {
		if bufLen >= len(buf) {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf[:bufLen])
			buf = newBuf
		}
		// if buffer length exceeds
		n, err := reader.Read(buf[bufLen:])
		if err != nil {
			return nil, err
		}
		bufLen += n
		readN, err := req.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}
		if readN > 0 {
			copy(buf, buf[readN:bufLen])
			bufLen -= readN
		}
		// HANDLE EOF
		if err == io.EOF {
			if req.State != ParserDone {
				return nil, fmt.Errorf("incomplete request: reached EOF before complete request line")
			}
			break
		}
	}
	return &req, nil
}

func parseRequestLine(reqBytes []byte) (RequestLine, error) {
	reqStr := string(reqBytes)

	lines := strings.Split(reqStr, "\r\n")

	if len(lines) == 1 {
		return RequestLine{}, nil
	}
	reqLine := lines[0]

	parts := strings.Split(reqLine, " ")
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("invalid request line format: expected 3 parts, got %d", len(parts))
	}
	method := parts[0]
	requestTarget := parts[1]
	httpVersion := parts[2]

	if !isValidMethod(method) {
		return RequestLine{}, fmt.Errorf("invalid method: %s (must contain only capital alphabetic characters)", method)
	}

	version, err := parseHttpVersion(httpVersion)
	if err != nil {
		return RequestLine{}, err
	}

	return RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   version,
	}, nil

}

func (r *Request) parse(data []byte) (int, error) {
	if r.State != ParserInitialized {
		return 0, nil
	}
	reqStr := string(data)
	crlfIndex := strings.Index(reqStr, "\r\n")
	if crlfIndex == -1 {
		return 0, nil
	}
	lines := strings.Split(reqStr, "\r\n")
	if len(lines) == 1 {
		return 0, nil
	}
	reqLine := reqStr[:crlfIndex]
	bytesConsumed := crlfIndex + 2 //includes \r\n

	parts := strings.Split(reqLine, " ")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid request line format: expected 3 parts, got %d", len(parts))
	}
	method := parts[0]
	requestTarget := parts[1]
	httpVersion := parts[2]

	if !isValidMethod(method) {
		return 0, fmt.Errorf("invalid method: %s (must contain only capital alphabetic characters)", method)
	}

	version, err := parseHttpVersion(httpVersion)
	if err != nil {
		return 0, err
	}

	r.RequestLine = RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   version,
	}

	r.State = ParserDone
	return bytesConsumed, nil
}

func isValidMethod(method string) bool {
	if method == "" {
		return false
	}
	for _, char := range method {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}

func parseHttpVersion(httpVersionStr string) (string, error) {
	// Expected format: "HTTP/1.1"
	if !strings.HasPrefix(httpVersionStr, "HTTP/") {
		return "", fmt.Errorf("invalid HTTP version format: %s (must start with 'HTTP/')", httpVersionStr)
	}

	version := strings.TrimPrefix(httpVersionStr, "HTTP/")
	if version != "1.1" {
		return "", fmt.Errorf("unsupported HTTP version: %s (only HTTP/1.1 is supported)", version)
	}

	return version, nil
}
