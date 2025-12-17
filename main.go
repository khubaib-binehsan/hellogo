package main

import (
	"fmt"

	"github.com/khubaib-binehsan/hellogo/internal/foo"
	"github.com/khubaib-binehsan/hellogo/pkg/bar"
)

func main() {
	fmt.Println("Hello, World from main.go")
	bar.Greet()
	foo.Greet()
}
