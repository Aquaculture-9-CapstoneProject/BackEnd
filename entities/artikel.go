package entities

import "time"

type Artikel struct {
	ID        int
	Judul     string
	Deskripsi string
	Kategori  string
	CreatedAt time.Time
	AdminID   int
	Admin     Admin
}
