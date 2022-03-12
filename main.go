package main
import (
	"fmt"
)

const intro = "Hello, "
func main() {
	fmt.Println(greeting("Nahom! "))
     
	// to print imported file jsonFile.go which have the same package with main.go
	message :="We are 2F-Capital Intership Go developers!"
	result := jsonMarshaling(message)
	fmt.Println(result)

	// run "go run main.go jsonFile.go

}
 
func greeting(name string) string{
	message := " How are you doing?"
	return intro + name + message
}
