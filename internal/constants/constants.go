package constants

type Resource string
type Role string
type Permission string

const (
	ResourceUsers    Resource = "users"
	ResourceProducts Resource = "products"

	RoleSuperAdmin Role = "super_admin"
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
	RoleViewer     Role = "viewer"

	PermissionAll    Permission = "all"
	PermissionRead   Permission = "read"
	PermissionWrite  Permission = "write"
	PermissionDelete Permission = "delete"
)

var Actions = []Permission{PermissionAll, PermissionRead, PermissionWrite, PermissionDelete}
var Roles = []Role{RoleSuperAdmin, RoleAdmin, RoleUser, RoleViewer}
var Resources = []Resource{ResourceUsers, ResourceProducts}
