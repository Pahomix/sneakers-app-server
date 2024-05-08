package routes

import (
	"github.com/gin-gonic/gin"
	"sneakers-app/internal/handlers/auth"
	"sneakers-app/internal/handlers/brand"
	"sneakers-app/internal/handlers/category"
	"sneakers-app/internal/handlers/promotion"
	"sneakers-app/internal/handlers/review"
	"sneakers-app/internal/handlers/size"
	"sneakers-app/internal/handlers/sneakers"
	"sneakers-app/internal/handlers/user"
)

func InitRoutes(routes *gin.Engine) {
	routes.POST("/login", auth.LoginHandler)
	routes.POST("/register", auth.RegisterHandler)

	userRoutes := routes.Group("/users")
	{
		userRoutes.GET("/:id", auth.RequireAuth, user.GetUserById)
		userRoutes.GET("/", auth.RequireAuth, user.GetUsers)
		userRoutes.POST("/", auth.RequireAuth, user.CreateUser)
		userRoutes.PUT("/:id", auth.RequireAuth, user.UpdateUser)
		userRoutes.DELETE("/:id", auth.RequireAuth, user.DeleteUser)
	}

	userRoutes = routes.Group("/sneakers")
	{
		userRoutes.GET("/", auth.RequireAuth, sneakers.GetSneakersHandler)
		userRoutes.GET("/:id", auth.RequireAuth, sneakers.GetSneakerHandler)
		userRoutes.POST("/", auth.RequireAuth, sneakers.CreateSneakerHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, sneakers.UpdateSneakerHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, sneakers.DeleteSneakerHandler)
	}

	userRoutes = routes.Group("/reviews")
	{
		userRoutes.GET("/", auth.RequireAuth, review.GetReviewsHandler)
		userRoutes.GET("/:id", auth.RequireAuth, review.GetReviewHandler)
		userRoutes.POST("/", auth.RequireAuth, review.CreateReviewHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, review.UpdateReviewHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, review.DeleteReviewHandler)
	}

	userRoutes = routes.Group("/sizes")
	{
		userRoutes.GET("/", auth.RequireAuth, size.GetSizesHandler)
		userRoutes.GET("/:id", auth.RequireAuth, size.GetSizeHandler)
		userRoutes.POST("/", auth.RequireAuth, size.CreateSizeHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, size.UpdateSizeHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, size.DeleteSizeHandler)
	}

	userRoutes = routes.Group("/promotions")
	{
		userRoutes.GET("/", auth.RequireAuth, promotion.GetPromotionsHandler)
		userRoutes.GET("/:id", auth.RequireAuth, promotion.GetPromotionHandler)
		userRoutes.POST("/", auth.RequireAuth, promotion.CreatePromotionHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, promotion.UpdatePromotionHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, promotion.DeletePromotionHandler)
	}

	userRoutes = routes.Group("/brands")
	{
		userRoutes.GET("/", auth.RequireAuth, brand.GetBrandsHandler)
		userRoutes.GET("/:id", auth.RequireAuth, brand.GetBrandsHandler)
		userRoutes.POST("/", auth.RequireAuth, brand.CreateBrandHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, brand.UpdateBrandHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, brand.DeleteBrandHandler)
	}

	userRoutes = routes.Group("/categories")
	{
		userRoutes.GET("/", auth.RequireAuth, category.GetCategoriesHandler)
		userRoutes.GET("/:id", auth.RequireAuth, category.GetCategoriesHandler)
		userRoutes.POST("/", auth.RequireAuth, category.CreateCategoryHandler)
		userRoutes.PUT("/:id", auth.RequireAuth, category.UpdateCategoryHandler)
		userRoutes.DELETE("/:id", auth.RequireAuth, category.DeleteCategoryHandler)
	}
}
