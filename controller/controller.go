package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maharaj2113/test/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectstring = "mongodb+srv://suhasrao:suhasrao@cluster0.evk2arw.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Prime"
const colName = "Movies"

var collection *mongo.Collection

func init() {
	conOption := options.Client().ApplyURI(connectstring)

	client, err := mongo.Connect(context.TODO(), conOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Connection Instance is Ready")
}
func insert(movie model.Prime) {
	insterted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One movie with ID ", insterted.InsertedID)
}
func update(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	upda := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, upda)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Movies in DB ", result.ModifiedCount)
}
func deleteOne(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie count is ", deleted)

}
func deleteMany() int64 {
	delet, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The deleted movie count is ", delet.DeletedCount)
	return delet.DeletedCount
}
func getMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	allmovie := getMovies()
	json.NewEncoder(w).Encode(allmovie)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var movie model.Prime
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insert(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")
	params := mux.Vars(r)

	update(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}
func DeleOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	params := mux.Vars(r)
	deleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count := deleteMany()
	json.NewEncoder(w).Encode(count)
}
