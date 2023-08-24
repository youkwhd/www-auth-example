package main

import (
	"os"

	"www-auth-example/db"
	"www-auth-example/db/user"

	"github.com/BurntSushi/toml"
)

type Config struct {
    Cookie Cookie
    Database Database
    Frontend Frontend
}

type Cookie struct {
    ExpiredAfter int `toml:"expired_after"`
}

type Database struct {
    Users []user.User
}

type Frontend struct {
    Urls []string
}

func InitConfig() Config {
    var config Config

    toml_content, err := os.ReadFile("config.toml")
    if err != nil {
        panic("Config file was not found")
    }

    // go is funny, makes me tickle
    meta, err := toml.Decode(string(toml_content), &config)
    if err != nil {
        // IGNORE
        meta.Undecoded()
        os.Exit(1)
    }

    return config;
}

func (c Config) AddUsers() {
    for _, val := range c.Database.Users {
        db.Data.Users.Add(val.Username, val.Password)
    }
}

func (c Config) GenerateAllowedOrigins() string {
    allowedOrigins := ""

    for idx, val := range c.Frontend.Urls {
        allowedOrigins += val

        if idx != len(c.Frontend.Urls) - 1 {
            allowedOrigins += ", "
        }
    }

    return allowedOrigins
}
