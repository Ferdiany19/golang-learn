package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// disiapkan utk mem-parsing penulisan tipe map tsb
type M map[string]interface{}

func main() {
	// memparsing semua file yang match dengan pattern yang ditentukan, dan fungsi ini mengembalikan 2 objek: *template.Template & error
	// Pattern path pada fungsi template.ParseGlob() nantinya akan diproses oleh filepath.Glob()
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		// param ke 1 obj, ke 2 nama template, ke3 data
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		// param ke 1 obj, ke 2 nama template, ke3 data
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
