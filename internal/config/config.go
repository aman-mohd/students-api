package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

//@TO_REMEMBER these `` backticks are called struct tags

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	//@TO_SEE: embedding of struct
	HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	// now we check if it's available from the terminal flags
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		//@TO_SEE: what is D reference in go
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config path is empty")
			// panic("CONFIG_PATH or -config flag is required")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file not found at %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err.Error())
	}

	return &cfg
}
