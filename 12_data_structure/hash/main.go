package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
	"os"
)

func main() {
	res, err :=http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err  !=nil{
		log.Println(err)
	}

	defer res.Body.Close()

	sc :=bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	words :=make(map[string]string)
	w := make([]string,0)
	for sc.Scan(){
		words[sc.Text()] = ""
		w = append(w,sc.Text())
	}

	if err :=sc.Err();err !=nil{
		fmt.Fprint(os.Stderr,"bład",err)
	}

	for i:=0;i<len(w);i++ {
		fmt.Println(w[i])
	}
}
