package entities

type Payment struct {
	ID           int
	InvoiceID    string `gorm:"uniqueIndex"`
	Status       string
	StatusBarang string
	Jumlah       float64
	OrderID      int
	Order        Order `json:"order" gorm:"foreignKey:OrderID"`
	PaymentUrl   string
}
