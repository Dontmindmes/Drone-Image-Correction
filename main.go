package main

import (
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("test.png")
	if err != nil {
		log.Fatal("Failed to load image.", err)
	}

	gamacor := imaging.AdjustGamma(src, 0.75)
	bright := imaging.AdjustBrightness(gamacor, 3)
	contrast := imaging.AdjustContrast(bright, 4)

	err = imaging.Save(contrast, "out.png")
	if err != nil {
		log.Fatal("Could not save the image.", err)
	}
}
