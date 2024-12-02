package entities

type Review struct {
	ID        int
	ProductID int
	UserID    uint
	Rating    float64
	Ulasan    string
	User      User `gorm:"foreignKey:UserID"`
}
