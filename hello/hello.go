package main

import (
	"fmt"
	"go-art/buildstring"
)

func main() {
	fmt.Printf("Hello, %q!\n", buildstring.Build("Whole", "New", "World"))
}
