package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Make request
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)

}
