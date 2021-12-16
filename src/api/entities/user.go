package entities

type User struct {
	Uuid           string `gorm:"column:uuid"` // omitempty for userID=0
	UserName       string `gorm:"column:username"`
	Email          string `gorm:"column:email" valid:"email"`
	HashedPassword string `gorm:"column:password"`
	Salt           string `gorm:"column:salt"`
}

type UserLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
