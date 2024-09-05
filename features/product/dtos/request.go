package dtos

import "time"

type InputProduct struct {
	Name string `json:"name" form:"name" validate:"required"`
	Images string `json:"images" form:"images"`
	Description string `json:"description" form:"description"`
	Composition string `json:"composition" form:"composition"`
	PomTR string `json:"pom_tr" form:"pom_tr"`
	MarketPlace string `json:"marketplace" form:"marketplace"`
	Whatsapp string `json:"whatsapp" form:"whatsapp"`
	Price string `json:"price" form:"price"`
	Quantity string `json:"quantity" form:"quantity"`
	ProductCreated time.Time `json:"product_created" form:"product_created"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}