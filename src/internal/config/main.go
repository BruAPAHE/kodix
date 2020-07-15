package config

import "os"

type (
	Config struct {
		Database Database
	}

	Database struct {
		Mongo Mongo
	}

	Mongo struct {
		Name     string
		Dsn      string
		Database string
	}
)

func NewConfig() *Config {
	return &Config{
		Database: Database{
			Mongo: Mongo{
				Name:     "AutoData",
				Dsn:      os.Getenv("DB_MONGO_DSN"),
				Database: os.Getenv("DB_MONGO_DATABASE"),
			}}}
}
