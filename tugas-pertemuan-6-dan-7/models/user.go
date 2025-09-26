package models

import "fmt"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	{ID: 1, Username: "admin", Password: "admin123", Role: "admin"},
	{ID: 2, Username: "student1", Password: "student123", Role: "student"},
}

func (s *User) Validate() error {
	if s.ID == 0 || s.Username == "" || s.Password == "" {
		return fmt.Errorf("ID, Username, and Password cannot be empty")
	}
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
