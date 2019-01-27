package avg

import (
	"github.com/Nr90/imgsim"
	"image"
)

func Difference(img1, img2 image.Image) int {
	// Hashing
	ahash1 := imgsim.AverageHash(img1)
	ahash2 := imgsim.AverageHash(img2)

	// distance
	avgDistance := imgsim.Distance(ahash1, ahash2)

	return avgDistance % 64 // because 64 bit hash
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
