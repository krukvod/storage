package main

import (
	"fmt"
	"github.com/krukvod/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it's work", st)
}
