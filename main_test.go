package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLinkedinHandler_ValidURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/linkedin?url=https://www.linkedin.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(linkedinHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	expectedURL := "https://www.linkedin.com"
	if location := rr.Header().Get("Location"); location != expectedURL {
		t.Errorf("handler returned wrong location header: got %v want %v",
			location, expectedURL)
	}
}

func TestLinkedinHandler_InvalidURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/linkedin?url=invalid-url", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(linkedinHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expectedBody := "Invalid URL\n"
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("handler returned wrong response body: got %v want %v",
			body, expectedBody)
	}
}

func TestLinkedinHandler_WrongURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/linkedin?url=https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(linkedinHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expectedBody := "Invalid URL\n"
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("handler returned wrong response body: got %v want %v",
			body, expectedBody)
	}
}
