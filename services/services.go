package services

import (
	"PowerPlantManagementApplication/constants"
	"PowerPlantManagementApplication/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Sample() {
	var Company models.Company
	Company.CompanyName = "benz"
	Company.CEO = "benz"
	fmt.Println(Company)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Connecting to MongoDb")
	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(constants.ConnectionString))

	if err != nil {
		log.Println("Error connecting to MongoDB: ", err)
	} else {
		log.Println("Conneted to MongoDb successfully")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Database("PowerPlant").Collection("Company")

	result, err := collection.InsertOne(ctx, Company)

	if err != nil {
		log.Println("Error connecting to MongoDB: ", err)
	} else {
		log.Println("Inserted to MongoDb successfully")
	}

	log.Println(result)

	
}
