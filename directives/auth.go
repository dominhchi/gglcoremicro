package directives

import (
	"context"
	"github.com/dominhchi/gglcoremicro/middlewares"
	"golang.org/x/exp/slices"

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

func IsAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsAdmin {

		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func IsSuper(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsSuperUser {

		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func IsStaff(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsStaff {

		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	check := slices.Contains(user.Permissions, role)
	if check {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func HasStaffRole(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	check := slices.Contains(user.StaffPermission, role)
	if check && user.IsStaff {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}
