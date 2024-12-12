package entities

type User struct {
	ID             int
	NamaLengkap    string
	Alamat         string
	NoTelpon       string
	Email          string
	Password       string
	KonfirmasiPass string
	Orders         []Order  `gorm:"foreignKey:UserID" json:"orders"`
	Reviews        []Review `gorm:"foreignKey:UserID" json:"reviews,omitempty"`
}
