package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Join folder dan file menjadi path
		var filepath = path.Join("views", "index.html")
		// mem-parsing fike template dan mengembalikan 2 data
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			// cek error 500 internel server
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// untuk menyisipkan ke template yang sudah di parsing
		var data = map[string]interface{}{
			"title": "Learning Golang",
			"name":  "Batman",
		}

		// Execute utk menyisipkan data pada template, utk ditampilkan ke browser
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("Server Started localhost:9000")
	http.ListenAndServe(":9000", nil)
}
