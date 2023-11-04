//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/kavkazov/homework-8/tests/fixtures"
)

func TestCreatePost(t *testing.T) {
	ctx := context.Background()
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()

		// arrange
		repo := postgresql.NewPosts(db.DB)

		// act
		resp, err := repo.Add(ctx, fixtures.Post().Valid().P())

		// assert
		require.NoError(t, err)
		assert.NotZero(t, resp)
	})
}

func TestGetPost(t *testing.T) {
	var (
		ctx       = context.Background()
		postValid = fixtures.Post().Valid().P()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t, postValid)
		defer db.TearDown()

		// arrange
		repo := postgresql.NewPosts(db.DB)

		resp, err := repo.Add(ctx, postValid)

		// assert
		require.NoError(t, err)
		assert.NotZero(t, resp)

		respGet, err := repo.GetByID(ctx, resp)

		// assert
		require.NoError(t, err)
		assert.Equal(t, postValid.Heading, respGet.Heading)
		assert.Equal(t, postValid.LikesCount, respGet.LikesCount)
	})
}
