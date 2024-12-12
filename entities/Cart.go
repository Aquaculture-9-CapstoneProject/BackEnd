package entities

type Cart struct {
	ID        int
	UserID    int
	ProductID int
	Kuantitas int
	Subtotal  float64
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}

