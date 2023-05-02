package controllers

import (
	"net/http"
	"strings"

	"test-case/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// request
type (
	RequestProduct struct {
		ID                 uint     `gorm:"primary_key" json:"id"`
		Title              string   `json:"title"`
		Desctiption        string   `json:"description"`
		Price              int      `json:"price"`
		DiscountPercentage float64  `json:"discountPercentage"`
		Rating             float64  `json:"rating"`
		Stock              int      `json:"stock"`
		Brand              string   `json:"brand"`
		Category           string   `json:"category"`
		Thumbnail          string   `json:"thumbnail"`
		Images             []string `json:"images"`
	}

	ProductRequest struct {
		RequestID int              `json:"request_id"`
		Data      []models.Product `json:"product"`
	}

	InputDummy struct {
		Total int `json:"total"`
	}
)

// GetAllProducts godoc
// @Summary Get all Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []RequestProduct
// @Router /products [get]
func GetAllProduct(c *gin.Context) {
	var result []RequestProduct
	var products []models.Product
	var images []models.Image

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Find(&products)
	for _, value := range products {

		if err := db.Where("product_id = ?", value.ID).Find(&images).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		var temp RequestProduct
		temp.ID = value.ID
		temp.Title = value.Title
		temp.Desctiption = value.Desctiption
		temp.Price = value.Price
		temp.DiscountPercentage = value.DiscountPercentage
		temp.Rating = value.Rating
		temp.Stock = value.Stock
		temp.Brand = value.Brand
		temp.Category = value.Category
		for _, value := range images {
			if value.Type == "thumbnail" {
				temp.Thumbnail = value.URL
			}
			temp.Images = append(temp.Images, value.URL)
		}
		result = append(result, temp)
	}

	c.JSON(http.StatusOK, gin.H{"product": result, "total": len(result), "skip": 0, "limit": 0})
}

// CreateProduct godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body RequestProduct true "the body to create a new Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var input RequestProduct

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Product
	product := models.Product{
		Title:              input.Title,
		Desctiption:        input.Desctiption,
		Price:              input.Price,
		DiscountPercentage: input.DiscountPercentage,
		Rating:             input.Rating,
		Stock:              input.Stock,
		Brand:              input.Brand,
		Category:           input.Category,
	}
	db.Create(&product)

	if err := db.Where("title = ?", input.Title).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	for _, value := range input.Images {

		var tipe string
		if strings.Contains(value, "thumbnail") {
			tipe = "thumbnail"
		} else {
			tipe = "image"
		}

		// Create Image
		image := models.Image{ProductID: product.ID, Type: tipe, URL: value}
		db.Create(&image)

	}

	input.ID = product.ID

	c.JSON(http.StatusOK, gin.H{"product": input})
}

// GetProductById godoc
// @Summary Get Product by Id.
// @Description Get an Product by Id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist

	var result RequestProduct
	var product models.Product
	var images []models.Image

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Where("product_id = ?", c.Param("id")).Find(&images).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	result.ID = product.ID
	result.Title = product.Title
	result.Desctiption = product.Desctiption
	result.Price = product.Price
	result.DiscountPercentage = product.DiscountPercentage
	result.Rating = product.Rating
	result.Stock = product.Stock
	result.Brand = product.Brand
	result.Category = product.Category
	for _, value := range images {
		if value.Type == "thumbnail" {
			result.Thumbnail = value.URL
		}
		result.Images = append(result.Images, value.URL)
	}

	c.JSON(http.StatusOK, gin.H{"products": result})
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Param Body body RequestProduct true "the body to update product"
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {

	var product models.Product
	var input RequestProduct

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Product
	updatedInput.ID = input.ID
	updatedInput.Title = input.Title
	updatedInput.Desctiption = input.Desctiption
	updatedInput.Price = input.Price
	updatedInput.DiscountPercentage = input.DiscountPercentage
	updatedInput.Rating = input.Rating
	updatedInput.Stock = input.Stock
	updatedInput.Brand = input.Brand
	updatedInput.Category = input.Category

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// DeleteProduct godoc
// @Summary Delete one Product. (Admin only)
// @Description Delete a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	var images []models.Image

	if err := db.Where("product_id = ?", c.Param("id")).Find(&images).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&images)

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"product": true})
}
