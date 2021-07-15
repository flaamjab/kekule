package main

import (
	"fmt"
	"os"

	"github.com/flaamjab/kekule/internal/api"
)

func main() {
	if len(os.Args) == 2 {
		addr := os.Args[1]
		api.Run(addr)
	}
	if len(os.Args) == 1 {
		fmt.Println("Please provide the adress to listen at (e.g. localhost:8080)")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	}
}
