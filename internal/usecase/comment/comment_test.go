package comment

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/usecase/comment/mocks"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var comment = &model.Comment{
	ID:        "1234",
	Content:   "Some test comment",
	UserID:    "4321",
	CreatedAt: time.Now(),
}

type CommentTestSuite struct {
	suite.Suite
	repo *mocks.Repository
	uc   *Usecase
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(CommentTestSuite))
}

func (s *CommentTestSuite) SetupSuite() {
	s.repo = mocks.NewRepository(s.T())
	s.uc = NewUsecase(s.repo)
}

func (s *CommentTestSuite) TestCreate() {
	input := model.NewComment{
		Content:   comment.Content,
		ArticleID: "1234",
	}

	s.repo.On("IsArticleClosed", mock.Anything,
		mock.AnythingOfType("string")).
		Return(false, nil).
		Once()

	s.repo.On("CreateComment", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.NewComment")).
		Return(comment, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Create(ctx, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), comment, got)
}

func (s *CommentTestSuite) TestCreateClosed() {
	input := model.NewComment{
		Content:   comment.Content,
		ArticleID: "1234",
	}

	s.repo.On("IsArticleClosed", mock.Anything,
		mock.AnythingOfType("string")).
		Return(true, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Create(ctx, input)
	require.ErrorIs(s.T(), err, errors.Forbidden)
	require.Nil(s.T(), got)
}

func (s *CommentTestSuite) TestCreateNoID() {
	input := model.NewComment{
		Content:   comment.Content,
		ArticleID: "1234",
	}

	got, err := s.uc.Create(context.Background(), input)
	require.ErrorIs(s.T(), err, errors.Unauthorized)
	require.Nil(s.T(), got)
}

func (s *CommentTestSuite) TestUpdate() {
	input := model.UpdateComment{
		ID:      "1234",
		Content: &comment.Content,
	}

	s.repo.On("GetCommentAuthorID", mock.Anything,
		mock.AnythingOfType("string")).
		Return("4321", nil).
		Once()

	s.repo.On("UpdateComment", mock.Anything,
		mock.AnythingOfType("model.UpdateComment")).
		Return(comment, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Update(ctx, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), comment, got)
}

func (s *CommentTestSuite) TestUpdateWrongAuthor() {
	input := model.UpdateComment{
		ID:      "1234",
		Content: &comment.Content,
	}

	s.repo.On("GetCommentAuthorID", mock.Anything,
		mock.AnythingOfType("string")).
		Return("1337", nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.Update(ctx, input)
	require.ErrorIs(s.T(), err, errors.Forbidden)
	require.Nil(s.T(), got)
}

func (s *CommentTestSuite) TestDelete() {
	s.repo.On("GetCommentAuthorID", mock.Anything,
		mock.AnythingOfType("string")).
		Return("4321", nil).
		Once()

	s.repo.On("DeleteComment", mock.Anything,
		mock.AnythingOfType("string")).
		Return(true, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	ok, err := s.uc.Delete(ctx, "1234")
	require.NoError(s.T(), err)
	require.True(s.T(), ok)
}

func (s *CommentTestSuite) TestDeleteWrongAuthor() {
	s.repo.On("GetCommentAuthorID", mock.Anything,
		mock.AnythingOfType("string")).
		Return("1337", nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	ok, err := s.uc.Delete(ctx, "1234")
	require.ErrorIs(s.T(), err, errors.Forbidden)
	require.False(s.T(), ok)
}
