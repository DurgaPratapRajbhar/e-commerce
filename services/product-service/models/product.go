package models

import (
	"html"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Product model
type Product struct {
	ID            uint64             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string             `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=255"`
	Slug          string             `gorm:"type:varchar(255);unique;not null" json:"slug" validate:"required,regexp=^[a-z0-9]+(?:-[a-z0-9]+)*$"`
	Description   string             `gorm:"type:text" json:"description" validate:"required,min=5"`
	Price         float64            `gorm:"type:decimal(10,2);not null" json:"price" validate:"required,gt=0"`
	Discount      float64            `gorm:"type:decimal(5,2);default:0" json:"discount" validate:"gte=0,lte=100"`
	Stock         int                `gorm:"type:int;not null" json:"stock" validate:"gte=0"`
	SKU           string             `gorm:"type:varchar(100);unique;not null" json:"sku" validate:"required"`
	Status        string             `gorm:"type:varchar(20);not null;default:'active'" json:"status" validate:"oneof=active inactive draft"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	CategoryID    uint64             `gorm:"index" json:"category_id,omitempty"`
	Brand         string             `gorm:"type:varchar(100)" json:"brand,omitempty"`
	UoMID         *uint64            `gorm:"index" json:"uom_id,omitempty"`
	QuantityValue *float64           `gorm:"type:decimal(10,2)" json:"quantity_value,omitempty" validate:"omitempty,gte=0"`
	PrimaryImage  *string            `json:"primary_image"`
	Images        []ProductImage     `gorm:"foreignKey:ProductID;references:ID" json:"images"`
	Category      Category           `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	Unit          UnitOfMeasurement  `gorm:"foreignKey:UoMID;references:ID" json:"unit,omitempty"`
	Attributes    []ProductAttribute `gorm:"foreignKey:ProductID;references:ID" json:"attributes,omitempty"`
	Variants      []ProductVariant   `gorm:"foreignKey:ProductID;references:ID" json:"variants,omitempty"`
}

type ProductAttribute struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint64 `gorm:"index;not null" json:"product_id"`
	Key       string `gorm:"type:varchar(100);not null" json:"key" validate:"required"`
	Value     string `gorm:"type:varchar(255);not null" json:"value" validate:"required"`
}

type ProductVariant struct {
	ID            uint64            `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     uint64            `gorm:"index;not null" json:"product_id"`
	UoMID         *uint64           `gorm:"index" json:"uom_id,omitempty"`
	Name          string            `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Price         float64           `gorm:"type:decimal(10,2);not null" json:"price" validate:"gt=0"`
	Stock         int               `gorm:"type:int;not null" json:"stock" validate:"gte=0"`
	SKU           string            `gorm:"type:varchar(100); " json:"sku"`
	QuantityValue float64           `gorm:"type:decimal(10,2)" json:"quantity_value,omitempty" validate:"omitempty,gte=0"`
	Unit          UnitOfMeasurement `gorm:"foreignKey:UoMID;references:ID" json:"unit,omitempty"`
	Size          *string           `gorm:"type:varchar(50)" json:"size,omitempty"`
	Weight        *string           `gorm:"type:varchar(50)" json:"weight,omitempty"`
	Color         *string           `gorm:"type:varchar(50)" json:"color,omitempty"`
}

// Table Names

func (Product) TableName() string          { return "products" }
func (ProductAttribute) TableName() string { return "product_attributes" }
func (ProductVariant) TableName() string   { return "product_variants" }

// Validation
func (p *Product) ValidateBasic(validate *validator.Validate) error {
	type ProductValidation struct {
		Name          string   `validate:"required,min=3,max=255"`
		Slug          string   `validate:"required,regexp=^[a-z0-9]+(?:-[a-z0-9]+)*$"`
		Description   string   `validate:"required,min=5"`
		Price         float64  `validate:"required,gt=0"`
		Discount      float64  `validate:"gte=0,lte=100"`
		Stock         int      `validate:"gte=0"`
		SKU           string   `validate:"required"`
		Status        string   `validate:"oneof=active inactive draft"`
		QuantityValue *float64 `validate:"omitempty,gte=0"`
	}

	validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		pattern := fl.Param()
		match, err := regexp.MatchString(pattern, fl.Field().String())
		return err == nil && match
	})

	pv := ProductValidation{
		Name:          strings.TrimSpace(p.Name),
		Slug:          p.Slug,
		Description:   html.EscapeString(strings.TrimSpace(p.Description)),
		Price:         p.Price,
		Discount:      p.Discount,
		Stock:         p.Stock,
		SKU:           p.SKU,
		Status:        p.Status,
		QuantityValue: p.QuantityValue,
	}

	return validate.Struct(pv)
}

// Hooks
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Name = strings.TrimSpace(p.Name)
	p.Brand = strings.TrimSpace(p.Brand)
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Name = strings.TrimSpace(p.Name)
	p.Brand = strings.TrimSpace(p.Brand)
	return nil
}
