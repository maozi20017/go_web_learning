package main

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role" gorm:"foreignkey:Role_ID"`
	Role_ID  int    `json:"role_id"`
}

type Role struct {
	Role_id     int    `json:"role_id" gorm:"primary_key"`
	Description string `json:"description"`
}
