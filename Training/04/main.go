package main

import ("fmt"
	"bufio"
	"os"
	"log"
	"regexp"
)

func main() {


	/*
	slownik.csv

	angielski,polski,overcome,przezwyciężać
	angielski,polski,whether,czy
	angielski,polski,iterative,wielokrotny
	angielski,polski,gleaned,zbierana
	angielski,polski,enhancements,udoskonalenia
	angielski,polski,near-term,w najbliższej przyszłości
	*/




	r:=`[a-zA-ZąćęłńóśźżĄĘŁŃÓŚŹŻ]+`
	var valid = regexp.MustCompile(r)

	file, err :=os.Open("slownik.csv")

	if err != nil {
		log.Fatal(err)
	}


	defer file.Close()


	sc :=bufio.NewScanner(file)

	var t = make([][]string,0)

	for sc.Scan(){
		t = append(t,valid.FindAllString(sc.Text(),4))

	}


	for i:=0;i<len(t);i++{
		for j:=0;j<len(t[0])/2;j++ {
			fmt.Print(t[i][j]+":"+t[i][j+2]+" ")
		}
		fmt.Println()
	}





}
