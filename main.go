package main

import (
	"os"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.devdungeon.com")
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
	
}