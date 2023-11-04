package fixtures

import "gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"

type PostBuilder struct {
	instance *repository.Post
}

func Post() *PostBuilder {
	return &PostBuilder{instance: &repository.Post{}}
}

func (b *PostBuilder) ID(v int64) *PostBuilder {
	b.instance.ID = v
	return b
}

func (b *PostBuilder) Heading(v string) *PostBuilder {
	b.instance.Heading = v
	return b
}

func (b *PostBuilder) Text(v string) *PostBuilder {
	b.instance.Text = v
	return b
}

func (b *PostBuilder) LikesCount(v int) *PostBuilder {
	b.instance.LikesCount = v
	return b
}

func (b *PostBuilder) Comments(v []repository.Comment) *PostBuilder {
	b.instance.Comments = v
	return b
}

func (b *PostBuilder) P() *repository.Post {
	return b.instance
}

func (b *PostBuilder) V() repository.Post {
	return *b.instance
}

func (b *PostBuilder) Valid() *PostBuilder {
	return b.Heading("test").ID(1)
}
