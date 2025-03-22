package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Role     string
}

type Database struct {
	Connection *gorm.DB
}
