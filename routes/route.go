package routes

import (
	"recipe-api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		// レシピ作成
		v1.POST("/recipes", handlers.NewRecipeHandler)
		// レシピ一覧の取得
		v1.GET("/recipes", handlers.ListRecipesHandler)
		// レシピの更新
		v1.PUT("/recipes/:id", handlers.UpdateRecipeHandler)
		// レシピの削除
		v1.DELETE("/recipes/:id", handlers.DeleteRecipeHandler)
		// レシピの検索
		v1.GET("/recipes/search", handlers.SearchRecipesHandler)
		return router
	}
}
