package main

import (
	"github.com/zuekveer/abcdWay/02_task/wordz"
	newcolor "github.com/zuekveer/abcdWay/02_task/color"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
		}

}