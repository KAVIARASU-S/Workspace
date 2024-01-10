package main

import (
	"PowerPlantManagementApplication/constants"
	"fmt"

	//"PowerPlantManagementApplication/services"
	"PowerPlantManagementApplication/models"
	"context"

	//"fmt"
	"log"
	"time"

	//"net/http"

	//"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// server := gin.Default()

	// 	// Define a route for handling POST requests
	// 	server.POST("/sample", func(c *gin.Context) {
	// 		var jsonInput struct {
	// 			Status string `json:"status"`
	// 		}
	
	// 		// Bind the JSON payload to the struct
	// 		if err := c.ShouldBindJSON(&jsonInput); err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{
	// 				"error": "Invalid JSON format",
	// 			})
	// 			return
	// 		}
	
	// 		services.Sample()
	// 		// Check if the status is "success"
		
	// 			c.JSON(http.StatusOK, gin.H{
	// 				"response": "happy",
	// 	})
				
	// server.Run(constants.Port)


	 var Company models.Company
	Company.CompanyName = "ford"
	// Company.CEO = "ford"
	// fmt.Println(Company)

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	log.Println("Connecting to MongoDb")
	mongo,err := mongo.Connect(ctx,options.Client().ApplyURI(constants.ConnectionString))

	if err != nil {
		log.Println("Error connecting to MongoDB: ",err)
	} else {
		log.Println("Conneted to MongoDb successfully")
	}

	ctx,cancel = context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	 collection := mongo.Database("PowerPlant").Collection("Company")
	
	
	// result,err := collection.InsertOne(ctx,Company)

	// if err != nil {
	// 	log.Println("Error connecting to MongoDB: ",err)
	// } else {
	// 	log.Println("Inserted to MongoDb successfully")
	// }

	// fmt.Println(result)


	result,err := collection.Find(ctx,bson.D{})

	if err != nil {
		log.Println("Error finding to MongoDB: ",err)
	} else {
		log.Println("finding to MongoDb successfully")
	}

	var results []bson.D

	result.All(ctx,&results)

	for _,i := range results{
		fmt.Println(i)
	}

	//fmt.Println(result)
}
