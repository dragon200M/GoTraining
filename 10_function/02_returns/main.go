package main

import "fmt"

func main() {

	fmt.Println(greet("Maciej","Michalik"))

	fmt.Println(greetValue("Łucja", "Michalik"))

}

func greet(fname, lname string) string{
	return fmt.Sprint(fname,lname) //return string
}

func greetValue(fname, lname string) (s string){
	s = fmt.Sprint(fname,lname) //return s value
	return
}