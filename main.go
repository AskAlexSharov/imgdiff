package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AskAlexSharov/imgdiff/avg"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("2 args required: imgdiff ./file1.jpg ./file2.jpeg")
		return
	}

	dist, err := avg.Distance(context.Background(), os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	fmt.Printf("Diffeernce: %d%%\n", dist)
}
