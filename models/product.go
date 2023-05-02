package models

type (

	// product
	Product struct {
		ID                 uint    `gorm:"primary_key" json:"id"`
		Title              string  `gorm:"uniqueIndex" json:"title"`
		Desctiption        string  `json:"description"`
		Price              int     `json:"price"`
		DiscountPercentage float64 `json:"discountPercentage"`
		Rating             float64 `json:"rating"`
		Stock              int     `json:"stock"`
		Brand              string  `json:"brand"`
		Category           string  `json:"category"`
		Images             []Image `json:"-"`
	}
)
