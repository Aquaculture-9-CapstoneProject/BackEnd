package entities

type Chat struct {
	ID        uint
	UserID    int
	UserInput string
	AiRespon  string
	User      User `gorm:"foreignKey:UserID"`
}
