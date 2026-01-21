package utils

// Pagination represents pagination parameters
type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// PaginationResponse represents a paginated response
type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// NewPagination creates a new pagination struct
func NewPagination(page, limit, total int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	
	pages := total / limit
	if total%limit != 0 {
		pages++
	}
	
	if page > pages {
		page = pages
	}
	
	return &Pagination{
		Page:  page,
		Limit: limit,
		Total: total,
	}
}

// GetOffset calculates the offset for database queries
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// GetLimit returns the limit for database queries
func (p *Pagination) GetLimit() int {
	return p.Limit
}