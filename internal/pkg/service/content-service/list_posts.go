package content_service

import (
	"context"
	"fmt"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// ListPosts получить данные нескольких постов
func (s *Service) ListPosts(ctx context.Context, filter content_dto.PostFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.Post, int64, error) {
	users, total, err := s.storage.ListPosts(ctx, filter.ToSql(), pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("s.storage.ListPosts: %w", err)
	}

	return users, total, nil
}
