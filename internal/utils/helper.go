package utils

import (
	"context"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(id string) (pgtype.UUID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return pgtype.UUID{
		Bytes: u,
		Valid: true,
	}, nil
}

func GetUserID(ctx context.Context) (pgtype.UUID, bool) {
	v, ok := ctx.Value(constants.ContextKeyUserID).(pgtype.UUID)
	return v, ok
}
