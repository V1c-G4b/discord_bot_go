package schemas

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Bots      []Bot     `json:"bots" gorm:"foreignKey:OwnerID"`
}

type Bot struct {
	gorm.Model
	ID      uint        `json:"id" gorm:"primaryKey"`
	Name    string      `json:"name"`
	APIKey  string      `json:"api_key"`
	OwnerID uint        `json:"owner_id"`
	Owner   User        `json:"owner" gorm:"foreignKey:OwnerID"`
	Members []BotMember `json:"members" gorm:"foreignKey:BotID"`
}

type BotMember struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	BotID  uint   `json:"bot_id"`
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}
