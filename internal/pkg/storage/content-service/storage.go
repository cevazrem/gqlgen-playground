package content_storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage cтруктура для работы с хранилищем
type Storage struct {
	pool *pgxpool.Pool
	sb   sq.StatementBuilderType
}

// NewStorage инициализирует пул коннектов к Postgres
func NewStorage(_ context.Context, pool *pgxpool.Pool) (*Storage, error) {
	return &Storage{
		pool: pool,
		sb:   sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

// builder возвращает sq builder для сборки SQL выражения
func (s *Storage) builder() sq.StatementBuilderType {
	return s.sb
}

// getx выполняет запрос (Sqlizer) и мапит результат в dest (структура/указатель на структуру)
func (s *Storage) getx(ctx context.Context, dest any, qb sq.Sqlizer) error {
	sqlStr, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	if err = pgxscan.Get(ctx, s.pool, dest, sqlStr, args...); err != nil {
		return fmt.Errorf("pgxscan.Get: %w", err)
	}

	return nil
}

// selectx выполняет запрос (Sqlizer) и мапит результат в массив
func (s *Storage) selectx(ctx context.Context, dest any, qb sq.Sqlizer) error {
	sqlStr, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	if err = pgxscan.Select(ctx, s.pool, dest, sqlStr, args...); err != nil {
		return fmt.Errorf("pgxscan.Select: %w", err)
	}

	return nil
}

// execx выполняет запрос (Sqlizer)
func (s *Storage) execx(ctx context.Context, qb sq.Sqlizer) (int64, error) {
	sqlStr, args, err := qb.ToSql()
	if err != nil {
		return 0, fmt.Errorf("qb.ToSql: %w", err)
	}

	cmd, err := s.pool.Exec(ctx, sqlStr, args...)
	if err != nil {
		return 0, handlePgError(err)
	}

	return cmd.RowsAffected(), nil
}
