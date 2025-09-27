package request

import (
	"fmt"
	"io"
	"strings"

	"github.com/roshinys/httpGo/internal/headers"
)

type ParserState int

const (
	ParserInitialized ParserState = iota
	ParserParsingHeaders
	ParserDone
)

type Request struct {
	RequestLine RequestLine
	Headers     headers.Headers
	State       ParserState
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	buf := make([]byte, 1024)
	bufLen := 0
	req := Request{
		State:   ParserInitialized,
		Headers: make(headers.Headers), // Initialize the headers map
	}
	for req.State != ParserDone {
		if bufLen >= len(buf) {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf[:bufLen])
			buf = newBuf
		}
		n, err := reader.Read(buf[bufLen:])
		if err != nil && err != io.EOF {
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

		// Handle EOF
		if err == io.EOF {
			if req.State != ParserDone {
				return nil, fmt.Errorf("incomplete request: reached EOF before complete request")
			}
			break
		}
	}
	return &req, nil
}

func (r *Request) parse(data []byte) (int, error) {
	switch r.State {
	case ParserInitialized:
		n, err := r.parseRequestLine(data)
		if err != nil {
			return n, err
		}
		if r.State == ParserParsingHeaders {
			// Parse headers from the remaining data
			headerN, err := r.parseHeaderLine(data[n:])
			if err != nil {
				return n, err // Return only request line bytes if header parsing fails
			}
			return n + headerN, nil // Return total bytes consumed
		}
		return n, nil
	case ParserParsingHeaders:
		return r.parseHeaderLine(data)
	default:
		return 0, nil
	}
}

func (r *Request) parseHeaderLine(data []byte) (int, error) {
	totalConsumed := 0
	for {
		n, done, err := r.Headers.Parse(data[totalConsumed:])
		if err != nil {
			return totalConsumed, err
		}

		if n == 0 && !done {
			// Need more data
			return 0, nil
		}
		totalConsumed += n
		if done {
			// Headers parsing complete (found empty line)
			r.State = ParserDone
			break
		}
	}
	return totalConsumed, nil
}

func (r *Request) parseRequestLine(data []byte) (int, error) {
	reqStr := string(data)
	crlfIndex := strings.Index(reqStr, "\r\n")
	if crlfIndex == -1 {
		return 0, nil // Need more data
	}

	reqLine := reqStr[:crlfIndex]
	bytesConsumed := crlfIndex + 2 // includes \r\n

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
	r.State = ParserParsingHeaders
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
