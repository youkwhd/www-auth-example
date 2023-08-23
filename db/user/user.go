package user

type User struct {
    Username string
    Password string
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

    u.Add("youkwhd", "youkwhd")
    u.Add("jake", "admin")
    u.Add("admin", "admin")

    return u
}
