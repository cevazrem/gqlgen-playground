package content_service

import (
	"context"
	"fmt"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// ListUsers получить данные нескольких пользователей
func (s *Service) ListUsers(ctx context.Context, filter content_dto.UserFilter, pagination content_dto.LimitOffsetPagination) ([]*content_model.User, int64, error) {
	users, total, err := s.storage.ListUsers(ctx, filter.ToSql(), pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("s.storage.ListUsers: %w", err)
	}

	return users, total, nil
}
