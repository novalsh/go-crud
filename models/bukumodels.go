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

func (p *BukuModel) FindAll() ([]entities.Buku, error) {
	rows, err := p.conn.Query("SELECT * FROM buku")
	if err != nil {
		return []entities.Buku{}, err
	}

	defer rows.Close()

	var bukus []entities.Buku

	for rows.Next() {
		var buku entities.Buku
		rows.Scan(&buku.Id, &buku.Judul, &buku.Jenis, &buku.Pengarang, &buku.Tahun, &buku.Harga)
		bukus = append(bukus, buku)
	}

	return bukus, nil
}

func (p *BukuModel) Create(buku entities.Buku) bool {

	result, err := p.conn.Exec("INSERT INTO buku (judul, jenis, pengarang, tahun, harga) VALUES (?, ?, ?, ?, ?)", buku.Judul, buku.Jenis, buku.Pengarang, buku.Tahun, buku.Harga)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
