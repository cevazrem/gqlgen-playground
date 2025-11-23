package content_dto

const (
	perPageLimit   = 1000
	perPageDefault = 100
)

// LimitOffsetPagination лимит-оффсетная пагинация
type LimitOffsetPagination struct {
	limit, offset int64
}

// BuildLimitOffsetPagination собирает лимит-оффсетную пагинацию по номеру страницы и кол-ву элементов на странице
func BuildLimitOffsetPagination(page, perPage int64) LimitOffsetPagination {
	if page <= 0 {
		page = 1
	}

	if perPage > perPageLimit {
		perPage = perPageLimit
	}

	if perPage <= 0 {
		perPage = perPageDefault
	}

	return LimitOffsetPagination{
		limit:  perPage,
		offset: perPage * (page - 1),
	}
}

// Limit возвращает значение limit
func (p LimitOffsetPagination) Limit() uint64 {
	if p.limit == 0 {
		return perPageDefault
	}

	return uint64(p.limit) // nolint:gosec
}

// Offset возвращает значение offset
func (p LimitOffsetPagination) Offset() uint64 {
	return uint64(p.offset) // nolint:gosec
}
