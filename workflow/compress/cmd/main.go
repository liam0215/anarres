package main

import (
	"fmt"

	"github.com/liam0215/anarres/workflow/compress"
)

func main() {
	input := []byte("lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	fmt.Printf("Input: %s\n", input)

	comp, err := compress.Compress(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Compressed from %d to %d bytes\n", len(input), len(comp))

	dec, err := compress.Decompress(comp, len(input))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decompressed: %s\n", dec)
}
