package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Tutor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}

var (
	students []Student
	tutors   []Tutor
	mu       sync.Mutex
	nextID   = 1
)

func addStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	// basic error handling to return relevant httpCodes
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	s.ID = nextID
	nextID++
	students = append(students, s)
	mu.Unlock()
	writeJSON(w, s)
}

func addTutor(w http.ResponseWriter, r *http.Request) {
	var t Tutor
	// basic error handling to return relevant httpCodes
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	t.ID = nextID
	nextID++
	tutors = append(tutors, t)
	mu.Unlock()
	writeJSON(w, t)
}

func matchTutor(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	// basic error handling to return relevant httpCodes
	if len(students) == 0 || len(tutors) == 0 {
		http.Error(w, "no students or tutors available", http.StatusNotFound)
		return
	}
	resp := map[string]any{
		"student": students[len(students)-1],
		"tutor":   tutors[len(tutors)-1],
	}
	writeJSON(w, resp)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	http.HandleFunc("/students", addStudent) // POST
	http.HandleFunc("/tutors", addTutor)     // POST
	http.HandleFunc("/match", matchTutor)    // GET
	http.HandleFunc("/health", health)       // GET

	log.Println("server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
