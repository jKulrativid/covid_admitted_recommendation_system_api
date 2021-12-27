package entities

type User struct {
	Uid            string `gorm:"primary_key;type:uid"`     // user uuid
	UserName       string `gorm:"unique;type:varchar(32);"` // just a username
	Email          string `gorm:"unique;type:varchar(320)"` // user email (required)
	HashedPassword string `gorm:"type:varchar(150)"`        // user password already hashed with bcrypt
}

// for sign-in sign-out
type UserSignIn struct {
	UserName string `json:"username" validate:"required,min=5,max=40"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type UserRegister struct {
	UserName string `json:"username" validate:"required,min=5,max=40"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}
