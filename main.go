package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"log"
	"os"
)

func main() {
	// Define command-line flags
	useMinified := true

	url := "http://localhost:8080"

	// Create a context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the website
	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent), // Capture the entire HTML content
	)
	if err != nil {
		log.Fatal(err)
	}

	// Prepend the DOCTYPE declaration
	capturedHTML := fmt.Sprintf("<!DOCTYPE html>%s", htmlContent)

	// Format the HTML content based on the flag
	var formattedHTML string
	if useMinified {
		// TODO: minifier dont work yet, it does not take multiline tags correctly
		//formattedHTML = minifyHTML(capturedHTML)
		formattedHTML = capturedHTML
	} else {
		//formattedHTML = prettifyHTML(capturedHTML)
	}

	// Save the HTML content to a file
	err = os.WriteFile("index.html", []byte(formattedHTML), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("index.html saved successfully")
}

func minifyHTML(inputHTML string) string {
	m := minify.New()

	// Configure HTML minification settings
	m.AddFunc("text/html", html.Minify)

	// Minify the HTML content
	minifiedHTML, err := m.String("text/html", inputHTML)
	if err != nil {
		log.Println("Error minifying HTML:", err)
		return inputHTML
	}

	return minifiedHTML
}

func minifyCSS(inputCSS string) string {
	m := minify.New()

	// Minify the CSS content
	minifiedCSS, err := m.String("text/css", inputCSS)
	if err != nil {
		log.Println("Error minifying CSS:", err)
		return inputCSS
	}

	return minifiedCSS
}

func minifyJS(inputJS string) string {
	m := minify.New()

	// Minify the JavaScript content
	minifiedJS, err := m.String("text/javascript", inputJS)
	if err != nil {
		log.Println("Error minifying JavaScript:", err)
		return inputJS
	}

	return minifiedJS
}
