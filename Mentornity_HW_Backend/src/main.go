package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type Item struct {
	//ID     string  `json:"id"`
	//Name   string  `json:"name"`
	//Email  string  `json:"email"`
	//Message string `json:"message"`
	Name string  `json:"Name"`
	Title string  `json:"email"`
	Body string  `json:"message"`

}
type Form struct {

	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"Name" bson:"Name"`
	//Name   string             `json:"Name"`
	Email  string              `json:"Email" bson:"Email"`
	//Email  string              `json:"Email"`
	Message  string             `json:"Message" bson:"Message"`
	//Message  string             `json:"Message"`
}
var items []Item
const connectionString = "mongodb+srv://ercanozturk:1234@cluster0.rtht0.mongodb.net/FormApp?retryWrites=true&w=majority"

// Database Name
const dbName = "FormApp"

// Collection name
const collName = "Forms"

// collection object/instance
var collection *mongo.Collection
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
}

func getItems(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//writer.Header().Set("Content-Type", "; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(writer).Encode(items)
}
/*
func getItem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	//(writer).Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(request)
	for _, i := range items{
		if i.ID == params["ID"]{
			json.NewEncoder(writer).Encode(i)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Item{})
}
*/
func createItem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(request.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(writer).Encode(item)


}

// CreateTask create task route
func CreateTask(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Type", "X-www-form-urlencoded")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "POST")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task Form


	_ = json.NewDecoder(r.Body).Decode(&task)
	 fmt.Println(task, r.Body)
	//insertOneTask(task)
	collection.InsertOne(context.Background(), task)

	json.NewEncoder(w).Encode(task)
}


// Insert one task in the DB
func insertOneTask(task Form) {
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", task)
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func main() {



	r := mux.NewRouter()
	//r.HandleFunc("/api/items", GetAllTask).Methods("GET")
	r.HandleFunc("/api/items", CreateTask).Methods("POST")
	//r.HandleFunc("/api/items", getItems).Methods("GET")
	//r.HandleFunc("/api/items", createItem).Methods("POST")
	//r.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))




}





