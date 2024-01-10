package wordz

import ("crypto/rand"
		"math/big"
)

var Hello = "This is package wordz" // says Hello
var Prefix = "random word: " //says PRefix
var Words = []string{"One", "Two", "Three", "Four", "Five"} //saysWords

func Random() string { //makes Func
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64()) 
	}

	func get(index int64) string {
		return Prefix + Words[index]
		}
	