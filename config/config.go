package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

func init() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	MediaURL  string `envconfig:"MEDIA_URL" default:"https://housebuildingstorage.blob.core.windows.net/media/"`
	SecretKey string `envconfig:"SECRET_KEY" default:"@ytbjco&cb^7hh!z3$dxih9+r$byc)s(!s-$_@f(cdt7))p4)x"`
}

var Cfg Config
