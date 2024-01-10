package main

import (
	"github.com/zuekveer/abcdWay/02_task/wordz"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello world")
	color.Red("Hello world again")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
		}

}