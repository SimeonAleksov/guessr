package models

import (
	"errors"
	"html"
	"strings"

	"guessr.net/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)



type User struct {
    gorm.Model
    Username string `gorm:"size:255;not null;unique" json:"username"`
    Password string `gorm:"size:255;not null" json:"password"`
}


func (u *User) SaveUser() (*User, error) {
  var err error
  DB := database.GetDB()

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
  if err != nil {
      return nil, err
  }
  u.Username = html.EscapeString(strings.TrimSpace(u.Username))
  u.Password = string(hashedPassword)

  err = DB.Create(&u).Error
  if err != nil {
    return &User{}, err
  }

  return u, nil
}


func GetUserByUsername (username string) (uint, error) {
  DB := database.GetDB()
  var user User

  if err := DB.First(&user, "username = ?", username).Error; err != nil {
    return 0, errors.New("Resource not found")
  }

  return user.ID, nil
}


func GetUserByID (id uint) (User, error) {
  DB := database.GetDB()
  var user User

  if err := DB.First(&user, id).Error; err != nil {
    return user, errors.New("Resource not found")
  }

  return user, nil
}
