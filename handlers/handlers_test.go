package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Change to project root for template/banner access
	os.Chdir("..")
	InitTemplates()
	os.Exit(m.Run())
}

func TestHomeHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	HomeHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "ASCII Art") {
		t.Error("Response should contain 'ASCII Art'")
	}
	if !strings.Contains(body, "standard") {
		t.Error("Response should contain banner options")
	}
}

func TestHomeHandler_NotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/invalid", nil)
	w := httptest.NewRecorder()

	HomeHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestHomeHandler_BadMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	HomeHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_Standard(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hi")
	form.Add("banner", "standard")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	AsciiArtHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "|") {
		t.Error("Response should contain ASCII art characters")
	}
}

func TestAsciiArtHandler_Shadow(t *testing.T) {
	form := url.Values{}
	form.Add("text", "A")
	form.Add("banner", "shadow")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	AsciiArtHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_Thinkertoy(t *testing.T) {
	form := url.Values{}
	form.Add("text", "B")
	form.Add("banner", "thinkertoy")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	AsciiArtHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_InvalidBanner(t *testing.T) {
	form := url.Values{}
	form.Add("text", "test")
	form.Add("banner", "invalid")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	AsciiArtHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_BadMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	w := httptest.NewRecorder()

	AsciiArtHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}
