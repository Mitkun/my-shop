package productdomain

import "my-shop/common"

type Product struct {
	common.BaseModel
	CategoryId    string
	TypeId        string
	PostId        string
	Name          string
	Thumbnail     string
	Qty           int
	SaleQty       int
	CostOfCapital float32
	Price         float32
	Discount      int8
	Status        string
}
