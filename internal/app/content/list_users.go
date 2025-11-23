package content

import (
	"context"
	"errors"
	"fmt"
	desc "gqlgen-playground/internal/pb/content/v1"
	content_dto "gqlgen-playground/internal/pkg/dto/content"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListUsers получить список пользователей
func (i *Implementation) ListUsers(ctx context.Context, req *desc.ListUsersRequest) (*desc.ListUsersResponse, error) {
	if err := validateListUsersRequest(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validateListUsersRequest: %v", err)
	}

	users, total, err := i.contentService.ListUsers(ctx, content_dto.BuildUserFilterByListReq(req.Filter), content_dto.BuildLimitOffsetPagination(req.Page, req.PerPage))
	if err != nil {
		return nil, fmt.Errorf("i.contentService.ListUsers: %w", err)
	}

	protoUsers := make([]*desc.User, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, user.ToProto())
	}

	return &desc.ListUsersResponse{
		Users: protoUsers,
		Total: total,
	}, nil
}

func validateListUsersRequest(req *desc.ListUsersRequest) error {
	if req == nil {
		return errors.New("nil request")
	}

	return nil
}
