package main

import (
	"context"
	"fmt"
	"github.com/AskAlexSharov/imgdiff/avg"
	"github.com/AskAlexSharov/imgdiff/loader"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 args required: imgdiff ./file1.jpg ./file2.jpeg")
		return
	}

	dist, err := readAndGetDistance(context.Background(), os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Diffeernce: %d%%\n", dist)
}

func readAndGetDistance(ctx context.Context, fileName1, fileName2 string) (int, error) {
	ch1 := loader.ImageAsync(ctx, fileName1)
	ch2 := loader.ImageAsync(ctx, fileName2)
	r1 := <-ch1
	r2 := <-ch2
	if r1.Err != nil {
		return 0, r1.Err
	}
	if r2.Err != nil {
		return 0, r2.Err
	}

	return avg.Difference(r1.Img, r2.Img), nil
}
