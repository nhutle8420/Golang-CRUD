package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	UserName    string       `json:"username"`
	PassWord    string       `json:"password"`
	Information *Information `json:"ionformation"`
}
type Information struct {
	Email string `json:"email"`
	Phone string `json:"Phone"`
}

var accounts []Account

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
func GetByUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	p := mux.Vars(r)
	for _, valude := range accounts {
		if valude.UserName == p["username"] {
			json.NewEncoder(w).Encode(valude)
			return
		}

	}
	f := "Not found"
	json.NewEncoder(w).Encode(f)
	return
}
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var accountt Account
	_ = json.NewDecoder(r.Body).Decode(&accountt)
	accounts = append(accounts, accountt)
	json.NewEncoder(w).Encode(accountt)
}
func UpdataAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	p := mux.Vars(r)
	for i, valude := range accounts {
		if valude.UserName == p["username"] {
			accounts = append(accounts[:i], accounts[i+1:]...)
			var account Account
			_ = json.NewDecoder(r.Body).Decode(&account)
			account.UserName = p["username"]
			accounts = append(accounts, account)
			json.NewEncoder(w).Encode(account)
			return
		}
	}

}
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application")
	p := mux.Vars(r)
	for i, valude := range accounts {
		if valude.UserName == p["username"] {
			accounts = append(accounts[:i], accounts[i+1:]...)
			t := "Delete Succesfull"
			json.NewEncoder(w).Encode(t)
			return
		}

	}
	t := "Delete Faile"
	json.NewEncoder(w).Encode(t)
	return

}
func main() {
	r := mux.NewRouter()
	accounts = append(accounts, Account{UserName: "admin", PassWord: "admin2", Information: &Information{Email: "admin@gmail.com", Phone: "098312312412"}})
	r.HandleFunc("/Accounts", getAllBooks).Methods("GET")
	r.HandleFunc("/Account/{username}", GetByUserName).Methods("GET")
	r.HandleFunc("/CreateAccount", CreateAccount).Methods("POST")
	r.HandleFunc("/UpdataAccount/{username}", UpdataAccount).Methods("PUT")
	r.HandleFunc("/Delete/{username}", DeleteAccount).Methods("DeLete")

	fmt.Printf("Listen To Server in Port 8088 !")
	log.Fatal(http.ListenAndServe(":8088", r))
}
