package dto

// DashboardData represents comprehensive dashboard data with related information
type DashboardData struct {
	OrdersWithUsers    []OrderWithUser `json:"orders_with_users"`
	CartsWithProducts  []CartWithProducts `json:"carts_with_products"`
	RecentUsers        []interface{} `json:"recent_users"`
	PopularProducts    []interface{} `json:"popular_products"`
	TotalUsers         int64         `json:"total_users"`
	TotalOrders        int64         `json:"total_orders"`
	TotalRevenue       float64       `json:"total_revenue"`
}

// OrderSummary represents a simplified order for dashboard display
type OrderSummary struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	Status       string  `json:"status"`
	TotalAmount  float64 `json:"total_amount"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

// UserSummary represents a simplified user for dashboard display
type UserSummary struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

// UserProfileWithOrders represents a user profile with their orders
type UserProfileWithOrders struct {
	User     interface{}     `json:"user"`
	Orders   []interface{}   `json:"orders"`
	Products []interface{}   `json:"products"`
}

// UserWithOrderCount represents a user with their order count
type UserWithOrderCount struct {
	User       UserSummary `json:"user"`
	OrderCount int         `json:"order_count"`
}