package main

import (
	"log"
	"net/http"
	"os"
)

// START OMIT
var (
	db     Store
	logger *log.Logger
)

func main() {
	var err error
	db, err = NewStore()
	if err != nil {
		panic(err)
	}

	logger = log.New(os.Stdout, "recipe-api: ", log.LstdFlags)

	http.Handle("/get_recipe", http.HandlerFunc(getRecipe))
	http.ListenAndServe(os.Getenv("ADDR"), nil)
}

// END OMIT
