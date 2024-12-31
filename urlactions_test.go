package main

import (
	"os"
	"testing"
)

func TestCheckValidUrl(t *testing.T) {
	tests := []struct {
		rawurl string
		valid  bool
	}{
		{"https://www.google.com", true},
		{"ftp://example.com", true},
		{"invalid-url", false},
		{"http://", false},
	}

	for _, tt := range tests {
		t.Run(tt.rawurl, func(t *testing.T) {
			if got := checkValidUrl(tt.rawurl); got != tt.valid {
				t.Errorf("checkValidUrl() = %v, want %v", got, tt.valid)
			}
		})
	}
}

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		rawurl  string
		domain  string
		wantErr bool
	}{
		{"https://www.google.com", "google.com", false},
		{"https://sub.example.co.uk", "example.co.uk", false},
		{"invalid-url", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.rawurl, func(t *testing.T) {
			got, err := extractDomain(tt.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.domain {
				t.Errorf("extractDomain() = %v, want %v", got, tt.domain)
			}
		})
	}
}

func TestGoqrgen(t *testing.T) {
	// Test QR code generation
	url := "https://www.google.com"
	filename := "test_qr.png"
	goqrgen(url, true, filename)
	// Check if the file is created
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("QR code file was not created")
	}
	// Clean up
	os.Remove(filename)
}

func TestGoshorten(t *testing.T) {
	// Test URL shortening
	originalURL := "https://google.com"
	shortenedURL := goshorten(originalURL)

	if shortenedURL == originalURL {
		t.Errorf("goshorten() did not shorten the URL, got %v", shortenedURL)
	}
}