package server

// Map for endpoints
func (r *Router) Map() {

	api := r.engine.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register", r.handler.Register)

			auth.POST("/login", r.handler.Login)
		}

		products := api.Group("/products") // product group without JWT
		{
			products.GET("", r.handler.ListProducts)

			products.GET("/:id", r.handler.GetProductByID)
		}

		cart := api.Group("/cart") // cart group with JWT
		cart.Use(r.md.AuthMiddleware())
		{
			cart.GET("", r.handler.GetCart)

			cart.POST("/items", r.handler.AddToCart)

			cart.DELETE("/items/:id", r.handler.RemoveFromCart)
		}

		orders := api.Group("/orders") // orders group with JWT
		orders.Use(r.md.AuthMiddleware())
		{
			orders.GET("", r.handler.)

			orders.GET("")

			orders.GET("/:id")

			orders.POST("/:id/pay")


		}

	}

}
