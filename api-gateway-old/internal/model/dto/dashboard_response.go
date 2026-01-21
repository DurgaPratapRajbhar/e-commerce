package dto

// DashboardResponse represents aggregated dashboard data
type DashboardResponse struct {
	TotalUsers      int64       `json:"total_users"`
	TotalOrders     int64       `json:"total_orders"`
	TotalRevenue    float64     `json:"total_revenue"`
	PendingOrders   int64       `json:"pending_orders"`
	RecentOrders    interface{} `json:"recent_orders"`
	MonthlyStats    interface{} `json:"monthly_stats"`
}