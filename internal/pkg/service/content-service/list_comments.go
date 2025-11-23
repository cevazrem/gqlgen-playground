package content_service

import (
	"context"
	"fmt"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// ListComments получить данные нескольких комментариев
func (s *Service) ListComments(ctx context.Context, filter content_dto.CommentFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.Comment, int64, error) {
	comments, total, err := s.storage.ListComments(ctx, filter.ToSql(), pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("s.storage.ListComments: %w", err)
	}

	return comments, total, nil
}
