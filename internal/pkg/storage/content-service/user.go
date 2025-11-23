package content_storage

import (
	"context"
	content_dto "gqlgen-playground/internal/pkg/dto/content"
	content_model "gqlgen-playground/internal/pkg/model/content"

	sq "github.com/Masterminds/squirrel"
)

// CreateUser создать пользователя в БД
func (s *Storage) CreateUser(ctx context.Context, data map[string]interface{}) (*content_model.User, error) {
	query := s.builder().
		Insert(UsersTable).
		SetMap(data).
		Suffix(returnAllUsersTabledFields)

	user := new(content_model.User)
	if pgErr := s.getx(ctx, user, query); pgErr != nil {
		return nil, handlePgError(pgErr)
	}

	return user, nil
}

// GetUser получить пользователя из БД
func (s *Storage) GetUser(ctx context.Context, whereCond sq.Sqlizer) (*content_model.User, error) {
	qb := s.builder().
		Select(usersTabledFields).
		From(UsersTable).
		Where(whereCond)

	user := new(content_model.User)
	if err := s.getx(ctx, user, qb); err != nil {
		return nil, handlePgError(err)
	}

	return user, nil
}

// ListUsers получить список пользователей из БД
func (s *Storage) ListUsers(ctx context.Context, whereCond sq.Sqlizer, pagination content_dto.LimitOffsetPagination) ([]*content_model.User, int64, error) {
	qb := s.builder().
		Select(usersTabledFields).
		From(UsersTable).
		Where(whereCond).
		OrderBy("created_at DESC").
		Limit(pagination.Limit()).
		Offset(pagination.Offset())

	var users []*content_model.User
	if err := s.selectx(ctx, &users, qb); err != nil {
		return nil, 0, handlePgError(err)
	}

	cntb := s.builder().
		Select("count (*) as total").
		From(UsersTable).
		Where(whereCond)

	var total int64
	if err := s.getx(ctx, &total, cntb); err != nil {
		return nil, 0, handlePgError(err)
	}

	return users, total, nil
}

/*
// DeleteUser удалить пользователя из БД
func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	qb := s.builder().
		Delete(UsersTable).
		Where(sq.Eq{"id": id})

	affected, err := s.execx(ctx, qb)
	if err != nil {
		return handlePgError(err)
	}

	if affected == 0 {
		return ErrNotFound
	}

	return nil
}
*/
