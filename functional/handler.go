package main

import (
	"log"
	"net/http"
)

// START OMIT
type RecipeReader interface {
	ReadRecipe(name string) ([]byte, error)
}

func getRecipeHandler(db RecipeReader, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		getRecipe(w, r, db, logger)
	})
}

func getRecipe(w http.ResponseWriter, r *http.Request, db RecipeReader, logger *log.Logger) {
	logger.Println("begin get recipe")

	name := r.URL.Query().Get("name")

	data, err := db.ReadRecipe(name)
	if err != nil {
		logger.Println("get errored: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

// END OMIT
