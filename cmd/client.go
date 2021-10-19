package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func getTask() int {
	resp, err := http.Get("http://localhost:5000/gettask")
	if err != nil {
		log.Fatal(err)
	}

	task := []Task{}

	json.NewDecoder(resp.Body).Decode(&task)
	for _, values := range task {
		fmt.Println(values.ID, "\t", values.Name, "\t", values.Status)
	}
	return resp.StatusCode

}
func createTask(t []string) int {
	task := &Task{
		Name:   t[0],
		Status: t[1],
	}
	data, _ := json.Marshal(&task)
	resp, err := http.Post("http://localhost:5000/createtask", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
func updateTask(t []string) int {
	task := &Task{
		Name:   t[0],
		Status: t[1],
	}
	data, _ := json.Marshal(&task)
	resp, err := http.Post("http://localhost:5000/updatetask", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
func deleteTask(t []string) int {
	task := &Task{
		Name: t[0],
	}
	data, _ := json.Marshal(&task)
	resp, err := http.Post("http://localhost:5000/deletetask", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}
