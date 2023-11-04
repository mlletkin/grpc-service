package postgresql

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

func TestPostRepo_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = int64(1)
	)
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().
			Get(gomock.Any(), gomock.Any(), "select id, heading, text, likes_count from posts where id=$1;", gomock.Any()).
			Return(nil)
		// act
		post, err := s.repo.GetByID(ctx, id)
		// assert

		require.NoError(t, err)
		assert.Equal(t, int64(0), post.ID)
	})
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().
				Get(gomock.Any(), gomock.Any(), "select id, heading, text, likes_count from posts where id=$1;", gomock.Any()).
				Return(repository.ErrZeroRows)

			post, err := s.repo.GetByID(ctx, id)

			require.EqualError(t, err, repository.ErrZeroRows.Error())
			assert.Nil(t, post)
		})
		t.Run("internal err", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().
				Get(gomock.Any(), gomock.Any(), "select id, heading, text, likes_count from posts where id=$1;", gomock.Any()).
				Return(assert.AnError)
			// act
			post, err := s.repo.GetByID(ctx, id)
			// assert
			require.EqualError(t, err, "assert.AnError general error for testing")

			assert.Nil(t, post)
		})
	})
}
