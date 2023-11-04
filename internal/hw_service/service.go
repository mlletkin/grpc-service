package hwservice

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
	pb "gitlab.ozon.dev/kavkazov/homework-8/pkg/hw_service"
	"gitlab.ozon.dev/kavkazov/homework-8/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrServer = errors.New("server error")
)

type Implementation struct {
	pb.UnimplementedHomeworkServiceServer
	server *server.Server
}

func New(server *server.Server) *Implementation {
	return &Implementation{
		server: server,
	}
}

func (i *Implementation) AddComment(ctx context.Context, comment *pb.CommentRequestWithEntity) (*pb.CommentResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: AddComment")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "AddComment")))
	logger.Infof(ctx, "%v", time.Now())

	entity := comment.GetEntity()
	post_id := comment.GetPostId()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: AddComment")
	data, status := i.server.AddComment(ctx, &repository.Comment{
		Text:       entity.GetText(),
		PostID:     int64(post_id),
		LikesCount: int(entity.GetLikesCount()),
	})
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}

	return &pb.CommentResponseWithEntity{
		Entity: &pb.Comment{
			Id:         uint64(data.ID),
			Text:       data.Text,
			LikesCount: uint64(data.LikesCount),
		},
	}, nil
}

func (i *Implementation) AddPost(ctx context.Context, post *pb.PostRequestWithEntity) (*pb.PostResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: AddPost")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "AddPost")))
	logger.Infof(ctx, "%v", time.Now())

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: AddPost")
	data, status := i.server.AddPost(ctx, &server.AddPostRequest{
		Heading: post.GetEntity().GetHeading(),
		Text:    post.GetEntity().GetText(),
	})
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}

	comms := mapper(data.Comments, func(comment repository.Comment) *pb.Comment {
		return &pb.Comment{
			Id:         uint64(comment.ID),
			Text:       comment.Text,
			LikesCount: uint64(comment.LikesCount),
		}
	})

	return &pb.PostResponseWithEntity{
		Entity: &pb.Post{
			Id:         uint64(data.ID),
			Heading:    data.Heading,
			Text:       data.Text,
			LikesCount: uint64(data.LikesCount),
			Comments:   comms,
		},
	}, nil
}

func (i *Implementation) GetPost(ctx context.Context, id *pb.PostRequestWithId) (*pb.PostResponseWithEntity, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: GetPost")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "GetPost")))
	logger.Infof(ctx, "%v", time.Now())

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: GetPost")
	data, status := i.server.GetPost(ctx, int64(id.GetId()))
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}

	comms := mapper(data.Comments, func(comment repository.Comment) *pb.Comment {
		return &pb.Comment{
			Id:         uint64(comment.ID),
			Text:       comment.Text,
			LikesCount: uint64(comment.LikesCount),
		}
	})

	return &pb.PostResponseWithEntity{
		Entity: &pb.Post{
			Id:         uint64(data.ID),
			Text:       data.Text,
			Heading:    data.Heading,
			LikesCount: uint64(data.LikesCount),
			Comments:   comms,
		},
	}, nil

}

func (i *Implementation) RemoveComment(ctx context.Context, id *pb.CommentRequestWithId) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: RemoveComment")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "RemoveComment")))
	logger.Infof(ctx, "%v", time.Now())

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: RemoveComment")
	status := i.server.RemoveComment(ctx, int64(id.GetId()))
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) RemovePost(ctx context.Context, id *pb.PostRequestWithId) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: RemovePost")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "RemovePost")))
	logger.Infof(ctx, "%v", time.Now())

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: RemovePost")
	status := i.server.RemovePost(ctx, int64(id.GetId()))
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}
	return &emptypb.Empty{}, nil
}

func (i *Implementation) UpdatePost(ctx context.Context, post *pb.PostRequestWithEntity) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "homework_service: UpdatePost")
	defer span.Finish()

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "UpdatePost")))
	logger.Infof(ctx, "%v", time.Now())

	entity := post.GetEntity()

	coreSpan, ctx := opentracing.StartSpanFromContext(ctx, "core: UpdatePost")
	status := i.server.UpdatePost(ctx, &server.UpdatePostRequest{
		ID: int64(entity.GetId()),
		AddPostRequest: server.AddPostRequest{
			Heading: entity.GetHeading(),
			Text:    entity.GetText(),
		},
	})
	coreSpan.Finish()

	if status != http.StatusOK {
		return nil, ErrServer
	}
	return &emptypb.Empty{}, nil
}
