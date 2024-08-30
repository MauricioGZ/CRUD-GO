package service

import (
	"context"
	"errors"
)

var (
	ErrRolesNotInitialized = errors.New("roles not initialized")
)

// emulating a cache
var rolesPermissions map[string][]string = make(map[string][]string)
var roleIDs map[string]int64 = make(map[string]int64)

func (s *serv) GetAllPermissionsRoles(ctx context.Context) error {
	pprr, err := s.repo.GetAllPermissionsRoles(ctx)
	if err != nil {
		return err
	}

	for _, pr := range pprr {
		rolesPermissions[pr.Role] = append(rolesPermissions[pr.Role], pr.Permission)
	}
	return nil
}

func (s *serv) GetAllRoles(ctx context.Context) error {
	rr, err := s.repo.GetAllRoles(ctx)

	if err != nil {
		return err
	}

	for _, r := range rr {
		roleIDs[r.Role] = r.ID
	}

	return nil
}

func getPermissions(role string) []string {
	if len(rolesPermissions) == 0 {
		return nil
	}
	return rolesPermissions[role]
}

func mayCreate(role string) bool {
	for _, p := range getPermissions(role) {
		if p == "Create" {
			return true
		}
	}
	return false
}

func mayUpdate(role string) bool {
	for _, p := range getPermissions(role) {
		if p == "Update" {
			return true
		}
	}
	return false
}

func mayDelete(role string) bool {
	for _, p := range getPermissions(role) {
		if p == "Delete" {
			return true
		}
	}
	return false
}

func getRoleID(role string) int64 {
	if len(roleIDs) == 0 {
		return 0
	}
	return roleIDs[role]
}

func getRole(id int64) string {
	for k, v := range roleIDs {
		if v == id {
			return k
		}
	}
	return ""
}
