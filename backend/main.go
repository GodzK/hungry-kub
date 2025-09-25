package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Food struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Price int                `json:"price" bson:"price"`
	Image string             `json:"image" bson:"image"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Foods     []Food             `json:"foods" bson:"foods"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

var client *mongo.Client
var foodCollection *mongo.Collection
var orderCollection *mongo.Collection

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func connectDB() {
	uri := os.Getenv("MONGO_URI") // read from .env

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("foodapp")
	foodCollection = db.Collection("foods")
	orderCollection = db.Collection("orders")
	fmt.Println("✅ Connected to MongoDB Atlas")
}

func getFoods(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := foodCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var foods []Food
	if err = cursor.All(ctx, &foods); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var foods []Food
	if err := json.NewDecoder(r.Body).Decode(&foods); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	order := Order{
		Foods:     foods,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"orderId": res.InsertedID,
		"message": "Order placed successfully",
	})
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback port
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/foods", getFoods)
	mux.HandleFunc("/order", createOrder)

	fmt.Println("🚀 Backend running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, enableCORS(mux)))
}
