package session

import (
	"errors"
	"time"
	"www-auth-example/db/user"
)

type Session struct {
    Id string
    User *user.User
    ExpiresIn time.Time
}

type Sessions map[string]Session

func (s *Sessions) Get(id string) (*Session, error) {
    session, found := (*s)[id]

    if !found {
        return nil, errors.New("Session not found")
    }

    if time.Now().Sub(session.ExpiresIn) <= 0 {
        return nil, errors.New("Session is expired")
    }

    return &session, nil
}

func (s *Sessions) Add(id string, user *user.User, expiresIn time.Time) {
    (*s)[id] = Session {
        Id: id,
        User: user,
        ExpiresIn: expiresIn,
    }
}

func Init() Sessions {
    s := Sessions{}
    return s
}
