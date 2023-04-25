package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

var envErr = godotenv.Load(".env")


type Info struct {
	name string
	value string
}

var savedInfo []*Info = []*Info{}

var uri = os.Getenv("MONGODB_URI")


func initDB() (*mongo.Client){
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return client
}



func index(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodGet){
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Welcome visitor")
}

func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to the X application, simply send values to save with your request or reload this page to see past values")
		listInfo(w, r)
	case http.MethodPost:
		sendInfo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}	

func listInfo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(savedInfo)
}

func sendInfo(w http.ResponseWriter, r *http.Request){
	newInfo := &Info{}
	err := json.NewDecoder(r.Body).Decode(newInfo)
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	savedInfo = append(savedInfo, newInfo)
}

func getDB(){
	client := initDB()
	collection := client.Database("DOP").Collection("info")
	filter := bson.D{}
	res, err := collection.Find(context.TODO(),filter)
	if err != nil{
		panic(err)
	}
	
}


func saveDB(i Info){
	client := initDB()
	collection := client.Database("DOP").Collection("info")
	res, err := collection.InsertOne(context.Background(),i)
	
	if err != nil { panic(err) }
	
	id := res.InsertedID

	fmt.Printf("Object inserted succesfully, id: %d", id)
}