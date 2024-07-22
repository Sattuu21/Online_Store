package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type details struct {
	ProductId string `json:"product_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

const dbName = "OnlineStore"
const colName = "ProductDetails"

var collection *mongo.Collection

func init() {
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection reference is ready")
}

func insertOneProductDetail(product details) {
	inserted, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted data:", inserted.InsertedID)

}
