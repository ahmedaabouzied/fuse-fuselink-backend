package utils

import (
	"context"
	"fmt"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
)

func ExtractContextUserID(ctx context.Context) (string, error) {
	ctxUserID := ctx.Value(entities.UserIDContextKey)
	if ctxUserID == nil {
		return "", fmt.Errorf("%s", "empty context user ID")
	}
	return ctxUserID.(string), nil
}
