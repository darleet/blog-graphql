package article

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/usecase/article/mocks"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var article = &model.Article{
	ID:      "1234",
	Title:   "Test Article",
	Content: "Some test article",
	Author: &model.User{
		ID:       "4321",
		Username: "username",
	},
	IsClosed:  false,
	Votes:     0,
	CreatedAt: time.Now(),
}

type ArticleTestSuite struct {
	suite.Suite
	repo *mocks.Repository
	uc   *Usecase
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}

func (s *ArticleTestSuite) SetupSuite() {
	s.repo = mocks.NewRepository(s.T())
	s.uc = NewUsecase(s.repo)
}

func (s *ArticleTestSuite) TestCreate() {
	input := model.NewArticle{
		Title:   "Test Article",
		Content: "Some test article",
	}

	s.repo.On("Create", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.NewArticle")).
		Return(article, nil)

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Create(ctx, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), article, got)
}

func (s *ArticleTestSuite) TestCreateNoID() {
	input := model.NewArticle{
		Title:   "Test Article",
		Content: "Some test article",
	}

	got, err := s.uc.Create(context.Background(), input)
	require.ErrorIs(s.T(), err, errors.Unauthorized)
	require.Nil(s.T(), got)
}

func (s *ArticleTestSuite) TestUpdate() {
	title := "Test Article"
	content := "Some test article"
	isClosed := false

	input := model.UpdateArticle{
		ID:       "1234",
		Title:    &title,
		Content:  &content,
		IsClosed: &isClosed,
	}

	s.repo.On("GetByID", mock.Anything,
		mock.AnythingOfType("string")).
		Return(article, nil)

	s.repo.On("Update", mock.Anything,
		mock.AnythingOfType("model.UpdateArticle")).
		Return(article, nil)

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Update(ctx, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), article, got)
}

func (s *ArticleTestSuite) TestUpdateWrongAuthor() {
	title := "Test Article"
	content := "Some test article"
	isClosed := false

	input := model.UpdateArticle{
		ID:       "1234",
		Title:    &title,
		Content:  &content,
		IsClosed: &isClosed,
	}

	s.repo.On("GetByID", mock.Anything,
		mock.AnythingOfType("string")).
		Return(article, nil)

	ctx := utils.SetUserID(context.Background(), "1337")
	got, err := s.uc.Update(ctx, input)
	require.ErrorIs(s.T(), err, errors.Forbidden)
	require.Nil(s.T(), got)
}

func (s *ArticleTestSuite) TestDelete() {
	s.repo.On("GetByID", mock.Anything,
		mock.AnythingOfType("string")).
		Return(article, nil)

	s.repo.On("Delete", mock.Anything,
		mock.AnythingOfType("string")).
		Return(true, nil)

	ctx := utils.SetUserID(context.Background(), "4321")
	ok, err := s.uc.Delete(ctx, "1234")
	require.NoError(s.T(), err)
	require.True(s.T(), ok)
}

func (s *ArticleTestSuite) TestDeleteWrongAuthor() {
	s.repo.On("GetByID", mock.Anything,
		mock.AnythingOfType("string")).
		Return(article, nil)

	ctx := utils.SetUserID(context.Background(), "1337")
	ok, err := s.uc.Delete(ctx, "1234")
	require.ErrorIs(s.T(), err, errors.Forbidden)
	require.False(s.T(), ok)
}

func (s *ArticleTestSuite) TestGetList() {
	s.repo.On("GetList", mock.Anything,
		mock.AnythingOfType("*string"),
		mock.AnythingOfType("*model.Sort")).
		Return([]*model.Article{article}, nil)

	got, err := s.uc.GetList(context.Background(), nil, nil)
	require.NoError(s.T(), err)
	require.Equal(s.T(), []*model.Article{article}, got)
}

func (s *ArticleTestSuite) TestGetByID() {
	s.repo.On("GetByID", mock.Anything,
		mock.AnythingOfType("string")).
		Return(article, nil)

	got, err := s.uc.GetByID(context.Background(), "1234")
	require.NoError(s.T(), err)
	require.Equal(s.T(), article, got)
}
