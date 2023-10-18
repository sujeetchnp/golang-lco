package main

// import (
// 	"context"
// 	"log"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type YourDocumentType struct {
// 	ReportService struct {
// 		QPAMessage string `bson:"qpamessage"`
// 	} `bson:"reportservice"`
// 	SimpleClaim struct {
// 		Services struct {
// 			AuditTrail []struct {
// 				QPAMessage string `bson:"qpamessage"`
// 			} `bson:"auditTrail"`
// 		} `bson:"services"`
// 	} `bson:"simpleClaim"`
// }

// func main() {
// 	// MongoDB connection parameters
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(context.Background())

// 	// Choose your database and collection
// 	collection := client.Database("your-database").Collection("your-collection")

// 	// Define a filter to get the document
// 	filter := bson.M{}

// 	// Define options for sorting in descending order (to get the max index)
// 	opts := options.Find()
// 	opts.SetSort(bson.D{{"simpleClaim.services.auditTrail.qpaMessage", -1}})
// 	opts.SetLimit(1)

// 	// Perform the query to find the document with the maximum qpaMessage index
// 	cursor, err := collection.Find(context.Background(), filter, opts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cursor.Close(context.Background())

// 	// Iterate over the result(s)
// 	var doc YourDocumentType
// 	if cursor.Next(context.Background()) {
// 		if err := cursor.Decode(&doc); err != nil {
// 			log.Fatal(err)
// 		}

// 		// Update reportservice.qpamessage with the maximum index
// 		maxIndex := len(doc.SimpleClaim.Services.AuditTrail) - 1
// 		doc.ReportService.QPAMessage = doc.SimpleClaim.Services.AuditTrail[maxIndex].QPAMessage

// 		// Update the document in the collection
// 		updateFilter := bson.M{"_id": doc.ID} // Assuming you have an "_id" field
// 		update := bson.M{
// 			"$set": bson.M{
// 				"reportservice.qpamessage": doc.ReportService.QPAMessage,
// 			},
// 		}
// 		_, err := collection.UpdateOne(context.Background(), updateFilter, update)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
