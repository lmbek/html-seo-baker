package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func main() {
	// Create a context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the website
	err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:8080"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Capture a screenshot of the entire page
	var screenshot []byte
	err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&screenshot))
	if err != nil {
		log.Fatal(err)
	}

	// Save the screenshot as an HTML file
	err = os.WriteFile("screenshot-of-website.png", screenshot, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("index.html saved successfully")
}
