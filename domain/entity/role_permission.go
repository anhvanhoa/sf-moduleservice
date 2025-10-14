package entity

import (
	"time"
)

// RolePermission represents a role-permission relationship entity
type RolePermission struct {
	tableName    struct{} `pg:"role_permissions"`
	RoleID       string
	PermissionID string
	CreatedAt    time.Time
}

func (r *RolePermission) TableName() any {
	return r.tableName
}

type RolePermissionFilter struct {
	RoleID       string
	PermissionID string
}
