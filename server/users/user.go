package users

type User struct {
	Name     string
	password string
	balance  float64
	Info     string
}

type UserRequest struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type UsersStorage struct {
	storage map[string]*User
}
