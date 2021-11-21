package utils

type Stock struct {
	//Id            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string  `json:"name" bson:"name"`
	Price         float32 `json:"price" bson:"price"`
	NumberOfStock uint16  `json:"numberOfStock" bson:"numberOfStock"`
}

// var Stocks = []Stock{
// 	{Name: "ABC1", Price: 85, NumberOfStock: 4},
// 	{Name: "ABC2", Price: 95, NumberOfStock: 45},
// }

const password = "12345"
const user = "amit"
const DBConnectionString = "mongodb+srv://" + user + ":" + password + "@cluster0.1foca.mongodb.net/stock?retryWrites=true&w=majority"
const LocalhostDBConnectionString = "mongodb://127.0.0.1:27017"

// type Stock struct {
// 	Name          string  `json:"name"`
// 	NumberOfStock uint16  `json:"numberofstock"`
// 	Price         float64 `json:"price"`
// }
