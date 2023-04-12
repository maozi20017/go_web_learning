package main

type User struct {
	ID         int64    `json:"id" gorm:"primary_key;auto_increment"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	IdentityID int      `json:"identity_id"`
	Identity   Identity `json:"identity" gorm:"foreignKey:IdentityID"`
}

type Identity struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
	Users       []User `json:"users" gorm:"foreignKey:IdentityID"`
}
