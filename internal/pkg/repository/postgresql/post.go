package postgresql

import (
	"context"

	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/db"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

type PostRepo struct {
	db db.DBops
}

func NewPosts(db db.DBops) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) Add(ctx context.Context, entity *repository.Post) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(
		ctx,
		`insert into posts(heading, text, likes_count) values ($1, $2, $3) returning id;`,
		entity.Heading,
		entity.Text,
		entity.LikesCount,
	).Scan(&id)
	return id, err
}

func (r *PostRepo) GetByID(ctx context.Context, id int64) (*repository.Post, error) {
	var post repository.Post
	err := r.db.Get(ctx,
		&post, "select id, heading, text, likes_count from posts where id=$1;", id)
	if err != nil {
		if err.Error() == "scanning one: no rows in result set" {
			return nil, repository.ErrZeroRows
		}
		return nil, err

	}
	return &post, nil
}

func (r *PostRepo) Remove(ctx context.Context, id int64) error {
	row, err := r.db.Exec(ctx, "delete from posts where id = $1;", id)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return repository.ErrZeroRows
	}
	return nil
}

func (r *PostRepo) Update(ctx context.Context, entity *repository.Post) error {
	row, err := r.db.Exec(
		ctx,
		`update posts set
    heading=$1,
    text=$2
    where id=$3;
    `,
		entity.Heading,
		entity.Text,
		entity.ID,
	)
	if err != nil {
		return err
	}
	if row.RowsAffected() == 0 {
		return repository.ErrZeroRows
	}
	return nil
}
