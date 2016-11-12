package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	words := make(map[string]string)
	w := make([]string, 0)
	buckets := make([]int, 200)


	for sc.Scan() {
		words[sc.Text()] = ""
		w = append(w, sc.Text())
		n := HashBucket(sc.Text())

		buckets[n]++

	}

	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stderr, "bład", err)
	}

	for i := 0; i < len(w); i++ {
		//fmt.Println(w[i])
		if i == 200 {
			log.Println("Break")
			break

		}
	}


	fmt.Println(buckets[97:123])

}

func HashBucket(w string) int {
	return int(w[0])
}
