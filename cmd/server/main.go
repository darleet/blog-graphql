package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/darleet/blog-graphql/internal/middleware/auth"
	"github.com/darleet/blog-graphql/internal/ports/gql/resolver"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
	"github.com/darleet/blog-graphql/internal/repository/pg"
	"github.com/darleet/blog-graphql/internal/usecase/article"
	"github.com/darleet/blog-graphql/internal/usecase/comment"
	"github.com/darleet/blog-graphql/internal/usecase/vote"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const defaultPort = "8888"

func main() {
	log := zap.Must(zap.NewDevelopment()).Sugar()

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
	votes := vote.NewUsecase(nil)

	res := resolver.NewRootResolvers(log, articles, comments, nil, votes)
	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(res))

	srv.AddTransport(&transport.Websocket{})

	authMW := auth.NewMiddleware()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authMW(srv))

	log.Info("Server started on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
