package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	tt := [...]struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "OK",
			queryParam:     "10",
			expectedStatus: http.StatusOK,
			expectedBody:   "20",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// NewRecorder returns an initialized [ResponseRecorder].
			// ResponseRecorder is an implementation of [http.ResponseWriter]
			rec := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodGet, "/double?v="+tc.queryParam, nil)

			doubleHandler(rec, req)

			require.Equal(t, tc.expectedStatus, rec.Code)
			body := rec.Body.String()

			require.Equal(t, tc.expectedBody, strings.TrimSpace(body))
		})
	}
}
