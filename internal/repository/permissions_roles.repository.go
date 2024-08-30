package repository

import (
	"context"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryGetAllPermissionsRoles = ` select 
																	ROLES.role, 
																	PERMISSIONS.permission 
																from ROLES 
																join PERMISSIONS_ROLES
																	on ROLES.id = PERMISSIONS_ROLES.roleId
																join PERMISSIONS
																	on PERMISSIONS_ROLES.permissionId = PERMISSIONS.id;`
	qryGetAllRoles = `select
											id,
											role
										from ROLES;`
)

func (r *repo) GetAllPermissionsRoles(ctx context.Context) ([]entity.PermissionRoles, error) {
	var permissionsRoles []entity.PermissionRoles
	var permissionRole entity.PermissionRoles

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllPermissionsRoles,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&permissionRole.Role,
			&permissionRole.Permission,
		)

		if err != nil {
			return nil, err
		}

		permissionsRoles = append(permissionsRoles, permissionRole)
	}

	return permissionsRoles, nil
}

func (r *repo) GetAllRoles(ctx context.Context) ([]entity.Role, error) {
	var roles []entity.Role
	var role entity.Role

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllRoles,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(
			&role.ID,
			&role.Role,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}
