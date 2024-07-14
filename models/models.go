package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Summary  string `json:"summary"`
}

type Response struct {
	Message string `json:"message"`
}
