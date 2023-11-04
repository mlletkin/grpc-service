package postgresql

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db/mocks"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

type postsRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.PostsRepo
	mockDb *mock_database.MockDBops
}

type mockRow struct{}

func (m *mockRow) Scan(dest ...interface{}) error {
	temp := 1
	dest[0] = &temp
	return nil
}

func setUp(t *testing.T) *postsRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDBops(ctrl)
	repo := NewPosts(mockDb)
	return &postsRepoFixture{
		ctrl:   ctrl,
		mockDb: mockDb,
		repo:   repo,
	}
}

func (p *postsRepoFixture) tearDown() {
	p.ctrl.Finish()
}
