package entities

type Buku struct {
	Id        int64  `json:"id"`
	Judul     string `json:"judul"`
	Jenis     string `json:"jenis"`
	Pengarang string `json:"Pengarang"`
	Tahun     string `json:"tahun"`
	Harga     string `json:"harga"`
}
