package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/kbinani/screenshot"
	"image/jpeg"
	"net/http"
)

func main() {
	ip := flag.String("ip", "", "IP address to upload the screenshot")
	port := flag.String("port", "", "Port to upload the screenshot")
	flag.Parse()

	if *ip == "" || *port == "" {
		fmt.Println("Please provide an IP address and port using the -ip and -port flags.")
		return
	}

	// Capture the screenshot
	screenBounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(screenBounds)
	if err != nil {
		fmt.Printf("Error capturing screenshot: %v\n", err)
		return
	}

	// Save the screenshot as a JPEG file
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 75})
	if err != nil {
		fmt.Printf("Error encoding screenshot: %v\n", err)
		return
	}

	// Upload the JPEG file to the specified IP and port
	url := fmt.Sprintf("http://%s:%s/upload", *ip, *port)
	resp, err := http.Post(url, "image/jpeg", &buf)
	if err != nil {
		fmt.Printf("Error uploading screenshot: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Screenshot uploaded successfully.")
}
