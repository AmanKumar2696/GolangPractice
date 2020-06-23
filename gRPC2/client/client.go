package main

import (
	"log"

	"github.com/GolangPractice/gRPC2/inventory"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.POST("/additem", inventory.AddItemHandler)
	g.GET("/getitem/:id", inventory.GetDetailsHandler)
	g.GET("/getall/:status", inventory.CheckInventoryHandler)
	g.DELETE("/remove/:id", inventory.RemoveItemHandler)
	g.POST("/updateitem", inventory.UpdateItemHandler)
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
