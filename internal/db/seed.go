package db

import (
	"context"
	"fmt"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes/auth"
	"github.com/sirupsen/logrus"
)

func Seed(ctx context.Context, db *repository.Queries) error {
	if err := seedRoles(ctx, db); err != nil {
		return fmt.Errorf("failed to seed roles: %w", err)
	}
	if err := seedPermissions(ctx, db); err != nil {
		return fmt.Errorf("failed to seed permissions: %w", err)
	}
	if err := seedSuperAdmin(ctx, db); err != nil {
		return fmt.Errorf("failed to seed super admin user: %w", err)
	}
	if err := seedAdminRolePermissions(ctx, db); err != nil {
		return fmt.Errorf("failed to seed admin role permissions: %w", err)
	}
	return nil
}

func seedRoles(ctx context.Context, db *repository.Queries) error {
	logrus.Info("seeding roles...")

	for _, roleName := range constants.Roles {
		role, err := db.GetRoleByName(ctx, string(roleName))
		if err == nil && role.ID != 0 {
			logrus.Infof("role %s already exists, skipping", roleName)
			continue
		}
		roleErr := db.CreateRole(ctx, string(roleName))
		if roleErr != nil {
			return fmt.Errorf("failed to insert role %s: %w", roleName, roleErr)
		}
	}
	logrus.Info("roles seeded successfully")
	return nil
}

func seedPermissions(ctx context.Context, db *repository.Queries) error {
	logrus.Info("seeding permissions...")
	for _, action := range constants.Actions {
		permission, err := db.GetPermissionByName(ctx, string(action))
		if err == nil && permission.ID != 0 {
			logrus.Infof("permission %s already exists, skipping", action)
			continue
		}
		permissionErr := db.CreatePermission(ctx, string(action))
		if permissionErr != nil {
			return fmt.Errorf("failed to insert permission %s: %w", action, permissionErr)
		}
	}
	logrus.Info("permissions seeded successfully")
	return nil
}

func seedSuperAdmin(ctx context.Context, db *repository.Queries) error {
	logrus.Info("seeding super admin user...")
	user := repository.CreateUserParams{
		Name:     "Arash",
		Email:    "arashn2y@gmail.com",
		Password: "Arash2026!", // Needs to pass the hashed password instead of plain text
		IsActive: true,
	}
	hash, err := auth.Hash(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash super admin password: %w", err)
	}

	role, err := db.GetRoleByName(ctx, string(constants.RoleSuperAdmin))
	if err != nil {
		return fmt.Errorf("failed to get super admin role: %w", err)
	}

	superAdmin, err := db.GetUserByEmail(ctx, user.Email)
	if err == nil && superAdmin.Email != "" {
		logrus.Infof("super admin user with email %s already exists, skipping", user.Email)
		return nil
	}

	id, err := db.CreateUser(ctx, repository.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
		IsActive: user.IsActive,
		RoleID:   role.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to insert super admin user: %w", err)
	}

	// assign super_admin role
	err = db.AssignRoleToUser(ctx, repository.AssignRoleToUserParams{
		UserID: id,
		RoleID: role.ID,
	})

	// assign all permissions to super_admin role
	permissions, err := db.GetAllPermissions(ctx)
	if err != nil {
		return fmt.Errorf("failed to get permissions: %w", err)
	}

	for _, perm := range permissions {
		for _, resource := range constants.Resources {
			err = db.AssignPermissionToRole(ctx, repository.AssignPermissionToRoleParams{
				RoleID:       role.ID,
				Resource:     string(resource),
				PermissionID: perm.ID,
			})
			if err != nil {
				return fmt.Errorf("failed to assign permission %s on resource %s to super admin role: %w", perm.Name, resource, err)
			}
		}
	}

	logrus.Info("super admin user seeded successfully")
	return err
}

func seedAdminRolePermissions(ctx context.Context, db *repository.Queries) error {
	// An Admin can read and write all resources but cannot do all and delete
	logrus.Info("seeding admin role permissions...")
	role, err := db.GetRoleByName(ctx, string(constants.RoleAdmin))
	if err != nil {
		return fmt.Errorf("failed to get admin role: %w", err)
	}

	permissions, err := db.GetAllPermissions(ctx)
	if err != nil {
		return fmt.Errorf("failed to get permissions: %w", err)
	}

	for _, perm := range permissions {
		if perm.Name == string(constants.PermissionDelete) || perm.Name == string(constants.PermissionAll) {
			continue // skip delete permission for admin role
		}
		for _, resource := range constants.Resources {
			permission, err := db.GetPermissionsByRoleID(ctx, repository.GetPermissionsByRoleIDParams{
				RoleID:   role.ID,
				Resource: string(resource),
			})
			if err == nil && len(permission) > 0 {
				logrus.Infof("permission %s on resource %s already assigned to admin role, skipping", perm.Name, resource)
				continue
			}
			err = db.AssignPermissionToRole(ctx, repository.AssignPermissionToRoleParams{
				RoleID:       role.ID,
				Resource:     string(resource),
				PermissionID: perm.ID,
			})
			if err != nil {
				return fmt.Errorf("failed to assign permission %s on resource %s to admin role: %w", perm.Name, resource, err)
			}
		}
	}

	logrus.Info("admin role permissions seeded successfully")
	return nil
}
