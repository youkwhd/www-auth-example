package config

import (
	"os"

	"www-auth-example/db/user"

	"github.com/BurntSushi/toml"
)

type Config struct {
    Cookie cookie
    Database database
    Frontend frontend
}

type cookie struct {
    Expired_after int
}

type database struct {
    Users []user.User
}

type frontend struct {
    Urls []string
}

var Conf Config

func Init() {
    toml_content, err := os.ReadFile("config.toml")
    if err != nil {
        panic("Config file was not found")
    }

    // go is funny, makes me tickle
    meta, err := toml.Decode(string(toml_content), &Conf)
    if err != nil {
        // IGNORE
        meta.Undecoded()
        os.Exit(1)
    }
}
