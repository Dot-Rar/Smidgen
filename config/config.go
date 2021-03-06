package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type(
	Config struct {
		Server Server
		Database Database
		Ratelimit Ratelimit
	}

	Server struct {
		Address string
	}

	Database struct {
		Host string
		Port int
		Username string
		Password string
		Database string
	}

	Ratelimit struct {
		PastesPerHour int `json:"pastes-per-hour"`
	}
)

var(
	Conf Config
)

func LoadConfig()  {
	log.Println("Loading config")

	bytes, err := ioutil.ReadFile("config.json"); if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &Conf); if err != nil {
		panic(err)
	}
}
