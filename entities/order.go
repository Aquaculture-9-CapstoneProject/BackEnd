package entities

type Order struct {
	ID               int
	Total            float64
	UserID           int
	User             User `gorm:"constraint:OnDelete:SET NULL;" json:"user"`
	BiayaLayanan     float64
	BiayaOngkir      float64
	MetodePembayaran string
	CreatedAt        string
	Details          []OrderDetail `gorm:"foreignKey:OrderID" json:"details"`
}
