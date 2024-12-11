package entities

type Product struct {
	ID              int
	Gambar          string
	Nama            string
	Deskripsi       string
	Keunggulan      string
	Harga           int
	Variasi         string
	Kategori        string
	KotaAsal        string
	Rating          float64
	Stok            int
	TotalReview     int
	TipsPenyimpanan string
	Reviews         []Review      `gorm:"foreignKey:ProductID" json:"reviews"`
	OrderDetails    []OrderDetail `gorm:"foreignKey:ProductID" json:"order_details,omitempty"`
}
