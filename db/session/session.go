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

    // Checking and saving the state of a session would expires in is 
    // kind of unnecessary, since the browser will clear off the cookie
    // automatically anyway.
    //
    // Nevertheless, maybe something didn't go as expected in the browser,
    // so, double checking considered ok.
    if session.ExpiresIn.Sub(time.Now()) <= 0 {
        delete(*s, id)
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
