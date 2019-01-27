package loader

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

type Result struct {
	Img image.Image
	Err error
}

func ImagesAsync(ctx context.Context, filePathOrUrls ...string) chan Result {
	resultCh := make(chan Result)
	wg := sync.WaitGroup{}
	wg.Add(len(filePathOrUrls))
	for _, filePathOrUrl := range filePathOrUrls {
		go func(filePathOrUrl string) {
			defer wg.Done()
			select {
			case resultCh <- Image(ctx, filePathOrUrl):
			case <-ctx.Done():
				resultCh <- Result{Err: ctx.Err()}
			}
		}(filePathOrUrl)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func Image(ctx context.Context, filePathOrUrl string) Result {
	// Read
	var imgReader io.Reader
	var fileName string

	parsedUrl, isUrl := parseUrl(filePathOrUrl)
	if isUrl {
		resp, err := getByUrl(ctx, filePathOrUrl)
		if err != nil {
			return Result{Err: err}
		}
		defer resp.Body.Close()
		imgReader = resp.Body
		fileName = parsedUrl.Path
	} else {
		var err error
		if imgReader, err = os.Open(filePathOrUrl); err != nil {
			return Result{Err: err}
		}
		fileName = filePathOrUrl
	}

	return decode(fileName, imgReader)
}

func parseUrl(rawurl string) (*url.URL, bool) {
	parsedUrl, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return parsedUrl, false
	}
	return parsedUrl, true
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

func getByUrl(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	return resp, nil
}
