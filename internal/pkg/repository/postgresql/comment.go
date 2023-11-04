package postgresql

import (
	"context"

	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

type CommentRepo struct {
	db db.DBops
}

func NewComments(db db.DBops) *CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) Add(ctx context.Context, entity *repository.Comment) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(
		ctx,
		`insert into comments(text, likes_count, post_id) values ($1, $2, $3) returning id;`,
		entity.Text,
		entity.LikesCount,
		entity.PostID,
	).Scan(&id)
	return id, err
}

// GetMany gets all comments by postID.
func (r *CommentRepo) GetMany(ctx context.Context, id int64) ([]repository.Comment, error) {
	comments := []repository.Comment{}

	err := r.db.GetMany(
		ctx,
		&comments,
		`select id, text, likes_count from comments where post_id = $1;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepo) Remove(ctx context.Context, id int64) error {
	row, err := r.db.Exec(ctx, "delete from comments where id = $1;", id)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return repository.ErrZeroRows
	}
	return nil
}
