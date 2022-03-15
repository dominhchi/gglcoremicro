package config

import "fmt"

type Config struct {
	Host     string `envconfig:"DB_HOST_SERVICE" default:"localhost"`
	Port     int    `envconfig:"DB_PORT_SERVICE" default:"5432"`
	User     string `envconfig:"DB_USER_SERVICE" default:"postgres"`
	DBName   string `envconfig:"DB_NAME_SERVICE" default:"service"`
	Password string `envconfig:"DB_PASSWORD_SERVICE" default:"postgres"`

	MediaURL  string `envconfig:"MEDIA_URL" default:"https://housebuildingstorage.blob.core.windows.net/media/"`
	SecretKey string `envconfig:"SECRET_KEY" default:"@ytbjco&cb^7hh!z3$dxih9+r$byc)s(!s-$_@f(cdt7))p4)x"`
}

var Cfg Config

func (cfg *Config) DbDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Saigon",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port,
	)
}
