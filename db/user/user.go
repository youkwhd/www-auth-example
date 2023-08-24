package user

type User struct {
    Username string `json:"username" xml:"username" form:"username"`
    Password string `json:"password" xml:"password" form:"password"`
}

type Users map[string]User

func (u *Users) Add(username string, password string) {
    (*u)[username] = User {
        Username: username,
        Password: password,
    }
}

func Init() Users {
    u := Users{}
    return u
}
