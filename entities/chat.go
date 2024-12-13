package entities

type Chat struct {
	ID        int
	UserID    int
	UserInput string
	AiRespon  string
	User      User `gorm:"foreignKey:UserID"`
}
