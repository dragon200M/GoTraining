package vis

import "fmt"


var x int = 42

func PrintVa() {

	fmt.Println(x)
	foo()
	fmt.Printf("Zasięg z tego samego pakietu: \x1b[32;1m %v \x1b[0m \n",name)
	fmt.Println(LastName)

}

func foo() {
	fmt.Println(x)
}
