package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

////////////////////////////////////////////////////////
func main() {

	log.Println("Starting the HTTP server on port 5000")
	r := mux.NewRouter()
	//log.Fatal(http.ListenAndServe("localhost:5000", router))
	http.HandleFunc("/", handler)
	r.HandleFunc("/gettask", logging(getTasks)).Methods("GET")
	r.HandleFunc("/createtask", logging(createTask)).Methods("POST")
	r.HandleFunc("/deletetask", logging(deleteTask)).Methods("POST")
	r.HandleFunc("/updatetask", logging(updateTask)).Methods("POST")
	//http.HandleFunc("/gettask", getTasks)
	http.ListenAndServe(":5000", r)

}

/////////////////////////////////
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

////////////////////////////////////////////////////////
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

////////////////////////////////////////////////////////
func getTasks(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	task := getTaskAll(db)
	json.NewEncoder(w).Encode(task)
}

////////////////////////////////////////////////////////
func createTask(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	decoder := json.NewDecoder(r.Body)
	var t Task
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "You've the task: %s with status %s\n", t.Name, t.Status)
	insert(db, t)
}

/////////////////////////////////////////////////////////
func deleteTask(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	decoder := json.NewDecoder(r.Body)
	var t Task
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Deleted the task: %s \n", t.Name)
	delete(db, t)
	task := getTaskAll(db)
	json.NewEncoder(w).Encode(task)
}

/////////////////////////////////////////////////////////
func updateTask(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	decoder := json.NewDecoder(r.Body)
	var t Task
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Updated the task: %s with status %s\n", t.Name, t.Status)
	update(db, t)
	task := getTaskAll(db)
	json.NewEncoder(w).Encode(task)
}
