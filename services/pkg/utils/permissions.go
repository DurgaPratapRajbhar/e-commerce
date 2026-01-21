package utils

type Permission string

type Role string

const (
	RoleMerchant Role       = "merchant"
	PermReadUsers Permission = "read:users"
	PermWriteUsers Permission = "write:users"
	PermDeleteUsers Permission = "delete:users"
	PermReadProducts Permission = "read:products"
	PermWriteProducts Permission = "write:products"
	PermDeleteProducts Permission = "delete:products"
	PermReadOrders Permission = "read:orders"
	PermWriteOrders Permission = "write:orders"
	PermDeleteOrders Permission = "delete:orders"
	PermReadCategories Permission = "read:categories"
	PermWriteCategories Permission = "write:categories"
	PermDeleteCategories Permission = "delete:categories"
)

var RolePermissions = map[Role][]Permission{
	Role("admin"): {
		PermReadUsers, PermWriteUsers, PermDeleteUsers,
		PermReadProducts, PermWriteProducts, PermDeleteProducts,
		PermReadOrders, PermWriteOrders, PermDeleteOrders,
		PermReadCategories, PermWriteCategories, PermDeleteCategories,
	},
	RoleMerchant: {
		PermReadProducts, PermWriteProducts, PermDeleteProducts,
		PermReadOrders, PermWriteOrders,
		PermReadCategories, PermWriteCategories,
	},
	Role("user"): {
		PermReadProducts,
		PermReadCategories,
		PermWriteOrders,
	},
}

// HasPermission checks if a role has a specific permission
func HasPermission(role Role, permission Permission) bool {
	permissions, exists := RolePermissions[role]
	if !exists {
		return false
	}

	for _, perm := range permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

// GetRolePermissions returns all permissions for a given role
func GetRolePermissions(role Role) []Permission {
	permissions, exists := RolePermissions[role]
	if !exists {
		return []Permission{}
	}
	return permissions
}