package hwservice_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	hwservice "gitlab.ozon.dev/kavkazov/homework-8/internal/hw_service"
	mock_repository "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/mocks"
)

func setUp(
	t *testing.T,
	expects func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo),
) (*hwservice.Implementation, *gomock.Controller) {
	ctrl := gomock.NewController(t)

	mp := mock_repository.NewMockPostsRepo(ctrl)
	cp := mock_repository.NewMockCommentsRepo(ctrl)
	expects(mp, cp)
	s := hwservice.New(mp, cp)

	return s, ctrl
}
