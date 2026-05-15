package server

import (
	"github.com/gin-gonic/gin"
)

// Map for endpoints
func (r *Router) Map() {

	api := r.engine.Group("/api")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register",r.handler.Register)

			auth.POST("/login", r.handler.Login)
		}


		product := r.engine.Group("/api/v1") // product group without JWT
		{
			product.GET("/products", func(ctx *gin.Context) {

			})
			product.GET("/products/:id", func(ctx *gin.Context) {

			})


		}
	}

}
