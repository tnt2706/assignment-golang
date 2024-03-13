package directive

import (
	"assignment/internal/model"
	"assignment/pkg/util"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/samber/lo"
)

func Auth(ctx context.Context, next graphql.Resolver, roles []*model.Role) (res interface{}, err error) {
	signature := ctx.Value("signature").(map[string]interface{})
	var userRoles []string = signature["roles"].([]string)
	var rolePermissions []string

	for _, role := range roles {
		rolePermissions = append(rolePermissions, role.String())
	}

	for _, role := range userRoles {
		roleMap := util.ToGraphqlRole(role)
		matchRole := lo.Contains(rolePermissions, roleMap)
		if matchRole {
			return next(ctx)
		}
	}

	return nil, errors.New("you are not authorized for this field")
}
