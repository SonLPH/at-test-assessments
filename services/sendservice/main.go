package main

import (
	"at-home-assessments/services/sendservice/handler"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var cred = options.Credential{
	Username: "admin",
	Password: "admin",
}

func main() {
	godotenv.Load()
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(cred)
	client, _ := mongo.Connect(ctx, clientOptions)
	collection = client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("EMPLOYEE_COLLECTION"))

	defer client.Disconnect(context.Background())

	sendService := handler.SendService{MongoCollection: collection}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	r.Handle("/send", http.HandlerFunc(sendService.Send)).Methods(http.MethodPost)
	r.Handle("/mock", http.HandlerFunc(sendService.MockEmployee)).Methods(http.MethodPost)

	log.Println("Service is running on port 4446")
	http.ListenAndServe(":4446", r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
