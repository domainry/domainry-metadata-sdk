package modulehost

import (
	"context"

	"github.com/domainry/domainry-orm/sqlhost"
)

type DBTX = sqlhost.DBTX

type executorContextKey struct{}

// WithExecutor binds a host-owned transaction to a Metadata operation. Module
// persistence resolves it before falling back to the borrowed host database.
func WithExecutor(ctx context.Context, executor DBTX) context.Context {
	if executor == nil {
		return ctx
	}
	return context.WithValue(ctx, executorContextKey{}, executor)
}

func ExecutorFromContext(ctx context.Context, fallback DBTX) DBTX {
	if ctx != nil {
		if executor, ok := ctx.Value(executorContextKey{}).(DBTX); ok && executor != nil {
			return executor
		}
	}
	return fallback
}
