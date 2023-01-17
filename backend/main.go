package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Todo struct {
	ID   int
	Text string
	Completed bool
}

var todos = []Todo{
	{
		ID:   1,
		Text: "Learn Svelte",
	},
	{
		ID:   2,
		Text: "Learn Go",
	},
	{
		ID:   3,
		Text: "Build a Todo App",
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		handleTodos(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleTodosGet(w, r)
	case "POST":
		handleTodoCreate(w, r)
	case "PUT":
		handleTodoUpdate(w, r)
	case "DELETE":
		handleTodoDelete(w, r)
	}
}

func handleTodosGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func handleTodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func handleTodoUpdate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	for i, t := range todos {
		if t.ID == todo.ID {
			todos[i] = todo
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func handleTodoDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/api/todos/"):])
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i + 1:]...)
			break
		}
	}
}

