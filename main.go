package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/Nr90/imgsim"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 args required: jpegdiff ./file1.jpg ./file2.jpeg")
		return
	}

	// Read
	file1, err := os.Open(os.Args[1])
	defer file1.Close()
	if err != nil {
		panic(err)
	}

	file2, err := os.Open(os.Args[2])
	defer file2.Close()
	if err != nil {
		panic(err)
	}

	// Decode
	img1, err := jpeg.Decode(file1)
	if err != nil {
		panic(err)
	}

	img2, err := jpeg.Decode(file2)
	if err != nil {
		panic(err)
	}

	// Hashing
	ahash1 := imgsim.AverageHash(img1)
	ahash2 := imgsim.AverageHash(img2)

	dhash1 := imgsim.DifferenceHash(img1)
	dhash2 := imgsim.DifferenceHash(img2)

	// Check
	if ahash1 == ahash2 || dhash1 == dhash2 {
		fmt.Println("Same")
		return
	}
	fmt.Println("Different")
}
