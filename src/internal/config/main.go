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
		Dsn      string
		Database string
	}
)

func NewConfig() *Config {
	return &Config{
		Database: Database{
			Mongo: Mongo{
				Dsn:      os.Getenv("DB_MONGO_DSN"),
				Database: os.Getenv("DB_MONGO_DATABASE"),
			}}}
}
