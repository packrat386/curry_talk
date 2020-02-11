package main

import (
	"log"
	"net/http"
	"os"
)

// START OMIT
func main() {
	db, err := NewStore()
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "recipe-api: ", log.LstdFlags)

	http.Handle("/get_recipe", getRecipeHandler(db, logger))
	http.ListenAndServe(os.Getenv("ADDR"), nil)
}

// END OMIT
