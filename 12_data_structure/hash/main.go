package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	for sc.Scan() {

		l := sc.Text()
		n := HashBucket(l, 12)

		buckets[n] = append(buckets[n], l)

	}

	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stderr, "bład", err)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	word := "zordon"
	location := HashBucket(word, 12)

	for _, v := range buckets[location] {
		if v == word {
			fmt.Println("Bucket:",location,"have word:",v)
		}

	}

}

func HashBucket(w string, buckets int) int {

	var sum int

	for _, v := range w {
		sum += int(v)
	}

	return sum % buckets
}
