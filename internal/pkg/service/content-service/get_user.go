package content_service

import (
	"context"
	"fmt"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"
)

// GetUser получить данные пользователя
func (s *Service) GetUser(ctx context.Context, filter content_dto.UserFilter) (*content_model.User, error) {
	user, err := s.storage.GetUser(ctx, filter.ToSql())
	if err != nil {
		return nil, fmt.Errorf("s.storage.GetUser: %w", err)
	}

	return user, nil
}
