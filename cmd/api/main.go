package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/config"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/database"
	"github.com/ihsanpraditya/docker-golang-postgres-api-boilerplate/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg.DB)

	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8080")
}
