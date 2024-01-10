package wordz
//wordz make sense
import ("crypto/rand"
		"math/big"
		"fmt"
)

var Hello = "This is package wordz" //Hello initialize a sting
var Prefix = "random word: " //Prefix makes it beatiful 
var Words = []string{"One", "Two", "Three", "Four", "Five"} //Words says

func Random() string {  
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64()) 
	} //Random aas

	func get(index int64) string {
		return Prefix + Words[index]
		}
	func init() {
		fmt.Println("Function init in package wordz")
	}
	