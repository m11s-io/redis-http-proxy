package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthAcceptsBearerToken(t *testing.T) {
	token = "secret"
	called := false
	handler := auth(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusNoContent)
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", "Bearer secret")
	response := httptest.NewRecorder()
	handler(response, request)

	if !called || response.Code != http.StatusNoContent {
		t.Fatalf("handler called=%v, status=%d; want called=true, status=%d", called, response.Code, http.StatusNoContent)
	}
}

func TestAuthAcceptsQueryToken(t *testing.T) {
	token = "secret"
	called := false
	handler := auth(func(w http.ResponseWriter, r *http.Request) { called = true })

	request := httptest.NewRequest(http.MethodGet, "/?_token=secret", nil)
	response := httptest.NewRecorder()
	handler(response, request)

	if !called || response.Code != http.StatusOK {
		t.Fatalf("handler called=%v, status=%d; want called=true, status=%d", called, response.Code, http.StatusOK)
	}
}

func TestAuthRejectsInvalidToken(t *testing.T) {
	token = "secret"
	handler := auth(func(http.ResponseWriter, *http.Request) {
		t.Fatal("protected handler called with invalid token")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Authorization", "Bearer wrong")
	response := httptest.NewRecorder()
	handler(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("status=%d, want %d", response.Code, http.StatusUnauthorized)
	}
	if !strings.Contains(response.Body.String(), `"error":"Unauthorized"`) {
		t.Fatalf("body=%q, want JSON unauthorized error", response.Body.String())
	}
}

func TestHandleCommandRejectsInvalidJSON(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	response := httptest.NewRecorder()
	handleCommand(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status=%d, want %d", response.Code, http.StatusBadRequest)
	}
}

func TestHandleCommandRejectsNonPost(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	handleCommand(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status=%d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
}
