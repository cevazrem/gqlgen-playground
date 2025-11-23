package content

import (
	"context"
	contentv1 "gqlgen-playground/internal/pb/content/v1"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

type (
	ContentService interface {
		CreateUser(ctx context.Context, user content_model.User) (*content_model.User, error)
		GetUser(ctx context.Context, filter content_dto.UserFilter) (*content_model.User, error)
		ListUsers(ctx context.Context, filter content_dto.UserFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.User, int64, error)

		CreatePost(ctx context.Context, post content_model.Post) (*content_model.Post, error)
		GetPost(ctx context.Context, filter content_dto.PostFilter) (*content_model.Post, error)
		ListPosts(ctx context.Context, filter content_dto.PostFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.Post, int64, error)

		CreateComment(ctx context.Context, post content_model.Comment) (*content_model.Comment, error)
		ListComments(ctx context.Context, filter content_dto.CommentFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.Comment, int64, error)
	}

	Implementation struct {
		contentv1.UnimplementedContentServiceServer

		contentService ContentService
	}
)

func NewImplementation(
	contentService ContentService,
) *Implementation {
	return &Implementation{
		contentService: contentService,
	}
}
