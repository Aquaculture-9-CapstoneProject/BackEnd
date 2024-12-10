package entities

type Payment struct {
	ID           int
	InvoiceID    int `gorm:"uniqueIndex"`
	Status       string
	StatusBarang string
	Jumlah       float64
	OrderID      int
	Order        Order `gorm:"foreignKey:OrderID" json:"order"`
	PaymentUrl   string
}
