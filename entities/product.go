package entities

type Product struct {
	ID              int
	Gambar          string
	Nama            string
	Deskripsi       string
	Keunggulan      string
	Harga           float64
	Variasi         string
	Kategori        string
	KotaAsal        string//input otomatis
	Rating          float64
	Stok            int
	TotalReview     int
	Status          string
	Terjual         int
	TipsPenyimpanan string
	Reviews         []Review      `gorm:"foreignKey:ProductID" json:"reviews"`
	OrderDetails    []OrderDetail `gorm:"foreignKey:ProductID" json:"order_details,omitempty"`
}
