package server

import "github.com/gin-gonic/gin"

// Map for endpoints
func (r *Router) Map() {

	auth := r.engine.Group("/api/v1/auth")
	{
		auth.POST("/register", func(ctx *gin.Context) {


		})
		auth.POST("/login", func(ctx *gin.Context) {

		})

	}

	product := r.engine.Group("/api/v1") // product group without JWT
	{
		product.GET("/products", func(ctx *gin.Context) {

		})
		product.GET("/products/:id", func(ctx *gin.Context) {

		})


	}
}
