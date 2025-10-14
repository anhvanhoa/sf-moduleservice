package entity

import (
	"time"
)

type ResourcePermission struct {
	tableName    struct{} `pg:"resource_permissions"`
	ID           string
	UserID       string
	ResourceType string
	ResourceID   string
	Action       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (r *ResourcePermission) TableName() any {
	return r.tableName
}

type ResourcePermissionFilter struct {
	UserID       string
	ResourceType string
	ResourceID   string
	Action       string
}
