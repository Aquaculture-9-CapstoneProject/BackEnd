package entities

type OrderDetail struct {
	ID        int
	OrderID   int
	ProductID int
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	UserID    int
	User      User `gorm:"foreignKey:UserID" json:"user"`
	Kuantitas int
	Subtotal  float64
}
