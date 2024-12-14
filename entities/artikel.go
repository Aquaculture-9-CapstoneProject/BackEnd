package entities

import "time"

type Artikel struct {
	ID        int
	Gambar    string
	Judul     string
	Deskripsi string
	Kategori  string
	CreatedAt time.Time
}
