package main

import (
	"net/http"
	"text/template"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/teste", TaskHandler)
	http.ListenAndServe(":8888", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprender Go",
		Done: true,
	}

	t := template.Must(template.ParseFiles("task.html"))
	t.Execute(w, task)
}