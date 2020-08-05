package avg

import (
	"context"
	"fmt"
	"image"
	"math"

	"github.com/AskAlexSharov/imgdiff/loader"
)

// Distance calculate average difference between 2 images
func Distance(ctx context.Context, fileName1, fileName2 string) (int, error) {
	ch := loader.ImagesAsync(ctx, fileName1, fileName2)

	// Hashing
	r1 := <-ch
	if r1.Err != nil {
		return 0, r1.Err
	}

	r2 := <-ch
	if r2.Err != nil {
		return 0, r2.Err
	}

	r, err := DistancePure(ctx, r1.Img, r2.Img)
	if err != nil {
		return 0, err
	}
	return r, nil
}

func DistancePure(ctx context.Context, img1, img2 image.Image) (int, error) {
	ahash1 := AverageHash(img1)
	ahash2 := AverageHash(img2)
	fmt.Printf("ahash1: %v\n",uint64(ahash1))
	fmt.Printf("ahash2: %v\n",uint64(ahash2))

	percents := float64(Distance2(ahash1, ahash2)) / 64 * 100 // because 64 bit hash

	return int(math.Round(percents)), nil
}
