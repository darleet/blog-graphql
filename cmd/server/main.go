package main

import (
	"github.com/darleet/blog-graphql/config"
	"github.com/darleet/blog-graphql/internal/server"
	"go.uber.org/zap"
)

func main() {
	log := zap.Must(zap.NewDevelopment()).Sugar()
	defer func(log *zap.SugaredLogger) {
		err := log.Sync()
		if err != nil {
			log.Error(err)
		}
	}(log)

	conf, err := config.NewConfig(".env")
	if err != nil {
		log.Fatal(err)
	}
	err = server.NewServer(log, conf).Start()
	if err != nil {
		log.Fatal(err)
	}
}
