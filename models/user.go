package models

//User models for user
type User struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-" gorm:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUser() *User {
	return &User{}
}
