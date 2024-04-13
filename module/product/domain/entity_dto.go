package productdomain

import "github.com/google/uuid"

type ProductUpdateDTO struct {
	Name          *string  `gorm:"column:name" json:"name"`
	CategoryId    *string  `gorm:"column:category_id" json:"category_id"`
	TypeId        *string  `gorm:"column:type" json:"type"`
	Thumbnail     *string  `gorm:"column:thumbnail" json:"Thumbnail"`
	PostId        *string  `gorm:"column:post_id" json:"post_id"`
	Qty           *int     `gorm:"column:q_ty" json:"q_ty"`
	SaleQty       *int     `gorm:"column:sale_q_ty" json:"sale_q_ty"`
	CostOfCapital *float32 `gorm:"column:cost_of_capital" json:"cost_of_capital"`
	Price         *float32 `gorm:"column:price" json:"price"`
	Discount      *int8    `gorm:"column:discount" json:"discount"`
	Status        *string  `gorm:"column:status" json:"status"`
}

type ProductCreationDTO struct {
	Id            uuid.UUID `gorm:"column:id" json:"id"`
	Name          string    `gorm:"column:name" json:"name"`
	CategoryId    string    `gorm:"column:category_id" json:"category_id"`
	TypeId        string    `gorm:"column:type" json:"type"`
	Thumbnail     string    `gorm:"column:thumbnail" json:"Thumbnail"`
	PostId        *string   `gorm:"column:post_id" json:"post_id"`
	Qty           int       `gorm:"column:q_ty" json:"q_ty"`
	SaleQty       *int      `gorm:"column:sale_q_ty" json:"sale_q_ty"`
	CostOfCapital float32   `gorm:"column:cost_of_capital" json:"cost_of_capital"`
	Price         float32   `gorm:"column:price" json:"price"`
	Discount      int8      `gorm:"column:discount" json:"discount"`
	Status        string    `gorm:"column:status" json:"status"`
}

func (ProductCreationDTO) TableName() string { return "products" }
func (ProductUpdateDTO) TableName() string   { return ProductCreationDTO{}.TableName() }
