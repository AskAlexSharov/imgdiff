package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Nr90/imgsim"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 args required: imgdiff ./file1.jpg ./file2.jpeg")
		return
	}

	fmt.Printf("Diffeernce: %d%%\n", readAndGetDistance(os.Args[1], os.Args[2]))
}

func readAndGetDistance(fileName1, fileName2 string) int {
	return distance(<-readAndDecode(fileName1), <-readAndDecode(fileName2))
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

func readAndDecode(filePathOrUrl string) chan image.Image {
	res := make(chan image.Image)
	go func() {
		defer close(res)

		// Read
		parsedUrl, isUrl := parseUrl(filePathOrUrl)

		if isUrl {
			body := mustGet(filePathOrUrl)
			defer body.Close()
			res <- decode(parsedUrl.Path, body)
			return

		}

		f, err := os.Open(filePathOrUrl)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		res <- decode(filePathOrUrl, f)
		return
	}()

	return res
}

func decode(fileName string, f io.Reader) image.Image {
	ext := filepath.Ext(fileName)

	if ext == ".jpeg" || ext == ".jpg" {
		img, err := jpeg.Decode(f)
		if err != nil {
			panic(err)
		}
		return img
	}

	if ext == ".png" {
		img, err := png.Decode(f)
		if err != nil {
			panic(err)
		}
		return img
	}

	log.Fatalf("Not supported file extension: %s", ext)
	return nil
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
