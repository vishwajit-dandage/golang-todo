package main

import (
	"net/http"
	"testing"
)

func TestGetTask(t *testing.T) {
	code := getTask()
	if code != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			code, http.StatusOK)
	}
}
func TestCreateTask(t *testing.T) {
	task := []string{"Task Test", "Task Test"}
	code := createTask(task)
	if code != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			code, http.StatusOK)
	}
}
