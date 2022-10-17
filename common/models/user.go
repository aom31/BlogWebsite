package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        uint   `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
	Password  []byte `json: "-"`
	Phone     string `json: "phone"`
}

//hash password
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

//check password is correct, it correspond new and exist in db
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
