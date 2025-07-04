package model

type User struct {
	ID       *uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     *string `json:"name" gorm:"not null;unique"`
	Position *string `json:"position" gorm:"not null"`
}
