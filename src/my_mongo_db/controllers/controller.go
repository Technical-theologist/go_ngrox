package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"my_mango/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rakesh:rakesh12@cluster0.gzipf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbname = "netflix"
const colname = "watchList"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbname).Collection(colname)

	fmt.Println("Collection instance is ready")

}

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with Id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("movie got delete with with delete count:", deleteCount)
}

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted :", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
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

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
