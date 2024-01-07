package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {

	// var myVar string = "Variable string" // (1)
	// var anotherVar string 		   		 // (2)

	// var noType = 100

	// var a,b,c int 				  		 // (3)
	// var d,e,f = "hello", 42, true 		 // (4)
	// var (						  		 // (5)
	// 	price 		int
	// 	qty		    int
	// 	isDeletable bool
	// )

	// pathToFile := getPath()				 // (6)

	// str, number, isExist := "new string", 42, false

	// str, number := "verride string", 10  // no new variables on left side of :=

	// var path string = "/path/to/file"
	// f, err := os.Open(path)

	// str, _ := getParams()

	// //constants

	// const statusCode int = 200
	// const (
	// 	orderStatusNew string = "new"
	// 	baseDiscount = 3.5
	// )

	// const (
	// 	untypedNum = 15
	// 	typedNum int = 10
	// )
	// var c int32 = 30

	// fmt,Println (c + untypedNum) 		//45
	// fmt.Println(c + typedNum) 			// invalid operation: c + typedNum

	// const (
	// 	num1 = 21
	// 	num2
	// 	num3 = 10
	// 	num4
	// )
	// fmt.Println(num1, num2, num3, num4)  // 21 1 10 10

	// const (
	// 	_ = 10 * iota
	// 	speed10 = 10 * iota
	// 	speed20 = 10 * iota
	// 	speed30 = 10 * iota
	// 	speed40 = 10 * iota
	// )
	// fmt.Println(speed10, speed20, speed30, speed40)

	// //----------

	// var num int = 100

	// var num8 int8 = 1 << 7 - 1
	// var unsignedNNum uint16 = 1 << 16 - 1

	// //float

	// fl := 4.54
	// fmt.Printf("%T\n", fl)  // float64

	// //bool

	// var b bool
	// var tr = true
	// isExist := true
	// fmt.Println(b,tr,isExist) // false true true

	//strings
	// var str string
	// var newStr string = "New string"

	// var strChar string = "!"
	// rCode, _ := utf8.DecodeRuneInString(strChar)
	// fmt.Printf("%U\n", rCode)

	// var runeChar rune = '\x21'
	// fmt.Println(string(runeChar))

	// var hello = "Hello"
	// fmt.Println(hello + " from NewYear!")

	// var greeting = "Привет!"
	// fmt.Println(len(greeting))
	// fmt.Println(utf8.RuneCountInString(greeting))

	// fmt.Println(greeting[3])
	// fmt.Println(greeting[4:6])

	// fmt.Println(greeting[:6])
	// fmt.Println(greeting[:])

	// // greeting[0] = "Л"
	// var convGreeting = []rune(greeting)
	// convGreeting[4] = 'E'
	// fmt.Println(string(convGreeting))
	// fmt.Println(len(convGreeting))

	// var cmplx complex128 = 1.1 + 2.1i
	// cmplx3 := complex(2.1, 2)
	// fmt.Println(cmplx, cmplx3)

	var defArr [3]string = [3]string{"one", "two", "three"}
	fmt.Println(defArr)

	var array [3]int
	fmt.Println(array)

	dynamic := [...]bool{4:true}
	fmt.Println(dynamic)
	fmt.Println(dynamic[2])

	arr := dynamic 
	arr[2] = true
	fmt.Println(dynamic) 

	arr2 := &dynamic
	arr2[2] = true
	fmt.Println(dynamic)

	s1 := dynamic[:2]
	s1[1] = true
	fmt.Println(dynamic)
}
