package main

import (
	"net/http"

	"github.com/novalsh/go-crud/controller/bukucontroller"
)

func main() {
	http.HandleFunc("/", bukucontroller.Index)
	http.HandleFunc("/buku", bukucontroller.Index)
	http.HandleFunc("/buku/index", bukucontroller.Index)
	http.HandleFunc("/buku/add", bukucontroller.Add)
	http.HandleFunc("/buku/edit", bukucontroller.Edit)
	http.HandleFunc("/buku/delete", bukucontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
