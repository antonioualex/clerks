package domain

type User struct {
	Name             Name    `json:"name"`
	Email            string  `json:"email"`
	PhoneNumber      string  `json:"phone"`
	Picture          Picture `json:"picture"`
	RegistrationDate string  `json:"registered_date"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Picture struct {
	Thumbnail string `json:"thumbnail"`
	Medium    string `json:"medium"`
	Large     string `json:"large"`
}

type UserService interface {
	Populate() error
	GetUsers(email string, limit, offset int) (users []User, startingAfter, endingBefore int, e error)
}

type UserRepository interface {
	AddUsers([]User) error
	GetUsers(email string, limit, offset int) (users []User, startingAfter, endingBefore int, e error)
}

type RandomUserRepository interface {
	FetchUsers() ([]User, error)
}
