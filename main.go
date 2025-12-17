package main

import (
	"fmt"

	"github.com/khubaib-binehsan/hellogo/bar/v1"
	v2 "github.com/khubaib-binehsan/hellogo/bar/v2"
	"github.com/khubaib-binehsan/hellogo/internal/foo"
)

func main() {
	fmt.Println("Hello, World from main.go")
	bar.Greet()
	v2.Greet()
	foo.Greet()
}
