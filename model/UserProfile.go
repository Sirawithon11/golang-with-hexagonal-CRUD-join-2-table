package model

type UserProfile struct {
	ID              *uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          *uint   `json:"user_id" gorm:"not null;index"`
	SkilledLanguage *string `json:"skilled_language" gorm:"not null"`
	Project1        *string `json:"project1"`
	Project2        *string `json:"project2"`
	Project3        *string `json:"project3"`
}
