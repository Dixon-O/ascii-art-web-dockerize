package handlers

import (
	artprinter "ascii-art/Artprinter"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// PageData holds data passed to the template
type PageData struct {
	Text    string
	Banner  string
	Result  string
	Banners []string
}

var templates *template.Template

// InitTemplates loads HTML templates
func InitTemplates() error {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	return err
}

// HomeHandler handles GET /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	data := PageData{
		Banners: []string{"standard", "shadow", "thinkertoy"},
	}

	if templates == nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

// AsciiArtHandler handles POST /ascii-art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate banner
	validBanners := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
	if !validBanners[banner] {
		http.Error(w, "400 Bad Request: Invalid banner", http.StatusBadRequest)
		return
	}

	// Read banner file
	bannerPath := "banners/" + banner + ".txt"
	data, err := os.ReadFile(bannerPath)
	if err != nil {
		http.Error(w, "404 Not Found: Banner file not found", http.StatusNotFound)
		return
	}

	// Handle input
	input := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", "\n"), "\n", "\n"), "\n")

	// Check for only newlines
	if artprinter.CheckOnlyNewLines(input) {
		var result strings.Builder
		for i := range input[:len(input)-1] {
			result.WriteString("\n")
			if i == len(input[:len(input)-1])-1 {
				break
			}
		}
		w.Write([]byte(result.String()))
		return
	}

	// Generate ASCII art
	bannerLines := strings.Split(string(data), "\n")
	result := artprinter.PrintAsciiArt(input, bannerLines)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
