package db

import (
    "www-auth-example/db/user"
    "www-auth-example/db/session"
)

var Data Database = Init()

type Database struct {
    Users user.Users
    Sessions session.Sessions
}

func Init() Database {
    d := Database{
        Users: user.Init(),
        Sessions: session.Init(),
    }

    return d
}
