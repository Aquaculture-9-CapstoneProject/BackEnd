package entities

type Review struct {
	ID        int
	ProductID int
	UserID    int
	Rating    float64
	Ulasan    string
	User      User `gorm:"foreignKey:UserID" json:"user"`
}
