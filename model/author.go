package model

type Author struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:AuthID" json:"books"`
}
