package entities

type Profil struct {
	ID     int
	Avatar string
	UserID int
	User   User `gorm:"foreignKey:UserID" json:"user"`
}
