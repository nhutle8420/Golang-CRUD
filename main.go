package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id       string  `json:"id"`
	NameBook string  `json:"namebook"`
	Title    string  `json:"title"`
	Author   *Author `json:"author"`
}
type Author struct {
	FistName string `json:"fistname"`
	LastName string `json:"lastname"`
}

var Listbooks []Book

func getAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(Listbooks)
}
func getBookID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	p := mux.Vars(r)
	for _, valude := range Listbooks {
		if valude.Id == p["id"] {
			json.NewEncoder(w).Encode(valude)
			return
		}

	}
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(1000000000))
	Listbooks = append(Listbooks, book)
	json.NewEncoder(w).Encode(book)
}
func UpdataBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	p := mux.Vars(r)
	for i, valude := range Listbooks {
		if valude.Id == p["id"] {
			Listbooks = append(Listbooks[:i], Listbooks[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = p["id"]
			Listbooks = append(Listbooks, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	p := mux.Vars(r)
	for i, valude := range Listbooks {
		if valude.Id == p["id"] {
			Listbooks = append(Listbooks[:i], Listbooks[i+1:]...)
			break
		}
	}

}
func main() {
	r := mux.NewRouter()
	Listbooks = append(Listbooks, Book{Id: "1", NameBook: "Conna", Title: "Conan chapter 12", Author: &Author{FistName: "asdv", LastName: "123asd"}})
	Listbooks = append(Listbooks, Book{Id: "2", NameBook: "Doaremon", Title: "Doraemon chapter 9", Author: &Author{FistName: "hadasdsd", LastName: "dsad23asd"}})
	r.HandleFunc("/books", getAllBook).Methods("GET")
	r.HandleFunc("/books/{id}", getBookID).Methods("GET")
	r.HandleFunc("/createBook", createBook).Methods("POST")
	r.HandleFunc("/upBooks/{id}", UpdataBook).Methods("PUT")
	r.HandleFunc("/deleteBook/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Listen to Server at Port 9081")
	log.Fatal(http.ListenAndServe(":9081", r))
}
