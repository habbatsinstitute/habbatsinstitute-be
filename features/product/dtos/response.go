package dtos

import "time"

type ResProduct struct {
	Name string `json:"name" form:"name"`
	Images string `json:"images" form:"images"`
	Description string `json:"description" form:"description"`
	Composition string `json:"composition" form:"composition"`
	PomTR string `json:"pom_tr" form:"pom_tr"`
	MarketPlace string `json:"marketplace" form:"marketplace"`
	Whatsapp string `json:"whatsapp" form:"whatsapp"`
	Price int64 `json:"price" form:"price"`
	Quantity int `json:"quantity" form:"quantity"`
	ProductCreated time.Time `json:"created_at" form:"created_at"`
}
