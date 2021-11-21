package dbconnector

import (
	"context"
	"first_assigment/utils"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabase(client *mongo.Client, ctx context.Context) []string {
	databases, error := client.ListDatabaseNames(ctx, bson.M{})
	if error != nil {
		log.Fatal("Error while geting the list of database : ", error)
		return []string{}
	}
	fmt.Printf("List of databases %v\n", databases)
	return databases

}

func GetStocks(client *mongo.Client) (results []utils.Stock) {
	collection := client.Database("stock").Collection("stock") // Stock is databse and stock is collection
	//filter := bson.M{}
	//var result utils.Stock
	// var result bson.M

	// To get single value from db
	// err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error while retriving the data", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var stock utils.Stock
		cursor.Decode(&stock)
		results = append(results, stock)

	}
	fmt.Printf("Document: %v\n", results)
	return results
}

func AddStock(client *mongo.Client, stock utils.Stock) bool {
	collection := client.Database("stock").Collection("stock")
	insetedResult, err := collection.InsertOne(context.TODO(), stock)
	if err != nil {
		fmt.Println("Error while inserting the data ", err)
		return false
	}
	fmt.Printf("Inserted result : %v\n", insetedResult)
	return true
}

func GetStockByName(client *mongo.Client, stockName string) (result utils.Stock) {
	collection := client.Database("stock").Collection("stock")
	err := collection.FindOne(context.TODO(), bson.M{"name": stockName}).Decode(&result)
	if err != nil {
		fmt.Println("Error while retrivng the data", err)
		return result
	}
	return result
}

func DeleteStockByName(client *mongo.Client, stockName string) bool {
	collection := client.Database("stock").Collection("stock")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"name": stockName})
	if err != nil {
		log.Fatal("Error while deleting the data")
		return false
	}
	return true

}
