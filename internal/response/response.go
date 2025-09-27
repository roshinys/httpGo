package response

import (
	"fmt"
	"io"
	"strconv"

	"github.com/roshinys/httpGo/internal/headers"
)

type StatusCode int

const (
	StatusOk       StatusCode = 200
	BadRequest     StatusCode = 400
	InternalServer StatusCode = 500
)

type HandlerError struct {
	StatusCode StatusCode
	Message    string
}

type Writer struct {
	ioWriter io.Writer
}

func NewWrite(w io.Writer) *Writer {
	return &Writer{
		ioWriter: w,
	}
}

func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
	var reason string
	switch statusCode {
	case StatusOk:
		reason = "OK"
	case BadRequest:
		reason = "Bad Request"
	case InternalServer:
		reason = "Internal Server Error"
	default:
		reason = "Unknown Error"
	}
	_, err := fmt.Fprintf(w.ioWriter, "HTTP/1.1 %d %s\r\n", statusCode, reason)
	return err
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	return headers.Headers{
		"Content-Length": strconv.Itoa(contentLen),
		"Connection":     "close",
		"Content-Type":   "text/html",
	}
}

func (w *Writer) WriteHeaders(headers headers.Headers) error {
	for k, v := range headers {
		_, err := fmt.Fprintf(w.ioWriter, "%s: %s\r\n", k, v)
		if err != nil {
			return err
		}
	}
	// End of headers
	_, err := fmt.Fprint(w.ioWriter, "\r\n")
	return err
}

func (w *Writer) WriteBody(p []byte) (int, error) {
	return w.ioWriter.Write(p)
}

func (w *Writer) WriteHandlerError(err HandlerError) error {
	body := []byte(err.Message) // Use the actual error message

	if writeErr := w.WriteStatusLine(err.StatusCode); writeErr != nil {
		return writeErr
	}

	headers := GetDefaultHeaders(len(body))
	if writeErr := w.WriteHeaders(headers); writeErr != nil {
		return writeErr
	}

	_, writeErr := w.WriteBody(body)
	return writeErr
}
