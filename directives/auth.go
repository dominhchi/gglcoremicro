package directives

import (
	"context"
	"github.com/dominhchi/gglcoremicro/middlewares"
	"github.com/dominhchi/gglcoremicro/utils"
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
	if user.IsAdmin || user.IsSuperUser {
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
	if user.IsStaff || user.IsSuperUser {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsAdmin || user.IsSuperUser {
		return next(ctx)
	}
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
	if user.IsSuperUser {
		return next(ctx)
	}
	check := slices.Contains(user.StaffPermission, role)
	if check && user.IsStaff {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func HasPermission(ctx context.Context, obj interface{}, next graphql.Resolver, permissions []string) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsAdmin || user.IsSuperUser {
		return next(ctx)
	}
	check := utils.Intersect(user.Permissions, permissions)
	if len(check) > 0 {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}

func HasStaffPermission(ctx context.Context, obj interface{}, next graphql.Resolver, permissions []string) (interface{}, error) {
	user := middlewares.CtxUser(ctx)
	if user.IsSuperUser {
		return next(ctx)
	}
	check := utils.Intersect(user.StaffPermission, permissions)
	if len(check) > 0 {
		return next(ctx)
	}
	return nil, &gqlerror.Error{
		Message: "Access Denied",
	}
}
