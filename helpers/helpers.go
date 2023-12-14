// helpers/helpers.go
package helpers

import (
	"context"
	"errors"
)

// GetUserIDFromContext retrieves the user ID from the request context
func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		return 0, ErrUserIDNotFound
	}
	return userID, nil
}

// ErrUserIDNotFound is an error indicating that the user ID is not found in the context
var ErrUserIDNotFound = errors.New("user ID not found in context")
