package loader

import (
	"context"
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
	"sync"
)

// Result image reading/downloading and decoding result
type Result struct {
	Img image.Image
	Err error
}

// ImagesAsync download/read and decode files in parallel
func ImagesAsync(ctx context.Context, filePathOrURLs ...string) chan Result {
	resultCh := make(chan Result)
	wg := sync.WaitGroup{}
	wg.Add(len(filePathOrURLs))
	for _, filePathOrURL := range filePathOrURLs {
		go func(filePathOrURL string) {
			defer wg.Done()
			select {
			case resultCh <- Image(ctx, filePathOrURL):
			case <-ctx.Done():
				resultCh <- Result{Err: ctx.Err()}
			}
		}(filePathOrURL)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

// Image download/read and decode file
func Image(ctx context.Context, filePathOrURL string) Result {
	// Read
	var img io.Reader
	var fileName string

	if parsedURL, ok := parseURL(filePathOrURL); ok {
		resp, closer, err := getByURL(ctx, filePathOrURL)
		if err != nil {
			return Result{Err: err}
		}
		defer closer()

		img = resp.Body
		fileName = parsedURL.Path
	} else { // then it's file
		var err error
		if img, err = os.Open(filepath.Clean(filePathOrURL)); err != nil {
			return Result{Err: err}
		}
		fileName = filePathOrURL
	}

	return decode(fileName, img)
}

func parseURL(rawurl string) (*url.URL, bool) {
	parsedURL, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return parsedURL, false
	}
	return parsedURL, true
}

func decode(fileName string, f io.Reader) Result {
	ext := filepath.Ext(fileName)

	result := Result{}
	if ext == ".jpeg" || ext == ".jpg" {
		result.Img, result.Err = jpeg.Decode(f)
		return result
	}

	if ext == ".png" {
		result.Img, result.Err = png.Decode(f)
		return result
	}

	result.Err = fmt.Errorf("not supported file extension: %s", ext)
	return result
}

func getByURL(ctx context.Context, url string) (*http.Response, func(), error) {
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return nil, nil, reqErr
	}

	resp, respErr := http.DefaultClient.Do(req.WithContext(ctx))
	closer := func() {
		if resp == nil || resp.Body == nil {
			return
		}

		if err := resp.Body.Close(); err != nil {
			log.Printf("WARN: can't close http request body by reason %s", err.Error())
		}
	}

	if respErr != nil {
		return nil, closer, respErr
	}
	if resp.StatusCode != http.StatusOK {
		closer()
		return nil, closer, fmt.Errorf("unexpected HTTP response code: %d, of file: %s", resp.StatusCode, url)
	}

	return resp, closer, nil
}
