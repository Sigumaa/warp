package main

import "log"

func main() {
	uri, err := loadURI()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(uri)
}
