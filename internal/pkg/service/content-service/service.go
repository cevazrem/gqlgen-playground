package content_service

import (
	"context"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"

	sq "github.com/Masterminds/squirrel"
)

type (
	Storage interface {
		CreateUser(ctx context.Context, data map[string]interface{}) (*content_model.User, error)
		GetUser(ctx context.Context, whereCond sq.Sqlizer) (*content_model.User, error)
		ListUsers(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.User, int64, error)

		CreatePost(ctx context.Context, data map[string]interface{}) (*content_model.Post, error)
		GetPost(ctx context.Context, whereCond sq.Sqlizer) (*content_model.Post, error)
		ListPosts(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.Post, int64, error)

		CreateComment(ctx context.Context, data map[string]interface{}) (*content_model.Comment, error)
		ListComments(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.Comment, int64, error)
	}

	// Service .
	Service struct {
		storage Storage
	}
)

// NewService инициализирует новый сервис
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
