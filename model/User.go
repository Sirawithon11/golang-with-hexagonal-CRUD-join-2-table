package model

type User struct {
	ID       *uint   `json:"id" gorm:"primaryKey key;autoIncrement"`
	Name     *string `json:"name" gorm:"not null"`
	Position *string `json:"position" gorm:"not null"`
}
