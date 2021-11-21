package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"first_assigment/dbconnector"
	"first_assigment/services"
	"first_assigment/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
	err    error
)

func getStocks(c *gin.Context) {

	// var a = services.GetStocks()
	c.JSON(http.StatusOK, services.GetStocks(client))
}

func postStocks(c *gin.Context) {
	//services.PostStocks()
	var requestBody utils.Stock

	// BindJSON to bind the received JSON.
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	stock := utils.Stock{
		Name:          requestBody.Name,
		NumberOfStock: requestBody.NumberOfStock,
		Price:         requestBody.Price,
	}
	fmt.Printf("%v\n", stock)
	if services.PostStocks(stock, client) {
		c.JSON(http.StatusCreated, stock)
	} else {
		c.JSON(http.StatusNotModified, gin.H{"message": "Data not updated"})
	}

	// c.JSON(200, gin.H{
	// 	"message": "Stocks are added",
	// })
}

func getByName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusFound, services.GetByName(name, client))
}

func deleteByName(c *gin.Context) {
	name := c.Param("name")
	if services.DeleteByName(name, client) {
		c.JSON(http.StatusAccepted, gin.H{"message": "Data deleted"})
	} else {
		c.JSON(http.StatusNotModified, gin.H{"message": "Not deleted"})
	}

}

func rootPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":      "Stock service started",
		"getStock":     "/stock",
		"postStock":    "/stock",
		"getByName":    "/stock/{{name}}",
		"deleteByName": "/stock/{{name}}",
	})
}

func main() {
	// Creates a gin router with default middleware
	// we can remove this part in other service
	//gin.SetMode(gin.ReleaseMode)
	client, ctx, cancel, err = dbconnector.Connect(utils.LocalhostDBConnectionString)
	// To Close the conneciton once the main exist
	defer dbconnector.Close(client, ctx, cancel)

	if err != nil {
		log.Fatal("Error in main while getting the connection : ", err)
	}
	err = dbconnector.Ping(client, ctx)
	if err != nil {
		log.Fatal("Error while pining the db : ", err)
	}

	var router *gin.Engine = gin.Default()
	router.GET("/", rootPage)
	router.GET("/stock", getStocks)
	router.POST("/stock", postStocks)
	router.GET("/stock/:name", getByName)
	router.DELETE("/stock/:name", deleteByName)
	router.Run(":8080") // listen and serve on port 8080

}
