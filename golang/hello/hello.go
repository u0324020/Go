package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // flag to disable printing
	
	names := []string{"Jane", "Andy", "Jady"}
	
	messages, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(messages)
}
