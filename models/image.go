package models

type (

	// image
	Image struct {
		ID        uint    `gorm:"primary_key" json:"id"`
		ProductID uint    `json:"product_id"`
		Type      string  `json:"type"`
		URL       string  `json:"url"`
		Product   Product `json:"-"`
	}
)
