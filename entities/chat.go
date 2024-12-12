package entities

type Chat struct {
	UserID    int
	UserInput string
	AiRespon  string
	User      User `gorm:"foreignKey:UserID"`
}
