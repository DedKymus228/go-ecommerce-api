package server

import (
	"github.com/gin-gonic/gin"
)

// Map for endpoints
func (r *Router) Map() {

	api := r.engine.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register", r.handler.Register)

			auth.POST("/login", r.handler.Login)
		}

		product := r.engine.Group("/product") // product group without JWT
		{
			product.GET("", r.handler.ListProducts)

			product.GET("/:id", r.handler.GetProductByID)
		}

		cart := r.engine.Group("/cart") // cart group with JWT
		cart.Use(r.md.AuthMiddleware())
		{
			cart.GET("", func(*gin.Context) {

			})

			cart.POST("/item", r. )

			cart.DELETE("/item/:id", func(*gin.Context) {})
		}

	}

}
