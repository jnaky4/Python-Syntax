package conditionals

import (
	"fmt"
	"math/rand"
)

type HTTPRequest struct {
	Method string
}
type contact struct {
	greeting string
	name     string
}

func main() {
	u1 := 1234
	u2 := 3456

	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}

	if u1 == u2 {
		println("Same")
	}

	//rolls 0,1
	b := rand.Intn(2)

	//food variable on in block scope
	if food := "Chocolate"; b > 0 {
		fmt.Println(food)
		println(b)
	}
	//error: fmt.Println(food)

	//switch statements
	//implicit breaking instead of fallthrough
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Select Request")
		//how to fallthrough
		fallthrough
		//multiple evals
	case "POST", "DELETE":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("Unhandled Method")
	}

	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}
