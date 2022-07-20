package model

type Product struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImageId       string `json:"image_id"`
	Price         string `json:"price"`
	CurrencyId    string `json:"currency_id"`
	Rating        string `json:"rating"`
	CategoryId    string `json:"category_id"`
	Specification string `json:"specification"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
