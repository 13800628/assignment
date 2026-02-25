package main

import (
	"assignment/internal/infrastructure/config"
	db_conn "assignment/internal/infrastructure/database"
	"assignment/internal/interfaces/controller"
	repo_impl "assignment/internal/interfaces/database"
	"assignment/internal/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. 設定読み込み
	cfg := config.Load()

	// 2. DB接続
	db := db_conn.NewMySQLDB(cfg)

	// 3. 依存関係の注入 (DI)
	itemRepo := repo_impl.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepo)
	itemController := controller.NewItemController(itemUsecase)

	// 4. Echo起動
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	e.GET("/items", itemController.GetItems)
	e.POST("/items", itemController.CreateItem)
	e.GET("/items/:id", itemController.GetItemByID)
	e.DELETE("/items/:id", itemController.DeleteItem)
	e.GET("/items/summary", itemController.GetSummary)
	e.PATCH("/items/:id", itemController.PatchItem)

	log.Fatal(e.Start(":8080"))
}
