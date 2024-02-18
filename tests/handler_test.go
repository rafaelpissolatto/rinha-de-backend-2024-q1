package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"rinha-backend-2024q1/internal/server"
	"testing"
)

func TestHandlers(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(s.RegisterRoutes())
	defer server.Close()

	tests := []struct {
		name     string
		method   string
		url      string
		body     io.Reader
		expected string
		status   int
	}{
		{
			name:     "Health",
			method:   "GET",
			url:      "/health",
			body:     nil,
			expected: `{"message":"It's healthy"}`,
			status:   http.StatusOK,
		},
		{
			name:     "Create Transaction",
			method:   "POST",
			url:      "/clientes/1/transacoes",
			body:     nil,
			expected: `{"message":"Transaction created"}`,
			status:   http.StatusOK,
		},
		{
			name:     "Get Extract",
			method:   "GET",
			url:      "/clientes/1/extrato",
			body:     nil,
			expected: `{"message":"Extract retrieved"}`,
			status:   http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, server.URL+tt.url, tt.body)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Errorf("expected status %d, got %d", tt.status, resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(body) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, string(body))
			}
		})
	}
}
