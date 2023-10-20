package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	bs := make([]byte, 100000)
	res.Body.Read(bs)
	//fmt.Println(string(bs))
	f := os.Stdout
	f.Write(bs)
}
