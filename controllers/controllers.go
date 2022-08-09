package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net/http"


	"github.com/PreBillionaire/mongoAPI/structure"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Add your own mongodb connection url
const connectionString = "mongodb+srv://prebillion:prebillion@cluster0.ge36fn2.mongodb.net/?retryWrites=true&w=majority"
const CreddatabaseName = "Cred"
const CredcollectionName = "cred"
const EmployeedatabaseName = "Employee"
const EmployeecollectionName = "employee"

var SECRET_KEY = []byte("123456")
var contextTODO = context.TODO()
var contextBackground = context.Background()

// make variable of mongoDB collection
// we gonna use it many times
var CredCollection *mongo.Collection
var EmployeeCollection *mongo.Collection

// creating connecting with mongoDB
func init() {

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(contextTODO, clientOption)
	checkErr(err)

	fmt.Println("Mongo Connected Successfully")

	// we have reached inside the database
	// will use this collection to perform operation
	CredCollection = client.Database(CreddatabaseName).Collection(CredcollectionName)
	EmployeeCollection = client.Database(EmployeedatabaseName).Collection(EmployeecollectionName)
	fmt.Println("Collection Instance is Ready")
}

// Controllers
func UserSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cred structure.Cred
	json.NewDecoder(r.Body).Decode(&cred)
	cred.Password = getHash([]byte(cred.Password))
	_, err := CredCollection.InsertOne(contextBackground, cred)
	checkErr(err)
	json.NewEncoder(w).Encode("Registration Successful!")
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cred structure.Cred		// receiving from user
	var dbCred structure.Cred	// already present in db

	json.NewDecoder(r.Body).Decode(&cred)

	// finding if username is already present in db
	filter := bson.M{"username":cred.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&dbCred)
	checkErr(err)

	// if username is present verify the password
	userPass := []byte(cred.Password) // password receiving from user
	dbPass := []byte(dbCred.Password) // password present in db

	// comparing password received from user and password in db
	passErr := bcrypt.CompareHashAndPassword(dbPass,userPass)
	checkErr(passErr)

	json.NewEncoder(w).Encode("Login Successful")
}

func AddEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// finding if username is already present in db
	var employee structure.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	filter := bson.M{"username":employee.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&employee)
	checkErr(err)
	EmployeeCollection.InsertOne(contextBackground,employee)
	json.NewEncoder(w).Encode("Employee Added")
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	// get all data from db using empty filter in Collection.Find
	// will get a cursor 

	var employee structure.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	filter := bson.M{"username":employee.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&employee)
	checkErr(err)
	cursor, err := EmployeeCollection.Find(contextBackground, bson.M{})
	checkErr(err)
	var employees []primitive.M
	// loop through cursor
	for cursor.Next(contextBackground) {
		var employee bson.M
		err := cursor.Decode(&employee)
		checkErr(err)
		employees = append(employees, employee)
	}

	json.NewEncoder(w).Encode(employees)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request){
	var employee structure.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	filter := bson.M{"username":employee.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&employee)
	checkErr(err)
	EmployeeCollection.UpdateOne(contextBackground, filter,
		bson.M{"$set": bson.M{
		"firstname":employee.Firstname,
		"lastname":employee.Lastname,
		"email":employee.Email,
		"salary":"1300000"}})
	json.NewEncoder(w).Encode("Updated")
}

func DeleteOneEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	var employee structure.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	filter := bson.M{"username":employee.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&employee)
	checkErr(err)

	params := mux.Vars(r)
	ok := params["username"]
	// username, _ := primitive.ObjectIDFromHex(ok)
	filter1 := bson.M{"username":ok}
	EmployeeCollection.DeleteOne(contextBackground,filter1)
	json.NewEncoder(w).Encode("Deleted")
}

func DeleteAllEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	var employee structure.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	filter := bson.M{"username":employee.Username}
	err := CredCollection.FindOne(contextBackground,filter).Decode(&employee)
	checkErr(err)
	deleteResult, err := EmployeeCollection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkErr(err)
	json.NewEncoder(w).Encode(deleteResult.DeletedCount)
}

// Extra functions
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
