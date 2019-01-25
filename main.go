package main

import (
	"fmt"
	"github.com/Nr90/imgsim"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 args required: imgdiff ./file1.jpg ./file2.jpeg")
		return
	}

	fileName1 := os.Args[1]
	fileName2 := os.Args[2]

	fmt.Printf("Diffeernce: %d%%\n", distance(fileName1, fileName2))
}

func distance(fname1, fname2 string) (distance int) {
	// Read
	file1, err := os.Open(fname1)
	defer file1.Close()
	if err != nil {
		panic(err)
	}

	file2, err := os.Open(fname2)
	defer file2.Close()
	if err != nil {
		panic(err)
	}

	// Decode
	var img1, img2 image.Image

	ext1 := filepath.Ext(fname1)
	ext2 := filepath.Ext(fname2)
	if (ext1 == ".jpeg" || ext1 == ".jpg") && (ext2 == ".jpeg" || ext2 == ".jpg") {
		img1, img2 = decodeJpeg(file1, file2)
	}

	if ext1 == ".png" && ext2 == ".png" {
		img1, img2 = decodePng(file1, file2)
	}

	// Hashing
	ahash1 := imgsim.AverageHash(img1)
	ahash2 := imgsim.AverageHash(img2)

	dhash1 := imgsim.DifferenceHash(img1)
	dhash2 := imgsim.DifferenceHash(img2)

	// distance
	avgDistance := imgsim.Distance(ahash1, ahash2)
	diffDistance := imgsim.Distance(dhash1, dhash2)

	return (avgDistance + diffDistance) % 128
}

func decodeJpeg(f1 io.Reader, f2 io.Reader) (image.Image, image.Image) {
	img1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}

	img2, err := jpeg.Decode(f2)
	if err != nil {
		panic(err)
	}

	return img1, img2
}

func decodePng(f1 io.Reader, f2 io.Reader) (image.Image, image.Image) {
	img1, err := png.Decode(f1)
	if err != nil {
		panic(err)
	}

	img2, err := png.Decode(f2)
	if err != nil {
		panic(err)
	}

	return img1, img2
}
