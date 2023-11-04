//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import (
	"context"
)

type PostsRepo interface {
	Add(ctx context.Context, entity *Post) (int64, error)
	GetByID(ctx context.Context, id int64) (*Post, error)
	Remove(ctx context.Context, id int64) error
	Update(ctx context.Context, entity *Post) error
}

type CommentsRepo interface {
	Add(ctx context.Context, entity *Comment) (int64, error)
	GetMany(ctx context.Context, id int64) ([]Comment, error)
	Remove(ctx context.Context, id int64) error
}
