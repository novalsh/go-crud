package bukucontroller

import (
	"html/template"
	"net/http"

	"github.com/novalsh/go-crud/entities"
	"github.com/novalsh/go-crud/models"
)

var bukuModel = models.NewBukuModel()

func Index(response http.ResponseWriter, request *http.Request) {

	buku, _ := bukuModel.FindAll()

	data := map[string]interface{}{
		"buku": buku,
	}

	temp, err := template.ParseFiles("views/buku/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/buku/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var buku entities.Buku
		buku.Judul = request.Form.Get("judul")
		buku.Jenis = request.Form.Get("jenis")
		buku.Pengarang = request.Form.Get("pengarang")
		buku.Tahun = request.Form.Get("tahun")
		buku.Harga = request.Form.Get("harga")

		bukuModel.Create(buku)
		data := map[string]interface{}{
			"message": "Data berhasil disimpan",
		}
		temp, _ := template.ParseFiles("views/buku/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
