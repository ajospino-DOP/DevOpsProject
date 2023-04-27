package server

import (
	"context"
	"encoding/json"
	"errors"
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
	Name string
	Value string
}

var savedInfo []*Info = []*Info{}



func initDB() (*mongo.Client){
	var uri = os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Client unnable to connect %v", err)
		panic(err)
	}
	
	return client
}

 func discDB(client *mongo.Client) {
	_ = client.Disconnect(context.TODO())
}

//New function that returns error
func indexReg(w http.ResponseWriter, r *http.Request)(error){
	if(r.Method != http.MethodGet){
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		err := errors.New("Method not allowed")
		return err
	}
	fmt.Fprintf(w, "Welcome visitor")
	return nil
}

//This is the index function
func index(w http.ResponseWriter, r *http.Request){
	indexReg(w,r)
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
		fmt.Fprintf(w, r.Method)
		return
	}
}	

func listInfo(w http.ResponseWriter, r *http.Request)(error){
	w.Header().Set("Content-Type", "applicaton/json")
	savedInfo = getDB()
	err := json.NewEncoder(w).Encode(savedInfo)
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return err
	}
	return nil
}

func sendInfo(w http.ResponseWriter, r *http.Request)(error){
	newInfo :=  &Info{}
	err := json.NewDecoder(r.Body).Decode(newInfo)
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return err
	}
	saveDB(newInfo)
	return nil
}

func getDB()([]*Info){
	client := initDB()
	collection := client.Database("DOP").Collection("info")
	filter := bson.D{}
	res, err := collection.Find(context.TODO(),filter)
	if err != nil{
		panic(err)
	}
	
	var results []*Info
	if err = res.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
	/**
	for _, result := range results {
		res.Decode(&result)
		output, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	*/
}


func saveDB(i *Info){
	client := initDB()
	collection := client.Database("DOP").Collection("info")
	res, err := collection.InsertOne(context.TODO(),i)
	
	if err != nil { 
		fmt.Printf("Insert error %v", err)
		panic(err) }
	
	id := res.InsertedID

	fmt.Printf("Object inserted succesfully, id: %d", id)
}