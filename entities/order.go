package entities

type Order struct {
	ID               int
	Total            float64
	UserID           int
	User             User `gorm:"foreignKey:UserID" json:"user"`
	BiayaLayanan     float64
	BiayaOngkir      float64
	MetodePembayaran string
	Details          []OrderDetail `gorm:"foreignKey:OrderID" json:"details"`
}
