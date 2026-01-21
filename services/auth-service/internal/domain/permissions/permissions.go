package permissions

type Permission string

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleUser     Role = "user"
	RoleMerchant Role = "merchant"
)

const (
	PermReadUsers        Permission = "read:users"
	PermWriteUsers       Permission = "write:users"
	PermDeleteUsers      Permission = "delete:users"
	PermReadProducts     Permission = "read:products"
	PermWriteProducts    Permission = "write:products"
	PermDeleteProducts   Permission = "delete:products"
	PermReadOrders       Permission = "read:orders"
	PermWriteOrders      Permission = "write:orders"
	PermDeleteOrders     Permission = "delete:orders"
	PermReadCategories   Permission = "read:categories"
	PermWriteCategories  Permission = "write:categories"
	PermDeleteCategories Permission = "delete:categories"
)

var RolePermissions = map[Role][]Permission{
	RoleAdmin: {
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
	RoleUser: {
		PermReadProducts,
		PermReadCategories,
		PermWriteOrders,
	},
}

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

func GetRolePermissions(role Role) []Permission {
	permissions, exists := RolePermissions[role]
	if !exists {
		return []Permission{}
	}
	return permissions
}