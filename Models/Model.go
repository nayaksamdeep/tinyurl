package Models

import (
   //"time"
)

type RedirectUrl struct {
	ID uint            `json:"id"`
	Url string       `form:"url" json:"url" binding:"required"`
	TinyUrl string `json:"tinyurl"`
	//CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	//ExpiryDate time.Time `gorm:"default.CURRENT_TIMESTAMP" json:"expiry_date"`
}

/*
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Contact) TableName() string {
	return "contact"
}
*/

