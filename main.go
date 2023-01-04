package main

import (
	barang "api/barang/controller"
	"api/barang/models"
	"api/config"
	user "api/user/controller"
	"api/user/model"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)
	model := model.UserModel{DB: db}
	controll := user.UserControll{Mdl: model}

	models := models.ItemsModel{DB: db}
	controllItems := barang.ItemControll{MDL: models}
	e.Pre(middleware.RemoveTrailingSlash()) // fungsi ini dijalankan sebelum routing

	e.Use(middleware.CORS()) // WAJIB DIPAKAI agar tidak terjadi masalah permission
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// e.Use(middleware.Logger()) // Dipakai untuk membuat log (catatan) ketika endpoint diakses

	e.POST("/users", controll.Insert())
	e.GET("/users", controll.GetAll())
	// e.GET("/users/barang", controll.GetAllItem())
	e.POST("/login", controll.Login())

	needLogin := e.Group("/users")
	needLogin.Use(middleware.JWT([]byte("BE!4a|t3rr4")))

	needLogin.GET("", controll.GetID())
	needLogin.PATCH("/patch", controll.Update())
	// PATCH localhost:8000/users/:id/patch
	needLogin.PUT("", controll.Update2())
	needLogin.DELETE("", controll.Delete())

	// e.POST("/barang", controllItems.InsertItem())
	needLogins := e.Group("/barang")
	needLogins.Use(middleware.JWT([]byte("BE!4a|t3rr4")))
	needLogins.POST("", controllItems.InsertItem())
	needLogins.GET("", controllItems.GetAllItem())
	needLogins.PUT("", controllItems.UpdateItems())
	needLogins.DELETE("", controllItems.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
