package session

import (
    "www-auth-example/db/user"
)

type Session struct {
    Id string
    User *user.User
}

type Sessions map[string]Session

func (u *Sessions) Add(id string, user *user.User) {
    (*u)[id] = Session {
        Id: id,
        User: user,
    }
}

func Init() Sessions {
    s := Sessions{}
    return s
}
