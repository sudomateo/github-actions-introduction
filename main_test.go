package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/sudomateo/github-actions-introduction"
)

func TestHandlers(t *testing.T) {
	testCases := map[string]struct {
		handler    http.Handler
		request    *http.Request
		statusCode int
		message    string
	}{
		"OKHandler": {
			handler:    http.HandlerFunc(main.OKHandler),
			request:    httptest.NewRequest(http.MethodGet, "/", nil),
			statusCode: http.StatusOK,
			message:    http.StatusText(http.StatusOK),
		},
		"NotFoundHandler": {
			handler:    http.HandlerFunc(main.NotFoundHandler),
			request:    httptest.NewRequest(http.MethodGet, "/fake", nil),
			statusCode: http.StatusNotFound,
			message:    http.StatusText(http.StatusNotFound),
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			w := httptest.NewRecorder()

			testCase.handler.ServeHTTP(w, testCase.request)

			if testCase.statusCode != w.Result().StatusCode {
				t.Fatalf("expected %v, got %v", testCase.statusCode, w.Result().StatusCode)
			}

			var resp main.Response
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatal("invalid response body")
			}

			if testCase.message != resp.Message {
				t.Fatalf("expected %v, got %v", testCase.message, resp.Message)
			}
		})
	}
}
