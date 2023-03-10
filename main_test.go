package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/sudomateo/github-actions-introduction"
)

func TestExampleApp(t *testing.T) {
	buf := new(bytes.Buffer)

	app := main.App{
		Log: log.New(buf, "test: ", 0),
	}

	testCases := map[string]struct {
		handler    http.Handler
		request    *http.Request
		statusCode int
		message    string
		log        string
	}{
		"OKHandler": {
			handler:    http.HandlerFunc(app.OKHandler),
			request:    httptest.NewRequest(http.MethodGet, "/", nil),
			statusCode: http.StatusOK,
			message:    http.StatusText(http.StatusOK),
			log:        "test: /\n",
		},
		"NotFoundHandler": {
			handler:    http.HandlerFunc(app.NotFoundHandler),
			request:    httptest.NewRequest(http.MethodGet, "/fake", nil),
			statusCode: http.StatusNotFound,
			message:    http.StatusText(http.StatusNotFound),
			log:        "test: /fake\n",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			// Reset the buffer from previous tests.
			buf.Reset()

			// Make the HTTP request and record its result.
			w := httptest.NewRecorder()
			testCase.handler.ServeHTTP(w, testCase.request)

			// Assert that the status code is what's expected.
			if testCase.statusCode != w.Result().StatusCode {
				t.Fatalf("expected %v, got %v", testCase.statusCode, w.Result().StatusCode)
			}

			// Assert that the response is valid JSON.
			var resp main.Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatal("invalid response body")
			}

			// Assert the response message is what's expected.
			if testCase.message != resp.Message {
				t.Fatalf("expected %v, got %v", testCase.message, resp.Message)
			}

			// Assert that the handler logged correctly.
			if testCase.log != buf.String() {
				t.Fatalf("expected %q, got %q", testCase.log, buf.String())
			}
		})
	}
}
