package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	bannerHeight = 8
)

// AsciiArtData holds the data for the template
type AsciiArtData struct {
	Result string
}

func main() {
	// Serve static files (CSS, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Initial load of the page, or after a submission
	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	// We can pass data here if needed, for now, it's nil
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	log.Printf("Received text: %q, banner: %q", text, banner)

	if text == "" || banner == "" {
		http.Error(w, "Bad Request: Missing text or banner", http.StatusBadRequest)
		return
	}

	bannerFile := "banners/" + banner + ".txt"
	bannerChars, err := LoadBannerFile(bannerFile)
	if err != nil {
		http.Error(w, "Internal Server Error: Could not load banner file", http.StatusInternalServerError)
		log.Println("Error loading banner file:", err)
		return
	}

	art := ConvertToAsciiArt(text, bannerChars)

	data := AsciiArtData{Result: art}

	tmpl, err := template.ParseFiles("static/html/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

// LoadBannerFile reads the banner file and parses character templates
func LoadBannerFile(filename string) (map[rune][]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	bannerChars := make(map[rune][]string)
	lines := strings.Split(string(content), "\n")

	// standard.txt works with the original logic.
	// shadow.txt and thinkertoy.txt are missing the 'space' character and start with '!'.
	if strings.Contains(filename, "shadow") || strings.Contains(filename, "thinkertoy") {
		// For these banners, the file content starts with ASCII 33 ('!').
		// We need to manually handle space and adjust the loop for other characters.
		
		// Provide a default, empty representation for the space character.
		bannerChars[' '] = make([]string, bannerHeight)
		for i := 0; i < bannerHeight; i++ {
			// Use a fixed width for the default space character
			bannerChars[' '][i] = "        "
		}

		// Load the rest of the characters, starting from '!' (33).
		for char := rune(33); char < 127; char++ {
			// The position in the file is based on a 0-index from character 33.
			fileIndex := int(char) - 33
			start := fileIndex * (bannerHeight + 1)
			
			if start+bannerHeight > len(lines) {
				break
			}
			charLines := lines[start : start+bannerHeight]
			bannerChars[char] = charLines
		}
	} else {
		// The original logic for standard.txt.
		for char := rune(32); char < 127; char++ {
			start := (int(char) - 32) * (bannerHeight + 1)
			
			if start+bannerHeight > len(lines) {
				break
			}
			charLines := lines[start : start+bannerHeight]
			bannerChars[char] = charLines
		}
	}

	// Final check to ensure space is there.
	if _, exists := bannerChars[' ']; !exists {
		return nil, fmt.Errorf("space character not found in banner file")
	}

	return bannerChars, nil
}

// ConvertToAsciiArt transforms the input text into ASCII art
func ConvertToAsciiArt(text string, bannerChars map[rune][]string) string {
	// Handle empty input
	if text == "" {
		return ""
	}

	// Handle newline characters
	lines := strings.Split(text, "\r\n")
	
	var result strings.Builder

	for lineIndex, line := range lines {
		// Skip empty lines
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Create each line of the ASCII art
		for height := 0; height < bannerHeight; height++ {
			for _, char := range line {
				// Use space as default if character not found
				charTemplate, exists := bannerChars[char]
				if !exists {
					charTemplate = bannerChars[' ']
				}
				
				// Defensive check to prevent index out of range
				if height < len(charTemplate) {
					result.WriteString(charTemplate[height])
				}
			}
			result.WriteString("\n")
		}

		// Add an extra newline between lines except for the last line
		if lineIndex < len(lines)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}