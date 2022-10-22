package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var tmpl *template.Template
var jsonData []byte

type Todo struct {
	Item string
	Done bool
	ID   int
}

type PageData struct {
	Title string
	Todos []Todo
}

var htmlData = PageData{
	Title: "TODO List",
	Todos: []Todo{
		{Item: "Install GO", Done: true, ID: 1},
		{Item: "Learn Go", Done: false, ID: 2},
		{Item: "Finish Todo-List", Done: false, ID: 3},
	},
}

func WriteFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(data))
	return err
}

func loadData() {
	// implement this
	// check if data already exists
	file, err := os.ReadFile("./database/data.json")
	// if no data found
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			jsonData, err := json.Marshal(htmlData)
			if err != nil {
				println(err.Error())
				return
			}
			WriteFile("/database/data.json", jsonData)
		}
	}
	//if data already exists
	//decode the data
	err = json.Unmarshal(file, &htmlData)
	//update PageData
}

func saveData() {
	// implement this
	// get Done data

	// update PageData

	// endcode data

	// write data to file
}

func todo(w http.ResponseWriter, r *http.Request) { //handler
	tmpl.Execute(w, htmlData)
}

func toggle(w http.ResponseWriter, r *http.Request) { //handler
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body is invalid"))
		return
	}

	ids, ok := r.PostForm["id"]

	if !ok || len(ids) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id doesn't exist"))
		return
	}

	for i := 0; i < len(htmlData.Todos); i++ {
		if ids[0] == strconv.Itoa(htmlData.Todos[i].ID) {
			htmlData.Todos[i].Done = !htmlData.Todos[i].Done

			w.WriteHeader(http.StatusOK)

			if htmlData.Todos[i].Done {
				w.Write([]byte("true"))
			} else {
				w.Write([]byte("false"))
			}

			saveData()

			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("not found"))
}

func main() {
	loadData()

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	mux.HandleFunc("/toggle", toggle)
	//if js changes, call another mux.HandleFunc() to change the template contents

	log.Fatal(http.ListenAndServe(":9091", mux))
}
