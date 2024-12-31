package main

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"strings"

	goclipboard "github.com/atotto/clipboard"
	goqr "github.com/skip2/go-qrcode"
	goshort "github.com/subosito/shorturl"
)

func extractDomain(rawurl string) (string, error) {
    parsedUrl, err := url.Parse(rawurl)
    if err != nil {
        return "", err
    }
    domainParts := strings.Split(parsedUrl.Hostname(), ".")
    if len(domainParts) < 2 {
        return "", fmt.Errorf("invalid domain")
    }
    // Handle cases like co.uk, com.au, etc.
    if len(domainParts) > 2 && (domainParts[len(domainParts)-2] == "co" || domainParts[len(domainParts)-2] == "com") {
        domain := strings.Join(domainParts[len(domainParts)-3:], ".")
        return domain, nil
    }
    domain := strings.Join(domainParts[len(domainParts)-2:], ".")
    return domain, nil
}

func checkValidUrl(rawurl string) bool {
	parsedUrl, err := url.ParseRequestURI(rawurl)
	if err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return false
	}
	return true
}

func govalidateURL(rawurl string) {
    if !checkValidUrl(rawurl) {
        fmt.Println("Invalid URL format")
        return
    }

    resp, err := http.Get(rawurl)
    if err != nil {
        fmt.Println("Server is not responding:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println("URL is valid and server is responding")
    } else {
        fmt.Printf("Server responded with status code: %d\n", resp.StatusCode)
    }
}

func goqrgen(url string, savefile bool, filename string) {
	if !checkValidUrl(url) {
		fmt.Println("Invalid URL")
		return
	}

	if filename == "" {
		domain, err := extractDomain(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		filename = domain + ".png"
	} else if !strings.HasSuffix(filename, ".png") {
		filename += ".png"
	}

	qrcode, err := goqr.New(url, goqr.High)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode.ToString(true))

	if savefile {
		err := qrcode.WriteFile(128, filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("QR code saved as " + filename)
	}
}

func goshorten(url string) string {
	if !checkValidUrl(url) {
		fmt.Println("Invalid URL")
		os.Exit(1)
	}

	shorturl, err := goshort.Shorten(url, "tinyurl")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	err = goclipboard.WriteAll(string(shorturl))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	result := string(shorturl)
	fmt.Println(result)
	fmt.Println("Shortened URL copied to clipboard")
	return string(result)
}