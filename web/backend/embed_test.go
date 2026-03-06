package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSPARouteFallsBackToIndex(t *testing.T) {
	mux := http.NewServeMux()
	registerEmbedRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/providers", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}
	if !strings.Contains(rr.Body.String(), `<div id="root"></div>`) {
		t.Fatalf("response does not look like index.html")
	}
}

func TestUnknownAPIPathStays404(t *testing.T) {
	mux := http.NewServeMux()
	registerEmbedRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/api/not-found", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusNotFound)
	}
}

func TestMissingAssetStays404(t *testing.T) {
	mux := http.NewServeMux()
	registerEmbedRoutes(mux)

	req := httptest.NewRequest(http.MethodGet, "/assets/not-found.js", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusNotFound)
	}
}
