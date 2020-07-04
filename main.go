package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello again"))
	})

	// Static Route file perlu didefiniskan
	// rute-nya dan handler-nya

	// Sedang untuk handler-nya bisa di-lihat,
	// ada pada parameter ke-2 yang isinya statement http.StripPrefix().
	// Sebenarnya actual handler nya berada pada http.FileServer(). Fungsi http.StripPrefix()
	// hanya digunakan untuk membungkus actual handler.

	// Fungsi http.StripPrefix() ini berguna untuk menghapus prefix dari endpoint yang di-request.

	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
