package models

import (
	"database/sql"
	"fmt"

	"github.com/novalsh/go-crud/config"
	"github.com/novalsh/go-crud/entities"
)

type BukuModel struct {
	conn *sql.DB
}

func NewBukuModel() *BukuModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &BukuModel{
		conn: conn,
	}
}

func (p *BukuModel) Create(buku entities.Buku) bool {

	result, err := p.conn.Exec("INSERT INTO buku (nama_buku, jenis_buku, pengarang, tahun_terbit, harga_buku) VALUES (?, ?, ?, ?, ?)", buku.Judul, buku.Jenis, buku.Pengarang, buku.Tahun, buku.Harga)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
