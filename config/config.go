package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const ModeMemory = "MEMORY"
const ModePostgres = "POSTGRES"

type Config struct {
	Host            string
	Port            string
	ComplexityLimit int
	ArticleLimit    int
	CommentLimit    int
	Mode            string
	PostgresURL     string
}

func NewConfig(paths ...string) (*Config, error) {
	err := godotenv.Load(paths...)
	if err != nil {
		return nil, err
	}

	var config Config
	var ok bool

	config.Host, ok = os.LookupEnv("SERVER_HOST")
	if !ok {
		config.Host = "127.0.0.1"
	}
	config.Port, ok = os.LookupEnv("SERVER_PORT")
	if !ok {
		config.Port = "8080"
	}

	complexLimit, ok := os.LookupEnv("COMPLEXITY_LIMIT")
	if !ok {
		config.ComplexityLimit = 300
	} else {
		config.ComplexityLimit, err = strconv.Atoi(complexLimit)
	}

	config.Mode, ok = os.LookupEnv("MODE")
	if !ok {
		config.Mode = ModeMemory
	}

	config.PostgresURL, ok = os.LookupEnv("PG_URL")
	if !ok && config.Mode == ModePostgres {
		return nil, fmt.Errorf("missing PG_URL")
	}

	arLimit, ok := os.LookupEnv("ARTICLE_LIMIT")
	if !ok {
		config.ArticleLimit = 10
	} else {
		config.ArticleLimit, err = strconv.Atoi(arLimit)
	}

	commLimit, ok := os.LookupEnv("COMMENT_LIMIT")
	if !ok {
		config.CommentLimit = 10
	} else {
		config.CommentLimit, err = strconv.Atoi(commLimit)
	}

	return &config, err
}
