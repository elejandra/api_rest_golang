package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elejandra/api_rest_golang/pkg/storage/inmem"
	sample "github.com/elejandra/api_rest_golang/cmd/sample-data"
	gopher "github.com/elejandra/api_rest_golang/pkg"

	"github.com/elejandra/api_rest_golang/pkg/server"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some gophers")
	flag.Parse()

	var gophers map[string]*gopher.Gopher
	if *withData {
		gophers = sample.Gophers
	}

	repo := inmem.NewGopherRepository(gophers)
	s := server.New(repo)

	fmt.Println("The gopher server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}