package main

import (
	"log"
	"os"
	user2 "os/user"
)

type Credentials struct {
	Username string
	Password string
}

func GetCredentialsFromEnvironmentVariables() *Credentials {
	username := getCurrentUsername()
	if len(username) < 1 {
		panic("no username set, set username via [username] environment variable")
	}

	password := getFromEnv("password")
	if len(password) < 1 {
		panic("no password set, set password via [password] environment variable")
	}

	return &Credentials{
		Username: username,
		Password: password,
	}
}

func getCurrentUsername() string {
	user, err := user2.Current()
	if err != nil {
		log.Print(err)
		return getFromEnv("username")
	}
	return useDefaultIfEmpty(user.Username, getFromEnv("username"))
}

func getFromEnv(name string) string {
	username, _ := os.LookupEnv(name)
	return username
}

func useDefaultIfEmpty(original, def string) string {
	if len(original) < 1 {
		return def
	}
	return original
}
