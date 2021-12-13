package entities

type User struct {
	UserId   uint64 `json:"id" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
