package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	expected := "Jai Mahishmathi ✊"
	if !strings.Contains(body, expected) {
		t.Errorf("expected body to contain %q, got %q", expected, body)
	}
}

func TestGetPort_WhenSet(t *testing.T) {
	_ = os.Setenv("PORT", "1234")
	port, err := getPort()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if port != "1234" {
		t.Errorf("expected port '1234', got '%s'", port)
	}
}
