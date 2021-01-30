package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

const (
	url string = "mongodb://localhost:27017"
)

var (
	_db *mgo.Session //store the db conneciton pool
)

func main() {
	printDetails()
	//start DB client, and storing connection client in context
	if err := startDBConnection(); err != nil {
		log.Fatalln(err)
	}
	//start API service
	if err := startWebService(); err != nil {
		log.Fatalln(err)
	}
}
