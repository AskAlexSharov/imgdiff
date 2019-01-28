package avg

import (
	"context"

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

	return imgsim.Distance(ahash1, ahash2) % 64, nil // because 64 bit hash
}

//func Difference(img1, img2 image.Image) int {
//	// Hashing
//	ahash1 := imgsim.AverageHash(img1)
//	ahash2 := imgsim.AverageHash(img2)
//
//	dhash1 := imgsim.DifferenceHash(img1)
//	dhash2 := imgsim.DifferenceHash(img2)
//
//	// distance
//	avgDistance := imgsim.Distance(ahash1, ahash2)
//	diffDistance := imgsim.Distance(dhash1, dhash2)
//
//	return (avgDistance + diffDistance) % 128
//}
