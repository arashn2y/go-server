package constants

type Resource string
type Role string
type Permission string
type contextKey string

const (
	ContextKeyUserID contextKey = "user_id"

	ResourceUsers    Resource = "users"
	ResourceProducts Resource = "products"

	RoleSuperAdmin Role = "super_admin"
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
	RoleViewer     Role = "viewer"

	PermissionAll    Permission = "all"
	PermissionRead   Permission = "read"
	PermissionCreate Permission = "create"
	PermissionUpdate Permission = "update"
	PermissionDelete Permission = "delete"
)

var Actions = []Permission{PermissionAll, PermissionRead, PermissionCreate, PermissionUpdate, PermissionDelete}
var Roles = []Role{RoleSuperAdmin, RoleAdmin, RoleUser, RoleViewer}
var Resources = []Resource{ResourceUsers, ResourceProducts}
