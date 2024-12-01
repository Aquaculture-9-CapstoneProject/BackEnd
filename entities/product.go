package entities

type Product struct {
	ID        int
	Nama      string
	Deskripsi string
	Variasi   string
	Kategori  string
	Harga     int
	Stok      int
	Gambar    string
	Ratings   []Rating
}

type Rating struct {
	ID        int
	ProductID int
	NamaUser  string
	Rating    int
	Ulasan    string
}
