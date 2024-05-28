package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/darleet/blog-graphql/internal/middleware/auth"
	"github.com/darleet/blog-graphql/internal/middleware/loader"
	"github.com/darleet/blog-graphql/internal/middleware/logging"
	"github.com/darleet/blog-graphql/internal/middleware/recoverer"
	"github.com/darleet/blog-graphql/internal/ports/gql/resolver"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
	"github.com/darleet/blog-graphql/internal/repository/pg"
	"github.com/darleet/blog-graphql/internal/usecase/article"
	"github.com/darleet/blog-graphql/internal/usecase/comment"
	"github.com/darleet/blog-graphql/internal/usecase/user"
	"github.com/darleet/blog-graphql/internal/usecase/vote"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	log := zap.Must(zap.NewDevelopment()).Sugar()
	defer func(log *zap.SugaredLogger) {
		err := log.Sync()
		if err != nil {
			log.Error(err)
		}
	}(log)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	repo := pg.NewRepository(pool)

	articles := article.NewUsecase(repo)
	comments := comment.NewUsecase(repo)
	votes := vote.NewUsecase(repo)
	users := user.NewUsecase()

	res := resolver.NewRootResolvers(log, articles, comments, users, votes)
	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(res))

	srv.AddTransport(&transport.Websocket{})

	srv.Use(extension.FixedComplexityLimit(300))

	recMW := recoverer.Middleware()
	authMW := auth.Middleware()
	logMW := logging.Middleware(log)

	http.Handle("/query", loader.Middleware(repo, authMW(logMW(recMW(srv)))))

	log.Info("Server started on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
