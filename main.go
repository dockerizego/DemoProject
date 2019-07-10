package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var users []User

const DATE_FORMAT = "20060102150405"
const PORT = "8080"
var logLocation string = "./logs"

func writeLog(message string) {
	f, err := os.OpenFile(logLocation + "/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	t := time.Now()
	log.Println(fmt.Sprintf("%s: %s", t.Format(DATE_FORMAT), message))
}

func main() {
	writeLog("Starting Application...")
	writeLog(fmt.Sprintf("Application is now running on port %s", PORT))
	fmt.Print("Started. Listening on port ", PORT)

	//creating local array
	users = append(users, User{ID: "1", Firstname: "Sachin", Lastname: "Mathew"})
	users = append(users, User{ID: "2", Firstname: "Shi Kai", Lastname: "Khiew"})
	users = append(users, User{ID: "3", Firstname: "Robin", Lastname: "Kedia"})
	users = append(users, User{ID: "4", Firstname: "Mukesh", Lastname: "Gadupudi"})

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}

func GetUserById(w http.ResponseWriter, req *http.Request) {
	writeLog("GetUserById() called")
	params := mux.Vars(req)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// if not found return empty object with User structure
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	writeLog("CreateUser() called")
	params := mux.Vars(req)
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	writeLog("GetUsers() called")
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	writeLog("DeleteUser() called")
	params := mux.Vars(req)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
