package content_storage

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	// ErrNotFound ошибка при отсутствии сущности
	ErrNotFound = errors.New("entity not found")
)

// handlePgError кастомная обработка ошибок postgre
func handlePgError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) || strings.Contains(err.Error(), pgx.ErrNoRows.Error()) {
		return ErrNotFound
	}

	return err
}
