package hwservice_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	hwservice "gitlab.ozon.dev/kavkazov/homework-8/internal/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
	mock_repository "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository/mocks"
	"gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
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

		res, err := s.GetPost(ctx, &hw_service.PostRequestWithId{Id: uint64(id)})

		require.NoError(t, err)
		assert.Equal(
			t,
			fixtures.Post().Valid().V(),
			repository.Post{
				ID:         int64(res.GetEntity().Id),
				Heading:    res.GetEntity().Heading,
				Text:       res.GetEntity().Text,
				LikesCount: int(res.GetEntity().LikesCount),
				Comments:   nil,
			},
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
			res, err := s.GetPost(ctx, &hw_service.PostRequestWithId{Id: uint64(id)})
			require.EqualError(t, err, hwservice.ErrNotFound.Error())
			assert.Nil(t, res)
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

			res, err := s.GetPost(ctx, &hw_service.PostRequestWithId{Id: uint64(id)})
			require.EqualError(t, err, hwservice.ErrServer.Error())
			assert.Nil(t, res)
		})
	})
}

func Test_AddPost(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		post = &hw_service.Post{Heading: "heading", Text: "text"}
	)
	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Add(gomock.Any(), &repository.Post{Heading: post.Heading, Text: post.Text}).Return(int64(1), nil)
			},
		)
		defer ctrl.Finish()
		res, err := s.AddPost(ctx, &hw_service.PostRequestWithEntity{Entity: post})
		require.NoError(t, err)
		assert.Equal(
			t,
			post.Heading,
			res.Entity.Heading,
		)
		assert.Equal(
			t,
			post.Text,
			res.Entity.Text,
		)
	})
	t.Run("failure", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Add(gomock.Any(), &repository.Post{Heading: post.Heading, Text: post.Text}).Return(int64(0), assert.AnError)
			},
		)
		defer ctrl.Finish()
		res, err := s.AddPost(ctx, &hw_service.PostRequestWithEntity{Entity: post})
		require.EqualError(t, err, hwservice.ErrServer.Error())
		assert.Nil(t, res)
	})
}

func Test_UpdatePost(t *testing.T) {
	var (
		ctx  = context.Background()
		post = &hw_service.PostRequestWithEntity{Entity: &hw_service.Post{Heading: "heading", Text: "text", Id: 1}}
	)
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().
					Update(gomock.Any(), &repository.Post{ID: int64(post.GetEntity().Id), Text: post.Entity.Text, Heading: post.GetEntity().Heading}).
					Return(nil)
			},
		)
		defer ctrl.Finish()
		_, err := s.UpdatePost(ctx, post)
		assert.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().
						Update(gomock.Any(), &repository.Post{ID: int64(post.GetEntity().Id), Text: post.Entity.Text, Heading: post.GetEntity().Heading}).
						Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()
			_, err := s.UpdatePost(ctx, post)
			assert.EqualError(t, err, hwservice.ErrNotFound.Error())
		})
		t.Run("internal server err", func(t *testing.T) {
			t.Parallel()
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().
						Update(gomock.Any(), &repository.Post{ID: int64(post.GetEntity().Id), Text: post.Entity.Text, Heading: post.GetEntity().Heading}).
						Return(assert.AnError)
				},
			)
			defer ctrl.Finish()
			_, err := s.UpdatePost(ctx, post)
			assert.EqualError(t, err, hwservice.ErrServer.Error())
		})
	})
}

func Test_RemovePost(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
		id  = &hw_service.PostRequestWithId{Id: uint64(1)}
	)
	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				mp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(nil)
			},
		)
		defer ctrl.Finish()

		_, err := s.RemovePost(ctx, id)
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()
			_, err := s.RemovePost(ctx, id)
			assert.EqualError(t, err, hwservice.ErrNotFound.Error())
		})
		t.Run("internal server err", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					mp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(assert.AnError)
				},
			)
			defer ctrl.Finish()
			_, err := s.RemovePost(ctx, id)
			assert.EqualError(t, err, hwservice.ErrServer.Error())
		})
	})
}

func Test_AddComment(t *testing.T) {
	t.Parallel()

	var (
		ctx     = context.Background()
		comment = &hw_service.CommentRequestWithEntity{Entity: &hw_service.Comment{Text: "test"}, PostId: 1}
	)
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Add(gomock.Any(), &repository.Comment{Text: comment.Entity.Text, PostID: int64(comment.PostId)}).Return(int64(1), nil)
			},
		)
		defer ctrl.Finish()
		res, err := s.AddComment(ctx, comment)
		require.NoError(t, err)
		assert.Equal(t, comment.Entity.Text, res.Entity.Text)
	})
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Add(gomock.Any(), &repository.Comment{Text: comment.Entity.Text, PostID: int64(comment.PostId)}).Return(int64(0), assert.AnError)
			},
		)
		defer ctrl.Finish()
		res, err := s.AddComment(ctx, comment)
		require.EqualError(t, err, hwservice.ErrServer.Error())
		assert.Nil(t, res)
	})
}

func Test_RemoveComment(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = &hw_service.CommentRequestWithId{Id: 1}
	)

	t.Run("ok", func(t *testing.T) {
		s, ctrl := setUp(
			t,
			func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
				cp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(nil)
			},
		)
		defer ctrl.Finish()

		_, err := s.RemoveComment(ctx, id)
		assert.NoError(t, err)
	})
	t.Run("failure", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					cp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(repository.ErrZeroRows)
				},
			)
			defer ctrl.Finish()

			_, err := s.RemoveComment(ctx, id)
			assert.EqualError(t, err, hwservice.ErrNotFound.Error())
		})
		t.Run("internal server err", func(t *testing.T) {
			s, ctrl := setUp(
				t,
				func(mp *mock_repository.MockPostsRepo, cp *mock_repository.MockCommentsRepo) {
					cp.EXPECT().Remove(gomock.Any(), int64(id.Id)).Return(assert.AnError)
				},
			)
			defer ctrl.Finish()

			_, err := s.RemoveComment(ctx, id)
			assert.EqualError(t, err, hwservice.ErrServer.Error())
		})
	})
}
