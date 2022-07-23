package model

import "time"

type Product struct {
	Id            string     `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ImageId       *string    `json:"image_id"`
	Price         int        `json:"price"`
	CurrencyId    int        `json:"currency_id"`
	Rating        *int       `json:"rating"`
	CategoryId    int        `json:"category_id"`
	Specification *string    `json:"specification"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
