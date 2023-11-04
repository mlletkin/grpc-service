package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
	mock_repository "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/mocks"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
	"gitlab.ozon.dev/kavkazov/homework-8/tests/fixtures"
)

func Test_GetPost(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().GetMany(gomock.Any(), int64(id)).Return([]repository.Comment{
					{
						ID:     1,
						PostID: 1,
						Text:   "Test",
					},
				}, nil)
				mp.EXPECT().
					GetByID(gomock.Any(), int64(id)).
					Return(fixtures.Post().Valid().P(), nil)
			},
		)
		defer ctrl.Finish()

		res, code := s.GetPost(ctx, int64(id))

		require.Equal(t, http.StatusOK, code)
		assert.Equal(
			t,
			fixtures.Post().Valid().P(),
			res,
		)
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()
			res, code := s.GetPost(ctx, int64(id))
			require.Equal(t, http.StatusNotFound, code)
			assert.Equal(t, []byte(nil), res)
		})
		t.Run("internal server err", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().GetByID(gomock.Any(), int64(id)).Return(nil, assert.AnError)
				},
			)
			defer ctrl.Finish()

			res, code := s.GetPost(ctx, int64(id))
			require.Equal(t, http.StatusInternalServerError, code)
			assert.Equal(t, nil, res)
		})
	})
}

func Test_AddPost(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		post = &repository.Post{
			Heading: "heading",
			Text:    "text",
		}
	)
	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Add(gomock.Any(), post).Return(int64(1), nil)
			},
		)
		defer ctrl.Finish()
		data, status := s.AddPost(ctx, &server.AddPostRequest{Heading: "heading", Text: "text"})
		require.Equal(t, http.StatusOK, status)
		assert.Equal(
			t,
			post,
			data,
		)
	})
	t.Run("failure", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Add(gomock.Any(), post).Return(int64(0), assert.AnError)
			},
		)
		defer ctrl.Finish()
		data, status := s.AddPost(ctx, &server.AddPostRequest{Heading: "heading", Text: "text"})
		require.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, nil, data)
	})
}

func Test_UpdatePost(t *testing.T) {
	var (
		ctx  = context.Background()
		post = &server.UpdatePostRequest{1, server.AddPostRequest{Heading: "heading", Text: "Text"}}
	)
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().
					Update(gomock.Any(), &repository.Post{ID: post.ID, Text: post.Text, Heading: post.Heading}).
					Return(nil)
			},
		)
		defer ctrl.Finish()
		status := s.UpdatePost(ctx, post)
		assert.Equal(t, http.StatusOK, status)
	})
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().
						Update(gomock.Any(), &repository.Post{ID: post.ID, Text: post.Text, Heading: post.Heading}).
						Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()
			status := s.UpdatePost(ctx, post)
			assert.Equal(t, http.StatusNotFound, status)
		})
		t.Run("internal server err", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().
						Update(gomock.Any(), &repository.Post{ID: post.ID, Text: post.Text, Heading: post.Heading}).
						Return(assert.AnError)
				},
			)
			defer ctrl.Finish()
			status := s.UpdatePost(ctx, post)
			assert.Equal(t, http.StatusInternalServerError, status)
		})
	})
}

func Test_RemovePos(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		id  = int64(1)
	)
	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Remove(gomock.Any(), id).Return(nil)
			},
		)
		defer ctrl.Finish()

		status := s.RemovePost(ctx, id)
		assert.Equal(t, http.StatusOK, status)
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().Remove(gomock.Any(), id).Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()
			status := s.RemovePost(ctx, id)
			assert.Equal(t, http.StatusNotFound, status)
		})
		t.Run("internal server err", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().Remove(gomock.Any(), id).Return(assert.AnError)
				},
			)
			defer ctrl.Finish()
			status := s.RemovePost(ctx, id)
			assert.Equal(t, http.StatusInternalServerError, status)
		})
	})
}

func Test_AddComment(t *testing.T) {
	t.Parallel()

	var (
		ctx     = context.Background()
		comment = &repository.Comment{PostID: 1, Text: "test"}
	)
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Add(gomock.Any(), comment).Return(int64(1), nil)
			},
		)
		defer ctrl.Finish()
		data, status := s.AddComment(ctx, comment)
		require.Equal(t, http.StatusOK, status)
		assert.Equal(t, comment, data)
	})
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Add(gomock.Any(), comment).Return(int64(0), assert.AnError)
			},
		)
		defer ctrl.Finish()
		data, status := s.AddComment(ctx, comment)
		require.Equal(t, http.StatusInternalServerError, status)
		assert.Equal(t, nil, data)
	})
}

func Test_RemoveComment(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = int64(1)
	)

	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Remove(gomock.Any(), id).Return(nil)
			},
		)
		defer ctrl.Finish()

		status := s.RemoveComment(ctx, id)
		assert.Equal(t, http.StatusOK, status)
	})
	t.Run("failure", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					cp.EXPECT().Remove(gomock.Any(), id).Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()

			status := s.RemoveComment(ctx, id)
			assert.Equal(t, http.StatusNotFound, status)
		})
		t.Run("internal server err", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					cp.EXPECT().Remove(gomock.Any(), id).Return(assert.AnError)
				},
			)
			defer ctrl.Finish()

			status := s.RemoveComment(ctx, id)
			assert.Equal(t, http.StatusInternalServerError, status)
		})
	})
}
