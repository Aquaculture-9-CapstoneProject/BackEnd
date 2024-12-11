package entities

type Product struct {
	ID           int
	Gambar       string
	Nama         string
	Deskripsi    string
	Keunggulan   string
	Harga        int
	Variasi      string
	Kategori     string
	KotaAsal     string
	Rating       float64
	Stok         int
	TotalReview  int
	Reviews      []Review      `gorm:"foreignKey:ProductID"`
	OrderDetails []OrderDetail `gorm:"foreignKey:ProductID"`
}
