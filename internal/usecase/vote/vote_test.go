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

func (s *VoteTestSuite) TestProcessVoteInsert() {
	s.repo.On("GetUserVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).
		Return(model.VoteValueNone, errors.NotFound).
		Once()

	s.repo.On("InsertVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.Vote")).
		Return(nil).
		Once()

	s.repo.On("GetArticleVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteArticle(ctx, model.Vote{
		ArticleID: "1234",
		Value:     model.VoteValueUp,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, got)
}

func (s *VoteTestSuite) TestProcessVoteUpdate() {
	s.repo.On("GetUserVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).
		Return(model.VoteValueUp, nil).
		Once()

	s.repo.On("SetVote", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("model.Vote")).
		Return(nil).
		Once()

	s.repo.On("GetArticleVotes", mock.Anything,
		mock.AnythingOfType("string")).
		Return(-1, nil).
		Once()

	ctx := utils.SetUserID(context.Background(), "4321")
	got, err := s.uc.VoteArticle(ctx, model.Vote{
		ArticleID: "1234",
		Value:     model.VoteValueDown,
	})
	require.NoError(s.T(), err)
	require.Equal(s.T(), -1, got)
}

func (s *VoteTestSuite) TestProcessVoteNoID() {
	got, err := s.uc.VoteArticle(context.Background(), model.Vote{
		ArticleID: "1234",
		Value:     model.VoteValueUp,
	})
	require.ErrorIs(s.T(), err, errors.Unauthorized)
	require.Equal(s.T(), 0, got)
}
