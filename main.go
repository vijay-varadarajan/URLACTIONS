package main

import (
	"fmt"

	goqr "github.com/skip2/go-qrcode"
	goshort "github.com/subosito/shorturl"
)

func main() {
	fmt.Println("TODO: URL CLI-app")

	originalURL := "https://google.com"
	qrcode, err := goqr.New(originalURL, goqr.Low)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode.ToString(true))
	qrcode.WriteFile(256, "qr.png")

	shorturl, err := goshort.Shorten(originalURL, "tinyurl")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(shorturl))
}