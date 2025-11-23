package content_dto

import (
	desc "gqlgen-playground/internal/pb/content/v1"

	sq "github.com/Masterminds/squirrel"
	"gopkg.in/guregu/null.v4/zero"
)

type UserFilter struct {
	IdsIn         []string
	UsernamesIn   []string
	CreatedAtFrom zero.Time
	CreatedAtTo   zero.Time
}

func BuildUserFilterByGetReq(req *desc.GetUserRequest) UserFilter {
	filter := UserFilter{}

	if x, ok := req.Filter.(*desc.GetUserRequest_Id); ok {
		filter.IdsIn = append(filter.IdsIn, x.Id)
	}

	if x, ok := req.Filter.(*desc.GetUserRequest_Username); ok {
		filter.UsernamesIn = append(filter.UsernamesIn, x.Username)
	}

	return filter
}

func BuildUserFilterByListReq(req *desc.ListUsersRequest_Filter) UserFilter {
	if req == nil {
		return UserFilter{}
	}

	filter := UserFilter{
		IdsIn:       req.IdsIn,
		UsernamesIn: req.UsernamesIn,
	}

	if req.CreatedAtFrom != nil {
		filter.CreatedAtFrom = zero.TimeFrom(req.CreatedAtFrom.AsTime())
	}

	if req.CreatedAtTo != nil {
		filter.CreatedAtTo = zero.TimeFrom(req.CreatedAtTo.AsTime())
	}

	return filter
}

// ToSql преобразует в SQL фильтр
func (u UserFilter) ToSql() sq.Sqlizer {
	var filterElem []sq.Sqlizer

	if len(u.IdsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"id": u.IdsIn})
	}

	if len(u.UsernamesIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"username": u.UsernamesIn})
	}

	if u.CreatedAtFrom.Valid {
		filterElem = append(filterElem, sq.GtOrEq{"created_at": u.CreatedAtFrom.Time})
	}

	if u.CreatedAtTo.Valid {
		filterElem = append(filterElem, sq.LtOrEq{"created_at": u.CreatedAtTo.Time})
	}

	filterElem = append(filterElem, sq.Eq{"deleted_at": nil})

	return sq.And(filterElem)
}

type PostFilter struct {
	IdsIn         []string
	AuthorIDsIn   []string
	CreatedAtFrom zero.Time
	CreatedAtTo   zero.Time
}

func BuildPostFilterByListReq(req *desc.ListPostsRequest_Filter) PostFilter {
	if req == nil {
		return PostFilter{}
	}

	filter := PostFilter{
		IdsIn:       req.IdsIn,
		AuthorIDsIn: req.AuthorIdsIn,
	}

	if req.CreatedAtFrom != nil {
		filter.CreatedAtFrom = zero.TimeFrom(req.CreatedAtFrom.AsTime())
	}

	if req.CreatedAtTo != nil {
		filter.CreatedAtTo = zero.TimeFrom(req.CreatedAtTo.AsTime())
	}

	return filter
}

// ToSql преобразует в SQL фильтр
func (u PostFilter) ToSql() sq.Sqlizer {
	var filterElem []sq.Sqlizer

	if len(u.IdsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"id": u.IdsIn})
	}

	if len(u.AuthorIDsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"author_id": u.AuthorIDsIn})
	}

	if u.CreatedAtFrom.Valid {
		filterElem = append(filterElem, sq.GtOrEq{"created_at": u.CreatedAtFrom.Time})
	}

	if u.CreatedAtTo.Valid {
		filterElem = append(filterElem, sq.LtOrEq{"created_at": u.CreatedAtTo.Time})
	}

	filterElem = append(filterElem, sq.Eq{"deleted_at": nil})

	return sq.And(filterElem)
}

type CommentFilter struct {
	IdsIn         []string
	PostIDsIn     []string
	AuthorIDsIn   []string
	CreatedAtFrom zero.Time
	CreatedAtTo   zero.Time
}

func BuildCommentFilterByListReq(req *desc.ListCommentsRequest_Filter) CommentFilter {
	if req == nil {
		return CommentFilter{}
	}

	filter := CommentFilter{
		IdsIn:       req.IdsIn,
		PostIDsIn:   req.PostIdsIn,
		AuthorIDsIn: req.AuthorIdsIn,
	}

	if req.CreatedAtFrom != nil {
		filter.CreatedAtFrom = zero.TimeFrom(req.CreatedAtFrom.AsTime())
	}

	if req.CreatedAtTo != nil {
		filter.CreatedAtTo = zero.TimeFrom(req.CreatedAtTo.AsTime())
	}

	return filter
}

// ToSql преобразует в SQL фильтр
func (u CommentFilter) ToSql() sq.Sqlizer {
	var filterElem []sq.Sqlizer

	if len(u.IdsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"id": u.IdsIn})
	}

	if len(u.PostIDsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"post_id": u.PostIDsIn})
	}

	if len(u.AuthorIDsIn) > 0 {
		filterElem = append(filterElem, sq.Eq{"author_id": u.AuthorIDsIn})
	}

	if u.CreatedAtFrom.Valid {
		filterElem = append(filterElem, sq.GtOrEq{"created_at": u.CreatedAtFrom.Time})
	}

	if u.CreatedAtTo.Valid {
		filterElem = append(filterElem, sq.LtOrEq{"created_at": u.CreatedAtTo.Time})
	}

	filterElem = append(filterElem, sq.Eq{"deleted_at": nil})

	return sq.And(filterElem)
}
