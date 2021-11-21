package services

import (
	"first_assigment/dbconnector"
	"first_assigment/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetStocks(client *mongo.Client) []utils.Stock {
	fmt.Println("In get call")
	return dbconnector.GetStocks(client)
	// a := utils.Stocks
	// fmt.Println("Value of a : ", a)
	// return utils.Stocks
}

func PostStocks(stocks utils.Stock, client *mongo.Client) bool {
	fmt.Println("In the post call")
	return dbconnector.AddStock(client, stocks)
	// utils.Stocks = append(utils.Stocks, stocks)
}

func GetByName(name string, client *mongo.Client) utils.Stock {
	return dbconnector.GetStockByName(client, name)

}

func DeleteByName(name string, client *mongo.Client) bool {
	return dbconnector.DeleteStockByName(client, name)
}
