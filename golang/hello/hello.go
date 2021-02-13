package main

import(
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Jane")
	fmt.Println(message)
}
