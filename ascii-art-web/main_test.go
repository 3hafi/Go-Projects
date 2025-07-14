package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestMainPageHandler tests the main page handler for correct status codes.
func TestMainPageHandler(t *testing.T) {
	// Test for 200 OK on root path
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	mainPageHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestAsciiArtHandler tests the ascii-art handler for various scenarios.
func TestAsciiArtHandler(t *testing.T) {
	// Test for BadRequest with missing form values
	req := httptest.NewRequest("POST", "/ascii-art", nil)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler did not return bad request for missing data: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Test for successful generation
	body := strings.NewReader("text=hello&banner=standard")
	req = httptest.NewRequest("POST", "/ascii-art", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for valid request: got %v want %v",
			status, http.StatusOK)
	}
	responseBody, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(responseBody), "<pre class=\"result\">") {
		t.Errorf("response body does not contain result block")
	}
}

// TestLoadBannerFile tests the banner file loading logic.
func TestLoadBannerFile(t *testing.T) {
	// Test loading a valid banner file
	bannerChars, err := LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("failed to load valid banner file 'standard.txt': %v", err)
	}
	if len(bannerChars) == 0 {
		t.Errorf("loaded banner file has no characters")
	}
}

// TestConvertToAsciiArt tests the text to ASCII art conversion logic.
func TestConvertToAsciiArt(t *testing.T) {
	bannerChars, err := LoadBannerFile("banners/standard.txt")
	if err != nil {
		t.Fatalf("could not load banner for testing: %v", err)
	}

	// Test with a simple word
	art := ConvertToAsciiArt("a", bannerChars)
	expectedArt := `        
        
  __ _  
 / _` + "`" + ` | 
| (_| | 
 \__,_| 
        
`
	if strings.TrimSpace(art) != strings.TrimSpace(expectedArt) {
		t.Errorf("ASCII art for 'a' was incorrect.\nGot:\n%s\nExpected:\n%s", art, expectedArt)
	}
}