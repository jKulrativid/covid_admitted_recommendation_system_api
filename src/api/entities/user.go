package entities

type User struct {
	UserId   uint64 `json:"userid" binding:"omitempty" gorm:"column:userid"` // omitempty for userID=0
	UserName string `json:"username" binding:"required" gorm:"column:username"`
	Password string `json:"password" binding:"required" gorm:"column:password"`
}
