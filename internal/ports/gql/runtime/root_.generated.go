// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package runtime

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/darleet/blog-graphql/internal/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
	Subscription() SubscriptionResolver
}

type DirectiveRoot struct {
	IsAuthenticated func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
}

type ComplexityRoot struct {
	Article struct {
		Author    func(childComplexity int) int
		Content   func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		IsClosed  func(childComplexity int) int
		Title     func(childComplexity int) int
		Votes     func(childComplexity int) int
	}

	Comment struct {
		Author    func(childComplexity int) int
		Content   func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Replies   func(childComplexity int) int
		Votes     func(childComplexity int) int
	}

	Mutation struct {
		CreateArticle func(childComplexity int, input model.NewArticle) int
		CreateComment func(childComplexity int, input model.NewComment) int
		DeleteArticle func(childComplexity int, id string) int
		DeleteComment func(childComplexity int, id string) int
		Login         func(childComplexity int, input model.LoginInput) int
		Register      func(childComplexity int, input model.RegisterInput) int
		UpdateArticle func(childComplexity int, input model.UpdateArticle) int
		UpdateComment func(childComplexity int, input *model.UpdateComment) int
		Vote          func(childComplexity int, input model.Vote) int
	}

	Query struct {
		Article      func(childComplexity int, articleID string) int
		ArticlesList func(childComplexity int, after *string, sort *model.Sort) int
		CommentsList func(childComplexity int, articleID string, after *string, sort *model.Sort) int
	}

	Subscription struct {
		NewComment func(childComplexity int, articleID string) int
	}

	User struct {
		AvatarURL func(childComplexity int) int
		ID        func(childComplexity int) int
		Username  func(childComplexity int) int
	}

	VoteCounter struct {
		Value func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Article.author":
		if e.complexity.Article.Author == nil {
			break
		}

		return e.complexity.Article.Author(childComplexity), true

	case "Article.content":
		if e.complexity.Article.Content == nil {
			break
		}

		return e.complexity.Article.Content(childComplexity), true

	case "Article.createdAt":
		if e.complexity.Article.CreatedAt == nil {
			break
		}

		return e.complexity.Article.CreatedAt(childComplexity), true

	case "Article.id":
		if e.complexity.Article.ID == nil {
			break
		}

		return e.complexity.Article.ID(childComplexity), true

	case "Article.isClosed":
		if e.complexity.Article.IsClosed == nil {
			break
		}

		return e.complexity.Article.IsClosed(childComplexity), true

	case "Article.title":
		if e.complexity.Article.Title == nil {
			break
		}

		return e.complexity.Article.Title(childComplexity), true

	case "Article.votes":
		if e.complexity.Article.Votes == nil {
			break
		}

		return e.complexity.Article.Votes(childComplexity), true

	case "Comment.author":
		if e.complexity.Comment.Author == nil {
			break
		}

		return e.complexity.Comment.Author(childComplexity), true

	case "Comment.content":
		if e.complexity.Comment.Content == nil {
			break
		}

		return e.complexity.Comment.Content(childComplexity), true

	case "Comment.createdAt":
		if e.complexity.Comment.CreatedAt == nil {
			break
		}

		return e.complexity.Comment.CreatedAt(childComplexity), true

	case "Comment.id":
		if e.complexity.Comment.ID == nil {
			break
		}

		return e.complexity.Comment.ID(childComplexity), true

	case "Comment.replies":
		if e.complexity.Comment.Replies == nil {
			break
		}

		return e.complexity.Comment.Replies(childComplexity), true

	case "Comment.votes":
		if e.complexity.Comment.Votes == nil {
			break
		}

		return e.complexity.Comment.Votes(childComplexity), true

	case "Mutation.createArticle":
		if e.complexity.Mutation.CreateArticle == nil {
			break
		}

		args, err := ec.field_Mutation_createArticle_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateArticle(childComplexity, args["input"].(model.NewArticle)), true

	case "Mutation.createComment":
		if e.complexity.Mutation.CreateComment == nil {
			break
		}

		args, err := ec.field_Mutation_createComment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateComment(childComplexity, args["input"].(model.NewComment)), true

	case "Mutation.deleteArticle":
		if e.complexity.Mutation.DeleteArticle == nil {
			break
		}

		args, err := ec.field_Mutation_deleteArticle_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteArticle(childComplexity, args["id"].(string)), true

	case "Mutation.deleteComment":
		if e.complexity.Mutation.DeleteComment == nil {
			break
		}

		args, err := ec.field_Mutation_deleteComment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteComment(childComplexity, args["id"].(string)), true

	case "Mutation.login":
		if e.complexity.Mutation.Login == nil {
			break
		}

		args, err := ec.field_Mutation_login_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.Login(childComplexity, args["input"].(model.LoginInput)), true

	case "Mutation.register":
		if e.complexity.Mutation.Register == nil {
			break
		}

		args, err := ec.field_Mutation_register_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.Register(childComplexity, args["input"].(model.RegisterInput)), true

	case "Mutation.updateArticle":
		if e.complexity.Mutation.UpdateArticle == nil {
			break
		}

		args, err := ec.field_Mutation_updateArticle_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateArticle(childComplexity, args["input"].(model.UpdateArticle)), true

	case "Mutation.updateComment":
		if e.complexity.Mutation.UpdateComment == nil {
			break
		}

		args, err := ec.field_Mutation_updateComment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateComment(childComplexity, args["input"].(*model.UpdateComment)), true

	case "Mutation.vote":
		if e.complexity.Mutation.Vote == nil {
			break
		}

		args, err := ec.field_Mutation_vote_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.Vote(childComplexity, args["input"].(model.Vote)), true

	case "Query.article":
		if e.complexity.Query.Article == nil {
			break
		}

		args, err := ec.field_Query_article_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Article(childComplexity, args["articleID"].(string)), true

	case "Query.articlesList":
		if e.complexity.Query.ArticlesList == nil {
			break
		}

		args, err := ec.field_Query_articlesList_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ArticlesList(childComplexity, args["after"].(*string), args["sort"].(*model.Sort)), true

	case "Query.commentsList":
		if e.complexity.Query.CommentsList == nil {
			break
		}

		args, err := ec.field_Query_commentsList_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.CommentsList(childComplexity, args["articleID"].(string), args["after"].(*string), args["sort"].(*model.Sort)), true

	case "Subscription.newComment":
		if e.complexity.Subscription.NewComment == nil {
			break
		}

		args, err := ec.field_Subscription_newComment_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.NewComment(childComplexity, args["articleID"].(string)), true

	case "User.avatarURL":
		if e.complexity.User.AvatarURL == nil {
			break
		}

		return e.complexity.User.AvatarURL(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.username":
		if e.complexity.User.Username == nil {
			break
		}

		return e.complexity.User.Username(childComplexity), true

	case "VoteCounter.value":
		if e.complexity.VoteCounter.Value == nil {
			break
		}

		return e.complexity.VoteCounter.Value(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputLoginInput,
		ec.unmarshalInputNewArticle,
		ec.unmarshalInputNewComment,
		ec.unmarshalInputRegisterInput,
		ec.unmarshalInputUpdateArticle,
		ec.unmarshalInputUpdateComment,
		ec.unmarshalInputVote,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Subscription:
		next := ec._Subscription(ctx, rc.Operation.SelectionSet)

		var buf bytes.Buffer
		return func(ctx context.Context) *graphql.Response {
			buf.Reset()
			data := next(ctx)

			if data == nil {
				return nil
			}
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../../api/article.graphql", Input: `type Article {
    id: ID!
    title: String!
    content: String!
    author: User!
    isClosed: Boolean!
    votes: Int!
    createdAt: Time!
}

input NewArticle {
    title: String!
    content: String!
    isClosed: Boolean!
}

input UpdateArticle {
    id: ID!
    title: String
    content: String
    isClosed: Boolean
}

extend type Query {
    articlesList(after: String, sort: Sort = NEW_DESC): [Article!]
    article(articleID: ID!): Article
}

extend type Mutation {
    createArticle(input: NewArticle!): Article! @isAuthenticated
    updateArticle(input: UpdateArticle!): Article! @isAuthenticated
    deleteArticle(id: ID!): Boolean! @isAuthenticated
}
`, BuiltIn: false},
	{Name: "../../../../api/comment.graphql", Input: `type Comment {
    id: ID!
    content: String!
    author: User!
    votes: Int!
    replies: [Comment!]
    createdAt: Time!
}

input NewComment {
    content: String!
    parentID: ID
}

input UpdateComment {
    id: ID!
    content: String # is nullable due to other fields may be added
}

extend type Query {
    commentsList(articleID: ID!, after: String, sort: Sort = NEW_DESC): [Comment!]
}

extend type Mutation {
    createComment(input: NewComment!): Comment! @isAuthenticated
    updateComment(input: UpdateComment): Comment! @isAuthenticated
    deleteComment(id: ID!): Boolean! @isAuthenticated
}

extend type Subscription {
    newComment(articleID: ID!): Comment! @isAuthenticated
}`, BuiltIn: false},
	{Name: "../../../../api/schema.graphql", Input: `schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

scalar Time

enum Sort {
    TOP_ASC
    TOP_DESC
    NEW_ASC
    NEW_DESC
}

type Query

type Mutation

type Subscription

directive @isAuthenticated on FIELD_DEFINITION
`, BuiltIn: false},
	{Name: "../../../../api/user.graphql", Input: `scalar URL
scalar Email

type User {
    id: ID!
    username: String!
    avatarURL: URL
}

input RegisterInput {
    username: String!
    email: Email!
    password: String!
    avatarURL: URL
}

input LoginInput {
    login: String!
    password: String!
}

extend type Mutation {
    login(input: LoginInput!): Boolean!
    register(input: RegisterInput!): Boolean!
}`, BuiltIn: false},
	{Name: "../../../../api/vote.graphql", Input: `type VoteCounter {
    value: Int!
}

enum VoteValue {
    NONE # used for canceling vote
    UP # +1 to votes
    DOWN # -1 to votes
}

input Vote {
    articleID: ID!
    value: VoteValue!
}

extend type Mutation {
    vote(input: Vote!): VoteCounter! @isAuthenticated
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
