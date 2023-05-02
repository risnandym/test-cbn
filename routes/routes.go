package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"test-case/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/products", controllers.GetAllProduct)
	r.POST("/product", controllers.CreateProduct)
	r.GET("/product/:id", controllers.GetProductById)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
