package entities

type Product struct {
	ID         int
	Nama       string
	Deskripsi  string
	Variasi    string
	KategoriID int
	RatingID   int
	Harga      int
	Stok       int
	Gambar     string
	Kategori   Category
	Rating     Rating
}

type Category struct {
	ID        int
	Nama      string
	Deskripsi string
}

type Rating struct {
	ID       int
	NamaUser string
	Rating   int
	Ulasan   string
}
