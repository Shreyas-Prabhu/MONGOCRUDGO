package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreyas-Prabhu/MONGOCRUDGO/initializers"
	"github.com/Shreyas-Prabhu/MONGOCRUDGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = initializers.MongoConfig()

func InsertMovie(movie models.Movie) string {
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return "Inserted"
}

func GetOneMovie(id string) *models.Movie {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var movie models.Movie
	filter := bson.M{"_id": movid}
	res := collection.FindOne(context.Background(), filter).Decode(&movie)
	fmt.Println(res)
	return &movie
}

func GetAllMovies() []bson.M {
	filter := bson.D{{}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies

}

func UpdateMovie(id string, movie models.Movie) *models.Movie {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": movid}
	update := bson.M{"$set": bson.M{"name": movie.Name, "year": movie.Year}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	m := GetOneMovie(id)
	return m
}

func DeleteOneMovie(id string) string {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": movid}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return "Deleted"
}

func DeleteAll() string {
	filter := bson.D{{}}
	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return "Deleted All"
}
