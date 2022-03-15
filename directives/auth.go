package directives

import (
	"context"
	"github.com/dominhchi/gglcoremicro/middlewares"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user == nil || user.ID == "" {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
