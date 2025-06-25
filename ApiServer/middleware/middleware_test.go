package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestLogging_NextHandlerCalled(t *testing.T) {
	called := false

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	var buffer bytes.Buffer
	loggingMiddlewareToTest := Logging(&buffer)(dummyHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	loggingMiddlewareToTest.ServeHTTP(rr, req)

	if !called {
		t.Error("expected handler to be called")
	}
}

func TestLogging_LoggingOrder(t *testing.T) {
	var buffer bytes.Buffer
	handlerText := "dummyHandler text"

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(&buffer, handlerText)
	})

	loggingMiddlewareToTest := Logging(&buffer)(dummyHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	loggingMiddlewareToTest.ServeHTTP(rr, req)

	logLines := strings.Split(strings.TrimSpace(buffer.String()), "\n")

	if len(logLines) != 3 {
		t.Fatalf("expected 3 log lines, got %d: %v", len(logLines), logLines)
	}

	if !strings.Contains(logLines[0], "Start") {
		t.Errorf("expected to have Start, got: %v", logLines[0])
	}

	if logLines[1] != handlerText {
		t.Errorf("Expected to have `%s` got: %v", handlerText, logLines[1])
	}

	if !strings.Contains(logLines[2], "Finish") {
		t.Errorf("expected to have Finish, got: %v", logLines[2])
	}
}

func TestLoggingMiddleware_StartFinishAndMethodPath(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"GET root", "GET", "/", http.StatusOK},
		{"POST endpoint", "POST", "/api/item", http.StatusCreated},
		{"DELETE endpoint", "DELETE", "/user/123", http.StatusNoContent},
		{"PUT endpoint", "PUT", "/user/123", http.StatusNoContent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			middleware := Logging(&buffer)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			})

			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			middleware(handler).ServeHTTP(rec, req)

			logLines := strings.Split(strings.TrimSpace(buffer.String()), "\n")

			if len(logLines) != 2 {
				t.Fatalf("expected 2 log lines, got %d: %v", len(logLines), logLines)
			}

			// start log
			{
				hasStart := strings.Contains(logLines[0], "Start")
				hasMethod := strings.Contains(logLines[0], tt.method)
				hasPath := strings.Contains(logLines[0], tt.path)

				if !(hasStart && hasMethod && hasPath) {
					t.Errorf("unexpected start log: %s \nexpected to have Start, Method, Path\n"+
						"got: start(%t), method(%t), path(%t)",
						logLines[0], hasStart, hasMethod, hasPath)
				}
			}

			// finish log
			{
				hasStart := strings.Contains(logLines[1], "Finish")
				hasMethod := strings.Contains(logLines[1], tt.method)
				hasPath := strings.Contains(logLines[1], tt.path)

				if !(hasStart && hasMethod && hasPath) {
					t.Errorf("unexpected start log: %s \nexpected to have Finish, Method, Path\n"+
						"got: start(%t), method(%t), path(%t)",
						logLines[0], hasStart, hasMethod, hasPath)
				}
			}
		})
	}
}

func TestCORS(t *testing.T) {
	expectedAllowedMethods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}
	expectedAllowedOrigins := "*"

	tests := []struct {
		name              string
		method            string
		path              string
		handlerStatusCode int
		desiredStatusCode int
	}{
		{"GET root", "GET", "/", http.StatusOK, http.StatusOK},
		{"POST endpoint", "POST", "/api/item", http.StatusCreated, http.StatusCreated},
		{"DELETE user", "DELETE", "/user/123", http.StatusOK, http.StatusOK},
		{"Preflight", "OPTIONS", "/", http.StatusOK, http.StatusNoContent},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.handlerStatusCode)
			})

			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			CORS(handler).ServeHTTP(rec, req)

			if rec.Result().StatusCode != tt.desiredStatusCode {
				t.Errorf("expected status code %d, got: %d", tt.desiredStatusCode, rec.Result().StatusCode)
			}

			responseAllowedMethods := strings.Split(strings.TrimSpace(rec.Result().Header.Get("Access-Control-Allow-Methods")), ", ")
			responseAllowedOrigins := rec.Result().Header.Get("Access-Control-Allow-Origin")
			if len(responseAllowedMethods) != len(expectedAllowedMethods) {
				t.Errorf("expected %d methods, got: %d", len(expectedAllowedMethods), len(responseAllowedMethods))
			}

			for _, allowedMethod := range expectedAllowedMethods {
				if !slices.Contains(responseAllowedMethods, allowedMethod) {
					t.Errorf("expected to have method %s, got: %s", allowedMethod, responseAllowedMethods)
				}
			}

			if responseAllowedOrigins != expectedAllowedOrigins {
				t.Errorf("expected allowed origin to be %s, got: %s", expectedAllowedOrigins, responseAllowedOrigins)
			}
		})
	}
}
