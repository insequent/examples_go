package main

import (
	"image"
	//"image/color"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/gift"
)

func main() {
	// 1. Load the image
	src := loadImage("images/nathan_fillion.png")

	// 2. Create a GIFT filter list and add filter(s)
	g := gift.New(
		// BLUE!
		gift.Colorize(240, 50, 100),
	)

	// 3. Create a destination image
	dst := image.NewRGBA(g.Bounds(src.Bounds()))

	// 4. Use Draw func to apply the filters to src and store the result in dst
	g.Draw(dst, src)

	// 5. Write the destination image
	saveImage("images/blue_nathan.png", dst)
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}
