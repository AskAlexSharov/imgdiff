package avg

import (
	"context"
	"math"

	"github.com/AskAlexSharov/imgdiff/loader"
	"github.com/Nr90/imgsim"
)

// Distance calculate average difference between 2 images
func Distance(ctx context.Context, fileName1, fileName2 string) (int, error) {
	ch := loader.ImagesAsync(ctx, fileName1, fileName2)

	// Hashing
	r1 := <-ch
	if r1.Err != nil {
		return 0, r1.Err
	}

	ahash1 := imgsim.AverageHash(r1.Img)

	r2 := <-ch
	if r2.Err != nil {
		return 0, r2.Err
	}
	ahash2 := imgsim.AverageHash(r2.Img)
	percents := float64(imgsim.Distance(ahash1, ahash2)) / 64 * 100 // because 64 bit hash

	return int(math.Round(percents)), nil
}
