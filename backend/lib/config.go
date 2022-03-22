package lib

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Api struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"api"`

	Sql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"sql"`
}

func NewConfig(path string) *Config {
	config := &Config{}
	config.LoadConfig(path)
	return config
}

func (c *Config) LoadConfig(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) GetDBConnection() string {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", c.Sql.Host, c.Sql.User, c.Sql.Password, c.Sql.Database, c.Sql.Port)
	return conn
}
