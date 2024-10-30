package foo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetFoo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Get Foo",
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"bar":"Hello, World!"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new gin router
			router := gin.Default()
			handler := &FooHandler{}
			router.GET("/foo", handler.GetFoo)

			// Create a new HTTP request
			req, err := http.NewRequest(http.MethodGet, "/foo", nil)
			assert.NoError(t, err)

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rr, req)

			// Check the status code
			assert.Equal(t, tt.expectedCode, rr.Code)

			// Check the response body
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestPostFoo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid Request",
			requestBody:  `{"id": 1, "bar": "Hello, World!"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"id":1,"bar":"Hello, World!"}`,
		},
		{
			name:         "Invalid Request",
			requestBody:  `{"id": "invalid", "bar": "Hello, World!"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"json: cannot unmarshal string into Go struct field Foo.id of type int"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new gin router
			router := gin.Default()
			handler := &FooHandler{}
			router.POST("/foo", handler.PostFoo)

			// Create a new HTTP request
			req, err := http.NewRequest(http.MethodPost, "/foo", bytes.NewBufferString(tt.requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rr, req)

			// Check the status code
			assert.Equal(t, tt.expectedCode, rr.Code)

			// Check the response body
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
		})
	}
}
