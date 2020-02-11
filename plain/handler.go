package main

import (
	"net/http"
)

// START OMIT
func getRecipe(w http.ResponseWriter, r *http.Request) {
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
