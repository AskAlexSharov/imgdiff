package main

import (
	"fmt"
	"github.com/Nr90/imgsim"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
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

	fmt.Printf("Diffeernce: %d%%\n", readAndGetDistance(fileName1, fileName2))
}

func readAndGetDistance(fileName1, fileName2 string) int {
	img1, img2 := readAndDecode(fileName1), readAndDecode(fileName2)
	return distance(img1, img2)
}

func parseUrl(rawurl string) (*url.URL, bool) {
	parsedUrl, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return parsedUrl, false
	}
	return parsedUrl, true
}

func mustGet(rawurl string) io.ReadCloser {
	resp, err := http.Get(rawurl)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Unexpected response from server. Status code: %d", resp.StatusCode))
	}
	return resp.Body
}

func readAndDecode(filePathOrUrl string) image.Image {
	// Read
	url1, isUrl := parseUrl(filePathOrUrl)

	if isUrl {
		body := mustGet(filePathOrUrl)
		defer body.Close()
		return decode(url1.Path, body)
	}

	f, err := os.Open(filePathOrUrl)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return decode(filePathOrUrl, f)
}

func decode(fileName string, f io.Reader) image.Image {
	var img image.Image
	var err error
	ext := filepath.Ext(fileName)

	if ext == ".jpeg" || ext == ".jpg" {
		img, err = jpeg.Decode(f)
		if err != nil {
			panic(err)
		}

	}

	if ext == ".png" {
		img, err = png.Decode(f)
		if err != nil {
			panic(err)
		}
	}

	return img
}

func distance(img1, img2 image.Image) int {
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
