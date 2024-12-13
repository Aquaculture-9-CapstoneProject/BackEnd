package entities

type Payment struct {
	ID           int    `gorm:"primaryKey"`
	InvoiceID    string `gorm:"type:varchar(255);uniqueIndex"`
	Status       string
	UserID       int
	StatusBarang string
	Jumlah       float64
	OrderID      int   `gorm:"not null"`
	Order        Order `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	PaymentUrl   string
}
