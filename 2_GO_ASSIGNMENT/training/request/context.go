package request

import "context"

// ContextKey prevents collisions with other context values.
type ContextKey string

const (
	RequestIDKey ContextKey = "request_id"
	UserIDKey    ContextKey = "user_id"
	CompanyIDKey ContextKey = "company_id"
	RoleKey      ContextKey = "role"
)

// Identity contains authenticated user information.
type Identity struct {
	UserID    int64
	CompanyID int64
	Role      string
}

// Store Request ID in context.
func WithRequestID(
	ctx context.Context,
	requestID string,
) context.Context {
	return context.WithValue(
		ctx,
		RequestIDKey,
		requestID,
	)
}

// Store authenticated user information.
func WithIdentity(
	ctx context.Context,
	identity Identity,
) context.Context {

	ctx = context.WithValue(
		ctx,
		UserIDKey,
		identity.UserID,
	)

	ctx = context.WithValue(
		ctx,
		CompanyIDKey,
		identity.CompanyID,
	)

	ctx = context.WithValue(
		ctx,
		RoleKey,
		identity.Role,
	)

	return ctx
}

// Get Request ID.
func GetRequestID(ctx context.Context) string {
	value, _ := ctx.Value(RequestIDKey).(string)
	return value
}

// Get User ID.
func GetUserID(ctx context.Context) int64 {
	value, _ := ctx.Value(UserIDKey).(int64)
	return value
}

// Get Company ID.
func GetCompanyID(ctx context.Context) int64 {
	value, _ := ctx.Value(CompanyIDKey).(int64)
	return value
}

// Get Role.
func GetRole(ctx context.Context) string {
	value, _ := ctx.Value(RoleKey).(string)
	return value
}
