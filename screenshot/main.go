package main

import (
	"image/png"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Fatal(err)
	}

	fname := time.Now().UTC().Format("20060102150405") + ".png"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
