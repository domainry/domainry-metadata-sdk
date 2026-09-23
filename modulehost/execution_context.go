package modulehost

import (
	"context"

	shareddefinition "github.com/domainry/domainry-foundation/definition"
	"github.com/domainry/domainry-orm/sqlhost"
)

type DBTX = sqlhost.DBTX

// WithExecutor binds a host-owned transaction to a Metadata operation. Module
// persistence resolves it before falling back to the borrowed host database.
func WithExecutor(ctx context.Context, executor DBTX) context.Context {
	return shareddefinition.WithExecutor(ctx, executor)
}

func ExecutorFromContext(ctx context.Context, fallback DBTX) DBTX {
	return shareddefinition.ExecutorFromContext(ctx, fallback)
}
