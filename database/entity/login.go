package entity

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Login) TableName() string {
	return "Login"
}
