package models

type Product struct {
	Id            int     `json:"id" gorm:"autoIncrement;column:id"`
	SourceId      int     `json:"source_id" gorm:"column:sourceid"`
	ProductTitle  string  `json:"product_title" gorm:"column:title"`
	ProductPrice  float32 `json:"product_price" gorm:"column:price"`
	StockQuantity uint16  `json:"stock_quantity" gorm:"column:stock"`
}

type CreateProduct struct {
	SourceId      int     `json:"id" gorm:"column:sourceid"`
	ProductTitle  string  `json:"product_title" gorm:"column:title"`
	ProductPrice  float32 `json:"product_price" gorm:"column:price"`
	StockQuantity uint16  `json:"stock_quantity" gorm:"column:stock"`
}
