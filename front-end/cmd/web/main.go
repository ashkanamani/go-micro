package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)




func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		render(w, "test.page.gohtml")
	})

	fmt.Println("starting frontend service on port 80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, page string) {
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var tempateSlice []string 

	tempateSlice = append(tempateSlice, fmt.Sprintf("./cmd/web/templates/%s", page))
	
	tempateSlice = append(tempateSlice, partials...)

	tmpl, err := template.ParseFiles(tempateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
