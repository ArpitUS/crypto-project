package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHashHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hash?data=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hashHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
