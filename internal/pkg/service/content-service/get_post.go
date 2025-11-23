package content_service

import (
	"context"
	"fmt"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// GetPost получить данные поста
func (s *Service) GetPost(ctx context.Context, filter content_dto.PostFilter) (*content_model.Post, error) {
	post, err := s.storage.GetPost(ctx, filter.ToSql())
	if err != nil {
		return nil, fmt.Errorf("s.storage.GetPost: %w", err)
	}

	return post, nil
}
