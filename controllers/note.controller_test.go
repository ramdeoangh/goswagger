package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {

	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{

		{
			description:  "get HTTP status 200",
			route:        "/",
			expectedCode: 200,
		},
	}

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)

		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

}
