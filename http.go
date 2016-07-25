package routeinjector

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

//Oops is an example error handler
func Oops() MyError {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

// WriteError sends a Error as a Response
func WriteError(w http.ResponseWriter, status int, message string) {
	// do the same as a regular response, but manually format as JSON
	WriteJSON(w, status, fmt.Sprintf("{\"error\":\"%s\"}", message))
}

// WriteJSON sends a JSON String as a Response
func WriteJSON(w http.ResponseWriter, status int, message string) {
	// write headers
	header := w.Header()
	header.Add("Content-Length", strconv.Itoa(len(message)))
	header.Add("Content-Type", "application/json")

	// write status code
	w.WriteHeader(status)

	// write data
	w.Write([]byte(message))
}
