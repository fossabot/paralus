package dao

import (
	"context"

	"github.com/RafaySystems/rcloud-base/internal/models"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetRolePermissions(ctx context.Context, db bun.IDB, id uuid.UUID) ([]models.ResourcePermission, error) {
	// Could possibly union them later for some speedup
	var r = []models.ResourcePermission{}
	err := db.NewSelect().Table("authsrv_resourcepermission").
		ColumnExpr("authsrv_resourcepermission.name as name").
		Join(`JOIN authsrv_resourcerolepermission ON authsrv_resourcerolepermission.resource_permission_id=authsrv_resourcepermission.id`).
		Where("authsrv_resourcerolepermission.resource_role_id = ?", id).
		Where("authsrv_resourcepermission.trash = ?", false).
		Where("authsrv_resourcerolepermission.trash = ?", false).
		Scan(ctx, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}