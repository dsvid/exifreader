package main

import (
	"fmt"
	"os"

	"github.com/dsvid/exifreader/internal/exifreader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: exifreader <file>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	model, err := exifreader.ExtractCameraModel(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading EXIF data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Camera Model: %s\n", model)
}
