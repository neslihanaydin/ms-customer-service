package main

import (
	"customer-service/db"
	"customer-service/service"

	"github.com/gin-gonic/gin"
)

func main() {

	secret := db.GetSecretValue()
	db := db.GetDB(secret)

	app := service.GetApp(db)
	r := gin.Default()

	r.POST("/customers", app.PostHandler)
	r.GET("/customers/:customerId", app.GetHandler)
	r.PUT("/customers/:customerId", app.PutHandler)
	r.DELETE("/customers/:customerId", app.DeleteHandler)

	r.Run("localhost:8080")

}
