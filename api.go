package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

type Student struct {
	Id int
	Name string
	Major string
}

var data = []Student {
	{2201234567, "John Doe", "Computer Science"},
	{2009876543, "Betty Powell", "Information Systems"},
	{2112343212, "Patrick Fowell", "Computer Science"},
	{1920212223, "Raymond Zimmerman", "Computer Science"},
	{2324252627, "Grace Small", "Information Systems"},
}

func students(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func student(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id, _ := strconv.Atoi(r.URL.Query().Get("Id"))

		for _, student_data := range data {
			if student_data.Id == id {
				var result, _ = json.Marshal(student_data)

				w.Write(result)
				return
			}
		}

		http.Error(w, "", http.StatusNotFound)
	}
}

func add_student(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		r.ParseForm()
		
		id := r.PostForm.Get("Id")
		id_int, _ := strconv.Atoi(id)

		name := r.PostForm.Get("Name")
		major := r.PostForm.Get("Major")

		if id == "" || name == "" || major == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		} else {
			var new_student Student = Student{id_int, name, major}
			data = append(data, new_student)
			
			var result, _ = json.Marshal(new_student)

			w.Write(result)
			return
		}
	}
}

func update_student(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "UPDATE" {
		r.ParseForm()
		
		id := r.URL.Query().Get("Id")
		id_int, _ := strconv.Atoi(id)

		name := r.PostForm.Get("Name")
		major := r.PostForm.Get("Major")

		if id == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		} else {
			var new_student Student = Student{id_int, name, major}

			for i, student_data := range data {
				fmt.Println(student_data.Id)
				if student_data.Id == id_int {
					data[i] = Student{id_int, name, major}

					var result, _ = json.Marshal(new_student)

					w.Write(result)
					return
				}
			}

			http.Error(w, "", http.StatusNotFound)
		}
	}
}

func delete_student(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		r.ParseForm()
		
		id := r.URL.Query().Get("Id")
		id_int, _ := strconv.Atoi(id)

		if id == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		} else {
			for i, student_data := range data {
				if student_data.Id == id_int {
					data = slices.Delete(data, i, i+1)

					w.WriteHeader(200)
					return
				}
			}

			http.Error(w, "", http.StatusNotFound)
		}
	}
}

func main() {
	http.HandleFunc("/students", students)
	http.HandleFunc("/student", student)
	http.HandleFunc("/student/add", add_student)
	http.HandleFunc("/student/update", update_student)
	http.HandleFunc("/student/delete", delete_student)

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}