package main

import (
	"bufio"
	"fmt"
	"regexp"

	"log"
	"math/rand"
	"os"
	"time"
)

func main() {


	r1 := `(p[0-9]+:)+`

	var valid = regexp.MustCompile(r1)

	file, err := os.Open("pytania")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var p [100]string
	i := 0

	for scanner.Scan() {
		if valid.MatchString(scanner.Text()) {
			p[i] = valid.FindString(scanner.Text())
			i++
		}
	}

	fmt.Println("Pytań jest:",i)

	r := rand.NewSource(time.Now().UnixNano())
	s := rand.New(r)
	z := 0

	file2, err2 := os.Create("zadanie")

	if err2 != nil {
		log.Fatal(err2)
	}

	defer file2.Close()

	for z < 5 {
		zm := s.Intn(i)

			fmt.Println(p[zm])
			file2.WriteString(p[zm] + "\n")

			z++


	}

}
