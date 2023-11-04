package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/repository"
)

type AddPostRequest struct {
	Heading string `json:"heading"`
	Text    string `json:"text"`
}
type AddCommentRequest struct {
	Text string `json:"text"`
}

type UpdatePostRequest struct {
	ID int64 `json:"id"`
	AddPostRequest
}

type Server struct {
	postRepo    repository.PostsRepo
	commentRepo repository.CommentsRepo
}

func NewServer(post repository.PostsRepo, comment repository.CommentsRepo) *Server {
	return &Server{postRepo: post, commentRepo: comment}
}

func (s *Server) ParsePostBody(req *http.Request) (*AddPostRequest, int) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	var postReq AddPostRequest
	if err = json.Unmarshal(body, &postReq); err != nil {
		return nil, http.StatusBadRequest
	}
	return &postReq, http.StatusOK
}

func (s *Server) AddPost(ctx context.Context, postReq *AddPostRequest) ([]byte, int) {
	post := &repository.Post{
		Heading: postReq.Heading,
		Text:    postReq.Text,
	}
	id, err := s.postRepo.Add(ctx, post)
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, http.StatusNotFound
		} else {
			log.Println(err)
			return nil, http.StatusInternalServerError

		}
	}
	post.ID = id
	postJson, err := json.Marshal(post)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return postJson, http.StatusOK
}

func (s *Server) ParseGetID(req *http.Request) (int64, int) {
	key := req.URL.Query().Get(PostIDKey)
	if key == "" {
		return 0, http.StatusBadRequest
	}
	postID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return postID, http.StatusOK
}

func (s *Server) GetPost(ctx context.Context, postID int64) ([]byte, int) {
	post, err := s.postRepo.GetByID(ctx, postID)
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}

	comments, err := s.commentRepo.GetMany(ctx, postID)
	if err != nil {
		log.Println(err)
		comments = nil
	}
	post.Comments = comments

	postJson, err := json.Marshal(post)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return postJson, http.StatusOK
}

func (s *Server) ParsePostBodyUpdate(req *http.Request) (*UpdatePostRequest, int) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	var postReq *UpdatePostRequest
	if err = json.Unmarshal(body, postReq); err != nil {
		return nil, http.StatusBadRequest
	}
	return postReq, http.StatusOK
}

func (s *Server) UpdatePost(ctx context.Context, postReq *UpdatePostRequest) int {
	if err := s.postRepo.Update(
		ctx,
		&repository.Post{Heading: postReq.Heading, Text: postReq.Text, ID: postReq.ID},
	); err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return http.StatusNotFound
		} else {
			log.Println(err)
			return http.StatusInternalServerError
		}
	}

	return http.StatusOK
}

func (s *Server) ParsePathID(req *http.Request) (int64, int) {
	key, ok := mux.Vars(req)[PostIDKey]
	if !ok {
		return int64(0), http.StatusBadRequest
	}
	postID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return int64(0), http.StatusBadRequest
	}
	return postID, http.StatusOK
}

func (s *Server) RemovePost(ctx context.Context, postID int64) int {
	err := s.postRepo.Remove(ctx, postID)
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return http.StatusNotFound
		} else {
			log.Println(err)
			return http.StatusInternalServerError
		}
	}
	return http.StatusOK
}

func (s *Server) ParseCommentReq(req *http.Request) (*repository.Comment, int) {
	key := req.URL.Query().Get(PostIDKey)
	if key == "" {
		return nil, http.StatusBadRequest
	}
	postID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return nil, http.StatusBadRequest
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	var commentReq AddCommentRequest
	if err = json.Unmarshal(body, &commentReq); err != nil {
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	comment := &repository.Comment{
		Text:   commentReq.Text,
		PostID: postID,
	}
	return comment, http.StatusOK
}

func (s *Server) AddComment(ctx context.Context, comment *repository.Comment) ([]byte, int) {
	id, err := s.commentRepo.Add(ctx, comment)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	comment.ID = id
	commentJson, err := json.Marshal(comment)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return commentJson, http.StatusOK
}

func (s *Server) ParseCommentID(req *http.Request) (int64, int) {
	key := req.URL.Query().Get(CommentIDKey)
	if key == "" {
		return 0, http.StatusBadRequest
	}
	commentID, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return commentID, http.StatusOK
}

func (s *Server) RemoveComment(ctx context.Context, commentID int64) int {
	err := s.commentRepo.Remove(ctx, commentID)
	if err != nil {
		if errors.Is(err, repository.ErrZeroRows) {
			return http.StatusNotFound
		} else {
			return http.StatusInternalServerError
		}
	}
	return http.StatusOK
}
