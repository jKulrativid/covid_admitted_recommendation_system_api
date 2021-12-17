package entities

type User struct {
	Uuid           string `gorm:"primary_key"`              // user uuid
	UserName       string `gorm:"unique;type:varchar(32);"` // just a username
	Email          string `gorm:"unique;type:varchar(320)"` // user email (required)
	HashedPassword string `gorm:"type:varchar(72)"`         // user password already hashed with salt
	Salt           string `gorm:"type:varchar(32)"`         // unique salt for every user
}

// for sign-in sign-out
type UserSignIn struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}
