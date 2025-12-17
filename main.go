package main

import (
	"fmt"

	"github.com/khubaib-binehsan/hellogo/bar"
	"github.com/khubaib-binehsan/hellogo/internal/foo"
)

func main() {
	fmt.Println("Hello, World from main.go")
	bar.Greet()
	foo.Greet()
}
