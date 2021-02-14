package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // flag to disable printing
	message, err := greetings.Hello("Jane")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
}
