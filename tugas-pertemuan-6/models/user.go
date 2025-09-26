package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var Users = []User{
	{ID: 1, Username: "admin", Password: "admin123", Role: "admin"},
	{ID: 2, Username: "student1", Password: "student123", Role: "student"},
}
