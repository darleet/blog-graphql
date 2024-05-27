package vote

import (
	"context"
	"github.com/darleet/blog-graphql/internal/model"
	"github.com/darleet/blog-graphql/internal/usecase/vote/mocks"
	"github.com/darleet/blog-graphql/pkg/errors"
	"github.com/darleet/blog-graphql/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VoteTestSuite struct {
	suite.Suite
	repo *mocks.Repository
	uc   *Usecase
}

func TestVoteTestSuite(t *testing.T) {
	suite.Run(t, new(VoteTestSuite))
}

func (s *VoteTestSuite) SetupSuite() {
	s.repo = mocks.NewRepository(s.T())
	s.uc = NewUsecase(s.repo)
}

func (s *VoteTestSuite) TestVoteArticleInsert() {
	s.repo.On("SetArticleVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.VoteArticle")).
		Return(nil).
		Once()

	s.repo.On("GetArticleVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteArticle(ctx, model.VoteArticle{
		ArticleID: "1234",
		Value:     model.VoteValueUp,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, got)
}

func (s *VoteTestSuite) TestVoteArticleUpdate() {
	s.repo.On("SetArticleVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.VoteArticle")).
		Return(nil).
		Once()

	s.repo.On("GetArticleVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(-1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteArticle(ctx, model.VoteArticle{
		ArticleID: "1234",
		Value:     model.VoteValueDown,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), -1, got)
}

func (s *VoteTestSuite) TestVoteArticleNoID() {
	got, err := s.uc.VoteArticle(context.Background(), model.VoteArticle{
		ArticleID: "1234",
		Value:     model.VoteValueUp,
	})
	require.ErrorIs(s.T(), err, errors.Unauthorized)
	require.Equal(s.T(), 0, got)
}

func (s *VoteTestSuite) TestVoteCommentInsert() {
	s.repo.On("SetCommentVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.VoteComment")).
		Return(nil).
		Once()

	s.repo.On("GetCommentVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteComment(ctx, model.VoteComment{
		CommentID: "1234",
		Value:     model.VoteValueUp,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, got)
}

func (s *VoteTestSuite) TestVoteCommentUpdate() {
	s.repo.On("SetCommentVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.VoteComment")).
		Return(nil).
		Once()

	s.repo.On("GetCommentVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(-1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteComment(ctx, model.VoteComment{
		CommentID: "1234",
		Value:     model.VoteValueDown,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), -1, got)
}

func (s *VoteTestSuite) TestVoteCommentNoID() {
	got, err := s.uc.VoteComment(context.Background(), model.VoteComment{
		CommentID: "1234",
		Value:     model.VoteValueUp,
	})
	require.ErrorIs(s.T(), err, errors.Unauthorized)
	require.Equal(s.T(), 0, got)
}
