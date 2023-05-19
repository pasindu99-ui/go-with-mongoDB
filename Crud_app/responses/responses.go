package responses

type Response struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	username string
	email    string
	password string
}
