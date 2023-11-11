package hwservice

import (
	"context"
	"errors"
	"log"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
	pb "gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrServer   = errors.New("server error")
	ErrNotFound = errors.New("not found")
)

type Implementation struct {
	pb.UnimplementedHomeworkServiceServer
	postRepo    repository.PostsRepo
	commentRepo repository.CommentsRepo
}

func New(post repository.PostsRepo, comment repository.CommentsRepo) *Implementation {
	return &Implementation{
		postRepo:    post,
		commentRepo: comment,
	}
}

func (i *Implementation) AddComment(ctx context.Context, req *pb.CommentRequestWithEntity) (*pb.CommentResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: AddComment")
	defer span.Finish()

	entity := req.GetEntity()
	post_id := req.GetPostId()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: AddComment")

	comment := &repository.Comment{
		Text:       entity.GetText(),
		PostID:     int64(post_id),
		LikesCount: int(entity.GetLikesCount()),
	}
	id, err := i.commentRepo.Add(ctx, comment)
	if err != nil {
		return nil, ErrServer
	}
	comment.ID = id

	coreSpan.Finish()

	return &pb.CommentResponseWithEntity{
		Entity: &pb.Comment{
			Id:         uint64(comment.ID),
			Text:       comment.Text,
			LikesCount: uint64(comment.LikesCount),
		},
	}, nil
}

func (i *Implementation) AddPost(ctx context.Context, req *pb.PostRequestWithEntity) (*pb.PostResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: AddPost")
	defer span.Finish()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: AddPost")

	post := &repository.Post{
		Heading: req.GetEntity().GetHeading(),
		Text:    req.Entity.GetText(),
	}
	id, err := i.postRepo.Add(ctx, post)
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, ErrNotFound
		} else {
			return nil, ErrServer

		}
	}
	post.ID = id

	coreSpan.Finish()

	comms := Mapper(post.Comments, func(comment repository.Comment) *pb.Comment {
		return &pb.Comment{
			Id:         uint64(comment.ID),
			Text:       comment.Text,
			LikesCount: uint64(comment.LikesCount),
		}
	})

	return &pb.PostResponseWithEntity{
		Entity: &pb.Post{
			Id:         uint64(post.ID),
			Heading:    post.Heading,
			Text:       post.Text,
			LikesCount: uint64(post.LikesCount),
			Comments:   comms,
		},
	}, nil
}

func (i *Implementation) GetPost(ctx context.Context, id *pb.PostRequestWithId) (*pb.PostResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: GetPost")
	defer span.Finish()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: GetPost")

	post, err := i.postRepo.GetByID(ctx, int64(id.GetId()))
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, ErrNotFound
		}
		return nil, ErrServer
	}

	comments, err := i.commentRepo.GetMany(ctx, int64(id.GetId()))
	if err != nil {
		log.Println(err)
		comments = nil
	}
	post.Comments = comments
	coreSpan.Finish()

	comms := Mapper(post.Comments, func(comment repository.Comment) *pb.Comment {
		return &pb.Comment{
			Id:         uint64(comment.ID),
			Text:       comment.Text,
			LikesCount: uint64(comment.LikesCount),
		}
	})

	return &pb.PostResponseWithEntity{
		Entity: &pb.Post{
			Id:         uint64(post.ID),
			Text:       post.Text,
			Heading:    post.Heading,
			LikesCount: uint64(post.LikesCount),
			Comments:   comms,
		},
	}, nil

}

func (i *Implementation) RemoveComment(ctx context.Context, id *pb.CommentRequestWithId) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: RemoveComment")
	defer span.Finish()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: RemoveComment")

	err := i.commentRepo.Remove(ctx, int64(id.GetId()))
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, ErrNotFound
		} else {
			return nil, ErrServer
		}
	}
	coreSpan.Finish()

	return &emptypb.Empty{}, nil
}

func (i *Implementation) RemovePost(ctx context.Context, id *pb.PostRequestWithId) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: RemovePost")
	defer span.Finish()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: RemovePost")

	err := i.postRepo.Remove(ctx, int64(id.GetId()))
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, ErrNotFound
		} else {
			return nil, ErrServer
		}
	}
	coreSpan.Finish()

	return &emptypb.Empty{}, nil
}

func (i *Implementation) UpdatePost(ctx context.Context, req *pb.PostRequestWithEntity) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: UpdatePost")
	defer span.Finish()

	entity := req.GetEntity()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: UpdatePost")

	if err := i.postRepo.Update(
		ctx,
		&repository.Post{Heading: entity.Heading, Text: entity.Text, ID: int64(entity.Id)},
	); err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, ErrNotFound
		} else {
			return nil, ErrServer
		}
	}
	coreSpan.Finish()

	return &emptypb.Empty{}, nil
}
