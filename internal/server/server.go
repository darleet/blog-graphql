package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/darleet/blog-graphql/config"
	"github.com/darleet/blog-graphql/internal/middleware/auth"
	"github.com/darleet/blog-graphql/internal/middleware/loader"
	"github.com/darleet/blog-graphql/internal/middleware/logging"
	"github.com/darleet/blog-graphql/internal/middleware/recoverer"
	"github.com/darleet/blog-graphql/internal/ports/gql/resolver"
	"github.com/darleet/blog-graphql/internal/ports/gql/runtime"
	"github.com/darleet/blog-graphql/internal/repository/local"
	"github.com/darleet/blog-graphql/internal/repository/pg"
	"github.com/darleet/blog-graphql/internal/usecase/article"
	"github.com/darleet/blog-graphql/internal/usecase/comment"
	"github.com/darleet/blog-graphql/internal/usecase/user"
	"github.com/darleet/blog-graphql/internal/usecase/vote"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	log    *zap.SugaredLogger
	config *config.Config
}

func NewServer(log *zap.SugaredLogger, config *config.Config) *Server {
	return &Server{
		log:    log,
		config: config,
	}
}

func (s *Server) resolverLocal() http.Handler {
	repo := local.NewRepository()
	res := resolver.NewRootResolvers(
		s.log,
		article.NewUsecase(repo),
		comment.NewUsecase(nil),
		user.NewUsecase(),
		vote.NewUsecase(repo),
	)

	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(res))
	srv.AddTransport(&transport.Websocket{})
	srv.Use(extension.FixedComplexityLimit(s.config.ComplexityLimit))

	recMW := recoverer.Middleware(s.log)
	authMW := auth.Middleware()
	logMW := logging.Middleware(s.log)

	return loader.Middleware(repo, authMW(logMW(recMW(srv))))
}

func (s *Server) resolverPG() http.Handler {
	pool, err := pgxpool.New(context.Background(), s.config.PostgresURL)
	if err != nil {
		s.log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		s.log.Fatal(err)
	}
	repo := pg.NewRepository(nil)
	res := resolver.NewRootResolvers(
		s.log,
		article.NewUsecase(repo),
		comment.NewUsecase(repo),
		user.NewUsecase(),
		vote.NewUsecase(repo),
	)

	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(res))
	srv.AddTransport(&transport.Websocket{})
	srv.Use(extension.FixedComplexityLimit(s.config.ComplexityLimit))

	recMW := recoverer.Middleware(s.log)
	authMW := auth.Middleware()
	logMW := logging.Middleware(s.log)

	return loader.Middleware(repo, authMW(logMW(recMW(srv))))
}

func (s *Server) Start() error {
	var h http.Handler

	if s.config.Mode == config.ModePostgres {
		h = s.resolverPG()
	} else {
		h = s.resolverLocal()
	}

	http.Handle("/query", h)
	s.log.Info("Server started on http://%s:%s/", s.config.Host, s.config.Port)
	return http.ListenAndServe(s.config.Host+":"+s.config.Port, nil)
}
