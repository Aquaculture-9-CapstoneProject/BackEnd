package entities

type Order struct {
	ID               int
	Total            float64
	UserID           int
	BiayaLayanan     float64
	BiayaOngkir      float64
	MetodePembayaran string
	CreatedAt        string
	Details          []OrderDetail `gorm:"foreignKey:" json:"details"`
}
