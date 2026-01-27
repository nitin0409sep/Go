package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}


//! In Go, struct tags must not contain spaces around ':' ->  Because of the space, cleanenv cannot read those fields, so HttpServer stays empty ({}).
//! Addr string `yaml: "address" env-required:"true"`  -> Wrong
//! Addr string `yaml:"address" env-required:"true"` -> Correct

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "Path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exists: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Can not read config file %s", err.Error())
	}

	// fmt.Printf("%+v\n", cfg)

	return &cfg
}
