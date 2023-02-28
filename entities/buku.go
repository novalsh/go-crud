package entities

type Buku struct {
	Id        int64  `json:"id"`
	Judul     string `json:"judul"`
	Jenis     string `json:"pengarang"`
	Pengarang string `json:"penerbit"`
	Tahun     string `json:"tahun"`
	Harga     string `json:"harga"`
}
