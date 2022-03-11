package main
import (
	"fmt"
)

const intro = "Hello, "
func main(){
	fmt.Println(greeting("Nahom! "))
   
}
 
func greeting(name string) string{
	message := " How are you doing?"
	return intro + name + message
}
