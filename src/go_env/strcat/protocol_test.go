package strcat

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	at := assert.New(t)

	tests := []struct {
		input  string
		code   int
		output string
	}{
		{"{\"strings\":[\"prefix\",\"suffix\"]}", http.StatusOK, "{\"string\":\"prefixsuffix\"}\n"},
		{"abc", http.StatusBadRequest, ""},
		{"{\"strings\":[\"1\",\"2\",\"3\"]}", http.StatusOK, "{\"string\":\"123\"}\n"},
	}

	for _, t := range tests {
		r, err := http.NewRequest("POST", "/strcat", strings.NewReader(t.input))
		at.Nil(err)

		w := httptest.NewRecorder()

		handler := New()
		handler.ServeHTTP(w, r)

		at.Equal(w.Code, t.code)
		if w.Code != http.StatusOK {
			continue
		}
		at.Equal(w.Body.String(), t.output)
	}
}

func TestHandler17(t *testing.T) {
	tests := []struct {
		input  string
		code   int
		output string
	}{
		{"{\"strings\":[\"prefix\",\"suffix\"]}", http.StatusOK, "{\"string\":\"prefixsuffix\"}\n"},
		{"abc", http.StatusBadRequest, ""},
		{"{\"strings\":[\"1\",\"2\",\"3\"]}", http.StatusOK, "{\"string\":\"123\"}\n"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			at := assert.New(t)

			r, err := http.NewRequest("POST", "/strcat", strings.NewReader(test.input))
			at.Nil(err)

			w := httptest.NewRecorder()

			handler := New()
			handler.ServeHTTP(w, r)

			at.Equal(w.Code, test.code)
			if w.Code != http.StatusOK {
				return
			}
			at.Equal(w.Body.String(), test.output)
		})
	}
}
