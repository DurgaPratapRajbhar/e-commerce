package utils

// HTTP Status Messages
const (
	MsgSuccess         = "Success"
	MsgCreated         = "Created successfully"
	MsgUpdated         = "Updated successfully"
	MsgDeleted         = "Deleted successfully"
	MsgNotFound        = "Resource not found"
	MsgUnauthorized    = "Unauthorized access"
	MsgForbidden       = "Forbidden"
	MsgBadRequest      = "Invalid request"
	MsgInternalError   = "Internal server error"
	MsgValidationError = "Validation failed"
)

// User Roles
const (
	RoleUser   = "user"
	RoleAdmin  = "admin"
	RoleVendor = "vendor"
)

// Order Status
const (
	OrderStatusPending    = "pending"
	OrderStatusConfirmed  = "confirmed"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"
	OrderStatusReturned   = "returned"
)

// Payment Status
const (
	PaymentStatusPending  = "pending"
	PaymentStatusSuccess  = "success"
	PaymentStatusFailed   = "failed"
	PaymentStatusRefunded = "refunded"
)

// Payment Methods
const (
	PaymentMethodCOD     = "cod"
	PaymentMethodCard    = "card"
	PaymentMethodUPI     = "upi"
	PaymentMethodWallet  = "wallet"
	PaymentMethodNetBanking = "netbanking"
)

// Shipment Status
const (
	ShipmentStatusPending     = "pending"
	ShipmentStatusPicked      = "picked"
	ShipmentStatusInTransit   = "in_transit"
	ShipmentStatusOutForDelivery = "out_for_delivery"
	ShipmentStatusDelivered   = "delivered"
	ShipmentStatusFailed      = "failed"
)

// Product Status
const (
	ProductStatusActive   = "active"
	ProductStatusInactive = "inactive"
	ProductStatusDraft    = "draft"
)

// File Size Limits (in MB)
const (
	MaxImageSizeMB     = 5
	MaxDocumentSizeMB  = 10
	MaxVideoSizeMB     = 50
)

// Pagination
const (
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// Cache Prefixes
const (
	CachePrefixProduct  = "product"
	CachePrefixCategory = "category"
	CachePrefixUser     = "user"
	CachePrefixCart     = "cart"
	CachePrefixOrder    = "order"
)

// Date Formats
const (
	DateFormatYYYYMMDD = "2006-01-02"
	DateFormatDDMMYYYY = "02-01-2006"
	DateTimeFormat     = "2006-01-02 15:04:05"
)

// GST Rates
const (
	GST5  = 5.0
	GST12 = 12.0
	GST18 = 18.0
	GST28 = 28.0
)

// Transaction Types
const (
	TransactionTypeIn  = "in"
	TransactionTypeOut = "out"
)