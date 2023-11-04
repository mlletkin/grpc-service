package server_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/mocks"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
)

func setUp(
	t *testing.T,
	expects func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo),
) (*server.Server, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mp := mock_repository.NewMockPostsRepo(ctrl)
	cp := mock_repository.NewMockCommentsRepo(ctrl)
	expects(mp, cp)
	s := server.NewServer(mp, cp)
	return s, ctrl
}
